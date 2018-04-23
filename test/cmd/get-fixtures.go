package main

import (
	"flag"
	"strings"

	"github.com/timvaillancourt/go-mongodb-replset/test"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mongodbUri  = flag.String("uri", "mongodb://localhost:27017", "mongodb server uri")
	fixturesDir = flag.String("dir", "./fixtures", "path to 'fixtures' directory")
	commands    = []string{
		"replSetGetConfig",
		"replSetGetStatus",
	}
)

func serverVersion(session *mgo.Session) (string, error) {
	buildInfo, err := session.BuildInfo()
	if err != nil {
		return "", err
	}
	if strings.Contains(buildInfo.Version, "-") {
		version := strings.SplitN(buildInfo.Version, "-", 2)
		return version[0], nil
	}
	return buildInfo.Version, nil
}

func main() {
	flag.Parse()

	session, err := mgo.Dial(*mongodbUri)
	if err != nil {
		panic(err)
	}

	version, err := serverVersion(session)
	if err != nil {
		panic(err)
	}

	for _, command := range commands {
		var data bson.Raw
		err = session.DB("admin").Run(bson.D{{command, "1"}}, &data)
		if err != nil {
			panic(err)
		}

		err = test.WriteFixture(*fixturesDir, version, command, data.Data)
		if err != nil {
			panic(err)
		}
	}
}

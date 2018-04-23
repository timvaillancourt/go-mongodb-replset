package main

import (
	"flag"
	"strings"

	"github.com/timvaillancourt/go-mongodb-replset/fixtures"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	mongodbUri  = flag.String("uri", "mongodb://localhost:27017", "mongodb server uri")
	baseDirName = flag.String("outDir", "./fixtures", "output directory for bson fixtures")
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
	version := strings.SplitN(buildInfo.Version, "-", 2)
	return version[0], nil
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

		err = fixtures.WriteFixture(version, command, data.Data)
		if err != nil {
			panic(err)
		}
	}
}

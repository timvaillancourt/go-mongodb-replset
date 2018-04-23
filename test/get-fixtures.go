package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

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

	versionDir := filepath.Join(*baseDirName, version)
	err = os.Mkdir(versionDir, 0755)
	if err != nil {
		panic(err)
	}

	for _, command := range commands {
		var data bson.Raw
		err = session.DB("admin").Run(bson.D{{command, "1"}}, &data)
		if err != nil {
			panic(err)
		}

		bsonData, err := bson.Marshal(data)
		if err != nil {
			panic(err)
		}

		fileName := versionDir + "/" + command + ".bson"
		err = ioutil.WriteFile(fileName, bsonData, 0640)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Wrote: %s\n", fileName)
	}
}

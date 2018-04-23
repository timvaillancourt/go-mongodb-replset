package test

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/mgo.v2/bson"
)

const fixturesDir = "./fixtures"

func LoadFixture(version, command string, out interface{}) error {
	filePath := filepath.Join(fixturesDir, version, command+".bson")
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return bson.Unmarshal(bytes, out)
}

func WriteFixture(version, command string, data []byte) error {
	versionDir := filepath.Join(fixturesDir, version)
	if _, err := os.Stat(versionDir); os.IsNotExist(err) {
		err = os.Mkdir(versionDir, 0755)
		if err != nil {
			return err
		}
	}
	filePath := filepath.Join(fixturesDir, version, command+".bson")
	return ioutil.WriteFile(filePath, data, 0640)
}

func FixtureVersions() []string {
	var versions []string
	subdirs, err := ioutil.ReadDir(fixturesDir)
	if err != nil {
		return versions
	}
	for _, subdir := range subdirs {
		versions = append(versions, subdir.Name())
	}
	return versions
}

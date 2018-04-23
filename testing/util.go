package testing

import (
	"io/ioutil"

	"gopkg.in/mgo.v2/bson"
)

func LoadFixture(filePath string, out interface{}) error {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	return bson.Unmarshal(bytes, out)
}

func WriteFixture(filePath string, data []byte) error {
	return ioutil.WriteFile(filePath, data, 0640)
}

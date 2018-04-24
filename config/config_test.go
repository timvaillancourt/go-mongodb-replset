package config

import (
	"testing"

	mongodb_fixtures "github.com/timvaillancourt/go-mongodb-fixtures"
)

var configCommand = "replSetGetConfig"

func getConfigFixture(version string) (*Config, error) {
	s := &Config{}
	err := mongodb_fixtures.LoadFixture(version, configCommand, s)
	return s, err
}

//func main() {
//for _, fixtureVersion := range mongodb_fixtures.FixtureVersions() {
//fixture, err := getStatusFixture(fixtureVersion)
//if err != nil {
//	fmt.Printf("Error loading fixture for %s: %s\n", fixtureVersion, err)
//}

func TestToJSON(t *testing.T) {
	output := `{
	"_id": "test",
	"version": 1,
	"members": [
		{
			"_id": 0,
			"host": "localhost:27017",
			"arbiterOnly": false,
			"buildIndexes": true,
			"hidden": false,
			"priority": 1,
			"tags": {
				"test": "123"
			},
			"slaveDelay": 0,
			"votes": 1
		}
	]
}`

	c := &Config{
		Name:    "test",
		Version: 1,
		Members: []*Member{
			&Member{
				Id:           0,
				Host:         "localhost:27017",
				BuildIndexes: true,
				Priority:     1,
				Tags: &ReplsetTags{
					"test": "123",
				},
				Votes: 1,
			},
		},
	}

	str, err := c.ToJSON()
	if err != nil {
		t.Errorf("Error running config.ToJSON(): %s", err)
	}
	if string(str) == "" {
		t.Errorf("config.ToJSON() returned empty string")
	}
	if string(str) != output {
		t.Error("config.ToJSON() does not match expected output")
	}
}

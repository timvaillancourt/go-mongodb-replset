package config

import (
	"testing"

	mongodb_fixtures "github.com/timvaillancourt/go-mongodb-fixtures"
)

var (
	configCommand = "replSetGetConfig"
	testConfig    = &Config{
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
	addMember = &Member{
		Id:           1,
		Host:         "localhost:27018",
		BuildIndexes: true,
		Priority:     1,
		Votes:        1,
	}
)

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

	str, err := testConfig.ToJSON()
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

func TestGetMember(t *testing.T) {
	member := testConfig.GetMember("localhost:27017")
	if member.Host != "localhost:27017" {
		t.Error("config.GetMember() returned wrong 'host'")
	}
}

func TestAddMember(t *testing.T) {
	testConfig.AddMember(addMember)
	member := testConfig.GetMember(addMember.Host)
	if member.Host != addMember.Host || member.Id != addMember.Id {
		t.Error("config.AddMember() failed, .GetMember() after add returns wrong data")
	}
}

func TestHasMember(t *testing.T) {
	if !testConfig.HasMember(addMember.Host) {
		t.Error("config.HasMember() did not return true")
	}
}

func TestIncrVersion(t *testing.T) {
	version := testConfig.Version
	testConfig.IncrVersion()
	if testConfig.Version != (version + 1) {
		t.Errorf("config.IncrVersion() did not increment the version from %d to %d", version, (version + 1))
	}
}

func TestRemoveMember(t *testing.T) {
	testConfig.RemoveMember(addMember)
	if testConfig.HasMember(addMember.Host) {
		t.Errorf("config.RemoveMember() did not succeed, %s is still in config", addMember.Host)
	}
}

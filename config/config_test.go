package config

import (
	"testing"

	mongodb_fixtures "github.com/timvaillancourt/go-mongodb-fixtures"
)

var (
	testConfig = &Config{
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

func getConfigFixture(t *testing.T, version string) *Config {
	rsgc := &ReplSetGetConfig{}
	err := mongodb_fixtures.Load(version, ConfigCommand, rsgc)
	if err != nil {
		t.Errorf("Error loading fixture for %s: %s\n", version, err)
	}
	return rsgc.Config
}

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

func TestFixtures(t *testing.T) {
	for _, version := range mongodb_fixtures.Versions() {
		t.Logf("Testing fixtures for mongodb version %s", version)

		c := getConfigFixture(t, version)
		if c.Name == "" {
			t.Error("config.Name cannot be an empty string")
		}
		if c.Version < 1 {
			t.Errorf("config.Version must be 1 or greater, got %d", c.Version)
		}
		if len(c.Members) < 1 {
			t.Errorf("config.Members must be greater than zero, got %d", len(c.Members))
			continue
		}

		member := c.GetMemberId(0)
		if member == nil {
			t.Errorf("config.GetMemberId(0) for %s returned nil", version)
		} else if member.Id != 0 {
			t.Errorf("config.GetMemberId(0) for %s returned a non-zero id", version)
		}

		getMember := c.GetMember(member.Host)
		if getMember == nil {
			t.Errorf("config.GetMember(\"%s\") for %s returned nil", member.Host, version)
		} else if getMember.Host != member.Host {
			t.Errorf("config.GetMember(\"%s\") for %s returned incorrect host", member.Host, version)
		}

		if !c.HasMember(member.Host) {
			t.Errorf("config.HasMember(\"%s\") for %s returned false", member.Host, version)
		}
	}
}

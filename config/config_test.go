package config

import (
	"testing"

	mongodb_fixtures "github.com/timvaillancourt/go-mongodb-fixtures"
)

var (
	testReplsetName = "test"
	testConfig      = &Config{
		Name:    testReplsetName,
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
)

func getConfigFixture(t *testing.T, version string) *Config {
	rsgc := &ReplSetGetConfig{}
	err := mongodb_fixtures.Load(version, ConfigCommand, rsgc)
	if err != nil {
		t.Errorf("Error loading fixture for %s: %s\n", version, err)
	}
	return rsgc.Config
}

func TestNewConfig(t *testing.T) {
	config := NewConfig(testReplsetName)
	if config.Name != testReplsetName {
		t.Errorf("config.NewConfig(\"%s\") returned a struct with 'Name' equal to %v, not %s", testReplsetName, config.Name, testReplsetName)
	}
	if config.Version != 1 {
		t.Errorf("config.NewConfig(\"%s\") returned a struct with 'Version' not equal to 1: %v", testReplsetName, config.Version)
	}
	if len(config.Members) > 0 {
		t.Errorf("config.NewConfig(\"%s\") returned a struct with a non-empty 'Members': %v", testReplsetName, config.Members)
	}
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

func TestIncrVersion(t *testing.T) {
	version := testConfig.Version
	testConfig.IncrVersion()
	if testConfig.Version != (version + 1) {
		t.Errorf("config.IncrVersion() did not increment the version from %d to %d", version, (version + 1))
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

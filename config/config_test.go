package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	assert.NoErrorf(t, err, "Error loading fixture for %s", version)
	return rsgc.Config
}

func TestNewConfig(t *testing.T) {
	config := NewConfig(testReplsetName)
	assert.Equal(t, testReplsetName, config.Name, "config.NewConfig(\"%s\") returned incorrect struct")
	assert.Equal(t, config.Version, 1, "config.NewConfig(\"%s\") returned incorrect struct")
	assert.Len(t, config.Members, 0, "config.NewConfig(\"%s\") returned a struct with non-empty 'Members'")
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

	bytes, err := testConfig.ToJSON()
	assert.NoError(t, err, "Error running config.ToJSON()")
	assert.NotZero(t, bytes, "config.ToJSON() returned zero bytes of json")
	assert.Equal(t, output, string(bytes), "config.ToJSON() does not match expected output")
}

func TestIncrVersion(t *testing.T) {
	version := testConfig.Version
	testConfig.IncrVersion()
	assert.Equal(t, version+1, testConfig.Version, "config.IncrVersion() did not increment")
}

func TestFixtures(t *testing.T) {
	for _, version := range mongodb_fixtures.Versions() {
		t.Logf("Testing fixtures for '%s' command on mongodb version %s", ConfigCommand, version)

		c := getConfigFixture(t, version)
		assert.NotEmpty(t, c.Name, "config.Name cannot be an empty string")
		assert.Truef(t, c.Version > 0, "config.Version must be 1 or greater, got %d", c.Version)
		assert.NotEmpty(t, c.Members, "config.Members must have at least 1 member")

		member := c.GetMemberId(0)
		assert.NotNil(t, member, "config.GetMemberId(0) returned nil")
		assert.Equal(t, 0, member.Id, "config.GetMemberId(0) returned a non-zero id")

		getMember := c.GetMember(member.Host)
		assert.NotNilf(t, getMember, "config.GetMember(\"%s\") returned nil", member.Host)
		assert.Equalf(t, member.Host, getMember.Host, "config.GetMember(\"%s\") returned incorrect host", member.Host)

		assert.Truef(t, c.HasMember(member.Host), "config.HasMember(\"%s\") returned false", member.Host)
	}
}

package status

import (
	"testing"

	"github.com/stretchr/testify/assert"
	mongodb_fixtures "github.com/timvaillancourt/go-mongodb-fixtures"
)

var (
	testStatus = &Status{
		Set:     "test",
		MyState: MemberStatePrimary,
		Ok:      1,
		Members: []*Member{
			testMember,
			testMemberSecondary,
		},
	}
)

func getStatusFixture(t *testing.T, version string) *Status {
	s := &Status{}
	err := mongodb_fixtures.Load(version, StatusCommand, s)
	assert.NoErrorf(t, err, "Error loading fixture for %s", version)
	return s
}

func TestToJSON(t *testing.T) {
	output := `{
	"set": "test",
	"date": "0001-01-01T00:00:00Z",
	"myState": 1,
	"members": [
		{
			"_id": 0,
			"name": "localhost:27017",
			"health": 1,
			"state": 1,
			"stateStr": "PRIMARY",
			"uptime": 1,
			"optime": {
				"ts": 0,
				"t": 0
			},
			"optimeDate": "0001-01-01T00:00:00Z",
			"configVersion": 0,
			"electionDate": "0001-01-01T00:00:00Z",
			"optimeDurableDate": "0001-01-01T00:00:00Z",
			"lastHeartbeat": "0001-01-01T00:00:00Z",
			"lastHeartbeatRecv": "0001-01-01T00:00:00Z",
			"self": true
		},
		{
			"_id": 1,
			"name": "localhost:27018",
			"health": 1,
			"state": 2,
			"stateStr": "SECONDARY",
			"uptime": 1,
			"optime": {
				"ts": 0,
				"t": 0
			},
			"optimeDate": "0001-01-01T00:00:00Z",
			"configVersion": 0,
			"electionDate": "0001-01-01T00:00:00Z",
			"optimeDurableDate": "0001-01-01T00:00:00Z",
			"lastHeartbeat": "0001-01-01T00:00:00Z",
			"lastHeartbeatRecv": "0001-01-01T00:00:00Z"
		}
	],
	"ok": 1
}`

	bytes, err := testStatus.ToJSON()
	assert.NoError(t, err, "Error running status.ToJSON()")
	assert.NotEmpty(t, bytes, "status.ToJSON() returned zero bytes of json")
	assert.Equal(t, output, string(bytes), "status.ToJSON() does not match expected output")
}

func TestFixtures(t *testing.T) {
	for _, version := range mongodb_fixtures.Versions() {
		t.Logf("Testing fixtures for '%s' command on mongodb version %s", StatusCommand, version)

		s := getStatusFixture(t, version)
		assert.NotEmpty(t, s.Members, "status.Members must return 1 or more members")

		self := s.GetSelf()
		assert.NotNil(t, self, "status.GetSelf() returned nil")

		if mongodb_fixtures.IsVersionMatch(version, ">= 3.2") {
			assert.NotNil(t, self.Optime, "status.Optime is nil")
		}

		assert.NotNilf(t, s.GetMemberId(self.Id), "status.GetMemberId(%d) returned nil", self.Id)
		assert.NotNilf(t, s.GetMember(self.Name), "status.GetMember(\"%s\") returned nil", self.Name)

		primary := s.Primary()
		assert.NotNil(t, primary, "status.Primary() returned nil")
		assert.Equal(t, MemberStatePrimary, primary.State, "status.Primary() did not return a Primary!")
	}
}

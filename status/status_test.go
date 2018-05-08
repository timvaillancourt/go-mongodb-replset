package status

import (
	"testing"

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
	if err != nil {
		t.Errorf("Error loading fixture for %s: %s", version, err)
	}
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

	str, err := testStatus.ToJSON()
	if err != nil {
		t.Errorf("Error running status.ToJSON(): %s", err)
	}
	if string(str) == "" {
		t.Error("status.ToJSON() returned empty string")
	}
	if string(str) != output {
		t.Error("status.ToJSON() does not match expected output")
	}
}

func TestFixtures(t *testing.T) {
	for _, version := range mongodb_fixtures.Versions() {
		t.Logf("Testing fixtures for '%s' command on mongodb version %s", StatusCommand, version)

		s := getStatusFixture(t, version)

		if len(s.Members) < 1 {
			t.Errorf("Error for %s: status.Members must return 1 or more members!", version)
			continue
		}

		self := s.GetSelf()
		if self == nil {
			t.Errorf("Error for %s: status.GetSelf() returned nil!", version)
			continue
		}

		if mongodb_fixtures.IsVersionMatch(version, ">= 3.2") {
			if self.Optime == nil {
				t.Errorf("Error for %s: status.Optime is nil!", version)
			}
		}

		if s.GetMemberId(self.Id) == nil {
			t.Errorf("Error for %s: status.GetMemberId(%d) returned nil!", version, self.Id)
		}

		if s.GetMember(self.Name) == nil {
			t.Errorf("Error for %s: status.GetMember(\"%s\") returned nil!", version, self.Name)
		}

		primary := s.Primary()
		if primary == nil {
			t.Errorf("Error for %s: status.Primary() returned nil!", version)
		} else if primary.State != MemberStatePrimary {
			t.Errorf("Error for %s: status.Primary() did not return a Primary!", version)
		}
	}
}

package status

import (
	"testing"

	fixtures "github.com/timvaillancourt/go-mongodb-fixtures"
)

func getStatusFixture(t *testing.T, version string) *Status {
	s := &Status{}
	err := fixtures.LoadFixture(version, statusCommand, s)
	if err != nil {
		t.Errorf("Error loading fixture for %s: %s", version, err)
	}
	return s
}

func TestGetMembers(t *testing.T) {
	for _, version := range fixtures.FixtureVersions() {
		s := getStatusFixture(t, version)
		t.Logf("Testing status.Members for %s", version)
		if len(s.Members) < 1 {
			t.Errorf("Error for %s: status.Members must return 1 or more members!", version)
		}
	}
}

func TestGetSelf(t *testing.T) {
	for _, version := range fixtures.FixtureVersions() {
		s := getStatusFixture(t, version)
		t.Logf("Testing status.GetSelf() for %s", version)
		if s.GetSelf() == nil {
			t.Errorf("Error for %s: status.GetSelf() returned nil!", version)
		}
	}
}

func TestGetMemberId0(t *testing.T) {
	for _, version := range fixtures.FixtureVersions() {
		s := getStatusFixture(t, version)
		t.Logf("Testing status.GetMemberId(0) for %s", version)
		if s.GetMemberId(0) == nil {
			t.Errorf("Error for %s: status.GetMemberId(0) returned nil!", version)
		}
	}
}

func TestPrimary(t *testing.T) {
	for _, version := range fixtures.FixtureVersions() {
		s := getStatusFixture(t, version)
		t.Logf("Testing status.Primary() for %s", version)
		primary := s.Primary()
		if primary == nil {
			t.Errorf("Error for %s: status.Primary() returned nil!", version)
		}
		if primary.State != MemberStatePrimary {
			t.Errorf("Error for %s: status.Primary() did not return a Primary!", version)
		}
	}
}

func TestJSONOutput(t *testing.T) {
	var output = `{
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
			"lastHeartbeatRecv": "0001-01-01T00:00:00Z"
		}
	],
	"ok": 1
}`

	s := &Status{
		Set:     "test",
		MyState: MemberStatePrimary,
		Ok:      1,
		Members: []*Member{
			&Member{
				Id:       0,
				Name:     "localhost:27017",
				Health:   MemberHealthUp,
				State:    MemberStatePrimary,
				StateStr: "PRIMARY",
				Optime:   &Optime{},
				Uptime:   1,
			},
		},
	}
	str, err := s.ToJSON()
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

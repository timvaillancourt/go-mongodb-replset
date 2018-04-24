package status

import (
	"testing"

	mongodb_fixtures "github.com/timvaillancourt/go-mongodb-fixtures"
)

var (
	testMember = &Member{
		Id:       0,
		Name:     "localhost:27017",
		Health:   MemberHealthUp,
		State:    MemberStatePrimary,
		StateStr: "PRIMARY",
		Optime:   &Optime{},
		Uptime:   1,
		Self:     true,
	}
	testMemberSecondary = &Member{
		Id:       1,
		Name:     "localhost:27018",
		Health:   MemberHealthUp,
		State:    MemberStateSecondary,
		StateStr: "SECONDARY",
		Optime:   &Optime{},
		Uptime:   1,
	}
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
	err := mongodb_fixtures.LoadFixture(version, statusCommand, s)
	if err != nil {
		t.Errorf("Error loading fixture for %s: %s", version, err)
	}
	return s
}

func TestGetSelf(t *testing.T) {
	if testStatus.GetSelf() == nil {
		t.Error("status.GetSelf() returned nil")
	}
}

func TestGetMemberId(t *testing.T) {
	if testStatus.GetMemberId(testMember.Id) == nil {
		t.Errorf("status.GetMemberId(%d) returned nil", testMember.Id)
	}
}

func TestGetMember(t *testing.T) {
	if testStatus.GetMember(testMember.Name) == nil {
		t.Errorf("status.GetMember(\"%s\") returned nil", testMember.Name)
	}
}

func TestPrimary(t *testing.T) {
	primary := testStatus.Primary()
	if primary == nil {
		t.Error("status.Primary() returned nil")
	} else if primary.State != MemberStatePrimary {
		t.Error("status.Primary() returned member with non-primary state")
	} else if primary.Name != testMember.Name || primary.Id != testMember.Id {
		t.Error("status.Primary() did not return the primary")
	}
}

func TestSecondary(t *testing.T) {
	secondaries := testStatus.Secondaries()
	if len(secondaries) != 1 {
		t.Error("status.Secondary() returned zero secondaries")
	} else if secondaries[0].State != MemberStateSecondary {
		t.Error("status.Secondary() returned member with non-secondary state")
	} else if secondaries[0].Name != testMemberSecondary.Name || secondaries[0].Id != testMemberSecondary.Id {
		t.Error("status.Secondary() did not return a secondary")
	}
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
	for _, version := range mongodb_fixtures.FixtureVersions() {
		t.Logf("Testing fixtures for mongodb version %s", version)

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

		if self.Optime == nil {
			t.Errorf("Error for %s: status.Optime is nil!", version)
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

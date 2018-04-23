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

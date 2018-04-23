package status

import (
	"testing"

	"github.com/timvaillancourt/go-mongodb-replset/test"
)

const fixturesDir = "../test/fixtures"

func getStatusFixture(version string) (*Status, error) {
	s := &Status{}
	err := test.LoadFixture(fixturesDir, version, statusCommand, s)
	return s, err
}

func TestGetSelf(t *testing.T) {
	for _, version := range test.FixtureVersions(fixturesDir) {
		s, err := getStatusFixture(version)
		t.Logf("Testing .GetSelf() for %s", version)
		if err != nil {
			t.Errorf("Error loading fixture for %s: %s", version, err)
		}
		if s.GetSelf() == nil {
			t.Errorf("Error for %s: .GetSelf() returned nil!", version)
		}
	}
}

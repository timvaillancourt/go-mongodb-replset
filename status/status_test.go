package status

import (
	"testing"

	"github.com/timvaillancourt/go-mongodb-replset/test"
)

func getStatusFixture(version string) (*Status, error) {
	s := &Status{}
	err := test.LoadFixture(version, statusCommand, s)
	return s, err
}

func TestGetSelf(t *testing.T) {
	for _, fixtureVersion := range test.Versions() {
		fixture, err := getStatusFixture(fixtureVersion)
		if err != nil {
			t.Errorf("Error loading fixture for %s: %s", fixtureVersion, err)
		}
		if fixture.GetSelf() == nil {
			t.Errorf("Error for %s: .GetSelf() returned nil!")
		}
	}
}

package status

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
)

func TestMemberStateString(t *testing.T) {
	assert.Equal(t, testMember.StateStr, testMember.State.String(), "member.State.String() has unexpected value")
	assert.Equal(t, testMemberSecondary.StateStr, testMemberSecondary.State.String(), "member.State.String() has unexpected value")
}

func TestGetSelf(t *testing.T) {
	assert.NotNil(t, testStatus.GetSelf(), "status.GetSelf() returned nil")
}

func TestGetMemberId(t *testing.T) {
	assert.NotNilf(t, testStatus.GetMemberId(testMember.Id), "status.GetMemberId(%d) returned nil", testMember.Id)
}

func TestGetMember(t *testing.T) {
	assert.NotNilf(t, testStatus.GetMember(testMember.Name), "status.GetMember(\"%s\") returned nil", testMember.Name)
}

func TestPrimary(t *testing.T) {
	primary := testStatus.Primary()
	assert.NotNil(t, primary, "status.Primary() returned nil")
	assert.Equal(t, MemberStatePrimary, primary.State, "status.Primary() returned member with non-primary state")
	assert.Equal(t, testMember.Name, primary.Name, "status.Primary() did not return the primary")
	assert.Equal(t, testMember.Id, primary.Id, "status.Primary() did not return the primary")
}

func TestSecondary(t *testing.T) {
	secondaries := testStatus.Secondaries()
	assert.Len(t, secondaries, 1, "status.Secondary() returned zero or more than one secondaries")
	assert.Equal(t, MemberStateSecondary, secondaries[0].State, "status.Secondary() returned member with non-secondary state")
	assert.Equal(t, testMemberSecondary.Name, secondaries[0].Name, "status.Secondary() did not return a secondary")
	assert.Equal(t, testMemberSecondary.Id, secondaries[0].Id, "status.Secondary() did not return a secondary")
}

func TestGetMembersByState(t *testing.T) {
	members := testStatus.GetMembersByState(MemberStatePrimary, 0)
	assert.Lenf(t, members, 1, "status.GetMembersByState(\"%s\", 0) returned %d members, not 1", MemberStatePrimary, len(members))
	members = testStatus.GetMembersByState(MemberStateUnknown, 0)
	assert.Lenf(t, members, 0, "status.GetMembersByState(\"%s\", 0) returned %d members, not 0", MemberStateUnknown, len(members))
}

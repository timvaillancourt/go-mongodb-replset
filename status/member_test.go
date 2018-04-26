package status

import "testing"

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
	if testMember.State.String() != "PRIMARY" {
		t.Errorf("member.State.String() returned %v, not %s", testMember.State.String(), "PRIMARY")
	}
	if testMemberSecondary.State.String() != "SECONDARY" {
		t.Errorf("member.State.String() returned %v, not %s", testMember.State.String(), "SECONDARY")
	}
}

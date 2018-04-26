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
	if testMember.State.String() != testMember.StateStr {
		t.Errorf("member.State.String() returned %v, not %s", testMember.State.String(), testMember.StateStr)
	}
	if testMemberSecondary.State.String() != testMemberSecondary.StateStr {
		t.Errorf("member.State.String() returned %v, not %s", testMember.State.String(), testMemberSecondary.StateStr)
	}
}

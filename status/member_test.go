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
		t.Error("status.Secondary() returned zero or more than one secondaries")
	} else if secondaries[0].State != MemberStateSecondary {
		t.Error("status.Secondary() returned member with non-secondary state")
	} else if secondaries[0].Name != testMemberSecondary.Name || secondaries[0].Id != testMemberSecondary.Id {
		t.Error("status.Secondary() did not return a secondary")
	}
}

func TestGetMembersByState(t *testing.T) {
	members := testStatus.GetMembersByState(MemberStatePrimary, 0)
	if len(members) != 1 {
		t.Errorf("status.GetMembersByState(\"%s\", 0) returned %d members, not 1", MemberStatePrimary, len(members))
	}
	members = testStatus.GetMembersByState(MemberStateUnknown, 0)
	if len(members) != 0 {
		t.Errorf("status.GetMembersByState(\"%s\", 0) returned %d members, not 0", MemberStateUnknown, len(members))
	}
}

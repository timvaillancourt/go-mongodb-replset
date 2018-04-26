package config

import (
	"testing"
)

var (
	testMember = &Member{
		Id:           1,
		Host:         "localhost:27018",
		BuildIndexes: true,
		Priority:     1,
		Votes:        1,
	}
)

func TestNewMember(t *testing.T) {
	member := NewMember("test:123456")
	if member.Host != "test:123456" {
		t.Errorf("config.NewMember(\"test:123456\") returned a struct with 'Host' not equal to test:123456: %v", member.Host)
	}
	if member.BuildIndexes != true {
		t.Error("config.NewMember(\"test:123456\") returned a struct with 'BuildIndexes' set to false")
	}
	if member.Priority != 1 {
		t.Errorf("config.NewMember(\"test:123456\") returned a struct with 'Priority' not equal to 1: %v", member.Priority)
	}
	if member.Votes != 1 {
		t.Errorf("config.NewMember(\"test:123456\") returned a struct with 'Votes' not equal to 1: %v", member.Votes)
	}
}

func TestGetMemberMaxId(t *testing.T) {
	config := testConfig
	member := NewMember("test:123456")
	member.Id = 99
	config.AddMember(member)
	if config.getMemberMaxId() != 99 {
		t.Errorf("config.getMemberMaxId() returned an value not equal to 99: %v", config.getMemberMaxId())
	}
}

func TestGetMember(t *testing.T) {
	member := testConfig.GetMember("localhost:27017")
	if member.Host != "localhost:27017" {
		t.Error("config.GetMember() returned wrong 'host'")
	}
}

func TestAddMember(t *testing.T) {
	testConfig.AddMember(testMember)
	member := testConfig.GetMember(testMember.Host)
	if member.Host != testMember.Host || member.Id != testMember.Id {
		t.Error("config.AddMember() failed, .GetMember() after add returns wrong data")
	}
}

func TestHasMember(t *testing.T) {
	if !testConfig.HasMember(testMember.Host) {
		t.Error("config.HasMember() did not return true")
	}
}

func TestRemoveMember(t *testing.T) {
	testConfig.RemoveMember(testMember)
	if testConfig.HasMember(testMember.Host) {
		t.Errorf("config.RemoveMember() did not succeed, %s is still in config", testMember.Host)
	}
}

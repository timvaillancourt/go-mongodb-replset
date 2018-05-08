package config

import (
	"testing"
)

func TestReplsetTagsHasKey(t *testing.T) {
	if !testMember.Tags.HasKey("test") {
		t.Errorf("member.Tags.HasKey() returned false for %v", "test")
	}
	if testMember.Tags.HasKey("does not exist") {
		t.Error("member.Tags.HasKey() returned true for missing key")
	}
}

func TestReplsetTagsHasMatch(t *testing.T) {
	if !testMember.Tags.HasMatch("test", "123456") {
		t.Errorf("member.Tags.HasMatch() returned false for %v=%v", "test", "123456")
	}
	if testMember.Tags.HasMatch("test", "1234567") {
		t.Error("member.Tags.HasMatch() returned true for missing match")
	}

}
func TestReplsetTagsGet(t *testing.T) {
	if testMember.Tags.Get("test") != "123456" {
		t.Errorf("member.Tags.GetTagValue(\"test\") returned false for %v=%v", "test", "123456")
	}
}

func TestReplsetTagsAdd(t *testing.T) {
	if testMember.Tags.HasKey("testaddtag") {
		t.Error("members.Tags should not contain 'testaddtag' yet")
	}
	testMember.Tags.Add("testaddtag", "123")
	if !testMember.Tags.HasKey("testaddtag") {
		t.Error("member.Tags does not contain 'tataddtag'")
	}
}

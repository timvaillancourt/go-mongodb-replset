package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplsetTagsHasKey(t *testing.T) {
	assert.True(t, testMemberAdd.Tags.HasKey("test"), "member.Tags.HasKey() returned false for existing key")
	assert.False(t, testMemberAdd.Tags.HasKey("does not exist"), "member.Tags.HasKey() returned true for missing key")
}

func TestReplsetTagsHasMatch(t *testing.T) {
	assert.Truef(t, testMemberAdd.Tags.HasMatch("test", "123456"), "member.Tags.HasMatch() returned false for %v=%v", "test", "123456")
	assert.False(t, testMemberAdd.Tags.HasMatch("test", "1234567"), "member.Tags.HasMatch() returned true for missing match")

}
func TestReplsetTagsGet(t *testing.T) {
	assert.Equal(t, "123456", testMemberAdd.Tags.Get("test"), "member.Tags.GetTagValue(\"test\") returned incorrect value")
}

func TestReplsetTagsAdd(t *testing.T) {
	assert.False(t, testMemberAdd.Tags.HasKey("testaddtag"), "members.Tags should not contain 'testaddtag' yet")
	testMemberAdd.Tags.Add("testaddtag", "123")
	assert.True(t, testMemberAdd.Tags.HasKey("testaddtag"), "member.Tags does not contain 'tataddtag'")
}

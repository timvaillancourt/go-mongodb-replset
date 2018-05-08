package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testMemberName = "test:123456"
	testMember     = &Member{}
	testMemberAdd  = &Member{
		Id:           1,
		Host:         "localhost:27018",
		BuildIndexes: true,
		Priority:     1,
		Votes:        1,
		Tags: &ReplsetTags{
			"test": "123456",
		},
	}
)

func TestNewMember(t *testing.T) {
	testMember = NewMember(testMemberName)
	assert.NotNil(t, testMember, "config.NewMember() returned nil")
	assert.Equal(t, testMemberName, testMember.Host, "config.NewMember(\"test:123456\") returned incorrect struct")
	assert.True(t, testMember.BuildIndexes, "config.NewMember(\"test:123456\") returned a struct with 'BuildIndexes' set to false")
	assert.Equal(t, 1, testMember.Priority, "config.NewMember(\"test:123456\") returned a struct with 'Priority' not equal to 1")
	assert.Equal(t, 1, testMember.Votes, "config.NewMember(\"test:123456\") returned a struct with 'Votes' not equal to 1")
}

func TestGetMemberMaxIdBeforeAdd(t *testing.T) {
	assert.Equal(t, 0, testConfig.getMemberMaxId(), "config.getMemberMaxId() returned an incorrect value")
}

func TestGetMember(t *testing.T) {
	member := testConfig.GetMember("localhost:27017")
	assert.NotNil(t, member, "config.GetMember() returned nil")
	assert.Equal(t, "localhost:27017", member.Host, "config.GetMember() returned wrong 'host'")
}

func TestAddMember(t *testing.T) {
	testConfig.AddMember(testMemberAdd)
	member := testConfig.GetMember(testMemberAdd.Host)
	assert.NotNil(t, member, "config.GetMember() after config.AddMember() failed")
	assert.Equal(t, testMemberAdd.Host, member.Host, "config.AddMember() failed, .GetMember() after add returns wrong data")
	assert.Equal(t, testMemberAdd.Id, member.Id, "config.AddMember() failed, .GetMember() after add returns wrong data")
}

func TestGetMemberMaxIdAfterAdd(t *testing.T) {
	assert.Equal(t, 1, testConfig.getMemberMaxId(), "config.getMemberMaxId() returned an incorrect value")
}

func TestHasMember(t *testing.T) {
	assert.True(t, testConfig.HasMember(testMemberAdd.Host), "config.HasMember() did not return true")
}

func TestRemoveMember(t *testing.T) {
	testConfig.RemoveMember(testMemberAdd)
	assert.Falsef(t, testConfig.HasMember(testMemberAdd.Host), "config.RemoveMember() did not succeed, %s is still in config", testMemberAdd.Host)
}

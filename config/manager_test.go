package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2"
)

var (
	testManager *ConfigManager
)

func TestNewManager(t *testing.T) {
	session, err := mgo.Dial("mongodb://localhost:65217")
	if err != nil {
		assert.FailNow(t, "cannot get session")
	}
	if session.Ping() != nil {
		assert.FailNow(t, "cannot ping session")
	}

	testManager = New(session)
	assert.Equal(t, session, testManager.session, "testManager.session is incorrect")
	assert.False(t, testManager.initiated, "testManager.initiated should be false")
	assert.Nil(t, testManager.config, "testManager.config should be nil")
}

func TestManagerGetNil(t *testing.T) {
	assert.Nil(t, testManager.Get(), "manager.Get() returned unexpected data")
}

func TestManagerLoad(t *testing.T) {
	err := testManager.Load()
	if err != nil {
		assert.FailNow(t, "manager.Load() returned error")
	}
	assert.True(t, testManager.initiated, "manager.initiated is not true")
	assert.NotNil(t, testManager.config, "manager.config is nil")
}

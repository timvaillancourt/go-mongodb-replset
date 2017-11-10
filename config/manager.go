package config

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ConfigManager struct {
	session   *mgo.Session
	config    *Config
	initiated bool
}

func New(session *mgo.Session) *ConfigManager {
	return &ConfigManager{
		session:   session,
		initiated: false,
	}
}

func (c *ConfigManager) Get() *Config {
	return c.config
}

func (c *ConfigManager) Set(config *Config) {
	c.config = config
}

func (c *ConfigManager) Load() error {
	resp := &ReplSetGetConfig{}
	err := c.session.Run(bson.D{{"replSetGetConfig", 1}}, resp)
	if err != nil {
		return err
	}
	if resp.Ok == 1 && resp.Config != nil {
		c.config = resp.Config
		c.initiated = true
	} else {
		return errors.New(resp.Errmsg)
	}
	return nil
}

func (c *ConfigManager) IsInitiated() bool {
	if c.initiated {
		return true
	}
	err := c.Load()
	if err != nil {
		return false
	}
	return true
}

func (c *ConfigManager) Initiate() error {
	if c.initiated {
		return nil
	}
	resp := &OkResponse{}
	err := c.session.Run(bson.D{{"replSetInitiate", c.config}}, resp)
	if err != nil {
		if err.Error() == "already initialized" {
			c.initiated = true
			return nil
		}
		return err
	}
	if resp.Ok == 1 {
		c.initiated = true
	}
	return nil
}

func (c *ConfigManager) Validate() error {
	if c.config.Name == "" {
		return ErrNoReplsetId
	}
	if len(c.config.Members) == 0 {
		return ErrNoReplsetMembers
	}
	return nil
}

func (c *ConfigManager) Save() error {
	err := c.Validate()
	if err != nil {
		return err
	}
	if c.IsInitiated() {
		resp := &OkResponse{}
		err = c.session.Run(bson.D{{"replSetReconfig", c}}, resp)
	} else {
		err = c.Initiate()
	}
	if err != nil {
		return err
	}
	return nil
}

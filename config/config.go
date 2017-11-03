package config

import (
	"errors"
	//"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	ErrNoReplsetId        = errors.New("replset config has no _id field!")
	ErrNoReplsetMembers   = errors.New("replset config has no members!")
	NoReplsetConfigString = "no replset config has been received"
)

type ReplsetTags map[string]string
type WriteConcern struct {
	WriteConcern string `bson:"w"`
	WriteTimeout int    `bson:"wtimeout"`
	Journal      bool   `bson:"j,omitempty"`
}

type OkResponse struct {
	Ok int `bson:"ok"`
}

type Settings struct {
	ChainingAllowed         bool                    `bson:"chainingAllowed,omitempty"`
	HeartbeatIntervalMillis int64                   `bson:"heartbeatIntervalMillis,omitempty"`
	HeartbeatTimeoutSecs    int                     `bson:"heartbeatTimeoutSecs,omitempty"`
	ElectionTimeoutMillis   int64                   `bson:"electionTimeoutMillis,omitempty"`
	CatchUpTimeoutMillis    int64                   `bson:"catchUpTimeoutMillis,omitempty"`
	GetLastErrorModes       map[string]*ReplsetTags `bson:"getLastErrorModes,omitempty"`
	GetLastErrorDefaults    *WriteConcern           `bson:"getLastErrorDefaults,omitempty"`
	ReplicaSetId            bson.ObjectId           `bson:"replicaSetId,omitempty"`
}

type Config struct {
	Name                               string    `bson:"_id"`
	Version                            int       `bson:"version"`
	Members                            []*Member `bson:"members"`
	Configsvr                          bool      `bson:"configsvr,omitempty"`
	ProtocolVersion                    int       `bson:"protocolVersion,omitempty"`
	Settings                           *Settings `bson:"settings,omitempty"`
	WriteConcernMajorityJournalDefault bool      `bson:"writeConcernMajorityJournalDefault,omitempty"`
}

type ReplSetGetConfig struct {
	Config *Config `bson:"config"`
	Ok     int     `bson:"ok"`
}

type ConfigHandler struct {
	session   *mgo.Session
	config    *Config
	initiated bool
}

func New(session *mgo.Session) *ConfigHandler {
	return &ConfigHandler{
		session:   session,
		initiated: false,
	}
}

func NewConfig(rsName string) *Config {
	return &Config{
		Name:    rsName,
		Members: make([]*Member, 0),
	}
}

func (c *ConfigHandler) Get() *Config {
	return c.config
}

func (c *ConfigHandler) Set(config *Config) {
	c.config = config
}

func (c *ConfigHandler) Load() error {
	resp := &ReplSetGetConfig{}
	err := c.session.Run(bson.D{{"replSetGetConfig", 1}}, resp)
	if err != nil {
		return err
	}
	if resp.Config != nil {
		c.config = resp.Config
		c.initiated = true
	}
	return nil
}

func (c *ConfigHandler) IsInitiated() bool {
	if c.initiated {
		return true
	}
	err := c.Load()
	if err != nil {
		return false
	}
	return true
}

func (c *ConfigHandler) Initiate() error {
	if c.initiated {
		return nil
	}
	resp := &OkResponse{}
	err := c.session.Run(bson.D{{"replSetInitiate", c.config}}, resp)
	if err != nil {
		return err
	}
	if resp.Ok == 1 {
		c.initiated = true
	}
	return nil
}

func (c *ConfigHandler) Validate() error {
	if c.config.Name == "" {
		return ErrNoReplsetId
	}
	if len(c.config.Members) == 0 {
		return ErrNoReplsetMembers
	}
	return nil
}

func (c *ConfigHandler) Save() error {
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

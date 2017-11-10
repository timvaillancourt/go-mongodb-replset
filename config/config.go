package config

import (
	"encoding/json"
	"errors"
	"sync"

	"gopkg.in/mgo.v2/bson"
)

var (
	ErrNoReplsetId        = errors.New("replset config has no _id field!")
	ErrNoReplsetMembers   = errors.New("replset config has no members!")
	NoReplsetConfigString = "no replset config has been received"
)

type ReplsetTags map[string]string
type WriteConcern struct {
	WriteConcern interface{} `bson:"w" json:"w"`
	WriteTimeout int         `bson:"wtimeout" json:"wtimeout"`
	Journal      bool        `bson:"j,omitempty" json:"j,omitempty"`
}

type OkResponse struct {
	Ok int `bson:"ok" json:"ok" json:"ok"`
}

type Settings struct {
	ChainingAllowed         bool                    `bson:"chainingAllowed,omitempty" json:"chainingAllowed,omitempty"`
	HeartbeatIntervalMillis int64                   `bson:"heartbeatIntervalMillis,omitempty" json:"heartbeatIntervalMillis,omitempty"`
	HeartbeatTimeoutSecs    int                     `bson:"heartbeatTimeoutSecs,omitempty" json:"heartbeatTimeoutSecs,omitempty"`
	ElectionTimeoutMillis   int64                   `bson:"electionTimeoutMillis,omitempty" json:"electionTimeoutMillis,omitempty"`
	CatchUpTimeoutMillis    int64                   `bson:"catchUpTimeoutMillis,omitempty" json:"catchUpTimeoutMillis,omitempty"`
	GetLastErrorModes       map[string]*ReplsetTags `bson:"getLastErrorModes,omitempty" json:"getLastErrorModes,omitempty"`
	GetLastErrorDefaults    *WriteConcern           `bson:"getLastErrorDefaults,omitempty" json:"getLastErrorDefaults,omitempty"`
	ReplicaSetId            bson.ObjectId           `bson:"replicaSetId,omitempty" json:"replicaSetId,omitempty"`
}

type Config struct {
	sync.Mutex
	Name                               string    `bson:"_id" json:"_id"`
	Version                            int       `bson:"version" json:"version"`
	Members                            []*Member `bson:"members" json:"members"`
	Configsvr                          bool      `bson:"configsvr,omitempty" json:"configsvr,omitempty"`
	ProtocolVersion                    int       `bson:"protocolVersion,omitempty" json:"protocolVersion,omitempty"`
	Settings                           *Settings `bson:"settings,omitempty" json:"settings,omitempty"`
	WriteConcernMajorityJournalDefault bool      `bson:"writeConcernMajorityJournalDefault,omitempty" json:"writeConcernMajorityJournalDefault,omitempty"`
}

type ReplSetGetConfig struct {
	Config *Config `bson:"config" json:"config"`
	Errmsg string  `bson:"errmsg,omitempty" json:"errmsg,omitempty"`
	Ok     int     `bson:"ok" json:"ok" json:"ok"`
}

func NewConfig(rsName string) *Config {
	return &Config{
		Name:    rsName,
		Members: make([]*Member, 0),
		Version: 1,
	}
}

func (c *Config) IncrVersion() {
	c.Version = c.Version + 1
}

func (c *Config) ToString() string {
	raw, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return ""
	}
	return string(raw)
}

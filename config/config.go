package config

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ReplsetTags map[string]string

type Member struct {
	Id          int          `bson:"_id"`
	Host        string       `bson:"host"`
	ArbiterOnly bool         `bson:"arbiterOnly"`
	Hidden      bool         `bson:"hidden"`
	Priority    int          `bson:"priority"`
	Tags        *ReplsetTags `bson:"tags"`
	SlaveDelay  int64        `bson:"slaveDelay"`
	Votes       int          `bson:"votes"`
}

type WriteConcern struct {
	WriteConcern string `bson:"w"`
	WriteTimeout int    `bson:"wtimeout"`
}

type Settings struct {
	ChainingAllowed         bool                    `bson:"chainingAllowed,omitempty"`
	HeartbeatIntervalMillis int64                   `bson:"heartbeatIntervalMillis,omitempty"`
	HeartbeatTimeoutSecs    int64                   `bson:"heartbeatTimeoutSecs,omitempty"`
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
	Settings                           *Settings `bson:"settings"`
	Configsvr                          bool      `bson:"configsvr,omitempty"`
	ProtocolVersion                    int       `bson:"protocolVersion,omitempty"`
	WriteConcernMajorityJournalDefault bool      `bson:"writeConcernMajorityJournalDefault,omitempty"`
}

type ReplSetGetConfig struct {
	Config *Config `bson:"config"`
	Ok     int     `bson:"ok"`
}

func New(session *mgo.Session) (*Config, error) {
	resp := &ReplSetGetConfig{}
	err := session.DB("admin").Run(bson.D{{"replSetGetConfig", 1}}, resp)
	if err != nil {
		return nil, err
	}
	if resp.Ok == 1 && resp.Config != nil {
		return resp.Config, nil
	}
	return nil, nil
}

func (c *Config) AddMember(member *Member) {
	c.Members = append(c.Members, member)
}

func (c *Config) RemoveMember(removeMember *Member) {
	for i, member := range c.Members {
		if member.Host == removeMember.Host {
			c.Members = append(c.Members[:i], c.Members[i+1])
		}
	}
}

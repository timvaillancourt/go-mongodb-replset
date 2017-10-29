package config

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

func (c *Config) AddMember(member *Member) {
	c.Members = append(c.Members, member)
}

func (c *Config) RemoveMember(removeMember *Member) {
	for i, member := range c.Members {
		if member.Host == removeMember.Host {
			c.Members = append(c.Members[:i], c.Members[i+1])
			return
		}
	}
}

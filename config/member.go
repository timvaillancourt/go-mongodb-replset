package config

type Member struct {
	Id           int          `bson:"_id"`
	Host         string       `bson:"host"`
	ArbiterOnly  bool         `bson:"arbiterOnly"`
	BuildIndexes bool         `bson:"buildIndexes"`
	Hidden       bool         `bson:"hidden"`
	Priority     int          `bson:"priority"`
	Tags         *ReplsetTags `bson:"tags"`
	SlaveDelay   int64        `bson:"slaveDelay"`
	Votes        int          `bson:"votes"`
}

func NewMember(host string) *Member {
	return &Member{
		Host:         host,
		BuildIndexes: true,
		Priority:     1,
		Votes:        1,
	}
}

func (c *Config) getMemberMaxId() int {
	var maxId int
	for _, member := range c.Members {
		if member.Id > maxId {
			maxId = member.Id
		}
	}
	return maxId
}

func (c *Config) AddMember(member *Member) {
	member.Id = c.getMemberMaxId()
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

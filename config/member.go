package config

type Member struct {
	Id           int          `bson:"_id" json:"_id"`
	Host         string       `bson:"host" json:"host"`
	ArbiterOnly  bool         `bson:"arbiterOnly" json:"arbiterOnly"`
	BuildIndexes bool         `bson:"buildIndexes" json:"buildIndexes"`
	Hidden       bool         `bson:"hidden" json:"hidden"`
	Priority     int          `bson:"priority" json:"priority"`
	Tags         *ReplsetTags `bson:"tags,omitempty" json:"tags,omitempty"`
	SlaveDelay   int64        `bson:"slaveDelay" json:"slaveDelay"`
	Votes        int          `bson:"votes" json:"votes"`
}

func (c *Config) NewMember(host string) *Member {
	return &Member{
		Id:           c.getMemberMaxId() + 1,
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
	if c.HasMember(member.Host) {
		return
	}
	memberMaxId := c.getMemberMaxId()
	if member.Id <= memberMaxId {
		member.Id = memberMaxId + 1
	}
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

func (c *Config) GetMember(name string) *Member {
	for _, member := range c.Members {
		if member.Host == name {
			return member
		}
	}
	return nil
}

func (c *Config) HasMember(name string) bool {
	return c.GetMember(name) != nil
}

func (c *Config) GetMemberId(id int) *Member {
	for _, member := range c.Members {
		if member.Id == id {
			return member
		}
	}
	return nil
}

package status

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type MemberHealth int
type MemberState int

var (
	MemberHealthDown      MemberHealth = 0
	MemberHealthUp        MemberHealth = 1
	MemberStateStartup    MemberState  = 0
	MemberStatePrimary    MemberState  = 1
	MemberStateSecondary  MemberState  = 2
	MemberStateRecovering MemberState  = 4
	MemberStateStartup2   MemberState  = 5
	MemberStateUnknown    MemberState  = 6
	MemberStateArbiter    MemberState  = 7
	MemberStateDown       MemberState  = 8
	MemberStateRollback   MemberState  = 9
	MemberStateRemoved    MemberState  = 10
)

type Member struct {
	Id                int                 `bson:"_id"`
	Name              string              `bson:"name"`
	Health            MemberHealth        `bson:"health"`
	State             MemberState         `bson:"state"`
	StateStr          string              `bson:"stateStr"`
	Uptime            int64               `bson:"uptime"`
	Optime            *Optime             `bson:"optime"`
	OptimeDate        time.Time           `bson:"optimeDate"`
	ConfigVersion     int                 `bson:"configVersion"`
	ElectionTime      bson.MongoTimestamp `bson:"electionTime,omitempty"`
	ElectionDate      time.Time           `bson:"electionDate,omitempty"`
	InfoMessage       string              `bson:"infoMessage,omitempty"`
	OptimeDurable     *Optime             `bson:"optimeDurable,omitempty"`
	OptimeDurableDate time.Time           `bson:"optimeDurableDate,omitempty"`
	LastHeartbeat     time.Time           `bson:"lastHeartbeat,omitempty"`
	LastHeartbeatRecv time.Time           `bson:"lastHeartbeatRecv,omitempty"`
	PingMs            int64               `bson:"pingMs,omitempty"`
	Self              bool                `bson:"self,omitempty"`
	SyncingTo         string              `bson:"syncingTo,omitempty"`
}

func (s *Status) GetSelf() *Member {
	for _, member := range s.Members {
		if member.Self == true {
			return member
		}
	}
	return nil
}

func (s *Status) GetMember(name string) *Member {
	for _, member := range s.Members {
		if member.Name == name {
			return member
		}
	}
	return nil
}

func (s *Status) HasMember(name string) bool {
	return s.GetMember(name) != nil
}

func (s *Status) GetMemberId(id int) *Member {
	for _, member := range s.Members {
		if member.Id == id {
			return member
		}
	}
	return nil
}

func (s *Status) GetMembersByState(state MemberState, limit int) []*Member {
	members := make([]*Member, 0)
	for _, member := range s.Members {
		if member.State == state {
			members = append(members, member)
			if limit > 0 && len(members) == limit {
				return members
			}
		}
	}
	return members
}

func (s *Status) Primary() *Member {
	primary := s.GetMembersByState(MemberStatePrimary, 1)
	if len(primary) == 1 {
		return primary[0]
	}
	return nil
}

func (s *Status) Secondaries() []*Member {
	return s.GetMembersByState(MemberStateSecondary, 0)
}

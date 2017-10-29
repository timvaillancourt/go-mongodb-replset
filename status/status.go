package status

import (
	"time"

	"gopkg.in/mgo.v2"
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
	MemberStateRemoved    MemberState  = 10
)

type Optime struct {
	Timestamp bson.MongoTimestamp `bson:"ts"`
	Term      int64               `bson:"t"`
}

type StatusOptimes struct {
	LastCommittedOpTime *Optime `bson:"lastCommittedOpTime"`
	AppliedOpTime       *Optime `bson:"appliedOpTime"`
	DurableOptime       *Optime `bson:"durableOpTime"`
}

type StatusMember struct {
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
	SyncingTo         string              `bson:"syncingTo,omitempty"`
	Self              bool                `bson:"self,omitempty"`
}

type Status struct {
	Set                     string          `bson:"set"`
	Date                    time.Time       `bson:"date"`
	MyState                 MemberState     `bson:"myState"`
	Members                 []*StatusMember `bson:"members"`
	Term                    int64           `bson:"term,omitempty"`
	HeartbeatIntervalMillis int64           `bson:"heartbeatIntervalMillis,omitempty"`
	Optimes                 *StatusOptimes  `bson:"optimes,omitempty"`
	Ok                      int             `bson:"ok"`
}

func New(session *mgo.Session) (*Status, error) {
	status := &Status{}
	err := s.session.DB("admin").Run(bson.D{{"replSetGetStatus", 1}}, status)
	if err != nil {
		return nil, err
	}
	if status.Ok == 1 {
		return status, nil
	}
	return nil, nil
}

func (s *Status) GetPrimary() *StatusMember {
	for _, member := range s.Members {
		if member.State == MemberStatePrimary {
			return member
		}
	}
	return nil
}

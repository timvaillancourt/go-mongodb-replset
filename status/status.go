package status

import (
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

type Status struct {
	Set                     string         `bson:"set"`
	Date                    time.Time      `bson:"date"`
	MyState                 MemberState    `bson:"myState"`
	Members                 []*Member      `bson:"members"`
	Term                    int64          `bson:"term,omitempty"`
	HeartbeatIntervalMillis int64          `bson:"heartbeatIntervalMillis,omitempty"`
	Optimes                 *StatusOptimes `bson:"optimes,omitempty"`
	Ok                      int            `bson:"ok"`
}

func New(session *mgo.Session) (*Status, error) {
	status := &Status{}
	err := session.DB("admin").Run(bson.D{{"replSetGetStatus", 1}}, status)
	if err != nil {
		return nil, err
	}
	if status.Ok == 1 {
		return status, nil
	}
	return nil, nil
}

package mongo

import (
	"testing"

	"time"

	"gopkg.in/mgo.v2/bson"
)

type user struct {
	Id     bson.ObjectId `bson:"_id"`
	Name   string        `bson:"name"`
	Passwd string        `bson:"password"`
	Model  `bson:",omitempty"`
}

func (u *user) indexKeys() []string {
	return []string{"name"}
}

func (u *user) cName() string {
	return "user"
}
func (u *user) dbAddr() string {
	return "localhost"
}

func (u *user) dbName() string {
	return "test"
}

func TestDB(t *testing.T) {
	u := &user{}
	u.Id = bson.NewObjectId()
	u.Name = "arion"
	u1 := &user{}
	u1.Id = bson.NewObjectId()
	u1.Name = "arion2"
	err := Session(u).Insert(u)
	err = Session(u).Insert(u1)
	if err != nil {
		t.Error(err)
	}
	err = Session(u).Remove(u)
	if err != nil {
		t.Error(err)
	}
	time.Sleep(50 * time.Second)
}

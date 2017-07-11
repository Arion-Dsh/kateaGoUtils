package mongo

import (
	"testing"

	"gopkg.in/mgo.v2/bson"
)

type user struct {
	ID     bson.ObjectId `bson:"_id"`
	Name   string        `bson:"name"`
	Passwd string        `bson:"password"`
	Model  `bson:",omitempty"`
}

func (u *user) IndexKeys() []string {
	return []string{"name"}
}

func (u *user) Mate() map[string]string {
	return map[string]string{
		"dbAddr": "localhost",
		"dbName": "test",
		"cName":  "user",
	}
}

func TestDB(t *testing.T) {
	a := map[string]string{
		"localhost": "mongodb://localhost",
	}
	Dial(a)
	u := &user{}
	u.ID = bson.NewObjectId()
	u.Name = "arion"
	u1 := &user{}
	u1.ID = bson.NewObjectId()
	u1.Name = "arion2"
	err := C(u).Insert(u)
	if err != nil {
		t.Error(err)
	}
	err = C(u).Remove(u)
	if err != nil {
		t.Error(err)
	}
	// time.Sleep(50 * time.Second)
}

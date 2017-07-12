package mongo

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

type model interface {
	Mate() map[string]string
	Indexs() []mgo.Index
}

var sessions map[string]*mgo.Session = make(map[string]*mgo.Session, 0)

var setIndex = make(map[string]bool, 0)

// Dial ...
func Dial(urls map[string]string) {
	for k, v := range urls {
		s, err := mgo.Dial(v)
		if err != err {
			log.Fatalf("please check out -- %s --, it's not be contected", v)
		}
		sessions[k] = s
	}
}

//Session ...
func Session(k string) *mgo.Session {
	if _, ok := sessions[k]; !ok {
		panic(fmt.Sprintf("please check Dial function, %s is not in it", k))
	}
	return sessions[k]
}

// DB ...
func DB(m model) *mgo.Database {
	mate := m.Mate()

	if mate["dbAddr"] == "" || mate["dbName"] == "" || mate["cName"] == "" {
		panic("check model mate, must include dbAddr, dbName, cName")
	}
	session := Session(mate["dbAddr"])
	return session.DB(mate["dbName"])

}

// C Collection alias
func C(m model) *mgo.Collection {
	db := DB(m)
	mate := m.Mate()
	c := db.C(mate["cName"])
	if !setIndex[mate["cName"]] {
		for _, index := range m.Indexs() {
			if err := c.EnsureIndex(index); err != nil {
				panic(err)
			}
		}
		setIndex[mate["cName"]] = true
	}
	return c
}

// GridFS GridFS alias
func GridFS(m model) *mgo.GridFS {
	db := DB(m)
	mate := m.Mate()
	return db.GridFS(mate["cName"])
}

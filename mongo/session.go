package mongo

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type model interface {
	Mate() map[string]string
	IndexKeys() []string
}

func index(keys []string, collection *mgo.Collection) {
	// Index
	_index := mgo.Index{
		Key:      keys,
		Unique:   true,
		DropDups: true,
		// Background: true,
		// Sparse:     true,
	}

	err := collection.EnsureIndex(_index)
	if err != nil {
		panic(err)
	}
}

// Session
func Session(m model) *db {
	mate := m.Mate()

	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s", mate["dbAddr"]))

	defer session.Close()
	if err != nil {
		panic(err)
	}
	db := &db{
		session: session.Copy(),
		dbName:  mate["dbName"],
		cName:   mate["cName"],
	}
	// db.collection = db.session.DB(m.dbName()).C(m.cName())
	// index(m.indexKeys(), db.collection)
	return db
}

type db struct {
	query   *mgo.Query
	session *mgo.Session
	dbName  string
	cName   string
}

func (db *db) Insert(docs ...interface{}) error {
	defer db.session.Close()
	return db.session.DB(db.dbName).C(db.cName).Insert(docs...)
}

func (db *db) Remove(selector interface{}) error {
	defer db.session.Close()
	return db.session.DB(db.dbName).C(db.cName).Remove(selector)

}
func (db *db) Find(query interface{}) *db {
	db.query = db.session.DB(db.dbName).C(db.cName).Find(query)
	return db
}

func (db *db) FindId(id interface{}) *db {
	db.query = db.session.DB(db.dbName).C(db.cName).FindId(id)
	return db
}

func (db *db) One(result interface{}) error {
	defer db.session.Close()
	return db.query.One(result)
}

func (db *db) Count() (int, error) {
	defer db.session.Close()
	return db.query.Count()
}

package mongo

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type model interface {
	dbAddr() string
	dbName() string
	cName() string
	indexKeys() []string
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

// mgodb Session
func Session(m model) *db {
	session, err := mgo.Dial(fmt.Sprintf("mongodb://%s", m.dbAddr()))
	defer session.Close()
	if err != nil {
		panic(err)
	}
	db := &db{
		session: session.Copy(),
		dbName:  m.dbName(),
		cName:   m.cName(),
	}
	// db.collection = db.session.DB(m.dbName()).C(m.cName())
	// index(m.indexKeys(), db.collection)
	return db
}

type db struct {
	query      *mgo.Query
	session    *mgo.Session
	collection *mgo.Collection
	dbName     string
	cName      string
}

func (db *db) Insert(docs ...interface{}) error {
	defer db.session.Close()
	err := db.session.DB(db.dbName).C(db.cName).Insert(docs...)
	if err != nil {
		return err
	}
	return nil
}
func (db *db) Remove(selector interface{}) error {
	defer db.session.Close()
	err := db.session.DB(db.dbName).C(db.cName).Remove(selector)
	if err != nil {
		return err
	}
	return nil
}

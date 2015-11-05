package persistence

import (
	"github.com/dst-hackathon/ring/config"
	"gopkg.in/mgo.v2"
)

type Database struct {
	session *mgo.Session
	name    string
}

func NewDatabase(name string) *Database {
	cfg := config.New()
	session, err := mgo.Dial(cfg.Database["dev"].Host())
	if err != nil {
		panic(err)
	}

	return &Database{session, name}
}

func (db Database) Session() *mgo.Session {
	return db.session.Copy()
}

func (db Database) Close() {
	db.session.Close()
}

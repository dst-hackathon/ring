package persistence

import (
	"github.com/dst-hackathon/ring/config"
	"gopkg.in/mgo.v2"
)

type Database interface {
	Session() *mgo.Session
	Name() string
	Close()
}

type MongoDB struct {
	session *mgo.Session
	name    string
}

func (m MongoDB) Session() *mgo.Session {
	return m.session.Copy()
}

func (m MongoDB) Name() string {
	return m.name
}

func (m MongoDB) Close() {
	m.session.Close()
}

func NewDatabase(name string) Database {
	cfg := config.New()
	session, err := mgo.Dial(cfg.Database["dev"].Host())
	if err != nil {
		panic(err)
	}

	return &MongoDB{session, name}
}

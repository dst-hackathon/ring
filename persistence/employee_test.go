package persistence

import (
	"gopkg.in/mgo.v2"
	"testing"
)

type TestDatabase struct {
}

func (t TestDatabase) Session() *mgo.Session {
	return nil
}

func (t TestDatabase) Name() string {
	return "TestDB"
}

func (t TestDatabase) Close() {
	// do nothing
}

func TestNewEmployeeRepo(t *testing.T) {
	e := NewEmployeeRepo(&TestDatabase{})

	if e == nil {
		t.Fail()
	}
}

package persistence

import (
	"github.com/dst-hackathon/ring/entity"
	"github.com/dst-hackathon/ring/repository"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EmployeeRepoHandler struct {
	DB      string
	session *mgo.Session
}

func NewEmployeeRepo(db Database) repository.Employee {
	return &EmployeeRepoHandler{db.Name(), db.Session()}
}

func (handler EmployeeRepoHandler) C() string {
	return "employees"
}

func (handler EmployeeRepoHandler) FindByFirstName(name string) entity.Employee {
	c := handler.session.DB(handler.DB).C(handler.C())
	employee := entity.Employee{}
	err := c.Find(bson.M{"first_name": name}).One(&employee)
	if err != nil {
		panic(err)
	}

	return employee
}

func (handler EmployeeRepoHandler) Save(employee entity.Employee) entity.Employee {
	c := handler.session.DB(handler.DB).C(handler.C())
	err := c.Insert(employee)
	if err != nil {
		panic(err)
	}

	return employee
}

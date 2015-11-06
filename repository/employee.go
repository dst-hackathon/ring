package repository

import (
	"github.com/dst-hackathon/ring/entity"
)

type Repository interface {
	Close()
}

type Employee interface {
	All() []entity.Employee
	Save(employee entity.Employee) entity.Employee
	FindByID(id string) entity.Employee
	FindByFirstName(name string) entity.Employee
}

package repository

import (
	"github.com/dst-hackathon/ring/entity"
)

type Repository interface {
	Close()
}

type Employee interface {
	Save(employee entity.Employee) entity.Employee
	FindByFirstName(name string) entity.Employee
}

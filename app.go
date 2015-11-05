package main

import (
	"fmt"
	"github.com/dst-hackathon/ring/entity"
	"github.com/dst-hackathon/ring/persistence"
)

func main() {
	db := persistence.NewDatabase("dst")
	defer db.Close()

	repo := persistence.NewEmployeeRepo(db)
	e := repo.FindByFirstName("Sid")
	fmt.Println(e)

	quin := &entity.Employee{FirstName: "Quin 2", LastName: "Liss"}
	repo.Save(*quin)
}

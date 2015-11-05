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
	e := repo.FindByFirstName("Peter")
	fmt.Println(e)

	quin := &entity.Employee{FirstName: "Peter", LastName: "Quin"}
	repo.Save(*quin)
}

package main

import (
	"github.com/codegangsta/negroni"
	"github.com/dst-hackathon/ring/handler"
	"github.com/dst-hackathon/ring/persistence"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := persistence.NewDatabase("dst")
	defer db.Close()
	mux := httprouter.New()

	repo := persistence.NewEmployeeRepo(db)
	employee := handler.Employee{repo}
	mux.GET("/employees", employee.Index)
	mux.GET("/employees/:id", employee.Show)
	mux.POST("/employees", employee.Create)

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

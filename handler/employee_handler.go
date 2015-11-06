package handler

import (
	"encoding/json"
	"github.com/dst-hackathon/ring/entity"
	"github.com/dst-hackathon/ring/repository"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Employee struct {
	Repo repository.Employee
}

func (h Employee) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(h.Repo.All()); err != nil {
		panic(err)
	}
}

func (h Employee) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(h.Repo.FindByID(id)); err != nil {
		panic(err)
	}
}

func (h Employee) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var emp entity.Employee
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&emp)
	if err != nil {
		panic(err)
	}

	h.Repo.Save(emp)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

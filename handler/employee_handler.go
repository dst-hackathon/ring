package handler

import (
	"github.com/dst-hackathon/ring/entity"
	"github.com/dst-hackathon/ring/repository"
	"github.com/julienschmidt/httprouter"
	"github.com/shwoodard/jsonapi"
	"net/http"
)

type Employee struct {
	Repo repository.Employee
}

func (h Employee) Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	emps := h.Repo.All()
	res := make([]interface{}, len(emps))
	for i, e := range emps {
		res[i] = &e
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)
	if err := jsonapi.MarshalManyPayload(w, res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Employee) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	emp := h.Repo.FindByID(ps.ByName("id"))
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusOK)
	if err := jsonapi.MarshalOnePayload(w, &emp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h Employee) Create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var emp entity.Employee
	if err := jsonapi.UnmarshalPayload(r.Body, &emp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	h.Repo.Save(emp)
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusCreated)
}

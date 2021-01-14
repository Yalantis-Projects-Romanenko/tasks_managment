package handlers

import (
	"fmt"
	"github.com/fdistorted/task_managment/handlers/projects"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	// r.Use(middleware.Recoverer)
	// r.Use(middlewares.RequestID)

	r.HandleFunc("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hi there ^._.^")
	}).Methods("GET").Schemes("http")

	// subpaths
	s := r.PathPrefix("/projects").Subrouter()
	s.HandleFunc("/", projects.GetAll).Methods("GET")
	s.HandleFunc("/{project_id}/", projects.Get).Methods("GET")
	//// "/products/{key}/details"
	//s.HandleFunc("/{key}/details", ProductDetailsHandler)

	return r
}

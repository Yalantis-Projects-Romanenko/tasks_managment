package handlers

import (
	"fmt"
	"github.com/fdistorted/task_managment/handlers/columns"
	"github.com/fdistorted/task_managment/handlers/comments"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/handlers/projects"
	"github.com/fdistorted/task_managment/handlers/tasks"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	//r.Use(middleware.Recoverer)
	r.Use(middlewares.RequestID)

	r.HandleFunc("/healthCheck", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "hi there ^._.^")
	}).Methods("GET").Schemes("http")

	// subpaths
	projectsRouter := r.PathPrefix("/projects").Subrouter()
	projectsRouter.Use(middlewares.Authorize) // enable authorization for subrouter
	projectsRouter.HandleFunc("/", projects.GetAll).Methods("GET")
	projectsRouter.HandleFunc("/", projects.Post).Methods("POST")
	projectsRouter.HandleFunc("/{id}/", projects.Get).Methods("GET")
	projectsRouter.HandleFunc("/{id}/", projects.Delete).Methods("DELETE")
	projectsRouter.HandleFunc("/{id}/", projects.Put).Methods("PUT")
	//// "/products/{key}/details"
	//projectsRouter.HandleFunc("/{key}/details", ProductDetailsHandler)

	columnsRouter := r.PathPrefix("/columns").Subrouter()
	columnsRouter.HandleFunc("/{id}/", columns.Get).Methods("GET")

	tasksRouter := r.PathPrefix("/tasks").Subrouter()
	tasksRouter.HandleFunc("/{id}/", tasks.Get).Methods("GET")

	commentsRouter := r.PathPrefix("/comments").Subrouter()
	commentsRouter.HandleFunc("/{id}/", comments.Get).Methods("GET")

	return r
}

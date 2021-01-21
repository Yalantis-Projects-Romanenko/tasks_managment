package handlers

import (
	"fmt"
	"github.com/fdistorted/task_managment/handlers/columns"
	"github.com/fdistorted/task_managment/handlers/middlewares"
	"github.com/fdistorted/task_managment/handlers/projects"
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
	}).Methods(http.MethodGet).Schemes("http")

	// subpaths
	projectsRouter := r.PathPrefix("/projects").Subrouter()
	projectsRouter.Use(middlewares.Authorize) // enable authorization for subrouter
	projectsRouter.HandleFunc("/", projects.GetAll).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/", projects.Post).Methods(http.MethodPost)
	projectsRouter.HandleFunc("/{id}/", projects.Get).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/{id}/", projects.Delete).Methods(http.MethodDelete)
	projectsRouter.HandleFunc("/{id}/", projects.Put).Methods(http.MethodPut)

	columnsRouter := r.PathPrefix("/columns").Subrouter()
	columnsRouter.Use(middlewares.Authorize) // enable authorization for subrouter
	columnsRouter.HandleFunc("/{id}/", columns.Get).Methods(http.MethodGet)

	//tasksRouter := r.PathPrefix("/tasks").Subrouter()
	//tasksRouter.HandleFunc("/{id}/", tasks.Get).Methods(http.MethodGet)

	//commentsRouter := r.PathPrefix("/comments").Subrouter()
	//commentsRouter.HandleFunc("/{id}/", comments.Get).Methods(http.MethodGet)

	return r
}

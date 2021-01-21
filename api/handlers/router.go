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
	projectsRouter.HandleFunc("/{projectId}/", projects.Get).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/{projectId}/", projects.Delete).Methods(http.MethodDelete)
	projectsRouter.HandleFunc("/{projectId}/", projects.Put).Methods(http.MethodPut)
	//columns part
	projectsRouter.HandleFunc("/{projectId}/columns/", columns.GetAll).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/{projectId}/columns/", columns.Post).Methods(http.MethodPost)
	projectsRouter.HandleFunc("/{projectId}/columns-move/", columns.Move).Methods(http.MethodPost)

	//tasksRouter := r.PathPrefix("/tasks").Subrouter()
	//tasksRouter.HandleFunc("/ {projectId}/", tasks.Get).Methods(http.MethodGet)

	//commentsRouter := r.PathPrefix("/comments").Subrouter()
	//commentsRouter.HandleFunc("/ {projectId}/", comments.Get).Methods(http.MethodGet)

	return r
}

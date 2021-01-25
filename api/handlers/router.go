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
	}).Methods(http.MethodGet).Schemes("http")

	// projects subrouter
	projectsRouter := r.PathPrefix("/projects").Subrouter()
	projectsRouter.Use(middlewares.Authorize) // enable authorization for subrouter
	projectsRouter.HandleFunc("/", projects.GetAll).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/", projects.Post).Methods(http.MethodPost)
	projectsRouter.HandleFunc("/{projectId}/", projects.Get).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/{projectId}/", projects.Delete).Methods(http.MethodDelete)
	projectsRouter.HandleFunc("/{projectId}/", projects.Put).Methods(http.MethodPut)

	//columns
	projectsRouter.HandleFunc("/{projectId}/columns/", columns.GetAll).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/{projectId}/columns/", columns.Post).Methods(http.MethodPost)
	projectsRouter.HandleFunc("/{projectId}/columns/{columnId}/", columns.Delete).Methods(http.MethodDelete)
	projectsRouter.HandleFunc("/{projectId}/columns/{columnId}/", columns.Put).Methods(http.MethodPut)
	projectsRouter.HandleFunc("/{projectId}/columns/{columnId}/", columns.Get).Methods(http.MethodGet)

	//tasks
	projectsRouter.HandleFunc("/{projectId}/columns/{columnId}/tasks/", tasks.Post).Methods(http.MethodPost)
	projectsRouter.HandleFunc("/{projectId}/columns/{columnId}/tasks/", tasks.GetAll).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/{projectId}/columns/{columnId}/tasks/{taskId}/", tasks.Put).Methods(http.MethodPut)
	projectsRouter.HandleFunc("/{projectId}/columns/{columnId}/tasks/{taskId}/", tasks.Delete).Methods(http.MethodDelete)
	projectsRouter.HandleFunc("/{projectId}/columns/{columnId}/tasks/{taskId}/", tasks.Get).Methods(http.MethodGet)

	//comments
	projectsRouter.HandleFunc("/{projectId}/columns/{columnId}/tasks/{taskId}/comments/", comments.Post).Methods(http.MethodPost)
	projectsRouter.HandleFunc("/{projectId}/columns/{columnId}/tasks/{taskId}/comments/", comments.GetAll).Methods(http.MethodGet)
	projectsRouter.HandleFunc("/{projectId}/columns/{columnId}/tasks/{taskId}/comments/{commentId}/", comments.Get).Methods(http.MethodGet)

	return r
}

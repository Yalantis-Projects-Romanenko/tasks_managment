package handlers

import (
	"fmt"
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

	// // subpaths
	// s := r.PathPrefix("/boards").Subrouter()
	// s.HandleFunc("/", ProductsHandler)
	// // "/products/{key}/"
	// s.HandleFunc("/{key}/", ProductHandler)
	// // "/products/{key}/details"
	// s.HandleFunc("/{key}/details", ProductDetailsHandler)

	return r
}

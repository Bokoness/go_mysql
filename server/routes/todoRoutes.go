package routes

import (
	"go_mysql/server/handlers/todo"
	"go_mysql/server/middleware"

	"github.com/gorilla/mux"
)

func TodoRoutes(r *mux.Router) {
	s := r.PathPrefix("/todo").Subrouter()
	s.HandleFunc("/", todo.Index).Methods("GET")
	s.HandleFunc("/{id}", todo.Show).Methods("GET")
	s.HandleFunc("/", todo.Store).Methods("POST")
	s.HandleFunc("/{id}", todo.Update).Methods("PUT")
	s.HandleFunc("/{id}", todo.Destroy).Methods("DELETE")
	s.Use(middleware.UserAuth)
}

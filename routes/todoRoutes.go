package routes

import (
	"go_mysql/db/models"
	"go_mysql/handlers"
	"go_mysql/middleware"

	"github.com/gorilla/mux"
)

// func TodoRoutes(r *mux.Router) {
// 	s := r.PathPrefix("/todo").Subrouter()
// 	s.HandleFunc("/", todo.Index).Methods("GET")
// 	s.HandleFunc("/{id}", todo.Show).Methods("GET")
// 	s.HandleFunc("/", todo.Store).Methods("POST")
// 	s.HandleFunc("/{id}", todo.Update).Methods("PUT")
// 	s.HandleFunc("/{id}", todo.Destroy).Methods("DELETE")
// 	s.Use(middleware.UserAuth)
// }
func TodoRoutes(r *mux.Router) {
	h := handlers.Handler{
		Model: &models.TodoModule{},
	}
	s := r.PathPrefix("/todo").Subrouter()
	s.HandleFunc("/", h.Index).Methods("GET")
	s.HandleFunc("/{id}", h.Show).Methods("GET")
	s.HandleFunc("/", h.Store).Methods("POST")
	s.HandleFunc("/{id}", h.Update).Methods("PUT")
	s.HandleFunc("/{id}", h.Destroy).Methods("DELETE")
	s.Use(middleware.UserAuth)
}

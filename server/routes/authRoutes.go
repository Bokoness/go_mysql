package routes

import (
	"go_mysql/server/handlers/auth"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/login", auth.Login).Methods("POST")
	s.HandleFunc("/register", auth.Register).Methods("POST")
	s.HandleFunc("/logout", auth.Logout).Methods("POST")
}

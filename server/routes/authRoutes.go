package routes

import (
	controller "go_mysql/server/controllers/auth"

	"github.com/gorilla/mux"
)

func CreateAuthRoutes(r *mux.Router) {
	authRoutes := r.PathPrefix("/auth").Subrouter()
	authRoutes.HandleFunc("/register", controller.Register).Methods("POST")
	authRoutes.HandleFunc("/login", controller.Login).Methods("POST")
	authRoutes.HandleFunc("/logout", controller.Logout).Methods("POST")
}

package routes

import (
	controllers "go_mysql/server/controllers/auth"

	"github.com/gorilla/mux"
)

func CreateAuthRoutes(r *mux.Router) {
	authRoutes := r.PathPrefix("/auth").Subrouter()
	authRoutes.HandleFunc("/register", controllers.Register).Methods("POST")
}

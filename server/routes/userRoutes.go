package routes

import (
	controller "go_mysql/server/controllers/user"

	"github.com/gorilla/mux"
)

func CreateUserRoutes(r *mux.Router) {
	userRoutes := r.PathPrefix("/user").Subrouter()
	userRoutes.HandleFunc("/{id}", controller.Destroy).Methods("DELETE")
}

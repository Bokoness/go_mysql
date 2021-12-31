package routes

import (
	controller "go_mysql/server/controllers/todo"
	auth "go_mysql/server/middleware"

	"github.com/gorilla/mux"
)

func CreateTodoRoutes(r *mux.Router) {
	todoRoutes := r.PathPrefix("/todo").Subrouter()
	todoRoutes.HandleFunc("/", controller.Store).Methods("POST")
	todoRoutes.Use(auth.UserAuth)
	// authRoutes.HandleFunc("/index", controller.Register).Methods("POST")
}

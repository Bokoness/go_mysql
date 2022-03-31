package server

import (
	"go_mysql/server/middleware"
	"go_mysql/server/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func LunchServer() {
	r := mux.NewRouter()
	routes.TodoRoutes(r)
	routes.AuthRoutes(r)
	r.Use(middleware.ServerHeaders)
	http.ListenAndServe(":8000", r)
}

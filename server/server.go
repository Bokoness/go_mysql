package server

import (
	"go_mysql/server/routes"
	"log"
	"net/http"

	"github.com/gookit/color"
	"github.com/gorilla/mux"
)

func LunchServer() {
	r := mux.NewRouter()
	routes.CreateAuthRoutes(r)
	routes.CreateUserRoutes(r)
	color.Cyan.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))
}

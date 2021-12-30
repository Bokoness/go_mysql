package server

import (
	"fmt"
	"go_mysql/server/routes"
	"log"
	"net/http"
	"os"

	"github.com/gookit/color"
	"github.com/gorilla/mux"
)

func LunchServer() {
	r := mux.NewRouter()
	routes.CreateAuthRoutes(r)
	routes.CreateUserRoutes(r)
	h := os.Getenv("HOST")
	p := os.Getenv("PORT")
	color.Cyan.Printf("Server is running on port %s", p)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", h, p), r))
}

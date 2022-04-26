package main

import (
	"fmt"
	"go_mysql/routes"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	LunchServer()
}

func LunchServer() {
	r := mux.NewRouter()
	routes.AuthRoutes(r)
	routes.TodoRoutes(r)
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	color.Green("Server is running on port %s", port)
	http.ListenAndServe(port, r)
}

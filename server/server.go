package server

import (
	"go_mysql/server/routes"
	"os"

	"github.com/go-martini/martini"
	"github.com/gookit/color"
)

func LunchServer() {
	m := martini.Classic()

	routes.AuthRoutes(m)
	routes.UserRoutes(m)
	routes.TodoRoutes(m)

	p := os.Getenv("PORT")
	color.Cyan.Printf("Server is running on port %s", p)
	m.Run()
}

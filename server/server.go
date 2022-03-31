package server

import (
	"go_mysql/server/routes"
	"net/http"

	"github.com/gorilla/mux"
)

// func LunchServer() {
// 	m := martini.Classic()

// 	routes.AuthRoutes(m)
// 	routes.UserRoutes(m)
// 	routes.TodoRoutes(m)

// 	p := os.Getenv("PORT")
// 	color.Cyan.Printf("Server is running on port %s", p)
// 	m.Run()
// }

func LunchServer() {
	r := mux.NewRouter()
	routes.TodoRoutes(r)
	routes.AuthRoutes(r)
	http.ListenAndServe(":8000", r)
}

package routes

import (
	handler "go_mysql/server/handlers/todo"
	"go_mysql/server/middleware"

	"github.com/go-martini/martini"
)

func TodoRoutes(m *martini.ClassicMartini) {
	m.Group("/todo", func(r martini.Router) {
		r.Get("/", middleware.UserAuth, handler.Index)
		r.Get("/:id", middleware.UserAuth, handler.Show)
		r.Post("/", middleware.UserAuth, handler.Store)
		r.Delete("/:id", middleware.UserAuth, handler.Destroy)
	})
}

package routes

import (
	handler "go_mysql/server/handlers/user"

	"github.com/go-martini/martini"
)

func UserRoutes(m *martini.ClassicMartini) {
	m.Group("/user", func(r martini.Router) {
		r.Delete("/:id", handler.Destroy)
	})
}

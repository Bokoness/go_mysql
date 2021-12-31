package routes

import (
	handler "go_mysql/server/handlers/auth"

	"github.com/go-martini/martini"
)

func AuthRoutes(m *martini.ClassicMartini) {
	m.Group("/auth", func(r martini.Router) {
		r.Post("/register", handler.Register)
		r.Post("/login", handler.Login)
		r.Post("/logout", handler.Logout)
	})
}

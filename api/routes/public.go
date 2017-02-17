package routes

import (
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
	"github.com/rs/cors"

	controller "../controllers"
	db "../dbs"
)

//Public Routes
func Public(s *db.Dispatch, cors *cors.Cors) func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.Logger)
		r.Use(middleware.Recoverer)
		r.Use(middleware.DefaultCompress)
		r.Use(cors.Handler)

		// home
		r.Get("/", controller.Home())

		// Authenticate user
		r.Post("/auth", controller.Auth())

		//CRUD User
		r.Post("/user", controller.CreateUser(s))
	}
}

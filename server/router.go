package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"usersvc.io/api/v1/server/handlers"
)

func BuildRouter() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.ListUsers)
		r.Post("/", handlers.CreateUser)

		r.Route("/{userID}", func(r chi.Router) {
			r.Use(handlers.UserCtx)
			r.Get("/", handlers.GetUser)
			r.Put("/", handlers.UpdateUser)
			r.Delete("/", handlers.DeleteUser)
		})

	})

	return r
}

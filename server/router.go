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

	// RESTy routes for "users" resource
	r.Route("/users", func(r chi.Router) {
		r.Get("/", handlers.ListUsers)
		r.Post("/", handlers.CreateUser) // POST /users

		r.Route("/{userID}", func(r chi.Router) {
			r.Use(handlers.UserCtx)            // Load the *User on the request context
			r.Get("/", handlers.GetUser)       // GET /users/123
			r.Put("/", handlers.UpdateUser)    // PUT /users/123
			r.Delete("/", handlers.DeleteUser) // DELETE /users/123
		})

	})

	return r
}

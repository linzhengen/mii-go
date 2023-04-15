package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/linzhengen/mii-go/app/interface/api/handler"
	"net/http"
)

func New(
	healthHandler handler.HealthHandler,
	userHandler handler.UserHandler,
) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/healthz", healthHandler.Get)
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/users/{userId}", userHandler.Get)
		r.Post("/users", userHandler.Post)
	})
	return r
}

package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/mateusrlopez/url-shortener-server/http/handlers"
	"net/http"
)

func NewRouter(context string, handler *handlers.Records) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"OPTIONS", "POST", "GET"},
		AllowedHeaders: []string{"Content-Type", "Accepts", "Origin"},
		MaxAge:         300,
	}))

	router.Route(context, func(r chi.Router) {
		r.Route("/records", func(r chi.Router) {
			r.Post("/", handler.Create)
			r.Get("/{key}", handler.FindOneByKey)
			r.Get("/{key}/redirect", handler.Redirect)
		})
	})

	return router
}

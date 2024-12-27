package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"

	"dbaas-api/api/resource/health"
	"dbaas-api/api/resource/instance"
	"dbaas-api/api/router/middleware"
	"dbaas-api/api/router/middleware/requestlog"
)

func New(l *zerolog.Logger, v *validator.Validate) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/livez", health.Read)

	r.Route("/v1", func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Use(middleware.ContentTypeJSON)

		instanceAPI := instance.New(l, v)
		r.Method(http.MethodGet, "/instances", requestlog.NewHandler(instanceAPI.List, l))
		r.Method(http.MethodPost, "/instances", requestlog.NewHandler(instanceAPI.Create, l))
		r.Method(http.MethodGet, "/instances/{id}", requestlog.NewHandler(instanceAPI.Get, l))
		r.Method(http.MethodDelete, "/instances/{id}", requestlog.NewHandler(instanceAPI.Delete, l))
	})

	return r
}

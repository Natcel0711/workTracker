package api

import "github.com/go-chi/chi/v5"

func SetupRoutes(r *chi.Mux) {
	r.HandleFunc("/", HealthHandler)
}

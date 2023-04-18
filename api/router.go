package api

import "github.com/go-chi/chi/v5"

func SetupRoutes(r *chi.Mux) {
	r.HandleFunc("/", GetTimesheet)
	r.HandleFunc("/timesheet", CreateTimesheet)
	r.HandleFunc("/timesheet", UpdateTimesheet)
	r.HandleFunc("/timesheet", DeleteTimesheet)
}

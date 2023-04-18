package main

import (
	"fmt"
	"net/http"

	"github.com/Natcel0711/workTracker/api"
	"github.com/Natcel0711/workTracker/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	err := db.InitDB()
	defer db.CloseDB()
	if err != nil {
		panic(err)
	}
	api.SetupRoutes(r)

	fmt.Println("Listening on port :3000")
	http.ListenAndServe(":3000", r)
}

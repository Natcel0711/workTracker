package main

import (
	"fmt"
	"net/http"

	"github.com/Natcel0711/workTracker/api"
	"github.com/Natcel0711/workTracker/config"
	"github.com/Natcel0711/workTracker/db"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	config, err := config.Init()
	if err != nil {
		panic(err)
	}
	r.Use(middleware.Logger)
	err = db.InitDB(config.ConnectionString)
	defer db.CloseDB()
	if err != nil {
		panic(err)
	}
	api.SetupRoutes(r)

	fmt.Println("Listening on port :3000")
	http.ListenAndServe(":"+config.Port, r)
}

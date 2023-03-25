package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:[]string{"https://*","http://*"},
		AllowedMethods:
	}))

}

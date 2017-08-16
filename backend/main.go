package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	router := httprouter.New()

	// Routes

	router.POST("/signup", signup)
	router.POST("/auth", authenticate)
	router.POST("/logout", logout)

	// Prefer white list domains with cors.New().Options({AllowedOrigins...})
	cors := cors.Default().Handler(router)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: cors,
	}
	server.ListenAndServe()
}

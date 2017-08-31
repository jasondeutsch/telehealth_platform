package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	loadConfig()
	router := httprouter.New()

	// Admin API
	router.GET("/admin/patients/all", adminAllPatients)

	// Auth API
	router.POST("/signup", signup)
	router.POST("/auth", authenticate)
	router.POST("/logout", logout)

	// Patient API
	router.POST("/patient/create/", createPatient)

	// Prefer white list domains with cors.New().Options({AllowedOrigins...})
	cors := cors.Default().Handler(router)

	server := &http.Server{
		Addr:    config.Address,
		Handler: cors,
	}
	fmt.Println("Listening on port " + config.Address)
	server.ListenAndServe()
}

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

	router.GET("/patient", indexPatient)
	router.GET("/patient/:id", showPatient)
	router.POST("/patient/create/", createPatient)

	router.POST("/provider/create", createProvider)
	//router.GET("/provider/patient/index", indexProviderPatient)
	//router.GET("/provider/patient/show", showProviderPatient)

	//router.POST("/admin/provider/update", adminUpdateProvider)

	// Auth API
	router.POST("/signup", signup)
	router.POST("/auth", authenticate)
	router.POST("/logout", logout)

	// Prefer white list domains with cors.New().Options({AllowedOrigins...})
	cors := cors.Default().Handler(router)

	server := &http.Server{
		Addr:    config.Address,
		Handler: cors,
	}
	fmt.Println("Listening on port " + config.Address)
	server.ListenAndServe()
}

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func main() {
	loadConfig()
	r := mux.NewRouter()

	r.HandleFunc("/", index)
	r.HandleFunc("/login", login)

	r.HandleFunc("/patient", indexPatient).Methods("GET")
	r.HandleFunc("/patient/show", showPatient).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/patient/create", createPatient).Methods("POST").Headers("Content-Type", "application/json")

	r.HandleFunc("/provider", indexProvider).Methods("GET")
	r.HandleFunc("/provider/create", createProvider).Methods("POST").Headers("Content-Type", "application/json")

	//Auth API
	r.HandleFunc("/signup", signup).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/auth", authenticate).Methods("POST")
	r.HandleFunc("/logout", logout).Methods("DELETE")

	// Prefer white list domains with cors.New().Options({AllowedOrigins...})
	cors := cors.Default().Handler(r)

	server := &http.Server{
		Addr:    config.Address,
		Handler: cors,
	}
	fmt.Println("Listening on port " + config.Address)
	server.ListenAndServe()
}

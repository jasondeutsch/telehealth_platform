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
	r.HandleFunc("/login", login).Methods("GET")
	r.HandleFunc("/signup", signup).Methods("GET")
	r.HandleFunc("/signupaccount", signupAccount).Methods("POST")
	r.HandleFunc("/auth", authenticate).Methods("POST")
	r.HandleFunc("/logout", logout).Methods("DELETE")

	r.HandleFunc("/admin", adminIndex)
	r.HandleFunc("/admin/patients/{id}", showPatient)
	r.HandleFunc("/admin/providers/{id}", showProvider)
	//r.HandleFunc("/admin/providers", indexProvider).Methods("GET")

	r.HandleFunc("/provider/patient", indexPatient).Methods("GET")
	r.HandleFunc("/provider/patient/{id}", showPatient).Methods("GET")
	r.HandleFunc("/providers/create", createProvider).Methods("POST")

	r.HandleFunc("/patient/create", createPatient).Methods("POST").Headers("Content-Type", "application/json")

	// Prefer white list domains with cors.New().Options({AllowedOrigins...})
	cors := cors.Default().Handler(r)

	server := &http.Server{
		Addr:    config.Address,
		Handler: cors,
	}
	fmt.Println("Listening on port " + config.Address)
	server.ListenAndServe()
}

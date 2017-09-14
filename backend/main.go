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

	r.HandleFunc("/patient", indexPatient).Methods("GET")
	r.HandleFunc("/patient/{id}", showPatient).Methods("GET")
	r.HandleFunc("/patient/create", createPatient).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/patient/providers", patientProvidersIndex).Methods("GET")

	r.HandleFunc("/provider", indexProvider).Methods("GET")
	r.HandleFunc("/provider/{id}", showProvider).Methods("GET")
	r.HandleFunc("/provider/create", createProvider).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/provider/patients", providerPatientsIndex).Methods("GET")
	r.HandleFunc("/provider/patients/{id}", showProviderPatient).Methods("GET")
	//r.HandleFunc("/provider/appointments/create", createAppointment).Methods("POST").Headers("Content-Type", "application/json")
	//r.HandleFunc("/provider/appointments, indexAppointment).Methods("GET")
	//r.HandleFunc("/provider/appointments/{id}", showAppointment).Methods("GET")
	//r.HandleFunc("/provider/appointments/{id}/cancel", cancelAppointment).Methods("PUT")
	//r.HandleFunc("/provider/appointments/{id}/complete", completeAppointment).Methods("PUT")
	//r.Handlefunc("/provider/appointments/{id}/update", updateAppointment).Methods("POST")

	//Auth API
	r.HandleFunc("/signup", signup).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/auth", authenticate).Methods("POST").Headers("Content-Type", "application/json")
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

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jasondeutsch/previ/backend/data"
	"net/http"
	"strconv"
)

type M map[string]interface{}

/**

Patients Resource

**/

func indexPatient(w http.ResponseWriter, r *http.Request) {
	// TODO Check if user is admin.
	// If user is admin then show all patients.
	// TODO Handle errors
	patients, err := data.Patients()

	m := map[string]interface{}{"error": err != nil, "message": "", "data": patients}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func createPatient(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)

	user, _ := data.UserById(sess.UserId)

	var patient *data.Patient
	err = json.NewDecoder(r.Body).Decode(&patient)
	patient.Id = user.Id
	err = patient.Create(user)

	m := map[string]interface{}{"error": err != nil, "message": "", "data": user}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func showPatient(w http.ResponseWriter, r *http.Request) {
	// TODO authorization
	vars := mux.Vars(r)
	sId, _ := vars["id"]
	id, _ := strconv.Atoi(sId)

	patient, err := data.PatientById(id)

	m := map[string]interface{}{"success": err == nil, "message": "", "data": patient}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)

}

func patientProvidersIndex(w http.ResponseWriter, r *http.Request) {
	sess, _ := session(w, r)
	patient, err := data.PatientById(sess.UserId)
	providers, err := patient.Providers()

	m := map[string]interface{}{"error": err != nil, "message": "", "data": providers}

	w.Header().Set("Contnet-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

/**

Provider Resource

**/

func createProvider(w http.ResponseWriter, r *http.Request) {
	sess, _ := session(w, r)
	user, _ := data.UserById(sess.UserId)

	var provider *data.Provider
	err := json.NewDecoder(r.Body).Decode(&provider)
	provider.Id = user.Id

	err = provider.Create(user)

	m := map[string]interface{}{"error": err != nil, "message": "", "data": nil}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func indexProvider(w http.ResponseWriter, r *http.Request) {
	sess, _ := session(w, r)

	fmt.Println(sess)

	providers, err := data.Providers()
	fmt.Println(err)
	fmt.Println(providers)

	m := map[string]interface{}{"error": err != nil, "message": "", "data": providers}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)

}

func showProvider(w http.ResponseWriter, r *http.Request) {
	sess, _ := session(w, r)

	fmt.Println(sess)

	vars := mux.Vars(r)
	sId, _ := vars["id"]
	id, _ := strconv.Atoi(sId)

	provider, err := data.ProviderById(id)

	m := map[string]interface{}{"error": err != nil, "message": "", "data": provider}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func providerPatientsIndex(w http.ResponseWriter, r *http.Request) {
	sess, _ := session(w, r)
	provider, err := data.ProviderById(sess.UserId)
	patients, err := provider.Patients()

	m := map[string]interface{}{"error": err != nil, "message": "", "data": patients}

	w.Header().Set("Contnet-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func showProviderPatient(w http.ResponseWriter, r *http.Request) {
	var m M

	sess, err := session(w, r)
	provider, err := data.ProviderById(sess.UserId)

	vars := mux.Vars(r)
	id, _ := vars["id"]
	patientId, _ := strconv.Atoi(id)

	if err = provider.HasPatient(patientId); err != nil {
		m = M{"error": true, "message": "", "data": nil}
	} else {
		patient, err := data.PatientById(patientId)
		m = M{"error": err != nil, "message": "", "data": patient}
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

/**

Appointments Resource

**/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jasondeutsch/previ/data"
	"html/template"
	"net/http"
)

/**

Patients Resource

**/

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {

		fmt.Println(sess)
		http.Redirect(w, r, "/login", 401)
	}
}

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

	// sess, _ := session(w, r)
	//	user, _ := sess.User()

	vars := mux.Vars(r)
	id := vars["id"]
	patient, _ := data.PatientById(id)

	files := []string{"templates/layout.html", "templates/private.patient.html"}

	var t *template.Template
	t = template.New("layout")
	t, _ = template.ParseFiles(files...)
	t.ExecuteTemplate(w, "layout", patient)
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

/**

Admin Resource

**/

func adminIndex(w http.ResponseWriter, r *http.Request) {
	sess, _ := session(w, r)
	user, _ := data.UserById(sess.UserId)
	fmt.Println(user)

	patients, err := data.Patients()
	fmt.Println(err)

	files := []string{"templates/layout.html", "templates/admin.index.html", "templates/admin.patient_index.html"}

	var t *template.Template
	t = template.New("layout")
	t, err = template.ParseFiles(files...)
	fmt.Println(err)
	t.ExecuteTemplate(w, "layout", patients)
}

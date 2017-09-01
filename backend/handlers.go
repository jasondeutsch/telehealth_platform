package main

import (
	"encoding/json"
	"fmt"
	"github.com/jasondeutsch/previ/backend/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

/**

Patients Resource

**/

func indexPatient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO Check if user is admin.
	// If user is admin then show all patients.
	// TODO Handle errors

	patients, _ := data.GetAllPatients()

	m := map[string]interface{}{"status": "ok", "message": "", "data": patients}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func createPatient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	sess, err := session(w, r)

	if err != nil {
		return
	}

	var patient *data.Patient
	user, _ := data.UserById(sess.UserId)
	err = patient.Create(user)

	if err != nil {
	}

	return
}

func showPatient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO Check if user is authorized to view patient.
	// TODO Handle err case of GetPatientById

	id := p.ByName("id")
	patient, _ := data.PatientById(id)

	m := map[string]interface{}{"status": "ok", "message": "", "data": patient}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)

}

/**

Auth Routes

**/

type authReponse struct {
	status  string
	message string
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// POST
// /authenticate
func authenticate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	currentRequest := AuthRequest{}

	err := json.NewDecoder(r.Body).Decode(&currentRequest)

	if err != nil {
		http.Error(w, err.Error(), 401)
	}

	user, err := data.UserByEmail(currentRequest.Email)

	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	fmt.Println(user.Password)

	if user.Password == data.Encrypt(currentRequest.Password) {

		_, err := user.CreateSession()
		if err != nil {
			// do stuff
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    strconv.Itoa(user.Id), // may need something better her, UUID?
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		w.Write([]byte("authorized\n"))

	} else {

		w.Write([]byte("not authorized\n"))
	}

}

// POST
// /signup
func signup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user *data.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = user.Create()

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Fprintf(w, "success")
}

// POST
// /logout
func logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cookie, err := r.Cookie("_cookie")

	if err != http.ErrNoCookie {
		id, _ := strconv.Atoi(cookie.Value)

		session := data.Session{Id: id}
		session.Delete()
		fmt.Fprintf(w, "success")
	}
	// handle no cookie
}

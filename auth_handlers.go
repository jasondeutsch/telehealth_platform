package main

import (
	"encoding/json"
	"fmt"
	"github.com/jasondeutsch/previ/backend/data"
	"net/http"
)

type authReponse struct {
	status  string
	message string
}

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// POST
// /auth
func authenticate(w http.ResponseWriter, r *http.Request) {

	// Prepare the JSON response
	m := map[string]interface{}{}

	currentRequest := AuthRequest{}

	err := json.NewDecoder(r.Body).Decode(&currentRequest)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)

		m["error"] = true
		m["message"] = "Could not understand request"
		m["data"] = nil
	}

	user, _ := data.UserByEmail(currentRequest.Email)

	w.Header().Set("Content-Type", "application/json")

	if user.Password == data.Encrypt(currentRequest.Password) {

		sess, err := user.CreateSession()
		if err != nil {
			// do stuff
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    sess.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)

		m["error"] = err != nil
		m["message"] = "authorized"
		m["data"] = nil

	} else {

		w.WriteHeader(http.StatusUnauthorized)

		m["error"] = true
		m["message"] = "authorization failed"
		m["data"] = nil

	}

	json.NewEncoder(w).Encode(m)

}

// POST
// /signup
func signup(w http.ResponseWriter, r *http.Request) {
	var user *data.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}

	err = user.Create()

	m := map[string]interface{}{"error": err != nil, "message": "", "data": ""}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

// POST
// /logout
func logout(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)

	fmt.Println(sess)

	if err != http.ErrNoCookie {
		sess.Delete()
	}

	w.Header().Set("Content-Type", "application/json")
	m := map[string]interface{}{"error": err != nil, "message": "session deleted", "data": nil}
	json.NewEncoder(w).Encode(m)
}

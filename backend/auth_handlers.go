package main

import (
	"encoding/json"
	"github.com/jasondeutsch/previ/backend/data"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
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

	w.Header().Set("Content-Type", "application/json")

	m := map[string]interface{}{}

	if user.Password == data.Encrypt(currentRequest.Password) {

		_, err := user.CreateSession()
		if err != nil {
			// do stuff
		}

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    uuid.NewV4().String(),
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
func signup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var user *data.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}

	err = user.Create()

	m := map[string]interface{}{"error": err != nil, "message": "user created", "data": ""}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

// POST
// /logout
func logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	cookie, err := r.Cookie("_cookie")

	if err != http.ErrNoCookie {
		id, _ := strconv.Atoi(cookie.Value)

		session := data.Session{Id: id}
		session.Delete()
	}

	w.Header().Set("Content-Type", "application/json")
	m := map[string]interface{}{"error": err != nil, "message": "session deleted", "data": nil}
	json.NewEncoder(w).Encode(m)
}

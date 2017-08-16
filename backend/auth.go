package main

import (
	"encoding/json"
	"github.com/jasondeutsch/previ/backend/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// POST /authenticate
func authenticate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	type authRequest struct {
		email    string
		password string
	}

	err := json.NewDecoder(r.Body).Decode(&authRequest)
	if err != nil {
		http.Error(w, err.Error(), 401)
	}

	user, err := data.UserByEmail(authRequest.email)

	if user.Password == data.Encript(authRequest.password) {
		session, err = user.CreateSession()
		if err != nil {
			// do stuff
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    user.Id, // may need something better her, UUID?
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		// send success response to client
	} else {
		// send failure response to client.
	}
}

// POST  /signup
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

}

// login

// logout

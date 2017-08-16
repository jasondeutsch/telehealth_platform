package main

import (
	"encoding/json"
	"fmt"
	"github.com/jasondeutsch/previ/backend/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type authReponse struct {
	status  string
	message string
}

// POST
// /authenticate
func authenticate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	type AuthRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	currentRequest := AuthRequest{}

	err := json.NewDecoder(r.Body).Decode(&currentRequest)

	if err != nil {
		http.Error(w, err.Error(), 401)
	}

	user, err := data.UserByEmail(currentRequest.Email)

	if err != nil {
		fmt.Println("user by email error")
		http.Error(w, err.Error(), 400)
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Println(user.Password)

	if user.Password == data.Encrypt(currentRequest.Password) {

		fmt.Println("Password match!")
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

		fmt.Println("Password is not a match!")
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

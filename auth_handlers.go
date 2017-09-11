package main

import (
	"encoding/json"
	"fmt"
	"github.com/jasondeutsch/previ/data"
	"html/template"
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

//Get
// /login

func login(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html", "templates/login.html"}

	var t *template.Template
	t = template.New("layout")
	t, _ = template.ParseFiles(files...)
	t.ExecuteTemplate(w, "layout", nil)
}

// POST
// /auth
func authenticate(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	user, _ := data.UserByEmail(r.PostFormValue("email"))

	if user.Password == data.Encrypt(r.PostFormValue("password")) {

		sess, _ := user.CreateSession()

		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    sess.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302)
		return
	} else {
		http.Redirect(w, r, "/login", 302)
	}
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

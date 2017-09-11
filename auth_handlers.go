package main

import (
	"encoding/json"
	"fmt"
	"github.com/jasondeutsch/previ/data"
	"html/template"
	"net/http"
)

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

// GET
// /signup
func signup(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html", "templates/signup.html"}

	var t *template.Template
	t = template.New("layout")
	t, _ = template.ParseFiles(files...)
	t.ExecuteTemplate(w, "layout", nil)

}

// POST
// /signupaccount
func signupAccount(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := data.User{
		Email:    r.PostFormValue("email"),
		Password: data.Encrypt(r.PostFormValue("password")),
	}
	err := user.Create()

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/signup", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

// POST
// /logout
func logout(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)

	fmt.Println(sess)

	if err != http.ErrNoCookie {
		sess.Delete()
	}
	http.Redirect(w, r, "/login", 302)

}

package main

import (
	"github.com/jasondeutsch/previ/backend/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()

	newUser := data.User{Email: r.Form["email"][0],
		Password: r.Form["password"][0]}

	newUser.Create()
}

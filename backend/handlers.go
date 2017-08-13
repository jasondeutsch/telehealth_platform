package main

import (
	"encoding/json"
	"fmt"
	"github.com/jasondeutsch/previ/backend/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func signup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var user data.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(user.Email)
	fmt.Println(user.Password)
	fmt.Println(user.Disabled)
	fmt.Println(user.CreatedAt)
}

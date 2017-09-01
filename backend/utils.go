package main

import (
	"encoding/json"
	"github.com/jasondeutsch/previ/backend/data"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Configuration struct {
	Address string
}

var config Configuration

func loadConfig() {
	file, err := os.Open("config.json")

	if err != nil {
		log.Fatalln("Error opening configuration file", err)
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}

	err = decoder.Decode(&config)

	if err != nil {
		log.Fatalln("Error loading config", err)
	}
}

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie")
	if err == nil {
		return
	}

	cookieValue, _ := strconv.Atoi(cookie.Value)

	sess = data.Session{Id: cookieValue}

	ok, _ := sess.Check()

	if !ok {
		return
	}

	return
}

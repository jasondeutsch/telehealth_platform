package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
	"text/template"
)

func getS(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("tomorrow"))
}

func main() {
	router := httprouter.New()

	// Routes
	router.GET("/string", getS)

	// Prefer white list domains with cors.New().Options({AllowedOrigins...})
	cors := cors.Default().Handler(router)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: cors,
	}
	server.ListenAndServe()
}

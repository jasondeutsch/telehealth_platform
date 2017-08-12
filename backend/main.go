package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
)

// dummy route for testing client/server communication.
func getS(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("tomorrow"))
}

func main() {
	router := httprouter.New()

	// Routes
	router.GET("/string", getS)

	router.POST("/signup", signup)

	// Prefer white list domains with cors.New().Options({AllowedOrigins...})
	cors := cors.Default().Handler(router)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: cors,
	}
	server.ListenAndServe()
}

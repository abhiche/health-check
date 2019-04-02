package main

import (
	"log"
	"net/http"
	"os"

	"github.com/globalsign/mgo"
	"github.com/gorilla/handlers"

	"github.com/abhiche/health-check/pkg/site"
)

var mongoURL = "MONGO_URL"

func main() {
	var connString = os.Getenv(mongoURL)
	if connString == "" {
		panic("Missing env var " + mongoURL)
	}
	session, err := mgo.Dial(os.Getenv(mongoURL))
	if err != nil {
		panic(err)
	}
	router := site.NewRouter(session)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PATCH"})

	// launch server
	log.Fatal(http.ListenAndServe(":9000",
		handlers.CORS(allowedOrigins, allowedMethods)(router)))
}

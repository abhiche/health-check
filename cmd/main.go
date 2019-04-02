package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/globalsign/mgo"
	"github.com/gorilla/handlers"

	"github.com/abhiche/health-check/pkg/site"
)

var mongoURL = "MONGO_URL"
var defaultPort = "9000"

func main() {
	var connString = os.Getenv(mongoURL)
	var port = os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
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

	s := &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Minute,
		WriteTimeout:   10 * time.Minute,
		MaxHeaderBytes: 0,
	}

	// launch server
	log.Fatal(s.ListenAndServe(),
		handlers.CORS(allowedOrigins, allowedMethods)(router))
}

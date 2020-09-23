package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	r = mux.NewRouter()
)

func Start() {
	configureRoutes()

	s := &http.Server{
		Addr:         "127.0.0.1:8082",
		Handler:      r,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}

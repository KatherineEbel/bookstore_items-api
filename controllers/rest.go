package controllers

import (
	"net/http"
)

type iRestController interface {
	New(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

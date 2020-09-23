package controllers

import (
	"fmt"
	"net/http"
)

const (
	pong = "pong"
)

var (
	Ping iPingController = &pingController{}
)

type iPingController interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type pingController struct{}

func (c pingController) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte(pong)); err != nil {
		fmt.Println(err)
	}
}

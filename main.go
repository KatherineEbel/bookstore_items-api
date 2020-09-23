package main

import (
	"github.com/gorilla/mux"

	"github.com/KatherineEbel/bookstore_items-api/app"
)

var (
	r = mux.NewRouter()
)

func main() {
	app.Start()
}

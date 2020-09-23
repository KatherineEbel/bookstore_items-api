package app

import (
	"net/http"

	"github.com/KatherineEbel/bookstore_items-api/controllers"
)

func configureRoutes() {
	r.HandleFunc("/ping", controllers.Ping.Ping)
	r.HandleFunc("/items", controllers.Items.New).Methods(http.MethodPost)
}

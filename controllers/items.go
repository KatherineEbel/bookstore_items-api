package controllers

import (
	"fmt"
	"net/http"

	"github.com/KatherineEbel/oauth-go/oauth"

	"github.com/KatherineEbel/bookstore_items-api/domain/items"
	items2 "github.com/KatherineEbel/bookstore_items-api/services/items"
)

var (
	Items iRestController = &itemsController{}
)

type itemsController struct {
}

func (c *itemsController) New(w http.ResponseWriter, r *http.Request) {

}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	if err := oauth.Authenticate(r); err != nil {

		// TODO: Return 401 status & json error
		return
	}
	item := items.Item{
		Seller: oauth.GetUserId(r),
	}
	res, err := items2.ItemsService.Add(item)
	if err != nil {
		// TODO: Return json error
		return
	}
	// TODO: Return 201 status & item json
	fmt.Printf("Result: %v\n", res)
}

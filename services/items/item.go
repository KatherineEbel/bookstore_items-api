package items

import (
	errors2 "errors"

	"github.com/KatherineEbel/bookstore-utils-go/rest/errors"

	"github.com/KatherineEbel/bookstore_items-api/domain/items"
)

var (
	ItemsService IService = &itemService{}
)

type IService interface {
	Add(item items.Item) (*items.Item, *errors.RestError)
	Fetch(string) (*items.Item, *errors.RestError)
}

type itemService struct {
}

func (s *itemService) Fetch(id string) (*items.Item, *errors.RestError) {
	return nil, errors.NewInternalServerError("internal server error", errors2.New("not implemented yet"))
}

func (s *itemService) Add(item items.Item) (*items.Item, *errors.RestError) {
	return nil, errors.NewInternalServerError("internal server error", errors2.New("not implemented yet"))
}

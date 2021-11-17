package services

import (
	"github.com/Lozovoi-Rodion/bookstore_items-api/domain/items"
	"github.com/Lozovoi-Rodion/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct {
}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	panic("implement me")
}

func (s *itemsService) Get(itemId string) (*items.Item, *rest_errors.RestErr) {
	panic("implement me")
}

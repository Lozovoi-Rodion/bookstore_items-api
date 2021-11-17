package controllers

import (
	"fmt"
	"github.com/Lozovoi-Rodion/bookstore-oauth-go/oauth"
	"github.com/Lozovoi-Rodion/bookstore_items-api/domain/items"
	"github.com/Lozovoi-Rodion/bookstore_items-api/services"
	"net/http"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct {
}

func (c *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO: Return error json to the user
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		// TODO: Return error json to the user
	}

	fmt.Println(result)
	//TODO: Return create item as json, 201
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
}

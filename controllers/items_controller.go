package controllers

import (
	"github.com/Lozovoi-Rodion/bookstore-oauth-go/oauth"
	"github.com/Lozovoi-Rodion/bookstore_items-api/domain/items"
	"github.com/Lozovoi-Rodion/bookstore_items-api/services"
	"github.com/Lozovoi-Rodion/bookstore_items-api/utils/http_utils"
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
		http_utils.RespondError(w, err)
		return
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		http_utils.RespondError(w, err)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (c *itemsController) Get(w http.ResponseWriter, r *http.Request) {
}

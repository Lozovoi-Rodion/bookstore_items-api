package items

import (
	"errors"
	"github.com/Lozovoi-Rodion/bookstore_items-api/clients/elasticsearch"
	"github.com/Lozovoi-Rodion/bookstore_utils-go/rest_errors"
)

const (
	indexItems = "items"
)

func (i *Item) Save() *rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(indexItems, i)

	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save an item", errors.New("database error"))
	}

	i.Id = result.Id
	return nil
}

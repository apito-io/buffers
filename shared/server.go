package shared

import (
	"encoding/json"
	"github.com/apito-cms/buffers/interfaces"
	"github.com/graph-gophers/dataloader/v7"
)

type GraphQLServer struct {
	ProjectDriver interfaces.ProjectDBInterface
	Dataloaders   DataLoaders
}

// DataLoaders Dataloaders
type DataLoaders struct {
	MultiLoader *dataloader.Loader[string, interface{}]
	//SingleLoader *dataloader.Loader[string, interface{}]
}

type Response struct {
	Data       interface{}            `json:"data,omitempty"`
	Errors     json.RawMessage        `json:"errors,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

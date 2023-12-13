package interfaces

import (
	"encoding/json"
	"github.com/apito-cms/buffers/protobuff"
	"github.com/graph-gophers/dataloader/v7"
)

type GraphQLExecutorInterface interface {
	SetProjectDriverAndParam(_project *protobuff.Project) error
	InitDataloaders(_func dataloader.BatchFunc[string, interface{}])

	GetProjectDriver() ProjectDBInterface
	SetProjectDriver(driver ProjectDBInterface)

	GetDataloaders() DataLoaders
	SetDataloaders(loader DataLoaders)
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

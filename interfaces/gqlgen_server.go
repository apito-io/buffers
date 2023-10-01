package interfaces

import (
	"context"
	"encoding/json"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/labstack/echo/v4"
	"github.com/vektah/gqlparser/v2/ast"
)

// Dynamic Dataloaders
type DataLoaders struct {
	MultiLoader  *dataloader.Loader[string, interface{}]
	SingleLoader *dataloader.Loader[string, interface{}]
}

// Server to Gql connect functions
type GqlServer interface {
	SolvePublicQuery(ctx context.Context, model string, _args map[string]interface{}, selectionSet *ast.SelectionSet, router echo.Context) ([]byte, error)
	SolvePublicQueryCount(ctx context.Context, model string, _args map[string]interface{}, router echo.Context) ([]byte, error)
}

// Marshal & Unmarshal []byte result

type Result[T any] struct {
	Result  []*T `json:"result"`
	Count   int  `json:"count"`
	Cached  bool `json:"cached"`
	HasMore bool `json:"hasMore"`
	Error   bool `json:"error"`
	Code    int  `json:"code"`
}

func UnmarshalDynamicCount[T any](body []byte) (*T, error) {
	var res Result[T]
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res.Result[0], nil
}

func UnmarshalDynamicResults[T any](body []byte) ([]*T, error) {
	var res Result[T]
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res.Result, nil
}

func UnmarshalDynamicResult[T any](body []byte) (*T, error) {
	var res Result[T]
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res.Result[0], nil
}

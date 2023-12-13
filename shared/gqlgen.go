package shared

import (
	"context"
	"encoding/json"
	"github.com/apito-cms/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

type SchemaLoaderInterface interface {
	SwitchSchema(projectId string, server GqlServer) graphql.ExecutableSchema
	LoadedSchemaName() []string
}

// GqlServer to Gql connect functions
type GqlServer interface {
	SolvePublicQuery(ctx context.Context, model string, _args interface{}, selectionSet *ast.SelectionSet, cache ApplicationCache) ([]byte, error)
	SolvePublicQueryCount(ctx context.Context, model string, _args interface{}, cache ApplicationCache) ([]byte, error)
	SolvePublicMutation(ctx context.Context, resolverName string, _id *string, _ids []*string, status *string, local *string, userInputPayload interface{}, connect interface{}, disconnect interface{}, cache ApplicationCache) ([]byte, error)
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

type DataloaderResult[T any] struct {
	Result  []T  `json:"result"`
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
	if res.Count > 0 {
		return res.Result[0], nil
	}
	return nil, nil
}

func UnmarshalDynamicResults[T any](body []byte) ([]*T, error) {
	var res Result[T]
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	if res.Count > 0 {
		return res.Result, nil
	}
	return nil, nil
}

func UnmarshalDynamicResult[T any](body []byte) (*T, error) {
	var res Result[T]
	err := json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	if res.Count > 0 {
		return res.Result[0], nil
	}
	return nil, nil
}

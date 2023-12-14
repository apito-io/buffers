package interfaces

import (
	"context"
	"github.com/apito-cms/buffers/protobuff"
	"github.com/apito-cms/buffers/shared"
	"github.com/apito-cms/gqlgen/graphql"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/vektah/gqlparser/v2/ast"
)

type GraphQLExecutorInterface interface {
	SetProjectDriverAndParam(_project *protobuff.Project) error
	InitDataloaders(_func dataloader.BatchFunc[string, interface{}])

	GetProjectDriver() ProjectDBInterface
	SetProjectDriver(driver ProjectDBInterface)

	GetDataloaders() shared.DataLoaders
	SetDataloaders(loader shared.DataLoaders)

	SolvePublicQuery(ctx context.Context, model string, _args interface{}, selectionSet *ast.SelectionSet, cache *shared.ApplicationCache) ([]byte, error)
	SolvePublicQueryCount(ctx context.Context, model string, _args interface{}, cache *shared.ApplicationCache) ([]byte, error)
	SolvePublicMutation(ctx context.Context, resolverName string, _id *string, _ids []*string, status *string, local *string, userInputPayload interface{}, connect interface{}, disconnect interface{}, cache *shared.ApplicationCache) ([]byte, error)
}

type SchemaLoaderInterface interface {
	SwitchSchema(projectId string, server GraphQLExecutorInterface) graphql.ExecutableSchema
	LoadedSchemaName() []string
}

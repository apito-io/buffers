package interfaces

import (
	"context"
	"github.com/apito-cms/buffers/protobuff"
	"github.com/apito-cms/buffers/shared"
	"github.com/apito-cms/gqlgen/codegen/config"
	"github.com/apito-cms/gqlgen/graphql"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/vektah/gqlparser/v2/ast"
)

type GraphQLExecutorInterface interface {
	Init(ctx context.Context, _driver *protobuff.InitParams) error

	CodeGeneration(ctx context.Context, projectId string, _cfg config.Config) error

	GetExecutorVersion(ctx context.Context) (string, error)

	SetApplicationCache(ctx context.Context, cache *shared.ApplicationCache)
	GetApplicationCache(ctx context.Context) *shared.ApplicationCache

	GetProjectDriver(ctx context.Context) ProjectDBInterface
	SetProjectDriver(ctx context.Context, driver ProjectDBInterface)

	GetSharedDBDriver(ctx context.Context) SharedDBInterface
	SetSharedDBDriver(ctx context.Context, driver SharedDBInterface)

	GetDataloaders(ctx context.Context) *shared.DataLoaders
	DataLoaderHandlr(ctx context.Context, keys []string) []*dataloader.Result[interface{}]

	SolvePublicQuery(ctx context.Context, model string, _args interface{}, selectionSet *ast.SelectionSet, cache *shared.ApplicationCache) ([]byte, error)
	SolvePublicQueryCount(ctx context.Context, model string, _args interface{}, cache *shared.ApplicationCache) ([]byte, error)
	SolvePublicMutation(ctx context.Context, resolverName string, _id *string, _ids []*string, status *string, local *string, userInputPayload interface{}, connect interface{}, disconnect interface{}, cache *shared.ApplicationCache) ([]byte, error)

	ConnectDisconnectParamBuilder(ctx context.Context, schema *protobuff.ProjectSchema, uid string, collectionName string, connectionIds map[string]interface{}, modelType *protobuff.ModelType) ([]*shared.ConnectDisconnectParam, error)
	HandlePayloadFormatting(ctx context.Context, param *shared.CommonSystemParams, isFaker bool, local string, fields []*protobuff.FieldInfo, inputPayload map[string]interface{}, dbPayload map[string]interface{}) (map[string]interface{}, error)

	UploadImageFromURL(ctx context.Context, projectId, modelName, imageUrl string) (*protobuff.FileDetails, error)
	HandleMediaURL(ctx context.Context, param *shared.CommonSystemParams, media map[string]interface{}) (interface{}, error)
}

type SchemaLoaderInterface interface {
	SwitchSchema(ctx context.Context, projectId string, server GraphQLExecutorInterface) graphql.ExecutableSchema
	LoadedSchemaName(ctx context.Context) []string
}

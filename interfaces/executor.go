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
	Init(_driver *protobuff.InitParams) error

	GetExecutorVersion() (string, error)
	GetMicroServicePort(id string) (string, error)

	SavePID(projectId, pid string) error
	GetPID(projectId string) (string, error)

	SetApplicationCache(cache *shared.ApplicationCache)
	GetApplicationCache() *shared.ApplicationCache

	GetProjectDriver() ProjectDBInterface
	SetProjectDriver(driver ProjectDBInterface)

	GetSharedDBDriver() SharedDBInterface
	SetSharedDBDriver(driver SharedDBInterface)

	GetDataloaders() *shared.DataLoaders
	DataLoaderHandlr(ctx context.Context, keys []string) []*dataloader.Result[interface{}]

	SolvePublicQuery(ctx context.Context, model string, _args interface{}, selectionSet *ast.SelectionSet, cache *shared.ApplicationCache) ([]byte, error)
	SolvePublicQueryCount(ctx context.Context, model string, _args interface{}, cache *shared.ApplicationCache) ([]byte, error)
	SolvePublicMutation(ctx context.Context, resolverName string, _id *string, _ids []*string, status *string, local *string, userInputPayload interface{}, connect interface{}, disconnect interface{}, cache *shared.ApplicationCache) ([]byte, error)

	ConnectDisconnectParamBuilder(ctx context.Context, schema *protobuff.ProjectSchema, uid string, collectionName string, connectionIds map[string]interface{}, modelType *protobuff.ModelType) ([]*shared.ConnectDisconnectParam, error)
	HandlePayloadFormatting(ctx context.Context, param *shared.CommonSystemParams, isFaker bool, local string, fields []*protobuff.FieldInfo, inputPayload map[string]interface{}, dbPayload map[string]interface{}) (map[string]interface{}, error)

	UploadImageFromURL(projectId, modelName, imageUrl string) (*protobuff.FileDetails, error)
	HandleMediaURL(ctx context.Context, param *shared.CommonSystemParams, media map[string]interface{}) (interface{}, error)
}

type SchemaLoaderInterface interface {
	SwitchSchema(projectId string, server GraphQLExecutorInterface) graphql.ExecutableSchema
	LoadedSchemaName() []string
}

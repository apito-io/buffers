package interfaces

import (
	"context"
	"github.com/apito-cms/gqlgen/graphql"
)

type SchemaLoaderInterface interface {
	SwitchSchema(ctx context.Context, server GraphQLExecutorInterface) graphql.ExecutableSchema
	LoadedSchemaName(ctx context.Context) []string
}

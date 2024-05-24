package interfaces

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
)

type SchemaLoaderInterface interface {
	SwitchSchema(ctx context.Context, server GraphQLExecutorInterface) graphql.ExecutableSchema
	LoadedSchemaName(ctx context.Context) []string
}

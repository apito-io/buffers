package interfaces

import (
	"context"
	"github.com/apito-cms/buffers/protobuff"
	"github.com/apito-cms/buffers/shared"
)

type CacheDB interface {
	GetProject(ctx context.Context, projectId string) (*protobuff.Project, error)
	SaveProject(ctx context.Context, project *protobuff.Project) (*protobuff.Project, error)

	ListKeys(ctx context.Context) ([]string, error)

	PutAppCache(ctx context.Context, projectId string, cache *shared.ApplicationCache) error
	GetAppCache(ctx context.Context, projectId string) (*shared.ApplicationCache, error)
	Expire(ctx context.Context, id string) error

	Put(ctx context.Context, id string, cache interface{}) error
	Get(ctx context.Context, id string) (interface{}, error)
}

package interfaces

import (
	"context"

	"github.com/apito-io/buffers/protobuff"
	"github.com/apito-io/buffers/shared"
)

// SystemDBInterface is the main interface for the system database operations
type SystemDBInterface interface {

	// RunMigration runs the database migrations
	RunMigration(ctx context.Context) error

	// GetSystemUser retrieves a system user by their ID
	GetSystemUser(ctx context.Context, id string) (*protobuff.SystemUser, error)

	// GetSystemUserByUsername retrieves a system user by their username
	GetSystemUserByUsername(ctx context.Context, username string) (*protobuff.SystemUser, error)

	// UpdateSystemUser updates a system user's profile
	UpdateSystemUser(ctx context.Context, user *protobuff.SystemUser, replace bool) error

	// SearchResource searches for system users based on a filter
	SearchResource(ctx context.Context, param *shared.CommonSystemParams) (*shared.SearchResponse[any], error)

	// GetProject retrieves a project by its ID
	GetProject(ctx context.Context, id string) (*protobuff.Project, error)

	// ListFunctions lists all the cloud functions for a given project
	ListFunctions(ctx context.Context, param *shared.CommonSystemParams) (*shared.SearchResponse[protobuff.CloudFunction], error)

	// UpdateProject updates a project
	UpdateProject(ctx context.Context, project *protobuff.Project, replace bool) error

	// CheckTokenBlacklisted checks if a token is blacklisted
	CheckTokenBlacklisted(ctx context.Context, tokenId string) error

	// BlacklistAToken blacklists a token
	BlacklistAToken(ctx context.Context, token map[string]interface{}) error

	// DeleteProjectFromSystem deletes a project from the system
	DeleteProjectFromSystem(ctx context.Context, projectId string) error

	// SaveRawData saves raw data related to payments
	SaveRawData(ctx context.Context, collection string, data map[string]interface{}) error
}

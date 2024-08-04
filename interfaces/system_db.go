package interfaces

import (
	"context"

	"github.com/apito-io/buffers/protobuff"
	"github.com/apito-io/buffers/shared"
)

// SystemDBInterface is the main interface for the system database operations
type SystemDBInterface interface {

	// RunMigration runs the database migrations
	RunMigration() error

	// GetOrganizations retrieves multiple user organizations by their IDs
	GetOrganizations(ctx context.Context, userId string) (*shared.SearchResponse[protobuff.Organization], error)

	// ListTeams lists all the teams for a given project
	//ListTeams(ctx context.Context, param *shared.CommonSystemParams) ([]*protobuff.SystemUser, error)

	// GetSystemUser retrieves a system user by their ID
	GetSystemUser(ctx context.Context, id string) (*protobuff.SystemUser, error)

	GetSystemUserByUsername(ctx context.Context, username string) (*protobuff.SystemUser, error)

	// FindOrganizationAdmin retrieves an organization admin by their ID
	FindOrganizationAdmin(ctx context.Context, orgId string) (*protobuff.SystemUser, error)

	// GetSystemUsers retrieves multiple system users by their IDs
	//GetSystemUsers(ctx context.Context, keys []string) (map[string]*protobuff.SystemUser, error)

	// CreateSystemUser creates a new system user
	CreateSystemUser(ctx context.Context, user *protobuff.SystemUser) (*protobuff.SystemUser, error)

	// UpdateSystemUser updates a system user's profile
	UpdateSystemUser(ctx context.Context, user *protobuff.SystemUser, replace bool) error

	// SearchResource searches for system users based on a filter
	SearchResource(ctx context.Context, param *shared.CommonSystemParams) (*shared.SearchResponse[any], error)

	// SearchUsers searches for system users based on a filter
	SearchUsers(ctx context.Context, param *shared.CommonSystemParams) (*shared.SearchResponse[protobuff.SystemUser], error)

	// GetProjectWithRolesAndPermission lists all the projects for a given user
	//GetProjectWithRolesAndPermission(ctx context.Context, userId string) ([]*protobuff.ProjectWithRoles, error)

	// ListAllUsers lists all system users
	//ListAllUsers(ctx context.Context) ([]*protobuff.SystemUser, error)

	// GetProject retrieves a project by its ID
	GetProject(ctx context.Context, id string) (*protobuff.Project, error)

	// CheckProjectName checks if a project name already exists
	CheckProjectName(ctx context.Context, name string) error

	// ListProjects lists all the projects for a given user
	ListProjects(ctx context.Context, param *shared.CommonSystemParams) (*shared.SearchResponse[protobuff.Project], error)

	// ListAllProjects lists all the projects for a given user
	//ListAllProjects(ctx context.Context, userId string) ([]*protobuff.Project, error)

	// ListFunctions lists all the cloud functions for a given project
	ListFunctions(ctx context.Context, param *shared.CommonSystemParams) (*shared.SearchResponse[protobuff.CloudFunction], error)

	// AddSystemUserMetaInfo adds metadata to a system user
	AddSystemUserMetaInfo(ctx context.Context, doc *shared.DefaultDocumentStructure) (*shared.DefaultDocumentStructure, error)

	// CreateProject creates a new project
	CreateProject(ctx context.Context, project *protobuff.Project) (*protobuff.Project, error)

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

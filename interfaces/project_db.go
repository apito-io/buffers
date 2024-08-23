package interfaces

import (
	"context"

	"github.com/apito-io/buffers/protobuff"
	"github.com/apito-io/buffers/shared"
)

// ProjectDBInterface represents the interface for interacting with the project database.
type ProjectDBInterface interface {
	// RunMigration runs the migration for a specific project.
	RunMigration(ctx context.Context, projectId string) error

	// DeleteProject deletes a project from the database.
	DeleteProject(ctx context.Context, projectId string) error

	// CheckCollectionExists checks if a collection exists for a specific project.
	CheckCollectionExists(ctx context.Context, projectId string) (bool, error)

	// AddCollection adds a new collection to a project.
	AddCollection(ctx context.Context, projectName string) (*string, error)

	// TransferProject transfers a project from one user to another.
	TransferProject(ctx context.Context, userId, from, to string) error

	// AddModel adds a new model to a project.
	AddModel(ctx context.Context, project *protobuff.Project, name string, singleRecord bool) (*protobuff.ProjectSchema, error)

	// AddFieldToModel adds a new field to a model.
	AddFieldToModel(ctx context.Context, param shared.CommonSystemParams, isUpdate bool, repeatedGroupIdentifier *string) (*protobuff.ModelType, error)

	// AddRelationFields creates a relation field (has one or has many) between two models.
	AddRelationFields(ctx context.Context, from *protobuff.ConnectionType, to *protobuff.ConnectionType) error

	// DropConnections drops the connections between two models, including pivot tables, relation keys, and collection tables.
	// This operation is not reversible.
	DropConnections(ctx context.Context, projectId string, from *protobuff.ConnectionType, to *protobuff.ConnectionType) error

	// ConnectBuilder connects a builder to a project.
	ConnectBuilder(ctx context.Context, param shared.CommonSystemParams) error

	// DisconnectBuilder disconnects a builder from a project.
	DisconnectBuilder(ctx context.Context, param shared.CommonSystemParams) error

	// GetProjectUser retrieves the user profile for a specific project.
	GetProjectUser(ctx context.Context, phone, email, projectId string) (*shared.DefaultDocumentStructure, error)

	// GetLoggedInProjectUser retrieves the user profile for the logged-in user in a specific project.
	GetLoggedInProjectUser(ctx context.Context, param *shared.CommonSystemParams) (*shared.DefaultDocumentStructure, error)

	// GetProjectUsers retrieves the metadata for multiple users in a project.
	GetProjectUsers(ctx context.Context, projectId string, keys []string) (map[string]*shared.DefaultDocumentStructure, error)

	// GetAllRelationDocumentsOfSingleDocument retrieves all relation data of a single document by its ID.
	GetAllRelationDocumentsOfSingleDocument(ctx context.Context, from string, arg *shared.CommonSystemParams) (interface{}, error)

	// GetSingleProjectDocumentBytes retrieves the byte representation of a single project document by its ID.
	GetSingleProjectDocumentBytes(ctx context.Context, param shared.CommonSystemParams) ([]byte, error)

	// GetSingleProjectDocument retrieves a single project document by its ID.
	GetSingleProjectDocument(ctx context.Context, param shared.CommonSystemParams) (*shared.DefaultDocumentStructure, error)

	// GetSingleProjectDocumentRevisions retrieves the revision history of a single project document by its ID.
	GetSingleProjectDocumentRevisions(ctx context.Context, param shared.CommonSystemParams) ([]*shared.DocumentRevisionHistory, error)

	// GetSingleRawDocumentFromProject retrieves a single raw document from a project by its ID.
	GetSingleRawDocumentFromProject(ctx context.Context, param shared.CommonSystemParams) (interface{}, error)

	// QueryMultiDocumentOfProjectBytes queries multiple documents in a project and retrieves the byte representation of the result.
	QueryMultiDocumentOfProjectBytes(ctx context.Context, param shared.CommonSystemParams) ([]byte, error)

	// QueryMultiDocumentOfProject queries multiple documents in a project and retrieves the result as a list of DefaultDocumentStructure.
	QueryMultiDocumentOfProject(ctx context.Context, param shared.CommonSystemParams) ([]*shared.DefaultDocumentStructure, error)

	// CountMultiDocumentOfProject counts the number of documents in a project that match the given parameters.
	CountMultiDocumentOfProject(ctx context.Context, param shared.CommonSystemParams, previewModel bool) (int, error)

	// AddDocumentToProject adds a document to a project.
	AddDocumentToProject(ctx context.Context, projectId string, modelName string, doc *shared.DefaultDocumentStructure) (interface{}, error)

	// UpdateDocumentOfProject updates a specific document in a project.
	UpdateDocumentOfProject(ctx context.Context, param shared.CommonSystemParams, doc *shared.DefaultDocumentStructure, replace bool) error

	// DeleteDocumentFromProject deletes a specific document from a project.
	DeleteDocumentFromProject(ctx context.Context, param shared.CommonSystemParams) error

	// DeleteDocumentRelation deletes the relations or data in pivot tables associated with a document in a project.
	DeleteDocumentRelation(ctx context.Context, param shared.CommonSystemParams) error

	// DeleteDocumentsFromProject deletes multiple documents from a project.
	DeleteDocumentsFromProject(ctx context.Context, param shared.CommonSystemParams) error

	// CreateRelation creates a relation between two models in a project.
	CreateRelation(ctx context.Context, projectId string, relation *shared.EdgeRelation) error

	// DeleteRelation deletes a relation between two models in a project.
	DeleteRelation(ctx context.Context, param *shared.ConnectDisconnectParam, id string) error

	// NewInsertableRelations returns a list of insertable relations for a specific model in a project.
	NewInsertableRelations(ctx context.Context, param *shared.ConnectDisconnectParam) ([]string, error)

	// CheckOneToOneRelationExists checks if a one-to-one relation exists between two models in a project.
	CheckOneToOneRelationExists(ctx context.Context, param *shared.ConnectDisconnectParam) (bool, error)

	// GetRelationIds retrieves the IDs of all documents related to a specific document in a project.
	GetRelationIds(ctx context.Context, param *shared.ConnectDisconnectParam) ([]string, error)

	// RelationshipDataLoader loads relationship data for a specific document in a project.
	RelationshipDataLoader(ctx context.Context, param *shared.CommonSystemParams, connection map[string]interface{}) (interface{}, error)

	// RelationshipDataLoaderBytes loads relationship data for a specific document in a project and returns the byte representation of the result.
	RelationshipDataLoaderBytes(ctx context.Context, param *shared.CommonSystemParams, connection map[string]interface{}) ([]byte, error)

	// CountDocOfProject counts the number of documents in a project that match the given parameters.
	CountDocOfProject(ctx context.Context, param *shared.CommonSystemParams) (interface{}, error)

	// CountDocOfProjectBytes counts the number of documents in a project that match the given parameters and returns the byte representation of the result.
	CountDocOfProjectBytes(ctx context.Context, param *shared.CommonSystemParams) ([]byte, error)

	// DropField drops a field and its associated data from a project.
	DropField(ctx context.Context, param shared.CommonSystemParams) error

	// RenameField renames a field in a model along with its data key.
	RenameField(ctx context.Context, oldFieldName string, repeatedFieldGroup *string, param shared.CommonSystemParams) error

	// --- MISC Functions for plugins and other services ---

	// CreateTableOrCollection creates a table for a specific project.
	CreateTableOrCollection(ctx context.Context, name string, properties map[string]string) error

	// DropTableOrCollection drops a table for a specific project.
	DropTableOrCollection(ctx context.Context, name string) error

	// AddDataToTableOrCollection adds data to a specific project table.
	AddDataToTableOrCollection(ctx context.Context, table string, data map[string]interface{}) error

	// UpdateDataToTableOrCollection updates data in a specific project table.
	UpdateDataToTableOrCollection(ctx context.Context, table string, data map[string]interface{}) (interface{}, error)

	// DeleteDataFromTableOrCollection deletes data from a specific project table.
	DeleteDataFromTableOrCollection(ctx context.Context, table string, id string) error
}

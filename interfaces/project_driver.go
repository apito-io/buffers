package interfaces

import (
	"context"
	"github.com/apito-cms/buffers/protobuff"
	"github.com/apito-cms/buffers/shared"
	"github.com/graph-gophers/dataloader"
	"github.com/tailor-inc/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

type ApitoProjectDB interface {
	DeleteProject(ctx context.Context, projectId string) error

	// CheckProjectExists Project Structure Related
	CheckProjectExists(ctx context.Context, projectId string) (bool, error)

	AddCollection(ctx context.Context, projectName string) (*string, error)

	TransferProject(ctx context.Context, userId, from, to string) error

	AddModel(ctx context.Context, project *protobuff.Project, name string, singleRecord bool) (*protobuff.ProjectSchema, error)

	AddFieldToModel(ctx context.Context, param shared.CommonSystemParams, isUpdate bool, repeatedGroupIdentifier *string) (*protobuff.ModelType, error)

	// AddRelationFields Creates a relation field has one or has many
	AddRelationFields(ctx context.Context, from *protobuff.ConnectionType, to *protobuff.ConnectionType) error

	// DropConnections Drop pivot tables, relation key or collection table and all document with it
	// NOT REVERSIBLE
	DropConnections(ctx context.Context, projectId string, from *protobuff.ConnectionType, to *protobuff.ConnectionType) error

	ConnectBuilder(ctx context.Context, param shared.CommonSystemParams) error

	DisconnectBuilder(ctx context.Context, param shared.CommonSystemParams) error

	//AddAuthAddOns(project *protobuff.Project, auth map[string]interface{}) error

	//RemoveAuthAddOns(project *protobuff.Project, option map[string]interface{}) error

	// GetProjectUser Get a user Profile
	GetProjectUser(ctx context.Context, phone, email, projectId string) (*shared.DefaultDocumentStructure, error)

	GetLoggedInProjectUser(ctx context.Context, param *shared.CommonSystemParams) (*shared.DefaultDocumentStructure, error)

	// GetProjectUsers meta data loader
	GetProjectUsers(ctx context.Context, projectId string, keys []string) (map[string]*shared.DefaultDocumentStructure, error)

	// GetAllPreviewDocumentsByModel Get all data by a model name
	//GetAllPreviewDocumentsByModel(param database.CommonSystemParams) ([]*protobuff.PreviewMode, error)

	// Get a Relation Data of a single document by id, it could be object or array
	GetAllRelationDocumentsOfSingleDocument(ctx context.Context, from string, arg *shared.CommonSystemParams) (interface{}, error)

	// List media data
	ListMedias(ctx context.Context, projectId string, param *graphql.ResolveParams) ([]*protobuff.FileDetails, error)

	// Count media data
	CountMedias(ctx context.Context, projectId string, param *graphql.ResolveParams) (int, error)

	// List media data
	DeleteMediaFile(ctx context.Context, param shared.CommonSystemParams) error

	// GetSingleProjectDocument Get a single Project document by id
	GetSingleProjectDocumentBytes(ctx context.Context, param shared.CommonSystemParams, s *ast.SelectionSet) ([]byte, error)

	// GetSingleProjectDocument Get a single Project document by id
	GetSingleProjectDocument(ctx context.Context, param shared.CommonSystemParams, s *ast.SelectionSet) (*shared.DefaultDocumentStructure, error)

	// GetSingleProjectDocumentRevisions Get a single Project document by id
	GetSingleProjectDocumentRevisions(ctx context.Context, param shared.CommonSystemParams) ([]*shared.DocumentRevisionHistory, error)

	// Same function as above just returns an interface, used in hooks doc
	GetSingleRawDocumentFromProject(ctx context.Context, param shared.CommonSystemParams) (interface{}, error)

	// QueryMultiDocumentOfProject Query Multiple Documents
	QueryMultiDocumentOfProjectBytes(ctx context.Context, param shared.CommonSystemParams, s *ast.SelectionSet) ([]byte, error)

	// QueryMultiDocumentOfProject Query Multiple Documents
	QueryMultiDocumentOfProject(ctx context.Context, param shared.CommonSystemParams, s *ast.SelectionSet) ([]*shared.DefaultDocumentStructure, error)

	CountMultiDocumentOfProject(ctx context.Context, param shared.CommonSystemParams, previewModel bool) (int, error)

	AddTeamMetaInfo(ctx context.Context, docs []*protobuff.UserMeta) ([]*protobuff.UserMeta, error)

	// -- Mutation --- //

	// Create a Media Document
	CreateMediaDocument(ctx context.Context, projectId string, media *protobuff.FileDetails) (*protobuff.FileDetails, error)

	// Create or Add a Document to project
	AddDocumentToProject(ctx context.Context, projectId string, modelName string, doc *shared.DefaultDocumentStructure) (interface{}, error)

	// Update a particular Document
	UpdateDocumentOfProject(ctx context.Context, param shared.CommonSystemParams, doc *shared.DefaultDocumentStructure, replace bool) error

	// Delete Doc from Project
	DeleteDocumentFromProject(ctx context.Context, param shared.CommonSystemParams) error

	// delete all the relations or data in pivot table from the project
	DeleteDocumentRelation(ctx context.Context, param shared.CommonSystemParams) error

	// Delete Doc from Project
	DeleteDocumentsFromProject(ctx context.Context, param shared.CommonSystemParams) error

	// Relation builder
	CreateRelation(ctx context.Context, projectId string, relation *shared.EdgeRelation) error

	// Relation builder
	DeleteRelation(ctx context.Context, param *shared.ConnectDisconnectParam, id string) error

	// !!! arent they same ?
	NewInsertableRelations(ctx context.Context, param *shared.ConnectDisconnectParam) ([]string, error)

	// !!
	CheckOneToOneRelationExists(ctx context.Context, param *shared.ConnectDisconnectParam) (bool, error)

	// Get ids of every document related to a document
	GetRelationIds(ctx context.Context, param *shared.ConnectDisconnectParam) ([]string, error)

	//
	RelationshipDataLoader(ctx context.Context, param *shared.CommonSystemParams, s *ast.SelectionSet, connection map[string]interface{}, keys []string) ([]*dataloader.Result, error)

	//
	RelationshipDataLoaderBytes(ctx context.Context, param *shared.CommonSystemParams, s *ast.SelectionSet, connection map[string]interface{}, keys []string) ([]byte, error)

	// Doc count resolver
	CountDocOfProject(ctx context.Context, param *shared.CommonSystemParams) (interface{}, error)

	// Doc count resolver
	CountDocOfProjectBytes(ctx context.Context, param *shared.CommonSystemParams) ([]byte, error)

	// Drop/delete a field and its data
	DropField(ctx context.Context, param shared.CommonSystemParams) error

	//DuplicateModel(project *protobuff.Project, modelName, newName string) (*protobuff.ProjectSchema, error)

	// Rename a Model and its type
	//RenameModel(project *protobuff.Project, modelName, newName string) error

	// Convert between single type to double type model or vice versa
	//ConvertModel(project *protobuff.Project, modelName string) error

	// DuplicateField rename a field in a model along with its data key
	//DuplicateField(oldFieldName string, repeatedFieldGroup *string, param models.CommonSystemParams) error

	// RenameField rename a field in a model along with its data key
	RenameField(ctx context.Context, oldFieldName string, repeatedFieldGroup *string, param shared.CommonSystemParams) error
}

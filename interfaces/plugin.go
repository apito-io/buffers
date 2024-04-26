package interfaces

import (
	"context"
	"github.com/apito-cms/buffers/extensions"
)

// NormalPluginInterface interface functions
type NormalPluginInterface interface {
	// Init This functions runs when any extension runs
	Init(env []*extensions.EnvVariables) error

	// Migration If your extension has any migration script you can put it here
	Migration() error

	// SchemaRegister Define the GraphQL Schema That Will be Added If this extension registers
	SchemaRegister() (*extensions.ThirdPartyGraphQLSchemas, error)

	// RESTApiRegister Define the REST api that will be added to the existing list
	RESTApiRegister() ([]*extensions.ThirdPartyRESTApi, error)

	// Execute calls when the function is called
	Execute(request interface{}) (interface{}, error)
}

// StoragePluginInterface interface functions
type StoragePluginInterface interface {
	// Init This functions runs when any extension runs
	Init(env []*extensions.EnvVariables) error

	// Migration If your extension has any migration script you can put it here
	Migration(ctx context.Context) error

	// SchemaRegister Define the GraphQL Schema That Will be Added If this extension registers
	SchemaRegister(ctx context.Context) (*extensions.ThirdPartyGraphQLSchemas, error)

	// RESTApiRegister Define the REST api that will be added to the existing list
	RESTApiRegister(ctx context.Context) ([]*extensions.ThirdPartyRESTApi, error)

	// UploadFile calls when the function is called
	UploadFile(ctx context.Context, file *extensions.FileDetails) (interface{}, error)

	// ListFiles list all the files that this plugin is serving
	ListFiles(ctx context.Context, filter map[string]interface{}) ([]*extensions.FileDetails, error)

	// CountFiles count files
	CountFiles(ctx context.Context, filter map[string]interface{}) (int64, error)

	// DeleteFile delete a single file
	DeleteFile(ctx context.Context, id string) error
}

// FunctionPluginInterface interface functions
type FunctionPluginInterface interface {
	// Init This functions runs when any extension runs
	Init(env []*extensions.EnvVariables) error

	// Migration If your extension has any migration script you can put it here
	Migration() error

	// SchemaRegister Define the GraphQL Schema That Will be Added If this extension registers
	SchemaRegister() (*extensions.ThirdPartyGraphQLSchemas, error)

	// RESTApiRegister Define the REST api that will be added to the existing list
	RESTApiRegister() ([]*extensions.ThirdPartyRESTApi, error)

	// Execute calls when the function is called
	Execute(request interface{}) (interface{}, error)

	// ListFunctions list all the files that this plugin is serving
	ListFunctions(filter interface{}) ([]*extensions.CloudFunction, error)

	// RemoveFunction delete a single file
	RemoveFunction(id string) error
}

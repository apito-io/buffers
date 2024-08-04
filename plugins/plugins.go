package plugins

import (
	echo "github.com/labstack/echo/v4"
	"github.com/tailor-inc/graphql"
)

type ThirdPartyGraphQLSchemas struct {
	Queries   graphql.Fields
	Mutations graphql.Fields
}

type ThirdPartyRESTApi struct {
	Path       string
	Method     string
	Controller func(c echo.Context) error
}

type EnvVariables struct {
	Key   string `json:"key,omitempty" firestore:"key,omitempty"`
	Value string `json:"value,omitempty" firestore:"value,omitempty"`
}

type FileDetails struct {
	Id               string        `json:"id,omitempty" firestore:"id,omitempty"`
	XKey             string        `json:"_key,omitempty" firestore:"_key,omitempty"`
	Type             string        `json:"type,omitempty" firestore:"type,omitempty"`
	FileExtension    string        `json:"file_extension,omitempty" firestore:"file_extension,omitempty"`
	FileName         string        `json:"file_name,omitempty" firestore:"file_name,omitempty"`
	RemoteFilePath   string        `json:"remote_file_path,omitempty" firestore:"remote_file_path,omitempty"`
	UploadedFullPath string        `json:"uploaded_full_path,omitempty" firestore:"uploaded_full_path,omitempty"`
	ContentType      string        `json:"content_type,omitempty" firestore:"content_type,omitempty"`
	Size             int64         `json:"size,omitempty" firestore:"size,omitempty"`
	S3Key            string        `json:"s3_key,omitempty" firestore:"s3_key,omitempty"`
	Url              string        `json:"url,omitempty" firestore:"url,omitempty"`
	CreatedAt        string        `json:"created_at,omitempty" firestore:"created_at,omitempty"`
	UploadParam      *UploadParams `json:"upload_param,omitempty" firestore:"upload_param,omitempty"`
	Buffer           []byte        `json:"buffer,omitempty" firestore:"upload_param,omitempty"`
}

type UploadParams struct {
	DocId      string `json:"doc_id,omitempty" firestore:"doc_id,omitempty"`
	ProjectId  string `json:"project_id,omitempty" firestore:"project_id,omitempty"`
	ModelName  string `json:"model_name,omitempty" firestore:"model_name,omitempty"`
	FieldName  string `json:"field_name,omitempty" firestore:"field_name,omitempty"`
	AllowMulti bool   `json:"allow_multi,omitempty" firestore:"allow_multi,omitempty"`
	Provider   string `json:"provider,omitempty" firestore:"provider,omitempty"`
}

type CloudFunction struct {
	Id                   string `json:"id,omitempty" firestore:"id,omitempty"`
	Name                 string `json:"name,omitempty" firestore:"name,omitempty"`
	Description          string `json:"description,omitempty" firestore:"description,omitempty"`
	FunctionPath         string `json:"function_path,omitempty" firestore:"function_path,omitempty"`
	FunctionProviderName string `json:"function_provider_name,omitempty" firestore:"function_provider_name,omitempty"`
	FunctionConnected    bool   `json:"function_connected,omitempty" firestore:"function_connected,omitempty"`
	UpdatedAt            string `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`
	CreatedAt            string `json:"created_at,omitempty" firestore:"created_at,omitempty"`
	ExportedVariable     string `json:"exported_variable,omitempty" firestore:"exported_func_to_call,omitempty"`
}

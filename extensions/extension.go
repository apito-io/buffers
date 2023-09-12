package extensions

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
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" firestore:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" firestore:"value,omitempty"`
}

type FileDetails struct {
	Id            string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`
	XKey          string        `protobuf:"bytes,2,opt,name=_key,json=Key,proto3" json:"_key,omitempty" firestore:"_key,omitempty"`
	Type          string        `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" firestore:"type,omitempty"`
	FileExtension string        `protobuf:"bytes,4,opt,name=file_extension,json=fileExtension,proto3" json:"file_extension,omitempty" file_extension:"title,omitempty"`
	FileName      string        `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty" firestore:"file_name,omitempty"`
	ContentType   string        `protobuf:"bytes,6,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty" firestore:"content_type,omitempty"`
	Size_         int64         `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty" firestore:"size,omitempty"`
	S3Key         string        `protobuf:"bytes,8,opt,name=s3_key,json=s3Key,proto3" json:"s3_key,omitempty" firestore:"s3_key,omitempty"`
	Url           string        `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty" firestore:"url,omitempty"`
	CreatedAt     string        `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" firestore:"created_at,omitempty"`
	UploadParam   *UploadParams `protobuf:"bytes,11,opt,name=upload_param,json=uploadParam,proto3" json:"upload_param,omitempty" firestore:"upload_param,omitempty"`
	Buffer        []byte        `protobuf:"bytes,12,opt,name=buffer,proto3" json:"buffer,omitempty" firestore:"upload_param,omitempty"`
}

type UploadParams struct {
	DocId      string `protobuf:"bytes,1,opt,name=doc_id,json=docId,proto3" json:"doc_id,omitempty" firestore:"doc_id,omitempty"`
	ProjectId  string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" firestore:"project_id,omitempty"`
	ModelName  string `protobuf:"bytes,3,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty" firestore:"model_name,omitempty"`
	FieldName  string `protobuf:"bytes,4,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty" firestore:"field_name,omitempty"`
	AllowMulti bool   `protobuf:"varint,5,opt,name=allow_multi,json=allowMulti,proto3" json:"allow_multi,omitempty" firestore:"allow_multi,omitempty"`
	Provider   string `protobuf:"bytes,6,opt,name=provider,proto3" json:"provider,omitempty" firestore:"provider,omitempty"`
}

type CloudFunction struct {
	Id                   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`
	Name                 string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" firestore:"name,omitempty"`
	Description          string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" firestore:"description,omitempty"`
	FunctionPath         string `protobuf:"bytes,5,opt,name=function_path,json=functionPath,proto3" json:"function_path,omitempty" firestore:"function_path,omitempty"`
	FunctionProviderName string `protobuf:"bytes,8,opt,name=function_provider_name,json=functionProviderName,proto3" json:"function_provider_name,omitempty" firestore:"function_provider_name,omitempty"`
	FunctionConnected    bool   `protobuf:"varint,10,opt,name=function_connected,json=functionConnected,proto3" json:"function_connected,omitempty" firestore:"function_connected,omitempty"`
	UpdatedAt            string `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" firestore:"updated_at,omitempty"`
	CreatedAt            string `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" firestore:"created_at,omitempty"`
	ExportedVariable     string `protobuf:"bytes,13,opt,name=exported_variable,json=exportedVariable,proto3" json:"exported_variable,omitempty" firestore:"exported_func_to_call,omitempty"`
}

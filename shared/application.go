package shared

import (
	"github.com/apito-cms/buffers/protobuff"
	dlv6 "github.com/graph-gophers/dataloader"
	"github.com/tailor-inc/graphql"
)

type ConnectDisconnectParam struct {
	DocCollectionName string
	DocRelationName   string

	ConnectionIds       []string
	ConnectionType      string
	ForwardConnectionId string

	ForwardConnectionType      *protobuff.ConnectionType
	ForwardConnectionModelType *protobuff.ModelType

	BackwardConnectionType     *protobuff.ConnectionType
	BackendConnectionModelType *protobuff.ModelType

	KnownAs string
}

type DefaultDocumentStructure struct {
	Key      string                 `json:"_key,omitempty" firestore:"_key,omitempty"`
	Id       string                 `json:"id,omitempty" firestore:"id,omitempty"`
	Type     string                 `json:"type,omitempty" firestore:"type,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty" firestore:"data,omitempty"`
	Meta     *protobuff.MetaField   `json:"meta,omitempty" firestore:"meta,omitempty"`
	ExpireAt string                 `json:"expire_at,omitempty" firestore:"expire_at,omitempty"`
}

type CommonSystemParams struct {
	Role          *protobuff.Role `json:"role,omitempty"`
	Plan          string          `json:"plan,omitempty"`
	UserId        string          `json:"user_id,omitempty"`
	RelationModel string          `json:"relation_model,omitempty"`
	Email         string          `json:"email,omitempty"`
	ProjectId     string          `json:"project_id,omitempty"`
	TenantId      string          `json:"tenant_id,omitempty"`

	ResolveParams *graphql.ResolveParams `json:"resolve_params,omitempty"`

	DocumentId  string                    `json:"document_id,omitempty"`
	Document    *DefaultDocumentStructure `json:"document,omitempty"`
	Model       *protobuff.ModelType      `json:"model_type,omitempty"`
	ConDisParam []*ConnectDisconnectParam `json:"con_dis_param,omitempty"`
	FieldInfo   *protobuff.FieldInfo      `json:"field_info,omitempty"`

	KnownAs        string `json:"known_as,omitempty"`
	Revision       bool   `json:"revision,omitempty"`
	SinglePageData bool   `json:"single_page_data,omitempty"`

	//Limit *protobuff.UsagesTracking `json:"limit,,omitempty"`

	DocPublishStatus string `json:"doc_publish_status,omitempty"`

	IsSystemRequest bool `json:"is_system_request,omitempty"`

	UnmarshalStructure interface{} `json:"unmarshal_structure"`
}

type RawSchema struct {
	Queries   graphql.Fields
	Mutations graphql.Fields
}

type ApplicationCache struct {
	Project       *protobuff.Project      `json:"project,omitempty"`
	Param         *CommonSystemParams     `json:"param,omitempty"`
	RawSchemas    *RawSchema              `json:"raw_schema,omitempty"`
	DataloadersV6 map[string]*dlv6.Loader `json:"dataloaders,omitempty"`

	IncomingRequest []*IncomingRequest `json:"incoming_request"`
}

type IncomingRequest struct {
	OperationType  string
	FilteredModels []*FilteredModel
	IsFunction     bool
}

type FilteredModel struct {
	Name              string
	WhereFilter       []string
	IsConnectionQuery bool
}

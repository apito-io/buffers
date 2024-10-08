package shared

import (
	"encoding/json"

	"github.com/apito-io/buffers/protobuff"
	dlv6 "github.com/graph-gophers/dataloader"
	"github.com/graph-gophers/dataloader/v7"
	"github.com/tailor-inc/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

type DataloaderResult[T any] struct {
	Result  []T  `json:"result"`
	Count   int  `json:"count"`
	Cached  bool `json:"cached"`
	HasMore bool `json:"hasMore"`
	Error   bool `json:"error"`
	Code    int  `json:"code"`
}

type SearchResponse[T any] struct {
	Results        []*T
	GroupedResults map[string][]*T
}

// DataLoaders Dataloaders
type DataLoaders struct {
	MultiLoader *dataloader.Loader[string, interface{}]
	//SingleLoader *dataloader.Loader[string, interface{}]
}

type Response struct {
	Data       interface{}            `json:"data,omitempty"`
	Errors     json.RawMessage        `json:"errors,omitempty"`
	Extensions map[string]interface{} `json:"extensions,omitempty"`
}

type RawSchema struct {
	Queries   graphql.Fields
	Mutations graphql.Fields
}

type ApplicationCache struct {
	Project     *protobuff.Project      `json:"project,omitempty"`
	Param       *CommonSystemParams     `json:"param,omitempty"`
	RawSchemas  *RawSchema              `json:"raw_schema,omitempty"`
	Dataloaders map[string]*dlv6.Loader `json:"dataloaders,omitempty"`

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
	HasMetaQuery      bool
}

type CommonSystemParams struct {
	UserId        string `json:"user_id,omitempty"`
	RelationModel string `json:"relation_model,omitempty"`
	Email         string `json:"email,omitempty"`
	ProjectId     string `json:"project_id,omitempty"`
	TenantId      string `json:"tenant_id,omitempty"`

	ResolveParams *graphql.ResolveParams `json:"resolve_params,omitempty"`

	SystemCollectionName string `json:"system_collection_name,omitempty"`

	DocumentId  string   `json:"document_id,omitempty"`
	DocumentIDs []string `json:"document_ids,omitempty"`

	Document    *DefaultDocumentStructure `json:"document,omitempty"`
	Model       *protobuff.ModelType      `json:"model_type,omitempty"`
	ConDisParam []*ConnectDisconnectParam `json:"con_dis_param,omitempty"`
	FieldInfo   *protobuff.FieldInfo      `json:"field_info,omitempty"`

	KnownAs        string `json:"known_as,omitempty"`
	Revision       bool   `json:"revision,omitempty"`
	SinglePageData bool   `json:"single_page_data,omitempty"`

	IsSystemRequest bool `json:"is_system_request,omitempty"`

	IsEntireCollectionSearchRequest bool `json:"is_entire_collection_search_request,omitempty"`

	IsDataloaderRequest bool `json:"is_dataloader_request,omitempty"`

	IsIntersectionResult bool `json:"is_intersection_result,omitempty"`

	OnlyReturnCount bool `json:"only_return_count,omitempty"`

	// these three used in intersection of two collections
	SkipSort       bool `json:"skip_sort,omitempty"`
	SkipPagination bool `json:"skip_pagination,omitempty"`
	ReturnOnlyID   bool `json:"return_only_id,omitempty"`

	QuerySelectionSets *ast.SelectionSet `json:"query_selection_sets,omitempty"`

	UnmarshalStructure interface{} `json:"unmarshal_structure"`
}

type DocumentRevisionHistory struct {
	Id         string `json:"id"`
	RevisionAt string `json:"revision_at"`
	Status     string `json:"status"`
}

type DefaultDocumentStructure struct {
	Key      string                 `json:"_key,omitempty" firestore:"_key,omitempty"`
	Id       string                 `json:"id,omitempty" firestore:"id,omitempty"`
	Type     string                 `json:"type,omitempty" firestore:"type,omitempty"`
	Data     map[string]interface{} `json:"data,omitempty" firestore:"data,omitempty"`
	Meta     *protobuff.MetaField   `json:"meta,omitempty" firestore:"meta,omitempty"`
	ExpireAt string                 `json:"expire_at,omitempty" firestore:"expire_at,omitempty"`
}

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

type EdgeRelation struct {
	XFrom       string   `json:"_from,omitempty"`
	XTo         string   `json:"_to,omitempty"`
	Key         string   `json:"_key,omitempty"`
	Relation    string   `json:"relation,omitempty"`
	From        string   `json:"from,omitempty"`
	To          string   `json:"to,omitempty"`
	CreatedAt   string   `json:"created_at,omitempty"`
	KnownAs     string   `json:"known_as,omitempty"`
	Permissions []string `json:"permissions,omitempty"`
}

type ModelDocsResponse struct {
	Docs  []*DefaultDocumentStructure `json:"docs"`
	Count int                         `json:"count"`
}

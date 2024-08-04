package shared

import (
	"reflect"

	"github.com/apito-io/buffers/protobuff"
	"github.com/vektah/gqlparser/v2/ast"
)

type Filter struct {
	KeyWrapperFunction     string      `json:"keyWrapperFunction"`     // LOWER(x.name)
	Variable               string      `json:"variable"`               // x
	Key                    string      `json:"key"`                    // name
	Condition              string      `json:"condition"`              // ==
	Value                  interface{} `json:"value"`                  // fahim
	ComplexPredefinedQuery string      `json:"complexPredefinedQuery"` // for array filter -> COUNT(array[* FILTER CONTAINS(name, CURRENT)])
}

type DBPaginationFilter struct {
	Total int `json:"total"`
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

type QueryBuilder struct {
	RawFilterData map[string]interface{}

	UserID                   string   `json:"user_id"`
	TenantID                 string   `json:"tenant_id"`
	ProjectID                string   `json:"project_id"`
	RootCollectionFilterType string   `json:"doc_filter_type"`
	DocumentID               string   `json:"document_id"`
	DocumentIDs              []string `json:"document_ids"`

	ParentVariableName string `json:"parent_variable_name"`
	VariablePrefix     string `json:"variable_prefix"`
	VariableName       string `json:"variable_name"`
	CollectionName     string `json:"main_collection_name"`

	DefaultFilterCondition string               `json:"filter_condition"`
	WhereFilter            []*FilterInformation `json:"where_filter"`
	SortFilter             []*Filter            `json:"sort_and_limit_param"`
	PaginationFilter       *DBPaginationFilter  `json:"limit_filter"`
	FilterByLocal          string               `json:"local"`
	FilterByStatus         string               `json:"status"`

	ConnectionFilter map[string]interface{} `json:"connection_filter"`

	ApitoFields []*protobuff.FieldInfo `json:"apito_fields"`

	//QueryFilters  []*FilterInformation     `json:"query_filters"`
	SubQueries    []*QueryBuilder          `json:"sub_queries"`
	NestedQueries []*QueryBuilder          `json:"nested_queries"`
	ReturnFields  map[string]*FieldDetails `json:"return_fields"`

	ReturnFieldsSelection *ast.SelectionSet `json:"return_fields_selection"`

	IncludeDefaultSortAndLimit bool `json:"include_default_sort_and_limit"`
	IntersectResult            bool `json:"intersect_result"`
	FetchRevisionDocumentsOnly bool `json:"fetch_revision_documents_only"`
	IsDataloaderQuery          bool `json:"is_dataloader_query"`
	IsSystemRequest            bool `json:"is_system_query"`
	IsSystemCollectionQuery    bool `json:"is_system_collection_query"`
	IsEntireCollectionQuery    bool `json:"is_entire_collection_query"`
	SkipSort                   bool `json:"skip_sort"`
	SkipPagination             bool `json:"skip_limit"`
	ReturnOnlyID               bool `json:"return_only_id"`
}

type FilterInformation struct {
	Condition string    `json:"condition"`
	Filters   []*Filter `json:"filters"`
}

type FieldDetails struct {
	Kind      reflect.Kind
	SubFields []*protobuff.FieldInfo
	Local     string

	FieldType  string
	Validation *protobuff.Validation
	Value      interface{}
}

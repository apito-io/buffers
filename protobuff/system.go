package protobuff

import (
	"encoding/json"
	"github.com/uptrace/bun"
)

type PluginType int32

const (
	PluginType_NormalPlugin PluginType = 0 // auth, customFunction etc
	PluginType_Function     PluginType = 1 // go-plugin, aws-lambda etc
	PluginType_Storage      PluginType = 2 // s3, cloudinary etc
)

type PluginLoadStatus int32

const (
	PluginLoadStatus_NotInstalled PluginLoadStatus = 0
	PluginLoadStatus_Installed    PluginLoadStatus = 1
	PluginLoadStatus_ReInstall    PluginLoadStatus = 2
	PluginLoadStatus_Loading      PluginLoadStatus = 3
	PluginLoadStatus_Loaded       PluginLoadStatus = 4
	PluginLoadStatus_LoadFailed   PluginLoadStatus = 5
)

type PluginActivateStatus int32

const (
	PluginActivateStatus_deactivated PluginActivateStatus = 0
	PluginActivateStatus_activated   PluginActivateStatus = 1
)

type APIPermission struct {
	Read   string `json:"read,omitempty" firestore:"read,omitempty"`     // @gotags: firestore:"read,omitempty"
	Create string `json:"create,omitempty" firestore:"create,omitempty"` // @gotags: firestore:"create,omitempty"
	Update string `json:"update,omitempty" firestore:"update,omitempty"` // @gotags: firestore:"update,omitempty"
	Delete string `json:"delete,omitempty" firestore:"delete,omitempty"` // @gotags: firestore:"delete,omitempty"
}

type Role struct {
	bun.BaseModel             `bun:"table:role,alias:r"`
	Id                        string                    `bun:"id,pk,notnull,type:uuid,default:gen_random_uuid()" json:"id,omitempty" firestore:"id,omitempty"` // @gotags: firestore:"id,omitempty"
	ApiPermissions            map[string]*APIPermission `bun:",type:jsonb" json:"api_permissions,omitempty" firestore:"permissions,omitempty"`                 // @gotags: firestore:"permissions,omitempty"
	AdministrativePermissions []string                  `json:"administrative_permissions,omitempty" firestore:"administrative_permissions,omitempty"`         // @gotags: firestore:"administrative_permissions,omitempty"
	LogicExecutions           []string                  `json:"logic_executions,omitempty" firestore:"logic_executions,omitempty"`                             // @gotags: firestore:"logic_executions,omitempty"
	SystemGenerated           bool                      `json:"system_generated,omitempty" firestore:"system_generated,omitempty"`                             // @gotags: firestore:"system_generated,omitempty"
	IsAdmin                   bool                      `json:"is_admin,omitempty" firestore:"is_admin,omitempty"`                                             // @gotags: firestore:"is_admin,omitempty"
	IsProjectUser             bool                      `json:"is_project_user,omitempty" firestore:"is_project_user,omitempty"`                               // @gotags: firestore:"is_project_user,omitempty"
	ReadOnlyProject           bool                      `json:"read_only_project,omitempty" firestore:"read_only_project,omitempty"`                           // @gotags: firestore:"read_only_project,omitempty"
	IsSuperAdmin              bool                      `json:"is_super_admin,omitempty" firestore:"read_only_project,omitempty"`                              // @gotags: firestore:"read_only_project,omitempty"
}

// MarshalApiPermissions serializes ApiPermissions to JSON.
func (u *Role) MarshalApiPermissions() ([]byte, error) {
	return json.Marshal(u.ApiPermissions)
}

// UnmarshalApiPermissions deserializes JSON to ApiPermissions.
func (u *Role) UnmarshalApiPermissions(data []byte) error {
	return json.Unmarshal(data, &u.ApiPermissions)
}

type UserSubscriptionType int32

const (
	UserSubscriptionType_Undefined  UserSubscriptionType = 0
	UserSubscriptionType_Trial      UserSubscriptionType = 1
	UserSubscriptionType_Free       UserSubscriptionType = 2
	UserSubscriptionType_Basic      UserSubscriptionType = 3
	UserSubscriptionType_Pro        UserSubscriptionType = 4
	UserSubscriptionType_Enterprise UserSubscriptionType = 5
)

type SystemUser struct {
	XKey string `json:"_key,omitempty" firestore:"_key,omitempty"`                // @gotags: firestore:"_key,omitempty"
	Id   string `bun:"type:uuid,pk" json:"id,omitempty" firestore:"id,omitempty"` // @gotags: firestore:"id,omitempty"

	Name             string `json:"name,omitempty" firestore:"name,omitempty"`                             // @gotags: firestore:"name,omitempty"
	Role             string `json:"role,omitempty" firestore:"role,omitempty"`                             // @gotags: firestore:"role,omitempty"
	Username         string `json:"username,omitempty" firestore:"username,omitempty"`                     // @gotags: firestore:"username,omitempty"
	Email            string `json:"email,omitempty" firestore:"email,omitempty"`                           // @gotags: firestore:"email,omitempty"
	Secret           string `json:"secret,omitempty" firestore:"secret,omitempty"`                         // @gotags: firestore:"secret,omitempty"
	Avatar           string `json:"avatar,omitempty" firestore:"avatar,omitempty"`                         // @gotags: firestore:"avatar,omitempty"
	CurrentProjectId string `json:"current_project_id,omitempty" firestore:"current_project_id,omitempty"` // @gotags: firestore:"current_project_id,omitempty"

	AdministrativePermissions []string `json:"administrative_permissions,omitempty" firestore:"email,omitempty"` // @gotags: firestore:"email,omitempty"

	OrganizationID string `json:"organization_id,omitempty" firestore:"organization_id,omitempty"` // @gotags: firestore:"organization_id,omitempty"
	IsSuperAdmin   bool   `json:"is_super_admin,omitempty" firestore:"is_super_admin,omitempty"`   // @gotags: firestore:"is_super_admin,omitempty"

	RefreshToken    string `json:"refresh_token,omitempty" firestore:"refresh_token,omitempty"`         // @gotags: firestore:"refresh_token,omitempty"
	AccessToken     string `json:"access_token,omitempty" firestore:"access_token,omitempty"`           // @gotags: firestore:"access_token,omitempty"
	ReadOnlyProject bool   `json:"read_only_project,omitempty" firestore:"read_only_project,omitempty"` // @gotags: firestore:"read_only_project,omitempty"
	LastLoggedIn    string `json:"last_logged_in,omitempty" firestore:"last_logged_in,omitempty"`       // @gotags: firestore:"last_logged_in,omitempty"

	ProjectLimit         uint32               `json:"project_limit,omitempty" firestore:"project_limit,omitempty"`                   // @gotags: firestore:"project_limit,omitempty"
	UserSubscriptionType UserSubscriptionType `json:"user_subscription_type,omitempty" firestore:"user_subscription_type,omitempty"` // @gotags: firestore:"user_subscription_type,omitempty"

	Projects      []*Project       `bun:"rel:has-many,join:id=owner_id" json:"projects,omitempty" firestore:"projects,omitempty"` // @gotags: firestore:"projects,omitempty"
	Organizations []*Organizations `bun:"rel:has-many" json:"organizations,omitempty" firestore:"organizations,omitempty"`        // @gotags: firestore:"organizations,omitempty"

	CreatedAt string `bun:"type:timestamp,notnull,default:current_timestamp" json:"created_at,omitempty" firestore:"created_at,omitempty"` // @gotags: firestore:"created_at,omitempty"
	UpdatedAt string `bun:"type:timestamp,notnull" json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                           // @gotags: firestore:"updated_at,omitempty"
}

type Organizations struct {
	XKey        string `json:"_key,omitempty" firestore:"_key,omitempty"`               // @gotags: firestore:"_key,omitempty"
	ProjectID   string `json:"project_id,omitempty" firestore:"project_id,omitempty"`   // @gotags: firestore:"project_id,omitempty"
	OwnerID     string `json:"owner_id,omitempty" firestore:"owner_id,omitempty"`       // @gotags: firestore:"owner_id,omitempty"
	Id          string `json:"id,omitempty" firestore:"id,omitempty"`                   // @gotags: firestore:"id,omitempty"
	Name        string `json:"name,omitempty" firestore:"name,omitempty"`               // @gotags: firestore:"name,omitempty"
	Description string `json:"description,omitempty" firestore:"description,omitempty"` // @gotags: firestore:"description,omitempty"
}

type AddOnsDetails struct {
	ProjectID             string   `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"`  // @gotags: firestore:"project_id,omitempty"
	Locals                []string `json:"locals,omitempty" firestore:"locals,omitempty"`                             // @gotags: firestore:"locals,omitempty"
	SystemGraphqlHooks    bool     `json:"system_graphql_hooks,omitempty" firestore:"system_graphql_hooks,omitempty"` // @gotags: firestore:"system_graphql_hooks,omitempty"
	EnableRevisionHistory bool     `json:"enable_revision_history,omitempty" firestore:"revision_history,omitempty"`  // @gotags: firestore:"revision_history,omitempty"
}

type DriverCredentials struct {
	// for sql migration purposes
	ProjectID string `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	// normal sql & nosql database
	Engine   string `json:"engine,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     string `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	Database string `json:"database,omitempty"`
	// for firebase
	FirebaseProjectId             string `json:"firebase_project_id,omitempty"`
	FirebaseProjectCredentialJson string `json:"firebase_project_credential_json,omitempty"`
	// for dynamodb
	AccessKey string `json:"access_key,omitempty"`
	SecretKey string `json:"secret_key,omitempty"`
	// for embedded database
	DatabaseDir string `json:"database_dir,omitempty"`
}

type UsagesTracking struct {
	// for sql migration purposes
	ProjectID       string  `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	ApiCalls        uint32  `json:"api_calls,omitempty" firestore:"api_calls,omitempty"`                      // @gotags: firestore:"api_calls,omitempty"
	ApiBandwidth    float64 `json:"api_bandwidth,omitempty" firestore:"api_bandwidth,omitempty"`              // @gotags: firestore:"api_bandwidth,omitempty"
	MediaStorage    float64 `json:"media_storage,omitempty" firestore:"media_storage,omitempty"`              // @gotags: firestore:"media_storage,omitempty"
	MediaBandwidth  float64 `json:"media_bandwidth,omitempty" firestore:"media_bandwidth,omitempty"`          // @gotags: firestore:"media_bandwidth,omitempty"
	NumberOfMedia   float64 `json:"number_of_media,omitempty" firestore:"number_of_media,omitempty"`          // @gotags: firestore:"number_of_media,omitempty"
	NumberOfRecords float64 `json:"number_of_records,omitempty" firestore:"number_of_records,omitempty"`      // @gotags: firestore:"number_of_records,omitempty"
}

type PluginDetails struct {
	ProjectID        string                  `json:"project_id,omitempty" firestore:"project_id,omitempty"`               // @gotags: firestore:"project_id,omitempty"
	Icon             string                  `json:"icon,omitempty" firestore:"icon,omitempty"`                           // @gotags: firestore:"icon,omitempty"
	Id               string                  `json:"id,omitempty" firestore:"id,omitempty"`                               // @gotags: firestore:"id,omitempty"
	Title            string                  `json:"title,omitempty" firestore:"title,omitempty"`                         // @gotags: firestore:"title,omitempty"
	Serial           int64                   `json:"serial,omitempty" firestore:"serial,omitempty"`                       // @gotags: firestore:"serial,omitempty"
	Version          string                  `json:"version,omitempty" firestore:"version,omitempty"`                     // @gotags: firestore:"version,omitempty"
	Description      string                  `json:"description,omitempty" firestore:"description,omitempty"`             // @gotags: firestore:"description,omitempty"
	Type             PluginType              `json:"type,omitempty" firestore:"type,omitempty"`                           // @gotags: firestore:"type,omitempty"
	Role             string                  `json:"role,omitempty" firestore:"role,omitempty"`                           // @gotags: firestore:"role,omitempty"
	EnvVars          []*FunctionEnvVariables `json:"env_vars,omitempty" firestore:"env_vars,omitempty"`                   // @gotags: firestore:"env_vars,omitempty"
	ExportedVariable string                  `json:"exported_variable,omitempty" firestore:"exported_variable,omitempty"` // @gotags: firestore:"exported_variable,omitempty"
	Enable           bool                    `json:"enable,omitempty" firestore:"enable,omitempty"`                       // @gotags: firestore:"enable,omitempty"
	RepositoryUrl    string                  `json:"repository_url,omitempty" firestore:"repository_url,omitempty"`       // @gotags: firestore:"repository_url,omitempty"
	Branch           string                  `json:"branch,omitempty" firestore:"branch,omitempty"`                       // @gotags: firestore:"branch,omitempty"
	Author           string                  `json:"author,omitempty" firestore:"author,omitempty"`                       // @gotags: firestore:"author,omitempty"
	LoadStatus       PluginLoadStatus        `json:"load_status,omitempty" firestore:"load_status,omitempty"`             // @gotags: firestore:"load_status,omitempty"
	ActivateStatus   PluginActivateStatus    `json:"activate_status,omitempty" firestore:"activate_status,omitempty"`     // @gotags: firestore:"activate_status,omitempty"
}

type APIToken struct {
	ProjectID string `json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	Name      string `json:"name,omitempty" firestore:"name,omitempty"`             // @gotags: firestore:"name,omitempty"
	Token     string `json:"token,omitempty" firestore:"token,omitempty"`           // @gotags: firestore:"token,omitempty"
	Role      string `json:"role,omitempty" firestore:"role,omitempty"`             // @gotags: firestore:"role,omitempty"
	Expire    string `json:"expire,omitempty" firestore:"expire,omitempty"`         // @gotags: firestore:"expire,omitempty"
}

type SubscriptionCheckRequest struct {
	SubscriptionType UserSubscriptionType `json:"subscription_type"`
	Name             string               `json:"name"`
	SubscriptionCode string               `json:"subscription_code"`
	FeaturePromise   []string             `json:"feature_promise"`
}

type ProjectCreateRequest struct {
	Id          string             `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Token       string             `json:"token"`
	Engine      string             `json:"engine"`
	Driver      *DriverCredentials `json:"driver"`
	Example     string             `json:"example"`
}

type ActivatedPlugins struct {
	Storage  string `json:"storage,omitempty"`
	Function string `json:"function,omitempty"`
}

type Workspace struct {
	ProjectID    string `json:"project_id,omitempty" firestore:"project_id,omitempty"`       // @gotags: firestore:"project_id,omitempty"
	Name         string `json:"name,omitempty" firestore:"name,omitempty"`                   // @gotags: firestore:"name,omitempty"
	Active       bool   `json:"active,omitempty" firestore:"active,omitempty"`               // @gotags: firestore:"active,omitempty"
	IsProduction bool   `json:"is_production,omitempty" firestore:"is_production,omitempty"` // @gotags: firestore:"is_production,omitempty"
	IsDefault    bool   `json:"is_default,omitempty" firestore:"is_default,omitempty"`       // @gotags: firestore:"is_default,omitempty"
}

type SystemMessage struct {
	ProjectID   string `json:"project_id,omitempty" firestore:"project_id,omitempty"`   // @gotags: firestore:"project_id,omitempty"
	Message     string `json:"message,omitempty" firestore:"message,omitempty"`         // @gotags: firestore:"message,omitempty"
	Code        string `json:"code,omitempty" firestore:"code,omitempty"`               // @gotags: firestore:"code,omitempty"
	Redirection string `json:"redirection,omitempty" firestore:"redirection,omitempty"` // @gotags: firestore:"redirection,omitempty"
	Hide        bool   `json:"hide,omitempty" firestore:"hide,omitempty"`               // @gotags: firestore:"hide,omitempty"
}

// Project user project
type Project struct {
	XKey               string             `json:"_key,omitempty" firestore:"_key,omitempty"` // @gotags: firestore:"_key,omitempty"
	Id                 string             `bun:"type:uuid,pk" json:"id,omitempty" firestore:"id,omitempty"`
	OwnerId            string             `json:"owner_id,omitempty" firestore:"owner_id,omitempty"`
	OrganizationID     string             `json:"organization_id,omitempty" firestore:"organization_id,omitempty"`                                                                                      // @gotags: firestore:"organization_id,omitempty"
	ProjectName        string             `json:"project_name,omitempty" firestore:"project_name,omitempty"`                                                                                            // @gotags: firestore:"project_name,omitempty"
	ProjectDescription string             `json:"project_description,omitempty" firestore:"project_description,omitempty"`                                                                              // @gotags: firestore:"project_description,omitempty"
	Schema             *ProjectSchema     `bun:"rel:belongs-to,join:id=project_id" json:"schema,omitempty" firestore:"schema,omitempty"`                                                                // @gotags: firestore:"schema,omitempty"
	CreatedAt          string             `json:"created_at,omitempty" firestore:"created_at,omitempty"`                                                                                                // @gotags: firestore:"created_at,omitempty"
	UpdatedAt          string             `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                                                                                                // @gotags: firestore:"updated_at,omitempty"
	ExpireAt           string             `json:"expire_at,omitempty" firestore:"expire_at,omitempty"`                                                                                                  // @gotags: firestore:"expire_at,omitempty"
	Plugins            []*PluginDetails   `bun:"rel:has-many" json:"plugins,omitempty" firestore:"plugins,omitempty"`                                                                                   // @gotags: firestore:"plugins,omitempty"
	Addons             *AddOnsDetails     `bun:"rel:belongs-to,join:id=project_id" json:"addons,omitempty"  firestore:"addons,omitempty"`                                                               // @gotags: gorm:"foreignKey:ProjectID" firestore:"addons,omitempty"
	Tokens             []*APIToken        `bun:"rel:has-many" json:"tokens,omitempty" firestore:"tokens,omitempty"`                                                                                     // @gotags: gorm:"foreignKey:ProjectID" firestore:"tokens,omitempty"
	Roles              map[string]*Role   `bun:"type:jsonb" json:"roles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"roles,omitempty"` // @gotags: firestore:"roles,omitempty"
	Driver             *DriverCredentials `bun:"rel:belongs-to,join:id=project_id" json:"driver,omitempty"  firestore:"driver,omitempty"`                                                               // @gotags: gorm:"foreignKey:ProjectID" firestore:"driver,omitempty"
	TempBanned         bool               `json:"temp_banned,omitempty" firestore:"temp_banned,omitempty"`                                                                                              // @gotags: firestore:"temp_banned,omitempty"

	TrialEnds   string          `json:"trial_ends,omitempty" firestore:"trial_ends,omitempty"`                                  // @gotags: firestore:"trial_ends,omitempty"
	FromExample string          `json:"from_example,omitempty" firestore:"from_example,omitempty"`                              // @gotags: firestore:"from_example,omitempty"
	Limits      *UsagesTracking `bun:"rel:belongs-to,join:id=project_id" json:"limits,omitempty"  firestore:"limits,omitempty"` // @gotags: gorm:"foreignKey:ProjectID" firestore:"limits,omitempty"

	Teams          []*TeamMembers   `bun:"rel:has-many" json:"teams,omitempty"  firestore:"teams,omitempty"`                    // @gotags: gorm:"foreignKey:ProjectID" firestore:"teams,omitempty"
	SystemMessages []*SystemMessage `bun:"rel:has-many" json:"system_messages,omitempty" firestore:"system_messages,omitempty"` // @gotags: firestore:"system_messages,omitempty"
	Workspaces     []*Workspace     `bun:"rel:has-many" json:"workspaces,omitempty" firestore:"workspaces,omitempty"`           // @gotags: firestore:"workspaces,omitempty"
	//ActivatedPluginsIds *ActivatedPlugins `bun:"rel:belongs-to,join:id=project_id" json:"activated_plugins_ids,omitempty" firestore:"activated_plugins_ids,omitempty"` // @gotags: firestore:"activated_plugins_ids,omitempty"

	UserSubscriptionType UserSubscriptionType `json:"user_subscription_type,omitempty" firestore:"user_subscription_type,omitempty"` // @gotags: firestore:"user_subscription_type,omitempty"
	IsPaymentDue         bool                 `json:"is_payment_due,omitempty" firestore:"is_payment_due,omitempty"`                 // @gotags: firestore:"is_payment_due,omitempty"

	// for microservice
	MicroServicePort string `json:"micro_service_port,omitempty" firestore:"micro_service_port,omitempty"`
}

type TeamMembers struct {
	// for sql migration purposes
	ProjectID         string   `json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	UserId            string   `json:"user_id,omitempty"`
	AssignedRole      string   `json:"assigned_role,omitempty"`
	AccessPermissions []string `json:"access_permissions,omitempty"`
}

type ProjectSchema struct {
	ProjectID string           `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	Models    []*ModelType     `bun:"rel:has-many" json:"models,omitempty" firestore:"models,omitempty"`         // @gotags: firestore:"models,omitempty"
	Functions []*CloudFunction `bun:"rel:has-many" json:"functions,omitempty" firestore:"functions,omitempty"`   // @gotags: firestore:"functions,omitempty"
}

type ModelType struct {
	ProjectID       string            `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	Name            string            `json:"name,omitempty" firestore:"name,omitempty"`                                // @gotags: firestore:"name,omitempty"
	Description     string            `json:"description,omitempty"`
	Fields          []*FieldInfo      `json:"fields,omitempty" firestore:"fields,omitempty"`                     // @gotags: firestore:"fields,omitempty"
	Connections     []*ConnectionType `json:"connections,omitempty" firestore:"connections,omitempty"`           // @gotags: firestore:"connections,omitempty"
	HookIds         []string          `json:"hook_ids,omitempty" firestore:"hook_ids,omitempty"`                 // @gotags: firestore:"hook_ids,omitempty"
	Locals          []string          `json:"locals,omitempty" firestore:"locals,omitempty"`                     // @gotags: firestore:"locals,omitempty"
	RepeatedGroups  []string          `json:"repeated_groups,omitempty" firestore:"locals,omitempty"`            // @gotags: firestore:"locals,omitempty"
	SystemGenerated bool              `json:"system_generated,omitempty" firestore:"system_generated,omitempty"` // @gotags: firestore:"system_generated,omitempty"
	SinglePage      bool              `json:"single_page,omitempty" firestore:"system_generated,omitempty"`      // @gotags: firestore:"system_generated,omitempty"
	SinglePageUuid  string            `json:"single_page_uuid,omitempty" firestore:"system_generated,omitempty"` // @gotags: firestore:"system_generated,omitempty"
	HasConnections  bool              `json:"has_connections,omitempty" firestore:"has_connections,omitempty"`   // @gotags: firestore:"has_connections,omitempty"
}

type FunctionEnvVariables struct {
	Key   string `json:"key,omitempty" firestore:"key,omitempty"`     // @gotags: firestore:"key,omitempty"
	Value string `json:"value,omitempty" firestore:"value,omitempty"` // @gotags: firestore:"value,omitempty"
}

type CloudFunctionRequestResponseType struct {
	Model  string       `json:"model,omitempty" firestore:"model,omitempty"`   // @gotags: firestore:"model,omitempty"
	Params []*FieldInfo `json:"params,omitempty" firestore:"params,omitempty"` // @gotags: firestore:"params,omitempty"
}

type CloudFunction struct {
	ProjectID                string                            `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"`              // @gotags: firestore:"project_id,omitempty"
	Id                       string                            `json:"id,omitempty" firestore:"id,omitempty"`                                                 // @gotags: firestore:"id,omitempty"
	Name                     string                            `json:"name,omitempty" firestore:"name,omitempty"`                                             // @gotags: firestore:"name,omitempty"
	Description              string                            `json:"description,omitempty" firestore:"description,omitempty"`                               // @gotags: firestore:"description,omitempty"
	EnvVars                  []*FunctionEnvVariables           `json:"env_vars,omitempty" firestore:"env_vars,omitempty"`                                     // @gotags: firestore:"env_vars,omitempty"`
	FunctionPath             string                            `json:"function_path,omitempty" firestore:"function_path,omitempty"`                           // @gotags: firestore:"function_path,omitempty"`
	Request                  *CloudFunctionRequestResponseType `json:"request,omitempty" firestore:"request,omitempty"`                                       // @gotags: firestore:"request,omitempty"`
	Response                 *CloudFunctionRequestResponseType `json:"response,omitempty" firestore:"response,omitempty"`                                     // @gotags: firestore:"response,omitempty"`
	FunctionProviderId       string                            `json:"function_provider_id,omitempty" firestore:"function_provider_id,omitempty"`             // @gotags: firestore:"function_provider_id,omitempty"`
	FunctionConnected        bool                              `json:"function_connected,omitempty" firestore:"function_connected,omitempty"`                 // @gotags: firestore:"function_connected,omitempty"`
	ProviderExportedVariable string                            `json:"provider_exported_variable,omitempty" firestore:"provider_exported_variable,omitempty"` // @gotags: firestore:"provider_exported_variable,omitempty"`
	FunctionExportedVariable string                            `json:"function_exported_variable,omitempty" firestore:"function_exported_variable,omitempty"` // @gotags: firestore:"function_exported_variable,omitempty"`
	RuntimeConfig            *FunctionRuntimeConfig            `json:"runtime_config,omitempty" firestore:"runtime_config,omitempty"`                         // @gotags: firestore:"runtime_config,omitempty"`
	UpdatedAt                string                            `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                                 // @gotags: firestore:"updated_at,omitempty"`
	CreatedAt                string                            `json:"created_at,omitempty" firestore:"created_at,omitempty"`                                 // @got
}

type ProjectWithRoles struct {
	Project     *Project `json:"project"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

type ProjectInvoices struct {
	XKey                  string  `json:"_key,omitempty" firestore:"_key,omitempty"`                                                                                                // @gotags: firestore:"_key,omitempty"
	Id                    string  `json:"id,omitempty" firestore:"id,omitempty"`                                                                                                    // @gotags: firestore:"id,omitempty"
	Type                  string  `json:"type,omitempty" firestore:"type,omitempty"`                                                                                                // @gotags: firestore:"type,omitempty"
	SubscriptionId        string  `json:"subscription_id,omitempty" firestore:"subscription_id,omitempty"`                                                                          // @gotags: firestore:"subscription_id,omitempty"
	CouponCode            string  `json:"coupon_code,omitempty" firestore:"coupon_code,omitempty"`                                                                                  // @gotags: firestore:"coupon_code,omitempty"
	DiscountAmount        float64 `protobuf:"fixed64,6,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty" firestore:"discount_amount,omitempty"` // @gotags: firestore:"discount_amount,omitempty"
	PaymentGateway        string  `json:"payment_gateway,omitempty" firestore:"payment_gateway,omitempty"`                                                                          // @gotags: firestore:"payment_gateway,omitempty"
	ReceiptUrl            string  `json:"receipt_url,omitempty" firestore:"receipt_url,omitempty"`                                                                                  // @gotags: firestore:"receipt_url,omitempty"
	ProductAmount         float64 `json:"product_amount,omitempty" firestore:"product_amount,omitempty"`                                                                            // @gotags: firestore:"product_amount,omitempty"
	DiscountPercentage    int64   `json:"discount_percentage,omitempty" firestore:"discount_percentage,omitempty"`                                                                  // @gotags: firestore:"discount_percentage,omitempty"
	Quantity              uint32  `json:"quantity,omitempty" firestore:"quantity,omitempty"`                                                                                        // @gotags: firestore:"quantity,omitempty"
	PaymentMethod         string  `json:"payment_method,omitempty" firestore:"payment_method,omitempty"`                                                                            // @gotags: firestore:"payment_method,omitempty"
	Earning               float64 `json:"earning,omitempty" firestore:"earning,omitempty"`                                                                                          // @gotags: firestore:"earning,omitempty"
	Fee                   float64 `json:"fee,omitempty" firestore:"fee,omitempty"`                                                                                                  // @gotags: firestore:"fee,omitempty"
	Tax                   float64 `json:"tax,omitempty" firestore:"tax,omitempty"`                                                                                                  // @gotags: firestore:"tax,omitempty"
	CreatedAt             string  `json:"created_at,omitempty" firestore:"created_at,omitempty"`                                                                                    // @gotags: firestore:"created_at,omitempty"
	ChargedAmount         float64 `json:"charged_amount,omitempty" firestore:"charged_amount,omitempty"`                                                                            // @gotags: firestore:"charged_amount,omitempty"
	NextAmountToBeCharged float64 `json:"next_amount_to_be_charged,omitempty" firestore:"next_amount_to_be_charged,omitempty"`                                                      // @gotags: firestore:"next_amount_to_be_charged,omitempty"
	SubscriptionPlanCode  string  `json:"subscription_plan_code,omitempty" firestore:"subscription_plan_code,omitempty"`                                                            // @gotags: firestore:"subscription_plan_code,omitempty"
}

type AccountUsage struct {
	Id              string                 `json:"id,omitempty" firestore:"id,omitempty"`                               // @gotags: firestore:"id,omitempty"
	XKey            string                 `json:"_key,omitempty" firestore:"_key,omitempty"`                           // @gotags: firestore:"_key,omitempty"
	Type            string                 `json:"type,omitempty" firestore:"type,omitempty"`                           // @gotags: firestore:"type,omitempty"
	Subscriptions   []*MonthlySubscription `json:"subscriptions,omitempty" firestore:"subscriptions,omitempty"`         // @gotags: firestore:"subscriptions,omitempty"
	Limits          *UsagesTracking        `json:"limits,omitempty" firestore:"limits,omitempty"`                       // @gotags: firestore:"limits,omitempty"
	Usages          []*UsagesTracking      `json:"usages,omitempty" firestore:"usages,omitempty"`                       // @gotags: firestore:"usages,omitempty"
	NumberOfProject uint32                 `json:"number_of_project,omitempty" firestore:"number_of_project,omitempty"` // @gotags: firestore:"number_of_project,omitempty"
}

type MonthlySubscription struct {
	Id                 string             `json:"id,omitempty" firestore:"id,omitempty"`                                   // @gotags: firestore:"id,omitempty"
	XKey               string             `json:"_key,omitempty" firestore:"_key,omitempty"`                               // @gotags: firestore:"_key,omitempty"
	Type               string             `json:"type,omitempty" firestore:"type,omitempty"`                               // @gotags: firestore:"type,omitempty"
	UserId             string             `json:"user_id,omitempty" firestore:"user_id,omitempty"`                         // @gotags: firestore:"user_id,omitempty"
	SubscriptionCode   string             `json:"subscription_code,omitempty" firestore:"plan_code,omitempty"`             // @gotags: firestore:"plan_code,omitempty"
	StartDate          string             `json:"start_date,omitempty" firestore:"start_date,omitempty"`                   // @gotags: firestore:"start_date,omitempty"
	EndDate            string             `json:"end_date,omitempty" firestore:"end_date,omitempty"`                       // @gotags: firestore:"end_date,omitempty"
	SubscriptionStatus string             `json:"subscription_status,omitempty" firestore:"subscription_status,omitempty"` // @gotags: firestore:"subscription_status,omitempty"
	PaymentStatus      string             `json:"payment_status,omitempty" firestore:"payment_status,omitempty"`           // @gotags: firestore:"payment_status,omitempty"
	Invoices           []*ProjectInvoices `json:"invoices,omitempty" firestore:"invoices,omitempty"`                       // @gotags: firestore:"invoices,omitempty"
	CreatedAt          string             `json:"created_at,omitempty" firestore:"created_at,omitempty"`                   // @gotags: firestore:"created_at,omitempty"
	UpdatedAt          string             `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                   // @gotags: firestore:"updated_at,omitempty"
	Title              string             `json:"title,omitempty" firestore:"title,omitempty"`                             // @gotags: firestore:"title,omitempty"
	CancelUrl          string             `json:"cancel_url,omitempty" firestore:"cancel_url,omitempty"`                   // @gotags: firestore:"cancel_url,omitempty"
	ProjectId          string             `json:"project_id,omitempty" firestore:"cancel_url,omitempty"`                   // @gotags: firestore:"cancel_url,omitempty"
}

type UserSubscription struct {
	Subscription *MonthlySubscription `json:"subscription,omitempty" firestore:"subscription,omitempty"` // @gotags: firestore:"subscription,omitempty"
	Invoices     []*ProjectInvoices   `json:"invoices,omitempty" firestore:"invoices,omitempty"`         // @gotags: firestore:"invoices,omitempty"
}

type ProjectUsages struct {
	Id        string          `json:"id,omitempty" firestore:"id,omitempty"`                 // @gotags: firestore:"id,omitempty"
	XKey      string          `json:"_key,omitempty" firestore:"_key,omitempty"`             // @gotags: firestore:"_key,omitempty"
	Type      string          `json:"type,omitempty" firestore:"type,omitempty"`             // @gotags: firestore:"type,omitempty"
	ProjectId string          `json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	Year      uint32          `json:"year,omitempty" firestore:"year,omitempty"`             // @gotags: firestore:"year,omitempty"
	Month     uint32          `json:"month,omitempty" firestore:"month,omitempty"`           // @gotags: firestore:"month,omitempty"
	Usages    *UsagesTracking `json:"usages,omitempty" firestore:"usage,omitempty"`          // @gotags: firestore:"usage,omitempty"
}

type FunctionRuntimeConfig struct {
	Runtime string `json:"runtime,omitempty" firestore:"runtime,omitempty"`   // @gotags: firestore:"runtime,omitempty"
	Memory  int64  `json:"memory,omitempty" firestore:"memory,omitempty"`     // @gotags: firestore:"memory,omitempty"
	Handler string `json:"handler,omitempty" firestore:"handler,omitempty"`   // @gotags: firestore:"handler,omitempty"
	TimeOut int64  `json:"time_out,omitempty" firestore:"time_out,omitempty"` // @gotags: firestore:"time_out,omitempty"
}

type Validation struct {
	Required          bool      `json:"required,omitempty" firestore:"required,omitempty"`                       // @gotags: firestore:"required,omitempty"
	Hide              bool      `json:"hide,omitempty" firestore:"hide,omitempty"`                               // @gotags: firestore:"hide,omitempty"
	AsTitle           bool      `json:"as_title,omitempty" firestore:"as_title,omitempty"`                       // @gotags: firestore:"as_title,omitempty"
	Locals            []string  `json:"locals,omitempty" firestore:"locals,omitempty"`                           // @gotags: firestore:"locals,omitempty"
	Unique            bool      `json:"unique,omitempty" firestore:"unique,omitempty"`                           // @gotags: firestore:"unique,omitempty"
	CharLimit         []uint32  `json:"char_limit,omitempty" firestore:"char_limit,omitempty"`                   // @gotags: firestore:"char_limit,omitempty"
	IntRangeLimit     []uint32  `json:"int_range_limit,omitempty" firestore:"int_range_limit,omitempty"`         // @gotags: firestore:"int_range_limit,omitempty"
	DoubleRangeLimit  []float64 `json:"double_range_limit,omitempty" firestore:"double_range_limit,omitempty"`   // @gotags: firestore:"double_range_limit,omitempty"
	IsEmail           bool      `json:"is_email,omitempty" firestore:"is_email,omitempty"`                       // @gotags: firestore:"is_email,omitempty"
	Placeholder       string    `json:"placeholder,omitempty" firestore:"placeholder,omitempty"`                 // @gotags: firestore:"placeholder,omitempty"
	FixedListElements []string  `json:"fixed_list_elements,omitempty" firestore:"fixed_list_elements,omitempty"` // @gotags: firestore:"fixed_list_elements,omitempty"
	IsMultiChoice     bool      `json:"is_multi_choice,omitempty" firestore:"is_multi_choice,omitempty"`         // @gotags: firestore:"is_multi_choice,omitempty"
	ListType          string    `json:"list_type,omitempty" firestore:"list_type,omitempty"`                     // @gotags: firestore:"list_type,omitempty"
	IsGallery         bool      `json:"is_gallery,omitempty" firestore:"is_gallery,omitempty"`                   // @gotags: firestore:"is_gallery,omitempty"
	IsPassword        bool      `json:"is_password,omitempty" firestore:"is_gallery,omitempty"`                  // @gotags: firestore:"is_gallery,omitempty"
	IsSystemRole      bool      `json:"is_system_role,omitempty" firestore:"is_gallery,omitempty"`               // @gotags: firestore:"is_gallery,omitempty"
	IsUrl             bool      `json:"is_url,omitempty" firestore:"is_url,omitempty"`                           // @gotags: firestore:"is_url,omitempty"
}

type FieldInfo struct {
	Identifier              string       `json:"identifier,omitempty" firestore:"identifier,omitempty"`   // @gotags: firestore:"identifier,omitempty"
	Description             string       `json:"description,omitempty" firestore:"description,omitempty"` // @gotags: firestore:"description,omitempty"
	InputType               string       `json:"input_type,omitempty" firestore:"input_type,omitempty"`   // @gotags: firestore:"input_type,omitempty"
	FieldType               string       `json:"field_type,omitempty" firestore:"field_type,omitempty"`   // @gotags: firestore:"field_type,omitempty"
	FieldSubType            string       `json:"field_sub_type,omitempty"`
	SubFieldInfo            []*FieldInfo `json:"sub_field_info,omitempty" firestore:"modules,omitempty"`                              // @gotags: firestore:"modules,omitempty"
	Validation              *Validation  `json:"validation,omitempty" firestore:"validation,omitempty"`                               // @gotags: firestore:"validation,omitempty"
	Serial                  uint32       `json:"serial,omitempty" firestore:"serial,omitempty"`                                       // @gotags: firestore:"serial,omitempty"
	Label                   string       `json:"label,omitempty" firestore:"label,omitempty"`                                         // @gotags: firestore:"label,omitempty"
	SystemGenerated         bool         `json:"system_generated,omitempty" firestore:"system_generated,omitempty"`                   // @gotags: firestore:"system_generated,omitempty"
	RepeatedGroupIdentifier string       `json:"repeated_group_identifier,omitempty" firestore:"repeated_group_identifier,omitempty"` // @gotags: firestore:"repeated_group_identifier,omitempty"
	IsObjectField           bool         `json:"is_object_field,omitempty" firestore:"is_object_field,omitempty"`                     // @gotags: firestore:"is_object_field,omitempty"
	ParentField             string       `json:"parent_field,omitempty" firestore:"parent_field,omitempty"`                           // @gotags: firestore:"parent_field,omitempty"
}

type ConnectionType struct {
	Model    string `json:"model,omitempty" firestore:"model,omitempty"`       // @gotags: firestore:"model,omitempty"
	Relation string `json:"relation,omitempty" firestore:"relation,omitempty"` // @gotags: firestore:"relation,omitempty"
	Type     string `json:"type,omitempty" firestore:"type,omitempty"`         // @gotags: firestore:"type,omitempty"
	KnownAs  string `json:"known_as,omitempty" firestore:"known_as,omitempty"` // @gotags: firestore:"known_as,omitempty"
}

type Webhook struct {
	Id              string   `json:"id,omitempty" firestore:"id,omitempty"`                             // @gotags: firestore:"id,omitempty"
	XKey            string   `json:"_key,omitempty" firestore:"_key,omitempty"`                         // @gotags: firestore:"_key,omitempty"
	Type            string   `json:"type,omitempty" firestore:"type,omitempty"`                         // @gotags: firestore:"type,omitempty"
	Model           string   `json:"model,omitempty" firestore:"model,omitempty"`                       // @gotags: firestore:"model,omitempty"
	ProjectId       string   `json:"project_id,omitempty" firestore:"project_id,omitempty"`             //
	Name            string   `json:"name,omitempty" firestore:"name,omitempty"`                         // @gotags: firestore:"name,omitempty"
	Events          []string `json:"events,omitempty" firestore:"events,omitempty"`                     // @gotags: firestore:"events,omitempty"
	Url             string   `json:"url,omitempty" firestore:"url,omitempty"`                           // @gotags: firestore:"url,omitempty"
	LogicExecutions []string `json:"logic_executions,omitempty" firestore:"logic_executions,omitempty"` // @gotags: firestore:"logic_executions,omitempty"
}

type SupportAndTicket struct {
	XKey             string         `json:"_key,omitempty"`
	Id               string         `json:"id,omitempty"`
	Type             string         `json:"type,omitempty"`
	ProjectId        string         `json:"project_id,omitempty"`
	Resolved         bool           `json:"resolved,omitempty"`
	Title            string         `json:"title,omitempty"`
	IssueDescription string         `json:"issue_description,omitempty"`
	CreatedAt        string         `json:"created_at,omitempty"`
	UpdatedAt        string         `json:"updated_at,omitempty"`
	Replies          []*TicketReply `json:"replies,omitempty"`
}

type TicketReply struct {
	Description string      `json:"description,omitempty"`
	User        *SystemUser `json:"user,omitempty"`
	CreatedAt   string      `json:"created_at,omitempty"`
	Edited      bool        `json:"edited,omitempty"`
}

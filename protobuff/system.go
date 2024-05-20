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

type SystemUser struct {
	XKey             string `json:"_key,omitempty" firestore:"_key,omitempty"`                            // @gotags: firestore:"_key,omitempty"
	Id               string `bun:"type:uuid,pk" json:"id,omitempty" firestore:"id,omitempty"`             // @gotags: firestore:"id,omitempty"
	Name             string `json:"name,omitempty" firestore:"name,omitempty"`                            // @gotags: firestore:"name,omitempty"
	Username         string `json:"username,omitempty" firestore:"username,omitempty"`                    // @gotags: firestore:"username,omitempty"
	Email            string `json:"email,omitempty" firestore:"email,omitempty"`                          // @gotags: firestore:"email,omitempty"
	Secret           string `json:"secret,omitempty" firestore:"secret,omitempty"`                        // @gotags: firestore:"secret,omitempty"
	Avatar           string `json:"avatar,omitempty" firestore:"avatar,omitempty"`                        // @gotags: firestore:"avatar,omitempty"
	CurrentProjectId string `son:"current_project_id,omitempty" firestore:"current_project_id,omitempty"` // @gotags: firestore:"current_project_id,omitempty"
	CreatedAt        string `json:"created_at,omitempty" firestore:"created_at,omitempty"`                // @gotags: firestore:"created_at,omitempty"
	UpdatedAt        string `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                // @gotags: firestore:"updated_at,omitempty"
	IsSuperAdmin     bool   `json:"is_super_admin,omitempty" firestore:"is_super_admin,omitempty"`        // @gotags: firestore:"is_super_admin,omitempty"
	RefreshToken     string `json:"refresh_token,omitempty" firestore:"refresh_token,omitempty"`          // @gotags: firestore:"refresh_token,omitempty"
	AccessToken      string `json:"access_token,omitempty" firestore:"access_token,omitempty"`            // @gotags: firestore:"access_token,omitempty"
	ReadOnlyProject  bool   `json:"read_only_project,omitempty" firestore:"read_only_project,omitempty"`  // @gotags: firestore:"read_only_project,omitempty"
	LastLoggedIn     string `json:"last_logged_in,omitempty" firestore:"last_logged_in,omitempty"`        // @gotags: firestore:"last_logged_in,omitempty"
	ProjectLimit     uint32 `json:"project_limit,omitempty" firestore:"project_limit,omitempty"`          // @gotags: firestore:"project_limit,omitempty"

	Projects []*Project `bun:"rel:has-many" json:"projects,omitempty" firestore:"projects,omitempty"` // @gotags: firestore:"projects,omitempty"
}

type AddOnsDetails struct {
	ProjectID             string   `json:"project_id,omitempty" firestore:"project_id,omitempty"`                     // @gotags: firestore:"project_id,omitempty"
	Locals                []string `json:"locals,omitempty" firestore:"locals,omitempty"`                             // @gotags: firestore:"locals,omitempty"
	SystemGraphqlHooks    bool     `json:"system_graphql_hooks,omitempty" firestore:"system_graphql_hooks,omitempty"` // @gotags: firestore:"system_graphql_hooks,omitempty"
	EnableRevisionHistory bool     `json:"enable_revision_history,omitempty" firestore:"revision_history,omitempty"`  // @gotags: firestore:"revision_history,omitempty"
}

type DriverCredentials struct {
	// for sql migration purposes
	ProjectID string `json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
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
	ProjectID       string  `json:"project_id,omitempty" firestore:"project_id,omitempty"`               // @gotags: firestore:"project_id,omitempty"
	ApiCalls        uint32  `json:"api_calls,omitempty" firestore:"api_calls,omitempty"`                 // @gotags: firestore:"api_calls,omitempty"
	ApiBandwidth    float64 `json:"api_bandwidth,omitempty" firestore:"api_bandwidth,omitempty"`         // @gotags: firestore:"api_bandwidth,omitempty"
	MediaStorage    float64 `json:"media_storage,omitempty" firestore:"media_storage,omitempty"`         // @gotags: firestore:"media_storage,omitempty"
	MediaBandwidth  float64 `json:"media_bandwidth,omitempty" firestore:"media_bandwidth,omitempty"`     // @gotags: firestore:"media_bandwidth,omitempty"
	NumberOfMedia   float64 `json:"number_of_media,omitempty" firestore:"number_of_media,omitempty"`     // @gotags: firestore:"number_of_media,omitempty"
	NumberOfRecords float64 `json:"number_of_records,omitempty" firestore:"number_of_records,omitempty"` // @gotags: firestore:"number_of_records,omitempty"
}

type PluginDetails struct {
	ProjectID        string                  `json:"project_id,omitempty" firestore:"project_id,omitempty"`               // @gotags: firestore:"project_id,omitempty"
	Icon             string                  `json:"icon,omitempty" firestore:"icon,omitempty"`                           // @gotags: firestore:"icon,omitempty"
	Id               string                  `json:"id,omitempty" firestore:"id,omitempty"`                               // @gotags: firestore:"id,omitempty"
	Title            string                  `json:"title,omitempty" firestore:"title,omitempty"`                         // @gotags: firestore:"title,omitempty"
	Serial           int64                   `json:"serial,omitempty" firestore:"serial,omitempty"`                       // @gotags: firestore:"serial,omitempty"
	Version          string                  `json:"version,omitempty" firestore:"version,omitempty"`                     // @gotags: firestore:"version,omitempty"
	Description      string                  `json:"description,omitempty" firestore:"description,omitempty"`             // @gotags: firestore:"description,omitempty"
	Type             *PluginType             `json:"type,omitempty" firestore:"type,omitempty"`                           // @gotags: firestore:"type,omitempty"
	Role             string                  `json:"role,omitempty" firestore:"role,omitempty"`                           // @gotags: firestore:"role,omitempty"
	EnvVars          []*FunctionEnvVariables `json:"env_vars,omitempty" firestore:"env_vars,omitempty"`                   // @gotags: firestore:"env_vars,omitempty"
	ExportedVariable string                  `json:"exported_variable,omitempty" firestore:"exported_variable,omitempty"` // @gotags: firestore:"exported_variable,omitempty"
	Enable           bool                    `json:"enable,omitempty" firestore:"enable,omitempty"`                       // @gotags: firestore:"enable,omitempty"
	RepositoryUrl    string                  `json:"repository_url,omitempty" firestore:"repository_url,omitempty"`       // @gotags: firestore:"repository_url,omitempty"
	Branch           string                  `json:"branch,omitempty" firestore:"branch,omitempty"`                       // @gotags: firestore:"branch,omitempty"
	Author           string                  `json:"author,omitempty" firestore:"author,omitempty"`                       // @gotags: firestore:"author,omitempty"
	LoadStatus       *PluginLoadStatus       `json:"load_status,omitempty" firestore:"load_status,omitempty"`             // @gotags: firestore:"load_status,omitempty"
	ActivateStatus   *PluginActivateStatus   `json:"activate_status,omitempty" firestore:"activate_status,omitempty"`     // @gotags: firestore:"activate_status,omitempty"
}

type APIToken struct {
	ProjectID string `json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	Name      string `json:"name,omitempty" firestore:"name,omitempty"`             // @gotags: firestore:"name,omitempty"
	Token     string `json:"token,omitempty" firestore:"token,omitempty"`           // @gotags: firestore:"token,omitempty"
	Role      string `json:"role,omitempty" firestore:"role,omitempty"`             // @gotags: firestore:"role,omitempty"
	Expire    string `json:"expire,omitempty" firestore:"expire,omitempty"`         // @gotags: firestore:"expire,omitempty"
}

// Project user project
type Project struct {
	XKey                string             `json:"_key,omitempty" firestore:"_key,omitempty"`                                                                                           // @gotags: firestore:"_key,omitempty"
	Id                  string             `bun:"type:uuid,pk" json:"id,omitempty" firestore:"id,omitempty"`                                                                            // @gotags: gorm:"primaryKey;autoIncrement:false" firestore:"id,omitempty"
	ProjectName         string             `json:"project_name,omitempty" firestore:"project_name,omitempty"`                                                                           // @gotags: firestore:"project_name,omitempty"
	ProjectDescription  string             `json:"project_description,omitempty" firestore:"project_description,omitempty"`                                                             // @gotags: firestore:"project_description,omitempty"
	Schema              *ProjectSchema     `json:"schema,omitempty" firestore:"schema,omitempty"`                                                                                       // @gotags: firestore:"schema,omitempty"
	CreatedAt           string             `json:"created_at,omitempty" firestore:"created_at,omitempty"`                                                                               // @gotags: firestore:"created_at,omitempty"
	UpdatedAt           string             `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                                                                               // @gotags: firestore:"updated_at,omitempty"
	ExpireAt            string             `json:"expire_at,omitempty" firestore:"expire_at,omitempty"`                                                                                 // @gotags: firestore:"expire_at,omitempty"
	Plugins             []*PluginDetails   `bun:"rel:has-many" json:"plugins,omitempty" firestore:"plugins,omitempty"`                                                                  // @gotags: firestore:"plugins,omitempty"
	Addons              *AddOnsDetails     `bun:"rel:belongs-to,join:id=project_id" json:"addons,omitempty"  firestore:"addons,omitempty"`                                              // @gotags: gorm:"foreignKey:ProjectID" firestore:"addons,omitempty"
	Tokens              []*APIToken        `bun:"rel:has-many" json:"tokens,omitempty" firestore:"tokens,omitempty"`                                                                    // @gotags: gorm:"foreignKey:ProjectID" firestore:"tokens,omitempty"
	Roles               map[string]*Role   `json:"roles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"roles,omitempty"` // @gotags: firestore:"roles,omitempty"
	Driver              *DriverCredentials `bun:"rel:belongs-to,join:id=project_id" json:"driver,omitempty"  firestore:"driver,omitempty"`                                              // @gotags: gorm:"foreignKey:ProjectID" firestore:"driver,omitempty"
	TempBanned          bool               `json:"temp_banned,omitempty" firestore:"temp_banned,omitempty"`                                                                             // @gotags: firestore:"temp_banned,omitempty"
	Plan                string             `json:"plan,omitempty" firestore:"plan,omitempty"`                                                                                           // @gotags: firestore:"plan,omitempty"
	TrialEnds           string             `json:"trial_ends,omitempty" firestore:"trial_ends,omitempty"`                                                                               // @gotags: firestore:"trial_ends,omitempty"
	FromExample         string             `json:"from_example,omitempty" firestore:"from_example,omitempty"`                                                                           // @gotags: firestore:"from_example,omitempty"
	Limits              *UsagesTracking    `bun:"rel:belongs-to,join:id=project_id" json:"limits,omitempty"  firestore:"limits,omitempty"`                                              // @gotags: gorm:"foreignKey:ProjectID" firestore:"limits,omitempty"
	OwnerId             string             `json:"owner_id,omitempty" firestore:"owner_id,omitempty"`                                                                                   // @gotags: firestore:"owner_id,omitempty"
	Teams               []*TeamMembers     `bun:"rel:has-many" json:"teams,omitempty"  firestore:"teams,omitempty"`                                                                     // @gotags: gorm:"foreignKey:ProjectID" firestore:"teams,omitempty"
	SystemMessages      []*SystemMessage   `bun:"rel:has-many" json:"system_messages,omitempty" firestore:"system_messages,omitempty"`                                                  // @gotags: firestore:"system_messages,omitempty"
	Workspaces          []*Workspace       `bun:"rel:has-many" json:"workspaces,omitempty" firestore:"workspaces,omitempty"`                                                            // @gotags: firestore:"workspaces,omitempty"
	ActivatedPluginsIds *ActivatedPlugins  `bun:"rel:belongs-to,join:id=project_id" json:"activated_plugins_ids,omitempty" firestore:"activated_plugins_ids,omitempty"`                 // @gotags: firestore:"activated_plugins_ids,omitempty"
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
	ProjectID string           `json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	Models    []*ModelType     `json:"models,omitempty" firestore:"models,omitempty"`         // @gotags: firestore:"models,omitempty"
	Functions []*CloudFunction `json:"functions,omitempty" firestore:"functions,omitempty"`   // @gotags: firestore:"functions,omitempty"
}

type ModelType struct {
	Name            string            `json:"name,omitempty" firestore:"name,omitempty"` // @gotags: firestore:"name,omitempty"
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

type CloudFunction struct {
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

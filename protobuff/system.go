package protobuff

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

type SystemUser struct {
	XKey string `json:"_key,omitempty" firestore:"_key,omitempty"`
	Id   string `bun:"type:uuid,pk" json:"id,omitempty" firestore:"id,omitempty"`

	FirstName string `json:"first_name,omitempty" firestore:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty" firestore:"last_name,omitempty"`
	//Role             string `json:"role,omitempty" firestore:"role,omitempty"`
	Username         string `json:"username,omitempty" firestore:"username,omitempty"`
	Email            string `json:"email,omitempty" firestore:"email,omitempty"`
	Secret           string `json:"secret,omitempty" firestore:"secret,omitempty"`
	Avatar           string `json:"avatar,omitempty" firestore:"avatar,omitempty"`
	CurrentProjectId string `json:"current_project_id,omitempty" firestore:"current_project_id,omitempty"`
	RegisterProvider string `json:"register_provider,omitempty" firestore:"register_provider,omitempty"`

	ProjectUser bool `json:"project_user,omitempty" firestore:"project_user,omitempty"`

	RefreshToken string `json:"refresh_token,omitempty" firestore:"refresh_token,omitempty"`
	AccessToken  string `json:"access_token,omitempty" firestore:"access_token,omitempty"`
	LastLoggedIn string `json:"last_logged_in,omitempty" firestore:"last_logged_in,omitempty"`

	CreatedAt string `bun:"type:timestamp,notnull,default:current_timestamp" json:"created_at,omitempty" firestore:"created_at,omitempty"`
	UpdatedAt string `bun:"type:timestamp,notnull" json:"updated_at,omitempty" firestore:"updated_at,omitempty"`
}

type DriverCredentials struct {
	// for sql migration purposes
	ProjectID string `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"`
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

type PluginDetails struct {
	ProjectID        string                  `json:"project_id,omitempty" firestore:"project_id,omitempty"`
	Icon             string                  `json:"icon,omitempty" firestore:"icon,omitempty"`
	Id               string                  `json:"id,omitempty" firestore:"id,omitempty"`
	Title            string                  `json:"title,omitempty" firestore:"title,omitempty"`
	Serial           int64                   `json:"serial,omitempty" firestore:"serial,omitempty"`
	Version          string                  `json:"version,omitempty" firestore:"version,omitempty"`
	Description      string                  `json:"description,omitempty" firestore:"description,omitempty"`
	Type             PluginType              `json:"type,omitempty" firestore:"type,omitempty"`
	EnvVars          []*FunctionEnvVariables `json:"env_vars,omitempty" firestore:"env_vars,omitempty"`
	ExportedVariable string                  `json:"exported_variable,omitempty" firestore:"exported_variable,omitempty"`
	Enable           bool                    `json:"enable,omitempty" firestore:"enable,omitempty"`
	RepositoryUrl    string                  `json:"repository_url,omitempty" firestore:"repository_url,omitempty"`
	Branch           string                  `json:"branch,omitempty" firestore:"branch,omitempty"`
	Author           string                  `json:"author,omitempty" firestore:"author,omitempty"`
	LoadStatus       PluginLoadStatus        `json:"load_status,omitempty" firestore:"load_status,omitempty"`
	ActivateStatus   PluginActivateStatus    `json:"activate_status,omitempty" firestore:"activate_status,omitempty"`
}

type APIToken struct {
	ProjectID string `json:"project_id,omitempty" firestore:"project_id,omitempty"`
	Name      string `json:"name,omitempty" firestore:"name,omitempty"`
	Token     string `json:"token,omitempty" firestore:"token,omitempty"`
	Expire    string `json:"expire,omitempty" firestore:"expire,omitempty"`
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

type SystemMessage struct {
	ProjectID   string `json:"project_id,omitempty" firestore:"project_id,omitempty"`
	Message     string `json:"message,omitempty" firestore:"message,omitempty"`
	Code        string `json:"code,omitempty" firestore:"code,omitempty"`
	Redirection string `json:"redirection,omitempty" firestore:"redirection,omitempty"`
	Hide        bool   `json:"hide,omitempty" firestore:"hide,omitempty"`
}

// Project user project
type Project struct {
	XKey    string `json:"_key,omitempty" firestore:"_key,omitempty"`
	Id      string `bun:"type:uuid,pk" json:"id,omitempty" firestore:"id,omitempty"`
	OwnerId string `json:"owner_id,omitempty" firestore:"owner_id,omitempty"`

	Name        string         `json:"name,omitempty" firestore:"name,omitempty"`
	Description string         `json:"description,omitempty" firestore:"description,omitempty"`
	Schema      *ProjectSchema `bun:"rel:belongs-to,join:id=project_id" json:"schema,omitempty" firestore:"schema,omitempty"`

	ExpireAt string           `json:"expire_at,omitempty" firestore:"expire_at,omitempty"`
	Plugins  []*PluginDetails `bun:"rel:has-many" json:"plugins,omitempty" firestore:"plugins,omitempty"`
	Tokens   []*APIToken      `bun:"rel:has-many" json:"tokens,omitempty" firestore:"tokens,omitempty"`

	ProjectTemplate string           `json:"project_template,omitempty" firestore:"project_template,omitempty"`
	SystemMessages  []*SystemMessage `bun:"rel:has-many" json:"system_messages,omitempty" firestore:"system_messages,omitempty"`
	//ActivatedPluginsIds *ActivatedPlugins `bun:"rel:belongs-to,join:id=project_id" json:"activated_plugins_ids,omitempty" firestore:"activated_plugins_ids,omitempty"`

	DefaultStoragePlugin  string `json:"default_storage_plugin,omitempty" firestore:"default_storage_plugin,omitempty"`
	DefaultFunctionPlugin string `json:"default_function_plugin,omitempty" firestore:"default_function_plugin,omitempty"`

	Locals []string `json:"locals,omitempty" firestore:"locals,omitempty"`

	CreatedAt string `json:"created_at,omitempty" firestore:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`
}

type ProjectSchema struct {
	ProjectID string           `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"`
	Models    []*ModelType     `bun:"rel:has-many" json:"models,omitempty" firestore:"models,omitempty"`
	Functions []*CloudFunction `bun:"rel:has-many" json:"functions,omitempty" firestore:"functions,omitempty"`
}

type ModelType struct {
	ProjectID       string            `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"`
	Name            string            `json:"name,omitempty" firestore:"name,omitempty"`
	Description     string            `json:"description,omitempty"`
	Fields          []*FieldInfo      `json:"fields,omitempty" firestore:"fields,omitempty"`
	Connections     []*ConnectionType `json:"connections,omitempty" firestore:"connections,omitempty"`
	HookIds         []string          `json:"hook_ids,omitempty" firestore:"hook_ids,omitempty"`
	Locals          []string          `json:"locals,omitempty" firestore:"locals,omitempty"`
	RepeatedGroups  []string          `json:"repeated_groups,omitempty" firestore:"locals,omitempty"`
	SystemGenerated bool              `json:"system_generated,omitempty" firestore:"system_generated,omitempty"`
	SinglePage      bool              `json:"single_page,omitempty" firestore:"system_generated,omitempty"`
	SinglePageUuid  string            `json:"single_page_uuid,omitempty" firestore:"system_generated,omitempty"`
	HasConnections  bool              `json:"has_connections,omitempty" firestore:"has_connections,omitempty"`
}

type FunctionEnvVariables struct {
	Key   string `json:"key,omitempty" firestore:"key,omitempty"`
	Value string `json:"value,omitempty" firestore:"value,omitempty"`
}

type CloudFunctionRequestResponseType struct {
	Model  string       `json:"model,omitempty" firestore:"model,omitempty"`
	Params []*FieldInfo `json:"params,omitempty" firestore:"params,omitempty"`
}

type CloudFunction struct {
	ProjectID                string                            `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"`
	Id                       string                            `json:"id,omitempty" firestore:"id,omitempty"`
	Name                     string                            `json:"name,omitempty" firestore:"name,omitempty"`
	Description              string                            `json:"description,omitempty" firestore:"description,omitempty"`
	EnvVars                  []*FunctionEnvVariables           `json:"env_vars,omitempty" firestore:"env_vars,omitempty"`
	FunctionPath             string                            `json:"function_path,omitempty" firestore:"function_path,omitempty"`
	Request                  *CloudFunctionRequestResponseType `json:"request,omitempty" firestore:"request,omitempty"`
	Response                 *CloudFunctionRequestResponseType `json:"response,omitempty" firestore:"response,omitempty"`
	FunctionProviderId       string                            `json:"function_provider_id,omitempty" firestore:"function_provider_id,omitempty"`
	FunctionConnected        bool                              `json:"function_connected,omitempty" firestore:"function_connected,omitempty"`
	ProviderExportedVariable string                            `json:"provider_exported_variable,omitempty" firestore:"provider_exported_variable,omitempty"`
	FunctionExportedVariable string                            `json:"function_exported_variable,omitempty" firestore:"function_exported_variable,omitempty"`
	RuntimeConfig            *FunctionRuntimeConfig            `json:"runtime_config,omitempty" firestore:"runtime_config,omitempty"`
	UpdatedAt                string                            `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`
	CreatedAt                string                            `json:"created_at,omitempty" firestore:"created_at,omitempty"`
}

type FunctionRuntimeConfig struct {
	Runtime string `json:"runtime,omitempty" firestore:"runtime,omitempty"`
	Memory  int64  `json:"memory,omitempty" firestore:"memory,omitempty"`
	Handler string `json:"handler,omitempty" firestore:"handler,omitempty"`
	TimeOut int64  `json:"time_out,omitempty" firestore:"time_out,omitempty"`
}

type Validation struct {
	Required          bool      `json:"required,omitempty" firestore:"required,omitempty"`
	Hide              bool      `json:"hide,omitempty" firestore:"hide,omitempty"`
	AsTitle           bool      `json:"as_title,omitempty" firestore:"as_title,omitempty"`
	Locals            []string  `json:"locals,omitempty" firestore:"locals,omitempty"`
	Unique            bool      `json:"unique,omitempty" firestore:"unique,omitempty"`
	CharLimit         []uint32  `json:"char_limit,omitempty" firestore:"char_limit,omitempty"`
	IntRangeLimit     []uint32  `json:"int_range_limit,omitempty" firestore:"int_range_limit,omitempty"`
	DoubleRangeLimit  []float64 `json:"double_range_limit,omitempty" firestore:"double_range_limit,omitempty"`
	IsEmail           bool      `json:"is_email,omitempty" firestore:"is_email,omitempty"`
	Placeholder       string    `json:"placeholder,omitempty" firestore:"placeholder,omitempty"`
	FixedListElements []string  `json:"fixed_list_elements,omitempty" firestore:"fixed_list_elements,omitempty"`
	IsMultiChoice     bool      `json:"is_multi_choice,omitempty" firestore:"is_multi_choice,omitempty"`
	ListType          string    `json:"list_type,omitempty" firestore:"list_type,omitempty"`
	IsGallery         bool      `json:"is_gallery,omitempty" firestore:"is_gallery,omitempty"`
	IsPassword        bool      `json:"is_password,omitempty" firestore:"is_gallery,omitempty"`
	IsSystemRole      bool      `json:"is_system_role,omitempty" firestore:"is_gallery,omitempty"`
	IsUrl             bool      `json:"is_url,omitempty" firestore:"is_url,omitempty"`
}

type FieldInfo struct {
	Identifier              string       `json:"identifier,omitempty" firestore:"identifier,omitempty"`
	Description             string       `json:"description,omitempty" firestore:"description,omitempty"`
	InputType               string       `json:"input_type,omitempty" firestore:"input_type,omitempty"`
	FieldType               string       `json:"field_type,omitempty" firestore:"field_type,omitempty"`
	FieldSubType            string       `json:"field_sub_type,omitempty"`
	SubFieldInfo            []*FieldInfo `json:"sub_field_info,omitempty" firestore:"modules,omitempty"`
	Validation              *Validation  `json:"validation,omitempty" firestore:"validation,omitempty"`
	Serial                  uint32       `json:"serial,omitempty" firestore:"serial,omitempty"`
	Label                   string       `json:"label,omitempty" firestore:"label,omitempty"`
	SystemGenerated         bool         `json:"system_generated,omitempty" firestore:"system_generated,omitempty"`
	RepeatedGroupIdentifier string       `json:"repeated_group_identifier,omitempty" firestore:"repeated_group_identifier,omitempty"`
	IsObjectField           bool         `json:"is_object_field,omitempty" firestore:"is_object_field,omitempty"`
	ParentField             string       `json:"parent_field,omitempty" firestore:"parent_field,omitempty"`
}

type ConnectionType struct {
	Model    string `json:"model,omitempty" firestore:"model,omitempty"`
	Relation string `json:"relation,omitempty" firestore:"relation,omitempty"`
	Type     string `json:"type,omitempty" firestore:"type,omitempty"`
	KnownAs  string `json:"known_as,omitempty" firestore:"known_as,omitempty"`
}

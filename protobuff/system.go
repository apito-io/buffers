package protobuff

import (
	"encoding/json"
	"github.com/uptrace/bun"
)

type Subscriber struct {
	Data   chan interface{}
	UserID string
}

type SubscriptionEvent struct {
	Type      string `json:"type"`
	ProjectID string `json:"project_id"`
	UserID    string `json:"user_id"`
	Message   string `json:"message"`
}

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
	Read   string `json:"read,omitempty" firestore:"read,omitempty"`     
	Create string `json:"create,omitempty" firestore:"create,omitempty"` 
	Update string `json:"update,omitempty" firestore:"update,omitempty"` 
	Delete string `json:"delete,omitempty" firestore:"delete,omitempty"` 
}

type Role struct {
	bun.BaseModel             `bun:"table:role,alias:r"`
	Id                        string                    `bun:"id,pk,notnull,type:uuid,default:gen_random_uuid()" json:"id,omitempty" firestore:"id,omitempty"` 
	ApiPermissions            map[string]*APIPermission `bun:",type:jsonb" json:"api_permissions,omitempty" firestore:"permissions,omitempty"`                 
	AdministrativePermissions []string                  `json:"administrative_permissions,omitempty" firestore:"administrative_permissions,omitempty"`         
	LogicExecutions           []string                  `json:"logic_executions,omitempty" firestore:"logic_executions,omitempty"`                             
	SystemGenerated           bool                      `json:"system_generated,omitempty" firestore:"system_generated,omitempty"`                             
	IsAdmin                   bool                      `json:"is_admin,omitempty" firestore:"is_admin,omitempty"`                                             
	IsProjectUser             bool                      `json:"is_project_user,omitempty" firestore:"is_project_user,omitempty"`                               
	ReadOnlyProject           bool                      `json:"read_only_project,omitempty" firestore:"read_only_project,omitempty"`                           
	IsSuperAdmin              bool                      `json:"is_super_admin,omitempty" firestore:"read_only_project,omitempty"`                              
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
	XKey string `json:"_key,omitempty" firestore:"_key,omitempty"`                
	Id   string `bun:"type:uuid,pk" json:"id,omitempty" firestore:"id,omitempty"` 

	FirstName        string `json:"first_name,omitempty" firestore:"first_name,omitempty"`                 
	LastName         string `json:"last_name,omitempty" firestore:"last_name,omitempty"`                   
	Role             string `json:"role,omitempty" firestore:"role,omitempty"`                             
	Username         string `json:"username,omitempty" firestore:"username,omitempty"`                     
	Email            string `json:"email,omitempty" firestore:"email,omitempty"`                           
	Secret           string `json:"secret,omitempty" firestore:"secret,omitempty"`                         
	Avatar           string `json:"avatar,omitempty" firestore:"avatar,omitempty"`                         
	CurrentProjectId string `json:"current_project_id,omitempty" firestore:"current_project_id,omitempty"` 
	RegisterProvider string `json:"register_provider,omitempty" firestore:"register_provider,omitempty"`   

	ProjectUser               bool     `json:"project_user,omitempty" firestore:"project_user,omitempty"`        
	AdministrativePermissions []string `json:"administrative_permissions,omitempty" firestore:"email,omitempty"` 

	ProjectAssignedRole      string   `json:"assigned_role,omitempty"`
	ProjectAccessPermissions []string `json:"access_permissions,omitempty"`

	OrganizationID string `json:"organization_id,omitempty" firestore:"organization_id,omitempty"` 

	IsAdmin      bool `json:"is_admin,omitempty" firestore:"is_admin,omitempty"`             
	IsSuperAdmin bool `json:"is_super_admin,omitempty" firestore:"is_super_admin,omitempty"` 

	RefreshToken    string `json:"refresh_token,omitempty" firestore:"refresh_token,omitempty"`         
	AccessToken     string `json:"access_token,omitempty" firestore:"access_token,omitempty"`           
	ReadOnlyProject bool   `json:"read_only_project,omitempty" firestore:"read_only_project,omitempty"` 
	LastLoggedIn    string `json:"last_logged_in,omitempty" firestore:"last_logged_in,omitempty"`       

	ProjectLimit         uint32               `json:"project_limit,omitempty" firestore:"project_limit,omitempty"`                   
	UserSubscriptionType UserSubscriptionType `json:"user_subscription_type,omitempty" firestore:"user_subscription_type,omitempty"` 

	//Projects      []*Project      `bun:"rel:has-many,join:id=owner_id" json:"projects,omitempty" firestore:"projects,omitempty"` 
	Teams []*Team `bun:"m2m:user_to_teams,join:SystemUser=Team" json:"organizations,omitempty" firestore:"organizations,omitempty"` 

	Organization []*Organization `bun:"rel:has-many,join:id=user_id" json:"organization,omitempty" firestore:"organization,omitempty"` 

	CreatedAt string `bun:"type:timestamp,notnull,default:current_timestamp" json:"created_at,omitempty" firestore:"created_at,omitempty"` 
	UpdatedAt string `bun:"type:timestamp,notnull" json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                           

	IsPaymentDue bool `json:"is_payment_due,omitempty" firestore:"is_payment_due,omitempty"` 
}

type Organization struct {
	XKey        string `json:"_key,omitempty" firestore:"_key,omitempty"`                
	Id          string `bun:"type:uuid,pk" json:"id,omitempty" firestore:"id,omitempty"` 
	UserID      string `json:"user_id,omitempty" firestore:"user_id,omitempty"`          
	Name        string `json:"name,omitempty" firestore:"name,omitempty"`                
	Description string `json:"description,omitempty" firestore:"description,omitempty"`  
}

type AuditLogs struct {
	XKey      string `json:"_key,omitempty" firestore:"_key,omitempty"` 
	Id        string `json:"id,omitempty" firestore:"id,omitempty"`     
	UserID    string `json:"user_id,omitempty" firestore:"user_id,omitempty"`
	ProjectID string `json:"project_id,omitempty" firestore:"project_id,omitempty"`

	RequestPayload  string `json:"request_payload,omitempty" firestore:"request_payload,omitempty"`
	RequestPath     string `json:"request_path,omitempty" firestore:"request_path,omitempty"`
	ResponseCode    int    `json:"response_code,omitempty" firestore:"response_code,omitempty"`
	ResponsePayload string `json:"response_payload,omitempty" firestore:"response_payload,omitempty"`

	Activity         string `json:"activity,omitempty" firestore:"activity,omitempty"`
	InternalFunction string `json:"internal_function,omitempty" firestore:"internal_function,omitempty"`

	GraphqlOperationName string `json:"graphql_operation_name,omitempty" firestore:"graphql_operation_name,omitempty"`
	GraphqlPayload       string `json:"graphql_payload,omitempty" firestore:"graphql_payload,omitempty"`
	GraphqlVariable      string `json:"graphql_variable,omitempty" firestore:"graphql_variable,omitempty"`

	InternalError string `json:"internal_error,omitempty" firestore:"internal_error,omitempty"`
	CreatedAt     string `json:"created_at,omitempty" firestore:"created_at,omitempty"`
}

type UserToTeams struct {
	UserID     string      `json:"user_id,omitempty" firestore:"user_id,omitempty"`                                              
	SystemUser *SystemUser `bun:"rel:belongs-to,join:user_id=id" json:"system_user,omitempty" firestore:"system_user,omitempty"` 
	TeamID     string      `json:"team_id,omitempty" firestore:"team_id,omitempty"`                                              
	Team       *Team       `bun:"rel:belongs-to,join:team_id=id" json:"team,omitempty" firestore:"team,omitempty"`               
}

type AddOnsDetails struct {
	ProjectID             string   `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"`  
	Locals                []string `json:"locals,omitempty" firestore:"locals,omitempty"`                             
	SystemGraphqlHooks    bool     `json:"system_graphql_hooks,omitempty" firestore:"system_graphql_hooks,omitempty"` 
	EnableRevisionHistory bool     `json:"enable_revision_history,omitempty" firestore:"revision_history,omitempty"`  
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

type UsagesTracking struct {
	// for sql migration purposes
	ProjectID       string  `bun:"type:uuid,pk" json:"project_id,omitempty" firestore:"project_id,omitempty"` 
	ApiCalls        uint32  `json:"api_calls,omitempty" firestore:"api_calls,omitempty"`                      
	ApiBandwidth    float64 `json:"api_bandwidth,omitempty" firestore:"api_bandwidth,omitempty"`              
	MediaStorage    float64 `json:"media_storage,omitempty" firestore:"media_storage,omitempty"`              
	MediaBandwidth  float64 `json:"media_bandwidth,omitempty" firestore:"media_bandwidth,omitempty"`          
	NumberOfMedia   float64 `json:"number_of_media,omitempty" firestore:"number_of_media,omitempty"`          
	NumberOfRecords float64 `json:"number_of_records,omitempty" firestore:"number_of_records,omitempty"`      
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
	Role             string                  `json:"role,omitempty" firestore:"role,omitempty"`                           
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
	Role      string `json:"role,omitempty" firestore:"role,omitempty"`             
	Expire    string `json:"expire,omitempty" firestore:"expire,omitempty"`         
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

type Workspace struct {
	ProjectID    string `json:"project_id,omitempty" firestore:"project_id,omitempty"`       
	Name         string `json:"name,omitempty" firestore:"name,omitempty"`                   
	Active       bool   `json:"active,omitempty" firestore:"active,omitempty"`               
	IsProduction bool   `json:"is_production,omitempty" firestore:"is_production,omitempty"` 
	IsDefault    bool   `json:"is_default,omitempty" firestore:"is_default,omitempty"`       
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
	XKey           string             `json:"_key,omitempty" firestore:"_key,omitempty"` 
	Id             string             `bun:"type:uuid,pk" json:"id,omitempty" firestore:"id,omitempty"`
	OwnerId        string             `json:"owner_id,omitempty" firestore:"owner_id,omitempty"`
	OrganizationID string             `json:"organization_id,omitempty" firestore:"organization_id,omitempty"`                                                                                      
	Name           string             `json:"name,omitempty" firestore:"name,omitempty"`                                                                                                            
	Description    string             `json:"description,omitempty" firestore:"description,omitempty"`                                                                                              
	Schema         *ProjectSchema     `bun:"rel:belongs-to,join:id=project_id" json:"schema,omitempty" firestore:"schema,omitempty"`                                                                
	CreatedAt      string             `json:"created_at,omitempty" firestore:"created_at,omitempty"`                                                                                                
	UpdatedAt      string             `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                                                                                                
	ExpireAt       string             `json:"expire_at,omitempty" firestore:"expire_at,omitempty"`                                                                                                  
	Plugins        []*PluginDetails   `bun:"rel:has-many" json:"plugins,omitempty" firestore:"plugins,omitempty"`                                                                                   
	Addons         *AddOnsDetails     `bun:"rel:belongs-to,join:id=project_id" json:"addons,omitempty"  firestore:"addons,omitempty"`                                                               // @gotags: gorm:"foreignKey:ProjectID" firestore:"addons,omitempty"
	Tokens         []*APIToken        `bun:"rel:has-many" json:"tokens,omitempty" firestore:"tokens,omitempty"`                                                                                     // @gotags: gorm:"foreignKey:ProjectID" firestore:"tokens,omitempty"
	Roles          map[string]*Role   `bun:"type:jsonb" json:"roles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"roles,omitempty"` 
	Driver         *DriverCredentials `bun:"rel:belongs-to,join:id=project_id" json:"driver,omitempty"  firestore:"driver,omitempty"`                                                               // @gotags: gorm:"foreignKey:ProjectID" firestore:"driver,omitempty"
	TempBanned     bool               `json:"temp_banned,omitempty" firestore:"temp_banned,omitempty"`                                                                                              

	TrialEnds       string          `json:"trial_ends,omitempty" firestore:"trial_ends,omitempty"`                                  
	ProjectTemplate string          `json:"project_template,omitempty" firestore:"project_template,omitempty"`                      
	Limits          *UsagesTracking `bun:"rel:belongs-to,join:id=project_id" json:"limits,omitempty"  firestore:"limits,omitempty"` // @gotags: gorm:"foreignKey:ProjectID" firestore:"limits,omitempty"

	Teams          []*SystemUser    `bun:"rel:has-many" json:"teams,omitempty"  firestore:"teams,omitempty"`                    // @gotags: gorm:"foreignKey:ProjectID" firestore:"teams,omitempty"
	SystemMessages []*SystemMessage `bun:"rel:has-many" json:"system_messages,omitempty" firestore:"system_messages,omitempty"` 
	Workspaces     []*Workspace     `bun:"rel:has-many" json:"workspaces,omitempty" firestore:"workspaces,omitempty"`           
	//ActivatedPluginsIds *ActivatedPlugins `bun:"rel:belongs-to,join:id=project_id" json:"activated_plugins_ids,omitempty" firestore:"activated_plugins_ids,omitempty"` 

	DefaultStoragePlugin  string `json:"default_storage_plugin,omitempty" firestore:"default_storage_plugin,omitempty"`   
	DefaultFunctionPlugin string `json:"default_function_plugin,omitempty" firestore:"default_function_plugin,omitempty"` 

	//UserSubscriptionType `json:"user_subscription_type,omitempty" firestore:"user_subscription_type,omitempty"` 

	// for microservice
	MicroServicePort string `json:"micro_service_port,omitempty" firestore:"micro_service_port,omitempty"`
}

type Team struct {
	XKey           string `json:"_key,omitempty" firestore:"_key,omitempty"`
	ID             string `bun:"type:uuid,pk" json:"id,omitempty" firestore:"id,omitempty"`        
	OrganizationID string `json:"organization_id,omitempty" firestore:"organization_id,omitempty"` 
	CreatedBy      string `json:"created_by,omitempty" firestore:"created_by,omitempty"`           

	Name        string `json:"name,omitempty" firestore:"name,omitempty"`               
	Description string `json:"description,omitempty" firestore:"description,omitempty"` 
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
	EnvVars                  []*FunctionEnvVariables           `json:"env_vars,omitempty" firestore:"env_vars,omitempty"`                                     `
	FunctionPath             string                            `json:"function_path,omitempty" firestore:"function_path,omitempty"`                           `
	Request                  *CloudFunctionRequestResponseType `json:"request,omitempty" firestore:"request,omitempty"`                                       `
	Response                 *CloudFunctionRequestResponseType `json:"response,omitempty" firestore:"response,omitempty"`                                     `
	FunctionProviderId       string                            `json:"function_provider_id,omitempty" firestore:"function_provider_id,omitempty"`             `
	FunctionConnected        bool                              `json:"function_connected,omitempty" firestore:"function_connected,omitempty"`                 `
	ProviderExportedVariable string                            `json:"provider_exported_variable,omitempty" firestore:"provider_exported_variable,omitempty"` `
	FunctionExportedVariable string                            `json:"function_exported_variable,omitempty" firestore:"function_exported_variable,omitempty"` `
	RuntimeConfig            *FunctionRuntimeConfig            `json:"runtime_config,omitempty" firestore:"runtime_config,omitempty"`                         `
	UpdatedAt                string                            `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                                 `
	CreatedAt                string                            `json:"created_at,omitempty" firestore:"created_at,omitempty"`                                 // @got
}

type ProjectWithRoles struct {
	Project     *Project `json:"project"`
	Role        string   `json:"role"`
	Permissions []string `json:"permissions"`
}

type ProjectInvoices struct {
	XKey                  string  `json:"_key,omitempty" firestore:"_key,omitempty"`                                                                                                
	Id                    string  `json:"id,omitempty" firestore:"id,omitempty"`                                                                                                    
	Type                  string  `json:"type,omitempty" firestore:"type,omitempty"`                                                                                                
	SubscriptionId        string  `json:"subscription_id,omitempty" firestore:"subscription_id,omitempty"`                                                                          
	CouponCode            string  `json:"coupon_code,omitempty" firestore:"coupon_code,omitempty"`                                                                                  
	DiscountAmount        float64 `protobuf:"fixed64,6,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty" firestore:"discount_amount,omitempty"` 
	PaymentGateway        string  `json:"payment_gateway,omitempty" firestore:"payment_gateway,omitempty"`                                                                          
	ReceiptUrl            string  `json:"receipt_url,omitempty" firestore:"receipt_url,omitempty"`                                                                                  
	ProductAmount         float64 `json:"product_amount,omitempty" firestore:"product_amount,omitempty"`                                                                            
	DiscountPercentage    int64   `json:"discount_percentage,omitempty" firestore:"discount_percentage,omitempty"`                                                                  
	Quantity              uint32  `json:"quantity,omitempty" firestore:"quantity,omitempty"`                                                                                        
	PaymentMethod         string  `json:"payment_method,omitempty" firestore:"payment_method,omitempty"`                                                                            
	Earning               float64 `json:"earning,omitempty" firestore:"earning,omitempty"`                                                                                          
	Fee                   float64 `json:"fee,omitempty" firestore:"fee,omitempty"`                                                                                                  
	Tax                   float64 `json:"tax,omitempty" firestore:"tax,omitempty"`                                                                                                  
	CreatedAt             string  `json:"created_at,omitempty" firestore:"created_at,omitempty"`                                                                                    
	ChargedAmount         float64 `json:"charged_amount,omitempty" firestore:"charged_amount,omitempty"`                                                                            
	NextAmountToBeCharged float64 `json:"next_amount_to_be_charged,omitempty" firestore:"next_amount_to_be_charged,omitempty"`                                                      
	SubscriptionPlanCode  string  `json:"subscription_plan_code,omitempty" firestore:"subscription_plan_code,omitempty"`                                                            
}

type AccountUsage struct {
	Id              string                 `json:"id,omitempty" firestore:"id,omitempty"`                               
	XKey            string                 `json:"_key,omitempty" firestore:"_key,omitempty"`                           
	Type            string                 `json:"type,omitempty" firestore:"type,omitempty"`                           
	Subscriptions   []*MonthlySubscription `json:"subscriptions,omitempty" firestore:"subscriptions,omitempty"`         
	Limits          *UsagesTracking        `json:"limits,omitempty" firestore:"limits,omitempty"`                       
	Usages          []*UsagesTracking      `json:"usages,omitempty" firestore:"usages,omitempty"`                       
	NumberOfProject uint32                 `json:"number_of_project,omitempty" firestore:"number_of_project,omitempty"` 
}

type MonthlySubscription struct {
	Id                 string             `json:"id,omitempty" firestore:"id,omitempty"`                                   
	XKey               string             `json:"_key,omitempty" firestore:"_key,omitempty"`                               
	Type               string             `json:"type,omitempty" firestore:"type,omitempty"`                               
	UserId             string             `json:"user_id,omitempty" firestore:"user_id,omitempty"`                         
	SubscriptionCode   string             `json:"subscription_code,omitempty" firestore:"plan_code,omitempty"`             
	StartDate          string             `json:"start_date,omitempty" firestore:"start_date,omitempty"`                   
	EndDate            string             `json:"end_date,omitempty" firestore:"end_date,omitempty"`                       
	SubscriptionStatus string             `json:"subscription_status,omitempty" firestore:"subscription_status,omitempty"` 
	PaymentStatus      string             `json:"payment_status,omitempty" firestore:"payment_status,omitempty"`           
	Invoices           []*ProjectInvoices `json:"invoices,omitempty" firestore:"invoices,omitempty"`                       
	CreatedAt          string             `json:"created_at,omitempty" firestore:"created_at,omitempty"`                   
	UpdatedAt          string             `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                   
	Title              string             `json:"title,omitempty" firestore:"title,omitempty"`                             
	CancelUrl          string             `json:"cancel_url,omitempty" firestore:"cancel_url,omitempty"`                   
	ProjectId          string             `json:"project_id,omitempty" firestore:"cancel_url,omitempty"`                   
}

type ProjectUsages struct {
	Id        string          `json:"id,omitempty" firestore:"id,omitempty"`                 
	XKey      string          `json:"_key,omitempty" firestore:"_key,omitempty"`             
	Type      string          `json:"type,omitempty" firestore:"type,omitempty"`             
	ProjectId string          `json:"project_id,omitempty" firestore:"project_id,omitempty"` 
	Year      uint32          `json:"year,omitempty" firestore:"year,omitempty"`             
	Month     uint32          `json:"month,omitempty" firestore:"month,omitempty"`           
	Usages    *UsagesTracking `json:"usages,omitempty" firestore:"usage,omitempty"`          
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

type Webhook struct {
	Id              string   `json:"id,omitempty" firestore:"id,omitempty"`                             
	XKey            string   `json:"_key,omitempty" firestore:"_key,omitempty"`                         
	Type            string   `json:"type,omitempty" firestore:"type,omitempty"`                         
	Model           string   `json:"model,omitempty" firestore:"model,omitempty"`                       
	ProjectId       string   `json:"project_id,omitempty" firestore:"project_id,omitempty"`             //
	Name            string   `json:"name,omitempty" firestore:"name,omitempty"`                         
	Events          []string `json:"events,omitempty" firestore:"events,omitempty"`                     
	Url             string   `json:"url,omitempty" firestore:"url,omitempty"`                           
	LogicExecutions []string `json:"logic_executions,omitempty" firestore:"logic_executions,omitempty"` 
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

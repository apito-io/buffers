package protobuff

type PictureDeleteRequest struct {
	Urls      []string `json:"urls,omitempty" firestore:"urls,omitempty"`             // @gotags: firestore:"urls,omitempty"
	Model     string   `json:"model,omitempty" firestore:"model,omitempty"`           // @gotags: firestore:"model,omitempty"
	Id        string   `json:"id,omitempty" firestore:"id,omitempty"`                 // @gotags: firestore:"id,omitempty"
	FieldName string   `json:"field_name,omitempty" firestore:"field_name,omitempty"` // @gotags: firestore:"field_name,omitempty"
}

type FileDetails struct {
	Id            string        `json:"id,omitempty" firestore:"id,omitempty"`                         // @gotags: firestore:"id,omitempty"
	XKey          string        `json:"_key,omitempty" firestore:"_key,omitempty"`                     // @gotags: firestore:"_key,omitempty"
	Type          string        `json:"type,omitempty" firestore:"type,omitempty"`                     // @gotags: firestore:"type,omitempty"
	FileExtension string        `json:"file_extension,omitempty" firestore:"file_extension,omitempty"` // @gotags: firestore:"file_extension,omitempty"
	FileName      string        `json:"file_name,omitempty" firestore:"file_name,omitempty"`           // @gotags: firestore:"file_name,omitempty"
	ContentType   string        `json:"content_type,omitempty" firestore:"content_type,omitempty"`     // @gotags: firestore:"content_type,omitempty"
	Size          int64         `json:"size,omitempty" firestore:"size,omitempty"`                     // @gotags: firestore:"size,omitempty"
	S3Key         string        `json:"s3_key,omitempty" firestore:"s3_key,omitempty"`                 // @gotags: firestore:"s3_key,omitempty"
	Url           string        `json:"url,omitempty" firestore:"url,omitempty"`                       // @gotags: firestore:"url,omitempty"
	CreatedAt     string        `json:"created_at,omitempty" firestore:"created_at,omitempty"`         // @gotags: firestore:"created_at,omitempty"
	UploadParam   *UploadParams `json:"upload_param,omitempty" firestore:"upload_param,omitempty"`     // @gotags: firestore:"upload_param,omitempty"
	Buffer        []byte        `json:"buffer,omitempty" firestore:"upload_param,omitempty"`           // @gotags: firestore:"upload_param,omitempty"
}

type UploadParams struct {
	DocId      string `json:"doc_id,omitempty" firestore:"doc_id,omitempty"`           // @gotags: firestore:"doc_id,omitempty"
	ProjectId  string `json:"project_id,omitempty" firestore:"project_id,omitempty"`   // @gotags: firestore:"project_id,omitempty"
	ModelName  string `json:"model_name,omitempty" firestore:"model_name,omitempty"`   // @gotags: firestore:"model_name,omitempty"
	FieldName  string `json:"field_name,omitempty" firestore:"field_name,omitempty"`   // @gotags: firestore:"field_name,omitempty"
	AllowMulti bool   `json:"allow_multi,omitempty" firestore:"allow_multi,omitempty"` // @gotags: firestore:"allow_multi,omitempty"
	Provider   string `json:"provider,omitempty" firestore:"provider,omitempty"`       // @gotags: firestore:"provider,omitempty"
}

type Filter struct {
	Page     uint32 `json:"page,omitempty" firestore:"page,omitempty"`         // @gotags: firestore:"page,omitempty"
	Offset   uint32 `json:"offset,omitempty" firestore:"offset,omitempty"`     // @gotags: firestore:"offset,omitempty"
	Limit    uint32 `json:"limit,omitempty" firestore:"limit,omitempty"`       // @gotags: firestore:"limit,omitempty"
	Order    string `json:"order,omitempty" firestore:"order,omitempty"`       // @gotags: firestore:"order,omitempty"
	Min      uint32 `json:"min,omitempty" firestore:"min,omitempty"`           // @gotags: firestore:"min,omitempty"
	Max      uint32 `json:"max,omitempty" firestore:"max,omitempty"`           // @gotags: firestore:"max,omitempty"
	Category string `json:"category,omitempty" firestore:"category,omitempty"` // @gotags: firestore:"category,omitempty"
}

type Request struct {
	Id           string  `json:"id,omitempty" firestore:"id,omitempty"`                       // @gotags: firestore:"id,omitempty"
	Type         string  `json:"type,omitempty" firestore:"type,omitempty"`                   // @gotags: firestore:"type,omitempty"
	Filter       *Filter `json:"filter,omitempty" firestore:"filter,omitempty"`               // @gotags: firestore:"filter,omitempty"
	SearchString string  `json:"search_string,omitempty" firestore:"search_string,omitempty"` // @gotags: firestore:"search_string,omitempty"
	Retry        bool    `json:"retry,omitempty" firestore:"retry,omitempty"`                 // @gotags: firestore:"retry,omitempty"
}

// File Picker
type FileLink struct {
	Link      string `json:"link,omitempty" firestore:"link,omitempty"`             // @gotags: firestore:"link,omitempty"
	Title     string `json:"title,omitempty" firestore:"title,omitempty"`           // @gotags: firestore:"title,omitempty"
	CreatedAt string `json:"created_at,omitempty" firestore:"created_at,omitempty"` // @gotags: firestore:"created_at,omitempty"
	UpdatedAt string `json:"updated_at,omitempty" firestore:"updated_at,omitempty"` // @gotags: firestore:"updated_at,omitempty"
}

type FilePickParameter struct {
	NumberOfImages uint32      `json:"number_of_images,omitempty" firestore:"number_of_images,omitempty"` // @gotags: firestore:"number_of_images,omitempty"
	S3Folder       string      `json:"s3_folder,omitempty" firestore:"s3_folder,omitempty"`               // @gotags: firestore:"s3_folder,omitempty"
	PickerTitle    string      `json:"picker_title,omitempty" firestore:"picker_title,omitempty"`         // @gotags: firestore:"picker_title,omitempty"
	Origin         *SystemUser `json:"origin,omitempty" firestore:"origin,omitempty"`                     // @gotags: firestore:"origin,omitempty"
}

type ImageMetaInfo struct {
	Identifier string `json:"identifier,omitempty" firestore:"identifier,omitempty"` // @gotags: firestore:"identifier,omitempty"
	Name       string `json:"name,omitempty" firestore:"name,omitempty"`             // @gotags: firestore:"name,omitempty"
	Width      uint32 `json:"width,omitempty" firestore:"width,omitempty"`           // @gotags: firestore:"width,omitempty"
	Height     uint32 `json:"height,omitempty" firestore:"height,omitempty"`         // @gotags: firestore:"height,omitempty"
	Type       string `json:"type,omitempty" firestore:"type,omitempty"`             // @gotags: firestore:"type,omitempty"
}

type RegisterRequest struct {
	Username         string `json:"username,omitempty" firestore:"username,omitempty"`            // @gotags: firestore:"username,omitempty"
	Email            string `json:"email,omitempty" firestore:"email,omitempty"`                  // @gotags: firestore:"email,omitempty"
	Secret           string `json:"secret,omitempty" firestore:"secret,omitempty"`                // @gotags: firestore:"secret,omitempty"
	FullName         string `json:"full_name,omitempty" firestore:"full_name,omitempty"`          // @gotags: firestore:"full_name,omitempty"
	Profession       string `json:"profession,omitempty" firestore:"profession,omitempty"`        // @gotags: firestore:"profession,omitempty"
	VerificationCode string `json:"verification_code,omitempty" firestore:"profession,omitempty"` // @gotags: firestore:"profession,omitempty"
}

type LoginRequest struct {
	Username string `json:"username,omitempty" firestore:"username,omitempty"` // @gotags: firestore:"username,omitempty"
	Email    string `json:"email,omitempty" firestore:"email,omitempty"`       // @gotags: firestore:"email,omitempty"
	Secret   string `json:"secret,omitempty" firestore:"secret,omitempty"`     // @gotags: firestore:"secret,omitempty"
}

type PassChangeRequest struct {
	OldPassword string `json:"old_password,omitempty" firestore:"old_password,omitempty"` // @gotags: firestore:"old_password,omitempty"
	NewPassword string `json:"new_password,omitempty" firestore:"new_password,omitempty"` // @gotags: firestore:"new_password,omitempty"
}

type ProjectLimit struct {
	Free  uint32 `protobuf:"varint,1,opt,name=free,proto3" json:"free,omitempty" firestore:"free,omitempty"`    // @gotags: firestore:"free,omitempty"
	Hobby uint32 `protobuf:"varint,2,opt,name=hobby,proto3" json:"hobby,omitempty" firestore:"hobby,omitempty"` // @gotags: firestore:"hobby,omitempty"
	Pro   uint32 `protobuf:"varint,3,opt,name=pro,proto3" json:"pro,omitempty" firestore:"pro,omitempty"`       // @gotags: firestore:"pro,omitempty"
}

type UserSubscription struct {
	Subscription *MonthlySubscription `json:"subscription,omitempty" firestore:"subscription,omitempty"` // @gotags: firestore:"subscription,omitempty"
	Invoices     []*ProjectInvoices   `json:"invoices,omitempty" firestore:"invoices,omitempty"`         // @gotags: firestore:"invoices,omitempty"
}

type UserMeta struct {
	Id                        string   `json:"id,omitempty" firestore:"id,omitempty"`                            // @gotags: firestore:"id,omitempty"
	Avatar                    string   `json:"avatar,omitempty" firestore:"avatar,omitempty"`                    // @gotags: firestore:"avatar,omitempty"
	Name                      string   `json:"name,omitempty" firestore:"name,omitempty"`                        // @gotags: firestore:"name,omitempty"
	Role                      string   `json:"role,omitempty" firestore:"role,omitempty"`                        // @gotags: firestore:"role,omitempty"
	Email                     string   `json:"email,omitempty" firestore:"email,omitempty"`                      // @gotags: firestore:"email,omitempty"
	ProjectUser               bool     `json:"project_user,omitempty" firestore:"email,omitempty"`               // @gotags: firestore:"email,omitempty"
	AdministrativePermissions []string `json:"administrative_permissions,omitempty" firestore:"email,omitempty"` // @gotags: firestore:"email,omitempty"
}

type MetaField struct {
	CreatedAt      string    `json:"created_at,omitempty" firestore:"created_at,omitempty"`             // @gotags: firestore:"created_at,omitempty"
	UpdatedAt      string    `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`             // @gotags: firestore:"updated_at,omitempty"
	CreatedBy      *UserMeta `json:"created_by,omitempty" firestore:"title,omitempty"`                  // @gotags: firestore:"title,omitempty"
	LastModifiedBy *UserMeta `json:"last_modified_by,omitempty" firestore:"created_by,omitempty"`       // @gotags: firestore:"created_by,omitempty"
	Status         string    `json:"status,omitempty" firestore:"status,omitempty"`                     // @gotags: firestore:"status,omitempty"
	TenantId       string    `json:"tenant_id,omitempty" firestore:"tenant_id,omitempty"`               // @gotags: firestore:"tenant_id,omitempty"
	RootRevisionId string    `json:"root_revision_id,omitempty" firestore:"root_revision_id,omitempty"` // @gotags: firestore:"root_revision_id,omitempty"
	Revision       bool      `json:"revision,omitempty" firestore:"revision,omitempty"`                 // @gotags: firestore:"revision,omitempty"
	RevisionAt     string    `json:"revision_at,omitempty" firestore:"revision_at,omitempty"`           // @gotags: firestore:"revision_at,omitempty"
}

type PreviewMode struct {
	Title  string `json:"title,omitempty" firestore:"title,omitempty"`   // @gotags: firestore:"title,omitempty"
	Icon   string `json:"icon,omitempty" firestore:"icon,omitempty"`     // @gotags: firestore:"icon,omitempty"
	Status string `json:"status,omitempty" firestore:"status,omitempty"` // @gotags: firestore:"status,omitempty"
	Id     string `json:"id,omitempty" firestore:"id,omitempty"`         // @gotags: firestore:"id,omitempty"
}

type ActivatedPlugins struct {
	Storage  string `json:"storage,omitempty"`
	Function string `json:"function,omitempty"`
}

type Workspace struct {
	Name         string `json:"name,omitempty" firestore:"name,omitempty"`                   // @gotags: firestore:"name,omitempty"
	Active       bool   `json:"active,omitempty" firestore:"active,omitempty"`               // @gotags: firestore:"active,omitempty"
	IsProduction bool   `json:"is_production,omitempty" firestore:"is_production,omitempty"` // @gotags: firestore:"is_production,omitempty"
	IsDefault    bool   `json:"is_default,omitempty" firestore:"is_default,omitempty"`       // @gotags: firestore:"is_default,omitempty"
}

type SystemMessage struct {
	Message     string `json:"message,omitempty" firestore:"message,omitempty"`         // @gotags: firestore:"message,omitempty"
	Code        string `json:"code,omitempty" firestore:"code,omitempty"`               // @gotags: firestore:"code,omitempty"
	Redirection string `json:"redirection,omitempty" firestore:"redirection,omitempty"` // @gotags: firestore:"redirection,omitempty"
	Hide        bool   `json:"hide,omitempty" firestore:"hide,omitempty"`               // @gotags: firestore:"hide,omitempty"
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
	Description string    `json:"description,omitempty"`
	User        *UserMeta `json:"user,omitempty"`
	CreatedAt   string    `json:"created_at,omitempty"`
	Edited      bool      `json:"edited,omitempty"`
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

type ProjectUsages struct {
	Id        string          `json:"id,omitempty" firestore:"id,omitempty"`                 // @gotags: firestore:"id,omitempty"
	XKey      string          `json:"_key,omitempty" firestore:"_key,omitempty"`             // @gotags: firestore:"_key,omitempty"
	Type      string          `json:"type,omitempty" firestore:"type,omitempty"`             // @gotags: firestore:"type,omitempty"
	ProjectId string          `json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	Year      uint32          `json:"year,omitempty" firestore:"year,omitempty"`             // @gotags: firestore:"year,omitempty"
	Month     uint32          `json:"month,omitempty" firestore:"month,omitempty"`           // @gotags: firestore:"month,omitempty"
	Usages    *UsagesTracking `json:"usages,omitempty" firestore:"usage,omitempty"`          // @gotags: firestore:"usage,omitempty"
}

type InitParams struct {
	ProjectID string             `json:"project_id"`
	ProjectDB *DriverCredentials `json:"system_credentials"`
	CacheDB   *DriverCredentials `json:"cache_db"`
	SharedDB  *DriverCredentials `json:"shared_db"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
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

type FunctionEnvVariables struct {
	Key   string `json:"key,omitempty" firestore:"key,omitempty"`     // @gotags: firestore:"key,omitempty"
	Value string `json:"value,omitempty" firestore:"value,omitempty"` // @gotags: firestore:"value,omitempty"
}

type CloudFunctionRequestResponseType struct {
	Model  string       `json:"model,omitempty" firestore:"model,omitempty"`   // @gotags: firestore:"model,omitempty"
	Params []*FieldInfo `json:"params,omitempty" firestore:"params,omitempty"` // @gotags: firestore:"params,omitempty"
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

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

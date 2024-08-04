package protobuff

type PictureDeleteRequest struct {
	Urls      []string `json:"urls,omitempty" firestore:"urls,omitempty"`
	Model     string   `json:"model,omitempty" firestore:"model,omitempty"`
	Id        string   `json:"id,omitempty" firestore:"id,omitempty"`
	FieldName string   `json:"field_name,omitempty" firestore:"field_name,omitempty"`
}

type FileDetails struct {
	Id            string        `json:"id,omitempty" firestore:"id,omitempty"`
	XKey          string        `json:"_key,omitempty" firestore:"_key,omitempty"`
	Type          string        `json:"type,omitempty" firestore:"type,omitempty"`
	FileExtension string        `json:"file_extension,omitempty" firestore:"file_extension,omitempty"`
	FileName      string        `json:"file_name,omitempty" firestore:"file_name,omitempty"`
	ContentType   string        `json:"content_type,omitempty" firestore:"content_type,omitempty"`
	Size          int64         `json:"size,omitempty" firestore:"size,omitempty"`
	S3Key         string        `json:"s3_key,omitempty" firestore:"s3_key,omitempty"`
	Url           string        `json:"url,omitempty" firestore:"url,omitempty"`
	CreatedAt     string        `json:"created_at,omitempty" firestore:"created_at,omitempty"`
	UploadParam   *UploadParams `json:"upload_param,omitempty" firestore:"upload_param,omitempty"`
	Buffer        []byte        `json:"buffer,omitempty" firestore:"upload_param,omitempty"`
}

type UploadParams struct {
	DocId      string `json:"doc_id,omitempty" firestore:"doc_id,omitempty"`
	ProjectId  string `json:"project_id,omitempty" firestore:"project_id,omitempty"`
	ModelName  string `json:"model_name,omitempty" firestore:"model_name,omitempty"`
	FieldName  string `json:"field_name,omitempty" firestore:"field_name,omitempty"`
	AllowMulti bool   `json:"allow_multi,omitempty" firestore:"allow_multi,omitempty"`
	Provider   string `json:"provider,omitempty" firestore:"provider,omitempty"`
}

type Filter struct {
	Page     uint32 `json:"page,omitempty" firestore:"page,omitempty"`
	Offset   uint32 `json:"offset,omitempty" firestore:"offset,omitempty"`
	Limit    uint32 `json:"limit,omitempty" firestore:"limit,omitempty"`
	Order    string `json:"order,omitempty" firestore:"order,omitempty"`
	Min      uint32 `json:"min,omitempty" firestore:"min,omitempty"`
	Max      uint32 `json:"max,omitempty" firestore:"max,omitempty"`
	Category string `json:"category,omitempty" firestore:"category,omitempty"`
}

type Request struct {
	Id           string  `json:"id,omitempty" firestore:"id,omitempty"`
	Type         string  `json:"type,omitempty" firestore:"type,omitempty"`
	Filter       *Filter `json:"filter,omitempty" firestore:"filter,omitempty"`
	SearchString string  `json:"search_string,omitempty" firestore:"search_string,omitempty"`
	Retry        bool    `json:"retry,omitempty" firestore:"retry,omitempty"`
}

// File Picker
type FileLink struct {
	Link      string `json:"link,omitempty" firestore:"link,omitempty"`
	Title     string `json:"title,omitempty" firestore:"title,omitempty"`
	CreatedAt string `json:"created_at,omitempty" firestore:"created_at,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`
}

type FilePickParameter struct {
	NumberOfImages uint32      `json:"number_of_images,omitempty" firestore:"number_of_images,omitempty"`
	S3Folder       string      `json:"s3_folder,omitempty" firestore:"s3_folder,omitempty"`
	PickerTitle    string      `json:"picker_title,omitempty" firestore:"picker_title,omitempty"`
	Origin         *SystemUser `json:"origin,omitempty" firestore:"origin,omitempty"`
}

type ImageMetaInfo struct {
	Identifier string `json:"identifier,omitempty" firestore:"identifier,omitempty"`
	Name       string `json:"name,omitempty" firestore:"name,omitempty"`
	Width      uint32 `json:"width,omitempty" firestore:"width,omitempty"`
	Height     uint32 `json:"height,omitempty" firestore:"height,omitempty"`
	Type       string `json:"type,omitempty" firestore:"type,omitempty"`
}

type RegisterRequest struct {
	User             *SystemUser `json:"user,omitempty" firestore:"user,omitempty"`
	VerificationCode string      `json:"verification_code,omitempty" firestore:"profession,omitempty"`
	AddedByAdmin     bool        `json:"added_by_admin,omitempty" firestore:"added_by_admin,omitempty"`
}

type LoginRequest struct {
	Username string `json:"username,omitempty" firestore:"username,omitempty"`
	Email    string `json:"email,omitempty" firestore:"email,omitempty"`
	Secret   string `json:"secret,omitempty" firestore:"secret,omitempty"`
}

type PassChangeRequest struct {
	OldPassword string `json:"old_password,omitempty" firestore:"old_password,omitempty"`
	NewPassword string `json:"new_password,omitempty" firestore:"new_password,omitempty"`
}

type MetaField struct {
	CreatedAt      string      `json:"created_at,omitempty" firestore:"created_at,omitempty"`
	UpdatedAt      string      `json:"updated_at,omitempty" firestore:"updated_at,omitempty"`
	CreatedBy      *SystemUser `json:"created_by,omitempty" firestore:"title,omitempty"`
	LastModifiedBy *SystemUser `json:"last_modified_by,omitempty" firestore:"created_by,omitempty"`
	Status         string      `json:"status,omitempty" firestore:"status,omitempty"`
	TenantId       string      `json:"tenant_id,omitempty" firestore:"tenant_id,omitempty"`
	RootRevisionId string      `json:"root_revision_id,omitempty" firestore:"root_revision_id,omitempty"`
	Revision       bool        `json:"revision,omitempty" firestore:"revision,omitempty"`
	RevisionAt     string      `json:"revision_at,omitempty" firestore:"revision_at,omitempty"`
}

type PreviewMode struct {
	Title  string `json:"title,omitempty" firestore:"title,omitempty"`
	Icon   string `json:"icon,omitempty" firestore:"icon,omitempty"`
	Status string `json:"status,omitempty" firestore:"status,omitempty"`
	Id     string `json:"id,omitempty" firestore:"id,omitempty"`
}

type InitParams struct {
	ProjectID string             `json:"project_id"`
	ProjectDB *DriverCredentials `json:"system_credentials"`
	CacheDB   *DriverCredentials `json:"cache_db"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

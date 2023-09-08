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

type PictureDeleteRequest struct {
	Urls      []string `json:"urls,omitempty" firestore:"urls,omitempty"`             // @gotags: firestore:"urls,omitempty"
	Model     string   `json:"model,omitempty" firestore:"model,omitempty"`           // @gotags: firestore:"model,omitempty"
	Id        string   `json:"id,omitempty" firestore:"id,omitempty"`                 // @gotags: firestore:"id,omitempty"
	FieldName string   `json:"field_name,omitempty" firestore:"field_name,omitempty"` // @gotags: firestore:"field_name,omitempty"
}

type FileDetails struct {
	Id            string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                                        // @gotags: firestore:"id,omitempty"
	XKey          string        `protobuf:"bytes,2,opt,name=_key,json=Key,proto3" json:"_key,omitempty" firestore:"_key,omitempty"`                                         // @gotags: firestore:"_key,omitempty"
	Type          string        `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" firestore:"type,omitempty"`                                                  // @gotags: firestore:"type,omitempty"
	FileExtension string        `protobuf:"bytes,4,opt,name=file_extension,json=fileExtension,proto3" json:"file_extension,omitempty" firestore:"file_extension,omitempty"` // @gotags: firestore:"file_extension,omitempty"
	FileName      string        `protobuf:"bytes,5,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty" firestore:"file_name,omitempty"`                     // @gotags: firestore:"file_name,omitempty"
	ContentType   string        `protobuf:"bytes,6,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty" firestore:"content_type,omitempty"`         // @gotags: firestore:"content_type,omitempty"
	Size          int64         `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty" firestore:"size,omitempty"`                                                 // @gotags: firestore:"size,omitempty"
	S3Key         string        `protobuf:"bytes,8,opt,name=s3_key,json=s3Key,proto3" json:"s3_key,omitempty" firestore:"s3_key,omitempty"`                                 // @gotags: firestore:"s3_key,omitempty"
	Url           string        `protobuf:"bytes,9,opt,name=url,proto3" json:"url,omitempty" firestore:"url,omitempty"`                                                     // @gotags: firestore:"url,omitempty"
	CreatedAt     string        `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" firestore:"created_at,omitempty"`                // @gotags: firestore:"created_at,omitempty"
	UploadParam   *UploadParams `protobuf:"bytes,11,opt,name=upload_param,json=uploadParam,proto3" json:"upload_param,omitempty" firestore:"upload_param,omitempty"`        // @gotags: firestore:"upload_param,omitempty"
	Buffer        []byte        `protobuf:"bytes,12,opt,name=buffer,proto3" json:"buffer,omitempty" firestore:"upload_param,omitempty"`                                     // @gotags: firestore:"upload_param,omitempty"
}

type UploadParams struct {
	DocId      string `protobuf:"bytes,1,opt,name=doc_id,json=docId,proto3" json:"doc_id,omitempty" firestore:"doc_id,omitempty"`                      // @gotags: firestore:"doc_id,omitempty"
	ProjectId  string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" firestore:"project_id,omitempty"`      // @gotags: firestore:"project_id,omitempty"
	ModelName  string `protobuf:"bytes,3,opt,name=model_name,json=modelName,proto3" json:"model_name,omitempty" firestore:"model_name,omitempty"`      // @gotags: firestore:"model_name,omitempty"
	FieldName  string `protobuf:"bytes,4,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty" firestore:"field_name,omitempty"`      // @gotags: firestore:"field_name,omitempty"
	AllowMulti bool   `protobuf:"varint,5,opt,name=allow_multi,json=allowMulti,proto3" json:"allow_multi,omitempty" firestore:"allow_multi,omitempty"` // @gotags: firestore:"allow_multi,omitempty"
	Provider   string `protobuf:"bytes,6,opt,name=provider,proto3" json:"provider,omitempty" firestore:"provider,omitempty"`                           // @gotags: firestore:"provider,omitempty"
}

type Filter struct {
	Page     uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty" firestore:"page,omitempty"`            // @gotags: firestore:"page,omitempty"
	Offset   uint32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty" firestore:"offset,omitempty"`      // @gotags: firestore:"offset,omitempty"
	Limit    uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty" firestore:"limit,omitempty"`         // @gotags: firestore:"limit,omitempty"
	Order    string `protobuf:"bytes,4,opt,name=order,proto3" json:"order,omitempty" firestore:"order,omitempty"`          // @gotags: firestore:"order,omitempty"
	Min      uint32 `protobuf:"varint,5,opt,name=min,proto3" json:"min,omitempty" firestore:"min,omitempty"`               // @gotags: firestore:"min,omitempty"
	Max      uint32 `protobuf:"varint,6,opt,name=max,proto3" json:"max,omitempty" firestore:"max,omitempty"`               // @gotags: firestore:"max,omitempty"
	Category string `protobuf:"bytes,7,opt,name=category,proto3" json:"category,omitempty" firestore:"category,omitempty"` // @gotags: firestore:"category,omitempty"
}

type Request struct {
	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                                    // @gotags: firestore:"id,omitempty"
	Type         string  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty" firestore:"type,omitempty"`                                              // @gotags: firestore:"type,omitempty"
	Filter       *Filter `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty" firestore:"filter,omitempty"`                                        // @gotags: firestore:"filter,omitempty"
	SearchString string  `protobuf:"bytes,5,opt,name=search_string,json=searchString,proto3" json:"search_string,omitempty" firestore:"search_string,omitempty"` // @gotags: firestore:"search_string,omitempty"
	Retry        bool    `protobuf:"varint,6,opt,name=retry,proto3" json:"retry,omitempty" firestore:"retry,omitempty"`                                          // @gotags: firestore:"retry,omitempty"
}

// File Picker
type FileLink struct {
	Link      string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty" firestore:"link,omitempty"`                                  // @gotags: firestore:"link,omitempty"
	Title     string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty" firestore:"title,omitempty"`                               // @gotags: firestore:"title,omitempty"
	CreatedAt string `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" firestore:"created_at,omitempty"` // @gotags: firestore:"created_at,omitempty"
	UpdatedAt string `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" firestore:"updated_at,omitempty"` // @gotags: firestore:"updated_at,omitempty"
}

type FilePickParameter struct {
	NumberOfImages uint32      `protobuf:"varint,1,opt,name=number_of_images,json=numberOfImages,proto3" json:"number_of_images,omitempty" firestore:"number_of_images,omitempty"` // @gotags: firestore:"number_of_images,omitempty"
	S3Folder       string      `protobuf:"bytes,2,opt,name=s3_folder,json=s3Folder,proto3" json:"s3_folder,omitempty" firestore:"s3_folder,omitempty"`                             // @gotags: firestore:"s3_folder,omitempty"
	PickerTitle    string      `protobuf:"bytes,3,opt,name=picker_title,json=pickerTitle,proto3" json:"picker_title,omitempty" firestore:"picker_title,omitempty"`                 // @gotags: firestore:"picker_title,omitempty"
	Origin         *SystemUser `protobuf:"bytes,4,opt,name=origin,proto3" json:"origin,omitempty" firestore:"origin,omitempty"`                                                    // @gotags: firestore:"origin,omitempty"
}

type ImageMetaInfo struct {
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty" firestore:"identifier,omitempty"` // @gotags: firestore:"identifier,omitempty"
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" firestore:"name,omitempty"`                   // @gotags: firestore:"name,omitempty"
	Width      uint32 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty" firestore:"width,omitempty"`               // @gotags: firestore:"width,omitempty"
	Height     uint32 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty" firestore:"height,omitempty"`            // @gotags: firestore:"height,omitempty"
	Type       string `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty" firestore:"type,omitempty"`                   // @gotags: firestore:"type,omitempty"
}

type LoginRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty" firestore:"username,omitempty"` // @gotags: firestore:"username,omitempty"
	Email    string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty" firestore:"email,omitempty"`          // @gotags: firestore:"email,omitempty"
	Secret   string `protobuf:"bytes,3,opt,name=secret,proto3" json:"secret,omitempty" firestore:"secret,omitempty"`       // @gotags: firestore:"secret,omitempty"
}

type Role struct {
	Id                        string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                                                                                                                                // @gotags: firestore:"id,omitempty"
	ApiPermissions            map[string]*APIPermission `protobuf:"bytes,2,rep,name=api_permissions,json=apiPermissions,proto3" json:"api_permissions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"permissions,omitempty"` // @gotags: firestore:"permissions,omitempty"
	AdministrativePermissions []string                  `protobuf:"bytes,3,rep,name=administrative_permissions,json=administrativePermissions,proto3" json:"administrative_permissions,omitempty" firestore:"administrative_permissions,omitempty"`                                         // @gotags: firestore:"administrative_permissions,omitempty"
	LogicExecutions           []string                  `protobuf:"bytes,4,rep,name=logic_executions,json=logicExecutions,proto3" json:"logic_executions,omitempty" firestore:"logic_executions,omitempty"`                                                                                 // @gotags: firestore:"logic_executions,omitempty"
	SystemGenerated           bool                      `protobuf:"varint,5,opt,name=system_generated,json=systemGenerated,proto3" json:"system_generated,omitempty" firestore:"system_generated,omitempty"`                                                                                // @gotags: firestore:"system_generated,omitempty"
	IsAdmin                   bool                      `protobuf:"varint,6,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty" firestore:"is_admin,omitempty"`                                                                                                                // @gotags: firestore:"is_admin,omitempty"
	IsProjectUser             bool                      `protobuf:"varint,7,opt,name=is_project_user,json=isProjectUser,proto3" json:"is_project_user,omitempty" firestore:"is_project_user,omitempty"`                                                                                     // @gotags: firestore:"is_project_user,omitempty"
	ReadOnlyProject           bool                      `protobuf:"varint,8,opt,name=read_only_project,json=readOnlyProject,proto3" json:"read_only_project,omitempty" firestore:"read_only_project,omitempty"`                                                                             // @gotags: firestore:"read_only_project,omitempty"
	IsSuperAdmin              bool                      `protobuf:"varint,9,opt,name=is_super_admin,json=isSuperAdmin,proto3" json:"is_super_admin,omitempty" firestore:"read_only_project,omitempty"`                                                                                      // @gotags: firestore:"read_only_project,omitempty"
}

type APIPermission struct {
	Read   string `protobuf:"bytes,2,opt,name=read,proto3" json:"read,omitempty" firestore:"read,omitempty"`       // @gotags: firestore:"read,omitempty"
	Create string `protobuf:"bytes,3,opt,name=create,proto3" json:"create,omitempty" firestore:"create,omitempty"` // @gotags: firestore:"create,omitempty"
	Update string `protobuf:"bytes,4,opt,name=update,proto3" json:"update,omitempty" firestore:"update,omitempty"` // @gotags: firestore:"update,omitempty"
	Delete string `protobuf:"bytes,5,opt,name=delete,proto3" json:"delete,omitempty" firestore:"delete,omitempty"` // @gotags: firestore:"delete,omitempty"
}

type SystemUser struct {
	XKey              string `protobuf:"bytes,1,opt,name=_key,json=Key,proto3" json:"_key,omitempty" firestore:"_key,omitempty"`                                                              // @gotags: firestore:"_key,omitempty"
	Id                string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                                                             // @gotags: firestore:"id,omitempty"
	Name              string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" firestore:"name,omitempty"`                                                                       // @gotags: firestore:"name,omitempty"
	Username          string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty" firestore:"username,omitempty"`                                                           // @gotags: firestore:"username,omitempty"
	Email             string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty" firestore:"email,omitempty"`                                                                    // @gotags: firestore:"email,omitempty"
	Secret            string `protobuf:"bytes,6,opt,name=secret,proto3" json:"secret,omitempty" firestore:"secret,omitempty"`                                                                 // @gotags: firestore:"secret,omitempty"
	Avatar            string `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty" firestore:"avatar,omitempty"`                                                                 // @gotags: firestore:"avatar,omitempty"
	CurrentProjectId  string `protobuf:"bytes,8,opt,name=current_project_id,json=currentProjectId,proto3" json:"current_project_id,omitempty" firestore:"current_project_id,omitempty"`       // @gotags: firestore:"current_project_id,omitempty"
	CreatedAt         string `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" firestore:"created_at,omitempty"`                                      // @gotags: firestore:"created_at,omitempty"
	UpdatedAt         string `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                                     // @gotags: firestore:"updated_at,omitempty"
	IsSuperAdmin      bool   `protobuf:"varint,11,opt,name=is_super_admin,json=isSuperAdmin,proto3" json:"is_super_admin,omitempty" firestore:"is_super_admin,omitempty"`                     // @gotags: firestore:"is_super_admin,omitempty"
	RefreshToken      string `protobuf:"bytes,12,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty" firestore:"refresh_token,omitempty"`                         // @gotags: firestore:"refresh_token,omitempty"
	AccessToken       string `protobuf:"bytes,13,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty" firestore:"access_token,omitempty"`                             // @gotags: firestore:"access_token,omitempty"
	ReadOnlyProject   bool   `protobuf:"varint,14,opt,name=read_only_project,json=readOnlyProject,proto3" json:"read_only_project,omitempty" firestore:"read_only_project,omitempty"`         // @gotags: firestore:"read_only_project,omitempty"
	LastLoggedIn      string `protobuf:"bytes,15,opt,name=last_logged_in,json=lastLoggedIn,proto3" json:"last_logged_in,omitempty" firestore:"last_logged_in,omitempty"`                      // @gotags: firestore:"last_logged_in,omitempty"
	ProjectLimit      uint32 `protobuf:"varint,16,opt,name=project_limit,json=projectLimit,proto3" json:"project_limit,omitempty" firestore:"project_limit,omitempty"`                        // @gotags: firestore:"project_limit,omitempty"
	ProjectLimitHobby uint32 `protobuf:"varint,17,opt,name=project_limit_hobby,json=projectLimitHobby,proto3" json:"project_limit_hobby,omitempty" firestore:"project_limit_hobby,omitempty"` // @gotags: firestore:"project_limit_hobby,omitempty"
	ProjectLimitPro   uint32 `protobuf:"varint,18,opt,name=project_limit_pro,json=projectLimitPro,proto3" json:"project_limit_pro,omitempty" firestore:"project_limit_pro,omitempty"`         // @gotags: firestore:"project_limit_pro,omitempty"
}

type ProjectLimit struct {
	Free  uint32 `protobuf:"varint,1,opt,name=free,proto3" json:"free,omitempty" firestore:"free,omitempty"`    // @gotags: firestore:"free,omitempty"
	Hobby uint32 `protobuf:"varint,2,opt,name=hobby,proto3" json:"hobby,omitempty" firestore:"hobby,omitempty"` // @gotags: firestore:"hobby,omitempty"
	Pro   uint32 `protobuf:"varint,3,opt,name=pro,proto3" json:"pro,omitempty" firestore:"pro,omitempty"`       // @gotags: firestore:"pro,omitempty"
}

type UserProjects struct {
	User     *SystemUser `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty" firestore:"user,omitempty"`             // @gotags: firestore:"user,omitempty"
	Projects []*Project  `protobuf:"bytes,2,rep,name=projects,proto3" json:"projects,omitempty" firestore:"projects,omitempty"` // @gotags: firestore:"projects,omitempty"
}

type UserSubscription struct {
	Subscription *MonthlySubscription `protobuf:"bytes,1,opt,name=subscription,proto3" json:"subscription,omitempty" firestore:"subscription,omitempty"` // @gotags: firestore:"subscription,omitempty"
	Invoices     []*ProjectInvoices   `protobuf:"bytes,2,rep,name=invoices,proto3" json:"invoices,omitempty" firestore:"invoices,omitempty"`             // @gotags: firestore:"invoices,omitempty"
}

type UserMeta struct {
	Id                        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                                                                   // @gotags: firestore:"id,omitempty"
	Avatar                    string   `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty" firestore:"avatar,omitempty"`                                                                       // @gotags: firestore:"avatar,omitempty"
	Name                      string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty" firestore:"name,omitempty"`                                                                             // @gotags: firestore:"name,omitempty"
	Role                      string   `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty" firestore:"role,omitempty"`                                                                             // @gotags: firestore:"role,omitempty"
	Email                     string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty" firestore:"email,omitempty"`                                                                          // @gotags: firestore:"email,omitempty"
	ProjectUser               bool     `protobuf:"varint,6,opt,name=project_user,json=projectUser,proto3" json:"project_user,omitempty" firestore:"email,omitempty"`                                          // @gotags: firestore:"email,omitempty"
	AdministrativePermissions []string `protobuf:"bytes,7,rep,name=administrative_permissions,json=administrativePermissions,proto3" json:"administrative_permissions,omitempty" firestore:"email,omitempty"` // @gotags: firestore:"email,omitempty"
}

type MetaField struct {
	CreatedAt      string    `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" firestore:"created_at,omitempty"`                        // @gotags: firestore:"created_at,omitempty"
	UpdatedAt      string    `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                        // @gotags: firestore:"updated_at,omitempty"
	CreatedBy      *UserMeta `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty" firestore:"title,omitempty"`                             // @gotags: firestore:"title,omitempty"
	LastModifiedBy *UserMeta `protobuf:"bytes,4,opt,name=last_modified_by,json=lastModifiedBy,proto3" json:"last_modified_by,omitempty" firestore:"created_by,omitempty"`       // @gotags: firestore:"created_by,omitempty"
	Status         string    `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty" firestore:"status,omitempty"`                                                   // @gotags: firestore:"status,omitempty"
	TenantId       string    `protobuf:"bytes,6,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty" firestore:"tenant_id,omitempty"`                            // @gotags: firestore:"tenant_id,omitempty"
	RootRevisionId string    `protobuf:"bytes,7,opt,name=root_revision_id,json=rootRevisionId,proto3" json:"root_revision_id,omitempty" firestore:"root_revision_id,omitempty"` // @gotags: firestore:"root_revision_id,omitempty"
	Revision       bool      `protobuf:"varint,9,opt,name=revision,proto3" json:"revision,omitempty" firestore:"revision,omitempty"`                                            // @gotags: firestore:"revision,omitempty"
	RevisionAt     string    `protobuf:"bytes,10,opt,name=revision_at,json=revisionAt,proto3" json:"revision_at,omitempty" firestore:"revision_at,omitempty"`                   // @gotags: firestore:"revision_at,omitempty"
}

type PreviewMode struct {
	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" firestore:"title,omitempty"`    // @gotags: firestore:"title,omitempty"
	Icon   string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon,omitempty" firestore:"icon,omitempty"`       // @gotags: firestore:"icon,omitempty"
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty" firestore:"status,omitempty"` // @gotags: firestore:"status,omitempty"
	Id     string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`             // @gotags: firestore:"id,omitempty"
}

type TeamMembers struct {

	// for sql migration purposes
	ProjectId         string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	UserId            string   `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AssignedRole      string   `protobuf:"bytes,3,opt,name=assigned_role,json=assignedRole,proto3" json:"assigned_role,omitempty"`
	AccessPermissions []string `protobuf:"bytes,4,rep,name=access_permissions,json=accessPermissions,proto3" json:"access_permissions,omitempty"`
}

type ProjectSchema struct {
	Models    []*ModelType     `protobuf:"bytes,1,rep,name=models,proto3" json:"models,omitempty" firestore:"models,omitempty"`          // @gotags: firestore:"models,omitempty"
	Functions []*CloudFunction `protobuf:"bytes,2,rep,name=functions,proto3" json:"functions,omitempty" firestore:"functions,omitempty"` // @gotags: firestore:"functions,omitempty"
}

// user project
type Project struct {
	XKey                string             `protobuf:"bytes,1,opt,name=_key,json=Key,proto3" json:"_key,omitempty" firestore:"_key,omitempty"`                                                                                    // @gotags: firestore:"_key,omitempty"
	Id                  string             `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey;autoIncrement:false" firestore:"id,omitempty"`                                                             // @gotags: gorm:"primaryKey;autoIncrement:false" firestore:"id,omitempty"
	ProjectName         string             `protobuf:"bytes,3,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty" firestore:"project_name,omitempty"`                                                    // @gotags: firestore:"project_name,omitempty"
	ProjectDescription  string             `protobuf:"bytes,4,opt,name=project_description,json=projectDescription,proto3" json:"project_description,omitempty" firestore:"project_description,omitempty"`                        // @gotags: firestore:"project_description,omitempty"
	Schema              *ProjectSchema     `protobuf:"bytes,5,opt,name=schema,proto3" json:"schema,omitempty" firestore:"schema,omitempty"`                                                                                       // @gotags: firestore:"schema,omitempty"
	CreatedAt           string             `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" firestore:"created_at,omitempty"`                                                            // @gotags: firestore:"created_at,omitempty"
	UpdatedAt           string             `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                                                            // @gotags: firestore:"updated_at,omitempty"
	ExpireAt            string             `protobuf:"bytes,8,opt,name=expire_at,json=expireAt,proto3" json:"expire_at,omitempty" firestore:"expire_at,omitempty"`                                                                // @gotags: firestore:"expire_at,omitempty"
	Plugins             []*PluginDetails   `protobuf:"bytes,9,rep,name=plugins,proto3" json:"plugins,omitempty" firestore:"plugins,omitempty"`                                                                                    // @gotags: firestore:"plugins,omitempty"
	Addons              *AddOnsDetails     `protobuf:"bytes,10,opt,name=addons,proto3" json:"addons,omitempty" gorm:"foreignKey:ProjectID" firestore:"addons,omitempty"`                                                          // @gotags: gorm:"foreignKey:ProjectID" firestore:"addons,omitempty"
	Tokens              []*APIToken        `protobuf:"bytes,11,rep,name=tokens,proto3" json:"tokens,omitempty" gorm:"foreignKey:ProjectID" firestore:"tokens,omitempty"`                                                          // @gotags: gorm:"foreignKey:ProjectID" firestore:"tokens,omitempty"
	Roles               map[string]*Role   `protobuf:"bytes,12,rep,name=roles,proto3" json:"roles,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" firestore:"roles,omitempty"` // @gotags: firestore:"roles,omitempty"
	Driver              *DriverCredentials `protobuf:"bytes,13,opt,name=driver,proto3" json:"driver,omitempty" gorm:"foreignKey:ProjectID" firestore:"driver,omitempty"`                                                          // @gotags: gorm:"foreignKey:ProjectID" firestore:"driver,omitempty"
	TempBanned          bool               `protobuf:"varint,14,opt,name=temp_banned,json=tempBanned,proto3" json:"temp_banned,omitempty" firestore:"temp_banned,omitempty"`                                                      // @gotags: firestore:"temp_banned,omitempty"
	Plan                string             `protobuf:"bytes,15,opt,name=plan,proto3" json:"plan,omitempty" firestore:"plan,omitempty"`                                                                                            // @gotags: firestore:"plan,omitempty"
	TrialEnds           string             `protobuf:"bytes,16,opt,name=trial_ends,json=trialEnds,proto3" json:"trial_ends,omitempty" firestore:"trial_ends,omitempty"`                                                           // @gotags: firestore:"trial_ends,omitempty"
	FromExample         string             `protobuf:"bytes,17,opt,name=from_example,json=fromExample,proto3" json:"from_example,omitempty" firestore:"from_example,omitempty"`                                                   // @gotags: firestore:"from_example,omitempty"
	Limits              *UsagesTracking    `protobuf:"bytes,18,opt,name=limits,proto3" json:"limits,omitempty" gorm:"foreignKey:ProjectID" firestore:"limits,omitempty"`                                                          // @gotags: gorm:"foreignKey:ProjectID" firestore:"limits,omitempty"
	OwnerId             string             `protobuf:"bytes,19,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty" firestore:"owner_id,omitempty"`                                                                   // @gotags: firestore:"owner_id,omitempty"
	Teams               []*TeamMembers     `protobuf:"bytes,20,rep,name=teams,proto3" json:"teams,omitempty" gorm:"foreignKey:ProjectID" firestore:"teams,omitempty"`                                                             // @gotags: gorm:"foreignKey:ProjectID" firestore:"teams,omitempty"
	SystemMessages      []*SystemMessage   `protobuf:"bytes,21,rep,name=system_messages,json=systemMessages,proto3" json:"system_messages,omitempty" firestore:"system_messages,omitempty"`                                       // @gotags: firestore:"system_messages,omitempty"
	Workspaces          []*Workspace       `protobuf:"bytes,22,rep,name=workspaces,proto3" json:"workspaces,omitempty" firestore:"workspaces,omitempty"`                                                                          // @gotags: firestore:"workspaces,omitempty"
	ActivatedPluginsIds *ActivatedPlugins  `protobuf:"bytes,23,opt,name=activated_plugins_ids,json=activatedPluginsIds,proto3" json:"activated_plugins_ids,omitempty" firestore:"activated_plugins_ids,omitempty"`                // @gotags: firestore:"activated_plugins_ids,omitempty"
}

type ActivatedPlugins struct {
	Storage  string `protobuf:"bytes,1,opt,name=storage,proto3" json:"storage,omitempty"`
	Function string `protobuf:"bytes,2,opt,name=function,proto3" json:"function,omitempty"`
}

type Workspace struct {
	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" firestore:"name,omitempty"`                                               // @gotags: firestore:"name,omitempty"
	Active       bool   `protobuf:"varint,2,opt,name=active,proto3" json:"active,omitempty" firestore:"active,omitempty"`                                        // @gotags: firestore:"active,omitempty"
	IsProduction bool   `protobuf:"varint,3,opt,name=is_production,json=isProduction,proto3" json:"is_production,omitempty" firestore:"is_production,omitempty"` // @gotags: firestore:"is_production,omitempty"
	IsDefault    bool   `protobuf:"varint,4,opt,name=is_default,json=isDefault,proto3" json:"is_default,omitempty" firestore:"is_default,omitempty"`             // @gotags: firestore:"is_default,omitempty"
}

type SystemMessage struct {
	Message     string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty" firestore:"message,omitempty"`             // @gotags: firestore:"message,omitempty"
	Code        string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty" firestore:"code,omitempty"`                      // @gotags: firestore:"code,omitempty"
	Redirection string `protobuf:"bytes,3,opt,name=redirection,proto3" json:"redirection,omitempty" firestore:"redirection,omitempty"` // @gotags: firestore:"redirection,omitempty"
	Hide        bool   `protobuf:"varint,4,opt,name=hide,proto3" json:"hide,omitempty" firestore:"hide,omitempty"`                     // @gotags: firestore:"hide,omitempty"
}

type SupportAndTicket struct {
	XKey             string         `protobuf:"bytes,1,opt,name=_key,json=Key,proto3" json:"_key,omitempty"`
	Id               string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type             string         `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	ProjectId        string         `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Resolved         bool           `protobuf:"varint,5,opt,name=resolved,proto3" json:"resolved,omitempty"`
	Title            string         `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	IssueDescription string         `protobuf:"bytes,7,opt,name=issue_description,json=issueDescription,proto3" json:"issue_description,omitempty"`
	CreatedAt        string         `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt        string         `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Replies          []*TicketReply `protobuf:"bytes,10,rep,name=replies,proto3" json:"replies,omitempty"`
}

type TicketReply struct {
	Description string    `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	User        *UserMeta `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt   string    `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Edited      bool      `protobuf:"varint,4,opt,name=edited,proto3" json:"edited,omitempty"`
}

type ProjectInvoices struct {
	XKey                  string  `protobuf:"bytes,1,opt,name=_key,json=Key,proto3" json:"_key,omitempty" firestore:"_key,omitempty"`                                                                                     // @gotags: firestore:"_key,omitempty"
	Id                    string  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                                                                                    // @gotags: firestore:"id,omitempty"
	Type                  string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" firestore:"type,omitempty"`                                                                                              // @gotags: firestore:"type,omitempty"
	SubscriptionId        string  `protobuf:"bytes,4,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty" firestore:"subscription_id,omitempty"`                                         // @gotags: firestore:"subscription_id,omitempty"
	CouponCode            string  `protobuf:"bytes,5,opt,name=coupon_code,json=couponCode,proto3" json:"coupon_code,omitempty" firestore:"coupon_code,omitempty"`                                                         // @gotags: firestore:"coupon_code,omitempty"
	DiscountAmount        float64 `protobuf:"fixed64,6,opt,name=discount_amount,json=discountAmount,proto3" json:"discount_amount,omitempty" firestore:"discount_amount,omitempty"`                                       // @gotags: firestore:"discount_amount,omitempty"
	PaymentGateway        string  `protobuf:"bytes,7,opt,name=payment_gateway,json=paymentGateway,proto3" json:"payment_gateway,omitempty" firestore:"payment_gateway,omitempty"`                                         // @gotags: firestore:"payment_gateway,omitempty"
	ReceiptUrl            string  `protobuf:"bytes,8,opt,name=receipt_url,json=receiptUrl,proto3" json:"receipt_url,omitempty" firestore:"receipt_url,omitempty"`                                                         // @gotags: firestore:"receipt_url,omitempty"
	ProductAmount         float64 `protobuf:"fixed64,9,opt,name=product_amount,json=productAmount,proto3" json:"product_amount,omitempty" firestore:"product_amount,omitempty"`                                           // @gotags: firestore:"product_amount,omitempty"
	DiscountPercentage    int64   `protobuf:"varint,10,opt,name=discount_percentage,json=discountPercentage,proto3" json:"discount_percentage,omitempty" firestore:"discount_percentage,omitempty"`                       // @gotags: firestore:"discount_percentage,omitempty"
	Quantity              uint32  `protobuf:"varint,11,opt,name=quantity,proto3" json:"quantity,omitempty" firestore:"quantity,omitempty"`                                                                                // @gotags: firestore:"quantity,omitempty"
	PaymentMethod         string  `protobuf:"bytes,12,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty" firestore:"payment_method,omitempty"`                                            // @gotags: firestore:"payment_method,omitempty"
	Earning               float64 `protobuf:"fixed64,13,opt,name=earning,proto3" json:"earning,omitempty" firestore:"earning,omitempty"`                                                                                  // @gotags: firestore:"earning,omitempty"
	Fee                   float64 `protobuf:"fixed64,14,opt,name=fee,proto3" json:"fee,omitempty" firestore:"fee,omitempty"`                                                                                              // @gotags: firestore:"fee,omitempty"
	Tax                   float64 `protobuf:"fixed64,15,opt,name=tax,proto3" json:"tax,omitempty" firestore:"tax,omitempty"`                                                                                              // @gotags: firestore:"tax,omitempty"
	CreatedAt             string  `protobuf:"bytes,16,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" firestore:"created_at,omitempty"`                                                            // @gotags: firestore:"created_at,omitempty"
	ChargedAmount         float64 `protobuf:"fixed64,17,opt,name=charged_amount,json=chargedAmount,proto3" json:"charged_amount,omitempty" firestore:"charged_amount,omitempty"`                                          // @gotags: firestore:"charged_amount,omitempty"
	NextAmountToBeCharged float64 `protobuf:"fixed64,18,opt,name=next_amount_to_be_charged,json=nextAmountToBeCharged,proto3" json:"next_amount_to_be_charged,omitempty" firestore:"next_amount_to_be_charged,omitempty"` // @gotags: firestore:"next_amount_to_be_charged,omitempty"
	SubscriptionPlanCode  string  `protobuf:"bytes,19,opt,name=subscription_plan_code,json=subscriptionPlanCode,proto3" json:"subscription_plan_code,omitempty" firestore:"subscription_plan_code,omitempty"`             // @gotags: firestore:"subscription_plan_code,omitempty"
}

type ProjectUsages struct {
	Id        string          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                        // @gotags: firestore:"id,omitempty"
	XKey      string          `protobuf:"bytes,2,opt,name=_key,json=Key,proto3" json:"_key,omitempty" firestore:"_key,omitempty"`                         // @gotags: firestore:"_key,omitempty"
	Type      string          `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" firestore:"type,omitempty"`                                  // @gotags: firestore:"type,omitempty"
	ProjectId string          `protobuf:"bytes,4,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	Year      uint32          `protobuf:"varint,5,opt,name=year,proto3" json:"year,omitempty" firestore:"year,omitempty"`                                 // @gotags: firestore:"year,omitempty"
	Month     uint32          `protobuf:"varint,6,opt,name=month,proto3" json:"month,omitempty" firestore:"month,omitempty"`                              // @gotags: firestore:"month,omitempty"
	Usages    *UsagesTracking `protobuf:"bytes,7,opt,name=usages,proto3" json:"usages,omitempty" firestore:"usage,omitempty"`                             // @gotags: firestore:"usage,omitempty"
}

type DriverCredentials struct {

	// for sql migration purposes
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	// normal sql & nosql database
	Engine   string `protobuf:"bytes,2,opt,name=engine,proto3" json:"engine,omitempty"`
	Host     string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port     string `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
	User     string `protobuf:"bytes,5,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
	Database string `protobuf:"bytes,7,opt,name=database,proto3" json:"database,omitempty"`
	// for firebase
	FirebaseProjectId             string `protobuf:"bytes,8,opt,name=firebase_project_id,json=firebaseProjectId,proto3" json:"firebase_project_id,omitempty"`
	FirebaseProjectCredentialJson string `protobuf:"bytes,9,opt,name=firebase_project_credential_json,json=firebaseProjectCredentialJson,proto3" json:"firebase_project_credential_json,omitempty"`
	// for dynamodb
	AccessKey string `protobuf:"bytes,10,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey string `protobuf:"bytes,11,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	// for embedded database
	DatabaseDir string `protobuf:"bytes,12,opt,name=database_dir,json=databaseDir,proto3" json:"database_dir,omitempty"`
}

type APIToken struct {

	// for sql migration purposes
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" firestore:"project_id,omitempty"` // @gotags: firestore:"project_id,omitempty"
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" firestore:"name,omitempty"`                                  // @gotags: firestore:"name,omitempty"
	Token     string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty" firestore:"token,omitempty"`                               // @gotags: firestore:"token,omitempty"
	Role      string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty" firestore:"role,omitempty"`                                  // @gotags: firestore:"role,omitempty"
	Expire    string `protobuf:"bytes,5,opt,name=expire,proto3" json:"expire,omitempty" firestore:"expire,omitempty"`                            // @gotags: firestore:"expire,omitempty"
}

type Response struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Code    string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
}

type PluginDetails struct {
	Icon             string                  `protobuf:"bytes,1,opt,name=icon,proto3" json:"icon,omitempty" firestore:"icon,omitempty"`                                                                                            // @gotags: firestore:"icon,omitempty"
	Id               string                  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                                                                                  // @gotags: firestore:"id,omitempty"
	Title            string                  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty" firestore:"title,omitempty"`                                                                                         // @gotags: firestore:"title,omitempty"
	Version          string                  `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty" firestore:"version,omitempty"`                                                                                   // @gotags: firestore:"version,omitempty"
	Description      string                  `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" firestore:"description,omitempty"`                                                                       // @gotags: firestore:"description,omitempty"
	Type             PluginType              `protobuf:"varint,6,opt,name=type,proto3,enum=protobuff.PluginType" json:"type,omitempty" firestore:"type,omitempty"`                                                                 // @gotags: firestore:"type,omitempty"
	Role             string                  `protobuf:"bytes,7,opt,name=role,proto3" json:"role,omitempty" firestore:"role,omitempty"`                                                                                            // @gotags: firestore:"role,omitempty"
	EnvVars          []*FunctionEnvVariables `protobuf:"bytes,10,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty" firestore:"env_vars,omitempty"`                                                                  // @gotags: firestore:"env_vars,omitempty"
	ExportedVariable string                  `protobuf:"bytes,8,opt,name=exported_variable,json=exportedVariable,proto3" json:"exported_variable,omitempty" firestore:"exported_variable,omitempty"`                               // @gotags: firestore:"exported_variable,omitempty"
	Enable           bool                    `protobuf:"varint,9,opt,name=enable,proto3" json:"enable,omitempty" firestore:"enable,omitempty"`                                                                                     // @gotags: firestore:"enable,omitempty"
	RepositoryUrl    string                  `protobuf:"bytes,11,opt,name=repository_url,json=repositoryUrl,proto3" json:"repository_url,omitempty" firestore:"repository_url,omitempty"`                                          // @gotags: firestore:"repository_url,omitempty"
	Branch           string                  `protobuf:"bytes,12,opt,name=branch,proto3" json:"branch,omitempty" firestore:"branch,omitempty"`                                                                                     // @gotags: firestore:"branch,omitempty"
	Author           string                  `protobuf:"bytes,13,opt,name=author,proto3" json:"author,omitempty" firestore:"author,omitempty"`                                                                                     // @gotags: firestore:"author,omitempty"
	LoadStatus       PluginLoadStatus        `protobuf:"varint,14,opt,name=load_status,json=loadStatus,proto3,enum=protobuff.PluginLoadStatus" json:"load_status,omitempty" firestore:"load_status,omitempty"`                     // @gotags: firestore:"load_status,omitempty"
	ActivateStatus   PluginActivateStatus    `protobuf:"varint,15,opt,name=activate_status,json=activateStatus,proto3,enum=protobuff.PluginActivateStatus" json:"activate_status,omitempty" firestore:"activate_status,omitempty"` // @gotags: firestore:"activate_status,omitempty"
}

type AddOnsDetails struct {

	// for sql migration purposes
	ProjectId             string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" firestore:"project_id,omitempty"`                                              // @gotags: firestore:"project_id,omitempty"
	Locals                []string `protobuf:"bytes,2,rep,name=locals,proto3" json:"locals,omitempty" firestore:"locals,omitempty"`                                                                         // @gotags: firestore:"locals,omitempty"
	SystemGraphqlHooks    bool     `protobuf:"varint,3,opt,name=system_graphql_hooks,json=systemGraphqlHooks,proto3" json:"system_graphql_hooks,omitempty" firestore:"system_graphql_hooks,omitempty"`      // @gotags: firestore:"system_graphql_hooks,omitempty"
	EnableRevisionHistory bool     `protobuf:"varint,4,opt,name=enable_revision_history,json=enableRevisionHistory,proto3" json:"enable_revision_history,omitempty" firestore:"revision_history,omitempty"` // @gotags: firestore:"revision_history,omitempty"
}

type AccountUsage struct {
	Id              string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                                                    // @gotags: firestore:"id,omitempty"
	XKey            string                 `protobuf:"bytes,2,opt,name=_key,json=Key,proto3" json:"_key,omitempty" firestore:"_key,omitempty"`                                                     // @gotags: firestore:"_key,omitempty"
	Type            string                 `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" firestore:"type,omitempty"`                                                              // @gotags: firestore:"type,omitempty"
	Subscriptions   []*MonthlySubscription `protobuf:"bytes,4,rep,name=subscriptions,proto3" json:"subscriptions,omitempty" firestore:"subscriptions,omitempty"`                                   // @gotags: firestore:"subscriptions,omitempty"
	Limits          *UsagesTracking        `protobuf:"bytes,5,opt,name=limits,proto3" json:"limits,omitempty" firestore:"limits,omitempty"`                                                        // @gotags: firestore:"limits,omitempty"
	Usages          []*UsagesTracking      `protobuf:"bytes,6,rep,name=usages,proto3" json:"usages,omitempty" firestore:"usages,omitempty"`                                                        // @gotags: firestore:"usages,omitempty"
	NumberOfProject uint32                 `protobuf:"varint,7,opt,name=number_of_project,json=numberOfProject,proto3" json:"number_of_project,omitempty" firestore:"number_of_project,omitempty"` // @gotags: firestore:"number_of_project,omitempty"
}

type MonthlySubscription struct {
	Id                 string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                                                            // @gotags: firestore:"id,omitempty"
	XKey               string             `protobuf:"bytes,2,opt,name=_key,json=Key,proto3" json:"_key,omitempty" firestore:"_key,omitempty"`                                                             // @gotags: firestore:"_key,omitempty"
	Type               string             `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" firestore:"type,omitempty"`                                                                      // @gotags: firestore:"type,omitempty"
	UserId             string             `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty" firestore:"user_id,omitempty"`                                                 // @gotags: firestore:"user_id,omitempty"
	SubscriptionCode   string             `protobuf:"bytes,5,opt,name=subscription_code,json=subscriptionCode,proto3" json:"subscription_code,omitempty" firestore:"plan_code,omitempty"`                 // @gotags: firestore:"plan_code,omitempty"
	StartDate          string             `protobuf:"bytes,6,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty" firestore:"start_date,omitempty"`                                     // @gotags: firestore:"start_date,omitempty"
	EndDate            string             `protobuf:"bytes,7,opt,name=end_date,json=endDate,proto3" json:"end_date,omitempty" firestore:"end_date,omitempty"`                                             // @gotags: firestore:"end_date,omitempty"
	SubscriptionStatus string             `protobuf:"bytes,8,opt,name=subscription_status,json=subscriptionStatus,proto3" json:"subscription_status,omitempty" firestore:"subscription_status,omitempty"` // @gotags: firestore:"subscription_status,omitempty"
	PaymentStatus      string             `protobuf:"bytes,9,opt,name=payment_status,json=paymentStatus,proto3" json:"payment_status,omitempty" firestore:"payment_status,omitempty"`                     // @gotags: firestore:"payment_status,omitempty"
	Invoices           []*ProjectInvoices `protobuf:"bytes,10,rep,name=invoices,proto3" json:"invoices,omitempty" firestore:"invoices,omitempty"`                                                         // @gotags: firestore:"invoices,omitempty"
	CreatedAt          string             `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" firestore:"created_at,omitempty"`                                    // @gotags: firestore:"created_at,omitempty"
	UpdatedAt          string             `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                                    // @gotags: firestore:"updated_at,omitempty"
	Title              string             `protobuf:"bytes,13,opt,name=title,proto3" json:"title,omitempty" firestore:"title,omitempty"`                                                                  // @gotags: firestore:"title,omitempty"
	CancelUrl          string             `protobuf:"bytes,14,opt,name=cancel_url,json=cancelUrl,proto3" json:"cancel_url,omitempty" firestore:"cancel_url,omitempty"`                                    // @gotags: firestore:"cancel_url,omitempty"
	ProjectId          string             `protobuf:"bytes,15,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" firestore:"cancel_url,omitempty"`                                    // @gotags: firestore:"cancel_url,omitempty"
}

type UsagesTracking struct {

	// for sql migration purposes
	ProjectId       string  `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" firestore:"project_id,omitempty"`                              // @gotags: firestore:"project_id,omitempty"
	ApiCalls        uint32  `protobuf:"varint,2,opt,name=api_calls,json=apiCalls,proto3" json:"api_calls,omitempty" firestore:"api_calls,omitempty"`                                 // @gotags: firestore:"api_calls,omitempty"
	ApiBandwidth    float64 `protobuf:"fixed64,3,opt,name=api_bandwidth,json=apiBandwidth,proto3" json:"api_bandwidth,omitempty" firestore:"api_bandwidth,omitempty"`                // @gotags: firestore:"api_bandwidth,omitempty"
	MediaStorage    float64 `protobuf:"fixed64,4,opt,name=media_storage,json=mediaStorage,proto3" json:"media_storage,omitempty" firestore:"media_storage,omitempty"`                // @gotags: firestore:"media_storage,omitempty"
	MediaBandwidth  float64 `protobuf:"fixed64,5,opt,name=media_bandwidth,json=mediaBandwidth,proto3" json:"media_bandwidth,omitempty" firestore:"media_bandwidth,omitempty"`        // @gotags: firestore:"media_bandwidth,omitempty"
	NumberOfMedia   float64 `protobuf:"fixed64,6,opt,name=number_of_media,json=numberOfMedia,proto3" json:"number_of_media,omitempty" firestore:"number_of_media,omitempty"`         // @gotags: firestore:"number_of_media,omitempty"
	NumberOfRecords float64 `protobuf:"fixed64,7,opt,name=number_of_records,json=numberOfRecords,proto3" json:"number_of_records,omitempty" firestore:"number_of_records,omitempty"` // @gotags: firestore:"number_of_records,omitempty"
}

type ModelType struct {
	Name            string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" firestore:"name,omitempty"`                                                           // @gotags: firestore:"name,omitempty"
	Fields          []*FieldInfo      `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" firestore:"fields,omitempty"`                                                     // @gotags: firestore:"fields,omitempty"
	Connections     []*ConnectionType `protobuf:"bytes,3,rep,name=connections,proto3" json:"connections,omitempty" firestore:"connections,omitempty"`                                      // @gotags: firestore:"connections,omitempty"
	HookIds         []string          `protobuf:"bytes,4,rep,name=hook_ids,json=hookIds,proto3" json:"hook_ids,omitempty" firestore:"hook_ids,omitempty"`                                  // @gotags: firestore:"hook_ids,omitempty"
	Locals          []string          `protobuf:"bytes,5,rep,name=locals,proto3" json:"locals,omitempty" firestore:"locals,omitempty"`                                                     // @gotags: firestore:"locals,omitempty"
	RepeatedGroups  []string          `protobuf:"bytes,6,rep,name=repeated_groups,json=repeatedGroups,proto3" json:"repeated_groups,omitempty" firestore:"locals,omitempty"`               // @gotags: firestore:"locals,omitempty"
	SystemGenerated bool              `protobuf:"varint,7,opt,name=system_generated,json=systemGenerated,proto3" json:"system_generated,omitempty" firestore:"system_generated,omitempty"` // @gotags: firestore:"system_generated,omitempty"
	SinglePage      bool              `protobuf:"varint,8,opt,name=single_page,json=singlePage,proto3" json:"single_page,omitempty" firestore:"system_generated,omitempty"`                // @gotags: firestore:"system_generated,omitempty"
	SinglePageUuid  string            `protobuf:"bytes,9,opt,name=single_page_uuid,json=singlePageUuid,proto3" json:"single_page_uuid,omitempty" firestore:"system_generated,omitempty"`   // @gotags: firestore:"system_generated,omitempty"
	HasConnections  bool              `protobuf:"varint,10,opt,name=has_connections,json=hasConnections,proto3" json:"has_connections,omitempty" firestore:"has_connections,omitempty"`    // @gotags: firestore:"has_connections,omitempty"
}

type CloudFunction struct {
	Id                       string                            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                                                                                        // @gotags: firestore:"id,omitempty"
	Name                     string                            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" firestore:"name,omitempty"`                                                                                                  // @gotags: firestore:"name,omitempty"
	Description              string                            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty" firestore:"description,omitempty"`                                                                             // @gotags: firestore:"description,omitempty"
	EnvVars                  []*FunctionEnvVariables           `protobuf:"bytes,4,rep,name=env_vars,json=envVars,proto3" json:"env_vars,omitempty" firestore:"env_vars,omitempty"`                                                                         // @gotags: firestore:"env_vars,omitempty"
	FunctionPath             string                            `protobuf:"bytes,5,opt,name=function_path,json=functionPath,proto3" json:"function_path,omitempty" firestore:"function_path,omitempty"`                                                     // @gotags: firestore:"function_path,omitempty"
	Request                  *CloudFunctionRequestResponseType `protobuf:"bytes,6,opt,name=request,proto3" json:"request,omitempty" firestore:"request,omitempty"`                                                                                         // @gotags: firestore:"request,omitempty"
	Response                 *CloudFunctionRequestResponseType `protobuf:"bytes,7,opt,name=response,proto3" json:"response,omitempty" firestore:"response,omitempty"`                                                                                      // @gotags: firestore:"response,omitempty"
	FunctionProviderId       string                            `protobuf:"bytes,8,opt,name=function_provider_id,json=functionProviderId,proto3" json:"function_provider_id,omitempty" firestore:"function_provider_id,omitempty"`                          // @gotags: firestore:"function_provider_id,omitempty"
	FunctionConnected        bool                              `protobuf:"varint,9,opt,name=function_connected,json=functionConnected,proto3" json:"function_connected,omitempty" firestore:"function_connected,omitempty"`                                // @gotags: firestore:"function_connected,omitempty"
	ProviderExportedVariable string                            `protobuf:"bytes,10,opt,name=provider_exported_variable,json=providerExportedVariable,proto3" json:"provider_exported_variable,omitempty" firestore:"provider_exported_variable,omitempty"` // @gotags: firestore:"provider_exported_variable,omitempty"
	FunctionExportedVariable string                            `protobuf:"bytes,11,opt,name=function_exported_variable,json=functionExportedVariable,proto3" json:"function_exported_variable,omitempty" firestore:"function_exported_variable,omitempty"` // @gotags: firestore:"function_exported_variable,omitempty"
	RuntimeConfig            *FunctionRuntimeConfig            `protobuf:"bytes,12,opt,name=runtime_config,json=runtimeConfig,proto3" json:"runtime_config,omitempty" firestore:"runtime_config,omitempty"`                                                // @gotags: firestore:"runtime_config,omitempty"
	UpdatedAt                string                            `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" firestore:"updated_at,omitempty"`                                                                // @gotags: firestore:"updated_at,omitempty"
	CreatedAt                string                            `protobuf:"bytes,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" firestore:"created_at,omitempty"`                                                                // @gotags: firestore:"created_at,omitempty"
}

type FunctionEnvVariables struct {
	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty" firestore:"key,omitempty"`       // @gotags: firestore:"key,omitempty"
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty" firestore:"value,omitempty"` // @gotags: firestore:"value,omitempty"
}

type CloudFunctionRequestResponseType struct {
	Model  string       `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty" firestore:"model,omitempty"`    // @gotags: firestore:"model,omitempty" // fixed model name Ex: model, JSON, CUSTOM
	Params []*FieldInfo `protobuf:"bytes,3,rep,name=params,proto3" json:"params,omitempty" firestore:"params,omitempty"` // @gotags: firestore:"params,omitempty" // for only fixed CUSTOM request type
}

type FunctionRuntimeConfig struct {
	Runtime string `protobuf:"bytes,1,opt,name=runtime,proto3" json:"runtime,omitempty" firestore:"runtime,omitempty"`                  // @gotags: firestore:"runtime,omitempty"
	Memory  int64  `protobuf:"varint,2,opt,name=memory,proto3" json:"memory,omitempty" firestore:"memory,omitempty"`                    // @gotags: firestore:"memory,omitempty"
	Handler string `protobuf:"bytes,3,opt,name=handler,proto3" json:"handler,omitempty" firestore:"handler,omitempty"`                  // @gotags: firestore:"handler,omitempty"
	TimeOut int64  `protobuf:"varint,4,opt,name=time_out,json=timeOut,proto3" json:"time_out,omitempty" firestore:"time_out,omitempty"` // @gotags: firestore:"time_out,omitempty"
}

type Validation struct {
	Required          bool      `protobuf:"varint,1,opt,name=required,proto3" json:"required,omitempty" firestore:"required,omitempty"`                                                              // @gotags: firestore:"required,omitempty"
	Hide              bool      `protobuf:"varint,2,opt,name=hide,proto3" json:"hide,omitempty" firestore:"hide,omitempty"`                                                                          // @gotags: firestore:"hide,omitempty"
	AsTitle           bool      `protobuf:"varint,3,opt,name=as_title,json=asTitle,proto3" json:"as_title,omitempty" firestore:"as_title,omitempty"`                                                 // @gotags: firestore:"as_title,omitempty"
	Locals            []string  `protobuf:"bytes,4,rep,name=locals,proto3" json:"locals,omitempty" firestore:"locals,omitempty"`                                                                     // @gotags: firestore:"locals,omitempty"
	Unique            bool      `protobuf:"varint,5,opt,name=unique,proto3" json:"unique,omitempty" firestore:"unique,omitempty"`                                                                    // @gotags: firestore:"unique,omitempty"
	CharLimit         []uint32  `protobuf:"varint,6,rep,packed,name=char_limit,json=charLimit,proto3" json:"char_limit,omitempty" firestore:"char_limit,omitempty"`                                  // @gotags: firestore:"char_limit,omitempty"
	IntRangeLimit     []uint32  `protobuf:"varint,10,rep,packed,name=int_range_limit,json=intRangeLimit,proto3" json:"int_range_limit,omitempty" firestore:"int_range_limit,omitempty"`              // @gotags: firestore:"int_range_limit,omitempty"
	DoubleRangeLimit  []float64 `protobuf:"fixed64,11,rep,packed,name=double_range_limit,json=doubleRangeLimit,proto3" json:"double_range_limit,omitempty" firestore:"double_range_limit,omitempty"` // @gotags: firestore:"double_range_limit,omitempty"
	IsEmail           bool      `protobuf:"varint,7,opt,name=is_email,json=isEmail,proto3" json:"is_email,omitempty" firestore:"is_email,omitempty"`                                                 // @gotags: firestore:"is_email,omitempty"
	Placeholder       string    `protobuf:"bytes,8,opt,name=placeholder,proto3" json:"placeholder,omitempty" firestore:"placeholder,omitempty"`                                                      // @gotags: firestore:"placeholder,omitempty"
	FixedListElements []string  `protobuf:"bytes,9,rep,name=fixed_list_elements,json=fixedListElements,proto3" json:"fixed_list_elements,omitempty" firestore:"fixed_list_elements,omitempty"`       // @gotags: firestore:"fixed_list_elements,omitempty"
	IsMultiChoice     bool      `protobuf:"varint,12,opt,name=is_multi_choice,json=isMultiChoice,proto3" json:"is_multi_choice,omitempty" firestore:"is_multi_choice,omitempty"`                     // @gotags: firestore:"is_multi_choice,omitempty"
	ListType          string    `protobuf:"bytes,13,opt,name=list_type,json=listType,proto3" json:"list_type,omitempty" firestore:"list_type,omitempty"`                                             // @gotags: firestore:"list_type,omitempty"
	IsGallery         bool      `protobuf:"varint,14,opt,name=is_gallery,json=isGallery,proto3" json:"is_gallery,omitempty" firestore:"is_gallery,omitempty"`                                        // @gotags: firestore:"is_gallery,omitempty"
	IsPassword        bool      `protobuf:"varint,15,opt,name=is_password,json=isPassword,proto3" json:"is_password,omitempty" firestore:"is_gallery,omitempty"`                                     // @gotags: firestore:"is_gallery,omitempty"
	IsSystemRole      bool      `protobuf:"varint,16,opt,name=is_system_role,json=isSystemRole,proto3" json:"is_system_role,omitempty" firestore:"is_gallery,omitempty"`                             // @gotags: firestore:"is_gallery,omitempty"
	IsUrl             bool      `protobuf:"varint,17,opt,name=is_url,json=isUrl,proto3" json:"is_url,omitempty" firestore:"is_url,omitempty"`                                                        // @gotags: firestore:"is_url,omitempty"
}

type FieldInfo struct {
	Identifier              string       `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty" firestore:"identifier,omitempty"`                                                                            // @gotags: firestore:"identifier,omitempty"
	Description             string       `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" firestore:"description,omitempty"`                                                                         // @gotags: firestore:"description,omitempty"
	InputType               string       `protobuf:"bytes,3,opt,name=input_type,json=inputType,proto3" json:"input_type,omitempty" firestore:"input_type,omitempty"`                                                             // @gotags: firestore:"input_type,omitempty" // string
	FieldType               string       `protobuf:"bytes,4,opt,name=field_type,json=fieldType,proto3" json:"field_type,omitempty" firestore:"field_type,omitempty"`                                                             // @gotags: firestore:"field_type,omitempty" // list
	SubFieldInfo            []*FieldInfo `protobuf:"bytes,5,rep,name=sub_field_info,json=subFieldInfo,proto3" json:"sub_field_info,omitempty" firestore:"modules,omitempty"`                                                     // @gotags: firestore:"modules,omitempty" // to avoid a loop
	Validation              *Validation  `protobuf:"bytes,6,opt,name=validation,proto3" json:"validation,omitempty" firestore:"validation,omitempty"`                                                                            // @gotags: firestore:"validation,omitempty"
	Serial                  uint32       `protobuf:"varint,7,opt,name=serial,proto3" json:"serial,omitempty" firestore:"serial,omitempty"`                                                                                       // @gotags: firestore:"serial,omitempty"
	Label                   string       `protobuf:"bytes,8,opt,name=label,proto3" json:"label,omitempty" firestore:"label,omitempty"`                                                                                           // @gotags: firestore:"label,omitempty"
	SystemGenerated         bool         `protobuf:"varint,9,opt,name=system_generated,json=systemGenerated,proto3" json:"system_generated,omitempty" firestore:"system_generated,omitempty"`                                    // @gotags: firestore:"system_generated,omitempty"
	RepeatedGroupIdentifier string       `protobuf:"bytes,10,opt,name=repeated_group_identifier,json=repeatedGroupIdentifier,proto3" json:"repeated_group_identifier,omitempty" firestore:"repeated_group_identifier,omitempty"` // @gotags: firestore:"repeated_group_identifier,omitempty"
	IsObjectField           bool         `protobuf:"varint,11,opt,name=is_object_field,json=isObjectField,proto3" json:"is_object_field,omitempty" firestore:"is_object_field,omitempty"`                                        // @gotags: firestore:"is_object_field,omitempty"
	ParentField             string       `protobuf:"bytes,12,opt,name=parent_field,json=parentField,proto3" json:"parent_field,omitempty" firestore:"parent_field,omitempty"`                                                    // @gotags: firestore:"parent_field,omitempty"
}

type ConnectionType struct {
	Model    string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty" firestore:"model,omitempty"`                       // @gotags: firestore:"model,omitempty"
	Relation string `protobuf:"bytes,2,opt,name=relation,proto3" json:"relation,omitempty" firestore:"relation,omitempty"`              // @gotags: firestore:"relation,omitempty"
	Type     string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" firestore:"type,omitempty"`                          // @gotags: firestore:"type,omitempty"
	KnownAs  string `protobuf:"bytes,4,opt,name=known_as,json=knownAs,proto3" json:"known_as,omitempty" firestore:"known_as,omitempty"` // @gotags: firestore:"known_as,omitempty"
}

type Webhook struct {
	Id              string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" firestore:"id,omitempty"`                                                                // @gotags: firestore:"id,omitempty"
	XKey            string   `protobuf:"bytes,2,opt,name=_key,json=Key,proto3" json:"_key,omitempty" firestore:"_key,omitempty"`                                                 // @gotags: firestore:"_key,omitempty"
	Type            string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty" firestore:"type,omitempty"`                                                          // @gotags: firestore:"type,omitempty"
	Model           string   `protobuf:"bytes,4,opt,name=model,proto3" json:"model,omitempty" firestore:"model,omitempty"`                                                       // @gotags: firestore:"model,omitempty"
	ProjectId       string   `protobuf:"bytes,5,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty" firestore:"project_id,omitempty"`                         // @gotags: firestore:"project_id,omitempty"
	Name            string   `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty" firestore:"name,omitempty"`                                                          // @gotags: firestore:"name,omitempty"
	Events          []string `protobuf:"bytes,7,rep,name=events,proto3" json:"events,omitempty" firestore:"events,omitempty"`                                                    // @gotags: firestore:"events,omitempty"
	Url             string   `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty" firestore:"url,omitempty"`                                                             // @gotags: firestore:"url,omitempty"
	LogicExecutions []string `protobuf:"bytes,9,rep,name=logic_executions,json=logicExecutions,proto3" json:"logic_executions,omitempty" firestore:"logic_executions,omitempty"` // @gotags: firestore:"logic_executions,omitempty"
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: schema.proto

package micro

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenCodesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Rebuild   bool   `protobuf:"varint,2,opt,name=rebuild,proto3" json:"rebuild,omitempty"`
}

func (x *GenCodesRequest) Reset() {
	*x = GenCodesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenCodesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenCodesRequest) ProtoMessage() {}

func (x *GenCodesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenCodesRequest.ProtoReflect.Descriptor instead.
func (*GenCodesRequest) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{0}
}

func (x *GenCodesRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GenCodesRequest) GetRebuild() bool {
	if x != nil {
		return x.Rebuild
	}
	return false
}

type CodeGenSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId                string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	HasGeneratedModel        bool   `protobuf:"varint,2,opt,name=has_generated_model,json=hasGeneratedModel,proto3" json:"has_generated_model,omitempty"`
	HasGeneratedCodes        bool   `protobuf:"varint,3,opt,name=has_generated_codes,json=hasGeneratedCodes,proto3" json:"has_generated_codes,omitempty"`
	HasGeneratedResolverRoot bool   `protobuf:"varint,4,opt,name=has_generated_resolver_root,json=hasGeneratedResolverRoot,proto3" json:"has_generated_resolver_root,omitempty"`
	HasGeneratedResolver     bool   `protobuf:"varint,5,opt,name=has_generated_resolver,json=hasGeneratedResolver,proto3" json:"has_generated_resolver,omitempty"`
}

func (x *CodeGenSuccess) Reset() {
	*x = CodeGenSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodeGenSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodeGenSuccess) ProtoMessage() {}

func (x *CodeGenSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodeGenSuccess.ProtoReflect.Descriptor instead.
func (*CodeGenSuccess) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{1}
}

func (x *CodeGenSuccess) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *CodeGenSuccess) GetHasGeneratedModel() bool {
	if x != nil {
		return x.HasGeneratedModel
	}
	return false
}

func (x *CodeGenSuccess) GetHasGeneratedCodes() bool {
	if x != nil {
		return x.HasGeneratedCodes
	}
	return false
}

func (x *CodeGenSuccess) GetHasGeneratedResolverRoot() bool {
	if x != nil {
		return x.HasGeneratedResolverRoot
	}
	return false
}

func (x *CodeGenSuccess) GetHasGeneratedResolver() bool {
	if x != nil {
		return x.HasGeneratedResolver
	}
	return false
}

type GenSchemaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId   string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Request     []byte `protobuf:"bytes,2,opt,name=request,proto3" json:"request,omitempty"`
	Cache       []byte `protobuf:"bytes,3,opt,name=cache,proto3" json:"cache,omitempty"`
	UpdateCache bool   `protobuf:"varint,4,opt,name=updateCache,proto3" json:"updateCache,omitempty"`
}

func (x *GenSchemaRequest) Reset() {
	*x = GenSchemaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenSchemaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenSchemaRequest) ProtoMessage() {}

func (x *GenSchemaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenSchemaRequest.ProtoReflect.Descriptor instead.
func (*GenSchemaRequest) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{2}
}

func (x *GenSchemaRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GenSchemaRequest) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *GenSchemaRequest) GetCache() []byte {
	if x != nil {
		return x.Cache
	}
	return nil
}

func (x *GenSchemaRequest) GetUpdateCache() bool {
	if x != nil {
		return x.UpdateCache
	}
	return false
}

type GenSchemaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Result    []byte `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *GenSchemaResponse) Reset() {
	*x = GenSchemaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenSchemaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenSchemaResponse) ProtoMessage() {}

func (x *GenSchemaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenSchemaResponse.ProtoReflect.Descriptor instead.
func (*GenSchemaResponse) Descriptor() ([]byte, []int) {
	return file_schema_proto_rawDescGZIP(), []int{3}
}

func (x *GenSchemaResponse) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GenSchemaResponse) GetResult() []byte {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_schema_proto protoreflect.FileDescriptor

var file_schema_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x22, 0x4a, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x22, 0x84, 0x02, 0x0a, 0x0e, 0x43, 0x6f, 0x64, 0x65, 0x47, 0x65, 0x6e, 0x53,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13, 0x68, 0x61, 0x73, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x68, 0x61, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x2e, 0x0a, 0x13, 0x68, 0x61, 0x73, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x11, 0x68, 0x61, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x1b, 0x68, 0x61, 0x73, 0x5f, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x5f,
	0x72, 0x6f, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x18, 0x68, 0x61, 0x73, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72,
	0x52, 0x6f, 0x6f, 0x74, 0x12, 0x34, 0x0a, 0x16, 0x68, 0x61, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x68, 0x61, 0x73, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x72, 0x22, 0x83, 0x01, 0x0a, 0x10, 0x47,
	0x65, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x63, 0x68,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x61, 0x63, 0x68, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x63, 0x68, 0x65,
	0x22, 0x4a, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x9e, 0x01, 0x0a,
	0x0c, 0x4d, 0x69, 0x63, 0x72, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a,
	0x15, 0x54, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x47, 0x65, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x47,
	0x65, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x42, 0x0a, 0x09, 0x45, 0x78, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x19, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e,
	0x2e, 0x47, 0x65, 0x6e, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x53,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x23, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x74,
	0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_proto_rawDescOnce sync.Once
	file_schema_proto_rawDescData = file_schema_proto_rawDesc
)

func file_schema_proto_rawDescGZIP() []byte {
	file_schema_proto_rawDescOnce.Do(func() {
		file_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_proto_rawDescData)
	})
	return file_schema_proto_rawDescData
}

var file_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_schema_proto_goTypes = []interface{}{
	(*GenCodesRequest)(nil),   // 0: codegen.GenCodesRequest
	(*CodeGenSuccess)(nil),    // 1: codegen.CodeGenSuccess
	(*GenSchemaRequest)(nil),  // 2: codegen.GenSchemaRequest
	(*GenSchemaResponse)(nil), // 3: codegen.GenSchemaResponse
}
var file_schema_proto_depIdxs = []int32{
	0, // 0: codegen.MicroService.TriggerCodeGeneration:input_type -> codegen.GenCodesRequest
	2, // 1: codegen.MicroService.ExeSchema:input_type -> codegen.GenSchemaRequest
	1, // 2: codegen.MicroService.TriggerCodeGeneration:output_type -> codegen.CodeGenSuccess
	3, // 3: codegen.MicroService.ExeSchema:output_type -> codegen.GenSchemaResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_schema_proto_init() }
func file_schema_proto_init() {
	if File_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenCodesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodeGenSuccess); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenSchemaRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenSchemaResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_schema_proto_goTypes,
		DependencyIndexes: file_schema_proto_depIdxs,
		MessageInfos:      file_schema_proto_msgTypes,
	}.Build()
	File_schema_proto = out.File
	file_schema_proto_rawDesc = nil
	file_schema_proto_goTypes = nil
	file_schema_proto_depIdxs = nil
}

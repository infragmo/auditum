// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: infragmo/auditum/v1alpha1/project.proto

package auditumv1alpha1

import (
	_ "github.com/infragmo/auditum/api/gen/go/google/api"
	_ "github.com/infragmo/auditum/api/gen/go/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Represents a project.
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Internal project identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Time when the project was created.
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Display name of the project.
	//
	// REQUIREMENTS.
	// The value must be 3-64 characters long.
	DisplayName string `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Whether to allow updating records for this project.
	// If set, overrides the global setting.
	// Defaults to unset.
	UpdateRecordEnabled *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=update_record_enabled,json=updateRecordEnabled,proto3" json:"update_record_enabled,omitempty"`
	// Whether to allow deleting records for this project.
	// If set, overrides the global setting.
	// Defaults to unset.
	DeleteRecordEnabled *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=delete_record_enabled,json=deleteRecordEnabled,proto3" json:"delete_record_enabled,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_infragmo_auditum_v1alpha1_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_infragmo_auditum_v1alpha1_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_infragmo_auditum_v1alpha1_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *Project) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Project) GetUpdateRecordEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.UpdateRecordEnabled
	}
	return nil
}

func (x *Project) GetDeleteRecordEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.DeleteRecordEnabled
	}
	return nil
}

var File_infragmo_auditum_v1alpha1_project_proto protoreflect.FileDescriptor

var file_infragmo_auditum_v1alpha1_project_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x75, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x6e, 0x66, 0x72, 0x61,
	0x67, 0x6d, 0x6f, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x12, 0x26, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16,
	0x92, 0x41, 0x10, 0xca, 0x3e, 0x0d, 0xfa, 0x02, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5f, 0x69, 0x64, 0xe0, 0x41, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x15, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x03, 0xe0, 0x41, 0x01, 0x52, 0x13, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x53, 0x0a, 0x15, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x03, 0xe0, 0x41, 0x01, 0x52, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x42, 0x85,
	0x02, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x6f, 0x2e,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x0c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66,
	0x72, 0x61, 0x67, 0x6d, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x75, 0x6d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x67, 0x6d,
	0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x75, 0x6d, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0xa2, 0x02, 0x03, 0x49, 0x41, 0x58, 0xaa, 0x02, 0x19, 0x49, 0x6e, 0x66, 0x72, 0x61,
	0x67, 0x6d, 0x6f, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x75, 0x6d, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x19, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x6f, 0x5c,
	0x41, 0x75, 0x64, 0x69, 0x74, 0x75, 0x6d, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x25, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x6f, 0x5c, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x75, 0x6d, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x49, 0x6e, 0x66, 0x72, 0x61,
	0x67, 0x6d, 0x6f, 0x3a, 0x3a, 0x41, 0x75, 0x64, 0x69, 0x74, 0x75, 0x6d, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_infragmo_auditum_v1alpha1_project_proto_rawDescOnce sync.Once
	file_infragmo_auditum_v1alpha1_project_proto_rawDescData = file_infragmo_auditum_v1alpha1_project_proto_rawDesc
)

func file_infragmo_auditum_v1alpha1_project_proto_rawDescGZIP() []byte {
	file_infragmo_auditum_v1alpha1_project_proto_rawDescOnce.Do(func() {
		file_infragmo_auditum_v1alpha1_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_infragmo_auditum_v1alpha1_project_proto_rawDescData)
	})
	return file_infragmo_auditum_v1alpha1_project_proto_rawDescData
}

var file_infragmo_auditum_v1alpha1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_infragmo_auditum_v1alpha1_project_proto_goTypes = []interface{}{
	(*Project)(nil),               // 0: infragmo.auditum.v1alpha1.Project
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*wrapperspb.BoolValue)(nil),  // 2: google.protobuf.BoolValue
}
var file_infragmo_auditum_v1alpha1_project_proto_depIdxs = []int32{
	1, // 0: infragmo.auditum.v1alpha1.Project.create_time:type_name -> google.protobuf.Timestamp
	2, // 1: infragmo.auditum.v1alpha1.Project.update_record_enabled:type_name -> google.protobuf.BoolValue
	2, // 2: infragmo.auditum.v1alpha1.Project.delete_record_enabled:type_name -> google.protobuf.BoolValue
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_infragmo_auditum_v1alpha1_project_proto_init() }
func file_infragmo_auditum_v1alpha1_project_proto_init() {
	if File_infragmo_auditum_v1alpha1_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_infragmo_auditum_v1alpha1_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
			RawDescriptor: file_infragmo_auditum_v1alpha1_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_infragmo_auditum_v1alpha1_project_proto_goTypes,
		DependencyIndexes: file_infragmo_auditum_v1alpha1_project_proto_depIdxs,
		MessageInfos:      file_infragmo_auditum_v1alpha1_project_proto_msgTypes,
	}.Build()
	File_infragmo_auditum_v1alpha1_project_proto = out.File
	file_infragmo_auditum_v1alpha1_project_proto_rawDesc = nil
	file_infragmo_auditum_v1alpha1_project_proto_goTypes = nil
	file_infragmo_auditum_v1alpha1_project_proto_depIdxs = nil
}

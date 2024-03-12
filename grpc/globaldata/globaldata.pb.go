// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v4.24.0
// source: globaldata.proto

package __

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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_globaldata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_globaldata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_globaldata_proto_rawDescGZIP(), []int{0}
}

type AllCountryList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AllCountryList) Reset() {
	*x = AllCountryList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_globaldata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllCountryList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllCountryList) ProtoMessage() {}

func (x *AllCountryList) ProtoReflect() protoreflect.Message {
	mi := &file_globaldata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllCountryList.ProtoReflect.Descriptor instead.
func (*AllCountryList) Descriptor() ([]byte, []int) {
	return file_globaldata_proto_rawDescGZIP(), []int{1}
}

func (x *AllCountryList) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AllCountryList) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AllCountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code        string            `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Description string            `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Result      []*AllCountryList `protobuf:"bytes,3,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *AllCountryResponse) Reset() {
	*x = AllCountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_globaldata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllCountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllCountryResponse) ProtoMessage() {}

func (x *AllCountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_globaldata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllCountryResponse.ProtoReflect.Descriptor instead.
func (*AllCountryResponse) Descriptor() ([]byte, []int) {
	return file_globaldata_proto_rawDescGZIP(), []int{2}
}

func (x *AllCountryResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *AllCountryResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AllCountryResponse) GetResult() []*AllCountryList {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_globaldata_proto protoreflect.FileDescriptor

var file_globaldata_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x22, 0x0e,
	0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x38,
	0x0a, 0x0e, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7e, 0x0a, 0x12, 0x41, 0x6c, 0x6c, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x60, 0x0a, 0x11, 0x47, 0x6c, 0x6f, 0x62,
	0x61, 0x6c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x18,
	0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x6c, 0x6f, 0x62, 0x61,
	0x6c, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x41, 0x6c, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_globaldata_proto_rawDescOnce sync.Once
	file_globaldata_proto_rawDescData = file_globaldata_proto_rawDesc
)

func file_globaldata_proto_rawDescGZIP() []byte {
	file_globaldata_proto_rawDescOnce.Do(func() {
		file_globaldata_proto_rawDescData = protoimpl.X.CompressGZIP(file_globaldata_proto_rawDescData)
	})
	return file_globaldata_proto_rawDescData
}

var file_globaldata_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_globaldata_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),       // 0: globaldata.EmptyRequest
	(*AllCountryList)(nil),     // 1: globaldata.AllCountryList
	(*AllCountryResponse)(nil), // 2: globaldata.AllCountryResponse
}
var file_globaldata_proto_depIdxs = []int32{
	1, // 0: globaldata.AllCountryResponse.result:type_name -> globaldata.AllCountryList
	0, // 1: globaldata.GlobalDataService.GetAllCountry:input_type -> globaldata.EmptyRequest
	2, // 2: globaldata.GlobalDataService.GetAllCountry:output_type -> globaldata.AllCountryResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_globaldata_proto_init() }
func file_globaldata_proto_init() {
	if File_globaldata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_globaldata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_globaldata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllCountryList); i {
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
		file_globaldata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllCountryResponse); i {
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
			RawDescriptor: file_globaldata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_globaldata_proto_goTypes,
		DependencyIndexes: file_globaldata_proto_depIdxs,
		MessageInfos:      file_globaldata_proto_msgTypes,
	}.Build()
	File_globaldata_proto = out.File
	file_globaldata_proto_rawDesc = nil
	file_globaldata_proto_goTypes = nil
	file_globaldata_proto_depIdxs = nil
}
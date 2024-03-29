// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v4.24.0
// source: rootadminweb.proto

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
		mi := &file_rootadminweb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rootadminweb_proto_msgTypes[0]
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
	return file_rootadminweb_proto_rawDescGZIP(), []int{0}
}

type DoLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *DoLoginRequest) Reset() {
	*x = DoLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rootadminweb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoLoginRequest) ProtoMessage() {}

func (x *DoLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rootadminweb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoLoginRequest.ProtoReflect.Descriptor instead.
func (*DoLoginRequest) Descriptor() ([]byte, []int) {
	return file_rootadminweb_proto_rawDescGZIP(), []int{1}
}

func (x *DoLoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DoLoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DoLoginList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Roleid    string `protobuf:"bytes,2,opt,name=roleid,proto3" json:"roleid,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone     string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Firstname string `protobuf:"bytes,5,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,6,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Fullname  string `protobuf:"bytes,7,opt,name=fullname,proto3" json:"fullname,omitempty"`
	Clientid  string `protobuf:"bytes,8,opt,name=clientid,proto3" json:"clientid,omitempty"`
}

func (x *DoLoginList) Reset() {
	*x = DoLoginList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rootadminweb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoLoginList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoLoginList) ProtoMessage() {}

func (x *DoLoginList) ProtoReflect() protoreflect.Message {
	mi := &file_rootadminweb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoLoginList.ProtoReflect.Descriptor instead.
func (*DoLoginList) Descriptor() ([]byte, []int) {
	return file_rootadminweb_proto_rawDescGZIP(), []int{2}
}

func (x *DoLoginList) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DoLoginList) GetRoleid() string {
	if x != nil {
		return x.Roleid
	}
	return ""
}

func (x *DoLoginList) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DoLoginList) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *DoLoginList) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *DoLoginList) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *DoLoginList) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

func (x *DoLoginList) GetClientid() string {
	if x != nil {
		return x.Clientid
	}
	return ""
}

type DoLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statuscode  string         `protobuf:"bytes,1,opt,name=statuscode,proto3" json:"statuscode,omitempty"`
	Description string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Session     string         `protobuf:"bytes,3,opt,name=session,proto3" json:"session,omitempty"`
	Result      []*DoLoginList `protobuf:"bytes,4,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *DoLoginResponse) Reset() {
	*x = DoLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rootadminweb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoLoginResponse) ProtoMessage() {}

func (x *DoLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rootadminweb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoLoginResponse.ProtoReflect.Descriptor instead.
func (*DoLoginResponse) Descriptor() ([]byte, []int) {
	return file_rootadminweb_proto_rawDescGZIP(), []int{3}
}

func (x *DoLoginResponse) GetStatuscode() string {
	if x != nil {
		return x.Statuscode
	}
	return ""
}

func (x *DoLoginResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DoLoginResponse) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

func (x *DoLoginResponse) GetResult() []*DoLoginList {
	if x != nil {
		return x.Result
	}
	return nil
}

type DoLogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *DoLogoutRequest) Reset() {
	*x = DoLogoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rootadminweb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoLogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoLogoutRequest) ProtoMessage() {}

func (x *DoLogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_rootadminweb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoLogoutRequest.ProtoReflect.Descriptor instead.
func (*DoLogoutRequest) Descriptor() ([]byte, []int) {
	return file_rootadminweb_proto_rawDescGZIP(), []int{4}
}

func (x *DoLogoutRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type DoLogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statuscode  string `protobuf:"bytes,1,opt,name=statuscode,proto3" json:"statuscode,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *DoLogoutResponse) Reset() {
	*x = DoLogoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rootadminweb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DoLogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DoLogoutResponse) ProtoMessage() {}

func (x *DoLogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rootadminweb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DoLogoutResponse.ProtoReflect.Descriptor instead.
func (*DoLogoutResponse) Descriptor() ([]byte, []int) {
	return file_rootadminweb_proto_rawDescGZIP(), []int{5}
}

func (x *DoLogoutResponse) GetStatuscode() string {
	if x != nil {
		return x.Statuscode
	}
	return ""
}

func (x *DoLogoutResponse) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_rootadminweb_proto protoreflect.FileDescriptor

var file_rootadminweb_proto_rawDesc = []byte{
	0x0a, 0x12, 0x72, 0x6f, 0x6f, 0x74, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x77, 0x65, 0x62, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x72, 0x6f, 0x6f, 0x74, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x77,
	0x65, 0x62, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x42, 0x0a, 0x0e, 0x44, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xd3, 0x01, 0x0a, 0x0b, 0x44, 0x6f, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x22, 0xa0, 0x01, 0x0a,
	0x0f, 0x44, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72,
	0x6f, 0x6f, 0x74, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x77, 0x65, 0x62, 0x2e, 0x44, 0x6f, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x27, 0x0a, 0x0f, 0x44, 0x6f, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x54, 0x0a, 0x10, 0x44, 0x6f, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xa8,
	0x01, 0x0a, 0x13, 0x52, 0x6f, 0x6f, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x57, 0x65, 0x62, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x07, 0x44, 0x6f, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x1c, 0x2e, 0x72, 0x6f, 0x6f, 0x74, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x77, 0x65, 0x62,
	0x2e, 0x44, 0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x72, 0x6f, 0x6f, 0x74, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x77, 0x65, 0x62, 0x2e, 0x44,
	0x6f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49,
	0x0a, 0x08, 0x44, 0x6f, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x1d, 0x2e, 0x72, 0x6f, 0x6f,
	0x74, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x77, 0x65, 0x62, 0x2e, 0x44, 0x6f, 0x4c, 0x6f, 0x67, 0x6f,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x72, 0x6f, 0x6f, 0x74,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x77, 0x65, 0x62, 0x2e, 0x44, 0x6f, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rootadminweb_proto_rawDescOnce sync.Once
	file_rootadminweb_proto_rawDescData = file_rootadminweb_proto_rawDesc
)

func file_rootadminweb_proto_rawDescGZIP() []byte {
	file_rootadminweb_proto_rawDescOnce.Do(func() {
		file_rootadminweb_proto_rawDescData = protoimpl.X.CompressGZIP(file_rootadminweb_proto_rawDescData)
	})
	return file_rootadminweb_proto_rawDescData
}

var file_rootadminweb_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rootadminweb_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),     // 0: rootadminweb.EmptyRequest
	(*DoLoginRequest)(nil),   // 1: rootadminweb.DoLoginRequest
	(*DoLoginList)(nil),      // 2: rootadminweb.DoLoginList
	(*DoLoginResponse)(nil),  // 3: rootadminweb.DoLoginResponse
	(*DoLogoutRequest)(nil),  // 4: rootadminweb.DoLogoutRequest
	(*DoLogoutResponse)(nil), // 5: rootadminweb.DoLogoutResponse
}
var file_rootadminweb_proto_depIdxs = []int32{
	2, // 0: rootadminweb.DoLoginResponse.result:type_name -> rootadminweb.DoLoginList
	1, // 1: rootadminweb.RootAdminWebService.DoLogin:input_type -> rootadminweb.DoLoginRequest
	4, // 2: rootadminweb.RootAdminWebService.DoLogout:input_type -> rootadminweb.DoLogoutRequest
	3, // 3: rootadminweb.RootAdminWebService.DoLogin:output_type -> rootadminweb.DoLoginResponse
	5, // 4: rootadminweb.RootAdminWebService.DoLogout:output_type -> rootadminweb.DoLogoutResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rootadminweb_proto_init() }
func file_rootadminweb_proto_init() {
	if File_rootadminweb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rootadminweb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_rootadminweb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoLoginRequest); i {
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
		file_rootadminweb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoLoginList); i {
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
		file_rootadminweb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoLoginResponse); i {
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
		file_rootadminweb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoLogoutRequest); i {
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
		file_rootadminweb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DoLogoutResponse); i {
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
			RawDescriptor: file_rootadminweb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rootadminweb_proto_goTypes,
		DependencyIndexes: file_rootadminweb_proto_depIdxs,
		MessageInfos:      file_rootadminweb_proto_msgTypes,
	}.Build()
	File_rootadminweb_proto = out.File
	file_rootadminweb_proto_rawDesc = nil
	file_rootadminweb_proto_goTypes = nil
	file_rootadminweb_proto_depIdxs = nil
}

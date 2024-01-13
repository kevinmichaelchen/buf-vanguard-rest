// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: bvr/v1beta1/api.proto

package bvrv1beta1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Foo) Reset() {
	*x = Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bvr_v1beta1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Foo) ProtoMessage() {}

func (x *Foo) ProtoReflect() protoreflect.Message {
	mi := &file_bvr_v1beta1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Foo.ProtoReflect.Descriptor instead.
func (*Foo) Descriptor() ([]byte, []int) {
	return file_bvr_v1beta1_api_proto_rawDescGZIP(), []int{0}
}

func (x *Foo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Foo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateFooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateFooRequest) Reset() {
	*x = CreateFooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bvr_v1beta1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFooRequest) ProtoMessage() {}

func (x *CreateFooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bvr_v1beta1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFooRequest.ProtoReflect.Descriptor instead.
func (*CreateFooRequest) Descriptor() ([]byte, []int) {
	return file_bvr_v1beta1_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateFooRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateFooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo *Foo `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (x *CreateFooResponse) Reset() {
	*x = CreateFooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bvr_v1beta1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFooResponse) ProtoMessage() {}

func (x *CreateFooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bvr_v1beta1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFooResponse.ProtoReflect.Descriptor instead.
func (*CreateFooResponse) Descriptor() ([]byte, []int) {
	return file_bvr_v1beta1_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFooResponse) GetFoo() *Foo {
	if x != nil {
		return x.Foo
	}
	return nil
}

type GetFooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetFooRequest) Reset() {
	*x = GetFooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bvr_v1beta1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFooRequest) ProtoMessage() {}

func (x *GetFooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bvr_v1beta1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFooRequest.ProtoReflect.Descriptor instead.
func (*GetFooRequest) Descriptor() ([]byte, []int) {
	return file_bvr_v1beta1_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetFooRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetFooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo *Foo `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
}

func (x *GetFooResponse) Reset() {
	*x = GetFooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bvr_v1beta1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFooResponse) ProtoMessage() {}

func (x *GetFooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bvr_v1beta1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFooResponse.ProtoReflect.Descriptor instead.
func (*GetFooResponse) Descriptor() ([]byte, []int) {
	return file_bvr_v1beta1_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetFooResponse) GetFoo() *Foo {
	if x != nil {
		return x.Foo
	}
	return nil
}

type ListFoosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parent    string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	PageSize  int32  `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListFoosRequest) Reset() {
	*x = ListFoosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bvr_v1beta1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFoosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFoosRequest) ProtoMessage() {}

func (x *ListFoosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bvr_v1beta1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFoosRequest.ProtoReflect.Descriptor instead.
func (*ListFoosRequest) Descriptor() ([]byte, []int) {
	return file_bvr_v1beta1_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListFoosRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListFoosRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListFoosRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListFoosResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foos []*Foo `protobuf:"bytes,1,rep,name=foos,proto3" json:"foos,omitempty"`
}

func (x *ListFoosResponse) Reset() {
	*x = ListFoosResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bvr_v1beta1_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFoosResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFoosResponse) ProtoMessage() {}

func (x *ListFoosResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bvr_v1beta1_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFoosResponse.ProtoReflect.Descriptor instead.
func (*ListFoosResponse) Descriptor() ([]byte, []int) {
	return file_bvr_v1beta1_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListFoosResponse) GetFoos() []*Foo {
	if x != nil {
		return x.Foos
	}
	return nil
}

type DeleteFooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFooRequest) Reset() {
	*x = DeleteFooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bvr_v1beta1_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFooRequest) ProtoMessage() {}

func (x *DeleteFooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bvr_v1beta1_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFooRequest.ProtoReflect.Descriptor instead.
func (*DeleteFooRequest) Descriptor() ([]byte, []int) {
	return file_bvr_v1beta1_api_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFooRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteFooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteFooResponse) Reset() {
	*x = DeleteFooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bvr_v1beta1_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFooResponse) ProtoMessage() {}

func (x *DeleteFooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bvr_v1beta1_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFooResponse.ProtoReflect.Descriptor instead.
func (*DeleteFooResponse) Descriptor() ([]byte, []int) {
	return file_bvr_v1beta1_api_proto_rawDescGZIP(), []int{8}
}

type PageToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PageToken) Reset() {
	*x = PageToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bvr_v1beta1_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageToken) ProtoMessage() {}

func (x *PageToken) ProtoReflect() protoreflect.Message {
	mi := &file_bvr_v1beta1_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageToken.ProtoReflect.Descriptor instead.
func (*PageToken) Descriptor() ([]byte, []int) {
	return file_bvr_v1beta1_api_proto_rawDescGZIP(), []int{9}
}

var File_bvr_v1beta1_api_proto protoreflect.FileDescriptor

var file_bvr_v1beta1_api_proto_rawDesc = []byte{
	0x0a, 0x15, 0x62, 0x76, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x62, 0x76, 0x72, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x29, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a,
	0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46,
	0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x03, 0x66, 0x6f,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x62, 0x76, 0x72, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x22, 0x1f,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x22, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x62, 0x76, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x6f, 0x6f,
	0x52, 0x03, 0x66, 0x6f, 0x6f, 0x22, 0x65, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x38, 0x0a, 0x10,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x24, 0x0a, 0x04, 0x66, 0x6f, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x62, 0x76, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x6f, 0x6f,
	0x52, 0x04, 0x66, 0x6f, 0x6f, 0x73, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0b, 0x0a, 0x09, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xca, 0x02, 0x0a,
	0x0a, 0x46, 0x6f, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x12, 0x1d, 0x2e, 0x62, 0x76, 0x72, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x76, 0x72, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6f, 0x12, 0x1a, 0x2e, 0x62, 0x76, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1b, 0x2e, 0x62, 0x76, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x66, 0x6f, 0x6f, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x12, 0x49, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x73, 0x12, 0x1c, 0x2e, 0x62,
	0x76, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x6f, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x62, 0x76, 0x72,
	0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x6f, 0x6f,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x09, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x12, 0x1d, 0x2e, 0x62, 0x76, 0x72, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x62, 0x76, 0x72, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xbb, 0x01, 0x0a, 0x0f, 0x63, 0x6f,
	0x6d, 0x2e, 0x62, 0x76, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x08, 0x41,
	0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x76, 0x69, 0x6e, 0x6d, 0x69, 0x63, 0x68, 0x61,
	0x65, 0x6c, 0x63, 0x68, 0x65, 0x6e, 0x2f, 0x62, 0x75, 0x66, 0x2d, 0x76, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x72, 0x64, 0x2d, 0x72, 0x65, 0x73, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x62, 0x76, 0x72, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x3b, 0x62, 0x76, 0x72, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x42,
	0x58, 0x58, 0xaa, 0x02, 0x0b, 0x42, 0x76, 0x72, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0xca, 0x02, 0x0b, 0x42, 0x76, 0x72, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02,
	0x17, 0x42, 0x76, 0x72, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x42, 0x76, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bvr_v1beta1_api_proto_rawDescOnce sync.Once
	file_bvr_v1beta1_api_proto_rawDescData = file_bvr_v1beta1_api_proto_rawDesc
)

func file_bvr_v1beta1_api_proto_rawDescGZIP() []byte {
	file_bvr_v1beta1_api_proto_rawDescOnce.Do(func() {
		file_bvr_v1beta1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_bvr_v1beta1_api_proto_rawDescData)
	})
	return file_bvr_v1beta1_api_proto_rawDescData
}

var file_bvr_v1beta1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_bvr_v1beta1_api_proto_goTypes = []interface{}{
	(*Foo)(nil),               // 0: bvr.v1beta1.Foo
	(*CreateFooRequest)(nil),  // 1: bvr.v1beta1.CreateFooRequest
	(*CreateFooResponse)(nil), // 2: bvr.v1beta1.CreateFooResponse
	(*GetFooRequest)(nil),     // 3: bvr.v1beta1.GetFooRequest
	(*GetFooResponse)(nil),    // 4: bvr.v1beta1.GetFooResponse
	(*ListFoosRequest)(nil),   // 5: bvr.v1beta1.ListFoosRequest
	(*ListFoosResponse)(nil),  // 6: bvr.v1beta1.ListFoosResponse
	(*DeleteFooRequest)(nil),  // 7: bvr.v1beta1.DeleteFooRequest
	(*DeleteFooResponse)(nil), // 8: bvr.v1beta1.DeleteFooResponse
	(*PageToken)(nil),         // 9: bvr.v1beta1.PageToken
}
var file_bvr_v1beta1_api_proto_depIdxs = []int32{
	0, // 0: bvr.v1beta1.CreateFooResponse.foo:type_name -> bvr.v1beta1.Foo
	0, // 1: bvr.v1beta1.GetFooResponse.foo:type_name -> bvr.v1beta1.Foo
	0, // 2: bvr.v1beta1.ListFoosResponse.foos:type_name -> bvr.v1beta1.Foo
	1, // 3: bvr.v1beta1.FooService.CreateFoo:input_type -> bvr.v1beta1.CreateFooRequest
	3, // 4: bvr.v1beta1.FooService.GetFoo:input_type -> bvr.v1beta1.GetFooRequest
	5, // 5: bvr.v1beta1.FooService.ListFoos:input_type -> bvr.v1beta1.ListFoosRequest
	7, // 6: bvr.v1beta1.FooService.DeleteFoo:input_type -> bvr.v1beta1.DeleteFooRequest
	2, // 7: bvr.v1beta1.FooService.CreateFoo:output_type -> bvr.v1beta1.CreateFooResponse
	4, // 8: bvr.v1beta1.FooService.GetFoo:output_type -> bvr.v1beta1.GetFooResponse
	6, // 9: bvr.v1beta1.FooService.ListFoos:output_type -> bvr.v1beta1.ListFoosResponse
	8, // 10: bvr.v1beta1.FooService.DeleteFoo:output_type -> bvr.v1beta1.DeleteFooResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_bvr_v1beta1_api_proto_init() }
func file_bvr_v1beta1_api_proto_init() {
	if File_bvr_v1beta1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bvr_v1beta1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Foo); i {
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
		file_bvr_v1beta1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFooRequest); i {
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
		file_bvr_v1beta1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFooResponse); i {
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
		file_bvr_v1beta1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFooRequest); i {
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
		file_bvr_v1beta1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFooResponse); i {
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
		file_bvr_v1beta1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFoosRequest); i {
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
		file_bvr_v1beta1_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFoosResponse); i {
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
		file_bvr_v1beta1_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFooRequest); i {
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
		file_bvr_v1beta1_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFooResponse); i {
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
		file_bvr_v1beta1_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageToken); i {
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
			RawDescriptor: file_bvr_v1beta1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bvr_v1beta1_api_proto_goTypes,
		DependencyIndexes: file_bvr_v1beta1_api_proto_depIdxs,
		MessageInfos:      file_bvr_v1beta1_api_proto_msgTypes,
	}.Build()
	File_bvr_v1beta1_api_proto = out.File
	file_bvr_v1beta1_api_proto_rawDesc = nil
	file_bvr_v1beta1_api_proto_goTypes = nil
	file_bvr_v1beta1_api_proto_depIdxs = nil
}

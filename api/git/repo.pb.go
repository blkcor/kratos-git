// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: api/git/repo.proto

package git

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

type ListRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *ListRepoRequest) Reset() {
	*x = ListRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_git_repo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepoRequest) ProtoMessage() {}

func (x *ListRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_git_repo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepoRequest.ProtoReflect.Descriptor instead.
func (*ListRepoRequest) Descriptor() ([]byte, []int) {
	return file_api_git_repo_proto_rawDescGZIP(), []int{0}
}

func (x *ListRepoRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRepoRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type ListRepoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cnt  int64           `protobuf:"varint,1,opt,name=cnt,proto3" json:"cnt,omitempty"`
	List []*ListRepoItem `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListRepoReply) Reset() {
	*x = ListRepoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_git_repo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepoReply) ProtoMessage() {}

func (x *ListRepoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_git_repo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepoReply.ProtoReflect.Descriptor instead.
func (*ListRepoReply) Descriptor() ([]byte, []int) {
	return file_api_git_repo_proto_rawDescGZIP(), []int{1}
}

func (x *ListRepoReply) GetCnt() int64 {
	if x != nil {
		return x.Cnt
	}
	return 0
}

func (x *ListRepoReply) GetList() []*ListRepoItem {
	if x != nil {
		return x.List
	}
	return nil
}

type ListRepoItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc     string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Path     string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Star     int32  `protobuf:"varint,5,opt,name=star,proto3" json:"star,omitempty"`
}

func (x *ListRepoItem) Reset() {
	*x = ListRepoItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_git_repo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepoItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepoItem) ProtoMessage() {}

func (x *ListRepoItem) ProtoReflect() protoreflect.Message {
	mi := &file_api_git_repo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepoItem.ProtoReflect.Descriptor instead.
func (*ListRepoItem) Descriptor() ([]byte, []int) {
	return file_api_git_repo_proto_rawDescGZIP(), []int{2}
}

func (x *ListRepoItem) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *ListRepoItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListRepoItem) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ListRepoItem) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ListRepoItem) GetStar() int32 {
	if x != nil {
		return x.Star
	}
	return 0
}

type CreateRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Type int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *CreateRepoRequest) Reset() {
	*x = CreateRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_git_repo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepoRequest) ProtoMessage() {}

func (x *CreateRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_git_repo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepoRequest.ProtoReflect.Descriptor instead.
func (*CreateRepoRequest) Descriptor() ([]byte, []int) {
	return file_api_git_repo_proto_rawDescGZIP(), []int{3}
}

func (x *CreateRepoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRepoRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreateRepoRequest) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateRepoRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type CreateRepoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc     string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Path     string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Star     int32  `protobuf:"varint,5,opt,name=star,proto3" json:"star,omitempty"`
}

func (x *CreateRepoReply) Reset() {
	*x = CreateRepoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_git_repo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRepoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRepoReply) ProtoMessage() {}

func (x *CreateRepoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_git_repo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRepoReply.ProtoReflect.Descriptor instead.
func (*CreateRepoReply) Descriptor() ([]byte, []int) {
	return file_api_git_repo_proto_rawDescGZIP(), []int{4}
}

func (x *CreateRepoReply) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *CreateRepoReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateRepoReply) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *CreateRepoReply) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *CreateRepoReply) GetStar() int32 {
	if x != nil {
		return x.Star
	}
	return 0
}

type UpdateRepoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Desc     string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Type     int32  `protobuf:"varint,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *UpdateRepoRequest) Reset() {
	*x = UpdateRepoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_git_repo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRepoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRepoRequest) ProtoMessage() {}

func (x *UpdateRepoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_git_repo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRepoRequest.ProtoReflect.Descriptor instead.
func (*UpdateRepoRequest) Descriptor() ([]byte, []int) {
	return file_api_git_repo_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateRepoRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *UpdateRepoRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateRepoRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *UpdateRepoRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

type UpdateRepoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateRepoReply) Reset() {
	*x = UpdateRepoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_git_repo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRepoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRepoReply) ProtoMessage() {}

func (x *UpdateRepoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_git_repo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRepoReply.ProtoReflect.Descriptor instead.
func (*UpdateRepoReply) Descriptor() ([]byte, []int) {
	return file_api_git_repo_proto_rawDescGZIP(), []int{6}
}

var File_api_git_repo_proto protoreflect.FileDescriptor

var file_api_git_repo_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x69, 0x74, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x69, 0x74, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x39, 0x0a, 0x0f, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x4c, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73,
	0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x69,
	0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x22, 0x7a, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x74, 0x61, 0x72,
	0x22, 0x63, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x7d, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x74, 0x61, 0x72, 0x22, 0x6b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73,
	0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x32, 0x80, 0x02, 0x0a, 0x04, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x4c, 0x0a,
	0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x69, 0x74, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x69, 0x74, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0e, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x12, 0x54, 0x0a, 0x0a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x69, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x69, 0x74, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a, 0x22, 0x05, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x12, 0x54, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x12,
	0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x69, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x70, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x69, 0x74, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x3a, 0x01, 0x2a,
	0x1a, 0x05, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x42, 0x23, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x2e, 0x67,
	0x69, 0x74, 0x50, 0x01, 0x5a, 0x16, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x69, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x69, 0x74, 0x3b, 0x67, 0x69, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_git_repo_proto_rawDescOnce sync.Once
	file_api_git_repo_proto_rawDescData = file_api_git_repo_proto_rawDesc
)

func file_api_git_repo_proto_rawDescGZIP() []byte {
	file_api_git_repo_proto_rawDescOnce.Do(func() {
		file_api_git_repo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_git_repo_proto_rawDescData)
	})
	return file_api_git_repo_proto_rawDescData
}

var file_api_git_repo_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_git_repo_proto_goTypes = []interface{}{
	(*ListRepoRequest)(nil),   // 0: api.git.ListRepoRequest
	(*ListRepoReply)(nil),     // 1: api.git.ListRepoReply
	(*ListRepoItem)(nil),      // 2: api.git.ListRepoItem
	(*CreateRepoRequest)(nil), // 3: api.git.CreateRepoRequest
	(*CreateRepoReply)(nil),   // 4: api.git.CreateRepoReply
	(*UpdateRepoRequest)(nil), // 5: api.git.UpdateRepoRequest
	(*UpdateRepoReply)(nil),   // 6: api.git.UpdateRepoReply
}
var file_api_git_repo_proto_depIdxs = []int32{
	2, // 0: api.git.ListRepoReply.list:type_name -> api.git.ListRepoItem
	0, // 1: api.git.Repo.ListRepo:input_type -> api.git.ListRepoRequest
	3, // 2: api.git.Repo.CreateRepo:input_type -> api.git.CreateRepoRequest
	5, // 3: api.git.Repo.UpdateRepo:input_type -> api.git.UpdateRepoRequest
	1, // 4: api.git.Repo.ListRepo:output_type -> api.git.ListRepoReply
	4, // 5: api.git.Repo.CreateRepo:output_type -> api.git.CreateRepoReply
	6, // 6: api.git.Repo.UpdateRepo:output_type -> api.git.UpdateRepoReply
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_git_repo_proto_init() }
func file_api_git_repo_proto_init() {
	if File_api_git_repo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_git_repo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepoRequest); i {
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
		file_api_git_repo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepoReply); i {
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
		file_api_git_repo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepoItem); i {
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
		file_api_git_repo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepoRequest); i {
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
		file_api_git_repo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRepoReply); i {
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
		file_api_git_repo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRepoRequest); i {
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
		file_api_git_repo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRepoReply); i {
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
			RawDescriptor: file_api_git_repo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_git_repo_proto_goTypes,
		DependencyIndexes: file_api_git_repo_proto_depIdxs,
		MessageInfos:      file_api_git_repo_proto_msgTypes,
	}.Build()
	File_api_git_repo_proto = out.File
	file_api_git_repo_proto_rawDesc = nil
	file_api_git_repo_proto_goTypes = nil
	file_api_git_repo_proto_depIdxs = nil
}

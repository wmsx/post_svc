// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: proto/post/post.proto

package go_micro_service_post_svc

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ErrorMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ErrorMsg) Reset() {
	*x = ErrorMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMsg) ProtoMessage() {}

func (x *ErrorMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorMsg.ProtoReflect.Descriptor instead.
func (*ErrorMsg) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SavePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type        int32       `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Title       string      `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string      `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	MengerId    int64       `protobuf:"varint,5,opt,name=mengerId,proto3" json:"mengerId,omitempty"`
	CategoryId  int64       `protobuf:"varint,6,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	Items       []*PostItem `protobuf:"bytes,7,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *SavePostRequest) Reset() {
	*x = SavePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePostRequest) ProtoMessage() {}

func (x *SavePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePostRequest.ProtoReflect.Descriptor instead.
func (*SavePostRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{1}
}

func (x *SavePostRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SavePostRequest) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *SavePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SavePostRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SavePostRequest) GetMengerId() int64 {
	if x != nil {
		return x.MengerId
	}
	return 0
}

func (x *SavePostRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *SavePostRequest) GetItems() []*PostItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type SavePostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMsg *ErrorMsg `protobuf:"bytes,999,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *SavePostResponse) Reset() {
	*x = SavePostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePostResponse) ProtoMessage() {}

func (x *SavePostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePostResponse.ProtoReflect.Descriptor instead.
func (*SavePostResponse) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{2}
}

func (x *SavePostResponse) GetErrorMsg() *ErrorMsg {
	if x != nil {
		return x.ErrorMsg
	}
	return nil
}

type GetPostListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryId int64 `protobuf:"varint,1,opt,name=categoryId,proto3" json:"categoryId,omitempty"`
	LastId     int64 `protobuf:"varint,2,opt,name=lastId,proto3" json:"lastId,omitempty"`
}

func (x *GetPostListRequest) Reset() {
	*x = GetPostListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostListRequest) ProtoMessage() {}

func (x *GetPostListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostListRequest.ProtoReflect.Descriptor instead.
func (*GetPostListRequest) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetPostListRequest) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *GetPostListRequest) GetLastId() int64 {
	if x != nil {
		return x.LastId
	}
	return 0
}

type GetPostListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostInfos []*PostInfo `protobuf:"bytes,1,rep,name=postInfos,proto3" json:"postInfos,omitempty"`
	ErrorMsg  *ErrorMsg   `protobuf:"bytes,999,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *GetPostListResponse) Reset() {
	*x = GetPostListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostListResponse) ProtoMessage() {}

func (x *GetPostListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostListResponse.ProtoReflect.Descriptor instead.
func (*GetPostListResponse) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{4}
}

func (x *GetPostListResponse) GetPostInfos() []*PostInfo {
	if x != nil {
		return x.PostInfos
	}
	return nil
}

func (x *GetPostListResponse) GetErrorMsg() *ErrorMsg {
	if x != nil {
		return x.ErrorMsg
	}
	return nil
}

type PostInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Type        int32       `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Title       string      `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string      `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	MengerId    int64       `protobuf:"varint,5,opt,name=mengerId,proto3" json:"mengerId,omitempty"`
	Item        []*PostItem `protobuf:"bytes,6,rep,name=item,proto3" json:"item,omitempty"`
}

func (x *PostInfo) Reset() {
	*x = PostInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostInfo) ProtoMessage() {}

func (x *PostInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostInfo.ProtoReflect.Descriptor instead.
func (*PostInfo) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{5}
}

func (x *PostInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PostInfo) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *PostInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PostInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PostInfo) GetMengerId() int64 {
	if x != nil {
		return x.MengerId
	}
	return 0
}

func (x *PostInfo) GetItem() []*PostItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type PostItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjectId int64 `protobuf:"varint,1,opt,name=objectId,proto3" json:"objectId,omitempty"`
	Index    int32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *PostItem) Reset() {
	*x = PostItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_post_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostItem) ProtoMessage() {}

func (x *PostItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_post_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostItem.ProtoReflect.Descriptor instead.
func (*PostItem) Descriptor() ([]byte, []int) {
	return file_proto_post_post_proto_rawDescGZIP(), []int{6}
}

func (x *PostItem) GetObjectId() int64 {
	if x != nil {
		return x.ObjectId
	}
	return 0
}

func (x *PostItem) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

var File_proto_post_post_proto protoreflect.FileDescriptor

var file_proto_post_post_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73,
	0x76, 0x63, 0x22, 0x1c, 0x0a, 0x08, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0xe4, 0x01, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x54, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0xe7, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x73, 0x67, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x4c, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x22, 0x9a, 0x01, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73,
	0x76, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x70, 0x6f, 0x73,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x40, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x73, 0x67, 0x18, 0xe7, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x08,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0xbb, 0x01, 0x0a, 0x08, 0x50, 0x6f, 0x73,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x12, 0x37, 0x0a,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x3c, 0x0a, 0x08, 0x50, 0x6f, 0x73, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x32, 0xdd, 0x01, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x65, 0x0a,
	0x08, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x2a, 0x2e, 0x67, 0x6f, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73,
	0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76,
	0x63, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2d, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_post_post_proto_rawDescOnce sync.Once
	file_proto_post_post_proto_rawDescData = file_proto_post_post_proto_rawDesc
)

func file_proto_post_post_proto_rawDescGZIP() []byte {
	file_proto_post_post_proto_rawDescOnce.Do(func() {
		file_proto_post_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_post_post_proto_rawDescData)
	})
	return file_proto_post_post_proto_rawDescData
}

var file_proto_post_post_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_post_post_proto_goTypes = []interface{}{
	(*ErrorMsg)(nil),            // 0: go.micro.service.post_svc.ErrorMsg
	(*SavePostRequest)(nil),     // 1: go.micro.service.post_svc.SavePostRequest
	(*SavePostResponse)(nil),    // 2: go.micro.service.post_svc.SavePostResponse
	(*GetPostListRequest)(nil),  // 3: go.micro.service.post_svc.GetPostListRequest
	(*GetPostListResponse)(nil), // 4: go.micro.service.post_svc.GetPostListResponse
	(*PostInfo)(nil),            // 5: go.micro.service.post_svc.PostInfo
	(*PostItem)(nil),            // 6: go.micro.service.post_svc.PostItem
}
var file_proto_post_post_proto_depIdxs = []int32{
	6, // 0: go.micro.service.post_svc.SavePostRequest.items:type_name -> go.micro.service.post_svc.PostItem
	0, // 1: go.micro.service.post_svc.SavePostResponse.errorMsg:type_name -> go.micro.service.post_svc.ErrorMsg
	5, // 2: go.micro.service.post_svc.GetPostListResponse.postInfos:type_name -> go.micro.service.post_svc.PostInfo
	0, // 3: go.micro.service.post_svc.GetPostListResponse.errorMsg:type_name -> go.micro.service.post_svc.ErrorMsg
	6, // 4: go.micro.service.post_svc.PostInfo.item:type_name -> go.micro.service.post_svc.PostItem
	1, // 5: go.micro.service.post_svc.Post.SavePost:input_type -> go.micro.service.post_svc.SavePostRequest
	3, // 6: go.micro.service.post_svc.Post.GetPostList:input_type -> go.micro.service.post_svc.GetPostListRequest
	2, // 7: go.micro.service.post_svc.Post.SavePost:output_type -> go.micro.service.post_svc.SavePostResponse
	4, // 8: go.micro.service.post_svc.Post.GetPostList:output_type -> go.micro.service.post_svc.GetPostListResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_post_post_proto_init() }
func file_proto_post_post_proto_init() {
	if File_proto_post_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_post_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ErrorMsg); i {
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
		file_proto_post_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePostRequest); i {
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
		file_proto_post_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SavePostResponse); i {
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
		file_proto_post_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostListRequest); i {
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
		file_proto_post_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostListResponse); i {
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
		file_proto_post_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostInfo); i {
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
		file_proto_post_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostItem); i {
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
			RawDescriptor: file_proto_post_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_post_post_proto_goTypes,
		DependencyIndexes: file_proto_post_post_proto_depIdxs,
		MessageInfos:      file_proto_post_post_proto_msgTypes,
	}.Build()
	File_proto_post_post_proto = out.File
	file_proto_post_post_proto_rawDesc = nil
	file_proto_post_post_proto_goTypes = nil
	file_proto_post_post_proto_depIdxs = nil
}
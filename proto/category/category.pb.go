// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.11.4
// source: proto/category/category.proto

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
		mi := &file_proto_category_category_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ErrorMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorMsg) ProtoMessage() {}

func (x *ErrorMsg) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[0]
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
	return file_proto_category_category_proto_rawDescGZIP(), []int{0}
}

func (x *ErrorMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type CreateCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	MengerId int64  `protobuf:"varint,2,opt,name=mengerId,proto3" json:"mengerId,omitempty"`
}

func (x *CreateCategoryRequest) Reset() {
	*x = CreateCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryRequest) ProtoMessage() {}

func (x *CreateCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryRequest.ProtoReflect.Descriptor instead.
func (*CreateCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCategoryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCategoryRequest) GetMengerId() int64 {
	if x != nil {
		return x.MengerId
	}
	return 0
}

type CreateCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ErrorMsg *ErrorMsg `protobuf:"bytes,999,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *CreateCategoryResponse) Reset() {
	*x = CreateCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCategoryResponse) ProtoMessage() {}

func (x *CreateCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCategoryResponse.ProtoReflect.Descriptor instead.
func (*CreateCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCategoryResponse) GetErrorMsg() *ErrorMsg {
	if x != nil {
		return x.ErrorMsg
	}
	return nil
}

type GetAllCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllCategoryRequest) Reset() {
	*x = GetAllCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCategoryRequest) ProtoMessage() {}

func (x *GetAllCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetAllCategoryRequest) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{3}
}

type CategoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MengerId int64  `protobuf:"varint,3,opt,name=mengerId,proto3" json:"mengerId,omitempty"`
}

func (x *CategoryInfo) Reset() {
	*x = CategoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CategoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryInfo) ProtoMessage() {}

func (x *CategoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryInfo.ProtoReflect.Descriptor instead.
func (*CategoryInfo) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{4}
}

func (x *CategoryInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CategoryInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryInfo) GetMengerId() int64 {
	if x != nil {
		return x.MengerId
	}
	return 0
}

type GetAllCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CategoryInfos []*CategoryInfo `protobuf:"bytes,1,rep,name=categoryInfos,proto3" json:"categoryInfos,omitempty"`
	ErrorMsg      *ErrorMsg       `protobuf:"bytes,999,opt,name=errorMsg,proto3" json:"errorMsg,omitempty"`
}

func (x *GetAllCategoryResponse) Reset() {
	*x = GetAllCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_category_category_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllCategoryResponse) ProtoMessage() {}

func (x *GetAllCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_category_category_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetAllCategoryResponse) Descriptor() ([]byte, []int) {
	return file_proto_category_category_proto_rawDescGZIP(), []int{5}
}

func (x *GetAllCategoryResponse) GetCategoryInfos() []*CategoryInfo {
	if x != nil {
		return x.CategoryInfos
	}
	return nil
}

func (x *GetAllCategoryResponse) GetErrorMsg() *ErrorMsg {
	if x != nil {
		return x.ErrorMsg
	}
	return nil
}

var File_proto_category_category_proto protoreflect.FileDescriptor

var file_proto_category_category_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x19, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x22, 0x1c, 0x0a, 0x08, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x47, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x67, 0x65, 0x72, 0x49,
	0x64, 0x22, 0x5a, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0xe7, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x4d, 0x73, 0x67, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x22, 0x17, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4e, 0x0a, 0x0c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65,
	0x6e, 0x67, 0x65, 0x72, 0x49, 0x64, 0x22, 0xa9, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4d, 0x0a, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74,
	0x5f, 0x73, 0x76, 0x63, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0d, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x73,
	0x12, 0x40, 0x0a, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x18, 0xe7, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x73, 0x67, 0x32, 0xfc, 0x01, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12,
	0x77, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x30, 0x2e, 0x67, 0x6f, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x67,
	0x6f, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_category_category_proto_rawDescOnce sync.Once
	file_proto_category_category_proto_rawDescData = file_proto_category_category_proto_rawDesc
)

func file_proto_category_category_proto_rawDescGZIP() []byte {
	file_proto_category_category_proto_rawDescOnce.Do(func() {
		file_proto_category_category_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_category_category_proto_rawDescData)
	})
	return file_proto_category_category_proto_rawDescData
}

var file_proto_category_category_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_category_category_proto_goTypes = []interface{}{
	(*ErrorMsg)(nil),               // 0: go.micro.service.post_svc.ErrorMsg
	(*CreateCategoryRequest)(nil),  // 1: go.micro.service.post_svc.CreateCategoryRequest
	(*CreateCategoryResponse)(nil), // 2: go.micro.service.post_svc.CreateCategoryResponse
	(*GetAllCategoryRequest)(nil),  // 3: go.micro.service.post_svc.GetAllCategoryRequest
	(*CategoryInfo)(nil),           // 4: go.micro.service.post_svc.CategoryInfo
	(*GetAllCategoryResponse)(nil), // 5: go.micro.service.post_svc.GetAllCategoryResponse
}
var file_proto_category_category_proto_depIdxs = []int32{
	0, // 0: go.micro.service.post_svc.CreateCategoryResponse.errorMsg:type_name -> go.micro.service.post_svc.ErrorMsg
	4, // 1: go.micro.service.post_svc.GetAllCategoryResponse.categoryInfos:type_name -> go.micro.service.post_svc.CategoryInfo
	0, // 2: go.micro.service.post_svc.GetAllCategoryResponse.errorMsg:type_name -> go.micro.service.post_svc.ErrorMsg
	1, // 3: go.micro.service.post_svc.Category.CreateCategory:input_type -> go.micro.service.post_svc.CreateCategoryRequest
	3, // 4: go.micro.service.post_svc.Category.GetAllCategory:input_type -> go.micro.service.post_svc.GetAllCategoryRequest
	2, // 5: go.micro.service.post_svc.Category.CreateCategory:output_type -> go.micro.service.post_svc.CreateCategoryResponse
	5, // 6: go.micro.service.post_svc.Category.GetAllCategory:output_type -> go.micro.service.post_svc.GetAllCategoryResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_category_category_proto_init() }
func file_proto_category_category_proto_init() {
	if File_proto_category_category_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_category_category_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_category_category_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryRequest); i {
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
		file_proto_category_category_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCategoryResponse); i {
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
		file_proto_category_category_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCategoryRequest); i {
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
		file_proto_category_category_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CategoryInfo); i {
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
		file_proto_category_category_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllCategoryResponse); i {
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
			RawDescriptor: file_proto_category_category_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_category_category_proto_goTypes,
		DependencyIndexes: file_proto_category_category_proto_depIdxs,
		MessageInfos:      file_proto_category_category_proto_msgTypes,
	}.Build()
	File_proto_category_category_proto = out.File
	file_proto_category_category_proto_rawDesc = nil
	file_proto_category_category_proto_goTypes = nil
	file_proto_category_category_proto_depIdxs = nil
}

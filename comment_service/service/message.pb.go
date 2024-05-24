// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: message.proto

package service

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

type State int32

const (
	State_NORMAL State = 0
	State_HIDDEN State = 1
)

// Enum value maps for State.
var (
	State_name = map[int32]string{
		0: "NORMAL",
		1: "HIDDEN",
	}
	State_value = map[string]int32{
		"NORMAL": 0,
		"HIDDEN": 1,
	}
)

func (x State) Enum() *State {
	p := new(State)
	*p = x
	return p
}

func (x State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (State) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (State) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use State.Descriptor instead.
func (State) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type ObjType int32

const (
	ObjType_Video   ObjType = 0
	ObjType_Article ObjType = 1
)

// Enum value maps for ObjType.
var (
	ObjType_name = map[int32]string{
		0: "Video",
		1: "Article",
	}
	ObjType_value = map[string]int32{
		"Video":   0,
		"Article": 1,
	}
)

func (x ObjType) Enum() *ObjType {
	p := new(ObjType)
	*p = x
	return p
}

func (x ObjType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[1].Descriptor()
}

func (ObjType) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[1]
}

func (x ObjType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjType.Descriptor instead.
func (ObjType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

type CreateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjId    int64   `protobuf:"varint,1,opt,name=obj_id,json=objId,proto3" json:"obj_id,omitempty"`
	MemberId int64   `protobuf:"varint,2,opt,name=member_id,json=memberId,proto3" json:"member_id,omitempty"`
	State    State   `protobuf:"varint,3,opt,name=state,proto3,enum=message.State" json:"state,omitempty"`
	ObjType  ObjType `protobuf:"varint,4,opt,name=obj_type,json=objType,proto3,enum=message.ObjType" json:"obj_type,omitempty"`
	Root     int64   `protobuf:"varint,5,opt,name=root,proto3" json:"root,omitempty"`
	Parent   int64   `protobuf:"varint,6,opt,name=parent,proto3" json:"parent,omitempty"`
	Floor    int32   `protobuf:"varint,7,opt,name=floor,proto3" json:"floor,omitempty"`
	Ip       int64   `protobuf:"varint,8,opt,name=ip,proto3" json:"ip,omitempty"`
	Comment  string  `protobuf:"bytes,9,opt,name=comment,proto3" json:"comment,omitempty"` // 评论内容
}

func (x *CreateMessageRequest) Reset() {
	*x = CreateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRequest) ProtoMessage() {}

func (x *CreateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMessageRequest) GetObjId() int64 {
	if x != nil {
		return x.ObjId
	}
	return 0
}

func (x *CreateMessageRequest) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

func (x *CreateMessageRequest) GetState() State {
	if x != nil {
		return x.State
	}
	return State_NORMAL
}

func (x *CreateMessageRequest) GetObjType() ObjType {
	if x != nil {
		return x.ObjType
	}
	return ObjType_Video
}

func (x *CreateMessageRequest) GetRoot() int64 {
	if x != nil {
		return x.Root
	}
	return 0
}

func (x *CreateMessageRequest) GetParent() int64 {
	if x != nil {
		return x.Parent
	}
	return 0
}

func (x *CreateMessageRequest) GetFloor() int32 {
	if x != nil {
		return x.Floor
	}
	return 0
}

func (x *CreateMessageRequest) GetIp() int64 {
	if x != nil {
		return x.Ip
	}
	return 0
}

func (x *CreateMessageRequest) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type CreateMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CreateMessageResponse) Reset() {
	*x = CreateMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageResponse) ProtoMessage() {}

func (x *CreateMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageResponse.ProtoReflect.Descriptor instead.
func (*CreateMessageResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *CreateMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ObjID   int64 `protobuf:"varint,4,opt,name=ObjID,proto3" json:"ObjID,omitempty"`
	ObjType int32 `protobuf:"varint,10,opt,name=ObjType,proto3" json:"ObjType,omitempty"`
}

func (x *GetCommentRequest) Reset() {
	*x = GetCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentRequest) ProtoMessage() {}

func (x *GetCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentRequest.ProtoReflect.Descriptor instead.
func (*GetCommentRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *GetCommentRequest) GetObjID() int64 {
	if x != nil {
		return x.ObjID
	}
	return 0
}

func (x *GetCommentRequest) GetObjType() int32 {
	if x != nil {
		return x.ObjType
	}
	return 0
}

type GetCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Root      int64  `protobuf:"varint,1,opt,name=Root,proto3" json:"Root,omitempty"`
	Parent    int64  `protobuf:"varint,2,opt,name=Parent,proto3" json:"Parent,omitempty"`
	MemberID  int64  `protobuf:"varint,3,opt,name=MemberID,proto3" json:"MemberID,omitempty"`
	ObjID     int64  `protobuf:"varint,4,opt,name=ObjID,proto3" json:"ObjID,omitempty"`
	Count     int32  `protobuf:"varint,5,opt,name=Count,proto3" json:"Count,omitempty"`
	RootCount int32  `protobuf:"varint,6,opt,name=RootCount,proto3" json:"RootCount,omitempty"`
	Like      int32  `protobuf:"varint,7,opt,name=Like,proto3" json:"Like,omitempty"`
	Hate      int32  `protobuf:"varint,8,opt,name=Hate,proto3" json:"Hate,omitempty"`
	State     int32  `protobuf:"varint,9,opt,name=State,proto3" json:"State,omitempty"`
	ObjType   int32  `protobuf:"varint,10,opt,name=ObjType,proto3" json:"ObjType,omitempty"`
	Floor     int32  `protobuf:"varint,11,opt,name=Floor,proto3" json:"Floor,omitempty"`
	IP        int64  `protobuf:"varint,12,opt,name=IP,proto3" json:"IP,omitempty"`
	Message   string `protobuf:"bytes,13,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *GetCommentResponse) Reset() {
	*x = GetCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentResponse) ProtoMessage() {}

func (x *GetCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentResponse.ProtoReflect.Descriptor instead.
func (*GetCommentResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetCommentResponse) GetRoot() int64 {
	if x != nil {
		return x.Root
	}
	return 0
}

func (x *GetCommentResponse) GetParent() int64 {
	if x != nil {
		return x.Parent
	}
	return 0
}

func (x *GetCommentResponse) GetMemberID() int64 {
	if x != nil {
		return x.MemberID
	}
	return 0
}

func (x *GetCommentResponse) GetObjID() int64 {
	if x != nil {
		return x.ObjID
	}
	return 0
}

func (x *GetCommentResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *GetCommentResponse) GetRootCount() int32 {
	if x != nil {
		return x.RootCount
	}
	return 0
}

func (x *GetCommentResponse) GetLike() int32 {
	if x != nil {
		return x.Like
	}
	return 0
}

func (x *GetCommentResponse) GetHate() int32 {
	if x != nil {
		return x.Hate
	}
	return 0
}

func (x *GetCommentResponse) GetState() int32 {
	if x != nil {
		return x.State
	}
	return 0
}

func (x *GetCommentResponse) GetObjType() int32 {
	if x != nil {
		return x.ObjType
	}
	return 0
}

func (x *GetCommentResponse) GetFloor() int32 {
	if x != nil {
		return x.Floor
	}
	return 0
}

func (x *GetCommentResponse) GetIP() int64 {
	if x != nil {
		return x.IP
	}
	return 0
}

func (x *GetCommentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x89, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x15, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x6f, 0x62, 0x6a, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x08, 0x6f,
	0x62, 0x6a, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x07, 0x6f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x66, 0x6c, 0x6f, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0x31, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x43, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x4f, 0x62, 0x6a, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4f, 0x62, 0x6a,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x4f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x22, 0xbe, 0x02, 0x0a,
	0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4f,
	0x62, 0x6a, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4f, 0x62, 0x6a, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x6f, 0x6f, 0x74, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x52, 0x6f, 0x6f, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x48, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x4f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x4f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x46, 0x6c, 0x6f, 0x6f, 0x72, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x46, 0x6c,
	0x6f, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x50, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x49, 0x50, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x1f, 0x0a,
	0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x49, 0x44, 0x44, 0x45, 0x4e, 0x10, 0x01, 0x2a, 0x21,
	0x0a, 0x07, 0x4f, 0x62, 0x6a, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x10,
	0x01, 0x32, 0xae, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_message_proto_goTypes = []interface{}{
	(State)(0),                    // 0: message.State
	(ObjType)(0),                  // 1: message.ObjType
	(*CreateMessageRequest)(nil),  // 2: message.CreateMessageRequest
	(*CreateMessageResponse)(nil), // 3: message.CreateMessageResponse
	(*GetCommentRequest)(nil),     // 4: message.GetCommentRequest
	(*GetCommentResponse)(nil),    // 5: message.GetCommentResponse
}
var file_message_proto_depIdxs = []int32{
	0, // 0: message.CreateMessageRequest.state:type_name -> message.State
	1, // 1: message.CreateMessageRequest.obj_type:type_name -> message.ObjType
	2, // 2: message.MessageService.CreateCommentMessage:input_type -> message.CreateMessageRequest
	4, // 3: message.MessageService.GetComment:input_type -> message.GetCommentRequest
	3, // 4: message.MessageService.CreateCommentMessage:output_type -> message.CreateMessageResponse
	5, // 5: message.MessageService.GetComment:output_type -> message.GetCommentResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageRequest); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageResponse); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentRequest); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentResponse); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}

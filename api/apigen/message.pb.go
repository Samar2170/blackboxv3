// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: message.proto

package blackboxv3

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Status int32

const (
	Status_PENDING    Status = 0
	Status_INPROGRESS Status = 1
	Status_SUCCESS    Status = 2
	Status_FAILED     Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "PENDING",
		1: "INPROGRESS",
		2: "SUCCESS",
		3: "FAILED",
	}
	Status_value = map[string]int32{
		"PENDING":    0,
		"INPROGRESS": 1,
		"SUCCESS":    2,
		"FAILED":     3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ChannelId int32  `protobuf:"varint,2,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Msg       string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SendMessageRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *SendMessageRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SendMessageResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	ChannelId int32  `protobuf:"varint,2,opt,name=channelId,proto3" json:"channelId,omitempty"`
}

func (x *GetMessagesRequest) Reset() {
	*x = GetMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesRequest) ProtoMessage() {}

func (x *GetMessagesRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use GetMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetMessagesRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *GetMessagesRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetMessagesRequest) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

type TextMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg    string               `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	SentAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=sentAt,proto3" json:"sentAt,omitempty"`
}

func (x *TextMessage) Reset() {
	*x = TextMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextMessage) ProtoMessage() {}

func (x *TextMessage) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TextMessage.ProtoReflect.Descriptor instead.
func (*TextMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *TextMessage) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TextMessage) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *TextMessage) GetSentAt() *timestamp.Timestamp {
	if x != nil {
		return x.SentAt
	}
	return nil
}

type GetMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success  bool           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg      string         `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Messages []*TextMessage `protobuf:"bytes,3,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GetMessagesResponse) Reset() {
	*x = GetMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessagesResponse) ProtoMessage() {}

func (x *GetMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetMessagesResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *GetMessagesResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *GetMessagesResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetMessagesResponse) GetMessages() []*TextMessage {
	if x != nil {
		return x.Messages
	}
	return nil
}

type FileMetaData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Size string `protobuf:"bytes,2,opt,name=size,proto3" json:"size,omitempty"`
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *FileMetaData) Reset() {
	*x = FileMetaData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetaData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetaData) ProtoMessage() {}

func (x *FileMetaData) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetaData.ProtoReflect.Descriptor instead.
func (*FileMetaData) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *FileMetaData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileMetaData) GetSize() string {
	if x != nil {
		return x.Size
	}
	return ""
}

func (x *FileMetaData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type SendMediaMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Media:
	//
	//	*SendMediaMessageRequest_FileMetaData
	//	*SendMediaMessageRequest_Token
	//	*SendMediaMessageRequest_ChannelId
	//	*SendMediaMessageRequest_Chunk
	Media isSendMediaMessageRequest_Media `protobuf_oneof:"media"`
}

func (x *SendMediaMessageRequest) Reset() {
	*x = SendMediaMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMediaMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMediaMessageRequest) ProtoMessage() {}

func (x *SendMediaMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMediaMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMediaMessageRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{6}
}

func (m *SendMediaMessageRequest) GetMedia() isSendMediaMessageRequest_Media {
	if m != nil {
		return m.Media
	}
	return nil
}

func (x *SendMediaMessageRequest) GetFileMetaData() *FileMetaData {
	if x, ok := x.GetMedia().(*SendMediaMessageRequest_FileMetaData); ok {
		return x.FileMetaData
	}
	return nil
}

func (x *SendMediaMessageRequest) GetToken() string {
	if x, ok := x.GetMedia().(*SendMediaMessageRequest_Token); ok {
		return x.Token
	}
	return ""
}

func (x *SendMediaMessageRequest) GetChannelId() int32 {
	if x, ok := x.GetMedia().(*SendMediaMessageRequest_ChannelId); ok {
		return x.ChannelId
	}
	return 0
}

func (x *SendMediaMessageRequest) GetChunk() []byte {
	if x, ok := x.GetMedia().(*SendMediaMessageRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isSendMediaMessageRequest_Media interface {
	isSendMediaMessageRequest_Media()
}

type SendMediaMessageRequest_FileMetaData struct {
	FileMetaData *FileMetaData `protobuf:"bytes,1,opt,name=fileMetaData,proto3,oneof"`
}

type SendMediaMessageRequest_Token struct {
	Token string `protobuf:"bytes,2,opt,name=token,proto3,oneof"`
}

type SendMediaMessageRequest_ChannelId struct {
	ChannelId int32 `protobuf:"varint,3,opt,name=channelId,proto3,oneof"`
}

type SendMediaMessageRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,4,opt,name=chunk,proto3,oneof"`
}

func (*SendMediaMessageRequest_FileMetaData) isSendMediaMessageRequest_Media() {}

func (*SendMediaMessageRequest_Token) isSendMediaMessageRequest_Media() {}

func (*SendMediaMessageRequest_ChannelId) isSendMediaMessageRequest_Media() {}

func (*SendMediaMessageRequest_Chunk) isSendMediaMessageRequest_Media() {}

type SendMediaMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Status  Status `protobuf:"varint,3,opt,name=status,proto3,enum=Status" json:"status,omitempty"`
}

func (x *SendMediaMessageResponse) Reset() {
	*x = SendMediaMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMediaMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMediaMessageResponse) ProtoMessage() {}

func (x *SendMediaMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMediaMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMediaMessageResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{7}
}

func (x *SendMediaMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *SendMediaMessageResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SendMediaMessageResponse) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_PENDING
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x5a, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x41, 0x0a, 0x13,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22,
	0x48, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x0b, 0x54, 0x65, 0x78,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x65,
	0x6e, 0x74, 0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x22, 0x6b,
	0x0a, 0x13, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x28, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x4a, 0x0a, 0x0c, 0x46,
	0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x0c, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x0c, 0x66, 0x69, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x1e, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64,
	0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x07, 0x0a, 0x05, 0x6d, 0x65, 0x64, 0x69,
	0x61, 0x22, 0x67, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x64, 0x69, 0x61, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1f, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x07, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x3e, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0a,
	0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x61, 0x6d, 0x61, 0x72, 0x32, 0x31,
	0x37, 0x30, 0x2f, 0x62, 0x6c, 0x61, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x76, 0x33, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_message_proto_goTypes = []any{
	(Status)(0),                      // 0: Status
	(*SendMessageRequest)(nil),       // 1: SendMessageRequest
	(*SendMessageResponse)(nil),      // 2: SendMessageResponse
	(*GetMessagesRequest)(nil),       // 3: GetMessagesRequest
	(*TextMessage)(nil),              // 4: TextMessage
	(*GetMessagesResponse)(nil),      // 5: GetMessagesResponse
	(*FileMetaData)(nil),             // 6: FileMetaData
	(*SendMediaMessageRequest)(nil),  // 7: SendMediaMessageRequest
	(*SendMediaMessageResponse)(nil), // 8: SendMediaMessageResponse
	(*timestamp.Timestamp)(nil),      // 9: google.protobuf.Timestamp
}
var file_message_proto_depIdxs = []int32{
	9, // 0: TextMessage.sentAt:type_name -> google.protobuf.Timestamp
	4, // 1: GetMessagesResponse.messages:type_name -> TextMessage
	6, // 2: SendMediaMessageRequest.fileMetaData:type_name -> FileMetaData
	0, // 3: SendMediaMessageResponse.status:type_name -> Status
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SendMessageRequest); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SendMessageResponse); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetMessagesRequest); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TextMessage); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetMessagesResponse); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*FileMetaData); i {
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
		file_message_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*SendMediaMessageRequest); i {
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
		file_message_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*SendMediaMessageResponse); i {
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
	file_message_proto_msgTypes[6].OneofWrappers = []any{
		(*SendMediaMessageRequest_FileMetaData)(nil),
		(*SendMediaMessageRequest_Token)(nil),
		(*SendMediaMessageRequest_ChannelId)(nil),
		(*SendMediaMessageRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
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

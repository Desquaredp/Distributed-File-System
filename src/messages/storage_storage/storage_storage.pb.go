// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: storage_storage.proto

package storage_storage

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

type StorageNodeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to StorageNodeMessage:
	//
	//	*StorageNodeMessage_PutCopy
	//	*StorageNodeMessage_PutCopyResponse
	//	*StorageNodeMessage_GetReplica
	//	*StorageNodeMessage_GetReplicaResponse
	StorageNodeMessage isStorageNodeMessage_StorageNodeMessage `protobuf_oneof:"storage_node_message"`
}

func (x *StorageNodeMessage) Reset() {
	*x = StorageNodeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageNodeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageNodeMessage) ProtoMessage() {}

func (x *StorageNodeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageNodeMessage.ProtoReflect.Descriptor instead.
func (*StorageNodeMessage) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{0}
}

func (m *StorageNodeMessage) GetStorageNodeMessage() isStorageNodeMessage_StorageNodeMessage {
	if m != nil {
		return m.StorageNodeMessage
	}
	return nil
}

func (x *StorageNodeMessage) GetPutCopy() *StorageNodeMessage_PUTCopy {
	if x, ok := x.GetStorageNodeMessage().(*StorageNodeMessage_PutCopy); ok {
		return x.PutCopy
	}
	return nil
}

func (x *StorageNodeMessage) GetPutCopyResponse() *StorageNodeMessage_PUTCopyResponse {
	if x, ok := x.GetStorageNodeMessage().(*StorageNodeMessage_PutCopyResponse); ok {
		return x.PutCopyResponse
	}
	return nil
}

func (x *StorageNodeMessage) GetGetReplica() *StorageNodeMessage_GETReplica {
	if x, ok := x.GetStorageNodeMessage().(*StorageNodeMessage_GetReplica); ok {
		return x.GetReplica
	}
	return nil
}

func (x *StorageNodeMessage) GetGetReplicaResponse() *StorageNodeMessage_GETReplicaResponse {
	if x, ok := x.GetStorageNodeMessage().(*StorageNodeMessage_GetReplicaResponse); ok {
		return x.GetReplicaResponse
	}
	return nil
}

type isStorageNodeMessage_StorageNodeMessage interface {
	isStorageNodeMessage_StorageNodeMessage()
}

type StorageNodeMessage_PutCopy struct {
	PutCopy *StorageNodeMessage_PUTCopy `protobuf:"bytes,1,opt,name=put_copy,json=putCopy,proto3,oneof"`
}

type StorageNodeMessage_PutCopyResponse struct {
	PutCopyResponse *StorageNodeMessage_PUTCopyResponse `protobuf:"bytes,2,opt,name=put_copy_response,json=putCopyResponse,proto3,oneof"`
}

type StorageNodeMessage_GetReplica struct {
	GetReplica *StorageNodeMessage_GETReplica `protobuf:"bytes,3,opt,name=get_replica,json=getReplica,proto3,oneof"`
}

type StorageNodeMessage_GetReplicaResponse struct {
	GetReplicaResponse *StorageNodeMessage_GETReplicaResponse `protobuf:"bytes,4,opt,name=get_replica_response,json=getReplicaResponse,proto3,oneof"`
}

func (*StorageNodeMessage_PutCopy) isStorageNodeMessage_StorageNodeMessage() {}

func (*StorageNodeMessage_PutCopyResponse) isStorageNodeMessage_StorageNodeMessage() {}

func (*StorageNodeMessage_GetReplica) isStorageNodeMessage_StorageNodeMessage() {}

func (*StorageNodeMessage_GetReplicaResponse) isStorageNodeMessage_StorageNodeMessage() {}

type StorageNodeMessage_PUTCopy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileData []byte `protobuf:"bytes,2,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`
	Checksum []byte `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (x *StorageNodeMessage_PUTCopy) Reset() {
	*x = StorageNodeMessage_PUTCopy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageNodeMessage_PUTCopy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageNodeMessage_PUTCopy) ProtoMessage() {}

func (x *StorageNodeMessage_PUTCopy) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageNodeMessage_PUTCopy.ProtoReflect.Descriptor instead.
func (*StorageNodeMessage_PUTCopy) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StorageNodeMessage_PUTCopy) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *StorageNodeMessage_PUTCopy) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

func (x *StorageNodeMessage_PUTCopy) GetChecksum() []byte {
	if x != nil {
		return x.Checksum
	}
	return nil
}

type StorageNodeMessage_PUTCopyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *StorageNodeMessage_PUTCopyResponse) Reset() {
	*x = StorageNodeMessage_PUTCopyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageNodeMessage_PUTCopyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageNodeMessage_PUTCopyResponse) ProtoMessage() {}

func (x *StorageNodeMessage_PUTCopyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageNodeMessage_PUTCopyResponse.ProtoReflect.Descriptor instead.
func (*StorageNodeMessage_PUTCopyResponse) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{0, 1}
}

func (x *StorageNodeMessage_PUTCopyResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type StorageNodeMessage_GETReplica struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *StorageNodeMessage_GETReplica) Reset() {
	*x = StorageNodeMessage_GETReplica{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageNodeMessage_GETReplica) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageNodeMessage_GETReplica) ProtoMessage() {}

func (x *StorageNodeMessage_GETReplica) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageNodeMessage_GETReplica.ProtoReflect.Descriptor instead.
func (*StorageNodeMessage_GETReplica) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{0, 2}
}

func (x *StorageNodeMessage_GETReplica) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type StorageNodeMessage_GETReplicaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileData []byte `protobuf:"bytes,2,opt,name=file_data,json=fileData,proto3" json:"file_data,omitempty"`
	Checksum []byte `protobuf:"bytes,3,opt,name=checksum,proto3" json:"checksum,omitempty"`
}

func (x *StorageNodeMessage_GETReplicaResponse) Reset() {
	*x = StorageNodeMessage_GETReplicaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_storage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageNodeMessage_GETReplicaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageNodeMessage_GETReplicaResponse) ProtoMessage() {}

func (x *StorageNodeMessage_GETReplicaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_storage_storage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageNodeMessage_GETReplicaResponse.ProtoReflect.Descriptor instead.
func (*StorageNodeMessage_GETReplicaResponse) Descriptor() ([]byte, []int) {
	return file_storage_storage_proto_rawDescGZIP(), []int{0, 3}
}

func (x *StorageNodeMessage_GETReplicaResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *StorageNodeMessage_GETReplicaResponse) GetFileData() []byte {
	if x != nil {
		return x.FileData
	}
	return nil
}

func (x *StorageNodeMessage_GETReplicaResponse) GetChecksum() []byte {
	if x != nil {
		return x.Checksum
	}
	return nil
}

var File_storage_storage_proto protoreflect.FileDescriptor

var file_storage_storage_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x04, 0x0a, 0x12, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x38,
	0x0a, 0x08, 0x70, 0x75, 0x74, 0x5f, 0x63, 0x6f, 0x70, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x55, 0x54, 0x43, 0x6f, 0x70, 0x79, 0x48, 0x00, 0x52,
	0x07, 0x70, 0x75, 0x74, 0x43, 0x6f, 0x70, 0x79, 0x12, 0x51, 0x0a, 0x11, 0x70, 0x75, 0x74, 0x5f,
	0x63, 0x6f, 0x70, 0x79, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x64,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x55, 0x54, 0x43, 0x6f, 0x70, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x75, 0x74, 0x43,
	0x6f, 0x70, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0b, 0x67,
	0x65, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x45, 0x54, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x48, 0x00, 0x52, 0x0a, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x5a,
	0x0a, 0x14, 0x67, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x47, 0x45, 0x54, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00, 0x52, 0x12, 0x67, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x5f, 0x0a, 0x07, 0x50, 0x55,
	0x54, 0x43, 0x6f, 0x70, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x1a, 0x2b, 0x0a, 0x0f, 0x50,
	0x55, 0x54, 0x43, 0x6f, 0x70, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x29, 0x0a, 0x0a, 0x47, 0x45, 0x54, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x6a, 0x0a, 0x12, 0x47, 0x45, 0x54, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x75, 0x6d, 0x42,
	0x16, 0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x2e, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_storage_proto_rawDescOnce sync.Once
	file_storage_storage_proto_rawDescData = file_storage_storage_proto_rawDesc
)

func file_storage_storage_proto_rawDescGZIP() []byte {
	file_storage_storage_proto_rawDescOnce.Do(func() {
		file_storage_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_storage_proto_rawDescData)
	})
	return file_storage_storage_proto_rawDescData
}

var file_storage_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_storage_storage_proto_goTypes = []interface{}{
	(*StorageNodeMessage)(nil),                    // 0: StorageNodeMessage
	(*StorageNodeMessage_PUTCopy)(nil),            // 1: StorageNodeMessage.PUTCopy
	(*StorageNodeMessage_PUTCopyResponse)(nil),    // 2: StorageNodeMessage.PUTCopyResponse
	(*StorageNodeMessage_GETReplica)(nil),         // 3: StorageNodeMessage.GETReplica
	(*StorageNodeMessage_GETReplicaResponse)(nil), // 4: StorageNodeMessage.GETReplicaResponse
}
var file_storage_storage_proto_depIdxs = []int32{
	1, // 0: StorageNodeMessage.put_copy:type_name -> StorageNodeMessage.PUTCopy
	2, // 1: StorageNodeMessage.put_copy_response:type_name -> StorageNodeMessage.PUTCopyResponse
	3, // 2: StorageNodeMessage.get_replica:type_name -> StorageNodeMessage.GETReplica
	4, // 3: StorageNodeMessage.get_replica_response:type_name -> StorageNodeMessage.GETReplicaResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_storage_storage_proto_init() }
func file_storage_storage_proto_init() {
	if File_storage_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageNodeMessage); i {
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
		file_storage_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageNodeMessage_PUTCopy); i {
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
		file_storage_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageNodeMessage_PUTCopyResponse); i {
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
		file_storage_storage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageNodeMessage_GETReplica); i {
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
		file_storage_storage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageNodeMessage_GETReplicaResponse); i {
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
	file_storage_storage_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*StorageNodeMessage_PutCopy)(nil),
		(*StorageNodeMessage_PutCopyResponse)(nil),
		(*StorageNodeMessage_GetReplica)(nil),
		(*StorageNodeMessage_GetReplicaResponse)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_storage_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_storage_proto_goTypes,
		DependencyIndexes: file_storage_storage_proto_depIdxs,
		MessageInfos:      file_storage_storage_proto_msgTypes,
	}.Build()
	File_storage_storage_proto = out.File
	file_storage_storage_proto_rawDesc = nil
	file_storage_storage_proto_goTypes = nil
	file_storage_storage_proto_depIdxs = nil
}

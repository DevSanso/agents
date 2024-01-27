// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: client_list.proto

package protos

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

type RedisClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Addr      string `protobuf:"bytes,2,opt,name=Addr,proto3" json:"Addr,omitempty"`
	LocalAddr string `protobuf:"bytes,3,opt,name=LocalAddr,proto3" json:"LocalAddr,omitempty"`
	FD        int64  `protobuf:"varint,4,opt,name=FD,proto3" json:"FD,omitempty"`
	Name      string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	Age       int64  `protobuf:"varint,6,opt,name=Age,proto3" json:"Age,omitempty"`
	Idle      int64  `protobuf:"varint,7,opt,name=Idle,proto3" json:"Idle,omitempty"`
	Flags     string `protobuf:"bytes,8,opt,name=Flags,proto3" json:"Flags,omitempty"`
	DB        int64  `protobuf:"varint,9,opt,name=DB,proto3" json:"DB,omitempty"`
	Sub       int64  `protobuf:"varint,10,opt,name=Sub,proto3" json:"Sub,omitempty"`
	PSub      int64  `protobuf:"varint,11,opt,name=PSub,proto3" json:"PSub,omitempty"`
	SSub      int64  `protobuf:"varint,12,opt,name=SSub,proto3" json:"SSub,omitempty"`
	Multi     int64  `protobuf:"varint,13,opt,name=Multi,proto3" json:"Multi,omitempty"`
	QBuf      string `protobuf:"bytes,14,opt,name=QBuf,proto3" json:"QBuf,omitempty"`
	QBufFree  int64  `protobuf:"varint,15,opt,name=QBufFree,proto3" json:"QBufFree,omitempty"`
	ArgvMem   int64  `protobuf:"varint,16,opt,name=ArgvMem,proto3" json:"ArgvMem,omitempty"`
	MultiMem  int64  `protobuf:"varint,17,opt,name=MultiMem,proto3" json:"MultiMem,omitempty"`
	RBS       int64  `protobuf:"varint,18,opt,name=RBS,proto3" json:"RBS,omitempty"`
	RBP       int64  `protobuf:"varint,19,opt,name=RBP,proto3" json:"RBP,omitempty"`
	OBL       int64  `protobuf:"varint,20,opt,name=OBL,proto3" json:"OBL,omitempty"`
	OLL       int64  `protobuf:"varint,21,opt,name=OLL,proto3" json:"OLL,omitempty"`
	OMem      int64  `protobuf:"varint,22,opt,name=OMem,proto3" json:"OMem,omitempty"`
	TotMem    int64  `protobuf:"varint,23,opt,name=TotMem,proto3" json:"TotMem,omitempty"`
	Events    string `protobuf:"bytes,24,opt,name=Events,proto3" json:"Events,omitempty"`
	Cmd       string `protobuf:"bytes,25,opt,name=Cmd,proto3" json:"Cmd,omitempty"`
	User      string `protobuf:"bytes,26,opt,name=User,proto3" json:"User,omitempty"`
	Redir     int64  `protobuf:"varint,27,opt,name=Redir,proto3" json:"Redir,omitempty"`
	Resp      int64  `protobuf:"varint,28,opt,name=Resp,proto3" json:"Resp,omitempty"`
	LibName   string `protobuf:"bytes,29,opt,name=LibName,proto3" json:"LibName,omitempty"`
	LibVer    string `protobuf:"bytes,30,opt,name=LibVer,proto3" json:"LibVer,omitempty"`
}

func (x *RedisClientInfo) Reset() {
	*x = RedisClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_list_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisClientInfo) ProtoMessage() {}

func (x *RedisClientInfo) ProtoReflect() protoreflect.Message {
	mi := &file_client_list_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisClientInfo.ProtoReflect.Descriptor instead.
func (*RedisClientInfo) Descriptor() ([]byte, []int) {
	return file_client_list_proto_rawDescGZIP(), []int{0}
}

func (x *RedisClientInfo) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *RedisClientInfo) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *RedisClientInfo) GetLocalAddr() string {
	if x != nil {
		return x.LocalAddr
	}
	return ""
}

func (x *RedisClientInfo) GetFD() int64 {
	if x != nil {
		return x.FD
	}
	return 0
}

func (x *RedisClientInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RedisClientInfo) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *RedisClientInfo) GetIdle() int64 {
	if x != nil {
		return x.Idle
	}
	return 0
}

func (x *RedisClientInfo) GetFlags() string {
	if x != nil {
		return x.Flags
	}
	return ""
}

func (x *RedisClientInfo) GetDB() int64 {
	if x != nil {
		return x.DB
	}
	return 0
}

func (x *RedisClientInfo) GetSub() int64 {
	if x != nil {
		return x.Sub
	}
	return 0
}

func (x *RedisClientInfo) GetPSub() int64 {
	if x != nil {
		return x.PSub
	}
	return 0
}

func (x *RedisClientInfo) GetSSub() int64 {
	if x != nil {
		return x.SSub
	}
	return 0
}

func (x *RedisClientInfo) GetMulti() int64 {
	if x != nil {
		return x.Multi
	}
	return 0
}

func (x *RedisClientInfo) GetQBuf() string {
	if x != nil {
		return x.QBuf
	}
	return ""
}

func (x *RedisClientInfo) GetQBufFree() int64 {
	if x != nil {
		return x.QBufFree
	}
	return 0
}

func (x *RedisClientInfo) GetArgvMem() int64 {
	if x != nil {
		return x.ArgvMem
	}
	return 0
}

func (x *RedisClientInfo) GetMultiMem() int64 {
	if x != nil {
		return x.MultiMem
	}
	return 0
}

func (x *RedisClientInfo) GetRBS() int64 {
	if x != nil {
		return x.RBS
	}
	return 0
}

func (x *RedisClientInfo) GetRBP() int64 {
	if x != nil {
		return x.RBP
	}
	return 0
}

func (x *RedisClientInfo) GetOBL() int64 {
	if x != nil {
		return x.OBL
	}
	return 0
}

func (x *RedisClientInfo) GetOLL() int64 {
	if x != nil {
		return x.OLL
	}
	return 0
}

func (x *RedisClientInfo) GetOMem() int64 {
	if x != nil {
		return x.OMem
	}
	return 0
}

func (x *RedisClientInfo) GetTotMem() int64 {
	if x != nil {
		return x.TotMem
	}
	return 0
}

func (x *RedisClientInfo) GetEvents() string {
	if x != nil {
		return x.Events
	}
	return ""
}

func (x *RedisClientInfo) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

func (x *RedisClientInfo) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *RedisClientInfo) GetRedir() int64 {
	if x != nil {
		return x.Redir
	}
	return 0
}

func (x *RedisClientInfo) GetResp() int64 {
	if x != nil {
		return x.Resp
	}
	return 0
}

func (x *RedisClientInfo) GetLibName() string {
	if x != nil {
		return x.LibName
	}
	return ""
}

func (x *RedisClientInfo) GetLibVer() string {
	if x != nil {
		return x.LibVer
	}
	return ""
}

type RedisClientList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clients   []*RedisClientInfo `protobuf:"bytes,1,rep,name=clients,proto3" json:"clients,omitempty"`
	UnixEpoch uint64             `protobuf:"varint,2,opt,name=unix_epoch,json=unixEpoch,proto3" json:"unix_epoch,omitempty"`
}

func (x *RedisClientList) Reset() {
	*x = RedisClientList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_list_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisClientList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisClientList) ProtoMessage() {}

func (x *RedisClientList) ProtoReflect() protoreflect.Message {
	mi := &file_client_list_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisClientList.ProtoReflect.Descriptor instead.
func (*RedisClientList) Descriptor() ([]byte, []int) {
	return file_client_list_proto_rawDescGZIP(), []int{1}
}

func (x *RedisClientList) GetClients() []*RedisClientInfo {
	if x != nil {
		return x.Clients
	}
	return nil
}

func (x *RedisClientList) GetUnixEpoch() uint64 {
	if x != nil {
		return x.UnixEpoch
	}
	return 0
}

var File_client_list_proto protoreflect.FileDescriptor

var file_client_list_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73,
	0x22, 0x87, 0x05, 0x0a, 0x0f, 0x52, 0x65, 0x64, 0x69, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x63, 0x61,
	0x6c, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x6f, 0x63,
	0x61, 0x6c, 0x41, 0x64, 0x64, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x46, 0x44, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x46, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x41, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x49, 0x64, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x49, 0x64, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x44, 0x42, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x44, 0x42, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x75, 0x62, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x53, 0x75, 0x62, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x53, 0x75, 0x62,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x53, 0x75, 0x62, 0x12, 0x12, 0x0a, 0x04,
	0x53, 0x53, 0x75, 0x62, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x53, 0x75, 0x62,
	0x12, 0x14, 0x0a, 0x05, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x51, 0x42, 0x75, 0x66, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x51, 0x42, 0x75, 0x66, 0x12, 0x1a, 0x0a, 0x08, 0x51, 0x42,
	0x75, 0x66, 0x46, 0x72, 0x65, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x51, 0x42,
	0x75, 0x66, 0x46, 0x72, 0x65, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x72, 0x67, 0x76, 0x4d, 0x65,
	0x6d, 0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x41, 0x72, 0x67, 0x76, 0x4d, 0x65, 0x6d,
	0x12, 0x1a, 0x0a, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4d, 0x65, 0x6d, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x4d, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03,
	0x52, 0x42, 0x53, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x52, 0x42, 0x53, 0x12, 0x10,
	0x0a, 0x03, 0x52, 0x42, 0x50, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x52, 0x42, 0x50,
	0x12, 0x10, 0x0a, 0x03, 0x4f, 0x42, 0x4c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x4f,
	0x42, 0x4c, 0x12, 0x10, 0x0a, 0x03, 0x4f, 0x4c, 0x4c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x4f, 0x4c, 0x4c, 0x12, 0x12, 0x0a, 0x04, 0x4f, 0x4d, 0x65, 0x6d, 0x18, 0x16, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x4f, 0x4d, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x6f, 0x74, 0x4d,
	0x65, 0x6d, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x54, 0x6f, 0x74, 0x4d, 0x65, 0x6d,
	0x12, 0x16, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x6d, 0x64, 0x18,
	0x19, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73,
	0x65, 0x72, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x52, 0x65, 0x64, 0x69, 0x72, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x52,
	0x65, 0x64, 0x69, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x65, 0x73, 0x70, 0x18, 0x1c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x04, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x4c, 0x69, 0x62, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x4c, 0x69, 0x62, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4c, 0x69, 0x62, 0x56, 0x65, 0x72, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x4c, 0x69, 0x62, 0x56, 0x65, 0x72, 0x22, 0x68, 0x0a, 0x0f, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x07, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c,
	0x2e, 0x61, 0x67, 0x6e, 0x65, 0x74, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x2e, 0x52, 0x65, 0x64,
	0x69, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x65, 0x70,
	0x6f, 0x63, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x75, 0x6e, 0x69, 0x78, 0x45,
	0x70, 0x6f, 0x63, 0x68, 0x42, 0x30, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0xaa, 0x02, 0x23, 0x49, 0x6e, 0x66, 0x6f, 0x47, 0x61, 0x74, 0x68, 0x65, 0x72, 0x48, 0x75, 0x62,
	0x2e, 0x48, 0x75, 0x62, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_list_proto_rawDescOnce sync.Once
	file_client_list_proto_rawDescData = file_client_list_proto_rawDesc
)

func file_client_list_proto_rawDescGZIP() []byte {
	file_client_list_proto_rawDescOnce.Do(func() {
		file_client_list_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_list_proto_rawDescData)
	})
	return file_client_list_proto_rawDescData
}

var file_client_list_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_client_list_proto_goTypes = []interface{}{
	(*RedisClientInfo)(nil), // 0: agnet.redis.RedisClientInfo
	(*RedisClientList)(nil), // 1: agnet.redis.RedisClientList
}
var file_client_list_proto_depIdxs = []int32{
	0, // 0: agnet.redis.RedisClientList.clients:type_name -> agnet.redis.RedisClientInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_client_list_proto_init() }
func file_client_list_proto_init() {
	if File_client_list_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_list_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisClientInfo); i {
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
		file_client_list_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisClientList); i {
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
			RawDescriptor: file_client_list_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_client_list_proto_goTypes,
		DependencyIndexes: file_client_list_proto_depIdxs,
		MessageInfos:      file_client_list_proto_msgTypes,
	}.Build()
	File_client_list_proto = out.File
	file_client_list_proto_rawDesc = nil
	file_client_list_proto_goTypes = nil
	file_client_list_proto_depIdxs = nil
}

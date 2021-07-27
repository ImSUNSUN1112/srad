// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.0
// source: sradrpc.proto

//包名

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

//注册服务
//0.1
//message RegisterRequest{
//  string servername = 1;
//  string serverstate = 2;
//}
//message RegisterResponse{
//  string msg = 1;
//}
//0.2
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName    string `protobuf:"bytes,1,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	ServiceId      string `protobuf:"bytes,2,opt,name=ServiceId,proto3" json:"ServiceId,omitempty"`
	ServiceAddress string `protobuf:"bytes,3,opt,name=ServiceAddress,proto3" json:"ServiceAddress,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sradrpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sradrpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_sradrpc_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *RegisterRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *RegisterRequest) GetServiceAddress() string {
	if x != nil {
		return x.ServiceAddress
	}
	return ""
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sradrpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sradrpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_sradrpc_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

//发现组件
type DiscoverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
}

func (x *DiscoverRequest) Reset() {
	*x = DiscoverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sradrpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverRequest) ProtoMessage() {}

func (x *DiscoverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sradrpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverRequest.ProtoReflect.Descriptor instead.
func (*DiscoverRequest) Descriptor() ([]byte, []int) {
	return file_sradrpc_proto_rawDescGZIP(), []int{2}
}

func (x *DiscoverRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

type DiscoverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg            string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
	ServiceAddress string `protobuf:"bytes,2,opt,name=ServiceAddress,proto3" json:"ServiceAddress,omitempty"`
}

func (x *DiscoverResponse) Reset() {
	*x = DiscoverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sradrpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiscoverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverResponse) ProtoMessage() {}

func (x *DiscoverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sradrpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverResponse.ProtoReflect.Descriptor instead.
func (*DiscoverResponse) Descriptor() ([]byte, []int) {
	return file_sradrpc_proto_rawDescGZIP(), []int{3}
}

func (x *DiscoverResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *DiscoverResponse) GetServiceAddress() string {
	if x != nil {
		return x.ServiceAddress
	}
	return ""
}

//心跳检测组件
type HeartbeatDetectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceName string `protobuf:"bytes,1,opt,name=ServiceName,proto3" json:"ServiceName,omitempty"`
	ServiceId   string `protobuf:"bytes,2,opt,name=ServiceId,proto3" json:"ServiceId,omitempty"`
}

func (x *HeartbeatDetectionRequest) Reset() {
	*x = HeartbeatDetectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sradrpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatDetectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatDetectionRequest) ProtoMessage() {}

func (x *HeartbeatDetectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sradrpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatDetectionRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatDetectionRequest) Descriptor() ([]byte, []int) {
	return file_sradrpc_proto_rawDescGZIP(), []int{4}
}

func (x *HeartbeatDetectionRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *HeartbeatDetectionRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

type HeartbeatDetectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (x *HeartbeatDetectionResponse) Reset() {
	*x = HeartbeatDetectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sradrpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HeartbeatDetectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatDetectionResponse) ProtoMessage() {}

func (x *HeartbeatDetectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sradrpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatDetectionResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatDetectionResponse) Descriptor() ([]byte, []int) {
	return file_sradrpc_proto_rawDescGZIP(), []int{5}
}

func (x *HeartbeatDetectionResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_sradrpc_proto protoreflect.FileDescriptor

var file_sradrpc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x72, 0x61, 0x64, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x05, 0x73, 0x73, 0x72, 0x61, 0x64, 0x22, 0x79, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x22, 0x24, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x33, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4c, 0x0a, 0x10,
	0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d,
	0x73, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5b, 0x0a, 0x19, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x1a, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x32, 0xe2, 0x01, 0x0a, 0x05, 0x52, 0x44, 0x52, 0x70,
	0x63, 0x12, 0x3d, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e,
	0x73, 0x73, 0x72, 0x61, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x73, 0x72, 0x61, 0x64, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x73,
	0x73, 0x72, 0x61, 0x64, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x73, 0x72, 0x61, 0x64, 0x2e, 0x44, 0x69, 0x73,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5b, 0x0a, 0x12, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x44, 0x65, 0x74, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x73, 0x73, 0x72, 0x61, 0x64, 0x2e, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x73, 0x73, 0x72, 0x61, 0x64, 0x2e,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02,
	0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sradrpc_proto_rawDescOnce sync.Once
	file_sradrpc_proto_rawDescData = file_sradrpc_proto_rawDesc
)

func file_sradrpc_proto_rawDescGZIP() []byte {
	file_sradrpc_proto_rawDescOnce.Do(func() {
		file_sradrpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_sradrpc_proto_rawDescData)
	})
	return file_sradrpc_proto_rawDescData
}

var file_sradrpc_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sradrpc_proto_goTypes = []interface{}{
	(*RegisterRequest)(nil),            // 0: ssrad.RegisterRequest
	(*RegisterResponse)(nil),           // 1: ssrad.RegisterResponse
	(*DiscoverRequest)(nil),            // 2: ssrad.DiscoverRequest
	(*DiscoverResponse)(nil),           // 3: ssrad.DiscoverResponse
	(*HeartbeatDetectionRequest)(nil),  // 4: ssrad.HeartbeatDetectionRequest
	(*HeartbeatDetectionResponse)(nil), // 5: ssrad.HeartbeatDetectionResponse
}
var file_sradrpc_proto_depIdxs = []int32{
	0, // 0: ssrad.RDRpc.Register:input_type -> ssrad.RegisterRequest
	2, // 1: ssrad.RDRpc.Discover:input_type -> ssrad.DiscoverRequest
	4, // 2: ssrad.RDRpc.HeartbeatDetection:input_type -> ssrad.HeartbeatDetectionRequest
	1, // 3: ssrad.RDRpc.Register:output_type -> ssrad.RegisterResponse
	3, // 4: ssrad.RDRpc.Discover:output_type -> ssrad.DiscoverResponse
	5, // 5: ssrad.RDRpc.HeartbeatDetection:output_type -> ssrad.HeartbeatDetectionResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sradrpc_proto_init() }
func file_sradrpc_proto_init() {
	if File_sradrpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sradrpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_sradrpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_sradrpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverRequest); i {
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
		file_sradrpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiscoverResponse); i {
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
		file_sradrpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatDetectionRequest); i {
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
		file_sradrpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HeartbeatDetectionResponse); i {
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
			RawDescriptor: file_sradrpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sradrpc_proto_goTypes,
		DependencyIndexes: file_sradrpc_proto_depIdxs,
		MessageInfos:      file_sradrpc_proto_msgTypes,
	}.Build()
	File_sradrpc_proto = out.File
	file_sradrpc_proto_rawDesc = nil
	file_sradrpc_proto_goTypes = nil
	file_sradrpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RDRpcClient is the client API for RDRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RDRpcClient interface {
	//注册
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	//发现
	Discover(ctx context.Context, in *DiscoverRequest, opts ...grpc.CallOption) (*DiscoverResponse, error)
	//心跳检测
	HeartbeatDetection(ctx context.Context, in *HeartbeatDetectionRequest, opts ...grpc.CallOption) (*HeartbeatDetectionResponse, error)
}

type rDRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRDRpcClient(cc grpc.ClientConnInterface) RDRpcClient {
	return &rDRpcClient{cc}
}

func (c *rDRpcClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/ssrad.RDRpc/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rDRpcClient) Discover(ctx context.Context, in *DiscoverRequest, opts ...grpc.CallOption) (*DiscoverResponse, error) {
	out := new(DiscoverResponse)
	err := c.cc.Invoke(ctx, "/ssrad.RDRpc/Discover", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rDRpcClient) HeartbeatDetection(ctx context.Context, in *HeartbeatDetectionRequest, opts ...grpc.CallOption) (*HeartbeatDetectionResponse, error) {
	out := new(HeartbeatDetectionResponse)
	err := c.cc.Invoke(ctx, "/ssrad.RDRpc/HeartbeatDetection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RDRpcServer is the server API for RDRpc service.
type RDRpcServer interface {
	//注册
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	//发现
	Discover(context.Context, *DiscoverRequest) (*DiscoverResponse, error)
	//心跳检测
	HeartbeatDetection(context.Context, *HeartbeatDetectionRequest) (*HeartbeatDetectionResponse, error)
}

// UnimplementedRDRpcServer can be embedded to have forward compatible implementations.
type UnimplementedRDRpcServer struct {
}

func (*UnimplementedRDRpcServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (*UnimplementedRDRpcServer) Discover(context.Context, *DiscoverRequest) (*DiscoverResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Discover not implemented")
}
func (*UnimplementedRDRpcServer) HeartbeatDetection(context.Context, *HeartbeatDetectionRequest) (*HeartbeatDetectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartbeatDetection not implemented")
}

func RegisterRDRpcServer(s *grpc.Server, srv RDRpcServer) {
	s.RegisterService(&_RDRpc_serviceDesc, srv)
}

func _RDRpc_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDRpcServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ssrad.RDRpc/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDRpcServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RDRpc_Discover_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDRpcServer).Discover(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ssrad.RDRpc/Discover",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDRpcServer).Discover(ctx, req.(*DiscoverRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RDRpc_HeartbeatDetection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatDetectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RDRpcServer).HeartbeatDetection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ssrad.RDRpc/HeartbeatDetection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RDRpcServer).HeartbeatDetection(ctx, req.(*HeartbeatDetectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RDRpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ssrad.RDRpc",
	HandlerType: (*RDRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RDRpc_Register_Handler,
		},
		{
			MethodName: "Discover",
			Handler:    _RDRpc_Discover_Handler,
		},
		{
			MethodName: "HeartbeatDetection",
			Handler:    _RDRpc_HeartbeatDetection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sradrpc.proto",
}
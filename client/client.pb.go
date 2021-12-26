// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: client.proto

package client

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	common "kubectl-cli/common"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetPodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *GetPodsRequest) Reset() {
	*x = GetPodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPodsRequest) ProtoMessage() {}

func (x *GetPodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPodsRequest.ProtoReflect.Descriptor instead.
func (*GetPodsRequest) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *GetPodsRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

// The response message containing the greetings
type GetPodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resp *common.CommonResponse `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
	Info []*PodInfo             `protobuf:"bytes,2,rep,name=info,proto3" json:"info,omitempty"`
}

func (x *GetPodsResponse) Reset() {
	*x = GetPodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPodsResponse) ProtoMessage() {}

func (x *GetPodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPodsResponse.ProtoReflect.Descriptor instead.
func (*GetPodsResponse) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

func (x *GetPodsResponse) GetResp() *common.CommonResponse {
	if x != nil {
		return x.Resp
	}
	return nil
}

func (x *GetPodsResponse) GetInfo() []*PodInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type PodInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Namespace string `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TotalPods int32  `protobuf:"varint,3,opt,name=totalPods,proto3" json:"totalPods,omitempty"`
	ReadyPods int32  `protobuf:"varint,4,opt,name=readyPods,proto3" json:"readyPods,omitempty"`
	Status    string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Restarts  int32  `protobuf:"varint,6,opt,name=restarts,proto3" json:"restarts,omitempty"`
	Age       int32  `protobuf:"varint,7,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *PodInfo) Reset() {
	*x = PodInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodInfo) ProtoMessage() {}

func (x *PodInfo) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodInfo.ProtoReflect.Descriptor instead.
func (*PodInfo) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{2}
}

func (x *PodInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PodInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PodInfo) GetTotalPods() int32 {
	if x != nil {
		return x.TotalPods
	}
	return 0
}

func (x *PodInfo) GetReadyPods() int32 {
	if x != nil {
		return x.ReadyPods
	}
	return 0
}

func (x *PodInfo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PodInfo) GetRestarts() int32 {
	if x != nil {
		return x.Restarts
	}
	return 0
}

func (x *PodInfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0x54, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x23, 0x0a, 0x04, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04,
	0x72, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x6f, 0x64, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x72, 0x65, 0x61, 0x64, 0x79, 0x50, 0x6f, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x72, 0x65, 0x61, 0x64, 0x79, 0x50, 0x6f, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x73, 0x74, 0x61, 0x72, 0x74, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61,
	0x67, 0x65, 0x32, 0x3f, 0x0a, 0x0d, 0x4b, 0x75, 0x62, 0x65, 0x63, 0x74, 0x6c, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x73, 0x12, 0x0f,
	0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x10, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x19, 0x5a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x72, 0x65, 0x70, 0x6f,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x3b, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_client_proto_goTypes = []interface{}{
	(*GetPodsRequest)(nil),        // 0: GetPodsRequest
	(*GetPodsResponse)(nil),       // 1: GetPodsResponse
	(*PodInfo)(nil),               // 2: PodInfo
	(*common.CommonResponse)(nil), // 3: CommonResponse
}
var file_client_proto_depIdxs = []int32{
	3, // 0: GetPodsResponse.resp:type_name -> CommonResponse
	2, // 1: GetPodsResponse.info:type_name -> PodInfo
	0, // 2: KubectlClient.GetPods:input_type -> GetPodsRequest
	1, // 3: KubectlClient.GetPods:output_type -> GetPodsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPodsRequest); i {
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
		file_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPodsResponse); i {
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
		file_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodInfo); i {
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
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}

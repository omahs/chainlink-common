// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: chain_reader.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetLatestValueRequest has arguments for [github.com/smartcontractkit/chainlink-common/pkg/types.ChainReader.GetLatestValue].
type GetLatestValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContractName string          `protobuf:"bytes,1,opt,name=contractName,proto3" json:"contractName,omitempty"`
	Method       string          `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Params       *VersionedBytes `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *GetLatestValueRequest) Reset() {
	*x = GetLatestValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_reader_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestValueRequest) ProtoMessage() {}

func (x *GetLatestValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chain_reader_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestValueRequest.ProtoReflect.Descriptor instead.
func (*GetLatestValueRequest) Descriptor() ([]byte, []int) {
	return file_chain_reader_proto_rawDescGZIP(), []int{0}
}

func (x *GetLatestValueRequest) GetContractName() string {
	if x != nil {
		return x.ContractName
	}
	return ""
}

func (x *GetLatestValueRequest) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *GetLatestValueRequest) GetParams() *VersionedBytes {
	if x != nil {
		return x.Params
	}
	return nil
}

// GetLatestValueReply has return arguments for [github.com/smartcontractkit/chainlink-common/pkg/types.ChainReader.GetLatestValue].
type GetLatestValueReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetVal *VersionedBytes `protobuf:"bytes,1,opt,name=retVal,proto3" json:"retVal,omitempty"`
}

func (x *GetLatestValueReply) Reset() {
	*x = GetLatestValueReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_reader_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLatestValueReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLatestValueReply) ProtoMessage() {}

func (x *GetLatestValueReply) ProtoReflect() protoreflect.Message {
	mi := &file_chain_reader_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLatestValueReply.ProtoReflect.Descriptor instead.
func (*GetLatestValueReply) Descriptor() ([]byte, []int) {
	return file_chain_reader_proto_rawDescGZIP(), []int{1}
}

func (x *GetLatestValueReply) GetRetVal() *VersionedBytes {
	if x != nil {
		return x.RetVal
	}
	return nil
}

type BindRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bindings []*BoundContract `protobuf:"bytes,1,rep,name=bindings,proto3" json:"bindings,omitempty"`
}

func (x *BindRequest) Reset() {
	*x = BindRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_reader_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BindRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BindRequest) ProtoMessage() {}

func (x *BindRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chain_reader_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BindRequest.ProtoReflect.Descriptor instead.
func (*BindRequest) Descriptor() ([]byte, []int) {
	return file_chain_reader_proto_rawDescGZIP(), []int{2}
}

func (x *BindRequest) GetBindings() []*BoundContract {
	if x != nil {
		return x.Bindings
	}
	return nil
}

// BoundContract represents a [github.com/smartcontractkit/chainlink-common/pkg/types.BoundContract].
type BoundContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address    string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Confidence int32  `protobuf:"varint,3,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *BoundContract) Reset() {
	*x = BoundContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chain_reader_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoundContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoundContract) ProtoMessage() {}

func (x *BoundContract) ProtoReflect() protoreflect.Message {
	mi := &file_chain_reader_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoundContract.ProtoReflect.Descriptor instead.
func (*BoundContract) Descriptor() ([]byte, []int) {
	return file_chain_reader_proto_rawDescGZIP(), []int{3}
}

func (x *BoundContract) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *BoundContract) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BoundContract) GetConfidence() int32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

var File_chain_reader_proto protoreflect.FileDescriptor

var file_chain_reader_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x6f, 0x6f, 0x70, 0x1a, 0x0b, 0x63, 0x6f, 0x64, 0x65,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2c, 0x0a, 0x06, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6c, 0x6f, 0x6f,
	0x70, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x42, 0x79, 0x74, 0x65, 0x73,
	0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x43, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x2c, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64,
	0x42, 0x79, 0x74, 0x65, 0x73, 0x52, 0x06, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x22, 0x3e, 0x0a,
	0x0b, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x08,
	0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x52, 0x08, 0x62, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x5d, 0x0a,
	0x0d, 0x42, 0x6f, 0x75, 0x6e, 0x64, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x32, 0x8e, 0x01, 0x0a,
	0x0b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1b,
	0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6c, 0x6f,
	0x6f, 0x70, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x33, 0x0a, 0x04, 0x42, 0x69, 0x6e, 0x64,
	0x12, 0x11, 0x2e, 0x6c, 0x6f, 0x6f, 0x70, 0x2e, 0x42, 0x69, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x43, 0x5a,
	0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6d, 0x61, 0x72,
	0x74, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6b, 0x69, 0x74, 0x2f, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x6c, 0x6f, 0x6f, 0x70, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chain_reader_proto_rawDescOnce sync.Once
	file_chain_reader_proto_rawDescData = file_chain_reader_proto_rawDesc
)

func file_chain_reader_proto_rawDescGZIP() []byte {
	file_chain_reader_proto_rawDescOnce.Do(func() {
		file_chain_reader_proto_rawDescData = protoimpl.X.CompressGZIP(file_chain_reader_proto_rawDescData)
	})
	return file_chain_reader_proto_rawDescData
}

var file_chain_reader_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_chain_reader_proto_goTypes = []interface{}{
	(*GetLatestValueRequest)(nil), // 0: loop.GetLatestValueRequest
	(*GetLatestValueReply)(nil),   // 1: loop.GetLatestValueReply
	(*BindRequest)(nil),           // 2: loop.BindRequest
	(*BoundContract)(nil),         // 3: loop.BoundContract
	(*VersionedBytes)(nil),        // 4: loop.VersionedBytes
	(*emptypb.Empty)(nil),         // 5: google.protobuf.Empty
}
var file_chain_reader_proto_depIdxs = []int32{
	4, // 0: loop.GetLatestValueRequest.params:type_name -> loop.VersionedBytes
	4, // 1: loop.GetLatestValueReply.retVal:type_name -> loop.VersionedBytes
	3, // 2: loop.BindRequest.bindings:type_name -> loop.BoundContract
	0, // 3: loop.ChainReader.GetLatestValue:input_type -> loop.GetLatestValueRequest
	2, // 4: loop.ChainReader.Bind:input_type -> loop.BindRequest
	1, // 5: loop.ChainReader.GetLatestValue:output_type -> loop.GetLatestValueReply
	5, // 6: loop.ChainReader.Bind:output_type -> google.protobuf.Empty
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_chain_reader_proto_init() }
func file_chain_reader_proto_init() {
	if File_chain_reader_proto != nil {
		return
	}
	file_codec_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chain_reader_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestValueRequest); i {
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
		file_chain_reader_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLatestValueReply); i {
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
		file_chain_reader_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BindRequest); i {
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
		file_chain_reader_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoundContract); i {
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
			RawDescriptor: file_chain_reader_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chain_reader_proto_goTypes,
		DependencyIndexes: file_chain_reader_proto_depIdxs,
		MessageInfos:      file_chain_reader_proto_msgTypes,
	}.Build()
	File_chain_reader_proto = out.File
	file_chain_reader_proto_rawDesc = nil
	file_chain_reader_proto_goTypes = nil
	file_chain_reader_proto_depIdxs = nil
}

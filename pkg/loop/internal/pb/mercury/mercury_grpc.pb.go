// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: mercury.proto

// the generate file in this dir specifies the import path of the relative proto files
// TODO: is there a better, more explicit to reference relayer.proto, which is really in the parent dir?
// naive attempt: import "github.com/smartcontractkit/chainlink-common/pkg/loop/internal/pb/relayer.proto"; and `../relayer.proto` did not work

package mercury_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	MercuryAdapter_NewMercuryV1Factory_FullMethodName = "/loop.internal.pb.mercury.MercuryAdapter/NewMercuryV1Factory"
	MercuryAdapter_NewMercuryV2Factory_FullMethodName = "/loop.internal.pb.mercury.MercuryAdapter/NewMercuryV2Factory"
	MercuryAdapter_NewMercuryV3Factory_FullMethodName = "/loop.internal.pb.mercury.MercuryAdapter/NewMercuryV3Factory"
)

// MercuryAdapterClient is the client API for MercuryAdapter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MercuryAdapterClient interface {
	NewMercuryV1Factory(ctx context.Context, in *NewMercuryV1FactoryRequest, opts ...grpc.CallOption) (*NewMercuryV1FactoryReply, error)
	NewMercuryV2Factory(ctx context.Context, in *NewMercuryV2FactoryRequest, opts ...grpc.CallOption) (*NewMercuryV2FactoryReply, error)
	NewMercuryV3Factory(ctx context.Context, in *NewMercuryV3FactoryRequest, opts ...grpc.CallOption) (*NewMercuryV3FactoryReply, error)
}

type mercuryAdapterClient struct {
	cc grpc.ClientConnInterface
}

func NewMercuryAdapterClient(cc grpc.ClientConnInterface) MercuryAdapterClient {
	return &mercuryAdapterClient{cc}
}

func (c *mercuryAdapterClient) NewMercuryV1Factory(ctx context.Context, in *NewMercuryV1FactoryRequest, opts ...grpc.CallOption) (*NewMercuryV1FactoryReply, error) {
	out := new(NewMercuryV1FactoryReply)
	err := c.cc.Invoke(ctx, MercuryAdapter_NewMercuryV1Factory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mercuryAdapterClient) NewMercuryV2Factory(ctx context.Context, in *NewMercuryV2FactoryRequest, opts ...grpc.CallOption) (*NewMercuryV2FactoryReply, error) {
	out := new(NewMercuryV2FactoryReply)
	err := c.cc.Invoke(ctx, MercuryAdapter_NewMercuryV2Factory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mercuryAdapterClient) NewMercuryV3Factory(ctx context.Context, in *NewMercuryV3FactoryRequest, opts ...grpc.CallOption) (*NewMercuryV3FactoryReply, error) {
	out := new(NewMercuryV3FactoryReply)
	err := c.cc.Invoke(ctx, MercuryAdapter_NewMercuryV3Factory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MercuryAdapterServer is the server API for MercuryAdapter service.
// All implementations must embed UnimplementedMercuryAdapterServer
// for forward compatibility
type MercuryAdapterServer interface {
	NewMercuryV1Factory(context.Context, *NewMercuryV1FactoryRequest) (*NewMercuryV1FactoryReply, error)
	NewMercuryV2Factory(context.Context, *NewMercuryV2FactoryRequest) (*NewMercuryV2FactoryReply, error)
	NewMercuryV3Factory(context.Context, *NewMercuryV3FactoryRequest) (*NewMercuryV3FactoryReply, error)
	mustEmbedUnimplementedMercuryAdapterServer()
}

// UnimplementedMercuryAdapterServer must be embedded to have forward compatible implementations.
type UnimplementedMercuryAdapterServer struct {
}

func (UnimplementedMercuryAdapterServer) NewMercuryV1Factory(context.Context, *NewMercuryV1FactoryRequest) (*NewMercuryV1FactoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewMercuryV1Factory not implemented")
}
func (UnimplementedMercuryAdapterServer) NewMercuryV2Factory(context.Context, *NewMercuryV2FactoryRequest) (*NewMercuryV2FactoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewMercuryV2Factory not implemented")
}
func (UnimplementedMercuryAdapterServer) NewMercuryV3Factory(context.Context, *NewMercuryV3FactoryRequest) (*NewMercuryV3FactoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewMercuryV3Factory not implemented")
}
func (UnimplementedMercuryAdapterServer) mustEmbedUnimplementedMercuryAdapterServer() {}

// UnsafeMercuryAdapterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MercuryAdapterServer will
// result in compilation errors.
type UnsafeMercuryAdapterServer interface {
	mustEmbedUnimplementedMercuryAdapterServer()
}

func RegisterMercuryAdapterServer(s grpc.ServiceRegistrar, srv MercuryAdapterServer) {
	s.RegisterService(&MercuryAdapter_ServiceDesc, srv)
}

func _MercuryAdapter_NewMercuryV1Factory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMercuryV1FactoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercuryAdapterServer).NewMercuryV1Factory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MercuryAdapter_NewMercuryV1Factory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercuryAdapterServer).NewMercuryV1Factory(ctx, req.(*NewMercuryV1FactoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MercuryAdapter_NewMercuryV2Factory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMercuryV2FactoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercuryAdapterServer).NewMercuryV2Factory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MercuryAdapter_NewMercuryV2Factory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercuryAdapterServer).NewMercuryV2Factory(ctx, req.(*NewMercuryV2FactoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MercuryAdapter_NewMercuryV3Factory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewMercuryV3FactoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercuryAdapterServer).NewMercuryV3Factory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MercuryAdapter_NewMercuryV3Factory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercuryAdapterServer).NewMercuryV3Factory(ctx, req.(*NewMercuryV3FactoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MercuryAdapter_ServiceDesc is the grpc.ServiceDesc for MercuryAdapter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MercuryAdapter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.internal.pb.mercury.MercuryAdapter",
	HandlerType: (*MercuryAdapterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewMercuryV1Factory",
			Handler:    _MercuryAdapter_NewMercuryV1Factory_Handler,
		},
		{
			MethodName: "NewMercuryV2Factory",
			Handler:    _MercuryAdapter_NewMercuryV2Factory_Handler,
		},
		{
			MethodName: "NewMercuryV3Factory",
			Handler:    _MercuryAdapter_NewMercuryV3Factory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mercury.proto",
}

const (
	OnchainConfigCodec_Encode_FullMethodName = "/loop.internal.pb.mercury.OnchainConfigCodec/Encode"
	OnchainConfigCodec_Decode_FullMethodName = "/loop.internal.pb.mercury.OnchainConfigCodec/Decode"
)

// OnchainConfigCodecClient is the client API for OnchainConfigCodec service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OnchainConfigCodecClient interface {
	Encode(ctx context.Context, in *EncodeOnchainConfigRequest, opts ...grpc.CallOption) (*EncodeOnchainConfigReply, error)
	Decode(ctx context.Context, in *DecodeOnchainConfigRequest, opts ...grpc.CallOption) (*DecodeOnchainConfigReply, error)
}

type onchainConfigCodecClient struct {
	cc grpc.ClientConnInterface
}

func NewOnchainConfigCodecClient(cc grpc.ClientConnInterface) OnchainConfigCodecClient {
	return &onchainConfigCodecClient{cc}
}

func (c *onchainConfigCodecClient) Encode(ctx context.Context, in *EncodeOnchainConfigRequest, opts ...grpc.CallOption) (*EncodeOnchainConfigReply, error) {
	out := new(EncodeOnchainConfigReply)
	err := c.cc.Invoke(ctx, OnchainConfigCodec_Encode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onchainConfigCodecClient) Decode(ctx context.Context, in *DecodeOnchainConfigRequest, opts ...grpc.CallOption) (*DecodeOnchainConfigReply, error) {
	out := new(DecodeOnchainConfigReply)
	err := c.cc.Invoke(ctx, OnchainConfigCodec_Decode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnchainConfigCodecServer is the server API for OnchainConfigCodec service.
// All implementations must embed UnimplementedOnchainConfigCodecServer
// for forward compatibility
type OnchainConfigCodecServer interface {
	Encode(context.Context, *EncodeOnchainConfigRequest) (*EncodeOnchainConfigReply, error)
	Decode(context.Context, *DecodeOnchainConfigRequest) (*DecodeOnchainConfigReply, error)
	mustEmbedUnimplementedOnchainConfigCodecServer()
}

// UnimplementedOnchainConfigCodecServer must be embedded to have forward compatible implementations.
type UnimplementedOnchainConfigCodecServer struct {
}

func (UnimplementedOnchainConfigCodecServer) Encode(context.Context, *EncodeOnchainConfigRequest) (*EncodeOnchainConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encode not implemented")
}
func (UnimplementedOnchainConfigCodecServer) Decode(context.Context, *DecodeOnchainConfigRequest) (*DecodeOnchainConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decode not implemented")
}
func (UnimplementedOnchainConfigCodecServer) mustEmbedUnimplementedOnchainConfigCodecServer() {}

// UnsafeOnchainConfigCodecServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OnchainConfigCodecServer will
// result in compilation errors.
type UnsafeOnchainConfigCodecServer interface {
	mustEmbedUnimplementedOnchainConfigCodecServer()
}

func RegisterOnchainConfigCodecServer(s grpc.ServiceRegistrar, srv OnchainConfigCodecServer) {
	s.RegisterService(&OnchainConfigCodec_ServiceDesc, srv)
}

func _OnchainConfigCodec_Encode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EncodeOnchainConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnchainConfigCodecServer).Encode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnchainConfigCodec_Encode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnchainConfigCodecServer).Encode(ctx, req.(*EncodeOnchainConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OnchainConfigCodec_Decode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeOnchainConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnchainConfigCodecServer).Decode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OnchainConfigCodec_Decode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnchainConfigCodecServer).Decode(ctx, req.(*DecodeOnchainConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OnchainConfigCodec_ServiceDesc is the grpc.ServiceDesc for OnchainConfigCodec service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OnchainConfigCodec_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.internal.pb.mercury.OnchainConfigCodec",
	HandlerType: (*OnchainConfigCodecServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encode",
			Handler:    _OnchainConfigCodec_Encode_Handler,
		},
		{
			MethodName: "Decode",
			Handler:    _OnchainConfigCodec_Decode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mercury.proto",
}

const (
	ReportCodecV1_BuildReport_FullMethodName               = "/loop.internal.pb.mercury.ReportCodecV1/BuildReport"
	ReportCodecV1_MaxReportLength_FullMethodName           = "/loop.internal.pb.mercury.ReportCodecV1/MaxReportLength"
	ReportCodecV1_CurrentBlockNumFromReport_FullMethodName = "/loop.internal.pb.mercury.ReportCodecV1/CurrentBlockNumFromReport"
)

// ReportCodecV1Client is the client API for ReportCodecV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportCodecV1Client interface {
	BuildReport(ctx context.Context, in *BuildReportRequestV1, opts ...grpc.CallOption) (*BuildReportReplyV1, error)
	MaxReportLength(ctx context.Context, in *MaxReportLengthRequest, opts ...grpc.CallOption) (*MaxReportLengthReply, error)
	CurrentBlockNumFromReport(ctx context.Context, in *CurrentBlockNumFromReportRequest, opts ...grpc.CallOption) (*CurrentBlockNumFromReportResponse, error)
}

type reportCodecV1Client struct {
	cc grpc.ClientConnInterface
}

func NewReportCodecV1Client(cc grpc.ClientConnInterface) ReportCodecV1Client {
	return &reportCodecV1Client{cc}
}

func (c *reportCodecV1Client) BuildReport(ctx context.Context, in *BuildReportRequestV1, opts ...grpc.CallOption) (*BuildReportReplyV1, error) {
	out := new(BuildReportReplyV1)
	err := c.cc.Invoke(ctx, ReportCodecV1_BuildReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportCodecV1Client) MaxReportLength(ctx context.Context, in *MaxReportLengthRequest, opts ...grpc.CallOption) (*MaxReportLengthReply, error) {
	out := new(MaxReportLengthReply)
	err := c.cc.Invoke(ctx, ReportCodecV1_MaxReportLength_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportCodecV1Client) CurrentBlockNumFromReport(ctx context.Context, in *CurrentBlockNumFromReportRequest, opts ...grpc.CallOption) (*CurrentBlockNumFromReportResponse, error) {
	out := new(CurrentBlockNumFromReportResponse)
	err := c.cc.Invoke(ctx, ReportCodecV1_CurrentBlockNumFromReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportCodecV1Server is the server API for ReportCodecV1 service.
// All implementations must embed UnimplementedReportCodecV1Server
// for forward compatibility
type ReportCodecV1Server interface {
	BuildReport(context.Context, *BuildReportRequestV1) (*BuildReportReplyV1, error)
	MaxReportLength(context.Context, *MaxReportLengthRequest) (*MaxReportLengthReply, error)
	CurrentBlockNumFromReport(context.Context, *CurrentBlockNumFromReportRequest) (*CurrentBlockNumFromReportResponse, error)
	mustEmbedUnimplementedReportCodecV1Server()
}

// UnimplementedReportCodecV1Server must be embedded to have forward compatible implementations.
type UnimplementedReportCodecV1Server struct {
}

func (UnimplementedReportCodecV1Server) BuildReport(context.Context, *BuildReportRequestV1) (*BuildReportReplyV1, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildReport not implemented")
}
func (UnimplementedReportCodecV1Server) MaxReportLength(context.Context, *MaxReportLengthRequest) (*MaxReportLengthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaxReportLength not implemented")
}
func (UnimplementedReportCodecV1Server) CurrentBlockNumFromReport(context.Context, *CurrentBlockNumFromReportRequest) (*CurrentBlockNumFromReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentBlockNumFromReport not implemented")
}
func (UnimplementedReportCodecV1Server) mustEmbedUnimplementedReportCodecV1Server() {}

// UnsafeReportCodecV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportCodecV1Server will
// result in compilation errors.
type UnsafeReportCodecV1Server interface {
	mustEmbedUnimplementedReportCodecV1Server()
}

func RegisterReportCodecV1Server(s grpc.ServiceRegistrar, srv ReportCodecV1Server) {
	s.RegisterService(&ReportCodecV1_ServiceDesc, srv)
}

func _ReportCodecV1_BuildReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildReportRequestV1)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecV1Server).BuildReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportCodecV1_BuildReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecV1Server).BuildReport(ctx, req.(*BuildReportRequestV1))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportCodecV1_MaxReportLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaxReportLengthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecV1Server).MaxReportLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportCodecV1_MaxReportLength_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecV1Server).MaxReportLength(ctx, req.(*MaxReportLengthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportCodecV1_CurrentBlockNumFromReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentBlockNumFromReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecV1Server).CurrentBlockNumFromReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportCodecV1_CurrentBlockNumFromReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecV1Server).CurrentBlockNumFromReport(ctx, req.(*CurrentBlockNumFromReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportCodecV1_ServiceDesc is the grpc.ServiceDesc for ReportCodecV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportCodecV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.internal.pb.mercury.ReportCodecV1",
	HandlerType: (*ReportCodecV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildReport",
			Handler:    _ReportCodecV1_BuildReport_Handler,
		},
		{
			MethodName: "MaxReportLength",
			Handler:    _ReportCodecV1_MaxReportLength_Handler,
		},
		{
			MethodName: "CurrentBlockNumFromReport",
			Handler:    _ReportCodecV1_CurrentBlockNumFromReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mercury.proto",
}

const (
	ReportCodecV2_BuildReport_FullMethodName                    = "/loop.internal.pb.mercury.ReportCodecV2/BuildReport"
	ReportCodecV2_MaxReportLength_FullMethodName                = "/loop.internal.pb.mercury.ReportCodecV2/MaxReportLength"
	ReportCodecV2_ObservationTimestampFromReport_FullMethodName = "/loop.internal.pb.mercury.ReportCodecV2/ObservationTimestampFromReport"
)

// ReportCodecV2Client is the client API for ReportCodecV2 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportCodecV2Client interface {
	BuildReport(ctx context.Context, in *BuildReportRequestV2, opts ...grpc.CallOption) (*BuildReportReplyV2, error)
	MaxReportLength(ctx context.Context, in *MaxReportLengthRequest, opts ...grpc.CallOption) (*MaxReportLengthReply, error)
	ObservationTimestampFromReport(ctx context.Context, in *ObservationTimestampFromReportRequest, opts ...grpc.CallOption) (*ObservationTimestampFromReportReply, error)
}

type reportCodecV2Client struct {
	cc grpc.ClientConnInterface
}

func NewReportCodecV2Client(cc grpc.ClientConnInterface) ReportCodecV2Client {
	return &reportCodecV2Client{cc}
}

func (c *reportCodecV2Client) BuildReport(ctx context.Context, in *BuildReportRequestV2, opts ...grpc.CallOption) (*BuildReportReplyV2, error) {
	out := new(BuildReportReplyV2)
	err := c.cc.Invoke(ctx, ReportCodecV2_BuildReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportCodecV2Client) MaxReportLength(ctx context.Context, in *MaxReportLengthRequest, opts ...grpc.CallOption) (*MaxReportLengthReply, error) {
	out := new(MaxReportLengthReply)
	err := c.cc.Invoke(ctx, ReportCodecV2_MaxReportLength_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportCodecV2Client) ObservationTimestampFromReport(ctx context.Context, in *ObservationTimestampFromReportRequest, opts ...grpc.CallOption) (*ObservationTimestampFromReportReply, error) {
	out := new(ObservationTimestampFromReportReply)
	err := c.cc.Invoke(ctx, ReportCodecV2_ObservationTimestampFromReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportCodecV2Server is the server API for ReportCodecV2 service.
// All implementations must embed UnimplementedReportCodecV2Server
// for forward compatibility
type ReportCodecV2Server interface {
	BuildReport(context.Context, *BuildReportRequestV2) (*BuildReportReplyV2, error)
	MaxReportLength(context.Context, *MaxReportLengthRequest) (*MaxReportLengthReply, error)
	ObservationTimestampFromReport(context.Context, *ObservationTimestampFromReportRequest) (*ObservationTimestampFromReportReply, error)
	mustEmbedUnimplementedReportCodecV2Server()
}

// UnimplementedReportCodecV2Server must be embedded to have forward compatible implementations.
type UnimplementedReportCodecV2Server struct {
}

func (UnimplementedReportCodecV2Server) BuildReport(context.Context, *BuildReportRequestV2) (*BuildReportReplyV2, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildReport not implemented")
}
func (UnimplementedReportCodecV2Server) MaxReportLength(context.Context, *MaxReportLengthRequest) (*MaxReportLengthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaxReportLength not implemented")
}
func (UnimplementedReportCodecV2Server) ObservationTimestampFromReport(context.Context, *ObservationTimestampFromReportRequest) (*ObservationTimestampFromReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ObservationTimestampFromReport not implemented")
}
func (UnimplementedReportCodecV2Server) mustEmbedUnimplementedReportCodecV2Server() {}

// UnsafeReportCodecV2Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportCodecV2Server will
// result in compilation errors.
type UnsafeReportCodecV2Server interface {
	mustEmbedUnimplementedReportCodecV2Server()
}

func RegisterReportCodecV2Server(s grpc.ServiceRegistrar, srv ReportCodecV2Server) {
	s.RegisterService(&ReportCodecV2_ServiceDesc, srv)
}

func _ReportCodecV2_BuildReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildReportRequestV2)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecV2Server).BuildReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportCodecV2_BuildReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecV2Server).BuildReport(ctx, req.(*BuildReportRequestV2))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportCodecV2_MaxReportLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaxReportLengthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecV2Server).MaxReportLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportCodecV2_MaxReportLength_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecV2Server).MaxReportLength(ctx, req.(*MaxReportLengthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportCodecV2_ObservationTimestampFromReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObservationTimestampFromReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecV2Server).ObservationTimestampFromReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportCodecV2_ObservationTimestampFromReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecV2Server).ObservationTimestampFromReport(ctx, req.(*ObservationTimestampFromReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportCodecV2_ServiceDesc is the grpc.ServiceDesc for ReportCodecV2 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportCodecV2_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.internal.pb.mercury.ReportCodecV2",
	HandlerType: (*ReportCodecV2Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildReport",
			Handler:    _ReportCodecV2_BuildReport_Handler,
		},
		{
			MethodName: "MaxReportLength",
			Handler:    _ReportCodecV2_MaxReportLength_Handler,
		},
		{
			MethodName: "ObservationTimestampFromReport",
			Handler:    _ReportCodecV2_ObservationTimestampFromReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mercury.proto",
}

const (
	MercuryServerFetcher_FetchInitialMaxFinalizedBlockNumber_FullMethodName = "/loop.internal.pb.mercury.MercuryServerFetcher/FetchInitialMaxFinalizedBlockNumber"
	MercuryServerFetcher_LatestPrice_FullMethodName                         = "/loop.internal.pb.mercury.MercuryServerFetcher/LatestPrice"
	MercuryServerFetcher_LatestTimestamp_FullMethodName                     = "/loop.internal.pb.mercury.MercuryServerFetcher/LatestTimestamp"
)

// MercuryServerFetcherClient is the client API for MercuryServerFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MercuryServerFetcherClient interface {
	FetchInitialMaxFinalizedBlockNumber(ctx context.Context, in *FetchInitialMaxFinalizedBlockNumberRequest, opts ...grpc.CallOption) (*FetchInitialMaxFinalizedBlockNumberReply, error)
	LatestPrice(ctx context.Context, in *LatestPriceRequest, opts ...grpc.CallOption) (*LatestPriceReply, error)
	LatestTimestamp(ctx context.Context, in *LatestTimestampRequest, opts ...grpc.CallOption) (*LatestTimestampReply, error)
}

type mercuryServerFetcherClient struct {
	cc grpc.ClientConnInterface
}

func NewMercuryServerFetcherClient(cc grpc.ClientConnInterface) MercuryServerFetcherClient {
	return &mercuryServerFetcherClient{cc}
}

func (c *mercuryServerFetcherClient) FetchInitialMaxFinalizedBlockNumber(ctx context.Context, in *FetchInitialMaxFinalizedBlockNumberRequest, opts ...grpc.CallOption) (*FetchInitialMaxFinalizedBlockNumberReply, error) {
	out := new(FetchInitialMaxFinalizedBlockNumberReply)
	err := c.cc.Invoke(ctx, MercuryServerFetcher_FetchInitialMaxFinalizedBlockNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mercuryServerFetcherClient) LatestPrice(ctx context.Context, in *LatestPriceRequest, opts ...grpc.CallOption) (*LatestPriceReply, error) {
	out := new(LatestPriceReply)
	err := c.cc.Invoke(ctx, MercuryServerFetcher_LatestPrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mercuryServerFetcherClient) LatestTimestamp(ctx context.Context, in *LatestTimestampRequest, opts ...grpc.CallOption) (*LatestTimestampReply, error) {
	out := new(LatestTimestampReply)
	err := c.cc.Invoke(ctx, MercuryServerFetcher_LatestTimestamp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MercuryServerFetcherServer is the server API for MercuryServerFetcher service.
// All implementations must embed UnimplementedMercuryServerFetcherServer
// for forward compatibility
type MercuryServerFetcherServer interface {
	FetchInitialMaxFinalizedBlockNumber(context.Context, *FetchInitialMaxFinalizedBlockNumberRequest) (*FetchInitialMaxFinalizedBlockNumberReply, error)
	LatestPrice(context.Context, *LatestPriceRequest) (*LatestPriceReply, error)
	LatestTimestamp(context.Context, *LatestTimestampRequest) (*LatestTimestampReply, error)
	mustEmbedUnimplementedMercuryServerFetcherServer()
}

// UnimplementedMercuryServerFetcherServer must be embedded to have forward compatible implementations.
type UnimplementedMercuryServerFetcherServer struct {
}

func (UnimplementedMercuryServerFetcherServer) FetchInitialMaxFinalizedBlockNumber(context.Context, *FetchInitialMaxFinalizedBlockNumberRequest) (*FetchInitialMaxFinalizedBlockNumberReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchInitialMaxFinalizedBlockNumber not implemented")
}
func (UnimplementedMercuryServerFetcherServer) LatestPrice(context.Context, *LatestPriceRequest) (*LatestPriceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestPrice not implemented")
}
func (UnimplementedMercuryServerFetcherServer) LatestTimestamp(context.Context, *LatestTimestampRequest) (*LatestTimestampReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestTimestamp not implemented")
}
func (UnimplementedMercuryServerFetcherServer) mustEmbedUnimplementedMercuryServerFetcherServer() {}

// UnsafeMercuryServerFetcherServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MercuryServerFetcherServer will
// result in compilation errors.
type UnsafeMercuryServerFetcherServer interface {
	mustEmbedUnimplementedMercuryServerFetcherServer()
}

func RegisterMercuryServerFetcherServer(s grpc.ServiceRegistrar, srv MercuryServerFetcherServer) {
	s.RegisterService(&MercuryServerFetcher_ServiceDesc, srv)
}

func _MercuryServerFetcher_FetchInitialMaxFinalizedBlockNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchInitialMaxFinalizedBlockNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercuryServerFetcherServer).FetchInitialMaxFinalizedBlockNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MercuryServerFetcher_FetchInitialMaxFinalizedBlockNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercuryServerFetcherServer).FetchInitialMaxFinalizedBlockNumber(ctx, req.(*FetchInitialMaxFinalizedBlockNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MercuryServerFetcher_LatestPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercuryServerFetcherServer).LatestPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MercuryServerFetcher_LatestPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercuryServerFetcherServer).LatestPrice(ctx, req.(*LatestPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MercuryServerFetcher_LatestTimestamp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestTimestampRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercuryServerFetcherServer).LatestTimestamp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MercuryServerFetcher_LatestTimestamp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercuryServerFetcherServer).LatestTimestamp(ctx, req.(*LatestTimestampRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MercuryServerFetcher_ServiceDesc is the grpc.ServiceDesc for MercuryServerFetcher service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MercuryServerFetcher_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.internal.pb.mercury.MercuryServerFetcher",
	HandlerType: (*MercuryServerFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchInitialMaxFinalizedBlockNumber",
			Handler:    _MercuryServerFetcher_FetchInitialMaxFinalizedBlockNumber_Handler,
		},
		{
			MethodName: "LatestPrice",
			Handler:    _MercuryServerFetcher_LatestPrice_Handler,
		},
		{
			MethodName: "LatestTimestamp",
			Handler:    _MercuryServerFetcher_LatestTimestamp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mercury.proto",
}

const (
	MercuryChainReader_LatestHeads_FullMethodName = "/loop.internal.pb.mercury.MercuryChainReader/LatestHeads"
)

// MercuryChainReaderClient is the client API for MercuryChainReader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MercuryChainReaderClient interface {
	LatestHeads(ctx context.Context, in *LatestHeadsRequest, opts ...grpc.CallOption) (*LatestHeadsReply, error)
}

type mercuryChainReaderClient struct {
	cc grpc.ClientConnInterface
}

func NewMercuryChainReaderClient(cc grpc.ClientConnInterface) MercuryChainReaderClient {
	return &mercuryChainReaderClient{cc}
}

func (c *mercuryChainReaderClient) LatestHeads(ctx context.Context, in *LatestHeadsRequest, opts ...grpc.CallOption) (*LatestHeadsReply, error) {
	out := new(LatestHeadsReply)
	err := c.cc.Invoke(ctx, MercuryChainReader_LatestHeads_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MercuryChainReaderServer is the server API for MercuryChainReader service.
// All implementations must embed UnimplementedMercuryChainReaderServer
// for forward compatibility
type MercuryChainReaderServer interface {
	LatestHeads(context.Context, *LatestHeadsRequest) (*LatestHeadsReply, error)
	mustEmbedUnimplementedMercuryChainReaderServer()
}

// UnimplementedMercuryChainReaderServer must be embedded to have forward compatible implementations.
type UnimplementedMercuryChainReaderServer struct {
}

func (UnimplementedMercuryChainReaderServer) LatestHeads(context.Context, *LatestHeadsRequest) (*LatestHeadsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LatestHeads not implemented")
}
func (UnimplementedMercuryChainReaderServer) mustEmbedUnimplementedMercuryChainReaderServer() {}

// UnsafeMercuryChainReaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MercuryChainReaderServer will
// result in compilation errors.
type UnsafeMercuryChainReaderServer interface {
	mustEmbedUnimplementedMercuryChainReaderServer()
}

func RegisterMercuryChainReaderServer(s grpc.ServiceRegistrar, srv MercuryChainReaderServer) {
	s.RegisterService(&MercuryChainReader_ServiceDesc, srv)
}

func _MercuryChainReader_LatestHeads_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LatestHeadsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MercuryChainReaderServer).LatestHeads(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MercuryChainReader_LatestHeads_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MercuryChainReaderServer).LatestHeads(ctx, req.(*LatestHeadsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MercuryChainReader_ServiceDesc is the grpc.ServiceDesc for MercuryChainReader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MercuryChainReader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.internal.pb.mercury.MercuryChainReader",
	HandlerType: (*MercuryChainReaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LatestHeads",
			Handler:    _MercuryChainReader_LatestHeads_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mercury.proto",
}

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: reportcodec.proto

package mercury_v1_pb

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
	ReportCodec_BuildReport_FullMethodName               = "/loop.internal.pb.mercury.v1.ReportCodec/BuildReport"
	ReportCodec_MaxReportLength_FullMethodName           = "/loop.internal.pb.mercury.v1.ReportCodec/MaxReportLength"
	ReportCodec_CurrentBlockNumFromReport_FullMethodName = "/loop.internal.pb.mercury.v1.ReportCodec/CurrentBlockNumFromReport"
)

// ReportCodecClient is the client API for ReportCodec service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportCodecClient interface {
	BuildReport(ctx context.Context, in *BuildReportRequest, opts ...grpc.CallOption) (*BuildReportReply, error)
	MaxReportLength(ctx context.Context, in *MaxReportLengthRequest, opts ...grpc.CallOption) (*MaxReportLengthReply, error)
	CurrentBlockNumFromReport(ctx context.Context, in *CurrentBlockNumFromReportRequest, opts ...grpc.CallOption) (*CurrentBlockNumFromReportResponse, error)
}

type reportCodecClient struct {
	cc grpc.ClientConnInterface
}

func NewReportCodecClient(cc grpc.ClientConnInterface) ReportCodecClient {
	return &reportCodecClient{cc}
}

func (c *reportCodecClient) BuildReport(ctx context.Context, in *BuildReportRequest, opts ...grpc.CallOption) (*BuildReportReply, error) {
	out := new(BuildReportReply)
	err := c.cc.Invoke(ctx, ReportCodec_BuildReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportCodecClient) MaxReportLength(ctx context.Context, in *MaxReportLengthRequest, opts ...grpc.CallOption) (*MaxReportLengthReply, error) {
	out := new(MaxReportLengthReply)
	err := c.cc.Invoke(ctx, ReportCodec_MaxReportLength_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reportCodecClient) CurrentBlockNumFromReport(ctx context.Context, in *CurrentBlockNumFromReportRequest, opts ...grpc.CallOption) (*CurrentBlockNumFromReportResponse, error) {
	out := new(CurrentBlockNumFromReportResponse)
	err := c.cc.Invoke(ctx, ReportCodec_CurrentBlockNumFromReport_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportCodecServer is the server API for ReportCodec service.
// All implementations must embed UnimplementedReportCodecServer
// for forward compatibility
type ReportCodecServer interface {
	BuildReport(context.Context, *BuildReportRequest) (*BuildReportReply, error)
	MaxReportLength(context.Context, *MaxReportLengthRequest) (*MaxReportLengthReply, error)
	CurrentBlockNumFromReport(context.Context, *CurrentBlockNumFromReportRequest) (*CurrentBlockNumFromReportResponse, error)
	mustEmbedUnimplementedReportCodecServer()
}

// UnimplementedReportCodecServer must be embedded to have forward compatible implementations.
type UnimplementedReportCodecServer struct {
}

func (UnimplementedReportCodecServer) BuildReport(context.Context, *BuildReportRequest) (*BuildReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuildReport not implemented")
}
func (UnimplementedReportCodecServer) MaxReportLength(context.Context, *MaxReportLengthRequest) (*MaxReportLengthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MaxReportLength not implemented")
}
func (UnimplementedReportCodecServer) CurrentBlockNumFromReport(context.Context, *CurrentBlockNumFromReportRequest) (*CurrentBlockNumFromReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CurrentBlockNumFromReport not implemented")
}
func (UnimplementedReportCodecServer) mustEmbedUnimplementedReportCodecServer() {}

// UnsafeReportCodecServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportCodecServer will
// result in compilation errors.
type UnsafeReportCodecServer interface {
	mustEmbedUnimplementedReportCodecServer()
}

func RegisterReportCodecServer(s grpc.ServiceRegistrar, srv ReportCodecServer) {
	s.RegisterService(&ReportCodec_ServiceDesc, srv)
}

func _ReportCodec_BuildReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecServer).BuildReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportCodec_BuildReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecServer).BuildReport(ctx, req.(*BuildReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportCodec_MaxReportLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaxReportLengthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecServer).MaxReportLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportCodec_MaxReportLength_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecServer).MaxReportLength(ctx, req.(*MaxReportLengthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReportCodec_CurrentBlockNumFromReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentBlockNumFromReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportCodecServer).CurrentBlockNumFromReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReportCodec_CurrentBlockNumFromReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportCodecServer).CurrentBlockNumFromReport(ctx, req.(*CurrentBlockNumFromReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReportCodec_ServiceDesc is the grpc.ServiceDesc for ReportCodec service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReportCodec_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.internal.pb.mercury.v1.ReportCodec",
	HandlerType: (*ReportCodecServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BuildReport",
			Handler:    _ReportCodec_BuildReport_Handler,
		},
		{
			MethodName: "MaxReportLength",
			Handler:    _ReportCodec_MaxReportLength_Handler,
		},
		{
			MethodName: "CurrentBlockNumFromReport",
			Handler:    _ReportCodec_CurrentBlockNumFromReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reportcodec.proto",
}

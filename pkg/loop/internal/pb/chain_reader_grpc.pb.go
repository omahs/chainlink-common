// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: chain_reader.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ChainReader_GetLatestValue_FullMethodName = "/loop.ChainReader/GetLatestValue"
	ChainReader_QueryKey_FullMethodName       = "/loop.ChainReader/QueryKey"
	ChainReader_Bind_FullMethodName           = "/loop.ChainReader/Bind"
	ChainReader_Unbind_FullMethodName         = "/loop.ChainReader/Unbind"
)

// ChainReaderClient is the client API for ChainReader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChainReaderClient interface {
	GetLatestValue(ctx context.Context, in *GetLatestValueRequest, opts ...grpc.CallOption) (*GetLatestValueReply, error)
	QueryKey(ctx context.Context, in *QueryKeyRequest, opts ...grpc.CallOption) (*QueryKeyReply, error)
	Bind(ctx context.Context, in *BindRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Unbind(ctx context.Context, in *BindRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type chainReaderClient struct {
	cc grpc.ClientConnInterface
}

func NewChainReaderClient(cc grpc.ClientConnInterface) ChainReaderClient {
	return &chainReaderClient{cc}
}

func (c *chainReaderClient) GetLatestValue(ctx context.Context, in *GetLatestValueRequest, opts ...grpc.CallOption) (*GetLatestValueReply, error) {
	out := new(GetLatestValueReply)
	err := c.cc.Invoke(ctx, ChainReader_GetLatestValue_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainReaderClient) QueryKey(ctx context.Context, in *QueryKeyRequest, opts ...grpc.CallOption) (*QueryKeyReply, error) {
	out := new(QueryKeyReply)
	err := c.cc.Invoke(ctx, ChainReader_QueryKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainReaderClient) Bind(ctx context.Context, in *BindRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChainReader_Bind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chainReaderClient) Unbind(ctx context.Context, in *BindRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ChainReader_Unbind_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainReaderServer is the server API for ChainReader service.
// All implementations must embed UnimplementedChainReaderServer
// for forward compatibility
type ChainReaderServer interface {
	GetLatestValue(context.Context, *GetLatestValueRequest) (*GetLatestValueReply, error)
	QueryKey(context.Context, *QueryKeyRequest) (*QueryKeyReply, error)
	Bind(context.Context, *BindRequest) (*emptypb.Empty, error)
	Unbind(context.Context, *BindRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedChainReaderServer()
}

// UnimplementedChainReaderServer must be embedded to have forward compatible implementations.
type UnimplementedChainReaderServer struct {
}

func (UnimplementedChainReaderServer) GetLatestValue(context.Context, *GetLatestValueRequest) (*GetLatestValueReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestValue not implemented")
}
func (UnimplementedChainReaderServer) QueryKey(context.Context, *QueryKeyRequest) (*QueryKeyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryKey not implemented")
}
func (UnimplementedChainReaderServer) Bind(context.Context, *BindRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bind not implemented")
}
func (UnimplementedChainReaderServer) Unbind(context.Context, *BindRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unbind not implemented")
}
func (UnimplementedChainReaderServer) mustEmbedUnimplementedChainReaderServer() {}

// UnsafeChainReaderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChainReaderServer will
// result in compilation errors.
type UnsafeChainReaderServer interface {
	mustEmbedUnimplementedChainReaderServer()
}

func RegisterChainReaderServer(s grpc.ServiceRegistrar, srv ChainReaderServer) {
	s.RegisterService(&ChainReader_ServiceDesc, srv)
}

func _ChainReader_GetLatestValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLatestValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainReaderServer).GetLatestValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChainReader_GetLatestValue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainReaderServer).GetLatestValue(ctx, req.(*GetLatestValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainReader_QueryKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainReaderServer).QueryKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChainReader_QueryKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainReaderServer).QueryKey(ctx, req.(*QueryKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainReader_Bind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainReaderServer).Bind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChainReader_Bind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainReaderServer).Bind(ctx, req.(*BindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChainReader_Unbind_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BindRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainReaderServer).Unbind(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChainReader_Unbind_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainReaderServer).Unbind(ctx, req.(*BindRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChainReader_ServiceDesc is the grpc.ServiceDesc for ChainReader service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChainReader_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.ChainReader",
	HandlerType: (*ChainReaderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestValue",
			Handler:    _ChainReader_GetLatestValue_Handler,
		},
		{
			MethodName: "QueryKey",
			Handler:    _ChainReader_QueryKey_Handler,
		},
		{
			MethodName: "Bind",
			Handler:    _ChainReader_Bind_Handler,
		},
		{
			MethodName: "Unbind",
			Handler:    _ChainReader_Unbind_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chain_reader.proto",
}

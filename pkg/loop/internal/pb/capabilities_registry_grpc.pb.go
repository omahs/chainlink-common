// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: loop/internal/pb/capabilities_registry.proto

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
	CapabilitiesRegistry_GetLocalNode_FullMethodName        = "/loop.CapabilitiesRegistry/GetLocalNode"
	CapabilitiesRegistry_ConfigForCapability_FullMethodName = "/loop.CapabilitiesRegistry/ConfigForCapability"
	CapabilitiesRegistry_Get_FullMethodName                 = "/loop.CapabilitiesRegistry/Get"
	CapabilitiesRegistry_GetTrigger_FullMethodName          = "/loop.CapabilitiesRegistry/GetTrigger"
	CapabilitiesRegistry_GetAction_FullMethodName           = "/loop.CapabilitiesRegistry/GetAction"
	CapabilitiesRegistry_GetConsensus_FullMethodName        = "/loop.CapabilitiesRegistry/GetConsensus"
	CapabilitiesRegistry_GetTarget_FullMethodName           = "/loop.CapabilitiesRegistry/GetTarget"
	CapabilitiesRegistry_List_FullMethodName                = "/loop.CapabilitiesRegistry/List"
	CapabilitiesRegistry_Add_FullMethodName                 = "/loop.CapabilitiesRegistry/Add"
)

// CapabilitiesRegistryClient is the client API for CapabilitiesRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CapabilitiesRegistryClient interface {
	GetLocalNode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLocalNodeReply, error)
	ConfigForCapability(ctx context.Context, in *ConfigForCapabilityRequest, opts ...grpc.CallOption) (*ConfigForCapabilityReply, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error)
	GetTrigger(ctx context.Context, in *GetTriggerRequest, opts ...grpc.CallOption) (*GetTriggerReply, error)
	GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionReply, error)
	GetConsensus(ctx context.Context, in *GetConsensusRequest, opts ...grpc.CallOption) (*GetConsensusReply, error)
	GetTarget(ctx context.Context, in *GetTargetRequest, opts ...grpc.CallOption) (*GetTargetReply, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListReply, error)
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type capabilitiesRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewCapabilitiesRegistryClient(cc grpc.ClientConnInterface) CapabilitiesRegistryClient {
	return &capabilitiesRegistryClient{cc}
}

func (c *capabilitiesRegistryClient) GetLocalNode(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetLocalNodeReply, error) {
	out := new(GetLocalNodeReply)
	err := c.cc.Invoke(ctx, CapabilitiesRegistry_GetLocalNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilitiesRegistryClient) ConfigForCapability(ctx context.Context, in *ConfigForCapabilityRequest, opts ...grpc.CallOption) (*ConfigForCapabilityReply, error) {
	out := new(ConfigForCapabilityReply)
	err := c.cc.Invoke(ctx, CapabilitiesRegistry_ConfigForCapability_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilitiesRegistryClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetReply, error) {
	out := new(GetReply)
	err := c.cc.Invoke(ctx, CapabilitiesRegistry_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilitiesRegistryClient) GetTrigger(ctx context.Context, in *GetTriggerRequest, opts ...grpc.CallOption) (*GetTriggerReply, error) {
	out := new(GetTriggerReply)
	err := c.cc.Invoke(ctx, CapabilitiesRegistry_GetTrigger_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilitiesRegistryClient) GetAction(ctx context.Context, in *GetActionRequest, opts ...grpc.CallOption) (*GetActionReply, error) {
	out := new(GetActionReply)
	err := c.cc.Invoke(ctx, CapabilitiesRegistry_GetAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilitiesRegistryClient) GetConsensus(ctx context.Context, in *GetConsensusRequest, opts ...grpc.CallOption) (*GetConsensusReply, error) {
	out := new(GetConsensusReply)
	err := c.cc.Invoke(ctx, CapabilitiesRegistry_GetConsensus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilitiesRegistryClient) GetTarget(ctx context.Context, in *GetTargetRequest, opts ...grpc.CallOption) (*GetTargetReply, error) {
	out := new(GetTargetReply)
	err := c.cc.Invoke(ctx, CapabilitiesRegistry_GetTarget_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilitiesRegistryClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListReply, error) {
	out := new(ListReply)
	err := c.cc.Invoke(ctx, CapabilitiesRegistry_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capabilitiesRegistryClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CapabilitiesRegistry_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CapabilitiesRegistryServer is the server API for CapabilitiesRegistry service.
// All implementations must embed UnimplementedCapabilitiesRegistryServer
// for forward compatibility
type CapabilitiesRegistryServer interface {
	GetLocalNode(context.Context, *emptypb.Empty) (*GetLocalNodeReply, error)
	ConfigForCapability(context.Context, *ConfigForCapabilityRequest) (*ConfigForCapabilityReply, error)
	Get(context.Context, *GetRequest) (*GetReply, error)
	GetTrigger(context.Context, *GetTriggerRequest) (*GetTriggerReply, error)
	GetAction(context.Context, *GetActionRequest) (*GetActionReply, error)
	GetConsensus(context.Context, *GetConsensusRequest) (*GetConsensusReply, error)
	GetTarget(context.Context, *GetTargetRequest) (*GetTargetReply, error)
	List(context.Context, *emptypb.Empty) (*ListReply, error)
	Add(context.Context, *AddRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCapabilitiesRegistryServer()
}

// UnimplementedCapabilitiesRegistryServer must be embedded to have forward compatible implementations.
type UnimplementedCapabilitiesRegistryServer struct {
}

func (UnimplementedCapabilitiesRegistryServer) GetLocalNode(context.Context, *emptypb.Empty) (*GetLocalNodeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocalNode not implemented")
}
func (UnimplementedCapabilitiesRegistryServer) ConfigForCapability(context.Context, *ConfigForCapabilityRequest) (*ConfigForCapabilityReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigForCapability not implemented")
}
func (UnimplementedCapabilitiesRegistryServer) Get(context.Context, *GetRequest) (*GetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCapabilitiesRegistryServer) GetTrigger(context.Context, *GetTriggerRequest) (*GetTriggerReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrigger not implemented")
}
func (UnimplementedCapabilitiesRegistryServer) GetAction(context.Context, *GetActionRequest) (*GetActionReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAction not implemented")
}
func (UnimplementedCapabilitiesRegistryServer) GetConsensus(context.Context, *GetConsensusRequest) (*GetConsensusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsensus not implemented")
}
func (UnimplementedCapabilitiesRegistryServer) GetTarget(context.Context, *GetTargetRequest) (*GetTargetReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTarget not implemented")
}
func (UnimplementedCapabilitiesRegistryServer) List(context.Context, *emptypb.Empty) (*ListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCapabilitiesRegistryServer) Add(context.Context, *AddRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCapabilitiesRegistryServer) mustEmbedUnimplementedCapabilitiesRegistryServer() {}

// UnsafeCapabilitiesRegistryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CapabilitiesRegistryServer will
// result in compilation errors.
type UnsafeCapabilitiesRegistryServer interface {
	mustEmbedUnimplementedCapabilitiesRegistryServer()
}

func RegisterCapabilitiesRegistryServer(s grpc.ServiceRegistrar, srv CapabilitiesRegistryServer) {
	s.RegisterService(&CapabilitiesRegistry_ServiceDesc, srv)
}

func _CapabilitiesRegistry_GetLocalNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilitiesRegistryServer).GetLocalNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CapabilitiesRegistry_GetLocalNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilitiesRegistryServer).GetLocalNode(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CapabilitiesRegistry_ConfigForCapability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigForCapabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilitiesRegistryServer).ConfigForCapability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CapabilitiesRegistry_ConfigForCapability_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilitiesRegistryServer).ConfigForCapability(ctx, req.(*ConfigForCapabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CapabilitiesRegistry_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilitiesRegistryServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CapabilitiesRegistry_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilitiesRegistryServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CapabilitiesRegistry_GetTrigger_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilitiesRegistryServer).GetTrigger(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CapabilitiesRegistry_GetTrigger_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilitiesRegistryServer).GetTrigger(ctx, req.(*GetTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CapabilitiesRegistry_GetAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilitiesRegistryServer).GetAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CapabilitiesRegistry_GetAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilitiesRegistryServer).GetAction(ctx, req.(*GetActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CapabilitiesRegistry_GetConsensus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsensusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilitiesRegistryServer).GetConsensus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CapabilitiesRegistry_GetConsensus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilitiesRegistryServer).GetConsensus(ctx, req.(*GetConsensusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CapabilitiesRegistry_GetTarget_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTargetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilitiesRegistryServer).GetTarget(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CapabilitiesRegistry_GetTarget_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilitiesRegistryServer).GetTarget(ctx, req.(*GetTargetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CapabilitiesRegistry_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilitiesRegistryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CapabilitiesRegistry_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilitiesRegistryServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CapabilitiesRegistry_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapabilitiesRegistryServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CapabilitiesRegistry_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapabilitiesRegistryServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CapabilitiesRegistry_ServiceDesc is the grpc.ServiceDesc for CapabilitiesRegistry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CapabilitiesRegistry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loop.CapabilitiesRegistry",
	HandlerType: (*CapabilitiesRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocalNode",
			Handler:    _CapabilitiesRegistry_GetLocalNode_Handler,
		},
		{
			MethodName: "ConfigForCapability",
			Handler:    _CapabilitiesRegistry_ConfigForCapability_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CapabilitiesRegistry_Get_Handler,
		},
		{
			MethodName: "GetTrigger",
			Handler:    _CapabilitiesRegistry_GetTrigger_Handler,
		},
		{
			MethodName: "GetAction",
			Handler:    _CapabilitiesRegistry_GetAction_Handler,
		},
		{
			MethodName: "GetConsensus",
			Handler:    _CapabilitiesRegistry_GetConsensus_Handler,
		},
		{
			MethodName: "GetTarget",
			Handler:    _CapabilitiesRegistry_GetTarget_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CapabilitiesRegistry_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _CapabilitiesRegistry_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loop/internal/pb/capabilities_registry.proto",
}

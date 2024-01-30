// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: warden/intent/tx.proto

package intent

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
	Msg_UpdateParams_FullMethodName  = "/warden.intent.Msg/UpdateParams"
	Msg_ApproveAction_FullMethodName = "/warden.intent.Msg/ApproveAction"
	Msg_NewIntent_FullMethodName     = "/warden.intent.Msg/NewIntent"
	Msg_RevokeAction_FullMethodName  = "/warden.intent.Msg/RevokeAction"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	// Add an approval to an existing Action.
	ApproveAction(ctx context.Context, in *MsgApproveAction, opts ...grpc.CallOption) (*MsgApproveActionResponse, error)
	// Create a new intent.
	NewIntent(ctx context.Context, in *MsgNewIntent, opts ...grpc.CallOption) (*MsgNewIntentResponse, error)
	// Revoke an existing Action while in pending state.
	RevokeAction(ctx context.Context, in *MsgRevokeAction, opts ...grpc.CallOption) (*MsgRevokeActionResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ApproveAction(ctx context.Context, in *MsgApproveAction, opts ...grpc.CallOption) (*MsgApproveActionResponse, error) {
	out := new(MsgApproveActionResponse)
	err := c.cc.Invoke(ctx, Msg_ApproveAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) NewIntent(ctx context.Context, in *MsgNewIntent, opts ...grpc.CallOption) (*MsgNewIntentResponse, error) {
	out := new(MsgNewIntentResponse)
	err := c.cc.Invoke(ctx, Msg_NewIntent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RevokeAction(ctx context.Context, in *MsgRevokeAction, opts ...grpc.CallOption) (*MsgRevokeActionResponse, error) {
	out := new(MsgRevokeActionResponse)
	err := c.cc.Invoke(ctx, Msg_RevokeAction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	// Add an approval to an existing Action.
	ApproveAction(context.Context, *MsgApproveAction) (*MsgApproveActionResponse, error)
	// Create a new intent.
	NewIntent(context.Context, *MsgNewIntent) (*MsgNewIntentResponse, error)
	// Revoke an existing Action while in pending state.
	RevokeAction(context.Context, *MsgRevokeAction) (*MsgRevokeActionResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) ApproveAction(context.Context, *MsgApproveAction) (*MsgApproveActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveAction not implemented")
}
func (UnimplementedMsgServer) NewIntent(context.Context, *MsgNewIntent) (*MsgNewIntentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewIntent not implemented")
}
func (UnimplementedMsgServer) RevokeAction(context.Context, *MsgRevokeAction) (*MsgRevokeActionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeAction not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ApproveAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgApproveAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApproveAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ApproveAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApproveAction(ctx, req.(*MsgApproveAction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_NewIntent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgNewIntent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).NewIntent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_NewIntent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).NewIntent(ctx, req.(*MsgNewIntent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RevokeAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRevokeAction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RevokeAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RevokeAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RevokeAction(ctx, req.(*MsgRevokeAction))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "warden.intent.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "ApproveAction",
			Handler:    _Msg_ApproveAction_Handler,
		},
		{
			MethodName: "NewIntent",
			Handler:    _Msg_NewIntent_Handler,
		},
		{
			MethodName: "RevokeAction",
			Handler:    _Msg_RevokeAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warden/intent/tx.proto",
}

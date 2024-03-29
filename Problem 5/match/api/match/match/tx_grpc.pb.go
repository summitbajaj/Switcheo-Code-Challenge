// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: match/match/tx.proto

package match

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
	Msg_UpdateParams_FullMethodName    = "/match.match.Msg/UpdateParams"
	Msg_CreateMatchInfo_FullMethodName = "/match.match.Msg/CreateMatchInfo"
	Msg_UpdateMatchInfo_FullMethodName = "/match.match.Msg/UpdateMatchInfo"
	Msg_DeleteMatchInfo_FullMethodName = "/match.match.Msg/DeleteMatchInfo"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreateMatchInfo(ctx context.Context, in *MsgCreateMatchInfo, opts ...grpc.CallOption) (*MsgCreateMatchInfoResponse, error)
	UpdateMatchInfo(ctx context.Context, in *MsgUpdateMatchInfo, opts ...grpc.CallOption) (*MsgUpdateMatchInfoResponse, error)
	DeleteMatchInfo(ctx context.Context, in *MsgDeleteMatchInfo, opts ...grpc.CallOption) (*MsgDeleteMatchInfoResponse, error)
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

func (c *msgClient) CreateMatchInfo(ctx context.Context, in *MsgCreateMatchInfo, opts ...grpc.CallOption) (*MsgCreateMatchInfoResponse, error) {
	out := new(MsgCreateMatchInfoResponse)
	err := c.cc.Invoke(ctx, Msg_CreateMatchInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateMatchInfo(ctx context.Context, in *MsgUpdateMatchInfo, opts ...grpc.CallOption) (*MsgUpdateMatchInfoResponse, error) {
	out := new(MsgUpdateMatchInfoResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateMatchInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteMatchInfo(ctx context.Context, in *MsgDeleteMatchInfo, opts ...grpc.CallOption) (*MsgDeleteMatchInfoResponse, error) {
	out := new(MsgDeleteMatchInfoResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteMatchInfo_FullMethodName, in, out, opts...)
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
	CreateMatchInfo(context.Context, *MsgCreateMatchInfo) (*MsgCreateMatchInfoResponse, error)
	UpdateMatchInfo(context.Context, *MsgUpdateMatchInfo) (*MsgUpdateMatchInfoResponse, error)
	DeleteMatchInfo(context.Context, *MsgDeleteMatchInfo) (*MsgDeleteMatchInfoResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreateMatchInfo(context.Context, *MsgCreateMatchInfo) (*MsgCreateMatchInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMatchInfo not implemented")
}
func (UnimplementedMsgServer) UpdateMatchInfo(context.Context, *MsgUpdateMatchInfo) (*MsgUpdateMatchInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMatchInfo not implemented")
}
func (UnimplementedMsgServer) DeleteMatchInfo(context.Context, *MsgDeleteMatchInfo) (*MsgDeleteMatchInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMatchInfo not implemented")
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

func _Msg_CreateMatchInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateMatchInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateMatchInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreateMatchInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateMatchInfo(ctx, req.(*MsgCreateMatchInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateMatchInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateMatchInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateMatchInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateMatchInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateMatchInfo(ctx, req.(*MsgUpdateMatchInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteMatchInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteMatchInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteMatchInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteMatchInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteMatchInfo(ctx, req.(*MsgDeleteMatchInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "match.match.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreateMatchInfo",
			Handler:    _Msg_CreateMatchInfo_Handler,
		},
		{
			MethodName: "UpdateMatchInfo",
			Handler:    _Msg_UpdateMatchInfo_Handler,
		},
		{
			MethodName: "DeleteMatchInfo",
			Handler:    _Msg_DeleteMatchInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "match/match/tx.proto",
}

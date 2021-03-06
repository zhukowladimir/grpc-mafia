// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package service

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

// MafiaClient is the client API for Mafia service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MafiaClient interface {
	SetUser(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*ACK, error)
	GetUsersList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UsersListResponse, error)
	ListenNotifications(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (Mafia_ListenNotificationsClient, error)
	Disconnect(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	StartGame(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LynchingVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*ACK, error)
	GoSleep(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	NightMurder(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*ACK, error)
	SneakPeek(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*ACK, error)
}

type mafiaClient struct {
	cc grpc.ClientConnInterface
}

func NewMafiaClient(cc grpc.ClientConnInterface) MafiaClient {
	return &mafiaClient{cc}
}

func (c *mafiaClient) SetUser(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*ACK, error) {
	out := new(ACK)
	err := c.cc.Invoke(ctx, "/service.Mafia/SetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) GetUsersList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*UsersListResponse, error) {
	out := new(UsersListResponse)
	err := c.cc.Invoke(ctx, "/service.Mafia/GetUsersList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) ListenNotifications(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (Mafia_ListenNotificationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Mafia_ServiceDesc.Streams[0], "/service.Mafia/ListenNotifications", opts...)
	if err != nil {
		return nil, err
	}
	x := &mafiaListenNotificationsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Mafia_ListenNotificationsClient interface {
	Recv() (*Notification, error)
	grpc.ClientStream
}

type mafiaListenNotificationsClient struct {
	grpc.ClientStream
}

func (x *mafiaListenNotificationsClient) Recv() (*Notification, error) {
	m := new(Notification)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mafiaClient) Disconnect(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.Mafia/Disconnect", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) StartGame(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.Mafia/StartGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) LynchingVote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*ACK, error) {
	out := new(ACK)
	err := c.cc.Invoke(ctx, "/service.Mafia/LynchingVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) GoSleep(ctx context.Context, in *StringRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.Mafia/GoSleep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) NightMurder(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*ACK, error) {
	out := new(ACK)
	err := c.cc.Invoke(ctx, "/service.Mafia/NightMurder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mafiaClient) SneakPeek(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*ACK, error) {
	out := new(ACK)
	err := c.cc.Invoke(ctx, "/service.Mafia/SneakPeek", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MafiaServer is the server API for Mafia service.
// All implementations must embed UnimplementedMafiaServer
// for forward compatibility
type MafiaServer interface {
	SetUser(context.Context, *StringRequest) (*ACK, error)
	GetUsersList(context.Context, *emptypb.Empty) (*UsersListResponse, error)
	ListenNotifications(*StringRequest, Mafia_ListenNotificationsServer) error
	Disconnect(context.Context, *StringRequest) (*emptypb.Empty, error)
	StartGame(context.Context, *StringRequest) (*emptypb.Empty, error)
	LynchingVote(context.Context, *VoteRequest) (*ACK, error)
	GoSleep(context.Context, *StringRequest) (*emptypb.Empty, error)
	NightMurder(context.Context, *VoteRequest) (*ACK, error)
	SneakPeek(context.Context, *VoteRequest) (*ACK, error)
	mustEmbedUnimplementedMafiaServer()
}

// UnimplementedMafiaServer must be embedded to have forward compatible implementations.
type UnimplementedMafiaServer struct {
}

func (UnimplementedMafiaServer) SetUser(context.Context, *StringRequest) (*ACK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUser not implemented")
}
func (UnimplementedMafiaServer) GetUsersList(context.Context, *emptypb.Empty) (*UsersListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersList not implemented")
}
func (UnimplementedMafiaServer) ListenNotifications(*StringRequest, Mafia_ListenNotificationsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListenNotifications not implemented")
}
func (UnimplementedMafiaServer) Disconnect(context.Context, *StringRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disconnect not implemented")
}
func (UnimplementedMafiaServer) StartGame(context.Context, *StringRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartGame not implemented")
}
func (UnimplementedMafiaServer) LynchingVote(context.Context, *VoteRequest) (*ACK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LynchingVote not implemented")
}
func (UnimplementedMafiaServer) GoSleep(context.Context, *StringRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GoSleep not implemented")
}
func (UnimplementedMafiaServer) NightMurder(context.Context, *VoteRequest) (*ACK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NightMurder not implemented")
}
func (UnimplementedMafiaServer) SneakPeek(context.Context, *VoteRequest) (*ACK, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SneakPeek not implemented")
}
func (UnimplementedMafiaServer) mustEmbedUnimplementedMafiaServer() {}

// UnsafeMafiaServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MafiaServer will
// result in compilation errors.
type UnsafeMafiaServer interface {
	mustEmbedUnimplementedMafiaServer()
}

func RegisterMafiaServer(s grpc.ServiceRegistrar, srv MafiaServer) {
	s.RegisterService(&Mafia_ServiceDesc, srv)
}

func _Mafia_SetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).SetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Mafia/SetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).SetUser(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_GetUsersList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).GetUsersList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Mafia/GetUsersList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).GetUsersList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_ListenNotifications_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StringRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MafiaServer).ListenNotifications(m, &mafiaListenNotificationsServer{stream})
}

type Mafia_ListenNotificationsServer interface {
	Send(*Notification) error
	grpc.ServerStream
}

type mafiaListenNotificationsServer struct {
	grpc.ServerStream
}

func (x *mafiaListenNotificationsServer) Send(m *Notification) error {
	return x.ServerStream.SendMsg(m)
}

func _Mafia_Disconnect_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).Disconnect(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Mafia/Disconnect",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).Disconnect(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_StartGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).StartGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Mafia/StartGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).StartGame(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_LynchingVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).LynchingVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Mafia/LynchingVote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).LynchingVote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_GoSleep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).GoSleep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Mafia/GoSleep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).GoSleep(ctx, req.(*StringRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_NightMurder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).NightMurder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Mafia/NightMurder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).NightMurder(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Mafia_SneakPeek_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MafiaServer).SneakPeek(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Mafia/SneakPeek",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MafiaServer).SneakPeek(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Mafia_ServiceDesc is the grpc.ServiceDesc for Mafia service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Mafia_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.Mafia",
	HandlerType: (*MafiaServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetUser",
			Handler:    _Mafia_SetUser_Handler,
		},
		{
			MethodName: "GetUsersList",
			Handler:    _Mafia_GetUsersList_Handler,
		},
		{
			MethodName: "Disconnect",
			Handler:    _Mafia_Disconnect_Handler,
		},
		{
			MethodName: "StartGame",
			Handler:    _Mafia_StartGame_Handler,
		},
		{
			MethodName: "LynchingVote",
			Handler:    _Mafia_LynchingVote_Handler,
		},
		{
			MethodName: "GoSleep",
			Handler:    _Mafia_GoSleep_Handler,
		},
		{
			MethodName: "NightMurder",
			Handler:    _Mafia_NightMurder_Handler,
		},
		{
			MethodName: "SneakPeek",
			Handler:    _Mafia_SneakPeek_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListenNotifications",
			Handler:       _Mafia_ListenNotifications_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}

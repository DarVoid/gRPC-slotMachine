// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: gamevitaedashboard.proto

package gameVitaeDashboard

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

// GameVitaeServiceClient is the client API for GameVitaeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GameVitaeServiceClient interface {
	// Create a Game
	RetrieveSessionData(ctx context.Context, in *SessionParameterRequest, opts ...grpc.CallOption) (*SessionParameterReply, error)
}

type gameVitaeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGameVitaeServiceClient(cc grpc.ClientConnInterface) GameVitaeServiceClient {
	return &gameVitaeServiceClient{cc}
}

func (c *gameVitaeServiceClient) RetrieveSessionData(ctx context.Context, in *SessionParameterRequest, opts ...grpc.CallOption) (*SessionParameterReply, error) {
	out := new(SessionParameterReply)
	err := c.cc.Invoke(ctx, "/GameVitaeService/RetrieveSessionData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GameVitaeServiceServer is the server API for GameVitaeService service.
// All implementations should embed UnimplementedGameVitaeServiceServer
// for forward compatibility
type GameVitaeServiceServer interface {
	// Create a Game
	RetrieveSessionData(context.Context, *SessionParameterRequest) (*SessionParameterReply, error)
}

// UnimplementedGameVitaeServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGameVitaeServiceServer struct {
}

func (UnimplementedGameVitaeServiceServer) RetrieveSessionData(context.Context, *SessionParameterRequest) (*SessionParameterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveSessionData not implemented")
}

// UnsafeGameVitaeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GameVitaeServiceServer will
// result in compilation errors.
type UnsafeGameVitaeServiceServer interface {
	mustEmbedUnimplementedGameVitaeServiceServer()
}

func RegisterGameVitaeServiceServer(s grpc.ServiceRegistrar, srv GameVitaeServiceServer) {
	s.RegisterService(&GameVitaeService_ServiceDesc, srv)
}

func _GameVitaeService_RetrieveSessionData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionParameterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GameVitaeServiceServer).RetrieveSessionData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/GameVitaeService/RetrieveSessionData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GameVitaeServiceServer).RetrieveSessionData(ctx, req.(*SessionParameterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GameVitaeService_ServiceDesc is the grpc.ServiceDesc for GameVitaeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GameVitaeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "GameVitaeService",
	HandlerType: (*GameVitaeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveSessionData",
			Handler:    _GameVitaeService_RetrieveSessionData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gamevitaedashboard.proto",
}
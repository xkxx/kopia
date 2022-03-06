// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: repository_server.proto

package grpcapi

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

// KopiaRepositoryClient is the client API for KopiaRepository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KopiaRepositoryClient interface {
	// Session starts a long-running repository session.
	Session(ctx context.Context, opts ...grpc.CallOption) (KopiaRepository_SessionClient, error)
}

type kopiaRepositoryClient struct {
	cc grpc.ClientConnInterface
}

func NewKopiaRepositoryClient(cc grpc.ClientConnInterface) KopiaRepositoryClient {
	return &kopiaRepositoryClient{cc}
}

func (c *kopiaRepositoryClient) Session(ctx context.Context, opts ...grpc.CallOption) (KopiaRepository_SessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &KopiaRepository_ServiceDesc.Streams[0], "/kopia_repository.KopiaRepository/Session", opts...)
	if err != nil {
		return nil, err
	}
	x := &kopiaRepositorySessionClient{stream}
	return x, nil
}

type KopiaRepository_SessionClient interface {
	Send(*SessionRequest) error
	Recv() (*SessionResponse, error)
	grpc.ClientStream
}

type kopiaRepositorySessionClient struct {
	grpc.ClientStream
}

func (x *kopiaRepositorySessionClient) Send(m *SessionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *kopiaRepositorySessionClient) Recv() (*SessionResponse, error) {
	m := new(SessionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KopiaRepositoryServer is the server API for KopiaRepository service.
// All implementations must embed UnimplementedKopiaRepositoryServer
// for forward compatibility
type KopiaRepositoryServer interface {
	// Session starts a long-running repository session.
	Session(KopiaRepository_SessionServer) error
	mustEmbedUnimplementedKopiaRepositoryServer()
}

// UnimplementedKopiaRepositoryServer must be embedded to have forward compatible implementations.
type UnimplementedKopiaRepositoryServer struct {
}

func (UnimplementedKopiaRepositoryServer) Session(KopiaRepository_SessionServer) error {
	return status.Errorf(codes.Unimplemented, "method Session not implemented")
}
func (UnimplementedKopiaRepositoryServer) mustEmbedUnimplementedKopiaRepositoryServer() {}

// UnsafeKopiaRepositoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KopiaRepositoryServer will
// result in compilation errors.
type UnsafeKopiaRepositoryServer interface {
	mustEmbedUnimplementedKopiaRepositoryServer()
}

func RegisterKopiaRepositoryServer(s grpc.ServiceRegistrar, srv KopiaRepositoryServer) {
	s.RegisterService(&KopiaRepository_ServiceDesc, srv)
}

func _KopiaRepository_Session_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(KopiaRepositoryServer).Session(&kopiaRepositorySessionServer{stream})
}

type KopiaRepository_SessionServer interface {
	Send(*SessionResponse) error
	Recv() (*SessionRequest, error)
	grpc.ServerStream
}

type kopiaRepositorySessionServer struct {
	grpc.ServerStream
}

func (x *kopiaRepositorySessionServer) Send(m *SessionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *kopiaRepositorySessionServer) Recv() (*SessionRequest, error) {
	m := new(SessionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// KopiaRepository_ServiceDesc is the grpc.ServiceDesc for KopiaRepository service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KopiaRepository_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kopia_repository.KopiaRepository",
	HandlerType: (*KopiaRepositoryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Session",
			Handler:       _KopiaRepository_Session_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "repository_server.proto",
}

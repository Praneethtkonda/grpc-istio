// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: go_signc.proto

package pb

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
	SigningRequest_SendRequest_FullMethodName    = "/pb.SigningRequest/SendRequest"
	SigningRequest_UploadToServer_FullMethodName = "/pb.SigningRequest/UploadToServer"
)

// SigningRequestClient is the client API for SigningRequest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SigningRequestClient interface {
	// Sends a signing request
	SendRequest(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error)
	UploadToServer(ctx context.Context, opts ...grpc.CallOption) (SigningRequest_UploadToServerClient, error)
}

type signingRequestClient struct {
	cc grpc.ClientConnInterface
}

func NewSigningRequestClient(cc grpc.ClientConnInterface) SigningRequestClient {
	return &signingRequestClient{cc}
}

func (c *signingRequestClient) SendRequest(ctx context.Context, in *SignRequest, opts ...grpc.CallOption) (*SignResponse, error) {
	out := new(SignResponse)
	err := c.cc.Invoke(ctx, SigningRequest_SendRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *signingRequestClient) UploadToServer(ctx context.Context, opts ...grpc.CallOption) (SigningRequest_UploadToServerClient, error) {
	stream, err := c.cc.NewStream(ctx, &SigningRequest_ServiceDesc.Streams[0], SigningRequest_UploadToServer_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &signingRequestUploadToServerClient{stream}
	return x, nil
}

type SigningRequest_UploadToServerClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type signingRequestUploadToServerClient struct {
	grpc.ClientStream
}

func (x *signingRequestUploadToServerClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *signingRequestUploadToServerClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SigningRequestServer is the server API for SigningRequest service.
// All implementations must embed UnimplementedSigningRequestServer
// for forward compatibility
type SigningRequestServer interface {
	// Sends a signing request
	SendRequest(context.Context, *SignRequest) (*SignResponse, error)
	UploadToServer(SigningRequest_UploadToServerServer) error
	mustEmbedUnimplementedSigningRequestServer()
}

// UnimplementedSigningRequestServer must be embedded to have forward compatible implementations.
type UnimplementedSigningRequestServer struct {
}

func (UnimplementedSigningRequestServer) SendRequest(context.Context, *SignRequest) (*SignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRequest not implemented")
}
func (UnimplementedSigningRequestServer) UploadToServer(SigningRequest_UploadToServerServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadToServer not implemented")
}
func (UnimplementedSigningRequestServer) mustEmbedUnimplementedSigningRequestServer() {}

// UnsafeSigningRequestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SigningRequestServer will
// result in compilation errors.
type UnsafeSigningRequestServer interface {
	mustEmbedUnimplementedSigningRequestServer()
}

func RegisterSigningRequestServer(s grpc.ServiceRegistrar, srv SigningRequestServer) {
	s.RegisterService(&SigningRequest_ServiceDesc, srv)
}

func _SigningRequest_SendRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SigningRequestServer).SendRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SigningRequest_SendRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SigningRequestServer).SendRequest(ctx, req.(*SignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SigningRequest_UploadToServer_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SigningRequestServer).UploadToServer(&signingRequestUploadToServerServer{stream})
}

type SigningRequest_UploadToServerServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type signingRequestUploadToServerServer struct {
	grpc.ServerStream
}

func (x *signingRequestUploadToServerServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *signingRequestUploadToServerServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SigningRequest_ServiceDesc is the grpc.ServiceDesc for SigningRequest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SigningRequest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SigningRequest",
	HandlerType: (*SigningRequestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendRequest",
			Handler:    _SigningRequest_SendRequest_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadToServer",
			Handler:       _SigningRequest_UploadToServer_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "go_signc.proto",
}

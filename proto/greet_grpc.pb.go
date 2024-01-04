// for syntax highlighting we use proto3 version

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: proto/greet.proto

// package name for our proto file

package proto

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
	GreetService_Product_FullMethodName                       = "/greet_service.GreetService/Product"
	GreetService_ProductServerStreaming_FullMethodName        = "/greet_service.GreetService/ProductServerStreaming"
	GreetService_ProductClientStreaming_FullMethodName        = "/greet_service.GreetService/ProductClientStreaming"
	GreetService_ProductBidirectionalStreaming_FullMethodName = "/greet_service.GreetService/ProductBidirectionalStreaming"
)

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	// simple unary RPC
	Product(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*ProductResponse, error)
	// server streaming RPC
	ProductServerStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_ProductServerStreamingClient, error)
	// client streaming RPC
	ProductClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_ProductClientStreamingClient, error)
	// bidirectional streaming RPC
	ProductBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_ProductBidirectionalStreamingClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Product(ctx context.Context, in *NoParam, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, GreetService_Product_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greetServiceClient) ProductServerStreaming(ctx context.Context, in *NamesList, opts ...grpc.CallOption) (GreetService_ProductServerStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], GreetService_ProductServerStreaming_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceProductServerStreamingClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_ProductServerStreamingClient interface {
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type greetServiceProductServerStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceProductServerStreamingClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) ProductClientStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_ProductClientStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[1], GreetService_ProductClientStreaming_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceProductClientStreamingClient{stream}
	return x, nil
}

type GreetService_ProductClientStreamingClient interface {
	Send(*ProductRequest) error
	CloseAndRecv() (*MessagesList, error)
	grpc.ClientStream
}

type greetServiceProductClientStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceProductClientStreamingClient) Send(m *ProductRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceProductClientStreamingClient) CloseAndRecv() (*MessagesList, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessagesList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greetServiceClient) ProductBidirectionalStreaming(ctx context.Context, opts ...grpc.CallOption) (GreetService_ProductBidirectionalStreamingClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[2], GreetService_ProductBidirectionalStreaming_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceProductBidirectionalStreamingClient{stream}
	return x, nil
}

type GreetService_ProductBidirectionalStreamingClient interface {
	Send(*ProductRequest) error
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type greetServiceProductBidirectionalStreamingClient struct {
	grpc.ClientStream
}

func (x *greetServiceProductBidirectionalStreamingClient) Send(m *ProductRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greetServiceProductBidirectionalStreamingClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	// simple unary RPC
	Product(context.Context, *NoParam) (*ProductResponse, error)
	// server streaming RPC
	ProductServerStreaming(*NamesList, GreetService_ProductServerStreamingServer) error
	// client streaming RPC
	ProductClientStreaming(GreetService_ProductClientStreamingServer) error
	// bidirectional streaming RPC
	ProductBidirectionalStreaming(GreetService_ProductBidirectionalStreamingServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) Product(context.Context, *NoParam) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Product not implemented")
}
func (UnimplementedGreetServiceServer) ProductServerStreaming(*NamesList, GreetService_ProductServerStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ProductServerStreaming not implemented")
}
func (UnimplementedGreetServiceServer) ProductClientStreaming(GreetService_ProductClientStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ProductClientStreaming not implemented")
}
func (UnimplementedGreetServiceServer) ProductBidirectionalStreaming(GreetService_ProductBidirectionalStreamingServer) error {
	return status.Errorf(codes.Unimplemented, "method ProductBidirectionalStreaming not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_Product_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Product(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GreetService_Product_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Product(ctx, req.(*NoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _GreetService_ProductServerStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NamesList)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).ProductServerStreaming(m, &greetServiceProductServerStreamingServer{stream})
}

type GreetService_ProductServerStreamingServer interface {
	Send(*ProductResponse) error
	grpc.ServerStream
}

type greetServiceProductServerStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceProductServerStreamingServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GreetService_ProductClientStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).ProductClientStreaming(&greetServiceProductClientStreamingServer{stream})
}

type GreetService_ProductClientStreamingServer interface {
	SendAndClose(*MessagesList) error
	Recv() (*ProductRequest, error)
	grpc.ServerStream
}

type greetServiceProductClientStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceProductClientStreamingServer) SendAndClose(m *MessagesList) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceProductClientStreamingServer) Recv() (*ProductRequest, error) {
	m := new(ProductRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GreetService_ProductBidirectionalStreaming_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreetServiceServer).ProductBidirectionalStreaming(&greetServiceProductBidirectionalStreamingServer{stream})
}

type GreetService_ProductBidirectionalStreamingServer interface {
	Send(*ProductResponse) error
	Recv() (*ProductRequest, error)
	grpc.ServerStream
}

type greetServiceProductBidirectionalStreamingServer struct {
	grpc.ServerStream
}

func (x *greetServiceProductBidirectionalStreamingServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greetServiceProductBidirectionalStreamingServer) Recv() (*ProductRequest, error) {
	m := new(ProductRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet_service.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Product",
			Handler:    _GreetService_Product_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ProductServerStreaming",
			Handler:       _GreetService_ProductServerStreaming_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ProductClientStreaming",
			Handler:       _GreetService_ProductClientStreaming_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ProductBidirectionalStreaming",
			Handler:       _GreetService_ProductBidirectionalStreaming_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/greet.proto",
}

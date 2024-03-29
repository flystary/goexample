// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: query.proto

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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	GetAge(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*AgeInfo, error)
	GetName(ctx context.Context, in *AgeInfo, opts ...grpc.CallOption) (*UserInfo, error)
	Update(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Watch(ctx context.Context, in *WatchTime, opts ...grpc.CallOption) (Query_WatchClient, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetAge(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*AgeInfo, error) {
	out := new(AgeInfo)
	err := c.cc.Invoke(ctx, "/pb.Query/GetAge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetName(ctx context.Context, in *AgeInfo, opts ...grpc.CallOption) (*UserInfo, error) {
	out := new(UserInfo)
	err := c.cc.Invoke(ctx, "/pb.Query/GetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Update(ctx context.Context, in *UserInfo, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.Query/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Watch(ctx context.Context, in *WatchTime, opts ...grpc.CallOption) (Query_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Query_ServiceDesc.Streams[0], "/pb.Query/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Query_WatchClient interface {
	Recv() (*UserInfo, error)
	grpc.ClientStream
}

type queryWatchClient struct {
	grpc.ClientStream
}

func (x *queryWatchClient) Recv() (*UserInfo, error) {
	m := new(UserInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	GetAge(context.Context, *UserInfo) (*AgeInfo, error)
	GetName(context.Context, *AgeInfo) (*UserInfo, error)
	Update(context.Context, *UserInfo) (*emptypb.Empty, error)
	Watch(*WatchTime, Query_WatchServer) error
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) GetAge(context.Context, *UserInfo) (*AgeInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAge not implemented")
}
func (UnimplementedQueryServer) GetName(context.Context, *AgeInfo) (*UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetName not implemented")
}
func (UnimplementedQueryServer) Update(context.Context, *UserInfo) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedQueryServer) Watch(*WatchTime, Query_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_GetAge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Query/GetAge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAge(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgeInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Query/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetName(ctx, req.(*AgeInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Query/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Update(ctx, req.(*UserInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchTime)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServer).Watch(m, &queryWatchServer{stream})
}

type Query_WatchServer interface {
	Send(*UserInfo) error
	grpc.ServerStream
}

type queryWatchServer struct {
	grpc.ServerStream
}

func (x *queryWatchServer) Send(m *UserInfo) error {
	return x.ServerStream.SendMsg(m)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAge",
			Handler:    _Query_GetAge_Handler,
		},
		{
			MethodName: "GetName",
			Handler:    _Query_GetName_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Query_Update_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Query_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "query.proto",
}

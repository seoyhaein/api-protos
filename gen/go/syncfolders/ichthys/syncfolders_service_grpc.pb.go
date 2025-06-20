// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: ichthys/v1/syncfolders_service.proto

package syncfoldersv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SyncFoldersService_SyncFolders_FullMethodName = "/ichthys.SyncFoldersService/SyncFolders"
)

// SyncFoldersServiceClient is the client API for SyncFoldersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SyncFoldersServiceClient interface {
	// 클라이언트의 요청에 따라 서버의 폴더와 DB를 비교한 후, 업데이트가 필요한 경우 수행하고 결과를 반환
	SyncFolders(ctx context.Context, in *SyncFoldersRequest, opts ...grpc.CallOption) (*SyncFoldersResponse, error)
}

type syncFoldersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSyncFoldersServiceClient(cc grpc.ClientConnInterface) SyncFoldersServiceClient {
	return &syncFoldersServiceClient{cc}
}

func (c *syncFoldersServiceClient) SyncFolders(ctx context.Context, in *SyncFoldersRequest, opts ...grpc.CallOption) (*SyncFoldersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SyncFoldersResponse)
	err := c.cc.Invoke(ctx, SyncFoldersService_SyncFolders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SyncFoldersServiceServer is the server API for SyncFoldersService service.
// All implementations must embed UnimplementedSyncFoldersServiceServer
// for forward compatibility.
type SyncFoldersServiceServer interface {
	// 클라이언트의 요청에 따라 서버의 폴더와 DB를 비교한 후, 업데이트가 필요한 경우 수행하고 결과를 반환
	SyncFolders(context.Context, *SyncFoldersRequest) (*SyncFoldersResponse, error)
	mustEmbedUnimplementedSyncFoldersServiceServer()
}

// UnimplementedSyncFoldersServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSyncFoldersServiceServer struct{}

func (UnimplementedSyncFoldersServiceServer) SyncFolders(context.Context, *SyncFoldersRequest) (*SyncFoldersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SyncFolders not implemented")
}
func (UnimplementedSyncFoldersServiceServer) mustEmbedUnimplementedSyncFoldersServiceServer() {}
func (UnimplementedSyncFoldersServiceServer) testEmbeddedByValue()                            {}

// UnsafeSyncFoldersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SyncFoldersServiceServer will
// result in compilation errors.
type UnsafeSyncFoldersServiceServer interface {
	mustEmbedUnimplementedSyncFoldersServiceServer()
}

func RegisterSyncFoldersServiceServer(s grpc.ServiceRegistrar, srv SyncFoldersServiceServer) {
	// If the following call pancis, it indicates UnimplementedSyncFoldersServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SyncFoldersService_ServiceDesc, srv)
}

func _SyncFoldersService_SyncFolders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncFoldersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SyncFoldersServiceServer).SyncFolders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SyncFoldersService_SyncFolders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SyncFoldersServiceServer).SyncFolders(ctx, req.(*SyncFoldersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SyncFoldersService_ServiceDesc is the grpc.ServiceDesc for SyncFoldersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SyncFoldersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ichthys.SyncFoldersService",
	HandlerType: (*SyncFoldersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SyncFolders",
			Handler:    _SyncFoldersService_SyncFolders_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ichthys/v1/syncfolders_service.proto",
}

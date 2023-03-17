// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: link.proto

package links

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

// LinksServiceClient is the client API for LinksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinksServiceClient interface {
	Create(ctx context.Context, in *CreateShortLinkRequest, opts ...grpc.CallOption) (*CreateShortLinkResponse, error)
	Retrive(ctx context.Context, in *RetriveOriginalLinkRequest, opts ...grpc.CallOption) (*RetriveOriginalLinkResponse, error)
}

type linksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLinksServiceClient(cc grpc.ClientConnInterface) LinksServiceClient {
	return &linksServiceClient{cc}
}

func (c *linksServiceClient) Create(ctx context.Context, in *CreateShortLinkRequest, opts ...grpc.CallOption) (*CreateShortLinkResponse, error) {
	out := new(CreateShortLinkResponse)
	err := c.cc.Invoke(ctx, "/links.LinksService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linksServiceClient) Retrive(ctx context.Context, in *RetriveOriginalLinkRequest, opts ...grpc.CallOption) (*RetriveOriginalLinkResponse, error) {
	out := new(RetriveOriginalLinkResponse)
	err := c.cc.Invoke(ctx, "/links.LinksService/Retrive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinksServiceServer is the server API for LinksService service.
// All implementations should embed UnimplementedLinksServiceServer
// for forward compatibility
type LinksServiceServer interface {
	Create(context.Context, *CreateShortLinkRequest) (*CreateShortLinkResponse, error)
	Retrive(context.Context, *RetriveOriginalLinkRequest) (*RetriveOriginalLinkResponse, error)
}

// UnimplementedLinksServiceServer should be embedded to have forward compatible implementations.
type UnimplementedLinksServiceServer struct {
}

func (UnimplementedLinksServiceServer) Create(context.Context, *CreateShortLinkRequest) (*CreateShortLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLinksServiceServer) Retrive(context.Context, *RetriveOriginalLinkRequest) (*RetriveOriginalLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrive not implemented")
}

// UnsafeLinksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinksServiceServer will
// result in compilation errors.
type UnsafeLinksServiceServer interface {
	mustEmbedUnimplementedLinksServiceServer()
}

func RegisterLinksServiceServer(s grpc.ServiceRegistrar, srv LinksServiceServer) {
	s.RegisterService(&LinksService_ServiceDesc, srv)
}

func _LinksService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateShortLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinksServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/links.LinksService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinksServiceServer).Create(ctx, req.(*CreateShortLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinksService_Retrive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetriveOriginalLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinksServiceServer).Retrive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/links.LinksService/Retrive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinksServiceServer).Retrive(ctx, req.(*RetriveOriginalLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinksService_ServiceDesc is the grpc.ServiceDesc for LinksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "links.LinksService",
	HandlerType: (*LinksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LinksService_Create_Handler,
		},
		{
			MethodName: "Retrive",
			Handler:    _LinksService_Retrive_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "link.proto",
}
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package posting_api_grpc

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

// PostingClient is the client API for Posting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostingClient interface {
	CreatePosting(ctx context.Context, in *PostingRequest, opts ...grpc.CallOption) (*PostingReply, error)
}

type postingClient struct {
	cc grpc.ClientConnInterface
}

func NewPostingClient(cc grpc.ClientConnInterface) PostingClient {
	return &postingClient{cc}
}

func (c *postingClient) CreatePosting(ctx context.Context, in *PostingRequest, opts ...grpc.CallOption) (*PostingReply, error) {
	out := new(PostingReply)
	err := c.cc.Invoke(ctx, "/posting_api_grpc.Posting/CreatePosting", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostingServer is the server API for Posting service.
// All implementations must embed UnimplementedPostingServer
// for forward compatibility
type PostingServer interface {
	CreatePosting(context.Context, *PostingRequest) (*PostingReply, error)
	mustEmbedUnimplementedPostingServer()
}

// UnimplementedPostingServer must be embedded to have forward compatible implementations.
type UnimplementedPostingServer struct {
}

func (UnimplementedPostingServer) CreatePosting(context.Context, *PostingRequest) (*PostingReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePosting not implemented")
}
func (UnimplementedPostingServer) mustEmbedUnimplementedPostingServer() {}

// UnsafePostingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostingServer will
// result in compilation errors.
type UnsafePostingServer interface {
	mustEmbedUnimplementedPostingServer()
}

func RegisterPostingServer(s grpc.ServiceRegistrar, srv PostingServer) {
	s.RegisterService(&Posting_ServiceDesc, srv)
}

func _Posting_CreatePosting_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostingServer).CreatePosting(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/posting_api_grpc.Posting/CreatePosting",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostingServer).CreatePosting(ctx, req.(*PostingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Posting_ServiceDesc is the grpc.ServiceDesc for Posting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Posting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "posting_api_grpc.Posting",
	HandlerType: (*PostingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePosting",
			Handler:    _Posting_CreatePosting_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "posting_api_grpc/posting.proto",
}
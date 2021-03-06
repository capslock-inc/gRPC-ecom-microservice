// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package mongoclientmodel

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

// MongoClientServiceClient is the client API for MongoClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MongoClientServiceClient interface {
	GetCartItemByUserId(ctx context.Context, in *Userid, opts ...grpc.CallOption) (*Cart, error)
	CreateNewCart(ctx context.Context, in *Userid, opts ...grpc.CallOption) (*Status, error)
	AddProductToCart(ctx context.Context, in *Addproduct, opts ...grpc.CallOption) (*Status, error)
	DeleteCartByUserId(ctx context.Context, in *Userid, opts ...grpc.CallOption) (*Status, error)
	DeleteCartItemByProductId(ctx context.Context, in *Addproduct, opts ...grpc.CallOption) (*Status, error)
}

type mongoClientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMongoClientServiceClient(cc grpc.ClientConnInterface) MongoClientServiceClient {
	return &mongoClientServiceClient{cc}
}

func (c *mongoClientServiceClient) GetCartItemByUserId(ctx context.Context, in *Userid, opts ...grpc.CallOption) (*Cart, error) {
	out := new(Cart)
	err := c.cc.Invoke(ctx, "/mongoclientmodel.MongoClientService/GetCartItemByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoClientServiceClient) CreateNewCart(ctx context.Context, in *Userid, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/mongoclientmodel.MongoClientService/CreateNewCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoClientServiceClient) AddProductToCart(ctx context.Context, in *Addproduct, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/mongoclientmodel.MongoClientService/AddProductToCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoClientServiceClient) DeleteCartByUserId(ctx context.Context, in *Userid, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/mongoclientmodel.MongoClientService/DeleteCartByUserId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mongoClientServiceClient) DeleteCartItemByProductId(ctx context.Context, in *Addproduct, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/mongoclientmodel.MongoClientService/DeleteCartItemByProductId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongoClientServiceServer is the server API for MongoClientService service.
// All implementations must embed UnimplementedMongoClientServiceServer
// for forward compatibility
type MongoClientServiceServer interface {
	GetCartItemByUserId(context.Context, *Userid) (*Cart, error)
	CreateNewCart(context.Context, *Userid) (*Status, error)
	AddProductToCart(context.Context, *Addproduct) (*Status, error)
	DeleteCartByUserId(context.Context, *Userid) (*Status, error)
	DeleteCartItemByProductId(context.Context, *Addproduct) (*Status, error)
	mustEmbedUnimplementedMongoClientServiceServer()
}

// UnimplementedMongoClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMongoClientServiceServer struct {
}

func (UnimplementedMongoClientServiceServer) GetCartItemByUserId(context.Context, *Userid) (*Cart, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartItemByUserId not implemented")
}
func (UnimplementedMongoClientServiceServer) CreateNewCart(context.Context, *Userid) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewCart not implemented")
}
func (UnimplementedMongoClientServiceServer) AddProductToCart(context.Context, *Addproduct) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductToCart not implemented")
}
func (UnimplementedMongoClientServiceServer) DeleteCartByUserId(context.Context, *Userid) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCartByUserId not implemented")
}
func (UnimplementedMongoClientServiceServer) DeleteCartItemByProductId(context.Context, *Addproduct) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCartItemByProductId not implemented")
}
func (UnimplementedMongoClientServiceServer) mustEmbedUnimplementedMongoClientServiceServer() {}

// UnsafeMongoClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MongoClientServiceServer will
// result in compilation errors.
type UnsafeMongoClientServiceServer interface {
	mustEmbedUnimplementedMongoClientServiceServer()
}

func RegisterMongoClientServiceServer(s grpc.ServiceRegistrar, srv MongoClientServiceServer) {
	s.RegisterService(&MongoClientService_ServiceDesc, srv)
}

func _MongoClientService_GetCartItemByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Userid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoClientServiceServer).GetCartItemByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongoclientmodel.MongoClientService/GetCartItemByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoClientServiceServer).GetCartItemByUserId(ctx, req.(*Userid))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoClientService_CreateNewCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Userid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoClientServiceServer).CreateNewCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongoclientmodel.MongoClientService/CreateNewCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoClientServiceServer).CreateNewCart(ctx, req.(*Userid))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoClientService_AddProductToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Addproduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoClientServiceServer).AddProductToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongoclientmodel.MongoClientService/AddProductToCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoClientServiceServer).AddProductToCart(ctx, req.(*Addproduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoClientService_DeleteCartByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Userid)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoClientServiceServer).DeleteCartByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongoclientmodel.MongoClientService/DeleteCartByUserId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoClientServiceServer).DeleteCartByUserId(ctx, req.(*Userid))
	}
	return interceptor(ctx, in, info, handler)
}

func _MongoClientService_DeleteCartItemByProductId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Addproduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoClientServiceServer).DeleteCartItemByProductId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mongoclientmodel.MongoClientService/DeleteCartItemByProductId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoClientServiceServer).DeleteCartItemByProductId(ctx, req.(*Addproduct))
	}
	return interceptor(ctx, in, info, handler)
}

// MongoClientService_ServiceDesc is the grpc.ServiceDesc for MongoClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MongoClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mongoclientmodel.MongoClientService",
	HandlerType: (*MongoClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCartItemByUserId",
			Handler:    _MongoClientService_GetCartItemByUserId_Handler,
		},
		{
			MethodName: "CreateNewCart",
			Handler:    _MongoClientService_CreateNewCart_Handler,
		},
		{
			MethodName: "AddProductToCart",
			Handler:    _MongoClientService_AddProductToCart_Handler,
		},
		{
			MethodName: "DeleteCartByUserId",
			Handler:    _MongoClientService_DeleteCartByUserId_Handler,
		},
		{
			MethodName: "DeleteCartItemByProductId",
			Handler:    _MongoClientService_DeleteCartItemByProductId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Protos/database/mongoclient/mongoclient.proto",
}

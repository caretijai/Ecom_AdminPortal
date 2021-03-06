// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protofiles

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ProductCatalogueServiceClient is the client API for ProductCatalogueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductCatalogueServiceClient interface {
	GetAllCategories(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error)
	CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error)
	EditCategory(ctx context.Context, in *EditCategoryRequest, opts ...grpc.CallOption) (*EditCategoryResponse, error)
	RemoveCategory(ctx context.Context, in *RemoveCategoryRequest, opts ...grpc.CallOption) (*RemoveCategoryResponse, error)
	EditSubCategory(ctx context.Context, in *EditSubCategoryRequest, opts ...grpc.CallOption) (*EditSubCategoryResponse, error)
	RemoveSubCategory(ctx context.Context, in *RemoveSubCategoryRequest, opts ...grpc.CallOption) (*RemoveSubCategoryResponse, error)
}

type productCatalogueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductCatalogueServiceClient(cc grpc.ClientConnInterface) ProductCatalogueServiceClient {
	return &productCatalogueServiceClient{cc}
}

func (c *productCatalogueServiceClient) GetAllCategories(ctx context.Context, in *GetAllCategoriesRequest, opts ...grpc.CallOption) (*GetAllCategoriesResponse, error) {
	out := new(GetAllCategoriesResponse)
	err := c.cc.Invoke(ctx, "/protofiles.ProductCatalogueService/GetAllCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogueServiceClient) CreateCategory(ctx context.Context, in *CreateCategoryRequest, opts ...grpc.CallOption) (*CreateCategoryResponse, error) {
	out := new(CreateCategoryResponse)
	err := c.cc.Invoke(ctx, "/protofiles.ProductCatalogueService/CreateCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogueServiceClient) EditCategory(ctx context.Context, in *EditCategoryRequest, opts ...grpc.CallOption) (*EditCategoryResponse, error) {
	out := new(EditCategoryResponse)
	err := c.cc.Invoke(ctx, "/protofiles.ProductCatalogueService/EditCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogueServiceClient) RemoveCategory(ctx context.Context, in *RemoveCategoryRequest, opts ...grpc.CallOption) (*RemoveCategoryResponse, error) {
	out := new(RemoveCategoryResponse)
	err := c.cc.Invoke(ctx, "/protofiles.ProductCatalogueService/RemoveCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogueServiceClient) EditSubCategory(ctx context.Context, in *EditSubCategoryRequest, opts ...grpc.CallOption) (*EditSubCategoryResponse, error) {
	out := new(EditSubCategoryResponse)
	err := c.cc.Invoke(ctx, "/protofiles.ProductCatalogueService/EditSubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productCatalogueServiceClient) RemoveSubCategory(ctx context.Context, in *RemoveSubCategoryRequest, opts ...grpc.CallOption) (*RemoveSubCategoryResponse, error) {
	out := new(RemoveSubCategoryResponse)
	err := c.cc.Invoke(ctx, "/protofiles.ProductCatalogueService/RemoveSubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductCatalogueServiceServer is the server API for ProductCatalogueService service.
// All implementations must embed UnimplementedProductCatalogueServiceServer
// for forward compatibility
type ProductCatalogueServiceServer interface {
	GetAllCategories(context.Context, *GetAllCategoriesRequest) (*GetAllCategoriesResponse, error)
	CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error)
	EditCategory(context.Context, *EditCategoryRequest) (*EditCategoryResponse, error)
	RemoveCategory(context.Context, *RemoveCategoryRequest) (*RemoveCategoryResponse, error)
	EditSubCategory(context.Context, *EditSubCategoryRequest) (*EditSubCategoryResponse, error)
	RemoveSubCategory(context.Context, *RemoveSubCategoryRequest) (*RemoveSubCategoryResponse, error)
	mustEmbedUnimplementedProductCatalogueServiceServer()
}

// UnimplementedProductCatalogueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductCatalogueServiceServer struct {
}

func (UnimplementedProductCatalogueServiceServer) GetAllCategories(context.Context, *GetAllCategoriesRequest) (*GetAllCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllCategories not implemented")
}
func (UnimplementedProductCatalogueServiceServer) CreateCategory(context.Context, *CreateCategoryRequest) (*CreateCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCategory not implemented")
}
func (UnimplementedProductCatalogueServiceServer) EditCategory(context.Context, *EditCategoryRequest) (*EditCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditCategory not implemented")
}
func (UnimplementedProductCatalogueServiceServer) RemoveCategory(context.Context, *RemoveCategoryRequest) (*RemoveCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCategory not implemented")
}
func (UnimplementedProductCatalogueServiceServer) EditSubCategory(context.Context, *EditSubCategoryRequest) (*EditSubCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditSubCategory not implemented")
}
func (UnimplementedProductCatalogueServiceServer) RemoveSubCategory(context.Context, *RemoveSubCategoryRequest) (*RemoveSubCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSubCategory not implemented")
}
func (UnimplementedProductCatalogueServiceServer) mustEmbedUnimplementedProductCatalogueServiceServer() {
}

// UnsafeProductCatalogueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductCatalogueServiceServer will
// result in compilation errors.
type UnsafeProductCatalogueServiceServer interface {
	mustEmbedUnimplementedProductCatalogueServiceServer()
}

func RegisterProductCatalogueServiceServer(s *grpc.Server, srv ProductCatalogueServiceServer) {
	s.RegisterService(&_ProductCatalogueService_serviceDesc, srv)
}

func _ProductCatalogueService_GetAllCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogueServiceServer).GetAllCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.ProductCatalogueService/GetAllCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogueServiceServer).GetAllCategories(ctx, req.(*GetAllCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogueService_CreateCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogueServiceServer).CreateCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.ProductCatalogueService/CreateCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogueServiceServer).CreateCategory(ctx, req.(*CreateCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogueService_EditCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogueServiceServer).EditCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.ProductCatalogueService/EditCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogueServiceServer).EditCategory(ctx, req.(*EditCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogueService_RemoveCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogueServiceServer).RemoveCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.ProductCatalogueService/RemoveCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogueServiceServer).RemoveCategory(ctx, req.(*RemoveCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogueService_EditSubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditSubCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogueServiceServer).EditSubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.ProductCatalogueService/EditSubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogueServiceServer).EditSubCategory(ctx, req.(*EditSubCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductCatalogueService_RemoveSubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSubCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductCatalogueServiceServer).RemoveSubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.ProductCatalogueService/RemoveSubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductCatalogueServiceServer).RemoveSubCategory(ctx, req.(*RemoveSubCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProductCatalogueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.ProductCatalogueService",
	HandlerType: (*ProductCatalogueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllCategories",
			Handler:    _ProductCatalogueService_GetAllCategories_Handler,
		},
		{
			MethodName: "CreateCategory",
			Handler:    _ProductCatalogueService_CreateCategory_Handler,
		},
		{
			MethodName: "EditCategory",
			Handler:    _ProductCatalogueService_EditCategory_Handler,
		},
		{
			MethodName: "RemoveCategory",
			Handler:    _ProductCatalogueService_RemoveCategory_Handler,
		},
		{
			MethodName: "EditSubCategory",
			Handler:    _ProductCatalogueService_EditSubCategory_Handler,
		},
		{
			MethodName: "RemoveSubCategory",
			Handler:    _ProductCatalogueService_RemoveSubCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "product.proto",
}

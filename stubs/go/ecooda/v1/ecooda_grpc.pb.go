// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: ecooda/v1/ecooda.proto

package v1

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
	LocationService_GetLocations_FullMethodName = "/ecooda.v1.LocationService/GetLocations"
	LocationService_GetLocation_FullMethodName  = "/ecooda.v1.LocationService/GetLocation"
	LocationService_PostLocation_FullMethodName = "/ecooda.v1.LocationService/PostLocation"
)

// LocationServiceClient is the client API for LocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationServiceClient interface {
	GetLocations(ctx context.Context, in *GetLocationsRequest, opts ...grpc.CallOption) (*GetLocationsResponse, error)
	GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error)
	PostLocation(ctx context.Context, in *PostLocationRequest, opts ...grpc.CallOption) (*PostLocationResponse, error)
}

type locationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationServiceClient(cc grpc.ClientConnInterface) LocationServiceClient {
	return &locationServiceClient{cc}
}

func (c *locationServiceClient) GetLocations(ctx context.Context, in *GetLocationsRequest, opts ...grpc.CallOption) (*GetLocationsResponse, error) {
	out := new(GetLocationsResponse)
	err := c.cc.Invoke(ctx, LocationService_GetLocations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error) {
	out := new(GetLocationResponse)
	err := c.cc.Invoke(ctx, LocationService_GetLocation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) PostLocation(ctx context.Context, in *PostLocationRequest, opts ...grpc.CallOption) (*PostLocationResponse, error) {
	out := new(PostLocationResponse)
	err := c.cc.Invoke(ctx, LocationService_PostLocation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationServiceServer is the server API for LocationService service.
// All implementations must embed UnimplementedLocationServiceServer
// for forward compatibility
type LocationServiceServer interface {
	GetLocations(context.Context, *GetLocationsRequest) (*GetLocationsResponse, error)
	GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error)
	PostLocation(context.Context, *PostLocationRequest) (*PostLocationResponse, error)
	mustEmbedUnimplementedLocationServiceServer()
}

// UnimplementedLocationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocationServiceServer struct {
}

func (UnimplementedLocationServiceServer) GetLocations(context.Context, *GetLocationsRequest) (*GetLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocations not implemented")
}
func (UnimplementedLocationServiceServer) GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedLocationServiceServer) PostLocation(context.Context, *PostLocationRequest) (*PostLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostLocation not implemented")
}
func (UnimplementedLocationServiceServer) mustEmbedUnimplementedLocationServiceServer() {}

// UnsafeLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationServiceServer will
// result in compilation errors.
type UnsafeLocationServiceServer interface {
	mustEmbedUnimplementedLocationServiceServer()
}

func RegisterLocationServiceServer(s grpc.ServiceRegistrar, srv LocationServiceServer) {
	s.RegisterService(&LocationService_ServiceDesc, srv)
}

func _LocationService_GetLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).GetLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationService_GetLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).GetLocations(ctx, req.(*GetLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationService_GetLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).GetLocation(ctx, req.(*GetLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_PostLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).PostLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LocationService_PostLocation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).PostLocation(ctx, req.(*PostLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationService_ServiceDesc is the grpc.ServiceDesc for LocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecooda.v1.LocationService",
	HandlerType: (*LocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLocations",
			Handler:    _LocationService_GetLocations_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _LocationService_GetLocation_Handler,
		},
		{
			MethodName: "PostLocation",
			Handler:    _LocationService_PostLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ecooda/v1/ecooda.proto",
}

const (
	ChallengeService_GetChallenges_FullMethodName    = "/ecooda.v1.ChallengeService/GetChallenges"
	ChallengeService_GetChallenge_FullMethodName     = "/ecooda.v1.ChallengeService/GetChallenge"
	ChallengeService_PostChallenge_FullMethodName    = "/ecooda.v1.ChallengeService/PostChallenge"
	ChallengeService_PutChallenge_FullMethodName     = "/ecooda.v1.ChallengeService/PutChallenge"
	ChallengeService_DeleteChallenge_FullMethodName  = "/ecooda.v1.ChallengeService/DeleteChallenge"
	ChallengeService_CommandChallenge_FullMethodName = "/ecooda.v1.ChallengeService/CommandChallenge"
	ChallengeService_GetCategories_FullMethodName    = "/ecooda.v1.ChallengeService/GetCategories"
	ChallengeService_GetCategory_FullMethodName      = "/ecooda.v1.ChallengeService/GetCategory"
	ChallengeService_PostCategory_FullMethodName     = "/ecooda.v1.ChallengeService/PostCategory"
	ChallengeService_PutCategory_FullMethodName      = "/ecooda.v1.ChallengeService/PutCategory"
	ChallengeService_DeleteCategory_FullMethodName   = "/ecooda.v1.ChallengeService/DeleteCategory"
)

// ChallengeServiceClient is the client API for ChallengeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChallengeServiceClient interface {
	GetChallenges(ctx context.Context, in *GetChallengesRequest, opts ...grpc.CallOption) (*GetChallengesResponse, error)
	GetChallenge(ctx context.Context, in *GetChallengeRequest, opts ...grpc.CallOption) (*GetChallengeResponse, error)
	PostChallenge(ctx context.Context, in *PostChallengeRequest, opts ...grpc.CallOption) (*PostChallengeResponse, error)
	PutChallenge(ctx context.Context, in *PutChallengeRequest, opts ...grpc.CallOption) (*PutChallengeResponse, error)
	DeleteChallenge(ctx context.Context, in *DeleteChallengeRequest, opts ...grpc.CallOption) (*DeleteChallengeResponse, error)
	CommandChallenge(ctx context.Context, in *CommandChallengeRequest, opts ...grpc.CallOption) (*CommandChallengeResponse, error)
	GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error)
	PostCategory(ctx context.Context, in *PostCategoryRequest, opts ...grpc.CallOption) (*PostCategoryResponse, error)
	PutCategory(ctx context.Context, in *PutCategoryRequest, opts ...grpc.CallOption) (*PutCategoryResponse, error)
	DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error)
}

type challengeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChallengeServiceClient(cc grpc.ClientConnInterface) ChallengeServiceClient {
	return &challengeServiceClient{cc}
}

func (c *challengeServiceClient) GetChallenges(ctx context.Context, in *GetChallengesRequest, opts ...grpc.CallOption) (*GetChallengesResponse, error) {
	out := new(GetChallengesResponse)
	err := c.cc.Invoke(ctx, ChallengeService_GetChallenges_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *challengeServiceClient) GetChallenge(ctx context.Context, in *GetChallengeRequest, opts ...grpc.CallOption) (*GetChallengeResponse, error) {
	out := new(GetChallengeResponse)
	err := c.cc.Invoke(ctx, ChallengeService_GetChallenge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *challengeServiceClient) PostChallenge(ctx context.Context, in *PostChallengeRequest, opts ...grpc.CallOption) (*PostChallengeResponse, error) {
	out := new(PostChallengeResponse)
	err := c.cc.Invoke(ctx, ChallengeService_PostChallenge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *challengeServiceClient) PutChallenge(ctx context.Context, in *PutChallengeRequest, opts ...grpc.CallOption) (*PutChallengeResponse, error) {
	out := new(PutChallengeResponse)
	err := c.cc.Invoke(ctx, ChallengeService_PutChallenge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *challengeServiceClient) DeleteChallenge(ctx context.Context, in *DeleteChallengeRequest, opts ...grpc.CallOption) (*DeleteChallengeResponse, error) {
	out := new(DeleteChallengeResponse)
	err := c.cc.Invoke(ctx, ChallengeService_DeleteChallenge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *challengeServiceClient) CommandChallenge(ctx context.Context, in *CommandChallengeRequest, opts ...grpc.CallOption) (*CommandChallengeResponse, error) {
	out := new(CommandChallengeResponse)
	err := c.cc.Invoke(ctx, ChallengeService_CommandChallenge_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *challengeServiceClient) GetCategories(ctx context.Context, in *GetCategoriesRequest, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, ChallengeService_GetCategories_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *challengeServiceClient) GetCategory(ctx context.Context, in *GetCategoryRequest, opts ...grpc.CallOption) (*GetCategoryResponse, error) {
	out := new(GetCategoryResponse)
	err := c.cc.Invoke(ctx, ChallengeService_GetCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *challengeServiceClient) PostCategory(ctx context.Context, in *PostCategoryRequest, opts ...grpc.CallOption) (*PostCategoryResponse, error) {
	out := new(PostCategoryResponse)
	err := c.cc.Invoke(ctx, ChallengeService_PostCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *challengeServiceClient) PutCategory(ctx context.Context, in *PutCategoryRequest, opts ...grpc.CallOption) (*PutCategoryResponse, error) {
	out := new(PutCategoryResponse)
	err := c.cc.Invoke(ctx, ChallengeService_PutCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *challengeServiceClient) DeleteCategory(ctx context.Context, in *DeleteCategoryRequest, opts ...grpc.CallOption) (*DeleteCategoryResponse, error) {
	out := new(DeleteCategoryResponse)
	err := c.cc.Invoke(ctx, ChallengeService_DeleteCategory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChallengeServiceServer is the server API for ChallengeService service.
// All implementations must embed UnimplementedChallengeServiceServer
// for forward compatibility
type ChallengeServiceServer interface {
	GetChallenges(context.Context, *GetChallengesRequest) (*GetChallengesResponse, error)
	GetChallenge(context.Context, *GetChallengeRequest) (*GetChallengeResponse, error)
	PostChallenge(context.Context, *PostChallengeRequest) (*PostChallengeResponse, error)
	PutChallenge(context.Context, *PutChallengeRequest) (*PutChallengeResponse, error)
	DeleteChallenge(context.Context, *DeleteChallengeRequest) (*DeleteChallengeResponse, error)
	CommandChallenge(context.Context, *CommandChallengeRequest) (*CommandChallengeResponse, error)
	GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error)
	GetCategory(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error)
	PostCategory(context.Context, *PostCategoryRequest) (*PostCategoryResponse, error)
	PutCategory(context.Context, *PutCategoryRequest) (*PutCategoryResponse, error)
	DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error)
	mustEmbedUnimplementedChallengeServiceServer()
}

// UnimplementedChallengeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChallengeServiceServer struct {
}

func (UnimplementedChallengeServiceServer) GetChallenges(context.Context, *GetChallengesRequest) (*GetChallengesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChallenges not implemented")
}
func (UnimplementedChallengeServiceServer) GetChallenge(context.Context, *GetChallengeRequest) (*GetChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChallenge not implemented")
}
func (UnimplementedChallengeServiceServer) PostChallenge(context.Context, *PostChallengeRequest) (*PostChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostChallenge not implemented")
}
func (UnimplementedChallengeServiceServer) PutChallenge(context.Context, *PutChallengeRequest) (*PutChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutChallenge not implemented")
}
func (UnimplementedChallengeServiceServer) DeleteChallenge(context.Context, *DeleteChallengeRequest) (*DeleteChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteChallenge not implemented")
}
func (UnimplementedChallengeServiceServer) CommandChallenge(context.Context, *CommandChallengeRequest) (*CommandChallengeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommandChallenge not implemented")
}
func (UnimplementedChallengeServiceServer) GetCategories(context.Context, *GetCategoriesRequest) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (UnimplementedChallengeServiceServer) GetCategory(context.Context, *GetCategoryRequest) (*GetCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategory not implemented")
}
func (UnimplementedChallengeServiceServer) PostCategory(context.Context, *PostCategoryRequest) (*PostCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostCategory not implemented")
}
func (UnimplementedChallengeServiceServer) PutCategory(context.Context, *PutCategoryRequest) (*PutCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutCategory not implemented")
}
func (UnimplementedChallengeServiceServer) DeleteCategory(context.Context, *DeleteCategoryRequest) (*DeleteCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCategory not implemented")
}
func (UnimplementedChallengeServiceServer) mustEmbedUnimplementedChallengeServiceServer() {}

// UnsafeChallengeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChallengeServiceServer will
// result in compilation errors.
type UnsafeChallengeServiceServer interface {
	mustEmbedUnimplementedChallengeServiceServer()
}

func RegisterChallengeServiceServer(s grpc.ServiceRegistrar, srv ChallengeServiceServer) {
	s.RegisterService(&ChallengeService_ServiceDesc, srv)
}

func _ChallengeService_GetChallenges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChallengesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).GetChallenges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChallengeService_GetChallenges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).GetChallenges(ctx, req.(*GetChallengesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChallengeService_GetChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).GetChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChallengeService_GetChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).GetChallenge(ctx, req.(*GetChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChallengeService_PostChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).PostChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChallengeService_PostChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).PostChallenge(ctx, req.(*PostChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChallengeService_PutChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).PutChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChallengeService_PutChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).PutChallenge(ctx, req.(*PutChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChallengeService_DeleteChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).DeleteChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChallengeService_DeleteChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).DeleteChallenge(ctx, req.(*DeleteChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChallengeService_CommandChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).CommandChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChallengeService_CommandChallenge_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).CommandChallenge(ctx, req.(*CommandChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChallengeService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChallengeService_GetCategories_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).GetCategories(ctx, req.(*GetCategoriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChallengeService_GetCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).GetCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChallengeService_GetCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).GetCategory(ctx, req.(*GetCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChallengeService_PostCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).PostCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChallengeService_PostCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).PostCategory(ctx, req.(*PostCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChallengeService_PutCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).PutCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChallengeService_PutCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).PutCategory(ctx, req.(*PutCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChallengeService_DeleteCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChallengeServiceServer).DeleteCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChallengeService_DeleteCategory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChallengeServiceServer).DeleteCategory(ctx, req.(*DeleteCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChallengeService_ServiceDesc is the grpc.ServiceDesc for ChallengeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChallengeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ecooda.v1.ChallengeService",
	HandlerType: (*ChallengeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChallenges",
			Handler:    _ChallengeService_GetChallenges_Handler,
		},
		{
			MethodName: "GetChallenge",
			Handler:    _ChallengeService_GetChallenge_Handler,
		},
		{
			MethodName: "PostChallenge",
			Handler:    _ChallengeService_PostChallenge_Handler,
		},
		{
			MethodName: "PutChallenge",
			Handler:    _ChallengeService_PutChallenge_Handler,
		},
		{
			MethodName: "DeleteChallenge",
			Handler:    _ChallengeService_DeleteChallenge_Handler,
		},
		{
			MethodName: "CommandChallenge",
			Handler:    _ChallengeService_CommandChallenge_Handler,
		},
		{
			MethodName: "GetCategories",
			Handler:    _ChallengeService_GetCategories_Handler,
		},
		{
			MethodName: "GetCategory",
			Handler:    _ChallengeService_GetCategory_Handler,
		},
		{
			MethodName: "PostCategory",
			Handler:    _ChallengeService_PostCategory_Handler,
		},
		{
			MethodName: "PutCategory",
			Handler:    _ChallengeService_PutCategory_Handler,
		},
		{
			MethodName: "DeleteCategory",
			Handler:    _ChallengeService_DeleteCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ecooda/v1/ecooda.proto",
}

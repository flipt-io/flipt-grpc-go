// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: auth/auth.proto

package auth

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

const (
	PublicAuthenticationService_ListAuthenticationMethods_FullMethodName = "/flipt.auth.PublicAuthenticationService/ListAuthenticationMethods"
)

// PublicAuthenticationServiceClient is the client API for PublicAuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PublicAuthenticationServiceClient interface {
	ListAuthenticationMethods(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListAuthenticationMethodsResponse, error)
}

type publicAuthenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicAuthenticationServiceClient(cc grpc.ClientConnInterface) PublicAuthenticationServiceClient {
	return &publicAuthenticationServiceClient{cc}
}

func (c *publicAuthenticationServiceClient) ListAuthenticationMethods(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListAuthenticationMethodsResponse, error) {
	out := new(ListAuthenticationMethodsResponse)
	err := c.cc.Invoke(ctx, PublicAuthenticationService_ListAuthenticationMethods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicAuthenticationServiceServer is the server API for PublicAuthenticationService service.
// All implementations must embed UnimplementedPublicAuthenticationServiceServer
// for forward compatibility
type PublicAuthenticationServiceServer interface {
	ListAuthenticationMethods(context.Context, *emptypb.Empty) (*ListAuthenticationMethodsResponse, error)
	mustEmbedUnimplementedPublicAuthenticationServiceServer()
}

// UnimplementedPublicAuthenticationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPublicAuthenticationServiceServer struct {
}

func (UnimplementedPublicAuthenticationServiceServer) ListAuthenticationMethods(context.Context, *emptypb.Empty) (*ListAuthenticationMethodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthenticationMethods not implemented")
}
func (UnimplementedPublicAuthenticationServiceServer) mustEmbedUnimplementedPublicAuthenticationServiceServer() {
}

// UnsafePublicAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicAuthenticationServiceServer will
// result in compilation errors.
type UnsafePublicAuthenticationServiceServer interface {
	mustEmbedUnimplementedPublicAuthenticationServiceServer()
}

func RegisterPublicAuthenticationServiceServer(s grpc.ServiceRegistrar, srv PublicAuthenticationServiceServer) {
	s.RegisterService(&PublicAuthenticationService_ServiceDesc, srv)
}

func _PublicAuthenticationService_ListAuthenticationMethods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicAuthenticationServiceServer).ListAuthenticationMethods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicAuthenticationService_ListAuthenticationMethods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicAuthenticationServiceServer).ListAuthenticationMethods(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicAuthenticationService_ServiceDesc is the grpc.ServiceDesc for PublicAuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicAuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipt.auth.PublicAuthenticationService",
	HandlerType: (*PublicAuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAuthenticationMethods",
			Handler:    _PublicAuthenticationService_ListAuthenticationMethods_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

const (
	AuthenticationService_GetAuthenticationSelf_FullMethodName    = "/flipt.auth.AuthenticationService/GetAuthenticationSelf"
	AuthenticationService_GetAuthentication_FullMethodName        = "/flipt.auth.AuthenticationService/GetAuthentication"
	AuthenticationService_ListAuthentications_FullMethodName      = "/flipt.auth.AuthenticationService/ListAuthentications"
	AuthenticationService_DeleteAuthentication_FullMethodName     = "/flipt.auth.AuthenticationService/DeleteAuthentication"
	AuthenticationService_ExpireAuthenticationSelf_FullMethodName = "/flipt.auth.AuthenticationService/ExpireAuthenticationSelf"
)

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	GetAuthenticationSelf(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Authentication, error)
	GetAuthentication(ctx context.Context, in *GetAuthenticationRequest, opts ...grpc.CallOption) (*Authentication, error)
	ListAuthentications(ctx context.Context, in *ListAuthenticationsRequest, opts ...grpc.CallOption) (*ListAuthenticationsResponse, error)
	DeleteAuthentication(ctx context.Context, in *DeleteAuthenticationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ExpireAuthenticationSelf(ctx context.Context, in *ExpireAuthenticationSelfRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) GetAuthenticationSelf(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Authentication, error) {
	out := new(Authentication)
	err := c.cc.Invoke(ctx, AuthenticationService_GetAuthenticationSelf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) GetAuthentication(ctx context.Context, in *GetAuthenticationRequest, opts ...grpc.CallOption) (*Authentication, error) {
	out := new(Authentication)
	err := c.cc.Invoke(ctx, AuthenticationService_GetAuthentication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ListAuthentications(ctx context.Context, in *ListAuthenticationsRequest, opts ...grpc.CallOption) (*ListAuthenticationsResponse, error) {
	out := new(ListAuthenticationsResponse)
	err := c.cc.Invoke(ctx, AuthenticationService_ListAuthentications_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) DeleteAuthentication(ctx context.Context, in *DeleteAuthenticationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthenticationService_DeleteAuthentication_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ExpireAuthenticationSelf(ctx context.Context, in *ExpireAuthenticationSelfRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AuthenticationService_ExpireAuthenticationSelf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
// All implementations must embed UnimplementedAuthenticationServiceServer
// for forward compatibility
type AuthenticationServiceServer interface {
	GetAuthenticationSelf(context.Context, *emptypb.Empty) (*Authentication, error)
	GetAuthentication(context.Context, *GetAuthenticationRequest) (*Authentication, error)
	ListAuthentications(context.Context, *ListAuthenticationsRequest) (*ListAuthenticationsResponse, error)
	DeleteAuthentication(context.Context, *DeleteAuthenticationRequest) (*emptypb.Empty, error)
	ExpireAuthenticationSelf(context.Context, *ExpireAuthenticationSelfRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedAuthenticationServiceServer()
}

// UnimplementedAuthenticationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (UnimplementedAuthenticationServiceServer) GetAuthenticationSelf(context.Context, *emptypb.Empty) (*Authentication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthenticationSelf not implemented")
}
func (UnimplementedAuthenticationServiceServer) GetAuthentication(context.Context, *GetAuthenticationRequest) (*Authentication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthentication not implemented")
}
func (UnimplementedAuthenticationServiceServer) ListAuthentications(context.Context, *ListAuthenticationsRequest) (*ListAuthenticationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAuthentications not implemented")
}
func (UnimplementedAuthenticationServiceServer) DeleteAuthentication(context.Context, *DeleteAuthenticationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAuthentication not implemented")
}
func (UnimplementedAuthenticationServiceServer) ExpireAuthenticationSelf(context.Context, *ExpireAuthenticationSelfRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExpireAuthenticationSelf not implemented")
}
func (UnimplementedAuthenticationServiceServer) mustEmbedUnimplementedAuthenticationServiceServer() {}

// UnsafeAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServiceServer will
// result in compilation errors.
type UnsafeAuthenticationServiceServer interface {
	mustEmbedUnimplementedAuthenticationServiceServer()
}

func RegisterAuthenticationServiceServer(s grpc.ServiceRegistrar, srv AuthenticationServiceServer) {
	s.RegisterService(&AuthenticationService_ServiceDesc, srv)
}

func _AuthenticationService_GetAuthenticationSelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).GetAuthenticationSelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_GetAuthenticationSelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).GetAuthenticationSelf(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_GetAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).GetAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_GetAuthentication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).GetAuthentication(ctx, req.(*GetAuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ListAuthentications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAuthenticationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ListAuthentications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_ListAuthentications_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ListAuthentications(ctx, req.(*ListAuthenticationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_DeleteAuthentication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAuthenticationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).DeleteAuthentication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_DeleteAuthentication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).DeleteAuthentication(ctx, req.(*DeleteAuthenticationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ExpireAuthenticationSelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpireAuthenticationSelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ExpireAuthenticationSelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationService_ExpireAuthenticationSelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ExpireAuthenticationSelf(ctx, req.(*ExpireAuthenticationSelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationService_ServiceDesc is the grpc.ServiceDesc for AuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipt.auth.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAuthenticationSelf",
			Handler:    _AuthenticationService_GetAuthenticationSelf_Handler,
		},
		{
			MethodName: "GetAuthentication",
			Handler:    _AuthenticationService_GetAuthentication_Handler,
		},
		{
			MethodName: "ListAuthentications",
			Handler:    _AuthenticationService_ListAuthentications_Handler,
		},
		{
			MethodName: "DeleteAuthentication",
			Handler:    _AuthenticationService_DeleteAuthentication_Handler,
		},
		{
			MethodName: "ExpireAuthenticationSelf",
			Handler:    _AuthenticationService_ExpireAuthenticationSelf_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

const (
	AuthenticationMethodTokenService_CreateToken_FullMethodName = "/flipt.auth.AuthenticationMethodTokenService/CreateToken"
)

// AuthenticationMethodTokenServiceClient is the client API for AuthenticationMethodTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationMethodTokenServiceClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
}

type authenticationMethodTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationMethodTokenServiceClient(cc grpc.ClientConnInterface) AuthenticationMethodTokenServiceClient {
	return &authenticationMethodTokenServiceClient{cc}
}

func (c *authenticationMethodTokenServiceClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := c.cc.Invoke(ctx, AuthenticationMethodTokenService_CreateToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationMethodTokenServiceServer is the server API for AuthenticationMethodTokenService service.
// All implementations must embed UnimplementedAuthenticationMethodTokenServiceServer
// for forward compatibility
type AuthenticationMethodTokenServiceServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	mustEmbedUnimplementedAuthenticationMethodTokenServiceServer()
}

// UnimplementedAuthenticationMethodTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationMethodTokenServiceServer struct {
}

func (UnimplementedAuthenticationMethodTokenServiceServer) CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateToken not implemented")
}
func (UnimplementedAuthenticationMethodTokenServiceServer) mustEmbedUnimplementedAuthenticationMethodTokenServiceServer() {
}

// UnsafeAuthenticationMethodTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationMethodTokenServiceServer will
// result in compilation errors.
type UnsafeAuthenticationMethodTokenServiceServer interface {
	mustEmbedUnimplementedAuthenticationMethodTokenServiceServer()
}

func RegisterAuthenticationMethodTokenServiceServer(s grpc.ServiceRegistrar, srv AuthenticationMethodTokenServiceServer) {
	s.RegisterService(&AuthenticationMethodTokenService_ServiceDesc, srv)
}

func _AuthenticationMethodTokenService_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationMethodTokenServiceServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationMethodTokenService_CreateToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationMethodTokenServiceServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationMethodTokenService_ServiceDesc is the grpc.ServiceDesc for AuthenticationMethodTokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationMethodTokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipt.auth.AuthenticationMethodTokenService",
	HandlerType: (*AuthenticationMethodTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _AuthenticationMethodTokenService_CreateToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

const (
	AuthenticationMethodOIDCService_AuthorizeURL_FullMethodName = "/flipt.auth.AuthenticationMethodOIDCService/AuthorizeURL"
	AuthenticationMethodOIDCService_Callback_FullMethodName     = "/flipt.auth.AuthenticationMethodOIDCService/Callback"
)

// AuthenticationMethodOIDCServiceClient is the client API for AuthenticationMethodOIDCService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationMethodOIDCServiceClient interface {
	AuthorizeURL(ctx context.Context, in *AuthorizeURLRequest, opts ...grpc.CallOption) (*AuthorizeURLResponse, error)
	Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error)
}

type authenticationMethodOIDCServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationMethodOIDCServiceClient(cc grpc.ClientConnInterface) AuthenticationMethodOIDCServiceClient {
	return &authenticationMethodOIDCServiceClient{cc}
}

func (c *authenticationMethodOIDCServiceClient) AuthorizeURL(ctx context.Context, in *AuthorizeURLRequest, opts ...grpc.CallOption) (*AuthorizeURLResponse, error) {
	out := new(AuthorizeURLResponse)
	err := c.cc.Invoke(ctx, AuthenticationMethodOIDCService_AuthorizeURL_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationMethodOIDCServiceClient) Callback(ctx context.Context, in *CallbackRequest, opts ...grpc.CallOption) (*CallbackResponse, error) {
	out := new(CallbackResponse)
	err := c.cc.Invoke(ctx, AuthenticationMethodOIDCService_Callback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationMethodOIDCServiceServer is the server API for AuthenticationMethodOIDCService service.
// All implementations must embed UnimplementedAuthenticationMethodOIDCServiceServer
// for forward compatibility
type AuthenticationMethodOIDCServiceServer interface {
	AuthorizeURL(context.Context, *AuthorizeURLRequest) (*AuthorizeURLResponse, error)
	Callback(context.Context, *CallbackRequest) (*CallbackResponse, error)
	mustEmbedUnimplementedAuthenticationMethodOIDCServiceServer()
}

// UnimplementedAuthenticationMethodOIDCServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationMethodOIDCServiceServer struct {
}

func (UnimplementedAuthenticationMethodOIDCServiceServer) AuthorizeURL(context.Context, *AuthorizeURLRequest) (*AuthorizeURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeURL not implemented")
}
func (UnimplementedAuthenticationMethodOIDCServiceServer) Callback(context.Context, *CallbackRequest) (*CallbackResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Callback not implemented")
}
func (UnimplementedAuthenticationMethodOIDCServiceServer) mustEmbedUnimplementedAuthenticationMethodOIDCServiceServer() {
}

// UnsafeAuthenticationMethodOIDCServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationMethodOIDCServiceServer will
// result in compilation errors.
type UnsafeAuthenticationMethodOIDCServiceServer interface {
	mustEmbedUnimplementedAuthenticationMethodOIDCServiceServer()
}

func RegisterAuthenticationMethodOIDCServiceServer(s grpc.ServiceRegistrar, srv AuthenticationMethodOIDCServiceServer) {
	s.RegisterService(&AuthenticationMethodOIDCService_ServiceDesc, srv)
}

func _AuthenticationMethodOIDCService_AuthorizeURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationMethodOIDCServiceServer).AuthorizeURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationMethodOIDCService_AuthorizeURL_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationMethodOIDCServiceServer).AuthorizeURL(ctx, req.(*AuthorizeURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationMethodOIDCService_Callback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationMethodOIDCServiceServer).Callback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationMethodOIDCService_Callback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationMethodOIDCServiceServer).Callback(ctx, req.(*CallbackRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationMethodOIDCService_ServiceDesc is the grpc.ServiceDesc for AuthenticationMethodOIDCService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationMethodOIDCService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipt.auth.AuthenticationMethodOIDCService",
	HandlerType: (*AuthenticationMethodOIDCServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AuthorizeURL",
			Handler:    _AuthenticationMethodOIDCService_AuthorizeURL_Handler,
		},
		{
			MethodName: "Callback",
			Handler:    _AuthenticationMethodOIDCService_Callback_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

const (
	AuthenticationMethodKubernetesService_VerifyServiceAccount_FullMethodName = "/flipt.auth.AuthenticationMethodKubernetesService/VerifyServiceAccount"
)

// AuthenticationMethodKubernetesServiceClient is the client API for AuthenticationMethodKubernetesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationMethodKubernetesServiceClient interface {
	VerifyServiceAccount(ctx context.Context, in *VerifyServiceAccountRequest, opts ...grpc.CallOption) (*VerifyServiceAccountResponse, error)
}

type authenticationMethodKubernetesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationMethodKubernetesServiceClient(cc grpc.ClientConnInterface) AuthenticationMethodKubernetesServiceClient {
	return &authenticationMethodKubernetesServiceClient{cc}
}

func (c *authenticationMethodKubernetesServiceClient) VerifyServiceAccount(ctx context.Context, in *VerifyServiceAccountRequest, opts ...grpc.CallOption) (*VerifyServiceAccountResponse, error) {
	out := new(VerifyServiceAccountResponse)
	err := c.cc.Invoke(ctx, AuthenticationMethodKubernetesService_VerifyServiceAccount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationMethodKubernetesServiceServer is the server API for AuthenticationMethodKubernetesService service.
// All implementations must embed UnimplementedAuthenticationMethodKubernetesServiceServer
// for forward compatibility
type AuthenticationMethodKubernetesServiceServer interface {
	VerifyServiceAccount(context.Context, *VerifyServiceAccountRequest) (*VerifyServiceAccountResponse, error)
	mustEmbedUnimplementedAuthenticationMethodKubernetesServiceServer()
}

// UnimplementedAuthenticationMethodKubernetesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationMethodKubernetesServiceServer struct {
}

func (UnimplementedAuthenticationMethodKubernetesServiceServer) VerifyServiceAccount(context.Context, *VerifyServiceAccountRequest) (*VerifyServiceAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyServiceAccount not implemented")
}
func (UnimplementedAuthenticationMethodKubernetesServiceServer) mustEmbedUnimplementedAuthenticationMethodKubernetesServiceServer() {
}

// UnsafeAuthenticationMethodKubernetesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationMethodKubernetesServiceServer will
// result in compilation errors.
type UnsafeAuthenticationMethodKubernetesServiceServer interface {
	mustEmbedUnimplementedAuthenticationMethodKubernetesServiceServer()
}

func RegisterAuthenticationMethodKubernetesServiceServer(s grpc.ServiceRegistrar, srv AuthenticationMethodKubernetesServiceServer) {
	s.RegisterService(&AuthenticationMethodKubernetesService_ServiceDesc, srv)
}

func _AuthenticationMethodKubernetesService_VerifyServiceAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyServiceAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationMethodKubernetesServiceServer).VerifyServiceAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthenticationMethodKubernetesService_VerifyServiceAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationMethodKubernetesServiceServer).VerifyServiceAccount(ctx, req.(*VerifyServiceAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationMethodKubernetesService_ServiceDesc is the grpc.ServiceDesc for AuthenticationMethodKubernetesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationMethodKubernetesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "flipt.auth.AuthenticationMethodKubernetesService",
	HandlerType: (*AuthenticationMethodKubernetesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyServiceAccount",
			Handler:    _AuthenticationMethodKubernetesService_VerifyServiceAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

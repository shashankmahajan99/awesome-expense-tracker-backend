// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package AwesomeExpenseTrackerApi

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

// HealthCheckClient is the client API for HealthCheck service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HealthCheckClient interface {
	Healthy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type healthCheckClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthCheckClient(cc grpc.ClientConnInterface) HealthCheckClient {
	return &healthCheckClient{cc}
}

func (c *healthCheckClient) Healthy(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/apidefinitions.HealthCheck/Healthy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthCheckServer is the server API for HealthCheck service.
// All implementations must embed UnimplementedHealthCheckServer
// for forward compatibility
type HealthCheckServer interface {
	Healthy(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedHealthCheckServer()
}

// UnimplementedHealthCheckServer must be embedded to have forward compatible implementations.
type UnimplementedHealthCheckServer struct {
}

func (UnimplementedHealthCheckServer) Healthy(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthy not implemented")
}
func (UnimplementedHealthCheckServer) mustEmbedUnimplementedHealthCheckServer() {}

// UnsafeHealthCheckServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HealthCheckServer will
// result in compilation errors.
type UnsafeHealthCheckServer interface {
	mustEmbedUnimplementedHealthCheckServer()
}

func RegisterHealthCheckServer(s grpc.ServiceRegistrar, srv HealthCheckServer) {
	s.RegisterService(&HealthCheck_ServiceDesc, srv)
}

func _HealthCheck_Healthy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthCheckServer).Healthy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.HealthCheck/Healthy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthCheckServer).Healthy(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// HealthCheck_ServiceDesc is the grpc.ServiceDesc for HealthCheck service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HealthCheck_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apidefinitions.HealthCheck",
	HandlerType: (*HealthCheckServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthy",
			Handler:    _HealthCheck_Healthy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// UserAuthenticationClient is the client API for UserAuthentication service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserAuthenticationClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type userAuthenticationClient struct {
	cc grpc.ClientConnInterface
}

func NewUserAuthenticationClient(cc grpc.ClientConnInterface) UserAuthenticationClient {
	return &userAuthenticationClient{cc}
}

func (c *userAuthenticationClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/apidefinitions.UserAuthentication/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userAuthenticationClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/apidefinitions.UserAuthentication/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserAuthenticationServer is the server API for UserAuthentication service.
// All implementations must embed UnimplementedUserAuthenticationServer
// for forward compatibility
type UserAuthenticationServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	mustEmbedUnimplementedUserAuthenticationServer()
}

// UnimplementedUserAuthenticationServer must be embedded to have forward compatible implementations.
type UnimplementedUserAuthenticationServer struct {
}

func (UnimplementedUserAuthenticationServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserAuthenticationServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserAuthenticationServer) mustEmbedUnimplementedUserAuthenticationServer() {}

// UnsafeUserAuthenticationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserAuthenticationServer will
// result in compilation errors.
type UnsafeUserAuthenticationServer interface {
	mustEmbedUnimplementedUserAuthenticationServer()
}

func RegisterUserAuthenticationServer(s grpc.ServiceRegistrar, srv UserAuthenticationServer) {
	s.RegisterService(&UserAuthentication_ServiceDesc, srv)
}

func _UserAuthentication_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthenticationServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.UserAuthentication/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthenticationServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserAuthentication_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserAuthenticationServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.UserAuthentication/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserAuthenticationServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserAuthentication_ServiceDesc is the grpc.ServiceDesc for UserAuthentication service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserAuthentication_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apidefinitions.UserAuthentication",
	HandlerType: (*UserAuthenticationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _UserAuthentication_Login_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _UserAuthentication_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// ExpenseManagementClient is the client API for ExpenseManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExpenseManagementClient interface {
	CreateExpense(ctx context.Context, in *CreateExpenseRequest, opts ...grpc.CallOption) (*CreateExpenseResponse, error)
	UpdateExpense(ctx context.Context, in *UpdateExpenseRequest, opts ...grpc.CallOption) (*UpdateExpenseResponse, error)
	DeleteExpense(ctx context.Context, in *DeleteExpenseRequest, opts ...grpc.CallOption) (*DeleteExpenseResponse, error)
	GetExpenses(ctx context.Context, in *GetExpensesRequest, opts ...grpc.CallOption) (*GetExpensesResponse, error)
}

type expenseManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewExpenseManagementClient(cc grpc.ClientConnInterface) ExpenseManagementClient {
	return &expenseManagementClient{cc}
}

func (c *expenseManagementClient) CreateExpense(ctx context.Context, in *CreateExpenseRequest, opts ...grpc.CallOption) (*CreateExpenseResponse, error) {
	out := new(CreateExpenseResponse)
	err := c.cc.Invoke(ctx, "/apidefinitions.ExpenseManagement/CreateExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseManagementClient) UpdateExpense(ctx context.Context, in *UpdateExpenseRequest, opts ...grpc.CallOption) (*UpdateExpenseResponse, error) {
	out := new(UpdateExpenseResponse)
	err := c.cc.Invoke(ctx, "/apidefinitions.ExpenseManagement/UpdateExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseManagementClient) DeleteExpense(ctx context.Context, in *DeleteExpenseRequest, opts ...grpc.CallOption) (*DeleteExpenseResponse, error) {
	out := new(DeleteExpenseResponse)
	err := c.cc.Invoke(ctx, "/apidefinitions.ExpenseManagement/DeleteExpense", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expenseManagementClient) GetExpenses(ctx context.Context, in *GetExpensesRequest, opts ...grpc.CallOption) (*GetExpensesResponse, error) {
	out := new(GetExpensesResponse)
	err := c.cc.Invoke(ctx, "/apidefinitions.ExpenseManagement/GetExpenses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpenseManagementServer is the server API for ExpenseManagement service.
// All implementations must embed UnimplementedExpenseManagementServer
// for forward compatibility
type ExpenseManagementServer interface {
	CreateExpense(context.Context, *CreateExpenseRequest) (*CreateExpenseResponse, error)
	UpdateExpense(context.Context, *UpdateExpenseRequest) (*UpdateExpenseResponse, error)
	DeleteExpense(context.Context, *DeleteExpenseRequest) (*DeleteExpenseResponse, error)
	GetExpenses(context.Context, *GetExpensesRequest) (*GetExpensesResponse, error)
	mustEmbedUnimplementedExpenseManagementServer()
}

// UnimplementedExpenseManagementServer must be embedded to have forward compatible implementations.
type UnimplementedExpenseManagementServer struct {
}

func (UnimplementedExpenseManagementServer) CreateExpense(context.Context, *CreateExpenseRequest) (*CreateExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateExpense not implemented")
}
func (UnimplementedExpenseManagementServer) UpdateExpense(context.Context, *UpdateExpenseRequest) (*UpdateExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateExpense not implemented")
}
func (UnimplementedExpenseManagementServer) DeleteExpense(context.Context, *DeleteExpenseRequest) (*DeleteExpenseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExpense not implemented")
}
func (UnimplementedExpenseManagementServer) GetExpenses(context.Context, *GetExpensesRequest) (*GetExpensesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExpenses not implemented")
}
func (UnimplementedExpenseManagementServer) mustEmbedUnimplementedExpenseManagementServer() {}

// UnsafeExpenseManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpenseManagementServer will
// result in compilation errors.
type UnsafeExpenseManagementServer interface {
	mustEmbedUnimplementedExpenseManagementServer()
}

func RegisterExpenseManagementServer(s grpc.ServiceRegistrar, srv ExpenseManagementServer) {
	s.RegisterService(&ExpenseManagement_ServiceDesc, srv)
}

func _ExpenseManagement_CreateExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseManagementServer).CreateExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.ExpenseManagement/CreateExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseManagementServer).CreateExpense(ctx, req.(*CreateExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseManagement_UpdateExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseManagementServer).UpdateExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.ExpenseManagement/UpdateExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseManagementServer).UpdateExpense(ctx, req.(*UpdateExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseManagement_DeleteExpense_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExpenseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseManagementServer).DeleteExpense(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.ExpenseManagement/DeleteExpense",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseManagementServer).DeleteExpense(ctx, req.(*DeleteExpenseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpenseManagement_GetExpenses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExpensesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpenseManagementServer).GetExpenses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.ExpenseManagement/GetExpenses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpenseManagementServer).GetExpenses(ctx, req.(*GetExpensesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExpenseManagement_ServiceDesc is the grpc.ServiceDesc for ExpenseManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExpenseManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apidefinitions.ExpenseManagement",
	HandlerType: (*ExpenseManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateExpense",
			Handler:    _ExpenseManagement_CreateExpense_Handler,
		},
		{
			MethodName: "UpdateExpense",
			Handler:    _ExpenseManagement_UpdateExpense_Handler,
		},
		{
			MethodName: "DeleteExpense",
			Handler:    _ExpenseManagement_DeleteExpense_Handler,
		},
		{
			MethodName: "GetExpenses",
			Handler:    _ExpenseManagement_GetExpenses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// ReportsClient is the client API for Reports service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReportsClient interface {
	GenerateReport(ctx context.Context, in *GenerateReportRequest, opts ...grpc.CallOption) (*GenerateReportResponse, error)
}

type reportsClient struct {
	cc grpc.ClientConnInterface
}

func NewReportsClient(cc grpc.ClientConnInterface) ReportsClient {
	return &reportsClient{cc}
}

func (c *reportsClient) GenerateReport(ctx context.Context, in *GenerateReportRequest, opts ...grpc.CallOption) (*GenerateReportResponse, error) {
	out := new(GenerateReportResponse)
	err := c.cc.Invoke(ctx, "/apidefinitions.Reports/GenerateReport", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReportsServer is the server API for Reports service.
// All implementations must embed UnimplementedReportsServer
// for forward compatibility
type ReportsServer interface {
	GenerateReport(context.Context, *GenerateReportRequest) (*GenerateReportResponse, error)
	mustEmbedUnimplementedReportsServer()
}

// UnimplementedReportsServer must be embedded to have forward compatible implementations.
type UnimplementedReportsServer struct {
}

func (UnimplementedReportsServer) GenerateReport(context.Context, *GenerateReportRequest) (*GenerateReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateReport not implemented")
}
func (UnimplementedReportsServer) mustEmbedUnimplementedReportsServer() {}

// UnsafeReportsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReportsServer will
// result in compilation errors.
type UnsafeReportsServer interface {
	mustEmbedUnimplementedReportsServer()
}

func RegisterReportsServer(s grpc.ServiceRegistrar, srv ReportsServer) {
	s.RegisterService(&Reports_ServiceDesc, srv)
}

func _Reports_GenerateReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReportsServer).GenerateReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.Reports/GenerateReport",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReportsServer).GenerateReport(ctx, req.(*GenerateReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Reports_ServiceDesc is the grpc.ServiceDesc for Reports service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Reports_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apidefinitions.Reports",
	HandlerType: (*ReportsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateReport",
			Handler:    _Reports_GenerateReport_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// UserProfileClient is the client API for UserProfile service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserProfileClient interface {
	GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*GetUserProfileResponse, error)
	UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UpdateUserProfileResponse, error)
}

type userProfileClient struct {
	cc grpc.ClientConnInterface
}

func NewUserProfileClient(cc grpc.ClientConnInterface) UserProfileClient {
	return &userProfileClient{cc}
}

func (c *userProfileClient) GetUserProfile(ctx context.Context, in *GetUserProfileRequest, opts ...grpc.CallOption) (*GetUserProfileResponse, error) {
	out := new(GetUserProfileResponse)
	err := c.cc.Invoke(ctx, "/apidefinitions.UserProfile/GetUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userProfileClient) UpdateUserProfile(ctx context.Context, in *UpdateUserProfileRequest, opts ...grpc.CallOption) (*UpdateUserProfileResponse, error) {
	out := new(UpdateUserProfileResponse)
	err := c.cc.Invoke(ctx, "/apidefinitions.UserProfile/UpdateUserProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserProfileServer is the server API for UserProfile service.
// All implementations must embed UnimplementedUserProfileServer
// for forward compatibility
type UserProfileServer interface {
	GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileResponse, error)
	UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UpdateUserProfileResponse, error)
	mustEmbedUnimplementedUserProfileServer()
}

// UnimplementedUserProfileServer must be embedded to have forward compatible implementations.
type UnimplementedUserProfileServer struct {
}

func (UnimplementedUserProfileServer) GetUserProfile(context.Context, *GetUserProfileRequest) (*GetUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserProfile not implemented")
}
func (UnimplementedUserProfileServer) UpdateUserProfile(context.Context, *UpdateUserProfileRequest) (*UpdateUserProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserProfile not implemented")
}
func (UnimplementedUserProfileServer) mustEmbedUnimplementedUserProfileServer() {}

// UnsafeUserProfileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserProfileServer will
// result in compilation errors.
type UnsafeUserProfileServer interface {
	mustEmbedUnimplementedUserProfileServer()
}

func RegisterUserProfileServer(s grpc.ServiceRegistrar, srv UserProfileServer) {
	s.RegisterService(&UserProfile_ServiceDesc, srv)
}

func _UserProfile_GetUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).GetUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.UserProfile/GetUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).GetUserProfile(ctx, req.(*GetUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserProfile_UpdateUserProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserProfileServer).UpdateUserProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.UserProfile/UpdateUserProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserProfileServer).UpdateUserProfile(ctx, req.(*UpdateUserProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserProfile_ServiceDesc is the grpc.ServiceDesc for UserProfile service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserProfile_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apidefinitions.UserProfile",
	HandlerType: (*UserProfileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserProfile",
			Handler:    _UserProfile_GetUserProfile_Handler,
		},
		{
			MethodName: "UpdateUserProfile",
			Handler:    _UserProfile_UpdateUserProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// SettingsClient is the client API for Settings service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SettingsClient interface {
	GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error)
	UpdateSettings(ctx context.Context, in *UpdateSettingsRequest, opts ...grpc.CallOption) (*UpdateSettingsResponse, error)
}

type settingsClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsClient(cc grpc.ClientConnInterface) SettingsClient {
	return &settingsClient{cc}
}

func (c *settingsClient) GetSettings(ctx context.Context, in *GetSettingsRequest, opts ...grpc.CallOption) (*GetSettingsResponse, error) {
	out := new(GetSettingsResponse)
	err := c.cc.Invoke(ctx, "/apidefinitions.Settings/GetSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsClient) UpdateSettings(ctx context.Context, in *UpdateSettingsRequest, opts ...grpc.CallOption) (*UpdateSettingsResponse, error) {
	out := new(UpdateSettingsResponse)
	err := c.cc.Invoke(ctx, "/apidefinitions.Settings/UpdateSettings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServer is the server API for Settings service.
// All implementations must embed UnimplementedSettingsServer
// for forward compatibility
type SettingsServer interface {
	GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error)
	UpdateSettings(context.Context, *UpdateSettingsRequest) (*UpdateSettingsResponse, error)
	mustEmbedUnimplementedSettingsServer()
}

// UnimplementedSettingsServer must be embedded to have forward compatible implementations.
type UnimplementedSettingsServer struct {
}

func (UnimplementedSettingsServer) GetSettings(context.Context, *GetSettingsRequest) (*GetSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSettings not implemented")
}
func (UnimplementedSettingsServer) UpdateSettings(context.Context, *UpdateSettingsRequest) (*UpdateSettingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSettings not implemented")
}
func (UnimplementedSettingsServer) mustEmbedUnimplementedSettingsServer() {}

// UnsafeSettingsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingsServer will
// result in compilation errors.
type UnsafeSettingsServer interface {
	mustEmbedUnimplementedSettingsServer()
}

func RegisterSettingsServer(s grpc.ServiceRegistrar, srv SettingsServer) {
	s.RegisterService(&Settings_ServiceDesc, srv)
}

func _Settings_GetSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).GetSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.Settings/GetSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).GetSettings(ctx, req.(*GetSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Settings_UpdateSettings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSettingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServer).UpdateSettings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apidefinitions.Settings/UpdateSettings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServer).UpdateSettings(ctx, req.(*UpdateSettingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Settings_ServiceDesc is the grpc.ServiceDesc for Settings service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Settings_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apidefinitions.Settings",
	HandlerType: (*SettingsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSettings",
			Handler:    _Settings_GetSettings_Handler,
		},
		{
			MethodName: "UpdateSettings",
			Handler:    _Settings_UpdateSettings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

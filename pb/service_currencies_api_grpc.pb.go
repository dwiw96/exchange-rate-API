// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service_currencies_api.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CurrenciesAPIClient is the client API for CurrenciesAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CurrenciesAPIClient interface {
	ListAllCurrencies(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListAllCurrenciesResponse, error)
	GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error)
}

type currenciesAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewCurrenciesAPIClient(cc grpc.ClientConnInterface) CurrenciesAPIClient {
	return &currenciesAPIClient{cc}
}

func (c *currenciesAPIClient) ListAllCurrencies(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ListAllCurrenciesResponse, error) {
	out := new(ListAllCurrenciesResponse)
	err := c.cc.Invoke(ctx, "/pb.CurrenciesAPI/ListAllCurrencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currenciesAPIClient) GetCurrency(ctx context.Context, in *GetCurrencyRequest, opts ...grpc.CallOption) (*GetCurrencyResponse, error) {
	out := new(GetCurrencyResponse)
	err := c.cc.Invoke(ctx, "/pb.CurrenciesAPI/GetCurrency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CurrenciesAPIServer is the server API for CurrenciesAPI service.
// All implementations must embed UnimplementedCurrenciesAPIServer
// for forward compatibility
type CurrenciesAPIServer interface {
	ListAllCurrencies(context.Context, *empty.Empty) (*ListAllCurrenciesResponse, error)
	GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error)
	mustEmbedUnimplementedCurrenciesAPIServer()
}

// UnimplementedCurrenciesAPIServer must be embedded to have forward compatible implementations.
type UnimplementedCurrenciesAPIServer struct {
}

func (UnimplementedCurrenciesAPIServer) ListAllCurrencies(context.Context, *empty.Empty) (*ListAllCurrenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAllCurrencies not implemented")
}
func (UnimplementedCurrenciesAPIServer) GetCurrency(context.Context, *GetCurrencyRequest) (*GetCurrencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrency not implemented")
}
func (UnimplementedCurrenciesAPIServer) mustEmbedUnimplementedCurrenciesAPIServer() {}

// UnsafeCurrenciesAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CurrenciesAPIServer will
// result in compilation errors.
type UnsafeCurrenciesAPIServer interface {
	mustEmbedUnimplementedCurrenciesAPIServer()
}

func RegisterCurrenciesAPIServer(s grpc.ServiceRegistrar, srv CurrenciesAPIServer) {
	s.RegisterService(&CurrenciesAPI_ServiceDesc, srv)
}

func _CurrenciesAPI_ListAllCurrencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrenciesAPIServer).ListAllCurrencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CurrenciesAPI/ListAllCurrencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrenciesAPIServer).ListAllCurrencies(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _CurrenciesAPI_GetCurrency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CurrenciesAPIServer).GetCurrency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CurrenciesAPI/GetCurrency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CurrenciesAPIServer).GetCurrency(ctx, req.(*GetCurrencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CurrenciesAPI_ServiceDesc is the grpc.ServiceDesc for CurrenciesAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CurrenciesAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CurrenciesAPI",
	HandlerType: (*CurrenciesAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListAllCurrencies",
			Handler:    _CurrenciesAPI_ListAllCurrencies_Handler,
		},
		{
			MethodName: "GetCurrency",
			Handler:    _CurrenciesAPI_GetCurrency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_currencies_api.proto",
}

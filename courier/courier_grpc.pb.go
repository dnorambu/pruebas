// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package courier

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CourierServiceClient is the client API for CourierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourierServiceClient interface {
	EnviarOrdenPyme(ctx context.Context, in *OrdenPyme, opts ...grpc.CallOption) (*Codigo, error)
	EnviarOrdenRetail(ctx context.Context, in *OrdenRetail, opts ...grpc.CallOption) (*Empty, error)
	Seguimiento(ctx context.Context, in *Codigo, opts ...grpc.CallOption) (*Mensaje, error)
	PedirRetail(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Paquete, error)
	PedirPrioritario(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Paquete, error)
	PedirNormal(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Paquete, error)
	ResultadoEntrega(ctx context.Context, in *Entrega, opts ...grpc.CallOption) (*Empty, error)
}

type courierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCourierServiceClient(cc grpc.ClientConnInterface) CourierServiceClient {
	return &courierServiceClient{cc}
}

func (c *courierServiceClient) EnviarOrdenPyme(ctx context.Context, in *OrdenPyme, opts ...grpc.CallOption) (*Codigo, error) {
	out := new(Codigo)
	err := c.cc.Invoke(ctx, "/CourierService/EnviarOrdenPyme", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) EnviarOrdenRetail(ctx context.Context, in *OrdenRetail, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/CourierService/EnviarOrdenRetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) Seguimiento(ctx context.Context, in *Codigo, opts ...grpc.CallOption) (*Mensaje, error) {
	out := new(Mensaje)
	err := c.cc.Invoke(ctx, "/CourierService/Seguimiento", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) PedirRetail(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Paquete, error) {
	out := new(Paquete)
	err := c.cc.Invoke(ctx, "/CourierService/PedirRetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) PedirPrioritario(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Paquete, error) {
	out := new(Paquete)
	err := c.cc.Invoke(ctx, "/CourierService/PedirPrioritario", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) PedirNormal(ctx context.Context, in *Mensaje, opts ...grpc.CallOption) (*Paquete, error) {
	out := new(Paquete)
	err := c.cc.Invoke(ctx, "/CourierService/PedirNormal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courierServiceClient) ResultadoEntrega(ctx context.Context, in *Entrega, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/CourierService/ResultadoEntrega", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourierServiceServer is the server API for CourierService service.
// All implementations must embed UnimplementedCourierServiceServer
// for forward compatibility
type CourierServiceServer interface {
	EnviarOrdenPyme(context.Context, *OrdenPyme) (*Codigo, error)
	EnviarOrdenRetail(context.Context, *OrdenRetail) (*Empty, error)
	Seguimiento(context.Context, *Codigo) (*Mensaje, error)
	PedirRetail(context.Context, *Mensaje) (*Paquete, error)
	PedirPrioritario(context.Context, *Mensaje) (*Paquete, error)
	PedirNormal(context.Context, *Mensaje) (*Paquete, error)
	ResultadoEntrega(context.Context, *Entrega) (*Empty, error)
	mustEmbedUnimplementedCourierServiceServer()
}

// UnimplementedCourierServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCourierServiceServer struct {
}

func (UnimplementedCourierServiceServer) EnviarOrdenPyme(context.Context, *OrdenPyme) (*Codigo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarOrdenPyme not implemented")
}
func (UnimplementedCourierServiceServer) EnviarOrdenRetail(context.Context, *OrdenRetail) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnviarOrdenRetail not implemented")
}
func (UnimplementedCourierServiceServer) Seguimiento(context.Context, *Codigo) (*Mensaje, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Seguimiento not implemented")
}
func (UnimplementedCourierServiceServer) PedirRetail(context.Context, *Mensaje) (*Paquete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirRetail not implemented")
}
func (UnimplementedCourierServiceServer) PedirPrioritario(context.Context, *Mensaje) (*Paquete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirPrioritario not implemented")
}
func (UnimplementedCourierServiceServer) PedirNormal(context.Context, *Mensaje) (*Paquete, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PedirNormal not implemented")
}
func (UnimplementedCourierServiceServer) ResultadoEntrega(context.Context, *Entrega) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResultadoEntrega not implemented")
}
func (UnimplementedCourierServiceServer) mustEmbedUnimplementedCourierServiceServer() {}

// UnsafeCourierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourierServiceServer will
// result in compilation errors.
type UnsafeCourierServiceServer interface {
	mustEmbedUnimplementedCourierServiceServer()
}

func RegisterCourierServiceServer(s *grpc.Server, srv CourierServiceServer) {
	s.RegisterService(&_CourierService_serviceDesc, srv)
}

func _CourierService_EnviarOrdenPyme_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdenPyme)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).EnviarOrdenPyme(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourierService/EnviarOrdenPyme",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).EnviarOrdenPyme(ctx, req.(*OrdenPyme))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_EnviarOrdenRetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdenRetail)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).EnviarOrdenRetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourierService/EnviarOrdenRetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).EnviarOrdenRetail(ctx, req.(*OrdenRetail))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_Seguimiento_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Codigo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).Seguimiento(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourierService/Seguimiento",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).Seguimiento(ctx, req.(*Codigo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_PedirRetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mensaje)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).PedirRetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourierService/PedirRetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).PedirRetail(ctx, req.(*Mensaje))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_PedirPrioritario_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mensaje)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).PedirPrioritario(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourierService/PedirPrioritario",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).PedirPrioritario(ctx, req.(*Mensaje))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_PedirNormal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mensaje)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).PedirNormal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourierService/PedirNormal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).PedirNormal(ctx, req.(*Mensaje))
	}
	return interceptor(ctx, in, info, handler)
}

func _CourierService_ResultadoEntrega_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entrega)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourierServiceServer).ResultadoEntrega(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CourierService/ResultadoEntrega",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourierServiceServer).ResultadoEntrega(ctx, req.(*Entrega))
	}
	return interceptor(ctx, in, info, handler)
}

var _CourierService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CourierService",
	HandlerType: (*CourierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnviarOrdenPyme",
			Handler:    _CourierService_EnviarOrdenPyme_Handler,
		},
		{
			MethodName: "EnviarOrdenRetail",
			Handler:    _CourierService_EnviarOrdenRetail_Handler,
		},
		{
			MethodName: "Seguimiento",
			Handler:    _CourierService_Seguimiento_Handler,
		},
		{
			MethodName: "PedirRetail",
			Handler:    _CourierService_PedirRetail_Handler,
		},
		{
			MethodName: "PedirPrioritario",
			Handler:    _CourierService_PedirPrioritario_Handler,
		},
		{
			MethodName: "PedirNormal",
			Handler:    _CourierService_PedirNormal_Handler,
		},
		{
			MethodName: "ResultadoEntrega",
			Handler:    _CourierService_ResultadoEntrega_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "courier/courier.proto",
}

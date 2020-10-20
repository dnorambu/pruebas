// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: courier/courier.proto

package courier

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Mensaje struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Mensaje) Reset() {
	*x = Mensaje{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_courier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Mensaje) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Mensaje) ProtoMessage() {}

func (x *Mensaje) ProtoReflect() protoreflect.Message {
	mi := &file_courier_courier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Mensaje.ProtoReflect.Descriptor instead.
func (*Mensaje) Descriptor() ([]byte, []int) {
	return file_courier_courier_proto_rawDescGZIP(), []int{0}
}

func (x *Mensaje) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

type OrdenPyme struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto  string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor     int64  `protobuf:"varint,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda    string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino   string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
	Prioridad int64  `protobuf:"varint,6,opt,name=prioridad,proto3" json:"prioridad,omitempty"`
}

func (x *OrdenPyme) Reset() {
	*x = OrdenPyme{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_courier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdenPyme) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdenPyme) ProtoMessage() {}

func (x *OrdenPyme) ProtoReflect() protoreflect.Message {
	mi := &file_courier_courier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdenPyme.ProtoReflect.Descriptor instead.
func (*OrdenPyme) Descriptor() ([]byte, []int) {
	return file_courier_courier_proto_rawDescGZIP(), []int{1}
}

func (x *OrdenPyme) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrdenPyme) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *OrdenPyme) GetValor() int64 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *OrdenPyme) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *OrdenPyme) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

func (x *OrdenPyme) GetPrioridad() int64 {
	if x != nil {
		return x.Prioridad
	}
	return 0
}

type OrdenRetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Producto string `protobuf:"bytes,2,opt,name=producto,proto3" json:"producto,omitempty"`
	Valor    int64  `protobuf:"varint,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda   string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Destino  string `protobuf:"bytes,5,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *OrdenRetail) Reset() {
	*x = OrdenRetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_courier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdenRetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdenRetail) ProtoMessage() {}

func (x *OrdenRetail) ProtoReflect() protoreflect.Message {
	mi := &file_courier_courier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdenRetail.ProtoReflect.Descriptor instead.
func (*OrdenRetail) Descriptor() ([]byte, []int) {
	return file_courier_courier_proto_rawDescGZIP(), []int{2}
}

func (x *OrdenRetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrdenRetail) GetProducto() string {
	if x != nil {
		return x.Producto
	}
	return ""
}

func (x *OrdenRetail) GetValor() int64 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *OrdenRetail) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *OrdenRetail) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type Codigo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cod int64 `protobuf:"varint,1,opt,name=cod,proto3" json:"cod,omitempty"`
}

func (x *Codigo) Reset() {
	*x = Codigo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_courier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Codigo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Codigo) ProtoMessage() {}

func (x *Codigo) ProtoReflect() protoreflect.Message {
	mi := &file_courier_courier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Codigo.ProtoReflect.Descriptor instead.
func (*Codigo) Descriptor() ([]byte, []int) {
	return file_courier_courier_proto_rawDescGZIP(), []int{3}
}

func (x *Codigo) GetCod() int64 {
	if x != nil {
		return x.Cod
	}
	return 0
}

//Si, existe el paquete que importa el mensaje Empty creado
//por Google, pero daba problemas de implementacion
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_courier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_courier_courier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_courier_courier_proto_rawDescGZIP(), []int{4}
}

type Paquete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tipo    string `protobuf:"bytes,2,opt,name=tipo,proto3" json:"tipo,omitempty"`
	Valor   int64  `protobuf:"varint,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Tienda  string `protobuf:"bytes,4,opt,name=tienda,proto3" json:"tienda,omitempty"`
	Estado  string `protobuf:"bytes,5,opt,name=estado,proto3" json:"estado,omitempty"`
	Destino string `protobuf:"bytes,6,opt,name=destino,proto3" json:"destino,omitempty"`
}

func (x *Paquete) Reset() {
	*x = Paquete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_courier_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Paquete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Paquete) ProtoMessage() {}

func (x *Paquete) ProtoReflect() protoreflect.Message {
	mi := &file_courier_courier_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Paquete.ProtoReflect.Descriptor instead.
func (*Paquete) Descriptor() ([]byte, []int) {
	return file_courier_courier_proto_rawDescGZIP(), []int{5}
}

func (x *Paquete) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Paquete) GetTipo() string {
	if x != nil {
		return x.Tipo
	}
	return ""
}

func (x *Paquete) GetValor() int64 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *Paquete) GetTienda() string {
	if x != nil {
		return x.Tienda
	}
	return ""
}

func (x *Paquete) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *Paquete) GetDestino() string {
	if x != nil {
		return x.Destino
	}
	return ""
}

type Entrega struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Dignipesos   int64  `protobuf:"varint,2,opt,name=dignipesos,proto3" json:"dignipesos,omitempty"`
	Estado       string `protobuf:"bytes,3,opt,name=estado,proto3" json:"estado,omitempty"`
	Intentos     int64  `protobuf:"varint,4,opt,name=intentos,proto3" json:"intentos,omitempty"`
	Fechaentrega string `protobuf:"bytes,5,opt,name=fechaentrega,proto3" json:"fechaentrega,omitempty"`
	Camion       string `protobuf:"bytes,6,opt,name=camion,proto3" json:"camion,omitempty"`
}

func (x *Entrega) Reset() {
	*x = Entrega{}
	if protoimpl.UnsafeEnabled {
		mi := &file_courier_courier_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entrega) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entrega) ProtoMessage() {}

func (x *Entrega) ProtoReflect() protoreflect.Message {
	mi := &file_courier_courier_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entrega.ProtoReflect.Descriptor instead.
func (*Entrega) Descriptor() ([]byte, []int) {
	return file_courier_courier_proto_rawDescGZIP(), []int{6}
}

func (x *Entrega) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Entrega) GetDignipesos() int64 {
	if x != nil {
		return x.Dignipesos
	}
	return 0
}

func (x *Entrega) GetEstado() string {
	if x != nil {
		return x.Estado
	}
	return ""
}

func (x *Entrega) GetIntentos() int64 {
	if x != nil {
		return x.Intentos
	}
	return 0
}

func (x *Entrega) GetFechaentrega() string {
	if x != nil {
		return x.Fechaentrega
	}
	return ""
}

func (x *Entrega) GetCamion() string {
	if x != nil {
		return x.Camion
	}
	return ""
}

var File_courier_courier_proto protoreflect.FileDescriptor

var file_courier_courier_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x07, 0x4d, 0x65, 0x6e, 0x73, 0x61,
	0x6a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x9d, 0x01, 0x0a, 0x09, 0x4f, 0x72, 0x64, 0x65, 0x6e,
	0x50, 0x79, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x6f,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6f,
	0x72, 0x69, 0x64, 0x61, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x64, 0x61, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x0b, 0x4f, 0x72, 0x64, 0x65, 0x6e,
	0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x65, 0x6e,
	0x64, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x22, 0x1a, 0x0a, 0x06, 0x43, 0x6f,
	0x64, 0x69, 0x67, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x03, 0x63, 0x6f, 0x64, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x8d, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x69, 0x70, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x65, 0x6e, 0x64, 0x61, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x6f, 0x22,
	0xa9, 0x01, 0x0a, 0x07, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x64,
	0x69, 0x67, 0x6e, 0x69, 0x70, 0x65, 0x73, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x64, 0x69, 0x67, 0x6e, 0x69, 0x70, 0x65, 0x73, 0x6f, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65,
	0x73, 0x74, 0x61, 0x64, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x73, 0x74,
	0x61, 0x64, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x6f, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x66, 0x65, 0x63, 0x68, 0x61, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x61, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x65, 0x63, 0x68, 0x61, 0x65, 0x6e, 0x74, 0x72,
	0x65, 0x67, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x6d, 0x69, 0x6f, 0x6e, 0x32, 0xa7, 0x02, 0x0a, 0x0e,
	0x43, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x28,
	0x0a, 0x0f, 0x45, 0x6e, 0x76, 0x69, 0x61, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x50, 0x79, 0x6d,
	0x65, 0x12, 0x0a, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x50, 0x79, 0x6d, 0x65, 0x1a, 0x07, 0x2e,
	0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x11, 0x45, 0x6e, 0x76, 0x69,
	0x61, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0c, 0x2e,
	0x4f, 0x72, 0x64, 0x65, 0x6e, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x22, 0x0a, 0x0b, 0x53, 0x65, 0x67, 0x75, 0x69, 0x6d, 0x69,
	0x65, 0x6e, 0x74, 0x6f, 0x12, 0x07, 0x2e, 0x43, 0x6f, 0x64, 0x69, 0x67, 0x6f, 0x1a, 0x08, 0x2e,
	0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0b, 0x50, 0x65, 0x64,
	0x69, 0x72, 0x52, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x08, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61,
	0x6a, 0x65, 0x1a, 0x08, 0x2e, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x22, 0x00, 0x12, 0x28,
	0x0a, 0x10, 0x50, 0x65, 0x64, 0x69, 0x72, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x61, 0x72,
	0x69, 0x6f, 0x12, 0x08, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a, 0x65, 0x1a, 0x08, 0x2e, 0x50,
	0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x0b, 0x50, 0x65, 0x64, 0x69,
	0x72, 0x4e, 0x6f, 0x72, 0x6d, 0x61, 0x6c, 0x12, 0x08, 0x2e, 0x4d, 0x65, 0x6e, 0x73, 0x61, 0x6a,
	0x65, 0x1a, 0x08, 0x2e, 0x50, 0x61, 0x71, 0x75, 0x65, 0x74, 0x65, 0x22, 0x00, 0x12, 0x26, 0x0a,
	0x10, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x64, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x67,
	0x61, 0x12, 0x08, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x65, 0x67, 0x61, 0x1a, 0x06, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6e, 0x6f, 0x72, 0x61, 0x6d, 0x62, 0x75, 0x2f, 0x70, 0x72, 0x75,
	0x65, 0x62, 0x61, 0x73, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x69, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_courier_courier_proto_rawDescOnce sync.Once
	file_courier_courier_proto_rawDescData = file_courier_courier_proto_rawDesc
)

func file_courier_courier_proto_rawDescGZIP() []byte {
	file_courier_courier_proto_rawDescOnce.Do(func() {
		file_courier_courier_proto_rawDescData = protoimpl.X.CompressGZIP(file_courier_courier_proto_rawDescData)
	})
	return file_courier_courier_proto_rawDescData
}

var file_courier_courier_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_courier_courier_proto_goTypes = []interface{}{
	(*Mensaje)(nil),     // 0: Mensaje
	(*OrdenPyme)(nil),   // 1: OrdenPyme
	(*OrdenRetail)(nil), // 2: OrdenRetail
	(*Codigo)(nil),      // 3: Codigo
	(*Empty)(nil),       // 4: Empty
	(*Paquete)(nil),     // 5: Paquete
	(*Entrega)(nil),     // 6: Entrega
}
var file_courier_courier_proto_depIdxs = []int32{
	1, // 0: CourierService.EnviarOrdenPyme:input_type -> OrdenPyme
	2, // 1: CourierService.EnviarOrdenRetail:input_type -> OrdenRetail
	3, // 2: CourierService.Seguimiento:input_type -> Codigo
	0, // 3: CourierService.PedirRetail:input_type -> Mensaje
	0, // 4: CourierService.PedirPrioritario:input_type -> Mensaje
	0, // 5: CourierService.PedirNormal:input_type -> Mensaje
	6, // 6: CourierService.ResultadoEntrega:input_type -> Entrega
	3, // 7: CourierService.EnviarOrdenPyme:output_type -> Codigo
	4, // 8: CourierService.EnviarOrdenRetail:output_type -> Empty
	0, // 9: CourierService.Seguimiento:output_type -> Mensaje
	5, // 10: CourierService.PedirRetail:output_type -> Paquete
	5, // 11: CourierService.PedirPrioritario:output_type -> Paquete
	5, // 12: CourierService.PedirNormal:output_type -> Paquete
	4, // 13: CourierService.ResultadoEntrega:output_type -> Empty
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_courier_courier_proto_init() }
func file_courier_courier_proto_init() {
	if File_courier_courier_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_courier_courier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Mensaje); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_courier_courier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdenPyme); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_courier_courier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdenRetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_courier_courier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Codigo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_courier_courier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_courier_courier_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Paquete); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_courier_courier_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entrega); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_courier_courier_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_courier_courier_proto_goTypes,
		DependencyIndexes: file_courier_courier_proto_depIdxs,
		MessageInfos:      file_courier_courier_proto_msgTypes,
	}.Build()
	File_courier_courier_proto = out.File
	file_courier_courier_proto_rawDesc = nil
	file_courier_courier_proto_goTypes = nil
	file_courier_courier_proto_depIdxs = nil
}

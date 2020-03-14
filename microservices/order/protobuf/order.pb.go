// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package orderpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Order struct {
	UUID                 string               `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	OrderedProducts      []*OrderedProduct    `protobuf:"bytes,2,rep,name=orderedProducts,proto3" json:"orderedProducts,omitempty"`
	UserUUID             string               `protobuf:"bytes,3,opt,name=userUUID,proto3" json:"userUUID,omitempty"`
	PaymentInfoUUID      string               `protobuf:"bytes,4,opt,name=paymentInfoUUID,proto3" json:"paymentInfoUUID,omitempty"`
	AddressUUID          string               `protobuf:"bytes,5,opt,name=addressUUID,proto3" json:"addressUUID,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,8,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{0}
}

func (m *Order) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Order.Unmarshal(m, b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Order.Marshal(b, m, deterministic)
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return xxx_messageInfo_Order.Size(m)
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func (m *Order) GetOrderedProducts() []*OrderedProduct {
	if m != nil {
		return m.OrderedProducts
	}
	return nil
}

func (m *Order) GetUserUUID() string {
	if m != nil {
		return m.UserUUID
	}
	return ""
}

func (m *Order) GetPaymentInfoUUID() string {
	if m != nil {
		return m.PaymentInfoUUID
	}
	return ""
}

func (m *Order) GetAddressUUID() string {
	if m != nil {
		return m.AddressUUID
	}
	return ""
}

func (m *Order) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Order) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *Order) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type OrderedProduct struct {
	OrderUUID            string               `protobuf:"bytes,1,opt,name=orderUUID,proto3" json:"orderUUID,omitempty"`
	ProductUUID          string               `protobuf:"bytes,2,opt,name=productUUID,proto3" json:"productUUID,omitempty"`
	Count                int32                `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Price                int32                `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt            *timestamp.Timestamp `protobuf:"bytes,7,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderedProduct) Reset()         { *m = OrderedProduct{} }
func (m *OrderedProduct) String() string { return proto.CompactTextString(m) }
func (*OrderedProduct) ProtoMessage()    {}
func (*OrderedProduct) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{1}
}

func (m *OrderedProduct) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderedProduct.Unmarshal(m, b)
}
func (m *OrderedProduct) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderedProduct.Marshal(b, m, deterministic)
}
func (m *OrderedProduct) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderedProduct.Merge(m, src)
}
func (m *OrderedProduct) XXX_Size() int {
	return xxx_messageInfo_OrderedProduct.Size(m)
}
func (m *OrderedProduct) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderedProduct.DiscardUnknown(m)
}

var xxx_messageInfo_OrderedProduct proto.InternalMessageInfo

func (m *OrderedProduct) GetOrderUUID() string {
	if m != nil {
		return m.OrderUUID
	}
	return ""
}

func (m *OrderedProduct) GetProductUUID() string {
	if m != nil {
		return m.ProductUUID
	}
	return ""
}

func (m *OrderedProduct) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *OrderedProduct) GetPrice() int32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *OrderedProduct) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *OrderedProduct) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *OrderedProduct) GetDeletedAt() *timestamp.Timestamp {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

type GetRequest struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type GetResponse struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{3}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type SetRequest struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetRequest) Reset()         { *m = SetRequest{} }
func (m *SetRequest) String() string { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()    {}
func (*SetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{4}
}

func (m *SetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetRequest.Unmarshal(m, b)
}
func (m *SetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetRequest.Marshal(b, m, deterministic)
}
func (m *SetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetRequest.Merge(m, src)
}
func (m *SetRequest) XXX_Size() int {
	return xxx_messageInfo_SetRequest.Size(m)
}
func (m *SetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetRequest proto.InternalMessageInfo

func (m *SetRequest) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type SetResponse struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetResponse) Reset()         { *m = SetResponse{} }
func (m *SetResponse) String() string { return proto.CompactTextString(m) }
func (*SetResponse) ProtoMessage()    {}
func (*SetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{5}
}

func (m *SetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetResponse.Unmarshal(m, b)
}
func (m *SetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetResponse.Marshal(b, m, deterministic)
}
func (m *SetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetResponse.Merge(m, src)
}
func (m *SetResponse) XXX_Size() int {
	return xxx_messageInfo_SetResponse.Size(m)
}
func (m *SetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetResponse proto.InternalMessageInfo

func (m *SetResponse) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type UpdateRequest struct {
	Order                *Order   `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{6}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetOrder() *Order {
	if m != nil {
		return m.Order
	}
	return nil
}

type DeleteRequest struct {
	UUID                 string   `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cd01338c35d87077, []int{7}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

func init() {
	proto.RegisterType((*Order)(nil), "orderpb.Order")
	proto.RegisterType((*OrderedProduct)(nil), "orderpb.OrderedProduct")
	proto.RegisterType((*GetRequest)(nil), "orderpb.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "orderpb.GetResponse")
	proto.RegisterType((*SetRequest)(nil), "orderpb.SetRequest")
	proto.RegisterType((*SetResponse)(nil), "orderpb.SetResponse")
	proto.RegisterType((*UpdateRequest)(nil), "orderpb.UpdateRequest")
	proto.RegisterType((*DeleteRequest)(nil), "orderpb.DeleteRequest")
}

func init() {
	proto.RegisterFile("order.proto", fileDescriptor_cd01338c35d87077)
}

var fileDescriptor_cd01338c35d87077 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0xcd, 0x6a, 0xdb, 0x40,
	0x14, 0x85, 0x6b, 0x3b, 0xf2, 0xcf, 0x15, 0x49, 0x60, 0x1a, 0x52, 0xa3, 0x16, 0xaa, 0xaa, 0x5d,
	0x78, 0xa5, 0x80, 0x42, 0xa1, 0x74, 0x67, 0x48, 0x09, 0x5e, 0x25, 0x48, 0xf5, 0x03, 0xc8, 0x9a,
	0x9b, 0x10, 0x88, 0x34, 0xd3, 0x99, 0xd1, 0x22, 0xaf, 0xd4, 0x47, 0xeb, 0xba, 0x0f, 0x50, 0x74,
	0x27, 0x91, 0x34, 0xae, 0xdb, 0xfc, 0xec, 0x3c, 0xe7, 0x9e, 0xef, 0x8e, 0xce, 0x91, 0x30, 0xf8,
	0x42, 0x71, 0x54, 0xb1, 0x54, 0xc2, 0x08, 0x36, 0xa1, 0x83, 0xdc, 0x04, 0x6f, 0xaf, 0x85, 0xb8,
	0xbe, 0xc5, 0x13, 0x92, 0x37, 0xf5, 0xd5, 0x09, 0x96, 0xd2, 0xdc, 0x59, 0x57, 0xf0, 0x7e, 0x7b,
	0x68, 0x6e, 0x4a, 0xd4, 0x26, 0x2f, 0xa5, 0x35, 0x44, 0xbf, 0x87, 0xe0, 0x5d, 0x34, 0x9b, 0x18,
	0x83, 0xbd, 0xf5, 0x7a, 0x75, 0x36, 0x1f, 0x84, 0x83, 0xc5, 0x2c, 0xa5, 0xdf, 0x6c, 0x09, 0x87,
	0x74, 0x0d, 0xf2, 0x4b, 0x25, 0x78, 0x5d, 0x18, 0x3d, 0x1f, 0x86, 0xa3, 0x85, 0x9f, 0xbc, 0x89,
	0xef, 0xaf, 0x8f, 0x2f, 0x9c, 0x79, 0xba, 0xed, 0x67, 0x01, 0x4c, 0x6b, 0x8d, 0x8a, 0x56, 0x8f,
	0x68, 0x75, 0x7b, 0x66, 0x0b, 0x38, 0x94, 0xf9, 0x5d, 0x89, 0x95, 0x59, 0x55, 0x57, 0x82, 0x2c,
	0x7b, 0x64, 0xd9, 0x96, 0x59, 0x08, 0x7e, 0xce, 0xb9, 0x42, 0xad, 0xc9, 0xe5, 0x91, 0xab, 0x2f,
	0xb1, 0x2f, 0x30, 0x2b, 0x14, 0xe6, 0x06, 0xf9, 0xd2, 0xcc, 0xc7, 0xe1, 0x60, 0xe1, 0x27, 0x41,
	0x6c, 0xd3, 0xc7, 0x0f, 0xe9, 0xe3, 0xef, 0x0f, 0xe9, 0xd3, 0xce, 0xdc, 0x90, 0xb5, 0xe4, 0xf7,
	0xe4, 0xe4, 0x71, 0xb2, 0x35, 0x37, 0x24, 0xc7, 0x5b, 0xb4, 0xe4, 0xf4, 0x71, 0xb2, 0x35, 0x47,
	0x3f, 0x87, 0x70, 0xe0, 0x36, 0xc7, 0xde, 0xc1, 0x8c, 0xba, 0xeb, 0xbd, 0x84, 0x4e, 0x68, 0x0a,
	0x90, 0xd6, 0x48, 0xf3, 0xa1, 0x2d, 0xa0, 0x27, 0xb1, 0x23, 0xf0, 0x0a, 0x51, 0x57, 0x86, 0x5a,
	0xf6, 0x52, 0x7b, 0x68, 0x54, 0xa9, 0x6e, 0x0a, 0xa4, 0x62, 0xbd, 0xd4, 0x1e, 0xdc, 0xb2, 0xbc,
	0x17, 0x97, 0x35, 0x7e, 0x71, 0x59, 0x93, 0xe7, 0x94, 0x15, 0x02, 0x9c, 0xa3, 0x49, 0xf1, 0x47,
	0x8d, 0xda, 0xec, 0xfa, 0x4e, 0xa3, 0x53, 0xf0, 0xc9, 0xa1, 0xa5, 0xa8, 0x34, 0xb2, 0x4f, 0xe0,
	0x51, 0x73, 0xe4, 0xf1, 0x93, 0x03, 0xf7, 0x63, 0x4d, 0xed, 0x30, 0x4a, 0x00, 0xb2, 0x6e, 0xed,
	0xd3, 0x98, 0x0f, 0xe0, 0x67, 0xbd, 0x8b, 0x76, 0x3d, 0xcb, 0x67, 0xd8, 0x5f, 0x53, 0xe8, 0xe7,
	0x6d, 0xfe, 0x08, 0xfb, 0x67, 0x94, 0xf8, 0x3f, 0x39, 0x93, 0x5f, 0x03, 0x98, 0x12, 0xb5, 0xbc,
	0x5c, 0xb1, 0x04, 0x46, 0xe7, 0x68, 0xd8, 0xeb, 0x76, 0x5f, 0x57, 0x52, 0x70, 0xe4, 0x8a, 0xf6,
	0x71, 0xa3, 0x57, 0x0d, 0x93, 0x39, 0x4c, 0xb6, 0x8b, 0xc9, 0x1c, 0xe6, 0x2b, 0x8c, 0x6d, 0x20,
	0x76, 0xdc, 0x3a, 0x9c, 0x84, 0xc1, 0xf1, 0x5f, 0xef, 0xf1, 0x5b, 0xf3, 0x1f, 0x64, 0x59, 0x9b,
	0xaa, 0xc7, 0x3a, 0x31, 0xff, 0xcd, 0x6e, 0xc6, 0xa4, 0x9c, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xa3, 0xe2, 0xfc, 0xf0, 0xf7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// OrderAPIClient is the client API for OrderAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderAPIClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type orderAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderAPIClient(cc grpc.ClientConnInterface) OrderAPIClient {
	return &orderAPIClient{cc}
}

func (c *orderAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/orderpb.OrderAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAPIClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/orderpb.OrderAPI/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAPIClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/orderpb.OrderAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderAPIClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/orderpb.OrderAPI/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderAPIServer is the server API for OrderAPI service.
type OrderAPIServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Update(context.Context, *UpdateRequest) (*empty.Empty, error)
	Delete(context.Context, *DeleteRequest) (*empty.Empty, error)
}

// UnimplementedOrderAPIServer can be embedded to have forward compatible implementations.
type UnimplementedOrderAPIServer struct {
}

func (*UnimplementedOrderAPIServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedOrderAPIServer) Set(ctx context.Context, req *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (*UnimplementedOrderAPIServer) Update(ctx context.Context, req *UpdateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedOrderAPIServer) Delete(ctx context.Context, req *DeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterOrderAPIServer(s *grpc.Server, srv OrderAPIServer) {
	s.RegisterService(&_OrderAPI_serviceDesc, srv)
}

func _OrderAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderpb.OrderAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderAPI_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAPIServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderpb.OrderAPI/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAPIServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderpb.OrderAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAPIServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderAPI_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderAPIServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/orderpb.OrderAPI/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderAPIServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "orderpb.OrderAPI",
	HandlerType: (*OrderAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _OrderAPI_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _OrderAPI_Set_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OrderAPI_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OrderAPI_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: servicea.proto

package test

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Atest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Atest) Reset()         { *m = Atest{} }
func (m *Atest) String() string { return proto.CompactTextString(m) }
func (*Atest) ProtoMessage()    {}
func (*Atest) Descriptor() ([]byte, []int) {
	return fileDescriptor_42782aa172659b74, []int{0}
}

func (m *Atest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Atest.Unmarshal(m, b)
}
func (m *Atest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Atest.Marshal(b, m, deterministic)
}
func (m *Atest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Atest.Merge(m, src)
}
func (m *Atest) XXX_Size() int {
	return xxx_messageInfo_Atest.Size(m)
}
func (m *Atest) XXX_DiscardUnknown() {
	xxx_messageInfo_Atest.DiscardUnknown(m)
}

var xxx_messageInfo_Atest proto.InternalMessageInfo

func (m *Atest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Atest)(nil), "test.Atest")
}

func init() { proto.RegisterFile("servicea.proto", fileDescriptor_42782aa172659b74) }

var fileDescriptor_42782aa172659b74 = []byte{
	// 104 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x4d, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x29, 0x49, 0x2d, 0x2e,
	0x51, 0x92, 0xe6, 0x62, 0x75, 0x04, 0x31, 0x84, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x23, 0x3d, 0x2e, 0x8e, 0x60, 0x88, 0x26, 0x47,
	0x21, 0x25, 0x2e, 0x16, 0xe7, 0xc4, 0x9c, 0x1c, 0x21, 0x6e, 0x3d, 0x90, 0x72, 0x3d, 0xb0, 0x26,
	0x29, 0x64, 0x8e, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0x64, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x34, 0xd8, 0x5f, 0x4f, 0x6b, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceAClient is the client API for ServiceA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceAClient interface {
	Call(ctx context.Context, in *Atest, opts ...grpc.CallOption) (*Atest, error)
}

type serviceAClient struct {
	cc *grpc.ClientConn
}

func NewServiceAClient(cc *grpc.ClientConn) ServiceAClient {
	return &serviceAClient{cc}
}

func (c *serviceAClient) Call(ctx context.Context, in *Atest, opts ...grpc.CallOption) (*Atest, error) {
	out := new(Atest)
	err := c.cc.Invoke(ctx, "/test.ServiceA/Call", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceAServer is the server API for ServiceA service.
type ServiceAServer interface {
	Call(context.Context, *Atest) (*Atest, error)
}

// UnimplementedServiceAServer can be embedded to have forward compatible implementations.
type UnimplementedServiceAServer struct {
}

func (*UnimplementedServiceAServer) Call(ctx context.Context, req *Atest) (*Atest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Call not implemented")
}

func RegisterServiceAServer(s *grpc.Server, srv ServiceAServer) {
	s.RegisterService(&_ServiceA_serviceDesc, srv)
}

func _ServiceA_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Atest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceAServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/test.ServiceA/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceAServer).Call(ctx, req.(*Atest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "test.ServiceA",
	HandlerType: (*ServiceAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _ServiceA_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servicea.proto",
}
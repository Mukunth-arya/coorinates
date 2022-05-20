// Code generated by protoc-gen-go. DO NOT EDIT.
// source: data.proto

package data

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

type CordinateRequest struct {
	City                 string   `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CordinateRequest) Reset()         { *m = CordinateRequest{} }
func (m *CordinateRequest) String() string { return proto.CompactTextString(m) }
func (*CordinateRequest) ProtoMessage()    {}
func (*CordinateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{0}
}

func (m *CordinateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CordinateRequest.Unmarshal(m, b)
}
func (m *CordinateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CordinateRequest.Marshal(b, m, deterministic)
}
func (m *CordinateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CordinateRequest.Merge(m, src)
}
func (m *CordinateRequest) XXX_Size() int {
	return xxx_messageInfo_CordinateRequest.Size(m)
}
func (m *CordinateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CordinateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CordinateRequest proto.InternalMessageInfo

func (m *CordinateRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

type CoordinateResponse struct {
	CoordinateData       float64  `protobuf:"fixed64,1,opt,name=CoordinateData,proto3" json:"CoordinateData,omitempty"`
	LatLongi             float64  `protobuf:"fixed64,2,opt,name=lat_longi,json=latLongi,proto3" json:"lat_longi,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CoordinateResponse) Reset()         { *m = CoordinateResponse{} }
func (m *CoordinateResponse) String() string { return proto.CompactTextString(m) }
func (*CoordinateResponse) ProtoMessage()    {}
func (*CoordinateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_871986018790d2fd, []int{1}
}

func (m *CoordinateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoordinateResponse.Unmarshal(m, b)
}
func (m *CoordinateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoordinateResponse.Marshal(b, m, deterministic)
}
func (m *CoordinateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoordinateResponse.Merge(m, src)
}
func (m *CoordinateResponse) XXX_Size() int {
	return xxx_messageInfo_CoordinateResponse.Size(m)
}
func (m *CoordinateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CoordinateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CoordinateResponse proto.InternalMessageInfo

func (m *CoordinateResponse) GetCoordinateData() float64 {
	if m != nil {
		return m.CoordinateData
	}
	return 0
}

func (m *CoordinateResponse) GetLatLongi() float64 {
	if m != nil {
		return m.LatLongi
	}
	return 0
}

func init() {
	proto.RegisterType((*CordinateRequest)(nil), "CordinateRequest")
	proto.RegisterType((*CoordinateResponse)(nil), "CoordinateResponse")
}

func init() { proto.RegisterFile("data.proto", fileDescriptor_871986018790d2fd) }

var fileDescriptor_871986018790d2fd = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0x49,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe3, 0x12, 0x70, 0xce, 0x2f, 0x4a, 0xc9, 0xcc,
	0x4b, 0x2c, 0x49, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x49, 0xce,
	0x2c, 0xa9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x22, 0xb9, 0x84, 0x9c,
	0xf3, 0x11, 0x0a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xd4, 0xb8, 0xf8, 0x10, 0xa2, 0x2e,
	0x89, 0x25, 0x89, 0x60, 0x3d, 0x8c, 0x41, 0x68, 0xa2, 0x42, 0xd2, 0x5c, 0x9c, 0x39, 0x89, 0x25,
	0xf1, 0x39, 0xf9, 0x79, 0xe9, 0x99, 0x12, 0x4c, 0x60, 0x25, 0x1c, 0x39, 0x89, 0x25, 0x3e, 0x20,
	0xbe, 0x91, 0x2f, 0x97, 0x20, 0x42, 0x79, 0x71, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90, 0x05,
	0xb2, 0xc9, 0x20, 0xf7, 0x0a, 0x09, 0xea, 0xa1, 0x3b, 0x54, 0x4a, 0x58, 0x0f, 0xd3, 0x4d, 0x49,
	0x6c, 0x60, 0x8f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc1, 0xa2, 0x4e, 0xfa, 0xe6, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CoordinateserviceClient is the client API for Coordinateservice service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CoordinateserviceClient interface {
	Coordinatedata(ctx context.Context, in *CordinateRequest, opts ...grpc.CallOption) (*CoordinateResponse, error)
}

type coordinateserviceClient struct {
	cc *grpc.ClientConn
}

func NewCoordinateserviceClient(cc *grpc.ClientConn) CoordinateserviceClient {
	return &coordinateserviceClient{cc}
}

func (c *coordinateserviceClient) Coordinatedata(ctx context.Context, in *CordinateRequest, opts ...grpc.CallOption) (*CoordinateResponse, error) {
	out := new(CoordinateResponse)
	err := c.cc.Invoke(ctx, "/Coordinateservice/Coordinatedata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoordinateserviceServer is the server API for Coordinateservice service.
type CoordinateserviceServer interface {
	Coordinatedata(context.Context, *CordinateRequest) (*CoordinateResponse, error)
}

// UnimplementedCoordinateserviceServer can be embedded to have forward compatible implementations.
type UnimplementedCoordinateserviceServer struct {
}

func (*UnimplementedCoordinateserviceServer) Coordinatedata(ctx context.Context, req *CordinateRequest) (*CoordinateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Coordinatedata not implemented")
}

func RegisterCoordinateserviceServer(s *grpc.Server, srv CoordinateserviceServer) {
	s.RegisterService(&_Coordinateservice_serviceDesc, srv)
}

func _Coordinateservice_Coordinatedata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CordinateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoordinateserviceServer).Coordinatedata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Coordinateservice/Coordinatedata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoordinateserviceServer).Coordinatedata(ctx, req.(*CordinateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Coordinateservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Coordinateservice",
	HandlerType: (*CoordinateserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Coordinatedata",
			Handler:    _Coordinateservice_Coordinatedata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data.proto",
}

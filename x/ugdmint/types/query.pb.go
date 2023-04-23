// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/ugdmint/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c1ad1733cf2faab, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c1ad1733cf2faab, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QuerySubsidyHalvingIntervalRequest is the request type for the Query/SubsidyHalvingInterval RPC method.
type QuerySubsidyHalvingIntervalRequest struct {
}

func (m *QuerySubsidyHalvingIntervalRequest) Reset()         { *m = QuerySubsidyHalvingIntervalRequest{} }
func (m *QuerySubsidyHalvingIntervalRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySubsidyHalvingIntervalRequest) ProtoMessage()    {}
func (*QuerySubsidyHalvingIntervalRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c1ad1733cf2faab, []int{2}
}
func (m *QuerySubsidyHalvingIntervalRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySubsidyHalvingIntervalRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubsidyHalvingIntervalRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySubsidyHalvingIntervalRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySubsidyHalvingIntervalRequest.Merge(m, src)
}
func (m *QuerySubsidyHalvingIntervalRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySubsidyHalvingIntervalRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySubsidyHalvingIntervalRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySubsidyHalvingIntervalRequest proto.InternalMessageInfo

// QuerySubsidyHalvingIntervalResponse is the response type for the Query/SubsidyHalvingInterval RPC method.
type QuerySubsidyHalvingIntervalResponse struct {
	// subsidy halving interval value.
	SubsidyHalvingInterval github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=subsidy_halving_interval,json=subsidyHalvingInterval,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"subsidy_halving_interval"`
}

func (m *QuerySubsidyHalvingIntervalResponse) Reset()         { *m = QuerySubsidyHalvingIntervalResponse{} }
func (m *QuerySubsidyHalvingIntervalResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySubsidyHalvingIntervalResponse) ProtoMessage()    {}
func (*QuerySubsidyHalvingIntervalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c1ad1733cf2faab, []int{3}
}
func (m *QuerySubsidyHalvingIntervalResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySubsidyHalvingIntervalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySubsidyHalvingIntervalResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySubsidyHalvingIntervalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySubsidyHalvingIntervalResponse.Merge(m, src)
}
func (m *QuerySubsidyHalvingIntervalResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySubsidyHalvingIntervalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySubsidyHalvingIntervalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySubsidyHalvingIntervalResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "cosmos.ugdmint.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "cosmos.ugdmint.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QuerySubsidyHalvingIntervalRequest)(nil), "cosmos.ugdmint.v1beta1.QuerySubsidyHalvingIntervalRequest")
	proto.RegisterType((*QuerySubsidyHalvingIntervalResponse)(nil), "cosmos.ugdmint.v1beta1.QuerySubsidyHalvingIntervalResponse")
}

func init() { proto.RegisterFile("cosmos/ugdmint/v1beta1/query.proto", fileDescriptor_2c1ad1733cf2faab) }

var fileDescriptor_2c1ad1733cf2faab = []byte{
	// 426 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x4f, 0x8b, 0x13, 0x31,
	0x18, 0xc6, 0x67, 0x8a, 0xf6, 0x10, 0xc5, 0x43, 0x2c, 0xa5, 0x8c, 0x32, 0xd5, 0xa9, 0x88, 0x8a,
	0x4e, 0x68, 0xab, 0x27, 0x6f, 0xa5, 0xf8, 0x07, 0x11, 0xb4, 0xe2, 0xc5, 0x4b, 0xc9, 0xb4, 0x21,
	0x0d, 0xed, 0x24, 0xd3, 0x49, 0xa6, 0x58, 0xf0, 0xe4, 0x17, 0x50, 0xf0, 0xfb, 0x78, 0x94, 0x1e,
	0x0b, 0x5e, 0x96, 0x3d, 0x94, 0xa5, 0xdd, 0x0f, 0xb2, 0x4c, 0x92, 0x5d, 0xb6, 0x7f, 0x76, 0x60,
	0x4f, 0x2d, 0x79, 0x7f, 0xef, 0xf3, 0x3e, 0xef, 0xfb, 0x0c, 0xf0, 0x06, 0x42, 0xc6, 0x42, 0xa2,
	0x8c, 0x0e, 0x63, 0xc6, 0x15, 0x9a, 0x66, 0x24, 0x9d, 0x87, 0x49, 0x2a, 0x94, 0x80, 0x77, 0x4c,
	0x2d, 0xb4, 0x35, 0xaf, 0x42, 0x05, 0x15, 0xba, 0x84, 0xf2, 0x7f, 0x86, 0xf2, 0xee, 0x53, 0x21,
	0xe8, 0x84, 0x20, 0x9c, 0x30, 0x84, 0x39, 0x17, 0x0a, 0x2b, 0x26, 0xb8, 0xb4, 0xd5, 0x67, 0x56,
	0x3f, 0xc2, 0x92, 0x18, 0x71, 0x34, 0x6b, 0x46, 0x44, 0xe1, 0x26, 0x4a, 0x30, 0x65, 0x5c, 0xc3,
	0x96, 0xbd, 0xb7, 0xe3, 0x25, 0xc1, 0x29, 0x8e, 0xad, 0x50, 0x50, 0x01, 0xf0, 0x73, 0xde, 0xfe,
	0x49, 0x3f, 0xf6, 0xc8, 0x34, 0x23, 0x52, 0x05, 0x1f, 0xc0, 0xdd, 0xad, 0x57, 0x99, 0x08, 0x2e,
	0x09, 0x7c, 0x09, 0xca, 0xa6, 0xb9, 0xe6, 0x3e, 0x70, 0x9f, 0xdc, 0x6a, 0x55, 0xc3, 0xed, 0x55,
	0x42, 0xc3, 0x77, 0x6e, 0x2c, 0x56, 0x75, 0xa7, 0x67, 0xd9, 0xe0, 0x11, 0x08, 0xb4, 0xd8, 0x97,
	0x2c, 0x92, 0x6c, 0x38, 0x7f, 0x87, 0x27, 0x33, 0xc6, 0xe9, 0x7b, 0xae, 0x48, 0x3a, 0xc3, 0x93,
	0xf3, 0x91, 0xbf, 0x5c, 0xd0, 0x28, 0xc4, 0xac, 0x87, 0x11, 0xa8, 0x49, 0x43, 0xf4, 0x47, 0x06,
	0xe9, 0x33, 0xcb, 0x68, 0x57, 0xb7, 0x3b, 0x61, 0x3e, 0xfd, 0x78, 0x55, 0x7f, 0x4c, 0x99, 0x1a,
	0x65, 0x51, 0x38, 0x10, 0x31, 0xb2, 0x27, 0x30, 0x3f, 0x2f, 0xe4, 0x70, 0x8c, 0xd4, 0x3c, 0x21,
	0x32, 0xec, 0x92, 0x41, 0xaf, 0x2a, 0x0f, 0x4e, 0x6c, 0xfd, 0x2b, 0x81, 0x9b, 0xda, 0x11, 0xfc,
	0x01, 0xca, 0x66, 0x33, 0x18, 0xec, 0x6e, 0xbc, 0x7f, 0x3c, 0xaf, 0x51, 0xc8, 0x98, 0x35, 0x82,
	0xa7, 0x3f, 0xff, 0x9f, 0xfe, 0x29, 0x35, 0xe0, 0x43, 0x44, 0xc7, 0x22, 0x4b, 0xb3, 0x84, 0x48,
	0xf4, 0xf5, 0x6d, 0xf7, 0x63, 0x1e, 0xd0, 0x76, 0x50, 0xf0, 0xaf, 0x0b, 0xaa, 0x87, 0x8f, 0x02,
	0x5b, 0x07, 0x47, 0x15, 0x1e, 0xda, 0x6b, 0x5f, 0xab, 0xc7, 0xda, 0x7d, 0xad, 0xed, 0xbe, 0x82,
	0xed, 0x02, 0xbb, 0x57, 0xc5, 0xd2, 0x79, 0xb3, 0x58, 0xfb, 0xee, 0x72, 0xed, 0xbb, 0x27, 0x6b,
	0xdf, 0xfd, 0xbd, 0xf1, 0x9d, 0xe5, 0xc6, 0x77, 0x8e, 0x36, 0xbe, 0xf3, 0xed, 0xf9, 0xa5, 0x88,
	0xf6, 0x85, 0xbf, 0x5f, 0x48, 0xeb, 0xb0, 0xa2, 0xb2, 0xfe, 0x64, 0xdb, 0x67, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xf5, 0xc6, 0x5e, 0x16, 0x5d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Subsidy halving interval
	SubsidyHalvingInterval(ctx context.Context, in *QuerySubsidyHalvingIntervalRequest, opts ...grpc.CallOption) (*QuerySubsidyHalvingIntervalResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/cosmos.ugdmint.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SubsidyHalvingInterval(ctx context.Context, in *QuerySubsidyHalvingIntervalRequest, opts ...grpc.CallOption) (*QuerySubsidyHalvingIntervalResponse, error) {
	out := new(QuerySubsidyHalvingIntervalResponse)
	err := c.cc.Invoke(ctx, "/cosmos.ugdmint.v1beta1.Query/SubsidyHalvingInterval", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Subsidy halving interval
	SubsidyHalvingInterval(context.Context, *QuerySubsidyHalvingIntervalRequest) (*QuerySubsidyHalvingIntervalResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) SubsidyHalvingInterval(ctx context.Context, req *QuerySubsidyHalvingIntervalRequest) (*QuerySubsidyHalvingIntervalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubsidyHalvingInterval not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.ugdmint.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SubsidyHalvingInterval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySubsidyHalvingIntervalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SubsidyHalvingInterval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.ugdmint.v1beta1.Query/SubsidyHalvingInterval",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SubsidyHalvingInterval(ctx, req.(*QuerySubsidyHalvingIntervalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.ugdmint.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "SubsidyHalvingInterval",
			Handler:    _Query_SubsidyHalvingInterval_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/ugdmint/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QuerySubsidyHalvingIntervalRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySubsidyHalvingIntervalRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySubsidyHalvingIntervalRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QuerySubsidyHalvingIntervalResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySubsidyHalvingIntervalResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySubsidyHalvingIntervalResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SubsidyHalvingInterval.Size()
		i -= size
		if _, err := m.SubsidyHalvingInterval.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QuerySubsidyHalvingIntervalRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QuerySubsidyHalvingIntervalResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SubsidyHalvingInterval.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySubsidyHalvingIntervalRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySubsidyHalvingIntervalRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubsidyHalvingIntervalRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QuerySubsidyHalvingIntervalResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QuerySubsidyHalvingIntervalResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySubsidyHalvingIntervalResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubsidyHalvingInterval", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SubsidyHalvingInterval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)

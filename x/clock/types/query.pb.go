// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: juno/clock/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// QueryClockContracts is the request type to get all contracts.
type QueryClockContracts struct {
}

func (m *QueryClockContracts) Reset()         { *m = QueryClockContracts{} }
func (m *QueryClockContracts) String() string { return proto.CompactTextString(m) }
func (*QueryClockContracts) ProtoMessage()    {}
func (*QueryClockContracts) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da208f579d775c8, []int{0}
}
func (m *QueryClockContracts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClockContracts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClockContracts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClockContracts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClockContracts.Merge(m, src)
}
func (m *QueryClockContracts) XXX_Size() int {
	return m.Size()
}
func (m *QueryClockContracts) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClockContracts.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClockContracts proto.InternalMessageInfo

// QueryClockContractsResponse is the response type for the Query/ClockContracts RPC method.
type QueryClockContractsResponse struct {
	ContractAddresses []string `protobuf:"bytes,1,rep,name=contract_addresses,json=contractAddresses,proto3" json:"contract_addresses,omitempty" yaml:"contract_addresses"`
}

func (m *QueryClockContractsResponse) Reset()         { *m = QueryClockContractsResponse{} }
func (m *QueryClockContractsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryClockContractsResponse) ProtoMessage()    {}
func (*QueryClockContractsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da208f579d775c8, []int{1}
}
func (m *QueryClockContractsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryClockContractsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryClockContractsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryClockContractsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryClockContractsResponse.Merge(m, src)
}
func (m *QueryClockContractsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryClockContractsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryClockContractsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryClockContractsResponse proto.InternalMessageInfo

func (m *QueryClockContractsResponse) GetContractAddresses() []string {
	if m != nil {
		return m.ContractAddresses
	}
	return nil
}

// QueryParams is the request type to get all module params.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da208f579d775c8, []int{2}
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

// QueryClockContractsResponse is the response type for the Query/ClockContracts RPC method.
type QueryParamsResponse struct {
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params" yaml:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7da208f579d775c8, []int{3}
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

func (m *QueryParamsResponse) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryClockContracts)(nil), "juno.clock.v1.QueryClockContracts")
	proto.RegisterType((*QueryClockContractsResponse)(nil), "juno.clock.v1.QueryClockContractsResponse")
	proto.RegisterType((*QueryParamsRequest)(nil), "juno.clock.v1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "juno.clock.v1.QueryParamsResponse")
}

func init() { proto.RegisterFile("juno/clock/v1/query.proto", fileDescriptor_7da208f579d775c8) }

var fileDescriptor_7da208f579d775c8 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xb1, 0x8f, 0xd3, 0x30,
	0x18, 0xc5, 0x6b, 0x10, 0x95, 0x30, 0x3a, 0x24, 0xcc, 0x55, 0xf4, 0xd2, 0x23, 0x2d, 0x9e, 0x4e,
	0x08, 0x62, 0xe5, 0xd8, 0x90, 0x18, 0x68, 0x07, 0x24, 0x26, 0xc8, 0xc8, 0x82, 0x1c, 0x9f, 0x15,
	0x02, 0x89, 0xbf, 0x5c, 0xec, 0x54, 0x64, 0xbd, 0x95, 0x05, 0x89, 0x7f, 0x8a, 0xf1, 0x24, 0x16,
	0xa6, 0x08, 0xb5, 0x4c, 0x37, 0xde, 0xcc, 0x80, 0xe2, 0xb8, 0x95, 0x72, 0x57, 0xc1, 0x96, 0xbc,
	0xdf, 0xcb, 0xf7, 0x9e, 0xbf, 0x18, 0x1f, 0x7c, 0xac, 0x14, 0x30, 0x91, 0x81, 0xf8, 0xc4, 0x96,
	0x21, 0x3b, 0xad, 0x64, 0x59, 0x07, 0x45, 0x09, 0x06, 0xc8, 0x5e, 0x8b, 0x02, 0x8b, 0x82, 0x65,
	0xe8, 0xed, 0x27, 0x90, 0x80, 0x25, 0xac, 0x7d, 0xea, 0x4c, 0xde, 0x61, 0x02, 0x90, 0x64, 0x92,
	0xf1, 0x22, 0x65, 0x5c, 0x29, 0x30, 0xdc, 0xa4, 0xa0, 0xb4, 0xa3, 0xbe, 0x00, 0x9d, 0x83, 0x66,
	0x31, 0xd7, 0x92, 0x2d, 0xc3, 0x58, 0x1a, 0x1e, 0x32, 0x01, 0xa9, 0x72, 0x7c, 0xd2, 0x4f, 0x4f,
	0xa4, 0x92, 0x3a, 0x75, 0x1f, 0xd3, 0x11, 0xbe, 0xff, 0xb6, 0xad, 0xb3, 0x68, 0xf1, 0x02, 0x94,
	0x29, 0xb9, 0x30, 0x9a, 0x7e, 0x41, 0x78, 0xb2, 0x43, 0x8f, 0xa4, 0x2e, 0x40, 0x69, 0x49, 0x32,
	0x4c, 0x84, 0x13, 0xdf, 0xf3, 0x93, 0x93, 0x52, 0x6a, 0x2d, 0xf5, 0x18, 0xcd, 0x6e, 0x1e, 0xdd,
	0x9e, 0xbf, 0xb8, 0x68, 0xa6, 0x87, 0xd7, 0xe9, 0x13, 0xc8, 0x53, 0x23, 0xf3, 0xc2, 0xd4, 0x97,
	0xcd, 0xf4, 0xa0, 0xe6, 0x79, 0xf6, 0x9c, 0x5e, 0x77, 0xd1, 0xe8, 0xde, 0x46, 0x7c, 0xb9, 0xd5,
	0xf6, 0x31, 0xb1, 0x65, 0xde, 0xf0, 0x92, 0xe7, 0x3a, 0x92, 0xa7, 0x95, 0xd4, 0x86, 0x72, 0x57,
	0x7d, 0xa3, 0xba, 0x6a, 0xaf, 0xf1, 0xb0, 0xb0, 0xca, 0x18, 0xcd, 0xd0, 0xd1, 0x9d, 0xe3, 0x51,
	0xd0, 0x5b, 0x71, 0xd0, 0xd9, 0xe7, 0x93, 0x8b, 0x66, 0xea, 0x8c, 0x97, 0xcd, 0x74, 0xaf, 0xeb,
	0xd3, 0xbd, 0xd3, 0xc8, 0x81, 0xe3, 0x3f, 0x08, 0xdf, 0xb2, 0x19, 0xe4, 0x0c, 0xe1, 0xbb, 0xfd,
	0x5d, 0x10, 0x7a, 0x65, 0xf0, 0x8e, 0x7d, 0x79, 0x8f, 0xff, 0xef, 0xd9, 0x14, 0xa7, 0xb3, 0xb3,
	0x1f, 0xbf, 0xbf, 0xdd, 0xf0, 0xc8, 0x98, 0xf5, 0x7f, 0x98, 0xd8, 0x26, 0x2a, 0x3c, 0xec, 0xda,
	0x93, 0x47, 0xbb, 0xe6, 0xf6, 0xd6, 0xe3, 0xd1, 0x7f, 0x59, 0x5c, 0xe4, 0x43, 0x1b, 0xf9, 0x80,
	0x8c, 0xae, 0x44, 0x76, 0xc7, 0x9f, 0xbf, 0xfa, 0xbe, 0xf2, 0xd1, 0xf9, 0xca, 0x47, 0xbf, 0x56,
	0x3e, 0xfa, 0xba, 0xf6, 0x07, 0xe7, 0x6b, 0x7f, 0xf0, 0x73, 0xed, 0x0f, 0xde, 0x3d, 0x4d, 0x52,
	0xf3, 0xa1, 0x8a, 0x03, 0x01, 0x39, 0x5b, 0xd8, 0xeb, 0xb7, 0x3d, 0x4f, 0x37, 0xea, 0xb3, 0x1b,
	0x66, 0xea, 0x42, 0xea, 0x78, 0x68, 0x2f, 0xdb, 0xb3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdb,
	0x57, 0xe2, 0x8b, 0x09, 0x03, 0x00, 0x00,
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
	// ClockContracts
	ClockContracts(ctx context.Context, in *QueryClockContracts, opts ...grpc.CallOption) (*QueryClockContractsResponse, error)
	// Params
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) ClockContracts(ctx context.Context, in *QueryClockContracts, opts ...grpc.CallOption) (*QueryClockContractsResponse, error) {
	out := new(QueryClockContractsResponse)
	err := c.cc.Invoke(ctx, "/juno.clock.v1.Query/ClockContracts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/juno.clock.v1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ClockContracts
	ClockContracts(context.Context, *QueryClockContracts) (*QueryClockContractsResponse, error)
	// Params
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) ClockContracts(ctx context.Context, req *QueryClockContracts) (*QueryClockContractsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClockContracts not implemented")
}
func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_ClockContracts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryClockContracts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ClockContracts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/juno.clock.v1.Query/ClockContracts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ClockContracts(ctx, req.(*QueryClockContracts))
	}
	return interceptor(ctx, in, info, handler)
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
		FullMethod: "/juno.clock.v1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "juno.clock.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClockContracts",
			Handler:    _Query_ClockContracts_Handler,
		},
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "juno/clock/v1/query.proto",
}

func (m *QueryClockContracts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClockContracts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClockContracts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryClockContractsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryClockContractsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryClockContractsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContractAddresses) > 0 {
		for iNdEx := len(m.ContractAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.ContractAddresses[iNdEx])
			copy(dAtA[i:], m.ContractAddresses[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.ContractAddresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
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
	if m.Params != nil {
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
	}
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
func (m *QueryClockContracts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryClockContractsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ContractAddresses) > 0 {
		for _, s := range m.ContractAddresses {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
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
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryClockContracts) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryClockContracts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClockContracts: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryClockContractsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryClockContractsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryClockContractsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddresses = append(m.ContractAddresses, string(dAtA[iNdEx:postIndex]))
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
			if m.Params == nil {
				m.Params = &Params{}
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

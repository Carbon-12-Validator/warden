// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: warden/gmp/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgSetParams defines the SetParams message type.
type MsgSetParams struct {
	// address that controls the module (defaults to x/gov).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the gmp parameters to update.
	Params *Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *MsgSetParams) Reset()         { *m = MsgSetParams{} }
func (m *MsgSetParams) String() string { return proto.CompactTextString(m) }
func (*MsgSetParams) ProtoMessage()    {}
func (*MsgSetParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_97dd76729f9c0f4f, []int{0}
}
func (m *MsgSetParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetParams.Merge(m, src)
}
func (m *MsgSetParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetParams proto.InternalMessageInfo

// MsgSetParamsResponse defines the SetParams response type.
type MsgSetParamsResponse struct {
}

func (m *MsgSetParamsResponse) Reset()         { *m = MsgSetParamsResponse{} }
func (m *MsgSetParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetParamsResponse) ProtoMessage()    {}
func (*MsgSetParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97dd76729f9c0f4f, []int{1}
}
func (m *MsgSetParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetParamsResponse.Merge(m, src)
}
func (m *MsgSetParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetParamsResponse proto.InternalMessageInfo

// MsgRelay defines the Relay message type.
type MsgBridge struct {
	// authority is the address that signs the message.
	Relayer string `protobuf:"bytes,1,opt,name=relayer,proto3" json:"relayer,omitempty"`
	// destination_chain defines the chain which this will be relayed to.
	DestinationChain string `protobuf:"bytes,2,opt,name=destination_chain,json=destinationChain,proto3" json:"destination_chain,omitempty"`
	// warden_contract_address defines the warden contract that GMP is calling.
	WardenContractAddress string `protobuf:"bytes,3,opt,name=warden_contract_address,json=wardenContractAddress,proto3" json:"warden_contract_address,omitempty"`
	// destination_contract_address defines the destination contract that warden is calling.
	DestinationContractAddress string `protobuf:"bytes,4,opt,name=destination_contract_address,json=destinationContractAddress,proto3" json:"destination_contract_address,omitempty"`
	// destination_contract_calldata defines the command to call.
	DestinationContractCalldata []byte `protobuf:"bytes,5,opt,name=destination_contract_calldata,json=destinationContractCalldata,proto3" json:"destination_contract_calldata,omitempty"`
	// token determines the IBC token that the user wants to relay via GMP.
	Token types.Coin `protobuf:"bytes,6,opt,name=token,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"token"`
}

func (m *MsgBridge) Reset()         { *m = MsgBridge{} }
func (m *MsgBridge) String() string { return proto.CompactTextString(m) }
func (*MsgBridge) ProtoMessage()    {}
func (*MsgBridge) Descriptor() ([]byte, []int) {
	return fileDescriptor_97dd76729f9c0f4f, []int{2}
}
func (m *MsgBridge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBridge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBridge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBridge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBridge.Merge(m, src)
}
func (m *MsgBridge) XXX_Size() int {
	return m.Size()
}
func (m *MsgBridge) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBridge.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBridge proto.InternalMessageInfo

// MsgRelay defines the Relay response type.
type MsgBridgeResponse struct {
}

func (m *MsgBridgeResponse) Reset()         { *m = MsgBridgeResponse{} }
func (m *MsgBridgeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBridgeResponse) ProtoMessage()    {}
func (*MsgBridgeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97dd76729f9c0f4f, []int{3}
}
func (m *MsgBridgeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBridgeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBridgeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBridgeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBridgeResponse.Merge(m, src)
}
func (m *MsgBridgeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBridgeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBridgeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBridgeResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetParams)(nil), "warden.gmp.MsgSetParams")
	proto.RegisterType((*MsgSetParamsResponse)(nil), "warden.gmp.MsgSetParamsResponse")
	proto.RegisterType((*MsgBridge)(nil), "warden.gmp.MsgBridge")
	proto.RegisterType((*MsgBridgeResponse)(nil), "warden.gmp.MsgBridgeResponse")
}

func init() { proto.RegisterFile("warden/gmp/tx.proto", fileDescriptor_97dd76729f9c0f4f) }

var fileDescriptor_97dd76729f9c0f4f = []byte{
	// 547 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xf6, 0x11, 0x12, 0x94, 0xa3, 0x03, 0xbd, 0xa6, 0xad, 0x6b, 0xa8, 0x13, 0x65, 0x8a, 0x82,
	0x62, 0x93, 0x20, 0x75, 0xa8, 0x10, 0x82, 0x64, 0x60, 0x8a, 0x84, 0x5c, 0x89, 0x81, 0x25, 0xba,
	0xd8, 0xa7, 0x8b, 0xd5, 0xf8, 0xce, 0xba, 0xbb, 0x96, 0x66, 0x43, 0x4c, 0x88, 0x89, 0x85, 0xbd,
	0x12, 0x1b, 0x53, 0x07, 0x7e, 0x44, 0xc7, 0x8a, 0x89, 0x09, 0x50, 0x32, 0x94, 0x9f, 0x81, 0xe2,
	0xbb, 0x10, 0x13, 0x05, 0x75, 0x72, 0xde, 0xfb, 0xbe, 0xf7, 0xbd, 0x97, 0xf7, 0xbe, 0x83, 0x5b,
	0x6f, 0xb0, 0x88, 0x08, 0xf3, 0x69, 0x92, 0xfa, 0xea, 0xcc, 0x4b, 0x05, 0x57, 0x1c, 0x41, 0x9d,
	0xf4, 0x68, 0x92, 0x3a, 0x15, 0xca, 0x29, 0xcf, 0xd2, 0xfe, 0xfc, 0x97, 0x66, 0x38, 0x95, 0x5c,
	0x19, 0x4d, 0x52, 0x93, 0xdd, 0x0b, 0xb9, 0x4c, 0xb8, 0x1c, 0x68, 0xba, 0x0e, 0x0c, 0xb4, 0xab,
	0x23, 0x3f, 0x91, 0xd4, 0x3f, 0x6d, 0xcf, 0x3f, 0x06, 0x70, 0x0d, 0x30, 0xc4, 0x92, 0xf8, 0xa7,
	0xed, 0x21, 0x51, 0xb8, 0xed, 0x87, 0x3c, 0x66, 0x1a, 0xaf, 0x7f, 0x00, 0x70, 0xa3, 0x2f, 0xe9,
	0x11, 0x51, 0x2f, 0xb1, 0xc0, 0x89, 0x44, 0x07, 0xb0, 0x8c, 0x4f, 0xd4, 0x88, 0x8b, 0x58, 0x4d,
	0x6c, 0x50, 0x03, 0x8d, 0x72, 0xd7, 0xfe, 0xf6, 0xb5, 0x55, 0x31, 0xed, 0x9e, 0x47, 0x91, 0x20,
	0x52, 0x1e, 0x29, 0x11, 0x33, 0x1a, 0x2c, 0xa9, 0xa8, 0x09, 0x4b, 0x69, 0xa6, 0x60, 0xdf, 0xaa,
	0x81, 0xc6, 0xdd, 0x0e, 0xf2, 0x96, 0xff, 0xd2, 0xd3, 0xda, 0x81, 0x61, 0x1c, 0xee, 0xbc, 0x3f,
	0xaf, 0x5a, 0xbf, 0xcf, 0xab, 0xd6, 0xbb, 0xeb, 0x8b, 0xe6, 0x52, 0xa3, 0xbe, 0x03, 0x2b, 0xf9,
	0x59, 0x02, 0x22, 0x53, 0xce, 0x24, 0xa9, 0x7f, 0x2e, 0xc0, 0x72, 0x5f, 0xd2, 0xae, 0x88, 0x23,
	0x4a, 0x50, 0x07, 0xde, 0x11, 0x64, 0x8c, 0x27, 0x44, 0xdc, 0x38, 0xdf, 0x82, 0x88, 0x1e, 0xc2,
	0xcd, 0x88, 0x48, 0x15, 0x33, 0xac, 0x62, 0xce, 0x06, 0xe1, 0x08, 0xc7, 0x2c, 0x1b, 0xb4, 0x1c,
	0xdc, 0xcb, 0x01, 0xbd, 0x79, 0x1e, 0x1d, 0xc0, 0x5d, 0x3d, 0xfb, 0x20, 0xe4, 0x4c, 0x09, 0x1c,
	0xaa, 0x01, 0xd6, 0xb2, 0x76, 0x21, 0x2b, 0xd9, 0xd6, 0x70, 0xcf, 0xa0, 0xa6, 0x27, 0x7a, 0x06,
	0x1f, 0xfc, 0xd3, 0x64, 0xb5, 0xf8, 0x76, 0x56, 0xec, 0xe4, 0xfb, 0xad, 0x28, 0x74, 0xe1, 0xfe,
	0x5a, 0x85, 0x10, 0x8f, 0xc7, 0x11, 0x56, 0xd8, 0x2e, 0xd6, 0x40, 0x63, 0x23, 0xb8, 0xbf, 0x46,
	0xa2, 0x67, 0x28, 0x08, 0xc3, 0xa2, 0xe2, 0xc7, 0x84, 0xd9, 0xa5, 0xec, 0x0e, 0x7b, 0x9e, 0xd9,
	0xcc, 0xdc, 0x01, 0x9e, 0x71, 0x80, 0xd7, 0xe3, 0x31, 0xeb, 0x3e, 0xba, 0xfc, 0x51, 0xb5, 0xbe,
	0xfc, 0xac, 0x36, 0x68, 0xac, 0x46, 0x27, 0x43, 0x2f, 0xe4, 0x89, 0x71, 0x95, 0xf9, 0xb4, 0x64,
	0x74, 0xec, 0xab, 0x49, 0x4a, 0x64, 0x56, 0x20, 0x03, 0xad, 0x7c, 0x58, 0xc9, 0xdf, 0x6f, 0xb1,
	0xe3, 0xfa, 0x16, 0xdc, 0xfc, 0x7b, 0xa4, 0xc5, 0xe9, 0x3a, 0x9f, 0x00, 0x2c, 0xf4, 0x25, 0x45,
	0x2f, 0x60, 0x79, 0xe9, 0x31, 0x3b, 0xef, 0x8d, 0xfc, 0xc5, 0x9d, 0xda, 0xff, 0x90, 0x85, 0x20,
	0x7a, 0x0a, 0x4b, 0xc6, 0x07, 0xdb, 0x2b, 0x5c, 0x9d, 0x76, 0xf6, 0xd7, 0xa6, 0x17, 0xf5, 0x4e,
	0xf1, 0xed, 0xf5, 0x45, 0x13, 0x74, 0x5f, 0x5d, 0x4e, 0x5d, 0x70, 0x35, 0x75, 0xc1, 0xaf, 0xa9,
	0x0b, 0x3e, 0xce, 0x5c, 0xeb, 0x6a, 0xe6, 0x5a, 0xdf, 0x67, 0xae, 0xf5, 0xfa, 0x49, 0x6e, 0x1b,
	0x5a, 0xa9, 0x95, 0x3d, 0x95, 0x90, 0x8f, 0x4d, 0xbc, 0x12, 0xfa, 0x67, 0xfa, 0x79, 0xcf, 0xf7,
	0x34, 0x2c, 0x65, 0xe0, 0xe3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x21, 0xeb, 0xbd, 0xb8, 0xf9,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SetParams sets the parameters for the gmp module.
	SetParams(ctx context.Context, in *MsgSetParams, opts ...grpc.CallOption) (*MsgSetParamsResponse, error)
	// Relay relays Warden data via GMP.
	Bridge(ctx context.Context, in *MsgBridge, opts ...grpc.CallOption) (*MsgBridgeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetParams(ctx context.Context, in *MsgSetParams, opts ...grpc.CallOption) (*MsgSetParamsResponse, error) {
	out := new(MsgSetParamsResponse)
	err := c.cc.Invoke(ctx, "/warden.gmp.Msg/SetParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Bridge(ctx context.Context, in *MsgBridge, opts ...grpc.CallOption) (*MsgBridgeResponse, error) {
	out := new(MsgBridgeResponse)
	err := c.cc.Invoke(ctx, "/warden.gmp.Msg/Bridge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SetParams sets the parameters for the gmp module.
	SetParams(context.Context, *MsgSetParams) (*MsgSetParamsResponse, error)
	// Relay relays Warden data via GMP.
	Bridge(context.Context, *MsgBridge) (*MsgBridgeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetParams(ctx context.Context, req *MsgSetParams) (*MsgSetParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetParams not implemented")
}
func (*UnimplementedMsgServer) Bridge(ctx context.Context, req *MsgBridge) (*MsgBridgeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bridge not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warden.gmp.Msg/SetParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetParams(ctx, req.(*MsgSetParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Bridge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBridge)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Bridge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warden.gmp.Msg/Bridge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Bridge(ctx, req.(*MsgBridge))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "warden.gmp.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetParams",
			Handler:    _Msg_SetParams_Handler,
		},
		{
			MethodName: "Bridge",
			Handler:    _Msg_Bridge_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warden/gmp/tx.proto",
}

func (m *MsgSetParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgBridge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBridge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBridge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Token.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.DestinationContractCalldata) > 0 {
		i -= len(m.DestinationContractCalldata)
		copy(dAtA[i:], m.DestinationContractCalldata)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DestinationContractCalldata)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DestinationContractAddress) > 0 {
		i -= len(m.DestinationContractAddress)
		copy(dAtA[i:], m.DestinationContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DestinationContractAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.WardenContractAddress) > 0 {
		i -= len(m.WardenContractAddress)
		copy(dAtA[i:], m.WardenContractAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.WardenContractAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DestinationChain) > 0 {
		i -= len(m.DestinationChain)
		copy(dAtA[i:], m.DestinationChain)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DestinationChain)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Relayer) > 0 {
		i -= len(m.Relayer)
		copy(dAtA[i:], m.Relayer)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Relayer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBridgeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBridgeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBridgeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSetParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgBridge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Relayer)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DestinationChain)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.WardenContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DestinationContractAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DestinationContractCalldata)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Token.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgBridgeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
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
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSetParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgSetParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBridge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBridge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBridge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationChain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationChain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WardenContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WardenContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DestinationContractCalldata", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DestinationContractCalldata = append(m.DestinationContractCalldata[:0], dAtA[iNdEx:postIndex]...)
			if m.DestinationContractCalldata == nil {
				m.DestinationContractCalldata = []byte{}
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Token.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBridgeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBridgeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBridgeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

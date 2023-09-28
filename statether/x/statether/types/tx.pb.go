// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: statether/statether/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

type MsgSaveEthaddressState struct {
	Creator         string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	EthAddress      string `protobuf:"bytes,2,opt,name=ethAddress,proto3" json:"ethAddress,omitempty"`
	Block           uint64 `protobuf:"varint,3,opt,name=block,proto3" json:"block,omitempty"`
	Nonce           uint64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	StoragePosition uint64 `protobuf:"varint,5,opt,name=storagePosition,proto3" json:"storagePosition,omitempty"`
	Active          bool   `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *MsgSaveEthaddressState) Reset()         { *m = MsgSaveEthaddressState{} }
func (m *MsgSaveEthaddressState) String() string { return proto.CompactTextString(m) }
func (*MsgSaveEthaddressState) ProtoMessage()    {}
func (*MsgSaveEthaddressState) Descriptor() ([]byte, []int) {
	return fileDescriptor_3454f3c906df7237, []int{0}
}
func (m *MsgSaveEthaddressState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSaveEthaddressState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSaveEthaddressState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSaveEthaddressState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSaveEthaddressState.Merge(m, src)
}
func (m *MsgSaveEthaddressState) XXX_Size() int {
	return m.Size()
}
func (m *MsgSaveEthaddressState) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSaveEthaddressState.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSaveEthaddressState proto.InternalMessageInfo

func (m *MsgSaveEthaddressState) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSaveEthaddressState) GetEthAddress() string {
	if m != nil {
		return m.EthAddress
	}
	return ""
}

func (m *MsgSaveEthaddressState) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *MsgSaveEthaddressState) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *MsgSaveEthaddressState) GetStoragePosition() uint64 {
	if m != nil {
		return m.StoragePosition
	}
	return 0
}

func (m *MsgSaveEthaddressState) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type MsgSaveEthaddressStateResponse struct {
	EthAddress string `protobuf:"bytes,1,opt,name=ethAddress,proto3" json:"ethAddress,omitempty"`
}

func (m *MsgSaveEthaddressStateResponse) Reset()         { *m = MsgSaveEthaddressStateResponse{} }
func (m *MsgSaveEthaddressStateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSaveEthaddressStateResponse) ProtoMessage()    {}
func (*MsgSaveEthaddressStateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3454f3c906df7237, []int{1}
}
func (m *MsgSaveEthaddressStateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSaveEthaddressStateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSaveEthaddressStateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSaveEthaddressStateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSaveEthaddressStateResponse.Merge(m, src)
}
func (m *MsgSaveEthaddressStateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSaveEthaddressStateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSaveEthaddressStateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSaveEthaddressStateResponse proto.InternalMessageInfo

func (m *MsgSaveEthaddressStateResponse) GetEthAddress() string {
	if m != nil {
		return m.EthAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgSaveEthaddressState)(nil), "statether.statether.MsgSaveEthaddressState")
	proto.RegisterType((*MsgSaveEthaddressStateResponse)(nil), "statether.statether.MsgSaveEthaddressStateResponse")
}

func init() { proto.RegisterFile("statether/statether/tx.proto", fileDescriptor_3454f3c906df7237) }

var fileDescriptor_3454f3c906df7237 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x3f, 0x4b, 0x03, 0x41,
	0x10, 0xc5, 0xb3, 0xe6, 0x8f, 0xba, 0x8d, 0xb0, 0x91, 0xb0, 0x88, 0x2c, 0x21, 0xd5, 0x81, 0x70,
	0x41, 0x83, 0xbd, 0x0a, 0x96, 0x41, 0xb9, 0x74, 0x76, 0x9b, 0xcb, 0x70, 0x77, 0x18, 0x77, 0xc2,
	0xce, 0x18, 0x63, 0xe3, 0x67, 0xf0, 0x23, 0x59, 0x5a, 0xa6, 0xb4, 0x94, 0xe4, 0x8b, 0x48, 0x2e,
	0x26, 0x86, 0xe3, 0x1a, 0xbb, 0x79, 0xbf, 0x37, 0xb0, 0x6f, 0xe7, 0xc9, 0x53, 0x62, 0xcb, 0xc0,
	0x29, 0xf8, 0xee, 0xdf, 0xc4, 0xb3, 0x70, 0xe2, 0x91, 0x51, 0x35, 0xb7, 0x2c, 0xdc, 0x4e, 0x9d,
	0x0f, 0x21, 0x5b, 0x7d, 0x4a, 0x06, 0x76, 0x0a, 0xb7, 0x9c, 0xda, 0xd1, 0xc8, 0x03, 0xd1, 0x60,
	0x65, 0x2b, 0x2d, 0xf7, 0x63, 0x0f, 0x96, 0xd1, 0x6b, 0xd1, 0x16, 0xc1, 0x61, 0xb4, 0x91, 0xca,
	0x48, 0x09, 0x9c, 0x5e, 0xaf, 0x97, 0xf5, 0x5e, 0x6e, 0xee, 0x10, 0x75, 0x2c, 0xeb, 0xc3, 0x31,
	0xc6, 0x8f, 0xba, 0xda, 0x16, 0x41, 0x2d, 0x5a, 0x8b, 0x15, 0x75, 0xe8, 0x62, 0xd0, 0xb5, 0x35,
	0xcd, 0x85, 0x0a, 0xe4, 0x11, 0x31, 0x7a, 0x9b, 0xc0, 0x3d, 0x52, 0xc6, 0x19, 0x3a, 0x5d, 0xcf,
	0xfd, 0x22, 0x56, 0x2d, 0xd9, 0xb0, 0x31, 0x67, 0x53, 0xd0, 0x8d, 0xb6, 0x08, 0x0e, 0xa2, 0x5f,
	0xd5, 0xb9, 0x92, 0xa6, 0xfc, 0x07, 0x11, 0xd0, 0x04, 0x1d, 0x41, 0x21, 0xaf, 0x28, 0xe6, 0xbd,
	0x78, 0x93, 0xd5, 0x3e, 0x25, 0xea, 0x45, 0x36, 0xcb, 0xee, 0x70, 0x16, 0x96, 0x1c, 0x2e, 0x2c,
	0x7f, 0xf2, 0xa4, 0xf7, 0x8f, 0xe5, 0x4d, 0xbe, 0x9b, 0xbb, 0xcf, 0x85, 0x11, 0xf3, 0x85, 0x11,
	0xdf, 0x0b, 0x23, 0xde, 0x97, 0xa6, 0x32, 0x5f, 0x9a, 0xca, 0xd7, 0xd2, 0x54, 0x1e, 0x2e, 0x93,
	0x8c, 0xd3, 0xe7, 0x61, 0x18, 0xe3, 0x53, 0xd7, 0xc1, 0x98, 0xd0, 0x11, 0xfb, 0x2e, 0xf6, 0xdc,
	0xf9, 0x4e, 0xc3, 0xb3, 0xdd, 0xb6, 0x5f, 0x27, 0x40, 0xc3, 0x46, 0xde, 0x78, 0xef, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x70, 0xe2, 0x57, 0x8d, 0x11, 0x02, 0x00, 0x00,
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
	SaveEthaddressState(ctx context.Context, in *MsgSaveEthaddressState, opts ...grpc.CallOption) (*MsgSaveEthaddressStateResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SaveEthaddressState(ctx context.Context, in *MsgSaveEthaddressState, opts ...grpc.CallOption) (*MsgSaveEthaddressStateResponse, error) {
	out := new(MsgSaveEthaddressStateResponse)
	err := c.cc.Invoke(ctx, "/statether.statether.Msg/SaveEthaddressState", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SaveEthaddressState(context.Context, *MsgSaveEthaddressState) (*MsgSaveEthaddressStateResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SaveEthaddressState(ctx context.Context, req *MsgSaveEthaddressState) (*MsgSaveEthaddressStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveEthaddressState not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SaveEthaddressState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSaveEthaddressState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SaveEthaddressState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/statether.statether.Msg/SaveEthaddressState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SaveEthaddressState(ctx, req.(*MsgSaveEthaddressState))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "statether.statether.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveEthaddressState",
			Handler:    _Msg_SaveEthaddressState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "statether/statether/tx.proto",
}

func (m *MsgSaveEthaddressState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSaveEthaddressState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSaveEthaddressState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.StoragePosition != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.StoragePosition))
		i--
		dAtA[i] = 0x28
	}
	if m.Nonce != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x20
	}
	if m.Block != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x18
	}
	if len(m.EthAddress) > 0 {
		i -= len(m.EthAddress)
		copy(dAtA[i:], m.EthAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EthAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSaveEthaddressStateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSaveEthaddressStateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSaveEthaddressStateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EthAddress) > 0 {
		i -= len(m.EthAddress)
		copy(dAtA[i:], m.EthAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.EthAddress)))
		i--
		dAtA[i] = 0xa
	}
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
func (m *MsgSaveEthaddressState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.EthAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovTx(uint64(m.Block))
	}
	if m.Nonce != 0 {
		n += 1 + sovTx(uint64(m.Nonce))
	}
	if m.StoragePosition != 0 {
		n += 1 + sovTx(uint64(m.StoragePosition))
	}
	if m.Active {
		n += 2
	}
	return n
}

func (m *MsgSaveEthaddressStateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EthAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSaveEthaddressState) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSaveEthaddressState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSaveEthaddressState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddress", wireType)
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
			m.EthAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Block |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoragePosition", wireType)
			}
			m.StoragePosition = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StoragePosition |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Active = bool(v != 0)
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
func (m *MsgSaveEthaddressStateResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSaveEthaddressStateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSaveEthaddressStateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddress", wireType)
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
			m.EthAddress = string(dAtA[iNdEx:postIndex])
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
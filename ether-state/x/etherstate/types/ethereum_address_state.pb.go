// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: etherstate/etherstate/ethereum_address_state.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type EthereumAddressState struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	State uint64 `protobuf:"varint,2,opt,name=state,proto3" json:"state,omitempty"`
	Block uint64 `protobuf:"varint,3,opt,name=block,proto3" json:"block,omitempty"`
	Nonce uint64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
}

func (m *EthereumAddressState) Reset()         { *m = EthereumAddressState{} }
func (m *EthereumAddressState) String() string { return proto.CompactTextString(m) }
func (*EthereumAddressState) ProtoMessage()    {}
func (*EthereumAddressState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7e09dc22c2637e56, []int{0}
}
func (m *EthereumAddressState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthereumAddressState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthereumAddressState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthereumAddressState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumAddressState.Merge(m, src)
}
func (m *EthereumAddressState) XXX_Size() int {
	return m.Size()
}
func (m *EthereumAddressState) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumAddressState.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumAddressState proto.InternalMessageInfo

func (m *EthereumAddressState) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *EthereumAddressState) GetState() uint64 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *EthereumAddressState) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *EthereumAddressState) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func init() {
	proto.RegisterType((*EthereumAddressState)(nil), "etherstate.etherstate.EthereumAddressState")
}

func init() {
	proto.RegisterFile("etherstate/etherstate/ethereum_address_state.proto", fileDescriptor_7e09dc22c2637e56)
}

var fileDescriptor_7e09dc22c2637e56 = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x4a, 0x2d, 0xc9, 0x48,
	0x2d, 0x2a, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x47, 0x67, 0xa6, 0x96, 0xe6, 0xc6, 0x27, 0xa6, 0xa4,
	0x14, 0xa5, 0x16, 0x17, 0xc7, 0x83, 0x85, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x44, 0x11,
	0x0a, 0xf5, 0x10, 0x4c, 0xa5, 0x3c, 0x2e, 0x11, 0x57, 0xa8, 0x36, 0x47, 0x88, 0xae, 0x60, 0x90,
	0xb8, 0x90, 0x08, 0x17, 0x6b, 0x66, 0x5e, 0x4a, 0x6a, 0x85, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x84, 0x03, 0x12, 0x05, 0x6b, 0x93, 0x60, 0x52, 0x60, 0xd4, 0x60, 0x09, 0x82, 0x70, 0x40,
	0xa2, 0x49, 0x39, 0xf9, 0xc9, 0xd9, 0x12, 0xcc, 0x10, 0x51, 0x30, 0x07, 0x24, 0x9a, 0x97, 0x9f,
	0x97, 0x9c, 0x2a, 0xc1, 0x02, 0x11, 0x05, 0x73, 0x9c, 0x82, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0,
	0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8,
	0xf1, 0x58, 0x8e, 0x21, 0xca, 0x22, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57,
	0x3f, 0x2f, 0x35, 0xa7, 0x38, 0x3f, 0xaf, 0xb8, 0xa4, 0x48, 0x3f, 0xdf, 0x38, 0xcf, 0x10, 0xe2,
	0x31, 0x5d, 0x88, 0x27, 0x2b, 0x90, 0x7d, 0x5c, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6,
	0xa1, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x60, 0x17, 0x20, 0x29, 0x17, 0x01, 0x00, 0x00,
}

func (m *EthereumAddressState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthereumAddressState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthereumAddressState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Nonce != 0 {
		i = encodeVarintEthereumAddressState(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x20
	}
	if m.Block != 0 {
		i = encodeVarintEthereumAddressState(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x18
	}
	if m.State != 0 {
		i = encodeVarintEthereumAddressState(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintEthereumAddressState(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEthereumAddressState(dAtA []byte, offset int, v uint64) int {
	offset -= sovEthereumAddressState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthereumAddressState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovEthereumAddressState(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovEthereumAddressState(uint64(m.State))
	}
	if m.Block != 0 {
		n += 1 + sovEthereumAddressState(uint64(m.Block))
	}
	if m.Nonce != 0 {
		n += 1 + sovEthereumAddressState(uint64(m.Nonce))
	}
	return n
}

func sovEthereumAddressState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEthereumAddressState(x uint64) (n int) {
	return sovEthereumAddressState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthereumAddressState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthereumAddressState
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
			return fmt.Errorf("proto: EthereumAddressState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthereumAddressState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumAddressState
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
				return ErrInvalidLengthEthereumAddressState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthereumAddressState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumAddressState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthereumAddressState
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
					return ErrIntOverflowEthereumAddressState
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
		default:
			iNdEx = preIndex
			skippy, err := skipEthereumAddressState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEthereumAddressState
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
func skipEthereumAddressState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEthereumAddressState
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
					return 0, ErrIntOverflowEthereumAddressState
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
					return 0, ErrIntOverflowEthereumAddressState
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
				return 0, ErrInvalidLengthEthereumAddressState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEthereumAddressState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEthereumAddressState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEthereumAddressState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEthereumAddressState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEthereumAddressState = fmt.Errorf("proto: unexpected end of group")
)

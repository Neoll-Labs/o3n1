// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storepositionether/storepositionether/ethaddress_storage_position.proto

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

type EthaddressStoragePosition struct {
	EthAddress      string `protobuf:"bytes,1,opt,name=ethAddress,proto3" json:"ethAddress,omitempty"`
	Block           uint64 `protobuf:"varint,2,opt,name=block,proto3" json:"block,omitempty"`
	Nonce           uint64 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	StoragePosition uint64 `protobuf:"varint,4,opt,name=storagePosition,proto3" json:"storagePosition,omitempty"`
	Active          bool   `protobuf:"varint,5,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *EthaddressStoragePosition) Reset()         { *m = EthaddressStoragePosition{} }
func (m *EthaddressStoragePosition) String() string { return proto.CompactTextString(m) }
func (*EthaddressStoragePosition) ProtoMessage()    {}
func (*EthaddressStoragePosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_dff876832ea6df86, []int{0}
}
func (m *EthaddressStoragePosition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EthaddressStoragePosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EthaddressStoragePosition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EthaddressStoragePosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthaddressStoragePosition.Merge(m, src)
}
func (m *EthaddressStoragePosition) XXX_Size() int {
	return m.Size()
}
func (m *EthaddressStoragePosition) XXX_DiscardUnknown() {
	xxx_messageInfo_EthaddressStoragePosition.DiscardUnknown(m)
}

var xxx_messageInfo_EthaddressStoragePosition proto.InternalMessageInfo

func (m *EthaddressStoragePosition) GetEthAddress() string {
	if m != nil {
		return m.EthAddress
	}
	return ""
}

func (m *EthaddressStoragePosition) GetBlock() uint64 {
	if m != nil {
		return m.Block
	}
	return 0
}

func (m *EthaddressStoragePosition) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *EthaddressStoragePosition) GetStoragePosition() uint64 {
	if m != nil {
		return m.StoragePosition
	}
	return 0
}

func (m *EthaddressStoragePosition) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

func init() {
	proto.RegisterType((*EthaddressStoragePosition)(nil), "storepositionether.storepositionether.EthaddressStoragePosition")
}

func init() {
	proto.RegisterFile("storepositionether/storepositionether/ethaddress_storage_position.proto", fileDescriptor_dff876832ea6df86)
}

var fileDescriptor_dff876832ea6df86 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x2f, 0x2e, 0xc9, 0x2f,
	0x4a, 0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0xcc, 0xcf, 0x4b, 0x2d, 0xc9, 0x48, 0x2d, 0xd2, 0xc7,
	0x22, 0x94, 0x5a, 0x92, 0x91, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x1c, 0x0f, 0x92, 0x4d, 0x4c,
	0x4f, 0x8d, 0x87, 0x29, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xc5, 0xd4, 0xa5, 0x87,
	0x29, 0xa4, 0xb4, 0x92, 0x91, 0x4b, 0xd2, 0x15, 0x6e, 0x58, 0x30, 0xc4, 0xac, 0x00, 0xa8, 0x12,
	0x21, 0x39, 0x2e, 0xae, 0xd4, 0x92, 0x0c, 0x47, 0x88, 0xa4, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x92, 0x88, 0x90, 0x08, 0x17, 0x6b, 0x52, 0x4e, 0x7e, 0x72, 0xb6, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0x4b, 0x10, 0x84, 0x03, 0x12, 0xcd, 0xcb, 0xcf, 0x4b, 0x4e, 0x95, 0x60, 0x86, 0x88, 0x82,
	0x39, 0x42, 0x1a, 0x5c, 0xfc, 0xc5, 0xa8, 0xc6, 0x4b, 0xb0, 0x80, 0xe5, 0xd1, 0x85, 0x85, 0xc4,
	0xb8, 0xd8, 0x12, 0x93, 0x4b, 0x32, 0xcb, 0x52, 0x25, 0x58, 0x15, 0x18, 0x35, 0x38, 0x82, 0xa0,
	0x3c, 0x27, 0xf7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0xc5, 0x12,
	0x44, 0x15, 0xd8, 0xc2, 0xad, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0x1c, 0x44, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x24, 0x01, 0xf6, 0x79, 0x6d, 0x01, 0x00, 0x00,
}

func (m *EthaddressStoragePosition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EthaddressStoragePosition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EthaddressStoragePosition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0x28
	}
	if m.StoragePosition != 0 {
		i = encodeVarintEthaddressStoragePosition(dAtA, i, uint64(m.StoragePosition))
		i--
		dAtA[i] = 0x20
	}
	if m.Nonce != 0 {
		i = encodeVarintEthaddressStoragePosition(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x18
	}
	if m.Block != 0 {
		i = encodeVarintEthaddressStoragePosition(dAtA, i, uint64(m.Block))
		i--
		dAtA[i] = 0x10
	}
	if len(m.EthAddress) > 0 {
		i -= len(m.EthAddress)
		copy(dAtA[i:], m.EthAddress)
		i = encodeVarintEthaddressStoragePosition(dAtA, i, uint64(len(m.EthAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEthaddressStoragePosition(dAtA []byte, offset int, v uint64) int {
	offset -= sovEthaddressStoragePosition(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EthaddressStoragePosition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EthAddress)
	if l > 0 {
		n += 1 + l + sovEthaddressStoragePosition(uint64(l))
	}
	if m.Block != 0 {
		n += 1 + sovEthaddressStoragePosition(uint64(m.Block))
	}
	if m.Nonce != 0 {
		n += 1 + sovEthaddressStoragePosition(uint64(m.Nonce))
	}
	if m.StoragePosition != 0 {
		n += 1 + sovEthaddressStoragePosition(uint64(m.StoragePosition))
	}
	if m.Active {
		n += 2
	}
	return n
}

func sovEthaddressStoragePosition(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEthaddressStoragePosition(x uint64) (n int) {
	return sovEthaddressStoragePosition(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EthaddressStoragePosition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEthaddressStoragePosition
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
			return fmt.Errorf("proto: EthaddressStoragePosition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EthaddressStoragePosition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EthAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthaddressStoragePosition
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
				return ErrInvalidLengthEthaddressStoragePosition
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEthaddressStoragePosition
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EthAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			m.Block = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthaddressStoragePosition
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
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthaddressStoragePosition
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoragePosition", wireType)
			}
			m.StoragePosition = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthaddressStoragePosition
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
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEthaddressStoragePosition
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
			skippy, err := skipEthaddressStoragePosition(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEthaddressStoragePosition
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
func skipEthaddressStoragePosition(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEthaddressStoragePosition
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
					return 0, ErrIntOverflowEthaddressStoragePosition
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
					return 0, ErrIntOverflowEthaddressStoragePosition
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
				return 0, ErrInvalidLengthEthaddressStoragePosition
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEthaddressStoragePosition
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEthaddressStoragePosition
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEthaddressStoragePosition        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEthaddressStoragePosition          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEthaddressStoragePosition = fmt.Errorf("proto: unexpected end of group")
)
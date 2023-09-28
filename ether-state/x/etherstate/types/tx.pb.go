// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: etherstate/etherstate/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("etherstate/etherstate/tx.proto", fileDescriptor_1a2220bdd040cbcf) }

var fileDescriptor_1a2220bdd040cbcf = []byte{
	// 107 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0x2d, 0xc9, 0x48,
	0x2d, 0x2a, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x47, 0x62, 0x96, 0x54, 0xe8, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0x89, 0x22, 0x04, 0xf5, 0x10, 0x4c, 0x23, 0x56, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0x27,
	0x8b, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x82, 0x98, 0xab, 0x0b, 0x31,
	0xad, 0x02, 0xc5, 0xe8, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0xf1, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x64, 0xae, 0xbf, 0x03, 0x80, 0x00, 0x00, 0x00,
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
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "etherstate.etherstate.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "etherstate/etherstate/tx.proto",
}

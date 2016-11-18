// Code generated by protoc-gen-go.
// source: peer/fabric_block.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Block contains a list of transactions and the crypto hash of previous block
type Block2 struct {
	PreviousBlockHash []byte `protobuf:"bytes,1,opt,name=PreviousBlockHash,proto3" json:"PreviousBlockHash,omitempty"`
	// transactions are stored in serialized form so that the concenters can avoid marshaling of transactions
	Transactions [][]byte `protobuf:"bytes,2,rep,name=Transactions,proto3" json:"Transactions,omitempty"`
}

func (m *Block2) Reset()                    { *m = Block2{} }
func (m *Block2) String() string            { return proto.CompactTextString(m) }
func (*Block2) ProtoMessage()               {}
func (*Block2) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func init() {
	proto.RegisterType((*Block2)(nil), "protos.Block2")
}

func init() { proto.RegisterFile("peer/fabric_block.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 150 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x48, 0x4d, 0x2d,
	0xd2, 0x4f, 0x4b, 0x4c, 0x2a, 0xca, 0x4c, 0x8e, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x53, 0xc5, 0x4a, 0x51, 0x5c, 0x6c, 0x4e, 0x20, 0x61, 0x23,
	0x21, 0x1d, 0x2e, 0xc1, 0x80, 0xa2, 0xd4, 0xb2, 0xcc, 0xfc, 0xd2, 0x62, 0xb0, 0x88, 0x47, 0x62,
	0x71, 0x86, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x4f, 0x10, 0xa6, 0x84, 0x90, 0x12, 0x17, 0x4f, 0x48,
	0x51, 0x62, 0x5e, 0x71, 0x62, 0x72, 0x49, 0x66, 0x7e, 0x5e, 0xb1, 0x04, 0x93, 0x02, 0xb3, 0x06,
	0x4f, 0x10, 0x8a, 0x98, 0x93, 0x76, 0x94, 0x66, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72,
	0x7e, 0xae, 0x7e, 0x46, 0x65, 0x41, 0x6a, 0x51, 0x4e, 0x6a, 0x4a, 0x3a, 0xdc, 0x41, 0xfa, 0x10,
	0x37, 0xe8, 0x83, 0xdc, 0x98, 0x04, 0x71, 0x90, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x40, 0x8e,
	0x0d, 0x09, 0xb2, 0x00, 0x00, 0x00,
}

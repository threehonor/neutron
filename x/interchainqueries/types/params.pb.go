// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/interchainqueries/params.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// Params defines the parameters for the module.
type Params struct {
	// Defines amount of blocks required before query becomes available for
	// removal by anybody
	QuerySubmitTimeout uint64 `protobuf:"varint,1,opt,name=query_submit_timeout,json=querySubmitTimeout,proto3" json:"query_submit_timeout,omitempty"`
	// Amount of coins deposited for the query.
	QueryDeposit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=query_deposit,json=queryDeposit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"query_deposit"`
	// Amount of tx hashes to be removed during a single EndBlock. Can vary to
	// balance between network cleaning speed and EndBlock duration. A zero value
	// means no limit.
	TxQueryRemovalLimit uint64 `protobuf:"varint,3,opt,name=tx_query_removal_limit,json=txQueryRemovalLimit,proto3" json:"tx_query_removal_limit,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_752a5f3346da64b1, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetQuerySubmitTimeout() uint64 {
	if m != nil {
		return m.QuerySubmitTimeout
	}
	return 0
}

func (m *Params) GetQueryDeposit() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.QueryDeposit
	}
	return nil
}

func (m *Params) GetTxQueryRemovalLimit() uint64 {
	if m != nil {
		return m.TxQueryRemovalLimit
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "neutron.interchainqueries.Params")
}

func init() {
	proto.RegisterFile("neutron/interchainqueries/params.proto", fileDescriptor_752a5f3346da64b1)
}

var fileDescriptor_752a5f3346da64b1 = []byte{
	// 331 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x10, 0xc7, 0x93, 0xaf, 0x55, 0x87, 0x7c, 0xb0, 0x84, 0x0a, 0xb5, 0x1d, 0xdc, 0x8a, 0x01, 0x75,
	0xa9, 0xdd, 0x52, 0x26, 0xd8, 0x0a, 0x23, 0x03, 0x14, 0x58, 0x58, 0xa2, 0x24, 0xb5, 0x52, 0x8b,
	0x3a, 0x17, 0xec, 0x4b, 0xd5, 0xbe, 0x05, 0x23, 0x23, 0x33, 0x4f, 0xd2, 0xb1, 0x23, 0x13, 0xa0,
	0x76, 0xe0, 0x35, 0x50, 0xec, 0x20, 0x21, 0xc1, 0xe4, 0x93, 0x7f, 0x77, 0xfe, 0xff, 0x74, 0xf6,
	0x0e, 0x53, 0x9e, 0xa3, 0x82, 0x94, 0x89, 0x14, 0xb9, 0x8a, 0xa7, 0xa1, 0x48, 0x1f, 0x72, 0xae,
	0x04, 0xd7, 0x2c, 0x0b, 0x55, 0x28, 0x35, 0xcd, 0x14, 0x20, 0xf8, 0xcd, 0xb2, 0x8f, 0xfe, 0xea,
	0x6b, 0x91, 0x18, 0xb4, 0x04, 0xcd, 0xa2, 0x50, 0x73, 0x36, 0x1f, 0x44, 0x1c, 0xc3, 0x01, 0x8b,
	0x41, 0xa4, 0x76, 0xb4, 0x55, 0x4f, 0x20, 0x01, 0x53, 0xb2, 0xa2, 0xb2, 0xb7, 0x07, 0x9f, 0xae,
	0x57, 0xbb, 0x34, 0x09, 0x7e, 0xdf, 0xab, 0x17, 0x6f, 0x2d, 0x03, 0x9d, 0x47, 0x52, 0x60, 0x80,
	0x42, 0x72, 0xc8, 0xb1, 0xe1, 0x76, 0xdc, 0x6e, 0x75, 0xec, 0x1b, 0x76, 0x6d, 0xd0, 0x8d, 0x25,
	0x7e, 0xe6, 0xed, 0xda, 0x89, 0x09, 0xcf, 0x40, 0x0b, 0x6c, 0xfc, 0xeb, 0x54, 0xba, 0xff, 0x8f,
	0x9a, 0xd4, 0xaa, 0xd0, 0x42, 0x85, 0x96, 0x2a, 0xf4, 0x0c, 0x44, 0x3a, 0xea, 0xaf, 0xde, 0xda,
	0xce, 0xcb, 0x7b, 0xbb, 0x9b, 0x08, 0x9c, 0xe6, 0x11, 0x8d, 0x41, 0xb2, 0xd2, 0xdb, 0x1e, 0x3d,
	0x3d, 0xb9, 0x67, 0xb8, 0xcc, 0xb8, 0x36, 0x03, 0x7a, 0xbc, 0x63, 0x12, 0xce, 0x6d, 0x80, 0x3f,
	0xf4, 0xf6, 0x71, 0x11, 0xd8, 0x50, 0xc5, 0x25, 0xcc, 0xc3, 0x59, 0x30, 0x13, 0x52, 0x60, 0xa3,
	0x62, 0x2c, 0xf7, 0x70, 0x71, 0x55, 0xc0, 0xb1, 0x65, 0x17, 0x05, 0x3a, 0xa9, 0x3e, 0x3d, 0xb7,
	0x9d, 0xd1, 0xed, 0x6a, 0x43, 0xdc, 0xf5, 0x86, 0xb8, 0x1f, 0x1b, 0xe2, 0x3e, 0x6e, 0x89, 0xb3,
	0xde, 0x12, 0xe7, 0x75, 0x4b, 0x9c, 0xbb, 0xd3, 0x1f, 0x32, 0xe5, 0x7e, 0x7b, 0xa0, 0x92, 0xef,
	0x9a, 0xcd, 0x8f, 0xd9, 0xe2, 0x8f, 0x8f, 0x31, 0x96, 0x51, 0xcd, 0xec, 0x71, 0xf8, 0x15, 0x00,
	0x00, 0xff, 0xff, 0xc7, 0x61, 0x00, 0x43, 0xc2, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TxQueryRemovalLimit != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.TxQueryRemovalLimit))
		i--
		dAtA[i] = 0x18
	}
	if len(m.QueryDeposit) > 0 {
		for iNdEx := len(m.QueryDeposit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.QueryDeposit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintParams(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.QuerySubmitTimeout != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.QuerySubmitTimeout))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.QuerySubmitTimeout != 0 {
		n += 1 + sovParams(uint64(m.QuerySubmitTimeout))
	}
	if len(m.QueryDeposit) > 0 {
		for _, e := range m.QueryDeposit {
			l = e.Size()
			n += 1 + l + sovParams(uint64(l))
		}
	}
	if m.TxQueryRemovalLimit != 0 {
		n += 1 + sovParams(uint64(m.TxQueryRemovalLimit))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field QuerySubmitTimeout", wireType)
			}
			m.QuerySubmitTimeout = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.QuerySubmitTimeout |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field QueryDeposit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.QueryDeposit = append(m.QueryDeposit, types.Coin{})
			if err := m.QueryDeposit[len(m.QueryDeposit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxQueryRemovalLimit", wireType)
			}
			m.TxQueryRemovalLimit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TxQueryRemovalLimit |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)

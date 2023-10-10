// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: neutron/incentives/gauge.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	types1 "github.com/neutron-org/neutron/x/dex/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Gauge is an object that describes an LP incentivization plan and its state.
type Gauge struct {
	// id is the unique ID of a Gauge
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// There are two kinds of gauges: perpetual and non-perpetual. Perpetual
	// gauges describe an incentivization program for which the token rewards
	// distributed on any given day must be added to the gauge prior to that day's
	// distribution using an AddToGauge message.  When distribute is called on a
	// perpetual gauge, all of the remaining rewards in the gauge are distributed.
	// Because of this, all perpetual gauges must have `num_epochs_paid_over` set
	// to 1.  A non-perpetual gauge by contrast distributes its rewards over a
	// schedule as determined by `num_epochs_paid_over`. If a non-perpetual gauge
	// is created with coins=[100atom] and num_epochs_paid_over=10, this means
	// that for 10 days (10 epochs) the gauge will distribute 10atom each day to
	// the staked LP positions qualifying for the gauge.
	IsPerpetual bool `protobuf:"varint,2,opt,name=is_perpetual,json=isPerpetual,proto3" json:"is_perpetual,omitempty"`
	// distribute_to describes a set of staked LP positions that should be
	// distributed to from this gauge.
	DistributeTo QueryCondition `protobuf:"bytes,3,opt,name=distribute_to,json=distributeTo,proto3" json:"distribute_to"`
	// coins describes the total amount of coins that have been added to this
	// gauge for distribution.
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,4,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
	// start_time describes when this gauge should begin distributing rewards.
	// This allows gauge creators to schedule gauges into the future, in the event
	// that an earlier gauge is expected to expire.
	StartTime time.Time `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	// num_epochs_paid_over is the number of total epochs (days) the rewards in
	// this gauge will be distributed over.
	NumEpochsPaidOver uint64 `protobuf:"varint,6,opt,name=num_epochs_paid_over,json=numEpochsPaidOver,proto3" json:"num_epochs_paid_over,omitempty"`
	// filled_epochs describes the number of epochs distribution have been completed
	// already
	FilledEpochs uint64 `protobuf:"varint,7,opt,name=filled_epochs,json=filledEpochs,proto3" json:"filled_epochs,omitempty"`
	// distributed_coins describes coins that have been distributed already from
	// this gauge.
	DistributedCoins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,8,rep,name=distributed_coins,json=distributedCoins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"distributed_coins"`
	// pricing_tick is necessary for fairly distributing rewards over a range of
	// ticks.  Without pricing_tick, we might naively distribute rewards in
	// proportion to the number of deposit shares staked within the gauge's
	// qualifying tick range.
	//
	// For example, a gauge with a distribute_to tick range of [-10, 10] would
	// distribute to staked LP tokens where both tick-fee and tick+fee are within
	// [-10, 10]. Let's say for pair "tokenA<>tokenB", the current trading tick is
	// 0. If Alice were to LP (10tokenA, 0tokenB) @ tick -8, fee 2, this would
	// mean Alice would be issued 10 shares (10 + 0 * 1.0001^-8), since shares are
	// in terms of token0. Let's further assume Bob LPs (0tokenA, 10tokenB) @ tick
	// 8, fee 2, such that Bob is issued 10.008 shares (0 + 10 * 1.0001^8). Under
	// this naive approach, if Alice and Bob were to stake their shares, Bob would
	// receive more in rewards, purely on the basis of the relative locations of
	// their liquidity.
	//
	// This disparity originates in the fact that LP deposit denominations are not
	// fungible across ticks. To avoid this, we can use a single price throughout
	// the gauge's tick range for relating the relative value of token0 and
	// token1, as specified by pricing_tick.
	//
	// Let's run through the earier example using the more sophisticated approach,
	// where the gauge has pricing_tick set to 0. For the purpose of calculating
	// reward distribution weight, Alice's shares are worth 10 + 0 * 1.0001^0 = 10
	// and Bob's shares are worth 0 + 10 * 1.0001^0 = 10. With the distribution
	// weight of both shares set according to a gauge-specific tick, we do not
	// distribute more or less rewards according to the relative location of
	// liquidity within the gauge's tick range, freeing users to place liquidity
	// whereever they deem most profitable in the gauge's range and still equally
	// qualify for rewards.
	PricingTick int64 `protobuf:"varint,9,opt,name=pricing_tick,json=pricingTick,proto3" json:"pricing_tick,omitempty"`
}

func (m *Gauge) Reset()         { *m = Gauge{} }
func (m *Gauge) String() string { return proto.CompactTextString(m) }
func (*Gauge) ProtoMessage()    {}
func (*Gauge) Descriptor() ([]byte, []int) {
	return fileDescriptor_2467fab98b594cb6, []int{0}
}
func (m *Gauge) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Gauge) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Gauge.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Gauge) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gauge.Merge(m, src)
}
func (m *Gauge) XXX_Size() int {
	return m.Size()
}
func (m *Gauge) XXX_DiscardUnknown() {
	xxx_messageInfo_Gauge.DiscardUnknown(m)
}

var xxx_messageInfo_Gauge proto.InternalMessageInfo

func (m *Gauge) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Gauge) GetIsPerpetual() bool {
	if m != nil {
		return m.IsPerpetual
	}
	return false
}

func (m *Gauge) GetDistributeTo() QueryCondition {
	if m != nil {
		return m.DistributeTo
	}
	return QueryCondition{}
}

func (m *Gauge) GetCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Coins
	}
	return nil
}

func (m *Gauge) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *Gauge) GetNumEpochsPaidOver() uint64 {
	if m != nil {
		return m.NumEpochsPaidOver
	}
	return 0
}

func (m *Gauge) GetFilledEpochs() uint64 {
	if m != nil {
		return m.FilledEpochs
	}
	return 0
}

func (m *Gauge) GetDistributedCoins() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DistributedCoins
	}
	return nil
}

func (m *Gauge) GetPricingTick() int64 {
	if m != nil {
		return m.PricingTick
	}
	return 0
}

// QueryCondition describes a set of staked LP positions that a gauge is
// configured to distribute to. LP tokens qualifying for a given QueryCondition
// must have both tick-fee and tick+fee fall within the range [startTick, endTick],
// such that all of the tradable liquidity for the pool is within that range.
type QueryCondition struct {
	// pairID is the token pair which should be distributed to.
	PairID *types1.PairID `protobuf:"bytes,1,opt,name=pairID,proto3" json:"pairID,omitempty"`
	// start_tick is the inclusive lower bound on the location of LP tokens that
	// qualify for a gauge's distribution.
	StartTick int64 `protobuf:"varint,2,opt,name=startTick,proto3" json:"startTick,omitempty"`
	// end_tick is the inclusive upper bound on the location of LP tokens that
	// qualify for a gauge's distribution.
	EndTick int64 `protobuf:"varint,3,opt,name=endTick,proto3" json:"endTick,omitempty"`
}

func (m *QueryCondition) Reset()         { *m = QueryCondition{} }
func (m *QueryCondition) String() string { return proto.CompactTextString(m) }
func (*QueryCondition) ProtoMessage()    {}
func (*QueryCondition) Descriptor() ([]byte, []int) {
	return fileDescriptor_2467fab98b594cb6, []int{1}
}
func (m *QueryCondition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryCondition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryCondition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryCondition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCondition.Merge(m, src)
}
func (m *QueryCondition) XXX_Size() int {
	return m.Size()
}
func (m *QueryCondition) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCondition.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCondition proto.InternalMessageInfo

func (m *QueryCondition) GetPairID() *types1.PairID {
	if m != nil {
		return m.PairID
	}
	return nil
}

func (m *QueryCondition) GetStartTick() int64 {
	if m != nil {
		return m.StartTick
	}
	return 0
}

func (m *QueryCondition) GetEndTick() int64 {
	if m != nil {
		return m.EndTick
	}
	return 0
}

func init() {
	proto.RegisterType((*Gauge)(nil), "neutron.incentives.Gauge")
	proto.RegisterType((*QueryCondition)(nil), "neutron.incentives.QueryCondition")
}

func init() { proto.RegisterFile("neutron/incentives/gauge.proto", fileDescriptor_2467fab98b594cb6) }

var fileDescriptor_2467fab98b594cb6 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0x8d, 0x9b, 0x9f, 0xb6, 0x93, 0xb4, 0xfa, 0x32, 0x5f, 0x17, 0x6e, 0x04, 0x4e, 0x08, 0x1b,
	0x4b, 0xa8, 0x33, 0x24, 0xec, 0x58, 0xa6, 0x20, 0x84, 0x10, 0x22, 0x58, 0x59, 0x20, 0x36, 0xd6,
	0xd8, 0x33, 0x75, 0x47, 0x89, 0x3d, 0xd6, 0xcc, 0x38, 0x4a, 0xde, 0xa2, 0xcf, 0xc1, 0x1b, 0xf0,
	0x06, 0x5d, 0x76, 0xc9, 0xaa, 0x45, 0xc9, 0x1b, 0xf0, 0x04, 0xc8, 0x63, 0x9b, 0x04, 0xd8, 0xb2,
	0xf2, 0xcc, 0x39, 0xf7, 0xce, 0xbd, 0xe7, 0xdc, 0x6b, 0xe0, 0x24, 0x2c, 0xd3, 0x52, 0x24, 0x98,
	0x27, 0x21, 0x4b, 0x34, 0x5f, 0x32, 0x85, 0x23, 0x92, 0x45, 0x0c, 0xa5, 0x52, 0x68, 0x01, 0x61,
	0xc9, 0xa3, 0x1d, 0xdf, 0x3b, 0x8b, 0x44, 0x24, 0x0c, 0x8d, 0xf3, 0x53, 0x11, 0xd9, 0x73, 0x22,
	0x21, 0xa2, 0x05, 0xc3, 0xe6, 0x16, 0x64, 0x57, 0x98, 0x66, 0x92, 0x68, 0x2e, 0x92, 0x92, 0xef,
	0xff, 0xc9, 0x6b, 0x1e, 0x33, 0xa5, 0x49, 0x9c, 0x56, 0x0f, 0x84, 0x42, 0xc5, 0x42, 0xe1, 0x80,
	0x28, 0x86, 0x97, 0xa3, 0x80, 0x69, 0x32, 0xc2, 0xa1, 0xe0, 0xd5, 0x03, 0xe7, 0x55, 0xab, 0x94,
	0xad, 0x70, 0x4a, 0xb8, 0xf4, 0x39, 0x2d, 0xa8, 0xe1, 0xd7, 0x06, 0x68, 0xbe, 0xc9, 0xbb, 0x86,
	0xa7, 0xe0, 0x80, 0x53, 0xdb, 0x1a, 0x58, 0x6e, 0xc3, 0x3b, 0xe0, 0x14, 0x3e, 0x01, 0x1d, 0xae,
	0xfc, 0x94, 0xc9, 0x94, 0xe9, 0x8c, 0x2c, 0xec, 0x83, 0x81, 0xe5, 0x1e, 0x79, 0x6d, 0xae, 0xa6,
	0x15, 0x04, 0xdf, 0x83, 0x13, 0xca, 0x95, 0x96, 0x3c, 0xc8, 0x34, 0xf3, 0xb5, 0xb0, 0xeb, 0x03,
	0xcb, 0x6d, 0x8f, 0x87, 0xe8, 0x6f, 0xe9, 0xe8, 0x63, 0xc6, 0xe4, 0xfa, 0x52, 0x24, 0x94, 0xe7,
	0xca, 0x26, 0x8d, 0xdb, 0xfb, 0x7e, 0xcd, 0xeb, 0xec, 0xd2, 0x67, 0x02, 0x12, 0xd0, 0xcc, 0x9b,
	0x56, 0x76, 0x63, 0x50, 0x77, 0xdb, 0xe3, 0x73, 0x54, 0xc8, 0x42, 0xb9, 0x2c, 0x54, 0xca, 0x42,
	0x97, 0x82, 0x27, 0x93, 0xe7, 0x79, 0xf6, 0x97, 0x87, 0xbe, 0x1b, 0x71, 0x7d, 0x9d, 0x05, 0x28,
	0x14, 0x31, 0x2e, 0x3d, 0x28, 0x3e, 0x17, 0x8a, 0xce, 0xb1, 0x5e, 0xa7, 0x4c, 0x99, 0x04, 0xe5,
	0x15, 0x2f, 0xc3, 0x4f, 0x00, 0x28, 0x4d, 0xa4, 0xf6, 0x73, 0x0b, 0xed, 0xa6, 0x69, 0xb7, 0x87,
	0x0a, 0x7f, 0x51, 0xe5, 0x2f, 0x9a, 0x55, 0xfe, 0x4e, 0x1e, 0xe7, 0x85, 0x7e, 0xdc, 0xf7, 0xbb,
	0x6b, 0x12, 0x2f, 0x5e, 0x0e, 0x77, 0xb9, 0xc3, 0x9b, 0x87, 0xbe, 0xe5, 0x1d, 0x1b, 0x20, 0x0f,
	0x87, 0x18, 0x9c, 0x25, 0x59, 0xec, 0xb3, 0x54, 0x84, 0xd7, 0xca, 0x4f, 0x09, 0xa7, 0xbe, 0x58,
	0x32, 0x69, 0xb7, 0x8c, 0xa1, 0xdd, 0x24, 0x8b, 0x5f, 0x1b, 0x6a, 0x4a, 0x38, 0xfd, 0xb0, 0x64,
	0x12, 0x3e, 0x05, 0x27, 0x57, 0x7c, 0xb1, 0x60, 0xb4, 0xcc, 0xb1, 0x0f, 0x4d, 0x64, 0xa7, 0x00,
	0x8b, 0x60, 0xb8, 0x02, 0xdd, 0x9d, 0x45, 0xd4, 0x2f, 0xec, 0x39, 0xfa, 0xf7, 0xf6, 0xfc, 0xb7,
	0x57, 0xc5, 0x20, 0xf9, 0xf8, 0x53, 0xc9, 0x43, 0x9e, 0x44, 0xbe, 0xe6, 0xe1, 0xdc, 0x3e, 0x1e,
	0x58, 0x6e, 0xdd, 0x6b, 0x97, 0xd8, 0x8c, 0x87, 0xf3, 0x61, 0x06, 0x4e, 0x7f, 0x9f, 0x2a, 0x7c,
	0x06, 0x5a, 0xf9, 0x7a, 0xbd, 0x7d, 0x65, 0xf6, 0xa8, 0x3d, 0xfe, 0xff, 0xd7, 0x26, 0x50, 0xb6,
	0x42, 0x53, 0x43, 0x79, 0x65, 0x08, 0x7c, 0x04, 0x2a, 0xfb, 0xc2, 0xb9, 0xd9, 0xae, 0xba, 0xb7,
	0x03, 0xa0, 0x0d, 0x0e, 0x59, 0x42, 0x0d, 0x57, 0x37, 0x5c, 0x75, 0x9d, 0xbc, 0xbb, 0xdd, 0x38,
	0xd6, 0xdd, 0xc6, 0xb1, 0xbe, 0x6f, 0x1c, 0xeb, 0x66, 0xeb, 0xd4, 0xee, 0xb6, 0x4e, 0xed, 0xdb,
	0xd6, 0xa9, 0x7d, 0x1e, 0xed, 0xe9, 0x2d, 0x0b, 0x5f, 0x08, 0x19, 0x55, 0x67, 0xbc, 0xda, 0xff,
	0x57, 0x8d, 0xfc, 0xa0, 0x65, 0x86, 0xfe, 0xe2, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf6, 0xde,
	0x74, 0xfa, 0xce, 0x03, 0x00, 0x00,
}

func (m *Gauge) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Gauge) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Gauge) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PricingTick != 0 {
		i = encodeVarintGauge(dAtA, i, uint64(m.PricingTick))
		i--
		dAtA[i] = 0x48
	}
	if len(m.DistributedCoins) > 0 {
		for iNdEx := len(m.DistributedCoins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DistributedCoins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGauge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.FilledEpochs != 0 {
		i = encodeVarintGauge(dAtA, i, uint64(m.FilledEpochs))
		i--
		dAtA[i] = 0x38
	}
	if m.NumEpochsPaidOver != 0 {
		i = encodeVarintGauge(dAtA, i, uint64(m.NumEpochsPaidOver))
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGauge(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGauge(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	{
		size, err := m.DistributeTo.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGauge(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.IsPerpetual {
		i--
		if m.IsPerpetual {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintGauge(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryCondition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryCondition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryCondition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EndTick != 0 {
		i = encodeVarintGauge(dAtA, i, uint64(m.EndTick))
		i--
		dAtA[i] = 0x18
	}
	if m.StartTick != 0 {
		i = encodeVarintGauge(dAtA, i, uint64(m.StartTick))
		i--
		dAtA[i] = 0x10
	}
	if m.PairID != nil {
		{
			size, err := m.PairID.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGauge(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGauge(dAtA []byte, offset int, v uint64) int {
	offset -= sovGauge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Gauge) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovGauge(uint64(m.Id))
	}
	if m.IsPerpetual {
		n += 2
	}
	l = m.DistributeTo.Size()
	n += 1 + l + sovGauge(uint64(l))
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovGauge(uint64(l))
		}
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovGauge(uint64(l))
	if m.NumEpochsPaidOver != 0 {
		n += 1 + sovGauge(uint64(m.NumEpochsPaidOver))
	}
	if m.FilledEpochs != 0 {
		n += 1 + sovGauge(uint64(m.FilledEpochs))
	}
	if len(m.DistributedCoins) > 0 {
		for _, e := range m.DistributedCoins {
			l = e.Size()
			n += 1 + l + sovGauge(uint64(l))
		}
	}
	if m.PricingTick != 0 {
		n += 1 + sovGauge(uint64(m.PricingTick))
	}
	return n
}

func (m *QueryCondition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PairID != nil {
		l = m.PairID.Size()
		n += 1 + l + sovGauge(uint64(l))
	}
	if m.StartTick != 0 {
		n += 1 + sovGauge(uint64(m.StartTick))
	}
	if m.EndTick != 0 {
		n += 1 + sovGauge(uint64(m.EndTick))
	}
	return n
}

func sovGauge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGauge(x uint64) (n int) {
	return sovGauge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Gauge) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGauge
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
			return fmt.Errorf("proto: Gauge: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Gauge: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsPerpetual", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
			m.IsPerpetual = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributeTo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
				return ErrInvalidLengthGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributeTo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
				return ErrInvalidLengthGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
				return ErrInvalidLengthGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumEpochsPaidOver", wireType)
			}
			m.NumEpochsPaidOver = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumEpochsPaidOver |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FilledEpochs", wireType)
			}
			m.FilledEpochs = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FilledEpochs |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributedCoins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
				return ErrInvalidLengthGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DistributedCoins = append(m.DistributedCoins, types.Coin{})
			if err := m.DistributedCoins[len(m.DistributedCoins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PricingTick", wireType)
			}
			m.PricingTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PricingTick |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGauge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGauge
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
func (m *QueryCondition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGauge
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
			return fmt.Errorf("proto: QueryCondition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryCondition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairID", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
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
				return ErrInvalidLengthGauge
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGauge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PairID == nil {
				m.PairID = &types1.PairID{}
			}
			if err := m.PairID.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTick", wireType)
			}
			m.StartTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartTick |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTick", wireType)
			}
			m.EndTick = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGauge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.EndTick |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGauge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGauge
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
func skipGauge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGauge
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
					return 0, ErrIntOverflowGauge
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
					return 0, ErrIntOverflowGauge
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
				return 0, ErrInvalidLengthGauge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGauge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGauge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGauge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGauge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGauge = fmt.Errorf("proto: unexpected end of group")
)

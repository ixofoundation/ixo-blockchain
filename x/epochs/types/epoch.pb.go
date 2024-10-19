// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/epochs/v1beta1/epoch.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
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

// EpochInfo is a struct that describes the data going into
// a timer defined by the x/epochs module.
type EpochInfo struct {
	// identifier is a unique reference to this particular timer.
	Identifier string `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	// start_time is the time at which the timer first ever ticks.
	// If start_time is in the future, the epoch will not begin until the start
	// time.
	StartTime time.Time `protobuf:"bytes,2,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time" yaml:"start_time"`
	// duration is the time in between epoch ticks.
	// In order for intended behavior to be met, duration should
	// be greater than the chains expected block time.
	// Duration must be non-zero.
	Duration time.Duration `protobuf:"bytes,3,opt,name=duration,proto3,stdduration" json:"duration,omitempty" yaml:"duration"`
	// current_epoch is the current epoch number, or in other words,
	// how many times has the timer 'ticked'.
	// The first tick (current_epoch=1) is defined as
	// the first block whose blocktime is greater than the EpochInfo start_time.
	CurrentEpoch int64 `protobuf:"varint,4,opt,name=current_epoch,json=currentEpoch,proto3" json:"current_epoch,omitempty"`
	// current_epoch_start_time describes the start time of the current timer
	// interval. The interval is (current_epoch_start_time,
	// current_epoch_start_time + duration] When the timer ticks, this is set to
	// current_epoch_start_time = last_epoch_start_time + duration only one timer
	// tick for a given identifier can occur per block.
	//
	// NOTE! The current_epoch_start_time may diverge significantly from the
	// wall-clock time the epoch began at. Wall-clock time of epoch start may be
	// >> current_epoch_start_time. Suppose current_epoch_start_time = 10,
	// duration = 5. Suppose the chain goes offline at t=14, and comes back online
	// at t=30, and produces blocks at every successive time. (t=31, 32, etc.)
	// * The t=30 block will start the epoch for (10, 15]
	// * The t=31 block will start the epoch for (15, 20]
	// * The t=32 block will start the epoch for (20, 25]
	// * The t=33 block will start the epoch for (25, 30]
	// * The t=34 block will start the epoch for (30, 35]
	// * The **t=36** block will start the epoch for (35, 40]
	CurrentEpochStartTime time.Time `protobuf:"bytes,5,opt,name=current_epoch_start_time,json=currentEpochStartTime,proto3,stdtime" json:"current_epoch_start_time" yaml:"current_epoch_start_time"`
	// epoch_counting_started is a boolean, that indicates whether this
	// epoch timer has began yet.
	EpochCountingStarted bool `protobuf:"varint,6,opt,name=epoch_counting_started,json=epochCountingStarted,proto3" json:"epoch_counting_started,omitempty"`
	// current_epoch_start_height is the block height at which the current epoch
	// started. (The block height at which the timer last ticked)
	CurrentEpochStartHeight int64 `protobuf:"varint,8,opt,name=current_epoch_start_height,json=currentEpochStartHeight,proto3" json:"current_epoch_start_height,omitempty"`
}

func (m *EpochInfo) Reset()         { *m = EpochInfo{} }
func (m *EpochInfo) String() string { return proto.CompactTextString(m) }
func (*EpochInfo) ProtoMessage()    {}
func (*EpochInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e993fff647fe5e0, []int{0}
}
func (m *EpochInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EpochInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EpochInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EpochInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EpochInfo.Merge(m, src)
}
func (m *EpochInfo) XXX_Size() int {
	return m.Size()
}
func (m *EpochInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_EpochInfo.DiscardUnknown(m)
}

var xxx_messageInfo_EpochInfo proto.InternalMessageInfo

func (m *EpochInfo) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *EpochInfo) GetStartTime() time.Time {
	if m != nil {
		return m.StartTime
	}
	return time.Time{}
}

func (m *EpochInfo) GetDuration() time.Duration {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *EpochInfo) GetCurrentEpoch() int64 {
	if m != nil {
		return m.CurrentEpoch
	}
	return 0
}

func (m *EpochInfo) GetCurrentEpochStartTime() time.Time {
	if m != nil {
		return m.CurrentEpochStartTime
	}
	return time.Time{}
}

func (m *EpochInfo) GetEpochCountingStarted() bool {
	if m != nil {
		return m.EpochCountingStarted
	}
	return false
}

func (m *EpochInfo) GetCurrentEpochStartHeight() int64 {
	if m != nil {
		return m.CurrentEpochStartHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*EpochInfo)(nil), "ixo.epochs.v1beta1.EpochInfo")
}

func init() { proto.RegisterFile("ixo/epochs/v1beta1/epoch.proto", fileDescriptor_8e993fff647fe5e0) }

var fileDescriptor_8e993fff647fe5e0 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x31, 0x8f, 0xd3, 0x30,
	0x1c, 0xc5, 0x6b, 0xae, 0x1c, 0xa9, 0x01, 0x01, 0xd6, 0x01, 0xa1, 0x12, 0x4e, 0x55, 0x96, 0x4a,
	0x40, 0xac, 0xc2, 0x2d, 0xc0, 0x56, 0x40, 0x02, 0xc6, 0x96, 0x01, 0xb1, 0x54, 0x49, 0xea, 0x24,
	0x16, 0x8d, 0xff, 0x51, 0xe2, 0x9c, 0xd2, 0x8d, 0x8f, 0xd0, 0x91, 0x8f, 0x74, 0xe3, 0x8d, 0x4c,
	0x01, 0xb5, 0x1b, 0xe3, 0x7d, 0x02, 0x14, 0x3b, 0x39, 0x0a, 0x05, 0xb1, 0xc5, 0xfe, 0xbd, 0xff,
	0x7b, 0xfe, 0x3f, 0x05, 0x53, 0x51, 0x02, 0xe3, 0x29, 0x04, 0x71, 0xce, 0x4e, 0xc6, 0x3e, 0x57,
	0xde, 0xd8, 0x1c, 0xdd, 0x34, 0x03, 0x05, 0x84, 0x88, 0x12, 0x5c, 0xc3, 0xdd, 0x86, 0xf7, 0x8f,
	0x22, 0x88, 0x40, 0x63, 0x56, 0x7f, 0x19, 0x65, 0x9f, 0x46, 0x00, 0xd1, 0x92, 0x33, 0x7d, 0xf2,
	0x8b, 0x90, 0x2d, 0x8a, 0xcc, 0x53, 0x02, 0x64, 0xc3, 0x9d, 0x3f, 0xb9, 0x12, 0x09, 0xcf, 0x95,
	0x97, 0xa4, 0x46, 0x30, 0x5c, 0x77, 0x71, 0xef, 0x75, 0x9d, 0xf4, 0x56, 0x86, 0x40, 0x28, 0xc6,
	0x62, 0xc1, 0xa5, 0x12, 0xa1, 0xe0, 0x99, 0x8d, 0x06, 0x68, 0xd4, 0x9b, 0xee, 0xdc, 0x90, 0x0f,
	0x18, 0xe7, 0xca, 0xcb, 0xd4, 0xbc, 0xb6, 0xb1, 0x2f, 0x0d, 0xd0, 0xe8, 0xea, 0x93, 0xbe, 0x6b,
	0x32, 0xdc, 0x36, 0xc3, 0x7d, 0xdf, 0x66, 0x4c, 0xee, 0x9f, 0x56, 0x4e, 0xe7, 0xbc, 0x72, 0x6e,
	0xad, 0xbc, 0x64, 0xf9, 0x7c, 0xf8, 0x6b, 0x76, 0xb8, 0xfe, 0xe6, 0xa0, 0x69, 0x4f, 0x5f, 0xd4,
	0x72, 0x12, 0x63, 0xab, 0x7d, 0xba, 0x7d, 0xa0, 0x7d, 0xef, 0xed, 0xf9, 0xbe, 0x6a, 0x04, 0x93,
	0x71, 0x6d, 0xfb, 0xa3, 0x72, 0x48, 0x3b, 0xf2, 0x08, 0x12, 0xa1, 0x78, 0x92, 0xaa, 0xd5, 0x79,
	0xe5, 0xdc, 0x30, 0x61, 0x2d, 0x1b, 0x7e, 0xa9, 0xa3, 0x2e, 0xdc, 0xc9, 0x03, 0x7c, 0x3d, 0x28,
	0xb2, 0x8c, 0x4b, 0x35, 0xd7, 0x15, 0xdb, 0xdd, 0x01, 0x1a, 0x1d, 0x4c, 0xaf, 0x35, 0x97, 0xba,
	0x0c, 0xf2, 0x19, 0x61, 0xfb, 0x37, 0xd5, 0x7c, 0x67, 0xef, 0xcb, 0xff, 0xdd, 0xfb, 0x61, 0xb3,
	0xb7, 0x63, 0x9e, 0xf2, 0x2f, 0x27, 0xd3, 0xc2, 0xed, 0xdd, 0xe4, 0xd9, 0x45, 0x23, 0xc7, 0xf8,
	0x8e, 0xd1, 0x07, 0x50, 0x48, 0x25, 0x64, 0x64, 0x06, 0xf9, 0xc2, 0x3e, 0x1c, 0xa0, 0x91, 0x35,
	0x3d, 0xd2, 0xf4, 0x65, 0x03, 0x67, 0x86, 0x91, 0x17, 0xb8, 0xff, 0xb7, 0xb4, 0x98, 0x8b, 0x28,
	0x56, 0xb6, 0xa5, 0x57, 0xbd, 0xbb, 0x17, 0xf8, 0x46, 0xe3, 0x77, 0x5d, 0xeb, 0xca, 0x4d, 0x6b,
	0x32, 0x3b, 0xdd, 0x50, 0x74, 0xb6, 0xa1, 0xe8, 0xfb, 0x86, 0xa2, 0xf5, 0x96, 0x76, 0xce, 0xb6,
	0xb4, 0xf3, 0x75, 0x4b, 0x3b, 0x1f, 0x9f, 0x45, 0x42, 0xc5, 0x85, 0xef, 0x06, 0x90, 0x30, 0x51,
	0x42, 0x08, 0x85, 0x5c, 0xe8, 0x52, 0xeb, 0xd3, 0x63, 0x7f, 0x09, 0xc1, 0xa7, 0x20, 0xf6, 0x84,
	0x64, 0x27, 0xc7, 0xac, 0x6c, 0x7f, 0x70, 0xb5, 0x4a, 0x79, 0xee, 0x1f, 0xea, 0x96, 0x9e, 0xfe,
	0x0c, 0x00, 0x00, 0xff, 0xff, 0x11, 0xda, 0x5a, 0x63, 0xfb, 0x02, 0x00, 0x00,
}

func (m *EpochInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EpochInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EpochInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CurrentEpochStartHeight != 0 {
		i = encodeVarintEpoch(dAtA, i, uint64(m.CurrentEpochStartHeight))
		i--
		dAtA[i] = 0x40
	}
	if m.EpochCountingStarted {
		i--
		if m.EpochCountingStarted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.CurrentEpochStartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CurrentEpochStartTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintEpoch(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if m.CurrentEpoch != 0 {
		i = encodeVarintEpoch(dAtA, i, uint64(m.CurrentEpoch))
		i--
		dAtA[i] = 0x20
	}
	n2, err2 := github_com_cosmos_gogoproto_types.StdDurationMarshalTo(m.Duration, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintEpoch(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x1a
	n3, err3 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintEpoch(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x12
	if len(m.Identifier) > 0 {
		i -= len(m.Identifier)
		copy(dAtA[i:], m.Identifier)
		i = encodeVarintEpoch(dAtA, i, uint64(len(m.Identifier)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEpoch(dAtA []byte, offset int, v uint64) int {
	offset -= sovEpoch(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EpochInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identifier)
	if l > 0 {
		n += 1 + l + sovEpoch(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovEpoch(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdDuration(m.Duration)
	n += 1 + l + sovEpoch(uint64(l))
	if m.CurrentEpoch != 0 {
		n += 1 + sovEpoch(uint64(m.CurrentEpoch))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CurrentEpochStartTime)
	n += 1 + l + sovEpoch(uint64(l))
	if m.EpochCountingStarted {
		n += 2
	}
	if m.CurrentEpochStartHeight != 0 {
		n += 1 + sovEpoch(uint64(m.CurrentEpochStartHeight))
	}
	return n
}

func sovEpoch(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEpoch(x uint64) (n int) {
	return sovEpoch(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EpochInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEpoch
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
			return fmt.Errorf("proto: EpochInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EpochInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifier", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifier = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdDurationUnmarshal(&m.Duration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpoch", wireType)
			}
			m.CurrentEpoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpoch |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
				return ErrInvalidLengthEpoch
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEpoch
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.CurrentEpochStartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochCountingStarted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
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
			m.EpochCountingStarted = bool(v != 0)
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentEpochStartHeight", wireType)
			}
			m.CurrentEpochStartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEpoch
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentEpochStartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEpoch(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEpoch
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
func skipEpoch(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEpoch
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
					return 0, ErrIntOverflowEpoch
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
					return 0, ErrIntOverflowEpoch
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
				return 0, ErrInvalidLengthEpoch
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEpoch
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEpoch
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEpoch        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEpoch          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEpoch = fmt.Errorf("proto: unexpected end of group")
)

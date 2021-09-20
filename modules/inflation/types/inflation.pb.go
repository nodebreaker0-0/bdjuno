// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: em/inflation/v1/inflation.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

type InflationAsset struct {
	Denom     string                                 `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty" yaml:"denom"`
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation" yaml:"inflation"`
	Accum     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=accum,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"accum" yaml:"accum"`
}

func (m *InflationAsset) Reset()         { *m = InflationAsset{} }
func (m *InflationAsset) String() string { return proto.CompactTextString(m) }
func (*InflationAsset) ProtoMessage()    {}
func (*InflationAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_25d8d858c54688c8, []int{0}
}
func (m *InflationAsset) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InflationAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InflationAsset.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InflationAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InflationAsset.Merge(m, src)
}
func (m *InflationAsset) XXX_Size() int {
	return m.Size()
}
func (m *InflationAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_InflationAsset.DiscardUnknown(m)
}

var xxx_messageInfo_InflationAsset proto.InternalMessageInfo

func (m *InflationAsset) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

type InflationState struct {
	LastAppliedTime   time.Time                              `protobuf:"bytes,1,opt,name=last_applied,json=lastApplied,proto3,stdtime" json:"last_applied" yaml:"last_applied"`
	LastAppliedHeight github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,2,opt,name=last_applied_height,json=lastAppliedHeight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"last_applied_height" yaml:"last_applied_height"`
	InflationAssets   []InflationAsset                       `protobuf:"bytes,3,rep,name=assets,proto3" json:"assets" yaml:"assets"`
}

func (m *InflationState) Reset()      { *m = InflationState{} }
func (*InflationState) ProtoMessage() {}
func (*InflationState) Descriptor() ([]byte, []int) {
	return fileDescriptor_25d8d858c54688c8, []int{1}
}
func (m *InflationState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InflationState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InflationState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InflationState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InflationState.Merge(m, src)
}
func (m *InflationState) XXX_Size() int {
	return m.Size()
}
func (m *InflationState) XXX_DiscardUnknown() {
	xxx_messageInfo_InflationState.DiscardUnknown(m)
}

var xxx_messageInfo_InflationState proto.InternalMessageInfo

func (m *InflationState) GetLastAppliedTime() time.Time {
	if m != nil {
		return m.LastAppliedTime
	}
	return time.Time{}
}

func (m *InflationState) GetInflationAssets() []InflationAsset {
	if m != nil {
		return m.InflationAssets
	}
	return nil
}

func init() {
	proto.RegisterType((*InflationAsset)(nil), "em.inflation.v1.InflationAsset")
	proto.RegisterType((*InflationState)(nil), "em.inflation.v1.InflationState")
}

func init() { proto.RegisterFile("em/inflation/v1/inflation.proto", fileDescriptor_25d8d858c54688c8) }

var fileDescriptor_25d8d858c54688c8 = []byte{
	// 473 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0xed, 0x86, 0x56, 0xaa, 0x53, 0x68, 0x71, 0x19, 0x22, 0x0f, 0xbe, 0xea, 0x86, 0x2a,
	0x4b, 0xee, 0x94, 0x30, 0x20, 0x75, 0x40, 0xaa, 0xc5, 0x40, 0xa5, 0x2e, 0x98, 0x4e, 0x2c, 0xe5,
	0xec, 0xbc, 0x75, 0x2c, 0x7c, 0x3e, 0x93, 0xbb, 0x44, 0x44, 0xe2, 0x43, 0x74, 0x64, 0xe4, 0xe3,
	0x74, 0xec, 0x88, 0x18, 0x0c, 0x4a, 0x36, 0x26, 0xe4, 0x4f, 0x80, 0x72, 0xe7, 0xd4, 0x2e, 0x9d,
	0x3a, 0xd9, 0xef, 0x1f, 0xff, 0xde, 0xf7, 0x79, 0xee, 0xec, 0x20, 0xe0, 0x34, 0xcd, 0xaf, 0x32,
	0xa6, 0x52, 0x91, 0xd3, 0xf9, 0xb0, 0x09, 0x48, 0x31, 0x15, 0x4a, 0xb8, 0xfb, 0xc0, 0x49, 0x93,
	0x9b, 0x0f, 0xbd, 0x17, 0x89, 0x48, 0x84, 0xae, 0xd1, 0xf5, 0x9b, 0x69, 0xf3, 0xfc, 0x58, 0x48,
	0x2e, 0x24, 0x8d, 0x98, 0x04, 0x3a, 0x1f, 0x46, 0xa0, 0xd8, 0x90, 0xc6, 0x22, 0xad, 0x31, 0x1e,
	0x4a, 0x84, 0x48, 0x32, 0xa0, 0x3a, 0x8a, 0x66, 0x57, 0x54, 0xa5, 0x1c, 0xa4, 0x62, 0xbc, 0x30,
	0x0d, 0xf8, 0xaf, 0xed, 0x3c, 0x3b, 0xdb, 0xcc, 0x39, 0x95, 0x12, 0x94, 0x7b, 0xec, 0x6c, 0x8f,
	0x21, 0x17, 0xbc, 0x67, 0x1f, 0xd9, 0xfd, 0xdd, 0xe0, 0xa0, 0x2a, 0xd1, 0xde, 0x82, 0xf1, 0xec,
	0x04, 0xeb, 0x34, 0x0e, 0x4d, 0xd9, 0xfd, 0xe8, 0xec, 0xde, 0x6d, 0xd8, 0xdb, 0xd2, 0xbd, 0xc1,
	0x4d, 0x89, 0xac, 0x9f, 0x25, 0x3a, 0x4e, 0x52, 0x35, 0x99, 0x45, 0x24, 0x16, 0x9c, 0xd6, 0x1b,
	0x9a, 0xc7, 0x40, 0x8e, 0x3f, 0x51, 0xb5, 0x28, 0x40, 0x92, 0x37, 0x10, 0x57, 0x25, 0x3a, 0x30,
	0xe4, 0x3b, 0x10, 0x0e, 0x1b, 0xa8, 0x7b, 0xe1, 0x6c, 0xb3, 0x38, 0x9e, 0xf1, 0x5e, 0x47, 0xd3,
	0x5f, 0x3f, 0x9a, 0x5e, 0xef, 0xad, 0x21, 0x38, 0x34, 0x30, 0xfc, 0x67, 0xab, 0x25, 0xf9, 0xbd,
	0x62, 0x0a, 0xdc, 0xcf, 0xce, 0x5e, 0xc6, 0xa4, 0xba, 0x64, 0x45, 0x91, 0xa5, 0x30, 0xd6, 0xca,
	0xbb, 0x23, 0x8f, 0x18, 0xf7, 0xc8, 0xc6, 0x3d, 0x72, 0xb1, 0x71, 0x2f, 0x18, 0xad, 0x77, 0x59,
	0x96, 0x68, 0xff, 0x9c, 0x49, 0x75, 0x6a, 0x3e, 0x5b, 0x57, 0xab, 0x12, 0x1d, 0x9a, 0xa1, 0x6d,
	0x20, 0xbe, 0xfe, 0x85, 0xec, 0xb0, 0x9b, 0x35, 0xbd, 0xee, 0x57, 0xe7, 0xb0, 0xdd, 0x71, 0x39,
	0x81, 0x34, 0x99, 0xa8, 0xda, 0xc7, 0xf3, 0x47, 0x28, 0x3d, 0xcb, 0x55, 0x55, 0x22, 0xef, 0xe1,
	0xd0, 0x1a, 0x89, 0xc3, 0xe7, 0xad, 0xb9, 0x6f, 0x75, 0xce, 0x65, 0xce, 0x0e, 0x5b, 0x1f, 0xb6,
	0xec, 0x75, 0x8e, 0x3a, 0xfd, 0xee, 0x08, 0x91, 0xff, 0xee, 0x1b, 0xb9, 0x7f, 0x29, 0x82, 0xfe,
	0x46, 0xef, 0xfd, 0xbc, 0xac, 0x4a, 0xf4, 0xb4, 0x36, 0x59, 0xc7, 0x38, 0xac, 0xc1, 0x27, 0x4f,
	0xbe, 0x7d, 0x47, 0x56, 0xf0, 0xee, 0x66, 0xe9, 0xdb, 0xb7, 0x4b, 0xdf, 0xfe, 0xbd, 0xf4, 0xed,
	0xeb, 0x95, 0x6f, 0xdd, 0xae, 0x7c, 0xeb, 0xc7, 0xca, 0xb7, 0x3e, 0xbc, 0x6a, 0x69, 0x83, 0x01,
	0x17, 0x39, 0x2c, 0x28, 0xf0, 0x41, 0x06, 0xe3, 0x04, 0xa6, 0xf4, 0x4b, 0xeb, 0xf7, 0x48, 0x73,
	0x05, 0xd3, 0x9c, 0x65, 0x46, 0x70, 0xb4, 0xa3, 0x8f, 0xe3, 0xe5, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x2b, 0x74, 0xef, 0x08, 0x44, 0x03, 0x00, 0x00,
}

func (m *InflationAsset) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InflationAsset) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InflationAsset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Accum.Size()
		i -= size
		if _, err := m.Accum.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintInflation(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *InflationState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InflationState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InflationState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.InflationAssets) > 0 {
		for iNdEx := len(m.InflationAssets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.InflationAssets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintInflation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	{
		size := m.LastAppliedHeight.Size()
		i -= size
		if _, err := m.LastAppliedHeight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintInflation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.LastAppliedTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.LastAppliedTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintInflation(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintInflation(dAtA []byte, offset int, v uint64) int {
	offset -= sovInflation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InflationAsset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovInflation(uint64(l))
	}
	l = m.Inflation.Size()
	n += 1 + l + sovInflation(uint64(l))
	l = m.Accum.Size()
	n += 1 + l + sovInflation(uint64(l))
	return n
}

func (m *InflationState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.LastAppliedTime)
	n += 1 + l + sovInflation(uint64(l))
	l = m.LastAppliedHeight.Size()
	n += 1 + l + sovInflation(uint64(l))
	if len(m.InflationAssets) > 0 {
		for _, e := range m.InflationAssets {
			l = e.Size()
			n += 1 + l + sovInflation(uint64(l))
		}
	}
	return n
}

func sovInflation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozInflation(x uint64) (n int) {
	return sovInflation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InflationAsset) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInflation
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
			return fmt.Errorf("proto: InflationAsset: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InflationAsset: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Accum", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Accum.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInflation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInflation
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
func (m *InflationState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowInflation
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
			return fmt.Errorf("proto: InflationState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InflationState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastAppliedTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.LastAppliedTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastAppliedHeight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastAppliedHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationAssets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowInflation
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
				return ErrInvalidLengthInflation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthInflation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InflationAssets = append(m.InflationAssets, InflationAsset{})
			if err := m.InflationAssets[len(m.InflationAssets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipInflation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthInflation
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
func skipInflation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowInflation
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
					return 0, ErrIntOverflowInflation
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
					return 0, ErrIntOverflowInflation
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
				return 0, ErrInvalidLengthInflation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupInflation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthInflation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthInflation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowInflation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupInflation = fmt.Errorf("proto: unexpected end of group")
)
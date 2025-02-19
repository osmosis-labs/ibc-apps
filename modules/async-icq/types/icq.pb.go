// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: icq/v1/icq.proto

package types

import (
	fmt "fmt"
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

// Params defines the set of on-chain interchain query parameters.
type Params struct {
	// host_enabled enables or disables the host submodule.
	HostEnabled bool `protobuf:"varint,2,opt,name=host_enabled,json=hostEnabled,proto3" json:"host_enabled,omitempty" yaml:"host_enabled"`
	// allow_queries defines a list of query paths allowed to be queried on a host chain.
	AllowQueries []string `protobuf:"bytes,3,rep,name=allow_queries,json=allowQueries,proto3" json:"allow_queries,omitempty" yaml:"allow_queries"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a9dc71eedc8bea6, []int{0}
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

func (m *Params) GetHostEnabled() bool {
	if m != nil {
		return m.HostEnabled
	}
	return false
}

func (m *Params) GetAllowQueries() []string {
	if m != nil {
		return m.AllowQueries
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "icq.v1.Params")
}

func init() { proto.RegisterFile("icq/v1/icq.proto", fileDescriptor_0a9dc71eedc8bea6) }

var fileDescriptor_0a9dc71eedc8bea6 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x4c, 0x2e, 0xd4,
	0x2f, 0x33, 0xd4, 0xcf, 0x4c, 0x2e, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x31,
	0xcb, 0x0c, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x42, 0xfa, 0x20, 0x16, 0x44, 0x56, 0xa9,
	0x99, 0x91, 0x8b, 0x2d, 0x20, 0xb1, 0x28, 0x31, 0xb7, 0x58, 0xc8, 0x8a, 0x8b, 0x27, 0x23, 0xbf,
	0xb8, 0x24, 0x3e, 0x35, 0x2f, 0x31, 0x29, 0x27, 0x35, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0xc3,
	0x49, 0xfc, 0xd3, 0x3d, 0x79, 0xe1, 0xca, 0xc4, 0xdc, 0x1c, 0x2b, 0x25, 0x64, 0x59, 0xa5, 0x20,
	0x6e, 0x10, 0xd7, 0x15, 0xc2, 0x13, 0xb2, 0xe5, 0xe2, 0x4d, 0xcc, 0xc9, 0xc9, 0x2f, 0x8f, 0x2f,
	0x2c, 0x4d, 0x2d, 0xca, 0x4c, 0x2d, 0x96, 0x60, 0x56, 0x60, 0xd6, 0xe0, 0x74, 0x92, 0xf8, 0x74,
	0x4f, 0x5e, 0x04, 0xa2, 0x19, 0x45, 0x5a, 0x29, 0x88, 0x07, 0xcc, 0x0f, 0x84, 0x70, 0x9d, 0xfc,
	0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5,
	0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x34, 0x3d, 0xb3, 0x24, 0xa3,
	0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x39, 0xbf, 0x38, 0x37, 0xbf, 0x58, 0x3f, 0x33, 0x29,
	0x59, 0x37, 0xb1, 0xa0, 0xa0, 0x58, 0x3f, 0x37, 0x3f, 0xa5, 0x34, 0x27, 0xb5, 0x58, 0x3f, 0xb1,
	0xb8, 0x32, 0x2f, 0x59, 0x17, 0xec, 0x71, 0x73, 0xfd, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36,
	0xb0, 0xef, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x99, 0x62, 0x76, 0x0f, 0x01, 0x00,
	0x00,
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
	if len(m.AllowQueries) > 0 {
		for iNdEx := len(m.AllowQueries) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.AllowQueries[iNdEx])
			copy(dAtA[i:], m.AllowQueries[iNdEx])
			i = encodeVarintIcq(dAtA, i, uint64(len(m.AllowQueries[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.HostEnabled {
		i--
		if m.HostEnabled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	return len(dAtA) - i, nil
}

func encodeVarintIcq(dAtA []byte, offset int, v uint64) int {
	offset -= sovIcq(v)
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
	if m.HostEnabled {
		n += 2
	}
	if len(m.AllowQueries) > 0 {
		for _, s := range m.AllowQueries {
			l = len(s)
			n += 1 + l + sovIcq(uint64(l))
		}
	}
	return n
}

func sovIcq(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIcq(x uint64) (n int) {
	return sovIcq(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIcq
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
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostEnabled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
			m.HostEnabled = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowQueries", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIcq
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
				return ErrInvalidLengthIcq
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIcq
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AllowQueries = append(m.AllowQueries, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIcq(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIcq
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
func skipIcq(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIcq
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
					return 0, ErrIntOverflowIcq
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
					return 0, ErrIntOverflowIcq
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
				return 0, ErrInvalidLengthIcq
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIcq
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIcq
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIcq        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIcq          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIcq = fmt.Errorf("proto: unexpected end of group")
)

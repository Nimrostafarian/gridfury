// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: guardian/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// MsgAddSuper defines the properties of add super account message
type MsgAddSuper struct {
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	AddedBy     string `protobuf:"bytes,3,opt,name=added_by,json=addedBy,proto3" json:"added_by,omitempty"`
}

func (m *MsgAddSuper) Reset()         { *m = MsgAddSuper{} }
func (m *MsgAddSuper) String() string { return proto.CompactTextString(m) }
func (*MsgAddSuper) ProtoMessage()    {}
func (*MsgAddSuper) Descriptor() ([]byte, []int) {
	return fileDescriptor_b62288115d705ce8, []int{0}
}
func (m *MsgAddSuper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddSuper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddSuper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddSuper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddSuper.Merge(m, src)
}
func (m *MsgAddSuper) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddSuper) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddSuper.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddSuper proto.InternalMessageInfo

func (m *MsgAddSuper) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgAddSuper) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgAddSuper) GetAddedBy() string {
	if m != nil {
		return m.AddedBy
	}
	return ""
}

// MsgAddSuperResponse defines the Msg/AddSuper response type
type MsgAddSuperResponse struct {
}

func (m *MsgAddSuperResponse) Reset()         { *m = MsgAddSuperResponse{} }
func (m *MsgAddSuperResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddSuperResponse) ProtoMessage()    {}
func (*MsgAddSuperResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b62288115d705ce8, []int{1}
}
func (m *MsgAddSuperResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddSuperResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddSuperResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddSuperResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddSuperResponse.Merge(m, src)
}
func (m *MsgAddSuperResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddSuperResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddSuperResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddSuperResponse proto.InternalMessageInfo

// MsgDeleteSuper defines the properties of delete super account message
type MsgDeleteSuper struct {
	Address   string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	DeletedBy string `protobuf:"bytes,3,opt,name=deleted_by,json=deletedBy,proto3" json:"deleted_by,omitempty"`
}

func (m *MsgDeleteSuper) Reset()         { *m = MsgDeleteSuper{} }
func (m *MsgDeleteSuper) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteSuper) ProtoMessage()    {}
func (*MsgDeleteSuper) Descriptor() ([]byte, []int) {
	return fileDescriptor_b62288115d705ce8, []int{2}
}
func (m *MsgDeleteSuper) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteSuper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteSuper.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteSuper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteSuper.Merge(m, src)
}
func (m *MsgDeleteSuper) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteSuper) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteSuper.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteSuper proto.InternalMessageInfo

func (m *MsgDeleteSuper) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgDeleteSuper) GetDeletedBy() string {
	if m != nil {
		return m.DeletedBy
	}
	return ""
}

// MsgDeleteSuperResponse defines the Msg/DeleteSuper response type
type MsgDeleteSuperResponse struct {
}

func (m *MsgDeleteSuperResponse) Reset()         { *m = MsgDeleteSuperResponse{} }
func (m *MsgDeleteSuperResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteSuperResponse) ProtoMessage()    {}
func (*MsgDeleteSuperResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b62288115d705ce8, []int{3}
}
func (m *MsgDeleteSuperResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteSuperResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteSuperResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteSuperResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteSuperResponse.Merge(m, src)
}
func (m *MsgDeleteSuperResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteSuperResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteSuperResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteSuperResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddSuper)(nil), "gridhub.guardian.MsgAddSuper")
	proto.RegisterType((*MsgAddSuperResponse)(nil), "gridhub.guardian.MsgAddSuperResponse")
	proto.RegisterType((*MsgDeleteSuper)(nil), "gridhub.guardian.MsgDeleteSuper")
	proto.RegisterType((*MsgDeleteSuperResponse)(nil), "gridhub.guardian.MsgDeleteSuperResponse")
}

func init() { proto.RegisterFile("guardian/tx.proto", fileDescriptor_b62288115d705ce8) }

var fileDescriptor_b62288115d705ce8 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x3b, 0x16, 0xb4, 0xbd, 0x05, 0xd1, 0x11, 0x25, 0x16, 0x3a, 0x94, 0x80, 0xd0, 0x55,
	0x82, 0xfa, 0x04, 0x16, 0x37, 0x22, 0x01, 0xa9, 0x2b, 0xdd, 0x48, 0xd2, 0x7b, 0x8d, 0x81, 0x36,
	0x13, 0x66, 0x26, 0x60, 0xde, 0xc2, 0x67, 0xf1, 0x29, 0x5c, 0x76, 0xe9, 0x52, 0x92, 0x17, 0x91,
	0xc6, 0xa6, 0x8e, 0xe2, 0xcf, 0x72, 0xee, 0x39, 0xf3, 0x9d, 0x3b, 0x73, 0x60, 0x37, 0xce, 0x43,
	0x85, 0x49, 0x98, 0xfa, 0xe6, 0xd1, 0xcb, 0x94, 0x34, 0x92, 0xef, 0xdc, 0xe7, 0xaa, 0x78, 0xc8,
	0x23, 0xaf, 0x91, 0x5c, 0x84, 0x5e, 0xa0, 0xe3, 0x33, 0xc4, 0xeb, 0x3c, 0x23, 0xc5, 0x87, 0xd0,
	0x43, 0xd2, 0x53, 0x95, 0x64, 0x26, 0x91, 0xa9, 0xc3, 0x86, 0x6c, 0xd4, 0x9d, 0xd8, 0x23, 0xee,
	0xc0, 0x56, 0x88, 0xa8, 0x48, 0x6b, 0x67, 0xa3, 0x56, 0x9b, 0x23, 0x3f, 0x84, 0x4e, 0x88, 0x48,
	0x78, 0x17, 0x15, 0x4e, 0x7b, 0x2d, 0x11, 0x8e, 0x0b, 0x77, 0x1f, 0xf6, 0xac, 0x94, 0x09, 0xe9,
	0x4c, 0xa6, 0x9a, 0xdc, 0x0b, 0xd8, 0x0e, 0x74, 0x7c, 0x4e, 0x33, 0x32, 0xf4, 0x91, 0xff, 0x3b,
	0x7d, 0x00, 0x80, 0xb5, 0xd1, 0xe2, 0x77, 0x57, 0x93, 0x71, 0xe1, 0x3a, 0x70, 0xf0, 0x15, 0xd5,
	0x84, 0x9c, 0x3c, 0x33, 0x68, 0x07, 0x3a, 0xe6, 0x57, 0xd0, 0x59, 0x3f, 0x73, 0xe0, 0x7d, 0xff,
	0x08, 0xcf, 0xda, 0xaf, 0x7f, 0xf4, 0xa7, 0xdc, 0x90, 0xf9, 0x0d, 0xf4, 0xec, 0xdd, 0x87, 0x3f,
	0xde, 0xb2, 0x1c, 0xfd, 0xd1, 0x7f, 0x8e, 0x06, 0x3d, 0xbe, 0x7c, 0x29, 0x05, 0x5b, 0x94, 0x82,
	0xbd, 0x95, 0x82, 0x3d, 0x55, 0xa2, 0xb5, 0xa8, 0x44, 0xeb, 0xb5, 0x12, 0xad, 0xdb, 0xe3, 0x38,
	0x31, 0x4b, 0xc2, 0x54, 0xce, 0xfd, 0x25, 0x2d, 0x25, 0xe3, 0xaf, 0xa8, 0xfe, 0x5c, 0x62, 0x3e,
	0x23, 0xed, 0x7f, 0x16, 0x5f, 0x64, 0xa4, 0xa3, 0xcd, 0xba, 0xfc, 0xd3, 0xf7, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x9e, 0xb6, 0xf0, 0xa8, 0x11, 0x02, 0x00, 0x00,
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
	// AddSuper defines a method for adding a super account
	AddSuper(ctx context.Context, in *MsgAddSuper, opts ...grpc.CallOption) (*MsgAddSuperResponse, error)
	// DeleteSuper defines a method for deleting a super account
	DeleteSuper(ctx context.Context, in *MsgDeleteSuper, opts ...grpc.CallOption) (*MsgDeleteSuperResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddSuper(ctx context.Context, in *MsgAddSuper, opts ...grpc.CallOption) (*MsgAddSuperResponse, error) {
	out := new(MsgAddSuperResponse)
	err := c.cc.Invoke(ctx, "/gridhub.guardian.Msg/AddSuper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteSuper(ctx context.Context, in *MsgDeleteSuper, opts ...grpc.CallOption) (*MsgDeleteSuperResponse, error) {
	out := new(MsgDeleteSuperResponse)
	err := c.cc.Invoke(ctx, "/gridhub.guardian.Msg/DeleteSuper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// AddSuper defines a method for adding a super account
	AddSuper(context.Context, *MsgAddSuper) (*MsgAddSuperResponse, error)
	// DeleteSuper defines a method for deleting a super account
	DeleteSuper(context.Context, *MsgDeleteSuper) (*MsgDeleteSuperResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddSuper(ctx context.Context, req *MsgAddSuper) (*MsgAddSuperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSuper not implemented")
}
func (*UnimplementedMsgServer) DeleteSuper(ctx context.Context, req *MsgDeleteSuper) (*MsgDeleteSuperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSuper not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddSuper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddSuper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddSuper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gridhub.guardian.Msg/AddSuper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddSuper(ctx, req.(*MsgAddSuper))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteSuper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteSuper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteSuper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gridhub.guardian.Msg/DeleteSuper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteSuper(ctx, req.(*MsgDeleteSuper))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gridhub.guardian.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSuper",
			Handler:    _Msg_AddSuper_Handler,
		},
		{
			MethodName: "DeleteSuper",
			Handler:    _Msg_DeleteSuper_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "guardian/tx.proto",
}

func (m *MsgAddSuper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddSuper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddSuper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AddedBy) > 0 {
		i -= len(m.AddedBy)
		copy(dAtA[i:], m.AddedBy)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AddedBy)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgAddSuperResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddSuperResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddSuperResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteSuper) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteSuper) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteSuper) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DeletedBy) > 0 {
		i -= len(m.DeletedBy)
		copy(dAtA[i:], m.DeletedBy)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DeletedBy)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteSuperResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteSuperResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteSuperResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgAddSuper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AddedBy)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgAddSuperResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteSuper) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DeletedBy)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteSuperResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgAddSuper) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddSuper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddSuper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddedBy", wireType)
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
			m.AddedBy = string(dAtA[iNdEx:postIndex])
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
func (m *MsgAddSuperResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddSuperResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddSuperResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *MsgDeleteSuper) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteSuper: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteSuper: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedBy", wireType)
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
			m.DeletedBy = string(dAtA[iNdEx:postIndex])
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
func (m *MsgDeleteSuperResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteSuperResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteSuperResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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

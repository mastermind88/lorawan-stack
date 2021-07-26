// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/applicationserver_integrations_storage.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type GetStoredApplicationUpRequest struct {
	// Query upstream messages from all end devices of an application. Cannot be used in conjunction with end_device_ids.
	ApplicationIds *ApplicationIdentifiers `protobuf:"bytes,1,opt,name=application_ids,json=applicationIds,proto3" json:"application_ids,omitempty"`
	// Query upstream messages from a single end device. Cannot be used in conjunction with application_ids.
	EndDeviceIds *EndDeviceIdentifiers `protobuf:"bytes,2,opt,name=end_device_ids,json=endDeviceIds,proto3" json:"end_device_ids,omitempty"`
	// Query upstream messages of a specific type. If not set, then all upstream messages are returned.
	Type string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	// Limit number of results.
	Limit *types.UInt32Value `protobuf:"bytes,4,opt,name=limit,proto3" json:"limit,omitempty"`
	// Query upstream messages after this timestamp only. Cannot be used in conjunction with last.
	After *types.Timestamp `protobuf:"bytes,5,opt,name=after,proto3" json:"after,omitempty"`
	// Query upstream messages before this timestamp only. Cannot be used in conjunction with last.
	Before *types.Timestamp `protobuf:"bytes,6,opt,name=before,proto3" json:"before,omitempty"`
	// Query uplinks on a specific FPort only.
	FPort *types.UInt32Value `protobuf:"bytes,7,opt,name=f_port,json=fPort,proto3" json:"f_port,omitempty"`
	// Order results.
	Order string `protobuf:"bytes,8,opt,name=order,proto3" json:"order,omitempty"`
	// The names of the upstream message fields that should be returned. See the API reference
	// for allowed field names for each type of upstream message.
	FieldMask *types.FieldMask `protobuf:"bytes,9,opt,name=field_mask,json=fieldMask,proto3" json:"field_mask,omitempty"`
	// Query upstream messages that have arrived in the last minutes or hours. Cannot be used in conjunction with after and before.
	Last                 *types.Duration `protobuf:"bytes,10,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetStoredApplicationUpRequest) Reset()      { *m = GetStoredApplicationUpRequest{} }
func (*GetStoredApplicationUpRequest) ProtoMessage() {}
func (*GetStoredApplicationUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ff0e9f52f73d254, []int{0}
}
func (m *GetStoredApplicationUpRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetStoredApplicationUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetStoredApplicationUpRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetStoredApplicationUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStoredApplicationUpRequest.Merge(m, src)
}
func (m *GetStoredApplicationUpRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetStoredApplicationUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStoredApplicationUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStoredApplicationUpRequest proto.InternalMessageInfo

func (m *GetStoredApplicationUpRequest) GetApplicationIds() *ApplicationIdentifiers {
	if m != nil {
		return m.ApplicationIds
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetEndDeviceIds() *EndDeviceIdentifiers {
	if m != nil {
		return m.EndDeviceIds
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetStoredApplicationUpRequest) GetLimit() *types.UInt32Value {
	if m != nil {
		return m.Limit
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetAfter() *types.Timestamp {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetBefore() *types.Timestamp {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetFPort() *types.UInt32Value {
	if m != nil {
		return m.FPort
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetOrder() string {
	if m != nil {
		return m.Order
	}
	return ""
}

func (m *GetStoredApplicationUpRequest) GetFieldMask() *types.FieldMask {
	if m != nil {
		return m.FieldMask
	}
	return nil
}

func (m *GetStoredApplicationUpRequest) GetLast() *types.Duration {
	if m != nil {
		return m.Last
	}
	return nil
}

func init() {
	proto.RegisterType((*GetStoredApplicationUpRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpRequest")
	golang_proto.RegisterType((*GetStoredApplicationUpRequest)(nil), "ttn.lorawan.v3.GetStoredApplicationUpRequest")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/applicationserver_integrations_storage.proto", fileDescriptor_6ff0e9f52f73d254)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/applicationserver_integrations_storage.proto", fileDescriptor_6ff0e9f52f73d254)
}

var fileDescriptor_6ff0e9f52f73d254 = []byte{
	// 789 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x6b, 0x24, 0x45,
	0x14, 0x9e, 0x8a, 0x99, 0x68, 0x2a, 0x31, 0x81, 0x46, 0xa4, 0x1d, 0x36, 0x9d, 0x10, 0x45, 0xf6,
	0x32, 0xdd, 0xcb, 0xcc, 0x45, 0x3d, 0x08, 0x1b, 0x56, 0x25, 0x8a, 0x28, 0xb5, 0xc6, 0xc3, 0x82,
	0x34, 0x35, 0x5d, 0xaf, 0x3b, 0xe5, 0xf4, 0x54, 0xd5, 0x56, 0xbd, 0xe9, 0x18, 0x42, 0x40, 0xfc,
	0x0b, 0x04, 0xff, 0x81, 0x3d, 0x2e, 0x78, 0xf0, 0xea, 0xc1, 0x83, 0x07, 0x0f, 0xe2, 0x49, 0x11,
	0xc1, 0xa3, 0x4e, 0x3c, 0x78, 0xf4, 0xbc, 0x78, 0x90, 0xe9, 0xee, 0xf9, 0x6d, 0x76, 0xf7, 0x56,
	0xef, 0xd5, 0xf7, 0x7d, 0xfd, 0xba, 0xde, 0xfb, 0x1e, 0x7d, 0x33, 0xd7, 0x96, 0x9f, 0x71, 0xd5,
	0x76, 0xc8, 0x93, 0x7e, 0xc4, 0x8d, 0x8c, 0xb8, 0x31, 0xb9, 0x4c, 0x38, 0x4a, 0xad, 0x1c, 0xd8,
	0x02, 0x6c, 0x2c, 0x15, 0x42, 0x66, 0xab, 0x4c, 0xec, 0x50, 0x5b, 0x9e, 0x41, 0x68, 0xac, 0x46,
	0xed, 0xed, 0x20, 0xaa, 0xb0, 0xd6, 0x08, 0x8b, 0x6e, 0xeb, 0x76, 0x26, 0xf1, 0x74, 0xd8, 0x0b,
	0x13, 0x3d, 0x88, 0x40, 0x15, 0xfa, 0xdc, 0x58, 0xfd, 0xd9, 0x79, 0x54, 0x82, 0x93, 0x76, 0x06,
	0xaa, 0x5d, 0xf0, 0x5c, 0x0a, 0x8e, 0x10, 0xad, 0x1c, 0x2a, 0xc9, 0x56, 0x7b, 0x4e, 0x22, 0xd3,
	0x99, 0xae, 0xc8, 0xbd, 0x61, 0x5a, 0x46, 0x65, 0x50, 0x9e, 0x6a, 0xf8, 0x8d, 0x4c, 0xeb, 0x2c,
	0x87, 0xaa, 0x74, 0xa5, 0x34, 0x56, 0x75, 0xd6, 0xb7, 0x07, 0xf5, 0xed, 0x54, 0x23, 0x95, 0x90,
	0x8b, 0x78, 0xc0, 0x5d, 0xbf, 0x46, 0xec, 0x2f, 0x23, 0x50, 0x0e, 0xc0, 0x21, 0x1f, 0x98, 0x1a,
	0x10, 0x2c, 0x03, 0xc4, 0xb0, 0x7a, 0x8b, 0xeb, 0xee, 0xcf, 0x2c, 0x37, 0x06, 0xec, 0xa4, 0x84,
	0x97, 0x57, 0x9f, 0x58, 0x0a, 0x50, 0x28, 0x53, 0x39, 0x03, 0x1d, 0xac, 0x82, 0x06, 0xe0, 0x1c,
	0xcf, 0xa0, 0x46, 0x1c, 0xfe, 0xdb, 0xa4, 0x7b, 0xef, 0x00, 0xde, 0x45, 0x6d, 0x41, 0xdc, 0x9e,
	0xf5, 0xe8, 0xc4, 0x30, 0xb8, 0x3f, 0x04, 0x87, 0xde, 0x07, 0x74, 0x77, 0xae, 0x77, 0xb1, 0x14,
	0xce, 0x27, 0x07, 0xe4, 0xe6, 0x56, 0xe7, 0xd5, 0x70, 0xb1, 0x4b, 0xe1, 0x1c, 0xfd, 0x78, 0x56,
	0x0a, 0xdb, 0xe1, 0xf3, 0x79, 0xe7, 0xbd, 0x4b, 0x77, 0x40, 0x89, 0x58, 0x40, 0x21, 0x13, 0x28,
	0xf5, 0xd6, 0x4a, 0xbd, 0x57, 0x96, 0xf5, 0xde, 0x52, 0xe2, 0x4e, 0x09, 0x9a, 0x57, 0xdb, 0x86,
	0x59, 0xd6, 0x79, 0x3f, 0x10, 0xba, 0x8e, 0xe7, 0x06, 0xfc, 0x67, 0x0e, 0xc8, 0xcd, 0xcd, 0xa3,
	0x6f, 0xc8, 0xa3, 0xa3, 0xaf, 0x89, 0x7d, 0x48, 0x58, 0x83, 0xed, 0x0c, 0x4d, 0x2e, 0x55, 0x3f,
	0xae, 0x7f, 0x98, 0x6d, 0x7d, 0xaa, 0xa5, 0x8a, 0x79, 0x92, 0x80, 0x41, 0xb6, 0x2d, 0xf4, 0x99,
	0x2a, 0xaf, 0x79, 0xd2, 0x67, 0xcf, 0x4f, 0x23, 0xb5, 0x18, 0x3a, 0x50, 0xc8, 0x76, 0xa7, 0x61,
	0xca, 0x65, 0x0e, 0x62, 0x2e, 0x71, 0x7f, 0x08, 0x43, 0x10, 0xac, 0xb5, 0x98, 0x88, 0xa5, 0x9a,
	0x0c, 0x9f, 0x60, 0xbb, 0xb9, 0xae, 0x5f, 0xce, 0xe9, 0xbc, 0x00, 0xc1, 0xb6, 0xc7, 0xe3, 0x3f,
	0xfe, 0x73, 0xc1, 0x91, 0xb3, 0xb2, 0x7a, 0xaf, 0x43, 0x9b, 0xb9, 0x1c, 0x48, 0xf4, 0xd7, 0xcb,
	0x97, 0xb8, 0x11, 0x56, 0xcd, 0x0f, 0x27, 0xcd, 0x0f, 0x4f, 0x8e, 0x15, 0x76, 0x3b, 0x1f, 0xf3,
	0x7c, 0x08, 0xac, 0x82, 0x7a, 0xb7, 0x68, 0x93, 0xa7, 0x08, 0xd6, 0x6f, 0x96, 0x9c, 0xd6, 0x0a,
	0xe7, 0xa3, 0xc9, 0xc4, 0xb1, 0x0a, 0xe8, 0x75, 0xe8, 0x46, 0x0f, 0x52, 0x6d, 0xc1, 0xdf, 0x78,
	0x22, 0xa5, 0x46, 0x7a, 0x5d, 0xba, 0x91, 0xc6, 0x46, 0x5b, 0xf4, 0x9f, 0x7d, 0x9a, 0xd2, 0xd2,
	0x0f, 0xb5, 0x45, 0xef, 0x35, 0xda, 0xd4, 0x56, 0x80, 0xf5, 0x9f, 0x2b, 0xbb, 0x72, 0xf8, 0xe8,
	0x68, 0xdf, 0xee, 0xb1, 0x06, 0xdb, 0x6e, 0x5b, 0x48, 0x40, 0x16, 0x20, 0x62, 0x8e, 0x6c, 0x6b,
	0x3e, 0xa8, 0x08, 0xde, 0xeb, 0x94, 0xce, 0xac, 0xe4, 0x6f, 0x5e, 0x53, 0xe6, 0xdb, 0x63, 0xc8,
	0xfb, 0xdc, 0xf5, 0xd9, 0x66, 0x3a, 0x39, 0x7a, 0x6d, 0xba, 0x9e, 0x73, 0x87, 0x3e, 0x2d, 0x49,
	0x2f, 0xad, 0x90, 0xee, 0xd4, 0xfe, 0x62, 0x25, 0xec, 0x8d, 0xf5, 0x6f, 0x1f, 0xec, 0x93, 0xce,
	0x6f, 0x6b, 0xf4, 0x85, 0x85, 0xa9, 0xbf, 0x5b, 0xed, 0x21, 0xef, 0xbb, 0x35, 0xfa, 0xe2, 0xff,
	0xfb, 0xc2, 0x6b, 0x2f, 0xcf, 0xe9, 0x63, 0xfd, 0xd3, 0xda, 0x7b, 0x8c, 0x4d, 0x4e, 0xcc, 0xe1,
	0x2f, 0xe4, 0x8b, 0x5f, 0xff, 0xfa, 0x6a, 0xed, 0x27, 0xe2, 0x5d, 0x44, 0xdc, 0x2d, 0xac, 0xc9,
	0xe8, 0x62, 0xd1, 0x27, 0xe1, 0x92, 0x0f, 0x97, 0xe2, 0xcb, 0xa8, 0x82, 0xae, 0xf2, 0xa6, 0xc7,
	0xcb, 0xc8, 0xf0, 0xa4, 0x3f, 0xb6, 0x7f, 0x54, 0x2f, 0xdc, 0xe8, 0x62, 0x3c, 0x80, 0x97, 0xf7,
	0xde, 0xf3, 0x8e, 0x57, 0x3f, 0xff, 0xa4, 0xef, 0x5d, 0x23, 0x76, 0x8b, 0x1c, 0x7d, 0xf2, 0xfb,
	0x9f, 0x41, 0xe3, 0xf3, 0x51, 0x40, 0x1e, 0x8e, 0x02, 0xf2, 0xe3, 0x28, 0x20, 0x3f, 0x8f, 0x02,
	0xf2, 0xc7, 0x28, 0x20, 0x7f, 0x8f, 0x82, 0xc6, 0x3f, 0xa3, 0x80, 0x7c, 0x79, 0x15, 0x34, 0x1e,
	0x5c, 0x05, 0x8d, 0xef, 0xaf, 0x02, 0x72, 0x2f, 0xca, 0x74, 0x88, 0xa7, 0x80, 0xa7, 0x52, 0x65,
	0x2e, 0x54, 0x80, 0x67, 0xda, 0xf6, 0xa3, 0xc5, 0xdd, 0x55, 0x74, 0x23, 0xd3, 0xcf, 0x22, 0x44,
	0x65, 0x7a, 0xbd, 0x8d, 0xb2, 0xab, 0xdd, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x11, 0x0a, 0x09,
	0x22, 0x68, 0x06, 0x00, 0x00,
}

func (this *GetStoredApplicationUpRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*GetStoredApplicationUpRequest)
	if !ok {
		that2, ok := that.(GetStoredApplicationUpRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.ApplicationIds.Equal(that1.ApplicationIds) {
		return false
	}
	if !this.EndDeviceIds.Equal(that1.EndDeviceIds) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if !this.Limit.Equal(that1.Limit) {
		return false
	}
	if !this.After.Equal(that1.After) {
		return false
	}
	if !this.Before.Equal(that1.Before) {
		return false
	}
	if !this.FPort.Equal(that1.FPort) {
		return false
	}
	if this.Order != that1.Order {
		return false
	}
	if !this.FieldMask.Equal(that1.FieldMask) {
		return false
	}
	if !this.Last.Equal(that1.Last) {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationUpStorageClient is the client API for ApplicationUpStorage service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationUpStorageClient interface {
	// Returns a stream of application messages that have been stored in the database.
	GetStoredApplicationUp(ctx context.Context, in *GetStoredApplicationUpRequest, opts ...grpc.CallOption) (ApplicationUpStorage_GetStoredApplicationUpClient, error)
}

type applicationUpStorageClient struct {
	cc *grpc.ClientConn
}

func NewApplicationUpStorageClient(cc *grpc.ClientConn) ApplicationUpStorageClient {
	return &applicationUpStorageClient{cc}
}

func (c *applicationUpStorageClient) GetStoredApplicationUp(ctx context.Context, in *GetStoredApplicationUpRequest, opts ...grpc.CallOption) (ApplicationUpStorage_GetStoredApplicationUpClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ApplicationUpStorage_serviceDesc.Streams[0], "/ttn.lorawan.v3.ApplicationUpStorage/GetStoredApplicationUp", opts...)
	if err != nil {
		return nil, err
	}
	x := &applicationUpStorageGetStoredApplicationUpClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApplicationUpStorage_GetStoredApplicationUpClient interface {
	Recv() (*ApplicationUp, error)
	grpc.ClientStream
}

type applicationUpStorageGetStoredApplicationUpClient struct {
	grpc.ClientStream
}

func (x *applicationUpStorageGetStoredApplicationUpClient) Recv() (*ApplicationUp, error) {
	m := new(ApplicationUp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ApplicationUpStorageServer is the server API for ApplicationUpStorage service.
type ApplicationUpStorageServer interface {
	// Returns a stream of application messages that have been stored in the database.
	GetStoredApplicationUp(*GetStoredApplicationUpRequest, ApplicationUpStorage_GetStoredApplicationUpServer) error
}

// UnimplementedApplicationUpStorageServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationUpStorageServer struct {
}

func (*UnimplementedApplicationUpStorageServer) GetStoredApplicationUp(req *GetStoredApplicationUpRequest, srv ApplicationUpStorage_GetStoredApplicationUpServer) error {
	return status.Errorf(codes.Unimplemented, "method GetStoredApplicationUp not implemented")
}

func RegisterApplicationUpStorageServer(s *grpc.Server, srv ApplicationUpStorageServer) {
	s.RegisterService(&_ApplicationUpStorage_serviceDesc, srv)
}

func _ApplicationUpStorage_GetStoredApplicationUp_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetStoredApplicationUpRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApplicationUpStorageServer).GetStoredApplicationUp(m, &applicationUpStorageGetStoredApplicationUpServer{stream})
}

type ApplicationUpStorage_GetStoredApplicationUpServer interface {
	Send(*ApplicationUp) error
	grpc.ServerStream
}

type applicationUpStorageGetStoredApplicationUpServer struct {
	grpc.ServerStream
}

func (x *applicationUpStorageGetStoredApplicationUpServer) Send(m *ApplicationUp) error {
	return x.ServerStream.SendMsg(m)
}

var _ApplicationUpStorage_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ApplicationUpStorage",
	HandlerType: (*ApplicationUpStorageServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetStoredApplicationUp",
			Handler:       _ApplicationUpStorage_GetStoredApplicationUp_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "lorawan-stack/api/applicationserver_integrations_storage.proto",
}

func (m *GetStoredApplicationUpRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetStoredApplicationUpRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetStoredApplicationUpRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Last != nil {
		{
			size, err := m.Last.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApplicationserverIntegrationsStorage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x52
	}
	if m.FieldMask != nil {
		{
			size, err := m.FieldMask.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApplicationserverIntegrationsStorage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Order) > 0 {
		i -= len(m.Order)
		copy(dAtA[i:], m.Order)
		i = encodeVarintApplicationserverIntegrationsStorage(dAtA, i, uint64(len(m.Order)))
		i--
		dAtA[i] = 0x42
	}
	if m.FPort != nil {
		{
			size, err := m.FPort.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApplicationserverIntegrationsStorage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Before != nil {
		{
			size, err := m.Before.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApplicationserverIntegrationsStorage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.After != nil {
		{
			size, err := m.After.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApplicationserverIntegrationsStorage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.Limit != nil {
		{
			size, err := m.Limit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApplicationserverIntegrationsStorage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintApplicationserverIntegrationsStorage(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x1a
	}
	if m.EndDeviceIds != nil {
		{
			size, err := m.EndDeviceIds.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApplicationserverIntegrationsStorage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ApplicationIds != nil {
		{
			size, err := m.ApplicationIds.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintApplicationserverIntegrationsStorage(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApplicationserverIntegrationsStorage(dAtA []byte, offset int, v uint64) int {
	offset -= sovApplicationserverIntegrationsStorage(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedGetStoredApplicationUpRequest(r randyApplicationserverIntegrationsStorage, easy bool) *GetStoredApplicationUpRequest {
	this := &GetStoredApplicationUpRequest{}
	if r.Intn(5) != 0 {
		this.ApplicationIds = NewPopulatedApplicationIdentifiers(r, easy)
	}
	if r.Intn(5) != 0 {
		this.EndDeviceIds = NewPopulatedEndDeviceIdentifiers(r, easy)
	}
	this.Type = string(randStringApplicationserverIntegrationsStorage(r))
	if r.Intn(5) != 0 {
		this.Limit = types.NewPopulatedUInt32Value(r, easy)
	}
	if r.Intn(5) != 0 {
		this.After = types.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Before = types.NewPopulatedTimestamp(r, easy)
	}
	if r.Intn(5) != 0 {
		this.FPort = types.NewPopulatedUInt32Value(r, easy)
	}
	this.Order = string(randStringApplicationserverIntegrationsStorage(r))
	if r.Intn(5) != 0 {
		this.FieldMask = types.NewPopulatedFieldMask(r, easy)
	}
	if r.Intn(5) != 0 {
		this.Last = types.NewPopulatedDuration(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyApplicationserverIntegrationsStorage interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneApplicationserverIntegrationsStorage(r randyApplicationserverIntegrationsStorage) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringApplicationserverIntegrationsStorage(r randyApplicationserverIntegrationsStorage) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneApplicationserverIntegrationsStorage(r)
	}
	return string(tmps)
}
func randUnrecognizedApplicationserverIntegrationsStorage(r randyApplicationserverIntegrationsStorage, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldApplicationserverIntegrationsStorage(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldApplicationserverIntegrationsStorage(dAtA []byte, r randyApplicationserverIntegrationsStorage, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateApplicationserverIntegrationsStorage(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *GetStoredApplicationUpRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ApplicationIds != nil {
		l = m.ApplicationIds.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.EndDeviceIds != nil {
		l = m.EndDeviceIds.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.Limit != nil {
		l = m.Limit.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.After != nil {
		l = m.After.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.Before != nil {
		l = m.Before.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.FPort != nil {
		l = m.FPort.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	l = len(m.Order)
	if l > 0 {
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.FieldMask != nil {
		l = m.FieldMask.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	if m.Last != nil {
		l = m.Last.Size()
		n += 1 + l + sovApplicationserverIntegrationsStorage(uint64(l))
	}
	return n
}

func sovApplicationserverIntegrationsStorage(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApplicationserverIntegrationsStorage(x uint64) (n int) {
	return sovApplicationserverIntegrationsStorage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *GetStoredApplicationUpRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&GetStoredApplicationUpRequest{`,
		`ApplicationIds:` + strings.Replace(fmt.Sprintf("%v", this.ApplicationIds), "ApplicationIdentifiers", "ApplicationIdentifiers", 1) + `,`,
		`EndDeviceIds:` + strings.Replace(fmt.Sprintf("%v", this.EndDeviceIds), "EndDeviceIdentifiers", "EndDeviceIdentifiers", 1) + `,`,
		`Type:` + fmt.Sprintf("%v", this.Type) + `,`,
		`Limit:` + strings.Replace(fmt.Sprintf("%v", this.Limit), "UInt32Value", "types.UInt32Value", 1) + `,`,
		`After:` + strings.Replace(fmt.Sprintf("%v", this.After), "Timestamp", "types.Timestamp", 1) + `,`,
		`Before:` + strings.Replace(fmt.Sprintf("%v", this.Before), "Timestamp", "types.Timestamp", 1) + `,`,
		`FPort:` + strings.Replace(fmt.Sprintf("%v", this.FPort), "UInt32Value", "types.UInt32Value", 1) + `,`,
		`Order:` + fmt.Sprintf("%v", this.Order) + `,`,
		`FieldMask:` + strings.Replace(fmt.Sprintf("%v", this.FieldMask), "FieldMask", "types.FieldMask", 1) + `,`,
		`Last:` + strings.Replace(fmt.Sprintf("%v", this.Last), "Duration", "types.Duration", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringApplicationserverIntegrationsStorage(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *GetStoredApplicationUpRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplicationserverIntegrationsStorage
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
			return fmt.Errorf("proto: GetStoredApplicationUpRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetStoredApplicationUpRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApplicationIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserverIntegrationsStorage
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
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ApplicationIds == nil {
				m.ApplicationIds = &ApplicationIdentifiers{}
			}
			if err := m.ApplicationIds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDeviceIds", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserverIntegrationsStorage
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
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndDeviceIds == nil {
				m.EndDeviceIds = &EndDeviceIdentifiers{}
			}
			if err := m.EndDeviceIds.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserverIntegrationsStorage
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
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserverIntegrationsStorage
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
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Limit == nil {
				m.Limit = &types.UInt32Value{}
			}
			if err := m.Limit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field After", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserverIntegrationsStorage
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
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.After == nil {
				m.After = &types.Timestamp{}
			}
			if err := m.After.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Before", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserverIntegrationsStorage
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
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Before == nil {
				m.Before = &types.Timestamp{}
			}
			if err := m.Before.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FPort", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserverIntegrationsStorage
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
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FPort == nil {
				m.FPort = &types.UInt32Value{}
			}
			if err := m.FPort.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Order", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserverIntegrationsStorage
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
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Order = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldMask", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserverIntegrationsStorage
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
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FieldMask == nil {
				m.FieldMask = &types.FieldMask{}
			}
			if err := m.FieldMask.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Last", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplicationserverIntegrationsStorage
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
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Last == nil {
				m.Last = &types.Duration{}
			}
			if err := m.Last.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplicationserverIntegrationsStorage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApplicationserverIntegrationsStorage
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
func skipApplicationserverIntegrationsStorage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApplicationserverIntegrationsStorage
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
					return 0, ErrIntOverflowApplicationserverIntegrationsStorage
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
					return 0, ErrIntOverflowApplicationserverIntegrationsStorage
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
				return 0, ErrInvalidLengthApplicationserverIntegrationsStorage
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApplicationserverIntegrationsStorage
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApplicationserverIntegrationsStorage
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApplicationserverIntegrationsStorage        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApplicationserverIntegrationsStorage          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApplicationserverIntegrationsStorage = fmt.Errorf("proto: unexpected end of group")
)

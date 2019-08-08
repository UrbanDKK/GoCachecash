// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: metrics.proto

package metrics

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// LogData is the data to be shipped over the wire
type Scrape struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scrape) Reset()         { *m = Scrape{} }
func (m *Scrape) String() string { return proto.CompactTextString(m) }
func (*Scrape) ProtoMessage()    {}
func (*Scrape) Descriptor() ([]byte, []int) {
	return fileDescriptor_6039342a2ba47b72, []int{0}
}
func (m *Scrape) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Scrape) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Scrape.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Scrape) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scrape.Merge(m, src)
}
func (m *Scrape) XXX_Size() int {
	return m.Size()
}
func (m *Scrape) XXX_DiscardUnknown() {
	xxx_messageInfo_Scrape.DiscardUnknown(m)
}

var xxx_messageInfo_Scrape proto.InternalMessageInfo

func (m *Scrape) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Scrape)(nil), "metrics.Scrape")
}

func init() { proto.RegisterFile("metrics.proto", fileDescriptor_6039342a2ba47b72) }

var fileDescriptor_6039342a2ba47b72 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0x4d, 0x2d, 0x29,
	0xca, 0x4c, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xa5, 0xa4, 0xd3,
	0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xc2, 0x49, 0xa5, 0x69, 0xfa, 0xa9, 0xb9, 0x05, 0x25,
	0x95, 0x10, 0x55, 0x4a, 0x32, 0x5c, 0x6c, 0xc1, 0xc9, 0x45, 0x89, 0x05, 0xa9, 0x42, 0x42, 0x5c,
	0x2c, 0x29, 0x89, 0x25, 0x89, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x91, 0x27,
	0x17, 0xbb, 0x2f, 0xc4, 0x14, 0x21, 0x3b, 0x2e, 0x5e, 0x28, 0x33, 0x20, 0x3f, 0x27, 0x27, 0xb5,
	0x48, 0x88, 0x5f, 0x0f, 0x66, 0x1f, 0xc4, 0x00, 0x29, 0x31, 0x3d, 0x88, 0x45, 0x7a, 0x30, 0x8b,
	0xf4, 0x5c, 0x41, 0x16, 0x29, 0x31, 0x68, 0x30, 0x1a, 0x30, 0x3a, 0x09, 0x9c, 0x78, 0x24, 0xc7,
	0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x33, 0x1e, 0xcb, 0x31, 0x24, 0xb1, 0x81,
	0xd5, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdf, 0x38, 0x34, 0x15, 0xb8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsClient is the client API for Metrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsClient interface {
	// MetricsPoller polls for metrics from the datacentre out to clients
	// Clients establish a connection and then send one scrape in protobuf
	// format each time the server sends a poll request
	MetricsPoller(ctx context.Context, opts ...grpc.CallOption) (Metrics_MetricsPollerClient, error)
}

type metricsClient struct {
	cc *grpc.ClientConn
}

func NewMetricsClient(cc *grpc.ClientConn) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) MetricsPoller(ctx context.Context, opts ...grpc.CallOption) (Metrics_MetricsPollerClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Metrics_serviceDesc.Streams[0], "/metrics.Metrics/MetricsPoller", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsMetricsPollerClient{stream}
	return x, nil
}

type Metrics_MetricsPollerClient interface {
	Send(*Scrape) error
	Recv() (*empty.Empty, error)
	grpc.ClientStream
}

type metricsMetricsPollerClient struct {
	grpc.ClientStream
}

func (x *metricsMetricsPollerClient) Send(m *Scrape) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsMetricsPollerClient) Recv() (*empty.Empty, error) {
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MetricsServer is the server API for Metrics service.
type MetricsServer interface {
	// MetricsPoller polls for metrics from the datacentre out to clients
	// Clients establish a connection and then send one scrape in protobuf
	// format each time the server sends a poll request
	MetricsPoller(Metrics_MetricsPollerServer) error
}

// UnimplementedMetricsServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServer struct {
}

func (*UnimplementedMetricsServer) MetricsPoller(srv Metrics_MetricsPollerServer) error {
	return status.Errorf(codes.Unimplemented, "method MetricsPoller not implemented")
}

func RegisterMetricsServer(s *grpc.Server, srv MetricsServer) {
	s.RegisterService(&_Metrics_serviceDesc, srv)
}

func _Metrics_MetricsPoller_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServer).MetricsPoller(&metricsMetricsPollerServer{stream})
}

type Metrics_MetricsPollerServer interface {
	Send(*empty.Empty) error
	Recv() (*Scrape, error)
	grpc.ServerStream
}

type metricsMetricsPollerServer struct {
	grpc.ServerStream
}

func (x *metricsMetricsPollerServer) Send(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsMetricsPollerServer) Recv() (*Scrape, error) {
	m := new(Scrape)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Metrics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "metrics.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "MetricsPoller",
			Handler:       _Metrics_MetricsPoller_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "metrics.proto",
}

func (m *Scrape) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Scrape) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMetrics(dAtA, i, uint64(len(m.Data)))
		i += copy(dAtA[i:], m.Data)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintMetrics(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Scrape) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovMetrics(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMetrics(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMetrics(x uint64) (n int) {
	return sovMetrics(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Scrape) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMetrics
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
			return fmt.Errorf("proto: Scrape: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Scrape: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMetrics
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMetrics
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMetrics(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthMetrics
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMetrics(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMetrics
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
					return 0, ErrIntOverflowMetrics
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMetrics
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
				return 0, ErrInvalidLengthMetrics
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthMetrics
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMetrics
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMetrics(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthMetrics
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMetrics = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMetrics   = fmt.Errorf("proto: integer overflow")
)
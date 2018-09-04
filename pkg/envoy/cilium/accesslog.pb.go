// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cilium/accesslog.proto

package cilium

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HttpProtocol int32

const (
	HttpProtocol_HTTP10 HttpProtocol = 0
	HttpProtocol_HTTP11 HttpProtocol = 1
	HttpProtocol_HTTP2  HttpProtocol = 2
)

var HttpProtocol_name = map[int32]string{
	0: "HTTP10",
	1: "HTTP11",
	2: "HTTP2",
}
var HttpProtocol_value = map[string]int32{
	"HTTP10": 0,
	"HTTP11": 1,
	"HTTP2":  2,
}

func (x HttpProtocol) String() string {
	return proto.EnumName(HttpProtocol_name, int32(x))
}
func (HttpProtocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_accesslog_149966dc1dc4329c, []int{0}
}

type EntryType int32

const (
	EntryType_Request  EntryType = 0
	EntryType_Response EntryType = 1
	EntryType_Denied   EntryType = 2
)

var EntryType_name = map[int32]string{
	0: "Request",
	1: "Response",
	2: "Denied",
}
var EntryType_value = map[string]int32{
	"Request":  0,
	"Response": 1,
	"Denied":   2,
}

func (x EntryType) String() string {
	return proto.EnumName(EntryType_name, int32(x))
}
func (EntryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_accesslog_149966dc1dc4329c, []int{1}
}

type KeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_accesslog_149966dc1dc4329c, []int{0}
}
func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (dst *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(dst, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type HttpLogEntry struct {
	HttpProtocol HttpProtocol `protobuf:"varint,1,opt,name=http_protocol,json=httpProtocol,proto3,enum=cilium.HttpProtocol" json:"http_protocol,omitempty"`
	// Request info that is also retained for the response
	Scheme string `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"`
	Host   string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Path   string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Method string `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	// Request headers not included above
	Headers []*KeyValue `protobuf:"bytes,6,rep,name=headers,proto3" json:"headers,omitempty"`
	// Response info
	Status               uint32   `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HttpLogEntry) Reset()         { *m = HttpLogEntry{} }
func (m *HttpLogEntry) String() string { return proto.CompactTextString(m) }
func (*HttpLogEntry) ProtoMessage()    {}
func (*HttpLogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_accesslog_149966dc1dc4329c, []int{1}
}
func (m *HttpLogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpLogEntry.Unmarshal(m, b)
}
func (m *HttpLogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpLogEntry.Marshal(b, m, deterministic)
}
func (dst *HttpLogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpLogEntry.Merge(dst, src)
}
func (m *HttpLogEntry) XXX_Size() int {
	return xxx_messageInfo_HttpLogEntry.Size(m)
}
func (m *HttpLogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpLogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_HttpLogEntry proto.InternalMessageInfo

func (m *HttpLogEntry) GetHttpProtocol() HttpProtocol {
	if m != nil {
		return m.HttpProtocol
	}
	return HttpProtocol_HTTP10
}

func (m *HttpLogEntry) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *HttpLogEntry) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *HttpLogEntry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *HttpLogEntry) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *HttpLogEntry) GetHeaders() []*KeyValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *HttpLogEntry) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type LogEntry struct {
	// The time that Cilium filter captured this log entry,
	// in, nanoseconds since 1/1/1970.
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 'true' if the request was received by an ingress listener,
	// 'false' if received by an egress listener
	IsIngress bool      `protobuf:"varint,15,opt,name=is_ingress,json=isIngress,proto3" json:"is_ingress,omitempty"`
	EntryType EntryType `protobuf:"varint,3,opt,name=entry_type,json=entryType,proto3,enum=cilium.EntryType" json:"entry_type,omitempty"`
	// Cilium network policy resource name
	PolicyName string `protobuf:"bytes,4,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// Cilium rule reference
	CiliumRuleRef string `protobuf:"bytes,5,opt,name=cilium_rule_ref,json=ciliumRuleRef,proto3" json:"cilium_rule_ref,omitempty"`
	// Cilium security ID of the source and destination
	SourceSecurityId      uint32 `protobuf:"varint,6,opt,name=source_security_id,json=sourceSecurityId,proto3" json:"source_security_id,omitempty"`
	DestinationSecurityId uint32 `protobuf:"varint,16,opt,name=destination_security_id,json=destinationSecurityId,proto3" json:"destination_security_id,omitempty"`
	// These fields record the original source and destination addresses,
	// stored in ipv4:port or [ipv6]:port format.
	SourceAddress      string `protobuf:"bytes,7,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress string `protobuf:"bytes,8,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	// Types that are valid to be assigned to L7:
	//	*LogEntry_Http
	L7 isLogEntry_L7 `protobuf_oneof:"l7"`
	//
	// Deprecated HTTP fields. Use the http field above instead.
	//
	HttpProtocol HttpProtocol `protobuf:"varint,2,opt,name=http_protocol,json=httpProtocol,proto3,enum=cilium.HttpProtocol" json:"http_protocol,omitempty"` // Deprecated: Do not use.
	// Request info that is also retained for the response
	Scheme string `protobuf:"bytes,9,opt,name=scheme,proto3" json:"scheme,omitempty"`   // Deprecated: Do not use.
	Host   string `protobuf:"bytes,10,opt,name=host,proto3" json:"host,omitempty"`      // Deprecated: Do not use.
	Path   string `protobuf:"bytes,11,opt,name=path,proto3" json:"path,omitempty"`      // Deprecated: Do not use.
	Method string `protobuf:"bytes,12,opt,name=method,proto3" json:"method,omitempty"`  // Deprecated: Do not use.
	Status uint32 `protobuf:"varint,13,opt,name=status,proto3" json:"status,omitempty"` // Deprecated: Do not use.
	// Request headers not included above
	Headers              []*KeyValue `protobuf:"bytes,14,rep,name=headers,proto3" json:"headers,omitempty"` // Deprecated: Do not use.
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LogEntry) Reset()         { *m = LogEntry{} }
func (m *LogEntry) String() string { return proto.CompactTextString(m) }
func (*LogEntry) ProtoMessage()    {}
func (*LogEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_accesslog_149966dc1dc4329c, []int{2}
}
func (m *LogEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogEntry.Unmarshal(m, b)
}
func (m *LogEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogEntry.Marshal(b, m, deterministic)
}
func (dst *LogEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogEntry.Merge(dst, src)
}
func (m *LogEntry) XXX_Size() int {
	return xxx_messageInfo_LogEntry.Size(m)
}
func (m *LogEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_LogEntry.DiscardUnknown(m)
}

var xxx_messageInfo_LogEntry proto.InternalMessageInfo

func (m *LogEntry) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *LogEntry) GetIsIngress() bool {
	if m != nil {
		return m.IsIngress
	}
	return false
}

func (m *LogEntry) GetEntryType() EntryType {
	if m != nil {
		return m.EntryType
	}
	return EntryType_Request
}

func (m *LogEntry) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *LogEntry) GetCiliumRuleRef() string {
	if m != nil {
		return m.CiliumRuleRef
	}
	return ""
}

func (m *LogEntry) GetSourceSecurityId() uint32 {
	if m != nil {
		return m.SourceSecurityId
	}
	return 0
}

func (m *LogEntry) GetDestinationSecurityId() uint32 {
	if m != nil {
		return m.DestinationSecurityId
	}
	return 0
}

func (m *LogEntry) GetSourceAddress() string {
	if m != nil {
		return m.SourceAddress
	}
	return ""
}

func (m *LogEntry) GetDestinationAddress() string {
	if m != nil {
		return m.DestinationAddress
	}
	return ""
}

type isLogEntry_L7 interface {
	isLogEntry_L7()
}

type LogEntry_Http struct {
	Http *HttpLogEntry `protobuf:"bytes,100,opt,name=http,proto3,oneof"`
}

func (*LogEntry_Http) isLogEntry_L7() {}

func (m *LogEntry) GetL7() isLogEntry_L7 {
	if m != nil {
		return m.L7
	}
	return nil
}

func (m *LogEntry) GetHttp() *HttpLogEntry {
	if x, ok := m.GetL7().(*LogEntry_Http); ok {
		return x.Http
	}
	return nil
}

// Deprecated: Do not use.
func (m *LogEntry) GetHttpProtocol() HttpProtocol {
	if m != nil {
		return m.HttpProtocol
	}
	return HttpProtocol_HTTP10
}

// Deprecated: Do not use.
func (m *LogEntry) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

// Deprecated: Do not use.
func (m *LogEntry) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

// Deprecated: Do not use.
func (m *LogEntry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

// Deprecated: Do not use.
func (m *LogEntry) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

// Deprecated: Do not use.
func (m *LogEntry) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

// Deprecated: Do not use.
func (m *LogEntry) GetHeaders() []*KeyValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LogEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LogEntry_OneofMarshaler, _LogEntry_OneofUnmarshaler, _LogEntry_OneofSizer, []interface{}{
		(*LogEntry_Http)(nil),
	}
}

func _LogEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LogEntry)
	// l7
	switch x := m.L7.(type) {
	case *LogEntry_Http:
		b.EncodeVarint(100<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Http); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LogEntry.L7 has unexpected type %T", x)
	}
	return nil
}

func _LogEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LogEntry)
	switch tag {
	case 100: // l7.http
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HttpLogEntry)
		err := b.DecodeMessage(msg)
		m.L7 = &LogEntry_Http{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LogEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LogEntry)
	// l7
	switch x := m.L7.(type) {
	case *LogEntry_Http:
		s := proto.Size(x.Http)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*KeyValue)(nil), "cilium.KeyValue")
	proto.RegisterType((*HttpLogEntry)(nil), "cilium.HttpLogEntry")
	proto.RegisterType((*LogEntry)(nil), "cilium.LogEntry")
	proto.RegisterEnum("cilium.HttpProtocol", HttpProtocol_name, HttpProtocol_value)
	proto.RegisterEnum("cilium.EntryType", EntryType_name, EntryType_value)
}

func init() { proto.RegisterFile("cilium/accesslog.proto", fileDescriptor_accesslog_149966dc1dc4329c) }

var fileDescriptor_accesslog_149966dc1dc4329c = []byte{
	// 578 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x6d, 0x8b, 0xd3, 0x40,
	0x10, 0xc7, 0x2f, 0xb9, 0x36, 0x6d, 0xa6, 0x0f, 0x17, 0xd7, 0xb3, 0x86, 0x43, 0xb1, 0x1c, 0x28,
	0xa5, 0x48, 0xef, 0x1a, 0x41, 0xf1, 0x85, 0x2f, 0x3c, 0x14, 0x7a, 0x28, 0x72, 0xac, 0xc5, 0xb7,
	0x21, 0x26, 0x73, 0xcd, 0x62, 0x9e, 0xcc, 0x6e, 0x84, 0x7c, 0x20, 0x3f, 0x9a, 0xdf, 0x43, 0x76,
	0x37, 0x49, 0x7b, 0x87, 0xbe, 0x9b, 0xf9, 0xcf, 0xc3, 0xce, 0x4c, 0x7e, 0x81, 0x59, 0xc8, 0x12,
	0x56, 0xa5, 0x17, 0x41, 0x18, 0x22, 0xe7, 0x49, 0xbe, 0x5b, 0x15, 0x65, 0x2e, 0x72, 0x62, 0x69,
	0xfd, 0xdc, 0x83, 0xe1, 0x27, 0xac, 0xbf, 0x05, 0x49, 0x85, 0xc4, 0x81, 0xe3, 0x1f, 0x58, 0xbb,
	0xc6, 0xdc, 0x58, 0xd8, 0x54, 0x9a, 0xe4, 0x14, 0xfa, 0xbf, 0x64, 0xc8, 0x35, 0x95, 0xa6, 0x9d,
	0xf3, 0x3f, 0x06, 0x8c, 0x37, 0x42, 0x14, 0x9f, 0xf3, 0xdd, 0xc7, 0x4c, 0x94, 0x35, 0x79, 0x0b,
	0x93, 0x58, 0x88, 0xc2, 0x57, 0xad, 0xc3, 0x3c, 0x51, 0x2d, 0xa6, 0xde, 0xe9, 0x4a, 0x3f, 0xb2,
	0x92, 0xc9, 0x37, 0x4d, 0x8c, 0x8e, 0xe3, 0x03, 0x8f, 0xcc, 0xc0, 0xe2, 0x61, 0x8c, 0x69, 0xfb,
	0x44, 0xe3, 0x11, 0x02, 0xbd, 0x38, 0xe7, 0xc2, 0x3d, 0x56, 0xaa, 0xb2, 0xa5, 0x56, 0x04, 0x22,
	0x76, 0x7b, 0x5a, 0x93, 0xb6, 0xac, 0x4f, 0x51, 0xc4, 0x79, 0xe4, 0xf6, 0x75, 0xbd, 0xf6, 0xc8,
	0x12, 0x06, 0x31, 0x06, 0x11, 0x96, 0xdc, 0xb5, 0xe6, 0xc7, 0x8b, 0x91, 0xe7, 0xb4, 0xc3, 0xb4,
	0xeb, 0xd2, 0x36, 0x41, 0xcd, 0x20, 0x02, 0x51, 0x71, 0x77, 0x30, 0x37, 0x16, 0x13, 0xda, 0x78,
	0xe7, 0xbf, 0xfb, 0x30, 0xec, 0x76, 0x7c, 0x02, 0xb6, 0x60, 0x29, 0x72, 0x11, 0xa4, 0x85, 0xda,
	0xaf, 0x47, 0xf7, 0x02, 0x79, 0x0a, 0xc0, 0xb8, 0xcf, 0xb2, 0x5d, 0x89, 0x9c, 0xbb, 0x27, 0x73,
	0x63, 0x31, 0xa4, 0x36, 0xe3, 0xd7, 0x5a, 0x20, 0x97, 0x00, 0x28, 0xbb, 0xf8, 0xa2, 0x2e, 0x50,
	0xed, 0x34, 0xf5, 0x1e, 0xb4, 0x03, 0xa9, 0xfe, 0xdb, 0xba, 0x40, 0x6a, 0x63, 0x6b, 0x92, 0x67,
	0x30, 0x2a, 0xf2, 0x84, 0x85, 0xb5, 0x9f, 0x05, 0x29, 0x36, 0x2b, 0x83, 0x96, 0xbe, 0x04, 0x29,
	0x92, 0x17, 0x70, 0xa2, 0xeb, 0xfd, 0xb2, 0x4a, 0xd0, 0x2f, 0xf1, 0xb6, 0xb9, 0xc0, 0x44, 0xcb,
	0xb4, 0x4a, 0x90, 0xe2, 0x2d, 0x79, 0x09, 0x84, 0xe7, 0x55, 0x19, 0xa2, 0xcf, 0x31, 0xac, 0x4a,
	0x26, 0x6a, 0x9f, 0x45, 0xae, 0xa5, 0x16, 0x75, 0x74, 0xe4, 0x6b, 0x13, 0xb8, 0x8e, 0xc8, 0x6b,
	0x78, 0x1c, 0x21, 0x17, 0x2c, 0x0b, 0x04, 0xcb, 0xb3, 0x3b, 0x25, 0x8e, 0x2a, 0x79, 0x74, 0x10,
	0x3e, 0xa8, 0x7b, 0x0e, 0xd3, 0xe6, 0x95, 0x20, 0x8a, 0xd4, 0x0d, 0x06, 0x7a, 0x18, 0xad, 0xbe,
	0xd7, 0x22, 0xb9, 0x80, 0x87, 0x87, 0xed, 0xdb, 0xdc, 0xa1, 0xca, 0x25, 0x07, 0xa1, 0xb6, 0x60,
	0x09, 0x3d, 0x89, 0x8b, 0x1b, 0xcd, 0x8d, 0xc5, 0xe8, 0x2e, 0x50, 0xed, 0x97, 0xd9, 0x1c, 0x51,
	0x95, 0x43, 0xde, 0xdd, 0xa7, 0xd0, 0xfc, 0x3f, 0x85, 0x57, 0xa6, 0x6b, 0xdc, 0x23, 0xf1, 0xac,
	0x23, 0xd1, 0x96, 0xe3, 0xa8, 0x8c, 0x96, 0xc6, 0x59, 0x43, 0x23, 0x74, 0x11, 0x4d, 0xe4, 0xac,
	0x21, 0x72, 0xb4, 0xd7, 0x15, 0x95, 0x67, 0x1d, 0x95, 0xe3, 0x7d, 0xaf, 0x86, 0xcc, 0xb3, 0x8e,
	0xb6, 0x89, 0xbc, 0x68, 0xf3, 0x8e, 0x52, 0xc8, 0x6a, 0x4f, 0xed, 0xf4, 0xdf, 0xd4, 0xaa, 0xf4,
	0x36, 0xe9, 0xaa, 0x07, 0x66, 0xf2, 0x66, 0xb9, 0xd6, 0xbf, 0x63, 0xb7, 0x09, 0x80, 0xb5, 0xd9,
	0x6e, 0x6f, 0xd6, 0x97, 0xce, 0x51, 0x67, 0xaf, 0x1d, 0x83, 0xd8, 0xd0, 0x97, 0xb6, 0xe7, 0x98,
	0x4b, 0x0f, 0xec, 0x0e, 0x3b, 0x32, 0x82, 0x01, 0xc5, 0x9f, 0x15, 0x72, 0xe1, 0x1c, 0x91, 0x31,
	0x0c, 0x29, 0xf2, 0x22, 0xcf, 0x38, 0x3a, 0x86, 0x2c, 0xff, 0x80, 0x19, 0xc3, 0xc8, 0x31, 0xbf,
	0x5b, 0xea, 0xb0, 0xaf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x11, 0x8d, 0x6f, 0xe8, 0x53, 0x04,
	0x00, 0x00,
}

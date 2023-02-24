package policy

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/imannamdari/v2ray-core/v5/common/protoext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Second struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint32 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Second) Reset() {
	*x = Second{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_policy_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Second) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Second) ProtoMessage() {}

func (x *Second) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Second.ProtoReflect.Descriptor instead.
func (*Second) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{0}
}

func (x *Second) GetValue() uint32 {
	if x != nil {
		return x.Value
	}
	return 0
}

type Policy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timeout *Policy_Timeout `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	Stats   *Policy_Stats   `protobuf:"bytes,2,opt,name=stats,proto3" json:"stats,omitempty"`
	Buffer  *Policy_Buffer  `protobuf:"bytes,3,opt,name=buffer,proto3" json:"buffer,omitempty"`
}

func (x *Policy) Reset() {
	*x = Policy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_policy_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy) ProtoMessage() {}

func (x *Policy) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy.ProtoReflect.Descriptor instead.
func (*Policy) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{1}
}

func (x *Policy) GetTimeout() *Policy_Timeout {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *Policy) GetStats() *Policy_Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *Policy) GetBuffer() *Policy_Buffer {
	if x != nil {
		return x.Buffer
	}
	return nil
}

type SystemPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stats *SystemPolicy_Stats `protobuf:"bytes,1,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (x *SystemPolicy) Reset() {
	*x = SystemPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_policy_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemPolicy) ProtoMessage() {}

func (x *SystemPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemPolicy.ProtoReflect.Descriptor instead.
func (*SystemPolicy) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{2}
}

func (x *SystemPolicy) GetStats() *SystemPolicy_Stats {
	if x != nil {
		return x.Stats
	}
	return nil
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Level  map[uint32]*Policy `protobuf:"bytes,1,rep,name=level,proto3" json:"level,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	System *SystemPolicy      `protobuf:"bytes,2,opt,name=system,proto3" json:"system,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_policy_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{3}
}

func (x *Config) GetLevel() map[uint32]*Policy {
	if x != nil {
		return x.Level
	}
	return nil
}

func (x *Config) GetSystem() *SystemPolicy {
	if x != nil {
		return x.System
	}
	return nil
}

// Timeout is a message for timeout settings in various stages, in seconds.
type Policy_Timeout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Handshake      *Second `protobuf:"bytes,1,opt,name=handshake,proto3" json:"handshake,omitempty"`
	ConnectionIdle *Second `protobuf:"bytes,2,opt,name=connection_idle,json=connectionIdle,proto3" json:"connection_idle,omitempty"`
	UplinkOnly     *Second `protobuf:"bytes,3,opt,name=uplink_only,json=uplinkOnly,proto3" json:"uplink_only,omitempty"`
	DownlinkOnly   *Second `protobuf:"bytes,4,opt,name=downlink_only,json=downlinkOnly,proto3" json:"downlink_only,omitempty"`
}

func (x *Policy_Timeout) Reset() {
	*x = Policy_Timeout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_policy_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy_Timeout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy_Timeout) ProtoMessage() {}

func (x *Policy_Timeout) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy_Timeout.ProtoReflect.Descriptor instead.
func (*Policy_Timeout) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Policy_Timeout) GetHandshake() *Second {
	if x != nil {
		return x.Handshake
	}
	return nil
}

func (x *Policy_Timeout) GetConnectionIdle() *Second {
	if x != nil {
		return x.ConnectionIdle
	}
	return nil
}

func (x *Policy_Timeout) GetUplinkOnly() *Second {
	if x != nil {
		return x.UplinkOnly
	}
	return nil
}

func (x *Policy_Timeout) GetDownlinkOnly() *Second {
	if x != nil {
		return x.DownlinkOnly
	}
	return nil
}

type Policy_Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserUplink   bool `protobuf:"varint,1,opt,name=user_uplink,json=userUplink,proto3" json:"user_uplink,omitempty"`
	UserDownlink bool `protobuf:"varint,2,opt,name=user_downlink,json=userDownlink,proto3" json:"user_downlink,omitempty"`
}

func (x *Policy_Stats) Reset() {
	*x = Policy_Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_policy_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy_Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy_Stats) ProtoMessage() {}

func (x *Policy_Stats) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy_Stats.ProtoReflect.Descriptor instead.
func (*Policy_Stats) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Policy_Stats) GetUserUplink() bool {
	if x != nil {
		return x.UserUplink
	}
	return false
}

func (x *Policy_Stats) GetUserDownlink() bool {
	if x != nil {
		return x.UserDownlink
	}
	return false
}

type Policy_Buffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Buffer size per connection, in bytes. -1 for unlimited buffer.
	Connection int32 `protobuf:"varint,1,opt,name=connection,proto3" json:"connection,omitempty"`
}

func (x *Policy_Buffer) Reset() {
	*x = Policy_Buffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_policy_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Policy_Buffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Policy_Buffer) ProtoMessage() {}

func (x *Policy_Buffer) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Policy_Buffer.ProtoReflect.Descriptor instead.
func (*Policy_Buffer) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{1, 2}
}

func (x *Policy_Buffer) GetConnection() int32 {
	if x != nil {
		return x.Connection
	}
	return 0
}

type SystemPolicy_Stats struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InboundUplink    bool `protobuf:"varint,1,opt,name=inbound_uplink,json=inboundUplink,proto3" json:"inbound_uplink,omitempty"`
	InboundDownlink  bool `protobuf:"varint,2,opt,name=inbound_downlink,json=inboundDownlink,proto3" json:"inbound_downlink,omitempty"`
	OutboundUplink   bool `protobuf:"varint,3,opt,name=outbound_uplink,json=outboundUplink,proto3" json:"outbound_uplink,omitempty"`
	OutboundDownlink bool `protobuf:"varint,4,opt,name=outbound_downlink,json=outboundDownlink,proto3" json:"outbound_downlink,omitempty"`
}

func (x *SystemPolicy_Stats) Reset() {
	*x = SystemPolicy_Stats{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_policy_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemPolicy_Stats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemPolicy_Stats) ProtoMessage() {}

func (x *SystemPolicy_Stats) ProtoReflect() protoreflect.Message {
	mi := &file_app_policy_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemPolicy_Stats.ProtoReflect.Descriptor instead.
func (*SystemPolicy_Stats) Descriptor() ([]byte, []int) {
	return file_app_policy_config_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SystemPolicy_Stats) GetInboundUplink() bool {
	if x != nil {
		return x.InboundUplink
	}
	return false
}

func (x *SystemPolicy_Stats) GetInboundDownlink() bool {
	if x != nil {
		return x.InboundDownlink
	}
	return false
}

func (x *SystemPolicy_Stats) GetOutboundUplink() bool {
	if x != nil {
		return x.OutboundUplink
	}
	return false
}

func (x *SystemPolicy_Stats) GetOutboundDownlink() bool {
	if x != nil {
		return x.OutboundDownlink
	}
	return false
}

var File_app_policy_config_proto protoreflect.FileDescriptor

var file_app_policy_config_proto_rawDesc = []byte{
	0x0a, 0x17, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x76, 0x32, 0x72, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78,
	0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x06, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xd0, 0x04, 0x0a, 0x06, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3f, 0x0a,
	0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x39,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x06, 0x62, 0x75, 0x66,
	0x66, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x42, 0x75, 0x66, 0x66, 0x65, 0x72, 0x52,
	0x06, 0x62, 0x75, 0x66, 0x66, 0x65, 0x72, 0x1a, 0x92, 0x02, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x12, 0x3b, 0x0a, 0x09, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x09, 0x68, 0x61, 0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65,
	0x12, 0x46, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x6c, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x75, 0x70, 0x6c, 0x69,
	0x6e, 0x6b, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x0a, 0x75, 0x70,
	0x6c, 0x69, 0x6e, 0x6b, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x42, 0x0a, 0x0d, 0x64, 0x6f, 0x77, 0x6e,
	0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x52, 0x0c,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x4f, 0x6e, 0x6c, 0x79, 0x1a, 0x4d, 0x0a, 0x05,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x75, 0x70,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64,
	0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x75,
	0x73, 0x65, 0x72, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x1a, 0x28, 0x0a, 0x06, 0x42,
	0x75, 0x66, 0x66, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x81, 0x02, 0x0a, 0x0c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x3f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x1a, 0xaf, 0x01, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x75, 0x70, 0x6c,
	0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x6e, 0x62, 0x6f, 0x75,
	0x6e, 0x64, 0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x29, 0x0a, 0x10, 0x69, 0x6e, 0x62, 0x6f,
	0x75, 0x6e, 0x64, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x69, 0x6e, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x69, 0x6e, 0x6b, 0x12, 0x27, 0x0a, 0x0f, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x75, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6f, 0x75,
	0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x55, 0x70, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x2b, 0x0a, 0x11,
	0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0xf9, 0x01, 0x0a, 0x06, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x3e, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x06, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x1a, 0x57, 0x0a, 0x0a, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x33, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x19, 0x82, 0xb5, 0x18, 0x09,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x82, 0xb5, 0x18, 0x08, 0x12, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x60, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x35, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0xaa,
	0x02, 0x15, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70,
	0x2e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_policy_config_proto_rawDescOnce sync.Once
	file_app_policy_config_proto_rawDescData = file_app_policy_config_proto_rawDesc
)

func file_app_policy_config_proto_rawDescGZIP() []byte {
	file_app_policy_config_proto_rawDescOnce.Do(func() {
		file_app_policy_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_policy_config_proto_rawDescData)
	})
	return file_app_policy_config_proto_rawDescData
}

var file_app_policy_config_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_app_policy_config_proto_goTypes = []interface{}{
	(*Second)(nil),             // 0: v2ray.core.app.policy.Second
	(*Policy)(nil),             // 1: v2ray.core.app.policy.Policy
	(*SystemPolicy)(nil),       // 2: v2ray.core.app.policy.SystemPolicy
	(*Config)(nil),             // 3: v2ray.core.app.policy.Config
	(*Policy_Timeout)(nil),     // 4: v2ray.core.app.policy.Policy.Timeout
	(*Policy_Stats)(nil),       // 5: v2ray.core.app.policy.Policy.Stats
	(*Policy_Buffer)(nil),      // 6: v2ray.core.app.policy.Policy.Buffer
	(*SystemPolicy_Stats)(nil), // 7: v2ray.core.app.policy.SystemPolicy.Stats
	nil,                        // 8: v2ray.core.app.policy.Config.LevelEntry
}
var file_app_policy_config_proto_depIdxs = []int32{
	4,  // 0: v2ray.core.app.policy.Policy.timeout:type_name -> v2ray.core.app.policy.Policy.Timeout
	5,  // 1: v2ray.core.app.policy.Policy.stats:type_name -> v2ray.core.app.policy.Policy.Stats
	6,  // 2: v2ray.core.app.policy.Policy.buffer:type_name -> v2ray.core.app.policy.Policy.Buffer
	7,  // 3: v2ray.core.app.policy.SystemPolicy.stats:type_name -> v2ray.core.app.policy.SystemPolicy.Stats
	8,  // 4: v2ray.core.app.policy.Config.level:type_name -> v2ray.core.app.policy.Config.LevelEntry
	2,  // 5: v2ray.core.app.policy.Config.system:type_name -> v2ray.core.app.policy.SystemPolicy
	0,  // 6: v2ray.core.app.policy.Policy.Timeout.handshake:type_name -> v2ray.core.app.policy.Second
	0,  // 7: v2ray.core.app.policy.Policy.Timeout.connection_idle:type_name -> v2ray.core.app.policy.Second
	0,  // 8: v2ray.core.app.policy.Policy.Timeout.uplink_only:type_name -> v2ray.core.app.policy.Second
	0,  // 9: v2ray.core.app.policy.Policy.Timeout.downlink_only:type_name -> v2ray.core.app.policy.Second
	1,  // 10: v2ray.core.app.policy.Config.LevelEntry.value:type_name -> v2ray.core.app.policy.Policy
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_app_policy_config_proto_init() }
func file_app_policy_config_proto_init() {
	if File_app_policy_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_policy_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Second); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_policy_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_policy_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemPolicy); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_policy_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Config); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_policy_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy_Timeout); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_policy_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy_Stats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_policy_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Policy_Buffer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_app_policy_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemPolicy_Stats); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_policy_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_policy_config_proto_goTypes,
		DependencyIndexes: file_app_policy_config_proto_depIdxs,
		MessageInfos:      file_app_policy_config_proto_msgTypes,
	}.Build()
	File_app_policy_config_proto = out.File
	file_app_policy_config_proto_rawDesc = nil
	file_app_policy_config_proto_goTypes = nil
	file_app_policy_config_proto_depIdxs = nil
}

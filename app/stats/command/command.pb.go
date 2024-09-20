package command

import (
	_ "github.com/imannamdari/v2ray-core/v5/common/protoext"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the stat counter.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Whether or not to reset the counter to fetching its value.
	Reset_ bool `protobuf:"varint,2,opt,name=reset,proto3" json:"reset,omitempty"`
}

func (x *GetStatsRequest) Reset() {
	*x = GetStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_stats_command_command_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsRequest) ProtoMessage() {}

func (x *GetStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_stats_command_command_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsRequest.ProtoReflect.Descriptor instead.
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return file_app_stats_command_command_proto_rawDescGZIP(), []int{0}
}

func (x *GetStatsRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetStatsRequest) GetReset_() bool {
	if x != nil {
		return x.Reset_
	}
	return false
}

type Stat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value int64  `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Stat) Reset() {
	*x = Stat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_stats_command_command_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stat) ProtoMessage() {}

func (x *Stat) ProtoReflect() protoreflect.Message {
	mi := &file_app_stats_command_command_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stat.ProtoReflect.Descriptor instead.
func (*Stat) Descriptor() ([]byte, []int) {
	return file_app_stats_command_command_proto_rawDescGZIP(), []int{1}
}

func (x *Stat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Stat) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type GetStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stat *Stat `protobuf:"bytes,1,opt,name=stat,proto3" json:"stat,omitempty"`
}

func (x *GetStatsResponse) Reset() {
	*x = GetStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_stats_command_command_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatsResponse) ProtoMessage() {}

func (x *GetStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_stats_command_command_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatsResponse.ProtoReflect.Descriptor instead.
func (*GetStatsResponse) Descriptor() ([]byte, []int) {
	return file_app_stats_command_command_proto_rawDescGZIP(), []int{2}
}

func (x *GetStatsResponse) GetStat() *Stat {
	if x != nil {
		return x.Stat
	}
	return nil
}

type QueryStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Deprecated, use Patterns instead
	Pattern  string   `protobuf:"bytes,1,opt,name=pattern,proto3" json:"pattern,omitempty"`
	Reset_   bool     `protobuf:"varint,2,opt,name=reset,proto3" json:"reset,omitempty"`
	Patterns []string `protobuf:"bytes,3,rep,name=patterns,proto3" json:"patterns,omitempty"`
	Regexp   bool     `protobuf:"varint,4,opt,name=regexp,proto3" json:"regexp,omitempty"`
}

func (x *QueryStatsRequest) Reset() {
	*x = QueryStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_stats_command_command_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryStatsRequest) ProtoMessage() {}

func (x *QueryStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_stats_command_command_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryStatsRequest.ProtoReflect.Descriptor instead.
func (*QueryStatsRequest) Descriptor() ([]byte, []int) {
	return file_app_stats_command_command_proto_rawDescGZIP(), []int{3}
}

func (x *QueryStatsRequest) GetPattern() string {
	if x != nil {
		return x.Pattern
	}
	return ""
}

func (x *QueryStatsRequest) GetReset_() bool {
	if x != nil {
		return x.Reset_
	}
	return false
}

func (x *QueryStatsRequest) GetPatterns() []string {
	if x != nil {
		return x.Patterns
	}
	return nil
}

func (x *QueryStatsRequest) GetRegexp() bool {
	if x != nil {
		return x.Regexp
	}
	return false
}

type QueryStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stat []*Stat `protobuf:"bytes,1,rep,name=stat,proto3" json:"stat,omitempty"`
}

func (x *QueryStatsResponse) Reset() {
	*x = QueryStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_stats_command_command_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryStatsResponse) ProtoMessage() {}

func (x *QueryStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_stats_command_command_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryStatsResponse.ProtoReflect.Descriptor instead.
func (*QueryStatsResponse) Descriptor() ([]byte, []int) {
	return file_app_stats_command_command_proto_rawDescGZIP(), []int{4}
}

func (x *QueryStatsResponse) GetStat() []*Stat {
	if x != nil {
		return x.Stat
	}
	return nil
}

type SysStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SysStatsRequest) Reset() {
	*x = SysStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_stats_command_command_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysStatsRequest) ProtoMessage() {}

func (x *SysStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_app_stats_command_command_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysStatsRequest.ProtoReflect.Descriptor instead.
func (*SysStatsRequest) Descriptor() ([]byte, []int) {
	return file_app_stats_command_command_proto_rawDescGZIP(), []int{5}
}

type SysStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NumGoroutine uint32 `protobuf:"varint,1,opt,name=NumGoroutine,proto3" json:"NumGoroutine,omitempty"`
	NumGC        uint32 `protobuf:"varint,2,opt,name=NumGC,proto3" json:"NumGC,omitempty"`
	Alloc        uint64 `protobuf:"varint,3,opt,name=Alloc,proto3" json:"Alloc,omitempty"`
	TotalAlloc   uint64 `protobuf:"varint,4,opt,name=TotalAlloc,proto3" json:"TotalAlloc,omitempty"`
	Sys          uint64 `protobuf:"varint,5,opt,name=Sys,proto3" json:"Sys,omitempty"`
	Mallocs      uint64 `protobuf:"varint,6,opt,name=Mallocs,proto3" json:"Mallocs,omitempty"`
	Frees        uint64 `protobuf:"varint,7,opt,name=Frees,proto3" json:"Frees,omitempty"`
	LiveObjects  uint64 `protobuf:"varint,8,opt,name=LiveObjects,proto3" json:"LiveObjects,omitempty"`
	PauseTotalNs uint64 `protobuf:"varint,9,opt,name=PauseTotalNs,proto3" json:"PauseTotalNs,omitempty"`
	Uptime       uint32 `protobuf:"varint,10,opt,name=Uptime,proto3" json:"Uptime,omitempty"`
}

func (x *SysStatsResponse) Reset() {
	*x = SysStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_stats_command_command_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysStatsResponse) ProtoMessage() {}

func (x *SysStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_app_stats_command_command_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysStatsResponse.ProtoReflect.Descriptor instead.
func (*SysStatsResponse) Descriptor() ([]byte, []int) {
	return file_app_stats_command_command_proto_rawDescGZIP(), []int{6}
}

func (x *SysStatsResponse) GetNumGoroutine() uint32 {
	if x != nil {
		return x.NumGoroutine
	}
	return 0
}

func (x *SysStatsResponse) GetNumGC() uint32 {
	if x != nil {
		return x.NumGC
	}
	return 0
}

func (x *SysStatsResponse) GetAlloc() uint64 {
	if x != nil {
		return x.Alloc
	}
	return 0
}

func (x *SysStatsResponse) GetTotalAlloc() uint64 {
	if x != nil {
		return x.TotalAlloc
	}
	return 0
}

func (x *SysStatsResponse) GetSys() uint64 {
	if x != nil {
		return x.Sys
	}
	return 0
}

func (x *SysStatsResponse) GetMallocs() uint64 {
	if x != nil {
		return x.Mallocs
	}
	return 0
}

func (x *SysStatsResponse) GetFrees() uint64 {
	if x != nil {
		return x.Frees
	}
	return 0
}

func (x *SysStatsResponse) GetLiveObjects() uint64 {
	if x != nil {
		return x.LiveObjects
	}
	return 0
}

func (x *SysStatsResponse) GetPauseTotalNs() uint64 {
	if x != nil {
		return x.PauseTotalNs
	}
	return 0
}

func (x *SysStatsResponse) GetUptime() uint32 {
	if x != nil {
		return x.Uptime
	}
	return 0
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_stats_command_command_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_stats_command_command_proto_msgTypes[7]
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
	return file_app_stats_command_command_proto_rawDescGZIP(), []int{7}
}

var File_app_stats_command_command_proto protoreflect.FileDescriptor

var file_app_stats_command_command_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x1a,
	0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x74,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x3b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x73, 0x65, 0x74, 0x22, 0x30,
	0x0a, 0x04, 0x53, 0x74, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x4a, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x73, 0x74, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x04, 0x73, 0x74, 0x61, 0x74, 0x22, 0x77, 0x0a, 0x11,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x72,
	0x65, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x65, 0x78, 0x70, 0x22, 0x4c, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x73,
	0x74, 0x61, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x32, 0x72, 0x61,
	0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x52, 0x04, 0x73,
	0x74, 0x61, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x53, 0x79, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xa2, 0x02, 0x0a, 0x10, 0x53, 0x79, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x4e,
	0x75, 0x6d, 0x47, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0c, 0x4e, 0x75, 0x6d, 0x47, 0x6f, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x4e, 0x75, 0x6d, 0x47, 0x43, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05,
	0x4e, 0x75, 0x6d, 0x47, 0x43, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x54,
	0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x53,
	0x79, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x53, 0x79, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x4d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x4d, 0x61, 0x6c, 0x6c, 0x6f, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x72, 0x65, 0x65, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x46, 0x72, 0x65, 0x65, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0b, 0x4c, 0x69, 0x76, 0x65, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x50, 0x61, 0x75, 0x73, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x4e, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x50, 0x61, 0x75, 0x73, 0x65, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x4e, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x55, 0x70, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x22, 0x0a, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x3a, 0x18, 0x82, 0xb5, 0x18, 0x14, 0x0a, 0x0b, 0x67, 0x72, 0x70,
	0x63, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x32,
	0xde, 0x02, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2d, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x71, 0x0a,
	0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x2f, 0x2e, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x74, 0x61,
	0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x6e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12,
	0x2d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x53,
	0x79, 0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x2e, 0x53, 0x79,
	0x73, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x75, 0x0a, 0x20, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x73, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x61, 0x6e, 0x64, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x73,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0xaa, 0x02, 0x1c, 0x56, 0x32, 0x52, 0x61, 0x79,
	0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x73, 0x2e,
	0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_stats_command_command_proto_rawDescOnce sync.Once
	file_app_stats_command_command_proto_rawDescData = file_app_stats_command_command_proto_rawDesc
)

func file_app_stats_command_command_proto_rawDescGZIP() []byte {
	file_app_stats_command_command_proto_rawDescOnce.Do(func() {
		file_app_stats_command_command_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_stats_command_command_proto_rawDescData)
	})
	return file_app_stats_command_command_proto_rawDescData
}

var file_app_stats_command_command_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_app_stats_command_command_proto_goTypes = []interface{}{
	(*GetStatsRequest)(nil),    // 0: v2ray.core.app.stats.command.GetStatsRequest
	(*Stat)(nil),               // 1: v2ray.core.app.stats.command.Stat
	(*GetStatsResponse)(nil),   // 2: v2ray.core.app.stats.command.GetStatsResponse
	(*QueryStatsRequest)(nil),  // 3: v2ray.core.app.stats.command.QueryStatsRequest
	(*QueryStatsResponse)(nil), // 4: v2ray.core.app.stats.command.QueryStatsResponse
	(*SysStatsRequest)(nil),    // 5: v2ray.core.app.stats.command.SysStatsRequest
	(*SysStatsResponse)(nil),   // 6: v2ray.core.app.stats.command.SysStatsResponse
	(*Config)(nil),             // 7: v2ray.core.app.stats.command.Config
}
var file_app_stats_command_command_proto_depIdxs = []int32{
	1, // 0: v2ray.core.app.stats.command.GetStatsResponse.stat:type_name -> v2ray.core.app.stats.command.Stat
	1, // 1: v2ray.core.app.stats.command.QueryStatsResponse.stat:type_name -> v2ray.core.app.stats.command.Stat
	0, // 2: v2ray.core.app.stats.command.StatsService.GetStats:input_type -> v2ray.core.app.stats.command.GetStatsRequest
	3, // 3: v2ray.core.app.stats.command.StatsService.QueryStats:input_type -> v2ray.core.app.stats.command.QueryStatsRequest
	5, // 4: v2ray.core.app.stats.command.StatsService.GetSysStats:input_type -> v2ray.core.app.stats.command.SysStatsRequest
	2, // 5: v2ray.core.app.stats.command.StatsService.GetStats:output_type -> v2ray.core.app.stats.command.GetStatsResponse
	4, // 6: v2ray.core.app.stats.command.StatsService.QueryStats:output_type -> v2ray.core.app.stats.command.QueryStatsResponse
	6, // 7: v2ray.core.app.stats.command.StatsService.GetSysStats:output_type -> v2ray.core.app.stats.command.SysStatsResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_app_stats_command_command_proto_init() }
func file_app_stats_command_command_proto_init() {
	if File_app_stats_command_command_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_stats_command_command_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsRequest); i {
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
		file_app_stats_command_command_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stat); i {
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
		file_app_stats_command_command_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatsResponse); i {
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
		file_app_stats_command_command_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryStatsRequest); i {
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
		file_app_stats_command_command_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryStatsResponse); i {
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
		file_app_stats_command_command_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysStatsRequest); i {
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
		file_app_stats_command_command_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysStatsResponse); i {
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
		file_app_stats_command_command_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_app_stats_command_command_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_app_stats_command_command_proto_goTypes,
		DependencyIndexes: file_app_stats_command_command_proto_depIdxs,
		MessageInfos:      file_app_stats_command_command_proto_msgTypes,
	}.Build()
	File_app_stats_command_command_proto = out.File
	file_app_stats_command_command_proto_rawDesc = nil
	file_app_stats_command_command_proto_goTypes = nil
	file_app_stats_command_command_proto_depIdxs = nil
}

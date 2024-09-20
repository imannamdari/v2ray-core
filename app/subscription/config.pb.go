package subscription

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

type ImportSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                 string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Url                  string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	TagPrefix            string `protobuf:"bytes,3,opt,name=tag_prefix,json=tagPrefix,proto3" json:"tag_prefix,omitempty"`
	ImportUsingTag       string `protobuf:"bytes,4,opt,name=import_using_tag,json=importUsingTag,proto3" json:"import_using_tag,omitempty"`
	DefaultExpireSeconds uint64 `protobuf:"varint,5,opt,name=default_expire_seconds,json=defaultExpireSeconds,proto3" json:"default_expire_seconds,omitempty"`
}

func (x *ImportSource) Reset() {
	*x = ImportSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_subscription_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImportSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImportSource) ProtoMessage() {}

func (x *ImportSource) ProtoReflect() protoreflect.Message {
	mi := &file_app_subscription_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImportSource.ProtoReflect.Descriptor instead.
func (*ImportSource) Descriptor() ([]byte, []int) {
	return file_app_subscription_config_proto_rawDescGZIP(), []int{0}
}

func (x *ImportSource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ImportSource) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *ImportSource) GetTagPrefix() string {
	if x != nil {
		return x.TagPrefix
	}
	return ""
}

func (x *ImportSource) GetImportUsingTag() string {
	if x != nil {
		return x.ImportUsingTag
	}
	return ""
}

func (x *ImportSource) GetDefaultExpireSeconds() uint64 {
	if x != nil {
		return x.DefaultExpireSeconds
	}
	return 0
}

// Config is the settings for Subscription Manager.
type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imports                       []*ImportSource `protobuf:"bytes,1,rep,name=imports,proto3" json:"imports,omitempty"`
	NonnativeConverterOverlay     []byte          `protobuf:"bytes,2,opt,name=nonnative_converter_overlay,json=nonnativeConverterOverlay,proto3" json:"nonnative_converter_overlay,omitempty"`
	NonnativeConverterOverlayFile string          `protobuf:"bytes,96002,opt,name=nonnative_converter_overlay_file,json=nonnativeConverterOverlayFile,proto3" json:"nonnative_converter_overlay_file,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_subscription_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_app_subscription_config_proto_msgTypes[1]
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
	return file_app_subscription_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetImports() []*ImportSource {
	if x != nil {
		return x.Imports
	}
	return nil
}

func (x *Config) GetNonnativeConverterOverlay() []byte {
	if x != nil {
		return x.NonnativeConverterOverlay
	}
	return nil
}

func (x *Config) GetNonnativeConverterOverlayFile() string {
	if x != nil {
		return x.NonnativeConverterOverlayFile
	}
	return ""
}

var File_app_subscription_config_proto protoreflect.FileDescriptor

var file_app_subscription_config_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1b, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3,
	0x01, 0x0a, 0x0c, 0x49, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x67, 0x5f, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x67, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x75,
	0x73, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x67, 0x12, 0x34,
	0x0a, 0x16, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x14,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x53, 0x65, 0x63,
	0x6f, 0x6e, 0x64, 0x73, 0x22, 0x98, 0x02, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x43, 0x0a, 0x07, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70,
	0x70, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x49,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x07, 0x69, 0x6d, 0x70,
	0x6f, 0x72, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x1b, 0x6e, 0x6f, 0x6e, 0x6e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x6f, 0x76, 0x65, 0x72,
	0x6c, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x19, 0x6e, 0x6f, 0x6e, 0x6e, 0x61,
	0x74, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x4f, 0x76, 0x65,
	0x72, 0x6c, 0x61, 0x79, 0x12, 0x6c, 0x0a, 0x20, 0x6e, 0x6f, 0x6e, 0x6e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x6f, 0x76, 0x65, 0x72,
	0x6c, 0x61, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x82, 0xee, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x21, 0x82, 0xb5, 0x18, 0x1d, 0x22, 0x1b, 0x6e, 0x6f, 0x6e, 0x6e, 0x61, 0x74, 0x69, 0x76,
	0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x6f, 0x76, 0x65, 0x72,
	0x6c, 0x61, 0x79, 0x52, 0x1d, 0x6e, 0x6f, 0x6e, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x43, 0x6f,
	0x6e, 0x76, 0x65, 0x72, 0x74, 0x65, 0x72, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x61, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x3a, 0x1b, 0x82, 0xb5, 0x18, 0x17, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x72, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x35, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0xaa, 0x02, 0x1b, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f,
	0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_subscription_config_proto_rawDescOnce sync.Once
	file_app_subscription_config_proto_rawDescData = file_app_subscription_config_proto_rawDesc
)

func file_app_subscription_config_proto_rawDescGZIP() []byte {
	file_app_subscription_config_proto_rawDescOnce.Do(func() {
		file_app_subscription_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_subscription_config_proto_rawDescData)
	})
	return file_app_subscription_config_proto_rawDescData
}

var file_app_subscription_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_subscription_config_proto_goTypes = []interface{}{
	(*ImportSource)(nil), // 0: v2ray.core.app.subscription.ImportSource
	(*Config)(nil),       // 1: v2ray.core.app.subscription.Config
}
var file_app_subscription_config_proto_depIdxs = []int32{
	0, // 0: v2ray.core.app.subscription.Config.imports:type_name -> v2ray.core.app.subscription.ImportSource
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_subscription_config_proto_init() }
func file_app_subscription_config_proto_init() {
	if File_app_subscription_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_subscription_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImportSource); i {
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
		file_app_subscription_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_app_subscription_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_subscription_config_proto_goTypes,
		DependencyIndexes: file_app_subscription_config_proto_depIdxs,
		MessageInfos:      file_app_subscription_config_proto_msgTypes,
	}.Build()
	File_app_subscription_config_proto = out.File
	file_app_subscription_config_proto_rawDesc = nil
	file_app_subscription_config_proto_goTypes = nil
	file_app_subscription_config_proto_depIdxs = nil
}

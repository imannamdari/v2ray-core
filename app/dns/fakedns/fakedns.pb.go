package fakedns

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

type FakeDnsPool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpPool  string `protobuf:"bytes,1,opt,name=ip_pool,json=ipPool,proto3" json:"ip_pool,omitempty"` //CIDR of IP pool used as fake DNS IP
	LruSize int64  `protobuf:"varint,2,opt,name=lruSize,proto3" json:"lruSize,omitempty"`            //Size of Pool for remembering relationship between domain name and IP address
}

func (x *FakeDnsPool) Reset() {
	*x = FakeDnsPool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dns_fakedns_fakedns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FakeDnsPool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FakeDnsPool) ProtoMessage() {}

func (x *FakeDnsPool) ProtoReflect() protoreflect.Message {
	mi := &file_app_dns_fakedns_fakedns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FakeDnsPool.ProtoReflect.Descriptor instead.
func (*FakeDnsPool) Descriptor() ([]byte, []int) {
	return file_app_dns_fakedns_fakedns_proto_rawDescGZIP(), []int{0}
}

func (x *FakeDnsPool) GetIpPool() string {
	if x != nil {
		return x.IpPool
	}
	return ""
}

func (x *FakeDnsPool) GetLruSize() int64 {
	if x != nil {
		return x.LruSize
	}
	return 0
}

type FakeDnsPoolMulti struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pools []*FakeDnsPool `protobuf:"bytes,1,rep,name=pools,proto3" json:"pools,omitempty"`
}

func (x *FakeDnsPoolMulti) Reset() {
	*x = FakeDnsPoolMulti{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_dns_fakedns_fakedns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FakeDnsPoolMulti) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FakeDnsPoolMulti) ProtoMessage() {}

func (x *FakeDnsPoolMulti) ProtoReflect() protoreflect.Message {
	mi := &file_app_dns_fakedns_fakedns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FakeDnsPoolMulti.ProtoReflect.Descriptor instead.
func (*FakeDnsPoolMulti) Descriptor() ([]byte, []int) {
	return file_app_dns_fakedns_fakedns_proto_rawDescGZIP(), []int{1}
}

func (x *FakeDnsPoolMulti) GetPools() []*FakeDnsPool {
	if x != nil {
		return x.Pools
	}
	return nil
}

var File_app_dns_fakedns_fakedns_proto protoreflect.FileDescriptor

var file_app_dns_fakedns_fakedns_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x70, 0x2f, 0x64, 0x6e, 0x73, 0x2f, 0x66, 0x61, 0x6b, 0x65, 0x64, 0x6e,
	0x73, 0x2f, 0x66, 0x61, 0x6b, 0x65, 0x64, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x64, 0x6e, 0x73, 0x2e, 0x66, 0x61, 0x6b, 0x65, 0x64, 0x6e, 0x73, 0x1a, 0x20, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x58, 0x0a,
	0x0b, 0x46, 0x61, 0x6b, 0x65, 0x44, 0x6e, 0x73, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x17, 0x0a, 0x07,
	0x69, 0x70, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69,
	0x70, 0x50, 0x6f, 0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x72, 0x75, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6c, 0x72, 0x75, 0x53, 0x69, 0x7a, 0x65, 0x3a,
	0x16, 0x82, 0xb5, 0x18, 0x12, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x07,
	0x66, 0x61, 0x6b, 0x65, 0x44, 0x6e, 0x73, 0x22, 0x6e, 0x0a, 0x10, 0x46, 0x61, 0x6b, 0x65, 0x44,
	0x6e, 0x73, 0x50, 0x6f, 0x6f, 0x6c, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x12, 0x3d, 0x0a, 0x05, 0x70,
	0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x32, 0x72,
	0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x6e, 0x73, 0x2e,
	0x66, 0x61, 0x6b, 0x65, 0x64, 0x6e, 0x73, 0x2e, 0x46, 0x61, 0x6b, 0x65, 0x44, 0x6e, 0x73, 0x50,
	0x6f, 0x6f, 0x6c, 0x52, 0x05, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x3a, 0x1b, 0x82, 0xb5, 0x18, 0x17,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0c, 0x66, 0x61, 0x6b, 0x65, 0x44,
	0x6e, 0x73, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x42, 0x6f, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x64, 0x6e,
	0x73, 0x2e, 0x66, 0x61, 0x6b, 0x65, 0x64, 0x6e, 0x73, 0x50, 0x01, 0x5a, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x61, 0x70, 0x70, 0x2f,
	0x64, 0x6e, 0x73, 0x2f, 0x66, 0x61, 0x6b, 0x65, 0x64, 0x6e, 0x73, 0xaa, 0x02, 0x1a, 0x56, 0x32,
	0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x2e, 0x44, 0x6e, 0x73,
	0x2e, 0x46, 0x61, 0x6b, 0x65, 0x64, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_dns_fakedns_fakedns_proto_rawDescOnce sync.Once
	file_app_dns_fakedns_fakedns_proto_rawDescData = file_app_dns_fakedns_fakedns_proto_rawDesc
)

func file_app_dns_fakedns_fakedns_proto_rawDescGZIP() []byte {
	file_app_dns_fakedns_fakedns_proto_rawDescOnce.Do(func() {
		file_app_dns_fakedns_fakedns_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_dns_fakedns_fakedns_proto_rawDescData)
	})
	return file_app_dns_fakedns_fakedns_proto_rawDescData
}

var file_app_dns_fakedns_fakedns_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_app_dns_fakedns_fakedns_proto_goTypes = []interface{}{
	(*FakeDnsPool)(nil),      // 0: v2ray.core.app.dns.fakedns.FakeDnsPool
	(*FakeDnsPoolMulti)(nil), // 1: v2ray.core.app.dns.fakedns.FakeDnsPoolMulti
}
var file_app_dns_fakedns_fakedns_proto_depIdxs = []int32{
	0, // 0: v2ray.core.app.dns.fakedns.FakeDnsPoolMulti.pools:type_name -> v2ray.core.app.dns.fakedns.FakeDnsPool
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_app_dns_fakedns_fakedns_proto_init() }
func file_app_dns_fakedns_fakedns_proto_init() {
	if File_app_dns_fakedns_fakedns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_dns_fakedns_fakedns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FakeDnsPool); i {
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
		file_app_dns_fakedns_fakedns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FakeDnsPoolMulti); i {
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
			RawDescriptor: file_app_dns_fakedns_fakedns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_dns_fakedns_fakedns_proto_goTypes,
		DependencyIndexes: file_app_dns_fakedns_fakedns_proto_depIdxs,
		MessageInfos:      file_app_dns_fakedns_fakedns_proto_msgTypes,
	}.Build()
	File_app_dns_fakedns_fakedns_proto = out.File
	file_app_dns_fakedns_fakedns_proto_rawDesc = nil
	file_app_dns_fakedns_fakedns_proto_goTypes = nil
	file_app_dns_fakedns_fakedns_proto_depIdxs = nil
}

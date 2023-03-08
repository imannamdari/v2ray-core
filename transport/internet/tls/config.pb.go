// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: transport/internet/tls/config.proto

package tls

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

type Certificate_Usage int32

const (
	Certificate_ENCIPHERMENT            Certificate_Usage = 0
	Certificate_AUTHORITY_VERIFY        Certificate_Usage = 1
	Certificate_AUTHORITY_ISSUE         Certificate_Usage = 2
	Certificate_AUTHORITY_VERIFY_CLIENT Certificate_Usage = 3
)

// Enum value maps for Certificate_Usage.
var (
	Certificate_Usage_name = map[int32]string{
		0: "ENCIPHERMENT",
		1: "AUTHORITY_VERIFY",
		2: "AUTHORITY_ISSUE",
		3: "AUTHORITY_VERIFY_CLIENT",
	}
	Certificate_Usage_value = map[string]int32{
		"ENCIPHERMENT":            0,
		"AUTHORITY_VERIFY":        1,
		"AUTHORITY_ISSUE":         2,
		"AUTHORITY_VERIFY_CLIENT": 3,
	}
)

func (x Certificate_Usage) Enum() *Certificate_Usage {
	p := new(Certificate_Usage)
	*p = x
	return p
}

func (x Certificate_Usage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Certificate_Usage) Descriptor() protoreflect.EnumDescriptor {
	return file_transport_internet_tls_config_proto_enumTypes[0].Descriptor()
}

func (Certificate_Usage) Type() protoreflect.EnumType {
	return &file_transport_internet_tls_config_proto_enumTypes[0]
}

func (x Certificate_Usage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Certificate_Usage.Descriptor instead.
func (Certificate_Usage) EnumDescriptor() ([]byte, []int) {
	return file_transport_internet_tls_config_proto_rawDescGZIP(), []int{0, 0}
}

type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TLS certificate in x509 format.
	Certificate []byte `protobuf:"bytes,1,opt,name=Certificate,proto3" json:"Certificate,omitempty"`
	// TLS key in x509 format.
	Key             []byte            `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
	Usage           Certificate_Usage `protobuf:"varint,3,opt,name=usage,proto3,enum=v2ray.core.transport.internet.tls.Certificate_Usage" json:"usage,omitempty"`
	CertificateFile string            `protobuf:"bytes,96001,opt,name=certificate_file,json=certificateFile,proto3" json:"certificate_file,omitempty"`
	KeyFile         string            `protobuf:"bytes,96002,opt,name=key_file,json=keyFile,proto3" json:"key_file,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_tls_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_tls_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_transport_internet_tls_config_proto_rawDescGZIP(), []int{0}
}

func (x *Certificate) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *Certificate) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *Certificate) GetUsage() Certificate_Usage {
	if x != nil {
		return x.Usage
	}
	return Certificate_ENCIPHERMENT
}

func (x *Certificate) GetCertificateFile() string {
	if x != nil {
		return x.CertificateFile
	}
	return ""
}

func (x *Certificate) GetKeyFile() string {
	if x != nil {
		return x.KeyFile
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether or not to allow self-signed certificates.
	AllowInsecure bool `protobuf:"varint,1,opt,name=allow_insecure,json=allowInsecure,proto3" json:"allow_insecure,omitempty"`
	// List of certificates to be served on server.
	Certificate []*Certificate `protobuf:"bytes,2,rep,name=certificate,proto3" json:"certificate,omitempty"`
	// Override server name.
	ServerName string `protobuf:"bytes,3,opt,name=server_name,json=serverName,proto3" json:"server_name,omitempty"`
	// Lists of string as ALPN values.
	NextProtocol []string `protobuf:"bytes,4,rep,name=next_protocol,json=nextProtocol,proto3" json:"next_protocol,omitempty"`
	// Whether or not to enable session (ticket) resumption.
	EnableSessionResumption bool `protobuf:"varint,5,opt,name=enable_session_resumption,json=enableSessionResumption,proto3" json:"enable_session_resumption,omitempty"`
	// If true, root certificates on the system will not be loaded for
	// verification.
	DisableSystemRoot bool `protobuf:"varint,6,opt,name=disable_system_root,json=disableSystemRoot,proto3" json:"disable_system_root,omitempty"`
	// @Document A pinned certificate chain sha256 hash.
	// @Document If the server's hash does not match this value, the connection will be aborted.
	// @Document This value replace allow_insecure.
	// @Critical
	PinnedPeerCertificateChainSha256 [][]byte `protobuf:"bytes,7,rep,name=pinned_peer_certificate_chain_sha256,json=pinnedPeerCertificateChainSha256,proto3" json:"pinned_peer_certificate_chain_sha256,omitempty"`
	// If true, the client is required to present a certificate.
	VerifyClientCertificate bool `protobuf:"varint,8,opt,name=verify_client_certificate,json=verifyClientCertificate,proto3" json:"verify_client_certificate,omitempty"`
	// Enable this if you want tls client hello encryption.
	EnableEch  bool        `protobuf:"varint,9,opt,name=enable_ech,json=enableEch,proto3" json:"enable_ech,omitempty"`
	EchSetting *ECHSetting `protobuf:"bytes,10,opt,name=ech_setting,json=echSetting,proto3" json:"ech_setting,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_tls_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_tls_config_proto_msgTypes[1]
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
	return file_transport_internet_tls_config_proto_rawDescGZIP(), []int{1}
}

func (x *Config) GetAllowInsecure() bool {
	if x != nil {
		return x.AllowInsecure
	}
	return false
}

func (x *Config) GetCertificate() []*Certificate {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *Config) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

func (x *Config) GetNextProtocol() []string {
	if x != nil {
		return x.NextProtocol
	}
	return nil
}

func (x *Config) GetEnableSessionResumption() bool {
	if x != nil {
		return x.EnableSessionResumption
	}
	return false
}

func (x *Config) GetDisableSystemRoot() bool {
	if x != nil {
		return x.DisableSystemRoot
	}
	return false
}

func (x *Config) GetPinnedPeerCertificateChainSha256() [][]byte {
	if x != nil {
		return x.PinnedPeerCertificateChainSha256
	}
	return nil
}

func (x *Config) GetVerifyClientCertificate() bool {
	if x != nil {
		return x.VerifyClientCertificate
	}
	return false
}

func (x *Config) GetEnableEch() bool {
	if x != nil {
		return x.EnableEch
	}
	return false
}

func (x *Config) GetEchSetting() *ECHSetting {
	if x != nil {
		return x.EchSetting
	}
	return nil
}

type ECHSetting struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DnsAddr    string `protobuf:"bytes,1,opt,name=dns_addr,json=dnsAddr,proto3" json:"dns_addr,omitempty"`
	InitEchKey string `protobuf:"bytes,2,opt,name=init_ech_key,json=initEchKey,proto3" json:"init_ech_key,omitempty"`
}

func (x *ECHSetting) Reset() {
	*x = ECHSetting{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transport_internet_tls_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ECHSetting) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ECHSetting) ProtoMessage() {}

func (x *ECHSetting) ProtoReflect() protoreflect.Message {
	mi := &file_transport_internet_tls_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ECHSetting.ProtoReflect.Descriptor instead.
func (*ECHSetting) Descriptor() ([]byte, []int) {
	return file_transport_internet_tls_config_proto_rawDescGZIP(), []int{2}
}

func (x *ECHSetting) GetDnsAddr() string {
	if x != nil {
		return x.DnsAddr
	}
	return ""
}

func (x *ECHSetting) GetInitEchKey() string {
	if x != nil {
		return x.InitEchKey
	}
	return ""
}

var File_transport_internet_tls_config_proto protoreflect.FileDescriptor

var file_transport_internet_tls_config_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2f, 0x74, 0x6c, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2e, 0x74, 0x6c, 0x73, 0x1a, 0x20, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x02, 0x0a, 0x0b, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x4a,
	0x0a, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x74, 0x6c,
	0x73, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x05, 0x75, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3e, 0x0a, 0x10, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x81,
	0xee, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x11, 0x82, 0xb5, 0x18, 0x0d, 0x22, 0x0b, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x0f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x08, 0x6b, 0x65,
	0x79, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x82, 0xee, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0x82, 0xb5, 0x18, 0x05, 0x22, 0x03, 0x4b, 0x65, 0x79, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x46, 0x69,
	0x6c, 0x65, 0x22, 0x61, 0x0a, 0x05, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x0c, 0x45,
	0x4e, 0x43, 0x49, 0x50, 0x48, 0x45, 0x52, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x00, 0x12, 0x14, 0x0a,
	0x10, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46,
	0x59, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x54, 0x59,
	0x5f, 0x49, 0x53, 0x53, 0x55, 0x45, 0x10, 0x02, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x55, 0x54, 0x48,
	0x4f, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x56, 0x45, 0x52, 0x49, 0x46, 0x59, 0x5f, 0x43, 0x4c, 0x49,
	0x45, 0x4e, 0x54, 0x10, 0x03, 0x22, 0xcf, 0x04, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x2d, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x42, 0x06, 0x82, 0xb5, 0x18, 0x02, 0x28, 0x01,
	0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x49, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12,
	0x50, 0x0a, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x65, 0x78, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x3a, 0x0a, 0x19, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6d, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6d, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x11, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x6f, 0x6f, 0x74, 0x12, 0x4e, 0x0a, 0x24, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x5f, 0x70, 0x65,
	0x65, 0x72, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5f, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x20, 0x70, 0x69, 0x6e, 0x6e, 0x65, 0x64, 0x50, 0x65, 0x65, 0x72, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x68, 0x61,
	0x32, 0x35, 0x36, 0x12, 0x3a, 0x0a, 0x19, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x65, 0x63, 0x68, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x63, 0x68, 0x12, 0x4e,
	0x0a, 0x0b, 0x65, 0x63, 0x68, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x2e, 0x74, 0x6c, 0x73, 0x2e, 0x45, 0x43, 0x48, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x52, 0x0a, 0x65, 0x63, 0x68, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x3a, 0x17,
	0x82, 0xb5, 0x18, 0x0a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x82, 0xb5,
	0x18, 0x05, 0x12, 0x03, 0x74, 0x6c, 0x73, 0x22, 0x49, 0x0a, 0x0a, 0x45, 0x43, 0x48, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x0a, 0x08, 0x64, 0x6e, 0x73, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x6e, 0x73, 0x41, 0x64, 0x64, 0x72,
	0x12, 0x20, 0x0a, 0x0c, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x65, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x45, 0x63, 0x68, 0x4b,
	0x65, 0x79, 0x42, 0x84, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x74, 0x6c, 0x73, 0x50, 0x01, 0x5a, 0x35,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79,
	0x2f, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x2f, 0x74, 0x6c, 0x73, 0xaa, 0x02, 0x21, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f,
	0x72, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x2e, 0x54, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_transport_internet_tls_config_proto_rawDescOnce sync.Once
	file_transport_internet_tls_config_proto_rawDescData = file_transport_internet_tls_config_proto_rawDesc
)

func file_transport_internet_tls_config_proto_rawDescGZIP() []byte {
	file_transport_internet_tls_config_proto_rawDescOnce.Do(func() {
		file_transport_internet_tls_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_transport_internet_tls_config_proto_rawDescData)
	})
	return file_transport_internet_tls_config_proto_rawDescData
}

var file_transport_internet_tls_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_transport_internet_tls_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_transport_internet_tls_config_proto_goTypes = []interface{}{
	(Certificate_Usage)(0), // 0: v2ray.core.transport.internet.tls.Certificate.Usage
	(*Certificate)(nil),    // 1: v2ray.core.transport.internet.tls.Certificate
	(*Config)(nil),         // 2: v2ray.core.transport.internet.tls.Config
	(*ECHSetting)(nil),     // 3: v2ray.core.transport.internet.tls.ECHSetting
}
var file_transport_internet_tls_config_proto_depIdxs = []int32{
	0, // 0: v2ray.core.transport.internet.tls.Certificate.usage:type_name -> v2ray.core.transport.internet.tls.Certificate.Usage
	1, // 1: v2ray.core.transport.internet.tls.Config.certificate:type_name -> v2ray.core.transport.internet.tls.Certificate
	3, // 2: v2ray.core.transport.internet.tls.Config.ech_setting:type_name -> v2ray.core.transport.internet.tls.ECHSetting
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_transport_internet_tls_config_proto_init() }
func file_transport_internet_tls_config_proto_init() {
	if File_transport_internet_tls_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transport_internet_tls_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
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
		file_transport_internet_tls_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_transport_internet_tls_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ECHSetting); i {
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
			RawDescriptor: file_transport_internet_tls_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transport_internet_tls_config_proto_goTypes,
		DependencyIndexes: file_transport_internet_tls_config_proto_depIdxs,
		EnumInfos:         file_transport_internet_tls_config_proto_enumTypes,
		MessageInfos:      file_transport_internet_tls_config_proto_msgTypes,
	}.Build()
	File_transport_internet_tls_config_proto = out.File
	file_transport_internet_tls_config_proto_rawDesc = nil
	file_transport_internet_tls_config_proto_goTypes = nil
	file_transport_internet_tls_config_proto_depIdxs = nil
}

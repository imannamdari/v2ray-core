package specs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ServerConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol          string     `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"`
	ProtocolSettings  *anypb.Any `protobuf:"bytes,2,opt,name=protocol_settings,json=protocolSettings,proto3" json:"protocol_settings,omitempty"`
	Transport         string     `protobuf:"bytes,3,opt,name=transport,proto3" json:"transport,omitempty"`
	TransportSettings *anypb.Any `protobuf:"bytes,4,opt,name=transport_settings,json=transportSettings,proto3" json:"transport_settings,omitempty"`
	Security          string     `protobuf:"bytes,5,opt,name=security,proto3" json:"security,omitempty"`
	SecuritySettings  *anypb.Any `protobuf:"bytes,6,opt,name=security_settings,json=securitySettings,proto3" json:"security_settings,omitempty"`
}

func (x *ServerConfiguration) Reset() {
	*x = ServerConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_subscription_specs_abstract_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfiguration) ProtoMessage() {}

func (x *ServerConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_app_subscription_specs_abstract_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfiguration.ProtoReflect.Descriptor instead.
func (*ServerConfiguration) Descriptor() ([]byte, []int) {
	return file_app_subscription_specs_abstract_spec_proto_rawDescGZIP(), []int{0}
}

func (x *ServerConfiguration) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ServerConfiguration) GetProtocolSettings() *anypb.Any {
	if x != nil {
		return x.ProtocolSettings
	}
	return nil
}

func (x *ServerConfiguration) GetTransport() string {
	if x != nil {
		return x.Transport
	}
	return ""
}

func (x *ServerConfiguration) GetTransportSettings() *anypb.Any {
	if x != nil {
		return x.TransportSettings
	}
	return nil
}

func (x *ServerConfiguration) GetSecurity() string {
	if x != nil {
		return x.Security
	}
	return ""
}

func (x *ServerConfiguration) GetSecuritySettings() *anypb.Any {
	if x != nil {
		return x.SecuritySettings
	}
	return nil
}

type SubscriptionServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Metadata      map[string]string    `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Configuration *ServerConfiguration `protobuf:"bytes,3,opt,name=configuration,proto3" json:"configuration,omitempty"`
}

func (x *SubscriptionServerConfig) Reset() {
	*x = SubscriptionServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_subscription_specs_abstract_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionServerConfig) ProtoMessage() {}

func (x *SubscriptionServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_app_subscription_specs_abstract_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionServerConfig.ProtoReflect.Descriptor instead.
func (*SubscriptionServerConfig) Descriptor() ([]byte, []int) {
	return file_app_subscription_specs_abstract_spec_proto_rawDescGZIP(), []int{1}
}

func (x *SubscriptionServerConfig) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubscriptionServerConfig) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SubscriptionServerConfig) GetConfiguration() *ServerConfiguration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

type SubscriptionDocument struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata map[string]string           `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Server   []*SubscriptionServerConfig `protobuf:"bytes,3,rep,name=server,proto3" json:"server,omitempty"`
}

func (x *SubscriptionDocument) Reset() {
	*x = SubscriptionDocument{}
	if protoimpl.UnsafeEnabled {
		mi := &file_app_subscription_specs_abstract_spec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionDocument) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionDocument) ProtoMessage() {}

func (x *SubscriptionDocument) ProtoReflect() protoreflect.Message {
	mi := &file_app_subscription_specs_abstract_spec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionDocument.ProtoReflect.Descriptor instead.
func (*SubscriptionDocument) Descriptor() ([]byte, []int) {
	return file_app_subscription_specs_abstract_spec_proto_rawDescGZIP(), []int{2}
}

func (x *SubscriptionDocument) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SubscriptionDocument) GetServer() []*SubscriptionServerConfig {
	if x != nil {
		return x.Server
	}
	return nil
}

var File_app_subscription_specs_abstract_spec_proto protoreflect.FileDescriptor

var file_app_subscription_specs_abstract_spec_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x61, 0x70, 0x70, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2f, 0x61, 0x62, 0x73, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x76, 0x32,
	0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02, 0x0a, 0x13, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x41,
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x43, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x12, 0x41, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x53, 0x65, 0x74, 0x74, 0x69,
	0x6e, 0x67, 0x73, 0x22, 0xac, 0x02, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x65, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x49, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x61, 0x70, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x5c, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x70, 0x65,
	0x63, 0x73, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x8b, 0x02, 0x0a, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x61, 0x0a, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x45, 0x2e,
	0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x70, 0x65, 0x63,
	0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x53,
	0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b,
	0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x70, 0x65,
	0x63, 0x73, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x84, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x32, 0x72, 0x61, 0x79, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x70, 0x65, 0x63, 0x73, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x32, 0x66, 0x6c, 0x79, 0x2f, 0x76,
	0x32, 0x72, 0x61, 0x79, 0x2d, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x35, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x70,
	0x65, 0x63, 0x73, 0xaa, 0x02, 0x21, 0x56, 0x32, 0x52, 0x61, 0x79, 0x2e, 0x43, 0x6f, 0x72, 0x65,
	0x2e, 0x41, 0x70, 0x70, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x53, 0x70, 0x65, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_app_subscription_specs_abstract_spec_proto_rawDescOnce sync.Once
	file_app_subscription_specs_abstract_spec_proto_rawDescData = file_app_subscription_specs_abstract_spec_proto_rawDesc
)

func file_app_subscription_specs_abstract_spec_proto_rawDescGZIP() []byte {
	file_app_subscription_specs_abstract_spec_proto_rawDescOnce.Do(func() {
		file_app_subscription_specs_abstract_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_app_subscription_specs_abstract_spec_proto_rawDescData)
	})
	return file_app_subscription_specs_abstract_spec_proto_rawDescData
}

var file_app_subscription_specs_abstract_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_app_subscription_specs_abstract_spec_proto_goTypes = []interface{}{
	(*ServerConfiguration)(nil),      // 0: v2ray.core.app.subscription.specs.ServerConfiguration
	(*SubscriptionServerConfig)(nil), // 1: v2ray.core.app.subscription.specs.SubscriptionServerConfig
	(*SubscriptionDocument)(nil),     // 2: v2ray.core.app.subscription.specs.SubscriptionDocument
	nil,                              // 3: v2ray.core.app.subscription.specs.SubscriptionServerConfig.MetadataEntry
	nil,                              // 4: v2ray.core.app.subscription.specs.SubscriptionDocument.MetadataEntry
	(*anypb.Any)(nil),                // 5: google.protobuf.Any
}
var file_app_subscription_specs_abstract_spec_proto_depIdxs = []int32{
	5, // 0: v2ray.core.app.subscription.specs.ServerConfiguration.protocol_settings:type_name -> google.protobuf.Any
	5, // 1: v2ray.core.app.subscription.specs.ServerConfiguration.transport_settings:type_name -> google.protobuf.Any
	5, // 2: v2ray.core.app.subscription.specs.ServerConfiguration.security_settings:type_name -> google.protobuf.Any
	3, // 3: v2ray.core.app.subscription.specs.SubscriptionServerConfig.metadata:type_name -> v2ray.core.app.subscription.specs.SubscriptionServerConfig.MetadataEntry
	0, // 4: v2ray.core.app.subscription.specs.SubscriptionServerConfig.configuration:type_name -> v2ray.core.app.subscription.specs.ServerConfiguration
	4, // 5: v2ray.core.app.subscription.specs.SubscriptionDocument.metadata:type_name -> v2ray.core.app.subscription.specs.SubscriptionDocument.MetadataEntry
	1, // 6: v2ray.core.app.subscription.specs.SubscriptionDocument.server:type_name -> v2ray.core.app.subscription.specs.SubscriptionServerConfig
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_app_subscription_specs_abstract_spec_proto_init() }
func file_app_subscription_specs_abstract_spec_proto_init() {
	if File_app_subscription_specs_abstract_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_app_subscription_specs_abstract_spec_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfiguration); i {
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
		file_app_subscription_specs_abstract_spec_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionServerConfig); i {
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
		file_app_subscription_specs_abstract_spec_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionDocument); i {
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
			RawDescriptor: file_app_subscription_specs_abstract_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_app_subscription_specs_abstract_spec_proto_goTypes,
		DependencyIndexes: file_app_subscription_specs_abstract_spec_proto_depIdxs,
		MessageInfos:      file_app_subscription_specs_abstract_spec_proto_msgTypes,
	}.Build()
	File_app_subscription_specs_abstract_spec_proto = out.File
	file_app_subscription_specs_abstract_spec_proto_rawDesc = nil
	file_app_subscription_specs_abstract_spec_proto_goTypes = nil
	file_app_subscription_specs_abstract_spec_proto_depIdxs = nil
}

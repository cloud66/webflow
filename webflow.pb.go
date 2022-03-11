// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: webflow.proto

package webflow

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family string `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	Major  string `protobuf:"bytes,2,opt,name=major,proto3" json:"major,omitempty"`
	Minor  string `protobuf:"bytes,3,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch  string `protobuf:"bytes,4,opt,name=patch,proto3" json:"patch,omitempty"`
}

func (x *Client) Reset() {
	*x = Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webflow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Client) ProtoMessage() {}

func (x *Client) ProtoReflect() protoreflect.Message {
	mi := &file_webflow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Client.ProtoReflect.Descriptor instead.
func (*Client) Descriptor() ([]byte, []int) {
	return file_webflow_proto_rawDescGZIP(), []int{0}
}

func (x *Client) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

func (x *Client) GetMajor() string {
	if x != nil {
		return x.Major
	}
	return ""
}

func (x *Client) GetMinor() string {
	if x != nil {
		return x.Minor
	}
	return ""
}

func (x *Client) GetPatch() string {
	if x != nil {
		return x.Patch
	}
	return ""
}

type OS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family     string `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	Major      string `protobuf:"bytes,2,opt,name=major,proto3" json:"major,omitempty"`
	Minor      string `protobuf:"bytes,3,opt,name=minor,proto3" json:"minor,omitempty"`
	Patch      string `protobuf:"bytes,4,opt,name=patch,proto3" json:"patch,omitempty"`
	PatchMinor string `protobuf:"bytes,5,opt,name=patch_minor,json=patchMinor,proto3" json:"patch_minor,omitempty"`
}

func (x *OS) Reset() {
	*x = OS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webflow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OS) ProtoMessage() {}

func (x *OS) ProtoReflect() protoreflect.Message {
	mi := &file_webflow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OS.ProtoReflect.Descriptor instead.
func (*OS) Descriptor() ([]byte, []int) {
	return file_webflow_proto_rawDescGZIP(), []int{1}
}

func (x *OS) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

func (x *OS) GetMajor() string {
	if x != nil {
		return x.Major
	}
	return ""
}

func (x *OS) GetMinor() string {
	if x != nil {
		return x.Minor
	}
	return ""
}

func (x *OS) GetPatch() string {
	if x != nil {
		return x.Patch
	}
	return ""
}

func (x *OS) GetPatchMinor() string {
	if x != nil {
		return x.PatchMinor
	}
	return ""
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Family string `protobuf:"bytes,1,opt,name=family,proto3" json:"family,omitempty"`
	Brand  string `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Model  string `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webflow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_webflow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_webflow_proto_rawDescGZIP(), []int{2}
}

func (x *Device) GetFamily() string {
	if x != nil {
		return x.Family
	}
	return ""
}

func (x *Device) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Device) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

type UserAgent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Client *Client `protobuf:"bytes,1,opt,name=client,proto3" json:"client,omitempty"`
	Os     *OS     `protobuf:"bytes,2,opt,name=os,proto3" json:"os,omitempty"`
	Device *Device `protobuf:"bytes,3,opt,name=device,proto3" json:"device,omitempty"`
}

func (x *UserAgent) Reset() {
	*x = UserAgent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webflow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAgent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAgent) ProtoMessage() {}

func (x *UserAgent) ProtoReflect() protoreflect.Message {
	mi := &file_webflow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAgent.ProtoReflect.Descriptor instead.
func (*UserAgent) Descriptor() ([]byte, []int) {
	return file_webflow_proto_rawDescGZIP(), []int{3}
}

func (x *UserAgent) GetClient() *Client {
	if x != nil {
		return x.Client
	}
	return nil
}

func (x *UserAgent) GetOs() *OS {
	if x != nil {
		return x.Os
	}
	return nil
}

func (x *UserAgent) GetDevice() *Device {
	if x != nil {
		return x.Device
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserAgentRaw string     `protobuf:"bytes,1,opt,name=user_agent_raw,json=userAgentRaw,proto3" json:"user_agent_raw,omitempty"`
	UserAgent    *UserAgent `protobuf:"bytes,2,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
	Url          string     `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Path         string     `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	Method       string     `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	Host         string     `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webflow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_webflow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_webflow_proto_rawDescGZIP(), []int{4}
}

func (x *Request) GetUserAgentRaw() string {
	if x != nil {
		return x.UserAgentRaw
	}
	return ""
}

func (x *Request) GetUserAgent() *UserAgent {
	if x != nil {
		return x.UserAgent
	}
	return nil
}

func (x *Request) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Request) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

type Origin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip            string  `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	ContinentCode string  `protobuf:"bytes,2,opt,name=continent_code,json=continentCode,proto3" json:"continent_code,omitempty"`
	CountryCode   string  `protobuf:"bytes,3,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	CityName      string  `protobuf:"bytes,4,opt,name=city_name,json=cityName,proto3" json:"city_name,omitempty"`
	Latitude      float32 `protobuf:"fixed32,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude     float32 `protobuf:"fixed32,6,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Timezone      string  `protobuf:"bytes,7,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Asn           string  `protobuf:"bytes,8,opt,name=asn,proto3" json:"asn,omitempty"`
}

func (x *Origin) Reset() {
	*x = Origin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webflow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Origin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Origin) ProtoMessage() {}

func (x *Origin) ProtoReflect() protoreflect.Message {
	mi := &file_webflow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Origin.ProtoReflect.Descriptor instead.
func (*Origin) Descriptor() ([]byte, []int) {
	return file_webflow_proto_rawDescGZIP(), []int{5}
}

func (x *Origin) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Origin) GetContinentCode() string {
	if x != nil {
		return x.ContinentCode
	}
	return ""
}

func (x *Origin) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *Origin) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *Origin) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Origin) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Origin) GetTimezone() string {
	if x != nil {
		return x.Timezone
	}
	return ""
}

func (x *Origin) GetAsn() string {
	if x != nil {
		return x.Asn
	}
	return ""
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	When    bool   `protobuf:"varint,1,opt,name=when,proto3" json:"when,omitempty"`
	With    int32  `protobuf:"varint,2,opt,name=with,proto3" json:"with,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webflow_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_webflow_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_webflow_proto_rawDescGZIP(), []int{6}
}

func (x *Block) GetWhen() bool {
	if x != nil {
		return x.When
	}
	return false
}

func (x *Block) GetWith() int32 {
	if x != nil {
		return x.With
	}
	return 0
}

func (x *Block) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Add    map[string]string     `protobuf:"bytes,1,rep,name=add,proto3" json:"add,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Remove []string              `protobuf:"bytes,2,rep,name=remove,proto3" json:"remove,omitempty"`
	Set    map[string]string     `protobuf:"bytes,3,rep,name=set,proto3" json:"set,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	When   *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=when,proto3" json:"when,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webflow_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_webflow_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_webflow_proto_rawDescGZIP(), []int{7}
}

func (x *Header) GetAdd() map[string]string {
	if x != nil {
		return x.Add
	}
	return nil
}

func (x *Header) GetRemove() []string {
	if x != nil {
		return x.Remove
	}
	return nil
}

func (x *Header) GetSet() map[string]string {
	if x != nil {
		return x.Set
	}
	return nil
}

func (x *Header) GetWhen() *wrapperspb.BoolValue {
	if x != nil {
		return x.When
	}
	return nil
}

type Redirect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string                `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   string                `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	With int32                 `protobuf:"varint,3,opt,name=with,proto3" json:"with,omitempty"`
	When *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=when,proto3" json:"when,omitempty"`
}

func (x *Redirect) Reset() {
	*x = Redirect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webflow_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Redirect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Redirect) ProtoMessage() {}

func (x *Redirect) ProtoReflect() protoreflect.Message {
	mi := &file_webflow_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Redirect.ProtoReflect.Descriptor instead.
func (*Redirect) Descriptor() ([]byte, []int) {
	return file_webflow_proto_rawDescGZIP(), []int{8}
}

func (x *Redirect) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Redirect) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Redirect) GetWith() int32 {
	if x != nil {
		return x.With
	}
	return 0
}

func (x *Redirect) GetWhen() *wrapperspb.BoolValue {
	if x != nil {
		return x.When
	}
	return nil
}

type Rewrite struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From string                `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To   string                `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	When *wrapperspb.BoolValue `protobuf:"bytes,3,opt,name=when,proto3" json:"when,omitempty"`
}

func (x *Rewrite) Reset() {
	*x = Rewrite{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webflow_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rewrite) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rewrite) ProtoMessage() {}

func (x *Rewrite) ProtoReflect() protoreflect.Message {
	mi := &file_webflow_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rewrite.ProtoReflect.Descriptor instead.
func (*Rewrite) Descriptor() ([]byte, []int) {
	return file_webflow_proto_rawDescGZIP(), []int{9}
}

func (x *Rewrite) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Rewrite) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Rewrite) GetWhen() *wrapperspb.BoolValue {
	if x != nil {
		return x.When
	}
	return nil
}

var File_webflow_proto protoreflect.FileDescriptor

var file_webflow_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x62, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x36, 0x36, 0x2e, 0x77, 0x65, 0x62, 0x66, 0x6c, 0x6f, 0x77,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x62, 0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x22, 0x7f, 0x0a, 0x02, 0x4f, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x61, 0x74, 0x63, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6d, 0x69,
	0x6e, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x4d, 0x69, 0x6e, 0x6f, 0x72, 0x22, 0x4c, 0x0a, 0x06, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x22, 0x92, 0x01, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x36, 0x36, 0x2e, 0x77, 0x65, 0x62, 0x66,
	0x6c, 0x6f, 0x77, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x02, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x36, 0x36, 0x2e, 0x77, 0x65, 0x62, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x4f, 0x53, 0x52, 0x02, 0x6f, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x36,
	0x36, 0x2e, 0x77, 0x65, 0x62, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0xbc, 0x01, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x72, 0x61, 0x77, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x61, 0x77, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x36, 0x36, 0x2e, 0x77, 0x65, 0x62, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x22, 0xe7, 0x01, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x69, 0x6e, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x63, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74,
	0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x7a, 0x6f, 0x6e, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x73, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x73,
	0x6e, 0x22, 0x49, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x68,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x77, 0x69, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x77, 0x69,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xa8, 0x02, 0x0a,
	0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x03, 0x61, 0x64, 0x64, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x36, 0x36, 0x2e, 0x77,
	0x65, 0x62, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x64,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x61, 0x64, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x12, 0x32, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x36, 0x36, 0x2e, 0x77, 0x65, 0x62, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x03, 0x73, 0x65, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x1a, 0x36, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x36, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x72, 0x0a, 0x08, 0x52, 0x65, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x69, 0x74, 0x68, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x77, 0x69, 0x74, 0x68, 0x12, 0x2e, 0x0a, 0x04, 0x77,
	0x68, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x22, 0x5d, 0x0a, 0x07, 0x52,
	0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2e, 0x0a, 0x04, 0x77, 0x68,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x77, 0x68, 0x65, 0x6e, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x36, 0x36,
	0x2f, 0x77, 0x65, 0x62, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_webflow_proto_rawDescOnce sync.Once
	file_webflow_proto_rawDescData = file_webflow_proto_rawDesc
)

func file_webflow_proto_rawDescGZIP() []byte {
	file_webflow_proto_rawDescOnce.Do(func() {
		file_webflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_webflow_proto_rawDescData)
	})
	return file_webflow_proto_rawDescData
}

var file_webflow_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_webflow_proto_goTypes = []interface{}{
	(*Client)(nil),               // 0: cloud66.webflow.Client
	(*OS)(nil),                   // 1: cloud66.webflow.OS
	(*Device)(nil),               // 2: cloud66.webflow.Device
	(*UserAgent)(nil),            // 3: cloud66.webflow.UserAgent
	(*Request)(nil),              // 4: cloud66.webflow.Request
	(*Origin)(nil),               // 5: cloud66.webflow.Origin
	(*Block)(nil),                // 6: cloud66.webflow.Block
	(*Header)(nil),               // 7: cloud66.webflow.Header
	(*Redirect)(nil),             // 8: cloud66.webflow.Redirect
	(*Rewrite)(nil),              // 9: cloud66.webflow.Rewrite
	nil,                          // 10: cloud66.webflow.Header.AddEntry
	nil,                          // 11: cloud66.webflow.Header.SetEntry
	(*wrapperspb.BoolValue)(nil), // 12: google.protobuf.BoolValue
}
var file_webflow_proto_depIdxs = []int32{
	0,  // 0: cloud66.webflow.UserAgent.client:type_name -> cloud66.webflow.Client
	1,  // 1: cloud66.webflow.UserAgent.os:type_name -> cloud66.webflow.OS
	2,  // 2: cloud66.webflow.UserAgent.device:type_name -> cloud66.webflow.Device
	3,  // 3: cloud66.webflow.Request.user_agent:type_name -> cloud66.webflow.UserAgent
	10, // 4: cloud66.webflow.Header.add:type_name -> cloud66.webflow.Header.AddEntry
	11, // 5: cloud66.webflow.Header.set:type_name -> cloud66.webflow.Header.SetEntry
	12, // 6: cloud66.webflow.Header.when:type_name -> google.protobuf.BoolValue
	12, // 7: cloud66.webflow.Redirect.when:type_name -> google.protobuf.BoolValue
	12, // 8: cloud66.webflow.Rewrite.when:type_name -> google.protobuf.BoolValue
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_webflow_proto_init() }
func file_webflow_proto_init() {
	if File_webflow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_webflow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Client); i {
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
		file_webflow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OS); i {
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
		file_webflow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Device); i {
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
		file_webflow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAgent); i {
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
		file_webflow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_webflow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Origin); i {
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
		file_webflow_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_webflow_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_webflow_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Redirect); i {
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
		file_webflow_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rewrite); i {
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
			RawDescriptor: file_webflow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_webflow_proto_goTypes,
		DependencyIndexes: file_webflow_proto_depIdxs,
		MessageInfos:      file_webflow_proto_msgTypes,
	}.Build()
	File_webflow_proto = out.File
	file_webflow_proto_rawDesc = nil
	file_webflow_proto_goTypes = nil
	file_webflow_proto_depIdxs = nil
}

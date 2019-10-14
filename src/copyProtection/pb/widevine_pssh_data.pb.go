// Code generated by protoc-gen-go. DO NOT EDIT.
// source: widevine_pssh_data.proto

package pb

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type WidevinePsshData_Algorithm int32

const (
	WidevinePsshData_UNENCRYPTED WidevinePsshData_Algorithm = 0
	WidevinePsshData_AESCTR      WidevinePsshData_Algorithm = 1
)

var WidevinePsshData_Algorithm_name = map[int32]string{
	0: "UNENCRYPTED",
	1: "AESCTR",
}

var WidevinePsshData_Algorithm_value = map[string]int32{
	"UNENCRYPTED": 0,
	"AESCTR":      1,
}

func (x WidevinePsshData_Algorithm) Enum() *WidevinePsshData_Algorithm {
	p := new(WidevinePsshData_Algorithm)
	*p = x
	return p
}

func (x WidevinePsshData_Algorithm) String() string {
	return proto.EnumName(WidevinePsshData_Algorithm_name, int32(x))
}

func (x *WidevinePsshData_Algorithm) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(WidevinePsshData_Algorithm_value, data, "WidevinePsshData_Algorithm")
	if err != nil {
		return err
	}
	*x = WidevinePsshData_Algorithm(value)
	return nil
}

func (WidevinePsshData_Algorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5be6ceb5afc57f18, []int{0, 0}
}

type WidevinePsshData struct {
	Algorithm *WidevinePsshData_Algorithm `protobuf:"varint,1,opt,name=algorithm,enum=widevine.WidevinePsshData_Algorithm" json:"algorithm,omitempty"`
	KeyId     [][]byte                    `protobuf:"bytes,2,rep,name=key_id,json=keyId" json:"key_id,omitempty"`
	// Content provider name.
	Provider *string `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	// A content identifier, specified by content provider.
	ContentId []byte `protobuf:"bytes,4,opt,name=content_id,json=contentId" json:"content_id,omitempty"`
	// The name of a registered policy to be used for this asset.
	Policy *string `protobuf:"bytes,6,opt,name=policy" json:"policy,omitempty"`
	// Crypto period index, for media using key rotation.
	CryptoPeriodIndex *uint32 `protobuf:"varint,7,opt,name=crypto_period_index,json=cryptoPeriodIndex" json:"crypto_period_index,omitempty"`
	// Optional protected context for group content. The grouped_license is a
	// serialized SignedMessage.
	GroupedLicense []byte `protobuf:"bytes,8,opt,name=grouped_license,json=groupedLicense" json:"grouped_license,omitempty"`
	// Protection scheme identifying the encryption algorithm. Represented as one
	// of the following 4CC values: 'cenc' (AES-CTR), 'cbc1' (AES-CBC),
	// 'cens' (AES-CTR subsample), 'cbcs' (AES-CBC subsample).
	ProtectionScheme     *uint32  `protobuf:"varint,9,opt,name=protection_scheme,json=protectionScheme" json:"protection_scheme,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WidevinePsshData) Reset()         { *m = WidevinePsshData{} }
func (m *WidevinePsshData) String() string { return proto.CompactTextString(m) }
func (*WidevinePsshData) ProtoMessage()    {}
func (*WidevinePsshData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be6ceb5afc57f18, []int{0}
}

func (m *WidevinePsshData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WidevinePsshData.Unmarshal(m, b)
}
func (m *WidevinePsshData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WidevinePsshData.Marshal(b, m, deterministic)
}
func (m *WidevinePsshData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WidevinePsshData.Merge(m, src)
}
func (m *WidevinePsshData) XXX_Size() int {
	return xxx_messageInfo_WidevinePsshData.Size(m)
}
func (m *WidevinePsshData) XXX_DiscardUnknown() {
	xxx_messageInfo_WidevinePsshData.DiscardUnknown(m)
}

var xxx_messageInfo_WidevinePsshData proto.InternalMessageInfo

func (m *WidevinePsshData) GetAlgorithm() WidevinePsshData_Algorithm {
	if m != nil && m.Algorithm != nil {
		return *m.Algorithm
	}
	return WidevinePsshData_UNENCRYPTED
}

func (m *WidevinePsshData) GetKeyId() [][]byte {
	if m != nil {
		return m.KeyId
	}
	return nil
}

func (m *WidevinePsshData) GetProvider() string {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return ""
}

func (m *WidevinePsshData) GetContentId() []byte {
	if m != nil {
		return m.ContentId
	}
	return nil
}

func (m *WidevinePsshData) GetPolicy() string {
	if m != nil && m.Policy != nil {
		return *m.Policy
	}
	return ""
}

func (m *WidevinePsshData) GetCryptoPeriodIndex() uint32 {
	if m != nil && m.CryptoPeriodIndex != nil {
		return *m.CryptoPeriodIndex
	}
	return 0
}

func (m *WidevinePsshData) GetGroupedLicense() []byte {
	if m != nil {
		return m.GroupedLicense
	}
	return nil
}

func (m *WidevinePsshData) GetProtectionScheme() uint32 {
	if m != nil && m.ProtectionScheme != nil {
		return *m.ProtectionScheme
	}
	return 0
}

// Derived from WidevinePsshData. The JSON format of this proto is used in
// Widevine HLS DRM signaling v1.
// We cannot build JSON from WidevinePsshData as |key_id| is required to be in
// hex format, while |bytes| type is translated to base64 by JSON formatter, so
// we have to use |string| type and do hex conversion in the code.
type WidevineHeader struct {
	KeyIds []string `protobuf:"bytes,2,rep,name=key_ids,json=keyIds" json:"key_ids,omitempty"`
	// Content provider name.
	Provider *string `protobuf:"bytes,3,opt,name=provider" json:"provider,omitempty"`
	// A content identifier, specified by content provider.
	ContentId            []byte   `protobuf:"bytes,4,opt,name=content_id,json=contentId" json:"content_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WidevineHeader) Reset()         { *m = WidevineHeader{} }
func (m *WidevineHeader) String() string { return proto.CompactTextString(m) }
func (*WidevineHeader) ProtoMessage()    {}
func (*WidevineHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_5be6ceb5afc57f18, []int{1}
}

func (m *WidevineHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WidevineHeader.Unmarshal(m, b)
}
func (m *WidevineHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WidevineHeader.Marshal(b, m, deterministic)
}
func (m *WidevineHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WidevineHeader.Merge(m, src)
}
func (m *WidevineHeader) XXX_Size() int {
	return xxx_messageInfo_WidevineHeader.Size(m)
}
func (m *WidevineHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_WidevineHeader.DiscardUnknown(m)
}

var xxx_messageInfo_WidevineHeader proto.InternalMessageInfo

func (m *WidevineHeader) GetKeyIds() []string {
	if m != nil {
		return m.KeyIds
	}
	return nil
}

func (m *WidevineHeader) GetProvider() string {
	if m != nil && m.Provider != nil {
		return *m.Provider
	}
	return ""
}

func (m *WidevineHeader) GetContentId() []byte {
	if m != nil {
		return m.ContentId
	}
	return nil
}

func init() {
	proto.RegisterEnum("widevine.WidevinePsshData_Algorithm", WidevinePsshData_Algorithm_name, WidevinePsshData_Algorithm_value)
	proto.RegisterType((*WidevinePsshData)(nil), "widevine.WidevinePsshData")
	proto.RegisterType((*WidevineHeader)(nil), "widevine.WidevineHeader")
}

func init() { proto.RegisterFile("widevine_pssh_data.proto", fileDescriptor_5be6ceb5afc57f18) }

var fileDescriptor_5be6ceb5afc57f18 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0x61, 0x4b, 0xc2, 0x40,
	0x18, 0xc7, 0x9b, 0xd6, 0x74, 0x4f, 0xa6, 0xf3, 0xa2, 0x3a, 0x82, 0x60, 0x48, 0xd0, 0x20, 0xd8,
	0x8b, 0xbe, 0x81, 0xa9, 0x90, 0x10, 0x22, 0xa7, 0x11, 0xbd, 0x1a, 0x63, 0xf7, 0xe0, 0x0e, 0x75,
	0x37, 0x76, 0x97, 0xb5, 0x6f, 0xdd, 0x47, 0x88, 0x9d, 0x9b, 0x42, 0x6f, 0x7b, 0xf9, 0xff, 0xff,
	0x9e, 0xff, 0xc6, 0xfd, 0x80, 0x7e, 0x09, 0x8e, 0x3b, 0x91, 0x62, 0x98, 0x29, 0x95, 0x84, 0x3c,
	0xd2, 0x51, 0x90, 0xe5, 0x52, 0x4b, 0xd2, 0xae, 0xc9, 0xe0, 0xa7, 0x01, 0xee, 0x7b, 0x15, 0xe6,
	0x4a, 0x25, 0xe3, 0x48, 0x47, 0xe4, 0x19, 0x9c, 0x68, 0xb3, 0x92, 0xb9, 0xd0, 0xc9, 0x96, 0x5a,
	0x9e, 0xe5, 0x77, 0x9f, 0xee, 0x83, 0x7a, 0x12, 0xfc, 0x3d, 0x0f, 0x86, 0xf5, 0x2d, 0x3b, 0xce,
	0xc8, 0x15, 0xd8, 0x6b, 0x2c, 0x42, 0xc1, 0x69, 0xc3, 0x6b, 0xfa, 0x1d, 0x76, 0xb6, 0xc6, 0x62,
	0xca, 0xc9, 0x2d, 0xb4, 0xb3, 0x5c, 0xee, 0x04, 0xc7, 0x9c, 0x36, 0x3d, 0xcb, 0x77, 0xd8, 0x21,
	0x93, 0x3b, 0x80, 0x58, 0xa6, 0x1a, 0x53, 0x5d, 0xce, 0x4e, 0x3d, 0xcb, 0xef, 0x30, 0xa7, 0x6a,
	0xa6, 0x9c, 0x5c, 0x83, 0x9d, 0xc9, 0x8d, 0x88, 0x0b, 0x6a, 0x9b, 0x61, 0x95, 0x48, 0x00, 0x97,
	0x71, 0x5e, 0x64, 0x5a, 0x86, 0x19, 0xe6, 0x42, 0xf2, 0x50, 0xa4, 0x1c, 0xbf, 0x69, 0xcb, 0xb3,
	0xfc, 0x0b, 0xd6, 0xdf, 0xa3, 0xb9, 0x21, 0xd3, 0x12, 0x90, 0x07, 0xe8, 0xad, 0x72, 0xf9, 0x99,
	0x21, 0x0f, 0x37, 0x22, 0xc6, 0x54, 0x21, 0x6d, 0x9b, 0x7f, 0x75, 0xab, 0xfa, 0x75, 0xdf, 0x92,
	0x47, 0xe8, 0x97, 0xba, 0x30, 0xd6, 0x42, 0xa6, 0xa1, 0x8a, 0x13, 0xdc, 0x22, 0x75, 0xcc, 0x67,
	0xdd, 0x23, 0x58, 0x98, 0x7e, 0xe0, 0x83, 0x73, 0xf0, 0x40, 0x7a, 0x70, 0xfe, 0x36, 0x9b, 0xcc,
	0x46, 0xec, 0x63, 0xbe, 0x9c, 0x8c, 0xdd, 0x13, 0x02, 0x60, 0x0f, 0x27, 0x8b, 0xd1, 0x92, 0xb9,
	0xd6, 0x80, 0x43, 0xb7, 0x56, 0xf8, 0x82, 0x51, 0xf9, 0xf0, 0x1b, 0x68, 0xed, 0x5d, 0x29, 0x23,
	0xcb, 0x61, 0xb6, 0x91, 0xa5, 0xfe, 0x61, 0xeb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa7, 0x2a, 0x15,
	0x98, 0xfd, 0x01, 0x00, 0x00,
}
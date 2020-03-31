// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/scanner/service.proto

package scanner

import (
	fmt "fmt"
	common "github.com/aquasecurity/trivy/rpc/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type ScanRequest struct {
	Target               string       `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	ImageId              string       `protobuf:"bytes,2,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	LayerIds             []string     `protobuf:"bytes,3,rep,name=layer_ids,json=layerIds,proto3" json:"layer_ids,omitempty"`
	Options              *ScanOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ScanRequest) Reset()         { *m = ScanRequest{} }
func (m *ScanRequest) String() string { return proto.CompactTextString(m) }
func (*ScanRequest) ProtoMessage()    {}
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_60d0e837512b18d4, []int{0}
}

func (m *ScanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanRequest.Unmarshal(m, b)
}
func (m *ScanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanRequest.Marshal(b, m, deterministic)
}
func (m *ScanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanRequest.Merge(m, src)
}
func (m *ScanRequest) XXX_Size() int {
	return xxx_messageInfo_ScanRequest.Size(m)
}
func (m *ScanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScanRequest proto.InternalMessageInfo

func (m *ScanRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *ScanRequest) GetImageId() string {
	if m != nil {
		return m.ImageId
	}
	return ""
}

func (m *ScanRequest) GetLayerIds() []string {
	if m != nil {
		return m.LayerIds
	}
	return nil
}

func (m *ScanRequest) GetOptions() *ScanOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type ScanOptions struct {
	VulnType             []string `protobuf:"bytes,1,rep,name=vuln_type,json=vulnType,proto3" json:"vuln_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScanOptions) Reset()         { *m = ScanOptions{} }
func (m *ScanOptions) String() string { return proto.CompactTextString(m) }
func (*ScanOptions) ProtoMessage()    {}
func (*ScanOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_60d0e837512b18d4, []int{1}
}

func (m *ScanOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanOptions.Unmarshal(m, b)
}
func (m *ScanOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanOptions.Marshal(b, m, deterministic)
}
func (m *ScanOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanOptions.Merge(m, src)
}
func (m *ScanOptions) XXX_Size() int {
	return xxx_messageInfo_ScanOptions.Size(m)
}
func (m *ScanOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ScanOptions proto.InternalMessageInfo

func (m *ScanOptions) GetVulnType() []string {
	if m != nil {
		return m.VulnType
	}
	return nil
}

type ScanResponse struct {
	Os                   *common.OS `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	Eosl                 bool       `protobuf:"varint,2,opt,name=eosl,proto3" json:"eosl,omitempty"`
	Results              []*Result  `protobuf:"bytes,3,rep,name=results,proto3" json:"results,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ScanResponse) Reset()         { *m = ScanResponse{} }
func (m *ScanResponse) String() string { return proto.CompactTextString(m) }
func (*ScanResponse) ProtoMessage()    {}
func (*ScanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_60d0e837512b18d4, []int{2}
}

func (m *ScanResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanResponse.Unmarshal(m, b)
}
func (m *ScanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanResponse.Marshal(b, m, deterministic)
}
func (m *ScanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanResponse.Merge(m, src)
}
func (m *ScanResponse) XXX_Size() int {
	return xxx_messageInfo_ScanResponse.Size(m)
}
func (m *ScanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ScanResponse proto.InternalMessageInfo

func (m *ScanResponse) GetOs() *common.OS {
	if m != nil {
		return m.Os
	}
	return nil
}

func (m *ScanResponse) GetEosl() bool {
	if m != nil {
		return m.Eosl
	}
	return false
}

func (m *ScanResponse) GetResults() []*Result {
	if m != nil {
		return m.Results
	}
	return nil
}

// Result is the same as github.com/aquasecurity/trivy/pkg/report.Result
type Result struct {
	Target               string                  `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	Vulnerabilities      []*common.Vulnerability `protobuf:"bytes,2,rep,name=vulnerabilities,proto3" json:"vulnerabilities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_60d0e837512b18d4, []int{3}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Result) GetVulnerabilities() []*common.Vulnerability {
	if m != nil {
		return m.Vulnerabilities
	}
	return nil
}

func init() {
	proto.RegisterType((*ScanRequest)(nil), "trivy.scanner.v1.ScanRequest")
	proto.RegisterType((*ScanOptions)(nil), "trivy.scanner.v1.ScanOptions")
	proto.RegisterType((*ScanResponse)(nil), "trivy.scanner.v1.ScanResponse")
	proto.RegisterType((*Result)(nil), "trivy.scanner.v1.Result")
}

func init() { proto.RegisterFile("rpc/scanner/service.proto", fileDescriptor_60d0e837512b18d4) }

var fileDescriptor_60d0e837512b18d4 = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x6b, 0xbb, 0x30,
	0x18, 0x45, 0x5b, 0x6a, 0x8d, 0x3f, 0xf8, 0x95, 0x1c, 0x86, 0x6d, 0xd9, 0x10, 0x4f, 0x65, 0x07,
	0x65, 0x0e, 0xb6, 0xfb, 0xa0, 0x87, 0x9e, 0x3a, 0xd2, 0xb1, 0xc3, 0x2e, 0x25, 0xd5, 0x0f, 0x17,
	0x50, 0x63, 0x93, 0x28, 0xf3, 0x1f, 0xd9, 0xdf, 0x3b, 0x4c, 0x2c, 0xac, 0x1d, 0xbd, 0x25, 0xef,
	0x3d, 0xbf, 0xf7, 0xde, 0x67, 0xd0, 0x5c, 0xd4, 0x69, 0x2c, 0x53, 0x5a, 0x55, 0x20, 0x62, 0x09,
	0xa2, 0x65, 0x29, 0x44, 0xb5, 0xe0, 0x8a, 0xe3, 0x99, 0x12, 0xac, 0xed, 0xa2, 0x81, 0x8c, 0xda,
	0x87, 0xc5, 0x53, 0xce, 0xd4, 0x67, 0x73, 0x88, 0x52, 0x5e, 0xc6, 0xf4, 0xd8, 0x50, 0x09, 0x69,
	0x23, 0x98, 0xea, 0x62, 0xad, 0x8c, 0xfb, 0x51, 0x29, 0x2f, 0x4b, 0x5e, 0x9d, 0x4f, 0x0a, 0xbf,
	0x2d, 0xe4, 0xed, 0x52, 0x5a, 0x11, 0x38, 0x36, 0x20, 0x15, 0xbe, 0x41, 0x13, 0x45, 0x45, 0x0e,
	0xca, 0xb7, 0x02, 0x6b, 0xe5, 0x92, 0xe1, 0x86, 0xe7, 0x68, 0xca, 0x4a, 0x9a, 0xc3, 0x9e, 0x65,
	0xbe, 0xad, 0x19, 0x47, 0xdf, 0x37, 0x19, 0x5e, 0x22, 0xb7, 0xa0, 0x1d, 0x88, 0x3d, 0xcb, 0xa4,
	0x3f, 0x0a, 0x46, 0x2b, 0x97, 0x4c, 0x35, 0xb0, 0xc9, 0x24, 0x7e, 0x46, 0x0e, 0xaf, 0x15, 0xe3,
	0x95, 0xf4, 0xc7, 0x81, 0xb5, 0xf2, 0x92, 0xdb, 0xe8, 0x32, 0x7b, 0xd4, 0xfb, 0x6f, 0x8d, 0x88,
	0x9c, 0xd4, 0xe1, 0xbd, 0xc9, 0x35, 0xe0, 0xbd, 0x49, 0xdb, 0x14, 0xd5, 0x5e, 0x75, 0x35, 0xf8,
	0x96, 0x31, 0xe9, 0x81, 0xb7, 0xae, 0x86, 0xf0, 0x0b, 0xfd, 0x33, 0x1d, 0x64, 0xcd, 0x2b, 0x09,
	0x38, 0x40, 0x36, 0x97, 0xba, 0x80, 0x97, 0xcc, 0x06, 0x3f, 0xd3, 0x3e, 0xda, 0xee, 0x88, 0xcd,
	0x25, 0xc6, 0x68, 0x0c, 0x5c, 0x16, 0xba, 0xca, 0x94, 0xe8, 0x33, 0x4e, 0x90, 0x23, 0x40, 0x36,
	0x85, 0x32, 0x2d, 0xbc, 0xc4, 0xff, 0x1b, 0x95, 0x68, 0x01, 0x39, 0x09, 0xc3, 0x1c, 0x4d, 0x0c,
	0x74, 0x75, 0x71, 0x6b, 0xf4, 0xbf, 0xcf, 0x09, 0x82, 0x1e, 0x58, 0xc1, 0x14, 0x03, 0xe9, 0xdb,
	0x7a, 0xfa, 0xf2, 0x3c, 0xd8, 0xfb, 0x2f, 0x51, 0x47, 0x2e, 0xbf, 0x49, 0x5e, 0x91, 0xb3, 0x33,
	0x31, 0xf0, 0x1a, 0x8d, 0xfb, 0x23, 0xbe, 0xb2, 0xc9, 0xe1, 0x4f, 0x2e, 0xee, 0xae, 0xd1, 0x66,
	0x49, 0x2f, 0xee, 0x87, 0x33, 0x50, 0x87, 0x89, 0x7e, 0x0b, 0x8f, 0x3f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x22, 0xf3, 0x23, 0xc0, 0x72, 0x02, 0x00, 0x00,
}
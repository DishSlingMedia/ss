// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sandstorm.proto

package sandstormpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SandstormEvent struct {
	Product                       string               `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Platform                      string               `protobuf:"bytes,2,opt,name=platform,proto3" json:"platform,omitempty"`
	AppVersion                    string               `protobuf:"bytes,3,opt,name=app_version,json=appVersion,proto3" json:"app_version,omitempty"`
	BuildInfo                     string               `protobuf:"bytes,4,opt,name=build_info,json=buildInfo,proto3" json:"build_info,omitempty"`
	Url                           string               `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
	TestFramework                 string               `protobuf:"bytes,6,opt,name=test_framework,json=testFramework,proto3" json:"test_framework,omitempty"`
	DeviceCount                   int32                `protobuf:"varint,7,opt,name=device_count,json=deviceCount,proto3" json:"device_count,omitempty"`
	Location                      string               `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	TriggeredBy                   string               `protobuf:"bytes,9,opt,name=triggered_by,json=triggeredBy,proto3" json:"triggered_by,omitempty"`
	RunType                       string               `protobuf:"bytes,10,opt,name=run_type,json=runType,proto3" json:"run_type,omitempty"`
	Environment                   string               `protobuf:"bytes,11,opt,name=environment,proto3" json:"environment,omitempty"`
	Suites                        string               `protobuf:"bytes,12,opt,name=suites,proto3" json:"suites,omitempty"`
	LocalTime                     *timestamp.Timestamp `protobuf:"bytes,13,opt,name=local_time,json=localTime,proto3" json:"local_time,omitempty"`
	AutomationRepoVersion         string               `protobuf:"bytes,14,opt,name=automation_repo_version,json=automationRepoVersion,proto3" json:"automation_repo_version,omitempty"`
	TestCasesInInput              int32                `protobuf:"varint,15,opt,name=test_cases_in_input,json=testCasesInInput,proto3" json:"test_cases_in_input,omitempty"`
	TestCasesNotExecuted          int32                `protobuf:"varint,16,opt,name=test_cases_not_executed,json=testCasesNotExecuted,proto3" json:"test_cases_not_executed,omitempty"`
	TestCasesExecuted             int32                `protobuf:"varint,17,opt,name=test_cases_executed,json=testCasesExecuted,proto3" json:"test_cases_executed,omitempty"`
	TotalTestCaseSkipped          int32                `protobuf:"varint,18,opt,name=total_test_case_skipped,json=totalTestCaseSkipped,proto3" json:"total_test_case_skipped,omitempty"`
	TestCasePassed                int32                `protobuf:"varint,19,opt,name=test_case_passed,json=testCasePassed,proto3" json:"test_case_passed,omitempty"`
	TestPassedInRetry             int32                `protobuf:"varint,20,opt,name=test_passed_in_retry,json=testPassedInRetry,proto3" json:"test_passed_in_retry,omitempty"`
	TestCaseFailed                int32                `protobuf:"varint,21,opt,name=test_case_failed,json=testCaseFailed,proto3" json:"test_case_failed,omitempty"`
	TestCaseFailedDueProductBugs  int32                `protobuf:"varint,22,opt,name=test_case_failed_due_product_bugs,json=testCaseFailedDueProductBugs,proto3" json:"test_case_failed_due_product_bugs,omitempty"`
	TestCaseFailuresYetClassified int32                `protobuf:"varint,23,opt,name=test_case_failures_yet_classified,json=testCaseFailuresYetClassified,proto3" json:"test_case_failures_yet_classified,omitempty"`
	TotalTimeTakenForRunInSecs    int32                `protobuf:"varint,24,opt,name=total_time_taken_for_run_in_secs,json=totalTimeTakenForRunInSecs,proto3" json:"total_time_taken_for_run_in_secs,omitempty"`
	ServerTimestamp               *timestamp.Timestamp `protobuf:"bytes,25,opt,name=server_timestamp,json=serverTimestamp,proto3" json:"server_timestamp,omitempty"`
	TotalTestCaseInZephyr         int32                `protobuf:"varint,26,opt,name=total_test_case_in_zephyr,json=totalTestCaseInZephyr,proto3" json:"total_test_case_in_zephyr,omitempty"`
	JobType                       string               `protobuf:"bytes,27,opt,name=job_type,json=jobType,proto3" json:"job_type,omitempty"`
	XXX_NoUnkeyedLiteral          struct{}             `json:"-"`
	XXX_unrecognized              []byte               `json:"-"`
	XXX_sizecache                 int32                `json:"-"`
}

func (m *SandstormEvent) Reset()         { *m = SandstormEvent{} }
func (m *SandstormEvent) String() string { return proto.CompactTextString(m) }
func (*SandstormEvent) ProtoMessage()    {}
func (*SandstormEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b845f3a62ba3779, []int{0}
}

func (m *SandstormEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SandstormEvent.Unmarshal(m, b)
}
func (m *SandstormEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SandstormEvent.Marshal(b, m, deterministic)
}
func (m *SandstormEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SandstormEvent.Merge(m, src)
}
func (m *SandstormEvent) XXX_Size() int {
	return xxx_messageInfo_SandstormEvent.Size(m)
}
func (m *SandstormEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SandstormEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SandstormEvent proto.InternalMessageInfo

func (m *SandstormEvent) GetProduct() string {
	if m != nil {
		return m.Product
	}
	return ""
}

func (m *SandstormEvent) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *SandstormEvent) GetAppVersion() string {
	if m != nil {
		return m.AppVersion
	}
	return ""
}

func (m *SandstormEvent) GetBuildInfo() string {
	if m != nil {
		return m.BuildInfo
	}
	return ""
}

func (m *SandstormEvent) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SandstormEvent) GetTestFramework() string {
	if m != nil {
		return m.TestFramework
	}
	return ""
}

func (m *SandstormEvent) GetDeviceCount() int32 {
	if m != nil {
		return m.DeviceCount
	}
	return 0
}

func (m *SandstormEvent) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *SandstormEvent) GetTriggeredBy() string {
	if m != nil {
		return m.TriggeredBy
	}
	return ""
}

func (m *SandstormEvent) GetRunType() string {
	if m != nil {
		return m.RunType
	}
	return ""
}

func (m *SandstormEvent) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *SandstormEvent) GetSuites() string {
	if m != nil {
		return m.Suites
	}
	return ""
}

func (m *SandstormEvent) GetLocalTime() *timestamp.Timestamp {
	if m != nil {
		return m.LocalTime
	}
	return nil
}

func (m *SandstormEvent) GetAutomationRepoVersion() string {
	if m != nil {
		return m.AutomationRepoVersion
	}
	return ""
}

func (m *SandstormEvent) GetTestCasesInInput() int32 {
	if m != nil {
		return m.TestCasesInInput
	}
	return 0
}

func (m *SandstormEvent) GetTestCasesNotExecuted() int32 {
	if m != nil {
		return m.TestCasesNotExecuted
	}
	return 0
}

func (m *SandstormEvent) GetTestCasesExecuted() int32 {
	if m != nil {
		return m.TestCasesExecuted
	}
	return 0
}

func (m *SandstormEvent) GetTotalTestCaseSkipped() int32 {
	if m != nil {
		return m.TotalTestCaseSkipped
	}
	return 0
}

func (m *SandstormEvent) GetTestCasePassed() int32 {
	if m != nil {
		return m.TestCasePassed
	}
	return 0
}

func (m *SandstormEvent) GetTestPassedInRetry() int32 {
	if m != nil {
		return m.TestPassedInRetry
	}
	return 0
}

func (m *SandstormEvent) GetTestCaseFailed() int32 {
	if m != nil {
		return m.TestCaseFailed
	}
	return 0
}

func (m *SandstormEvent) GetTestCaseFailedDueProductBugs() int32 {
	if m != nil {
		return m.TestCaseFailedDueProductBugs
	}
	return 0
}

func (m *SandstormEvent) GetTestCaseFailuresYetClassified() int32 {
	if m != nil {
		return m.TestCaseFailuresYetClassified
	}
	return 0
}

func (m *SandstormEvent) GetTotalTimeTakenForRunInSecs() int32 {
	if m != nil {
		return m.TotalTimeTakenForRunInSecs
	}
	return 0
}

func (m *SandstormEvent) GetServerTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.ServerTimestamp
	}
	return nil
}

func (m *SandstormEvent) GetTotalTestCaseInZephyr() int32 {
	if m != nil {
		return m.TotalTestCaseInZephyr
	}
	return 0
}

func (m *SandstormEvent) GetJobType() string {
	if m != nil {
		return m.JobType
	}
	return ""
}

type SandstormEventResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SandstormEventResponse) Reset()         { *m = SandstormEventResponse{} }
func (m *SandstormEventResponse) String() string { return proto.CompactTextString(m) }
func (*SandstormEventResponse) ProtoMessage()    {}
func (*SandstormEventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b845f3a62ba3779, []int{1}
}

func (m *SandstormEventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SandstormEventResponse.Unmarshal(m, b)
}
func (m *SandstormEventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SandstormEventResponse.Marshal(b, m, deterministic)
}
func (m *SandstormEventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SandstormEventResponse.Merge(m, src)
}
func (m *SandstormEventResponse) XXX_Size() int {
	return xxx_messageInfo_SandstormEventResponse.Size(m)
}
func (m *SandstormEventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SandstormEventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SandstormEventResponse proto.InternalMessageInfo

func (m *SandstormEventResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*SandstormEvent)(nil), "sandstormpb.SandstormEvent")
	proto.RegisterType((*SandstormEventResponse)(nil), "sandstormpb.SandstormEventResponse")
}

func init() { proto.RegisterFile("sandstorm.proto", fileDescriptor_5b845f3a62ba3779) }

var fileDescriptor_5b845f3a62ba3779 = []byte{
	// 802 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0x96, 0x13, 0xb2, 0x59, 0xf7, 0x64, 0x6d, 0xa7, 0x37, 0xbb, 0xdb, 0xeb, 0x24, 0xc4, 0x31,
	0x42, 0xb2, 0x90, 0x62, 0x43, 0x10, 0x08, 0x72, 0x41, 0xda, 0x3f, 0xf0, 0x05, 0xad, 0xc6, 0x16,
	0x12, 0xb9, 0xb4, 0x7a, 0x66, 0x6a, 0x9c, 0xce, 0xce, 0x74, 0x37, 0xfd, 0x63, 0x30, 0xc7, 0xbc,
	0x02, 0x2f, 0xc4, 0x3b, 0xf0, 0x0a, 0x9c, 0x78, 0x0a, 0xd4, 0x3d, 0x3f, 0x6b, 0xef, 0x01, 0x6e,
	0xae, 0xef, 0xa7, 0xba, 0x5c, 0x35, 0x55, 0xa8, 0x6f, 0x98, 0xc8, 0x8c, 0x95, 0xba, 0x9c, 0x2a,
	0x2d, 0xad, 0xc4, 0x51, 0x0b, 0xa8, 0x64, 0xf8, 0x62, 0x25, 0xe5, 0xaa, 0x80, 0x59, 0xa0, 0x12,
	0x97, 0xcf, 0x2c, 0x2f, 0xc1, 0x58, 0x56, 0xaa, 0x4a, 0x3d, 0x7c, 0x56, 0x0b, 0x98, 0xe2, 0x33,
	0x26, 0x84, 0xb4, 0xcc, 0x72, 0x29, 0x4c, 0xcd, 0xf6, 0x92, 0x5f, 0xa8, 0x65, 0x49, 0x01, 0x55,
	0x3c, 0xfe, 0xb3, 0x8b, 0x7a, 0x8b, 0x26, 0xfd, 0xe5, 0x1a, 0x84, 0xc5, 0x04, 0x3d, 0x54, 0x5a,
	0x66, 0x2e, 0xb5, 0xa4, 0x33, 0xea, 0x4c, 0xba, 0x71, 0x13, 0xe2, 0x21, 0xda, 0x57, 0x05, 0xb3,
	0xb9, 0xd4, 0x25, 0xb9, 0x17, 0xa8, 0x36, 0xc6, 0x2f, 0x50, 0xc4, 0x94, 0xa2, 0x6b, 0xd0, 0x86,
	0x4b, 0x41, 0xee, 0x07, 0x1a, 0x31, 0xa5, 0x7e, 0xaa, 0x10, 0xfc, 0x1c, 0xa1, 0xc4, 0xf1, 0x22,
	0xa3, 0x5c, 0xe4, 0x92, 0x7c, 0x14, 0xf8, 0x6e, 0x40, 0xe6, 0x22, 0x97, 0x78, 0x80, 0xee, 0x3b,
	0x5d, 0x90, 0x07, 0x01, 0xf7, 0x3f, 0xf1, 0xa7, 0xa8, 0x67, 0xc1, 0x58, 0x9a, 0x6b, 0x56, 0xc2,
	0xaf, 0x52, 0xdf, 0x90, 0xbd, 0x40, 0x1e, 0x78, 0xf4, 0xaa, 0x01, 0xf1, 0x4b, 0xf4, 0x28, 0x83,
	0x35, 0x4f, 0x81, 0xa6, 0xd2, 0x09, 0x4b, 0x1e, 0x8e, 0x3a, 0x93, 0x07, 0x71, 0x54, 0x61, 0xe7,
	0x1e, 0xf2, 0x75, 0x17, 0x32, 0x0d, 0x7d, 0x20, 0xfb, 0x55, 0xdd, 0x4d, 0xec, 0xed, 0x56, 0xf3,
	0xd5, 0x0a, 0x34, 0x64, 0x34, 0xd9, 0x90, 0x6e, 0xe0, 0xa3, 0x16, 0x3b, 0xdb, 0xe0, 0x53, 0xb4,
	0xaf, 0x9d, 0xa0, 0x76, 0xa3, 0x80, 0xa0, 0xaa, 0x23, 0xda, 0x89, 0xe5, 0x46, 0x01, 0x1e, 0xa1,
	0x08, 0xc4, 0x9a, 0x6b, 0x29, 0x4a, 0x10, 0x96, 0x44, 0x95, 0x79, 0x0b, 0xc2, 0xc7, 0x68, 0xcf,
	0x38, 0x6e, 0xc1, 0x90, 0x47, 0x81, 0xac, 0x23, 0xfc, 0x2d, 0x42, 0xbe, 0x86, 0x82, 0xfa, 0xf9,
	0x91, 0x83, 0x51, 0x67, 0x12, 0xbd, 0x1e, 0x4e, 0xab, 0xd9, 0x4d, 0x9b, 0xe1, 0x4e, 0x97, 0xcd,
	0x70, 0xe3, 0x6e, 0x50, 0xfb, 0x18, 0x7f, 0x8d, 0x4e, 0x98, 0xb3, 0xb2, 0x0c, 0x7f, 0x80, 0x6a,
	0x50, 0xb2, 0x6d, 0x7b, 0x2f, 0xbc, 0x71, 0x74, 0x4b, 0xc7, 0xa0, 0x64, 0x33, 0x81, 0x57, 0xe8,
	0x30, 0x34, 0x34, 0x65, 0x06, 0x0c, 0xe5, 0x82, 0x72, 0xa1, 0x9c, 0x25, 0xfd, 0xd0, 0xb0, 0x81,
	0xa7, 0xce, 0x3d, 0x33, 0x17, 0x73, 0x8f, 0xe3, 0xaf, 0xd0, 0xc9, 0x96, 0x5c, 0x48, 0x4b, 0xe1,
	0x37, 0x48, 0x9d, 0x85, 0x8c, 0x0c, 0x82, 0xe5, 0x49, 0x6b, 0xf9, 0x51, 0xda, 0xcb, 0x9a, 0xc3,
	0xd3, 0x9d, 0x57, 0x5a, 0xcb, 0xe3, 0x60, 0x79, 0xdc, 0x5a, 0x5a, 0xbd, 0x7f, 0x46, 0x5a, 0xdf,
	0x88, 0xc6, 0x45, 0xcd, 0x0d, 0x57, 0x0a, 0x32, 0x82, 0xeb, 0x67, 0x3c, 0xbd, 0xac, 0x8d, 0x8b,
	0x8a, 0xc3, 0x13, 0x34, 0xb8, 0x35, 0x28, 0x66, 0x0c, 0x64, 0xe4, 0x30, 0xe8, 0x7b, 0xcd, 0x1b,
	0xd7, 0x01, 0xc5, 0x33, 0x14, 0x0a, 0xad, 0x45, 0xfe, 0x7f, 0x6b, 0xb0, 0x7a, 0x43, 0x9e, 0xdc,
	0x56, 0x54, 0x29, 0xe7, 0x22, 0xf6, 0xc4, 0x6e, 0xea, 0x9c, 0xf1, 0x02, 0x32, 0x72, 0xb4, 0x9b,
	0xfa, 0x2a, 0xa0, 0xf8, 0x7b, 0xf4, 0xf2, 0xae, 0x92, 0x66, 0x0e, 0x68, 0xbd, 0x30, 0x34, 0x71,
	0x2b, 0x43, 0x8e, 0x83, 0xf5, 0xd9, 0xae, 0xf5, 0xc2, 0xc1, 0x75, 0x25, 0x3a, 0x73, 0x2b, 0x83,
	0x7f, 0xb8, 0x9b, 0xc8, 0x69, 0x30, 0x74, 0x03, 0x96, 0xa6, 0x05, 0x33, 0x86, 0xe7, 0x1c, 0x32,
	0x72, 0x12, 0x12, 0x3d, 0xdf, 0x4e, 0xe4, 0x65, 0x3f, 0x83, 0x3d, 0x6f, 0x45, 0xf8, 0x02, 0x8d,
	0xea, 0x76, 0xf2, 0x12, 0xa8, 0x65, 0x37, 0x20, 0x68, 0x2e, 0x35, 0xf5, 0x5f, 0x30, 0x17, 0xd4,
	0x40, 0x6a, 0x08, 0x09, 0x89, 0x86, 0x55, 0x5f, 0x79, 0x09, 0x4b, 0xaf, 0xba, 0x92, 0x3a, 0x76,
	0x62, 0x2e, 0x16, 0x90, 0x1a, 0x7c, 0x89, 0x06, 0x06, 0xf4, 0x1a, 0x34, 0x6d, 0xcf, 0x0b, 0x39,
	0xfd, 0xdf, 0x6f, 0xb4, 0x5f, 0x79, 0x5a, 0x00, 0x7f, 0x83, 0x4e, 0xef, 0xce, 0x96, 0x0b, 0xfa,
	0x3b, 0xa8, 0x77, 0x1b, 0x4d, 0x86, 0xa1, 0x8a, 0xa3, 0x9d, 0xe9, 0xce, 0xc5, 0xdb, 0x40, 0xfa,
	0x9d, 0x7b, 0x2f, 0x93, 0x6a, 0xe7, 0x9e, 0x56, 0x3b, 0xf7, 0x5e, 0x26, 0x7e, 0xe7, 0xde, 0xf4,
	0xfe, 0xf9, 0x2e, 0x42, 0xdd, 0xf6, 0x26, 0x8e, 0x3f, 0x47, 0xc7, 0xbb, 0x17, 0x2c, 0x06, 0xa3,
	0xa4, 0x30, 0xe0, 0x77, 0x4f, 0x83, 0x71, 0x45, 0x73, 0xc8, 0xea, 0xe8, 0xf5, 0x87, 0x0e, 0xea,
	0x5f, 0x30, 0xcb, 0xae, 0xa5, 0xb1, 0x0b, 0xd0, 0xfe, 0x50, 0x60, 0x89, 0xa2, 0x18, 0x52, 0xe0,
	0x6b, 0xf0, 0x0c, 0x7e, 0x3a, 0xdd, 0x3a, 0xba, 0xd3, 0xdd, 0xfc, 0xc3, 0x4f, 0xfe, 0x83, 0x6c,
	0x1e, 0x1f, 0x7f, 0xfc, 0xe1, 0xaf, 0xbf, 0xff, 0xb8, 0x47, 0xc6, 0x87, 0xb3, 0xf5, 0x17, 0xaf,
	0x58, 0xa1, 0xde, 0xb1, 0x59, 0xeb, 0x7a, 0xd3, 0xf9, 0xec, 0xec, 0xe0, 0xed, 0xf6, 0x5d, 0x4f,
	0xf6, 0x42, 0x3f, 0xbf, 0xfc, 0x37, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xb2, 0x0f, 0x24, 0xfe, 0x05,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataPostServiceClient is the client API for DataPostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataPostServiceClient interface {
	ReceiveData(ctx context.Context, in *SandstormEvent, opts ...grpc.CallOption) (*SandstormEventResponse, error)
}

type dataPostServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataPostServiceClient(cc *grpc.ClientConn) DataPostServiceClient {
	return &dataPostServiceClient{cc}
}

func (c *dataPostServiceClient) ReceiveData(ctx context.Context, in *SandstormEvent, opts ...grpc.CallOption) (*SandstormEventResponse, error) {
	out := new(SandstormEventResponse)
	err := c.cc.Invoke(ctx, "/sandstormpb.DataPostService/ReceiveData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataPostServiceServer is the server API for DataPostService service.
type DataPostServiceServer interface {
	ReceiveData(context.Context, *SandstormEvent) (*SandstormEventResponse, error)
}

// UnimplementedDataPostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataPostServiceServer struct {
}

func (*UnimplementedDataPostServiceServer) ReceiveData(ctx context.Context, req *SandstormEvent) (*SandstormEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReceiveData not implemented")
}

func RegisterDataPostServiceServer(s *grpc.Server, srv DataPostServiceServer) {
	s.RegisterService(&_DataPostService_serviceDesc, srv)
}

func _DataPostService_ReceiveData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SandstormEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataPostServiceServer).ReceiveData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sandstormpb.DataPostService/ReceiveData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataPostServiceServer).ReceiveData(ctx, req.(*SandstormEvent))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataPostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sandstormpb.DataPostService",
	HandlerType: (*DataPostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReceiveData",
			Handler:    _DataPostService_ReceiveData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sandstorm.proto",
}

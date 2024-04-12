// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: capabilities/consensus/ocr3/types/ocr3_types.proto

package types

import (
	pb "github.com/smartcontractkit/chainlink-common/pkg/values/pb"
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

// per-workflow aggregation outcome
type AggregationOutcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId       string  `protobuf:"bytes,1,opt,name=workflowId,proto3" json:"workflowId,omitempty"`
	EncodableOutcome *pb.Map `protobuf:"bytes,2,opt,name=encodableOutcome,proto3" json:"encodableOutcome,omitempty"` // passed to encoders
	Metadata         []byte  `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`                 // internal to the aggregator
	ShouldReport     bool    `protobuf:"varint,4,opt,name=shouldReport,proto3" json:"shouldReport,omitempty"`
}

func (x *AggregationOutcome) Reset() {
	*x = AggregationOutcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregationOutcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregationOutcome) ProtoMessage() {}

func (x *AggregationOutcome) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregationOutcome.ProtoReflect.Descriptor instead.
func (*AggregationOutcome) Descriptor() ([]byte, []int) {
	return file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescGZIP(), []int{0}
}

func (x *AggregationOutcome) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *AggregationOutcome) GetEncodableOutcome() *pb.Map {
	if x != nil {
		return x.EncodableOutcome
	}
	return nil
}

func (x *AggregationOutcome) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AggregationOutcome) GetShouldReport() bool {
	if x != nil {
		return x.ShouldReport
	}
	return false
}

type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the requests to get consensus on.
	Ids []*Id `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescGZIP(), []int{1}
}

func (x *Query) GetIds() []*Id {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowExecutionId string `protobuf:"bytes,1,opt,name=workflowExecutionId,proto3" json:"workflowExecutionId,omitempty"`
	WorkflowId          string `protobuf:"bytes,2,opt,name=workflowId,proto3" json:"workflowId,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescGZIP(), []int{2}
}

func (x *Id) GetWorkflowExecutionId() string {
	if x != nil {
		return x.WorkflowExecutionId
	}
	return ""
}

func (x *Id) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

type Observation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *Id `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// list of observations defined in inputs.observations
	Observations *pb.List `protobuf:"bytes,4,opt,name=observations,proto3" json:"observations,omitempty"`
}

func (x *Observation) Reset() {
	*x = Observation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Observation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observation) ProtoMessage() {}

func (x *Observation) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Observation.ProtoReflect.Descriptor instead.
func (*Observation) Descriptor() ([]byte, []int) {
	return file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescGZIP(), []int{3}
}

func (x *Observation) GetId() *Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Observation) GetObservations() *pb.List {
	if x != nil {
		return x.Observations
	}
	return nil
}

type Observations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// batched observations for multiple workflow execution IDs
	Observations []*Observation `protobuf:"bytes,1,rep,name=observations,proto3" json:"observations,omitempty"`
}

func (x *Observations) Reset() {
	*x = Observations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Observations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observations) ProtoMessage() {}

func (x *Observations) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Observations.ProtoReflect.Descriptor instead.
func (*Observations) Descriptor() ([]byte, []int) {
	return file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescGZIP(), []int{4}
}

func (x *Observations) GetObservations() []*Observation {
	if x != nil {
		return x.Observations
	}
	return nil
}

type Report struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *Id                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Outcome *AggregationOutcome `protobuf:"bytes,2,opt,name=outcome,proto3" json:"outcome,omitempty"`
}

func (x *Report) Reset() {
	*x = Report{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Report) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Report) ProtoMessage() {}

func (x *Report) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Report.ProtoReflect.Descriptor instead.
func (*Report) Descriptor() ([]byte, []int) {
	return file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescGZIP(), []int{5}
}

func (x *Report) GetId() *Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Report) GetOutcome() *AggregationOutcome {
	if x != nil {
		return x.Outcome
	}
	return nil
}

type ReportInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           *Id  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ShouldReport bool `protobuf:"varint,2,opt,name=should_report,json=shouldReport,proto3" json:"should_report,omitempty"`
}

func (x *ReportInfo) Reset() {
	*x = ReportInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportInfo) ProtoMessage() {}

func (x *ReportInfo) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportInfo.ProtoReflect.Descriptor instead.
func (*ReportInfo) Descriptor() ([]byte, []int) {
	return file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescGZIP(), []int{6}
}

func (x *ReportInfo) GetId() *Id {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ReportInfo) GetShouldReport() bool {
	if x != nil {
		return x.ShouldReport
	}
	return false
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NOTE: list needs to be sorted by the order that IDs appeared in the query to ensure deterministic encoding
	CurrentReports []*Report `protobuf:"bytes,2,rep,name=current_reports,json=currentReports,proto3" json:"current_reports,omitempty"`
	// NOTE: list needs to be sorted by workflowId to ensure deterministic encoding
	Outcomes []*AggregationOutcome `protobuf:"bytes,3,rep,name=outcomes,proto3" json:"outcomes,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescGZIP(), []int{7}
}

func (x *Outcome) GetCurrentReports() []*Report {
	if x != nil {
		return x.CurrentReports
	}
	return nil
}

func (x *Outcome) GetOutcomes() []*AggregationOutcome {
	if x != nil {
		return x.Outcomes
	}
	return nil
}

var File_capabilities_consensus_ocr3_types_ocr3_types_proto protoreflect.FileDescriptor

var file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDesc = []byte{
	0x0a, 0x32, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x6f, 0x63, 0x72, 0x33, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x6f, 0x63, 0x72, 0x33, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x63, 0x72, 0x33, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x1a, 0x16, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x12, 0x41, 0x67, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12,
	0x37, 0x0a, 0x10, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x4f, 0x75, 0x74, 0x63,
	0x6f, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x2e, 0x4d, 0x61, 0x70, 0x52, 0x10, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x61, 0x62, 0x6c,
	0x65, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x75,
	0x6c, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x29, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x12, 0x20, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6f, 0x63, 0x72, 0x33, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x22, 0x56, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x0b, 0x4f,
	0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x63, 0x72, 0x33, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x0c, 0x6f, 0x62,
	0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0c,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4b, 0x0a, 0x0c,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3b, 0x0a, 0x0c,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x63, 0x72, 0x33, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x4f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x6f, 0x62, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x62, 0x0a, 0x06, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x6f, 0x63, 0x72, 0x33, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x38, 0x0a, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x63, 0x72, 0x33, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x52, 0x07, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x22, 0x51, 0x0a,
	0x0a, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6f, 0x63, 0x72, 0x33, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x68, 0x6f, 0x75, 0x6c, 0x64, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x88, 0x01, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6f, 0x63, 0x72, 0x33, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x0e, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x3a, 0x0a, 0x08, 0x6f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6f, 0x63,
	0x72, 0x33, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x52, 0x08, 0x6f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x42, 0x23, 0x5a, 0x21, 0x63,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x73,
	0x65, 0x6e, 0x73, 0x75, 0x73, 0x2f, 0x6f, 0x63, 0x72, 0x33, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescOnce sync.Once
	file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescData = file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDesc
)

func file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescGZIP() []byte {
	file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescOnce.Do(func() {
		file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescData)
	})
	return file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDescData
}

var file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_capabilities_consensus_ocr3_types_ocr3_types_proto_goTypes = []interface{}{
	(*AggregationOutcome)(nil), // 0: ocr3_types.AggregationOutcome
	(*Query)(nil),              // 1: ocr3_types.Query
	(*Id)(nil),                 // 2: ocr3_types.Id
	(*Observation)(nil),        // 3: ocr3_types.Observation
	(*Observations)(nil),       // 4: ocr3_types.Observations
	(*Report)(nil),             // 5: ocr3_types.Report
	(*ReportInfo)(nil),         // 6: ocr3_types.ReportInfo
	(*Outcome)(nil),            // 7: ocr3_types.Outcome
	(*pb.Map)(nil),             // 8: values.Map
	(*pb.List)(nil),            // 9: values.List
}
var file_capabilities_consensus_ocr3_types_ocr3_types_proto_depIdxs = []int32{
	8,  // 0: ocr3_types.AggregationOutcome.encodableOutcome:type_name -> values.Map
	2,  // 1: ocr3_types.Query.ids:type_name -> ocr3_types.Id
	2,  // 2: ocr3_types.Observation.id:type_name -> ocr3_types.Id
	9,  // 3: ocr3_types.Observation.observations:type_name -> values.List
	3,  // 4: ocr3_types.Observations.observations:type_name -> ocr3_types.Observation
	2,  // 5: ocr3_types.Report.id:type_name -> ocr3_types.Id
	0,  // 6: ocr3_types.Report.outcome:type_name -> ocr3_types.AggregationOutcome
	2,  // 7: ocr3_types.ReportInfo.id:type_name -> ocr3_types.Id
	5,  // 8: ocr3_types.Outcome.current_reports:type_name -> ocr3_types.Report
	0,  // 9: ocr3_types.Outcome.outcomes:type_name -> ocr3_types.AggregationOutcome
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_capabilities_consensus_ocr3_types_ocr3_types_proto_init() }
func file_capabilities_consensus_ocr3_types_ocr3_types_proto_init() {
	if File_capabilities_consensus_ocr3_types_ocr3_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregationOutcome); i {
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
		file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Observation); i {
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
		file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Observations); i {
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
		file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Report); i {
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
		file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportInfo); i {
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
		file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
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
			RawDescriptor: file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_capabilities_consensus_ocr3_types_ocr3_types_proto_goTypes,
		DependencyIndexes: file_capabilities_consensus_ocr3_types_ocr3_types_proto_depIdxs,
		MessageInfos:      file_capabilities_consensus_ocr3_types_ocr3_types_proto_msgTypes,
	}.Build()
	File_capabilities_consensus_ocr3_types_ocr3_types_proto = out.File
	file_capabilities_consensus_ocr3_types_ocr3_types_proto_rawDesc = nil
	file_capabilities_consensus_ocr3_types_ocr3_types_proto_goTypes = nil
	file_capabilities_consensus_ocr3_types_ocr3_types_proto_depIdxs = nil
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.26.1
// source: move_v1.proto

package v1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Messages with the following topic contain a message matching this
// specification:
// _telemetry/broker/trace/move/v1[/additional/topic/levels]
// Note that the topic allows for additional topic levels to be added in the
// future. Receiving clients must not assume there are no additional topic
// levels.
//
// This message describes telemetry data that a Solace PubSub+ broker captures
// as a result of moving a message from one queue to another; for example,
// moving a message to a dead message queue if TTL expires.
//
// Fields with names that end in "time_unix_nano" are 64-bit timestamps, in
// nanoseconds, since midnight, Jan. 1, 1970 UTC.
type SpanData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 16-byte globally unique trace ID. Any two spans with the same trace ID
	// are part of the same trace.
	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// 8-byte span ID, unique within the scope of a trace.
	SpanId []byte `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	// If not present, this is a root span. If present, this is an 8-byte span
	// ID of the parent span.
	ParentSpanId []byte `protobuf:"bytes,3,opt,name=parent_span_id,json=parentSpanId,proto3,oneof" json:"parent_span_id,omitempty"`
	// The start and end timestamps of the receive span. The start of the span
	// is when Guaranteed Messaging processing begins in the broker.
	StartTimeUnixNano int64 `protobuf:"fixed64,4,opt,name=start_time_unix_nano,json=startTimeUnixNano,proto3" json:"start_time_unix_nano,omitempty"`
	EndTimeUnixNano   int64 `protobuf:"fixed64,5,opt,name=end_time_unix_nano,json=endTimeUnixNano,proto3" json:"end_time_unix_nano,omitempty"`
	// The name of the queue or topic endpoint the message is being moved from.
	//
	// Types that are assignable to Source:
	//
	//	*SpanData_SourceQueueName
	//	*SpanData_SourceTopicEndpointName
	Source isSpanData_Source `protobuf_oneof:"source"`
	// The destination queue or topic endpoint the message is being moved to.
	//
	// Types that are assignable to Destination:
	//
	//	*SpanData_DestinationQueueName
	//	*SpanData_DestinationTopicEndpointName
	Destination isSpanData_Destination `protobuf_oneof:"destination"`
	// The nested "info" message below provides the following information:
	// * The reason for the message being moved.
	// * Any additional information associated with that particular reason.
	// Currently, the defined "info" messages contain no information. An enum
	// could have been used to express this instead. The "info" message technique
	// was chosen to allow for information related specifically to each scenario
	// to be easily added in the future.
	//
	// Types that are assignable to TypeInfo:
	//
	//	*SpanData_MaxRedeliveriesInfo
	//	*SpanData_TtlExpiredInfo
	//	*SpanData_RejectedOutcomeInfo
	TypeInfo isSpanData_TypeInfo `protobuf_oneof:"type_info"`
	// The router-name of the broker generating this message at the time the
	// message was generated.
	RouterName string `protobuf:"bytes,13,opt,name=router_name,json=routerName,proto3" json:"router_name,omitempty"`
	// The broker's message-vpn name. This field may be removed in the future
	// without a major version change since the field is specified as optional.
	//
	// Rather than rely on this field, receiving clients should obtain the VPN
	// by using an SMF API to extract the VPN_NAME_IN_USE from the API's Session
	// object. The message_vpn_name of all messages received from via an SMF
	// API's session will match the session's VPN_NAME_IN_USE.
	MessageVpnName *string `protobuf:"bytes,14,opt,name=message_vpn_name,json=messageVpnName,proto3,oneof" json:"message_vpn_name,omitempty"`
	// The SolOS version of the broker generating the message. All elements of
	// egress_spans will always have been created by the same broker version.
	SolosVersion              string `protobuf:"bytes,15,opt,name=solos_version,json=solosVersion,proto3" json:"solos_version,omitempty"`
	ReplicationGroupMessageId []byte `protobuf:"bytes,16,opt,name=replication_group_message_id,json=replicationGroupMessageId,proto3,oneof" json:"replication_group_message_id,omitempty"`
	// Partition number of queue the message is being moved from,
	// if the associated queue is a partitioned queue
	SourcePartitionNumber *uint32 `protobuf:"varint,17,opt,name=source_partition_number,json=sourcePartitionNumber,proto3,oneof" json:"source_partition_number,omitempty"`
	// Partition number of the destination DMQ the message is being moved to,
	// if the associated queue is a partitioned queue
	DestinationPartitionNumber *uint32 `protobuf:"varint,18,opt,name=destination_partition_number,json=destinationPartitionNumber,proto3,oneof" json:"destination_partition_number,omitempty"`
}

func (x *SpanData) Reset() {
	*x = SpanData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_v1_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpanData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpanData) ProtoMessage() {}

func (x *SpanData) ProtoReflect() protoreflect.Message {
	mi := &file_move_v1_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpanData.ProtoReflect.Descriptor instead.
func (*SpanData) Descriptor() ([]byte, []int) {
	return file_move_v1_proto_rawDescGZIP(), []int{0}
}

func (x *SpanData) GetTraceId() []byte {
	if x != nil {
		return x.TraceId
	}
	return nil
}

func (x *SpanData) GetSpanId() []byte {
	if x != nil {
		return x.SpanId
	}
	return nil
}

func (x *SpanData) GetParentSpanId() []byte {
	if x != nil {
		return x.ParentSpanId
	}
	return nil
}

func (x *SpanData) GetStartTimeUnixNano() int64 {
	if x != nil {
		return x.StartTimeUnixNano
	}
	return 0
}

func (x *SpanData) GetEndTimeUnixNano() int64 {
	if x != nil {
		return x.EndTimeUnixNano
	}
	return 0
}

func (m *SpanData) GetSource() isSpanData_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *SpanData) GetSourceQueueName() string {
	if x, ok := x.GetSource().(*SpanData_SourceQueueName); ok {
		return x.SourceQueueName
	}
	return ""
}

func (x *SpanData) GetSourceTopicEndpointName() string {
	if x, ok := x.GetSource().(*SpanData_SourceTopicEndpointName); ok {
		return x.SourceTopicEndpointName
	}
	return ""
}

func (m *SpanData) GetDestination() isSpanData_Destination {
	if m != nil {
		return m.Destination
	}
	return nil
}

func (x *SpanData) GetDestinationQueueName() string {
	if x, ok := x.GetDestination().(*SpanData_DestinationQueueName); ok {
		return x.DestinationQueueName
	}
	return ""
}

func (x *SpanData) GetDestinationTopicEndpointName() string {
	if x, ok := x.GetDestination().(*SpanData_DestinationTopicEndpointName); ok {
		return x.DestinationTopicEndpointName
	}
	return ""
}

func (m *SpanData) GetTypeInfo() isSpanData_TypeInfo {
	if m != nil {
		return m.TypeInfo
	}
	return nil
}

func (x *SpanData) GetMaxRedeliveriesInfo() *MaxRedeliveriesInfo {
	if x, ok := x.GetTypeInfo().(*SpanData_MaxRedeliveriesInfo); ok {
		return x.MaxRedeliveriesInfo
	}
	return nil
}

func (x *SpanData) GetTtlExpiredInfo() *TtlExpiredInfo {
	if x, ok := x.GetTypeInfo().(*SpanData_TtlExpiredInfo); ok {
		return x.TtlExpiredInfo
	}
	return nil
}

func (x *SpanData) GetRejectedOutcomeInfo() *RejectedOutcomeInfo {
	if x, ok := x.GetTypeInfo().(*SpanData_RejectedOutcomeInfo); ok {
		return x.RejectedOutcomeInfo
	}
	return nil
}

func (x *SpanData) GetRouterName() string {
	if x != nil {
		return x.RouterName
	}
	return ""
}

func (x *SpanData) GetMessageVpnName() string {
	if x != nil && x.MessageVpnName != nil {
		return *x.MessageVpnName
	}
	return ""
}

func (x *SpanData) GetSolosVersion() string {
	if x != nil {
		return x.SolosVersion
	}
	return ""
}

func (x *SpanData) GetReplicationGroupMessageId() []byte {
	if x != nil {
		return x.ReplicationGroupMessageId
	}
	return nil
}

func (x *SpanData) GetSourcePartitionNumber() uint32 {
	if x != nil && x.SourcePartitionNumber != nil {
		return *x.SourcePartitionNumber
	}
	return 0
}

func (x *SpanData) GetDestinationPartitionNumber() uint32 {
	if x != nil && x.DestinationPartitionNumber != nil {
		return *x.DestinationPartitionNumber
	}
	return 0
}

type isSpanData_Source interface {
	isSpanData_Source()
}

type SpanData_SourceQueueName struct {
	SourceQueueName string `protobuf:"bytes,6,opt,name=source_queue_name,json=sourceQueueName,proto3,oneof"`
}

type SpanData_SourceTopicEndpointName struct {
	SourceTopicEndpointName string `protobuf:"bytes,7,opt,name=source_topic_endpoint_name,json=sourceTopicEndpointName,proto3,oneof"`
}

func (*SpanData_SourceQueueName) isSpanData_Source() {}

func (*SpanData_SourceTopicEndpointName) isSpanData_Source() {}

type isSpanData_Destination interface {
	isSpanData_Destination()
}

type SpanData_DestinationQueueName struct {
	DestinationQueueName string `protobuf:"bytes,8,opt,name=destination_queue_name,json=destinationQueueName,proto3,oneof"`
}

type SpanData_DestinationTopicEndpointName struct {
	DestinationTopicEndpointName string `protobuf:"bytes,9,opt,name=destination_topic_endpoint_name,json=destinationTopicEndpointName,proto3,oneof"`
}

func (*SpanData_DestinationQueueName) isSpanData_Destination() {}

func (*SpanData_DestinationTopicEndpointName) isSpanData_Destination() {}

type isSpanData_TypeInfo interface {
	isSpanData_TypeInfo()
}

type SpanData_MaxRedeliveriesInfo struct {
	MaxRedeliveriesInfo *MaxRedeliveriesInfo `protobuf:"bytes,10,opt,name=max_redeliveries_info,json=maxRedeliveriesInfo,proto3,oneof"`
}

type SpanData_TtlExpiredInfo struct {
	TtlExpiredInfo *TtlExpiredInfo `protobuf:"bytes,11,opt,name=ttl_expired_info,json=ttlExpiredInfo,proto3,oneof"`
}

type SpanData_RejectedOutcomeInfo struct {
	RejectedOutcomeInfo *RejectedOutcomeInfo `protobuf:"bytes,12,opt,name=rejected_outcome_info,json=rejectedOutcomeInfo,proto3,oneof"`
}

func (*SpanData_MaxRedeliveriesInfo) isSpanData_TypeInfo() {}

func (*SpanData_TtlExpiredInfo) isSpanData_TypeInfo() {}

func (*SpanData_RejectedOutcomeInfo) isSpanData_TypeInfo() {}

// The presence of this message implies the reason for the span is that a
// message exceeded the maximum number of redeliveries.
type MaxRedeliveriesInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MaxRedeliveriesInfo) Reset() {
	*x = MaxRedeliveriesInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_v1_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaxRedeliveriesInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaxRedeliveriesInfo) ProtoMessage() {}

func (x *MaxRedeliveriesInfo) ProtoReflect() protoreflect.Message {
	mi := &file_move_v1_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaxRedeliveriesInfo.ProtoReflect.Descriptor instead.
func (*MaxRedeliveriesInfo) Descriptor() ([]byte, []int) {
	return file_move_v1_proto_rawDescGZIP(), []int{1}
}

// The presence of this message implies the reason for the span is that the
// message's TTL has expired.
type TtlExpiredInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TtlExpiredInfo) Reset() {
	*x = TtlExpiredInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_v1_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TtlExpiredInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TtlExpiredInfo) ProtoMessage() {}

func (x *TtlExpiredInfo) ProtoReflect() protoreflect.Message {
	mi := &file_move_v1_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TtlExpiredInfo.ProtoReflect.Descriptor instead.
func (*TtlExpiredInfo) Descriptor() ([]byte, []int) {
	return file_move_v1_proto_rawDescGZIP(), []int{2}
}

// The presence of this message implies the reason for the span is that a
// consuming client settled the message with an outcome of "rejected".
type RejectedOutcomeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RejectedOutcomeInfo) Reset() {
	*x = RejectedOutcomeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_move_v1_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RejectedOutcomeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RejectedOutcomeInfo) ProtoMessage() {}

func (x *RejectedOutcomeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_move_v1_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RejectedOutcomeInfo.ProtoReflect.Descriptor instead.
func (*RejectedOutcomeInfo) Descriptor() ([]byte, []int) {
	return file_move_v1_proto_rawDescGZIP(), []int{3}
}

var File_move_v1_proto protoreflect.FileDescriptor

var file_move_v1_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x76, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x2b, 0x73, 0x6f, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x74,
	0x72, 0x61, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xf9, 0x09, 0x0a,
	0x08, 0x53, 0x70, 0x61, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x72, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a,
	0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x03, 0x52, 0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x53,
	0x70, 0x61, 0x6e, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x2f, 0x0a, 0x14, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x6e, 0x61, 0x6e, 0x6f,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x10, 0x52, 0x11, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d,
	0x65, 0x55, 0x6e, 0x69, 0x78, 0x4e, 0x61, 0x6e, 0x6f, 0x12, 0x2b, 0x0a, 0x12, 0x65, 0x6e, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x78, 0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x10, 0x52, 0x0f, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e,
	0x69, 0x78, 0x4e, 0x61, 0x6e, 0x6f, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x1a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x17, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x14, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x47, 0x0a, 0x1f, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x70, 0x69, 0x63,
	0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x1c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x76, 0x0a, 0x15, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x76, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x4d, 0x61, 0x78, 0x52, 0x65, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x02, 0x52, 0x13, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x67, 0x0a, 0x10,
	0x74, 0x74, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x63, 0x65, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x76,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x74, 0x6c, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x49,
	0x6e, 0x66, 0x6f, 0x48, 0x02, 0x52, 0x0e, 0x74, 0x74, 0x6c, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x76, 0x0a, 0x15, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x5f, 0x6f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x0c,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x40, 0x2e, 0x73, 0x6f, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x62, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x65, 0x2e, 0x6d, 0x6f, 0x76, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x02, 0x52, 0x13, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2d,
	0x0a, 0x10, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x76, 0x70, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x56, 0x70, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x23, 0x0a,
	0x0d, 0x73, 0x6f, 0x6c, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x6c, 0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x1c, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x05, 0x52, 0x19, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x3b, 0x0a, 0x17, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x06, 0x52, 0x15, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x45, 0x0a, 0x1c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x07, 0x52, 0x1a, 0x64,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0b, 0x0a, 0x09, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x70,
	0x61, 0x6e, 0x5f, 0x69, 0x64, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x76, 0x70, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x1f, 0x0a, 0x1d, 0x5f, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x42, 0x1a, 0x0a, 0x18, 0x5f,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x1f, 0x0a, 0x1d, 0x5f, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x4d, 0x61, 0x78, 0x52,
	0x65, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x69, 0x65, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x22,
	0x10, 0x0a, 0x0e, 0x54, 0x74, 0x6c, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x15, 0x0a, 0x13, 0x52, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4f, 0x75, 0x74,
	0x63, 0x6f, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_move_v1_proto_rawDescOnce sync.Once
	file_move_v1_proto_rawDescData = file_move_v1_proto_rawDesc
)

func file_move_v1_proto_rawDescGZIP() []byte {
	file_move_v1_proto_rawDescOnce.Do(func() {
		file_move_v1_proto_rawDescData = protoimpl.X.CompressGZIP(file_move_v1_proto_rawDescData)
	})
	return file_move_v1_proto_rawDescData
}

var file_move_v1_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_move_v1_proto_goTypes = []interface{}{
	(*SpanData)(nil),            // 0: solace.messaging.proto.broker.trace.move.v1.SpanData
	(*MaxRedeliveriesInfo)(nil), // 1: solace.messaging.proto.broker.trace.move.v1.MaxRedeliveriesInfo
	(*TtlExpiredInfo)(nil),      // 2: solace.messaging.proto.broker.trace.move.v1.TtlExpiredInfo
	(*RejectedOutcomeInfo)(nil), // 3: solace.messaging.proto.broker.trace.move.v1.RejectedOutcomeInfo
}
var file_move_v1_proto_depIdxs = []int32{
	1, // 0: solace.messaging.proto.broker.trace.move.v1.SpanData.max_redeliveries_info:type_name -> solace.messaging.proto.broker.trace.move.v1.MaxRedeliveriesInfo
	2, // 1: solace.messaging.proto.broker.trace.move.v1.SpanData.ttl_expired_info:type_name -> solace.messaging.proto.broker.trace.move.v1.TtlExpiredInfo
	3, // 2: solace.messaging.proto.broker.trace.move.v1.SpanData.rejected_outcome_info:type_name -> solace.messaging.proto.broker.trace.move.v1.RejectedOutcomeInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_move_v1_proto_init() }
func file_move_v1_proto_init() {
	if File_move_v1_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_move_v1_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpanData); i {
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
		file_move_v1_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaxRedeliveriesInfo); i {
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
		file_move_v1_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TtlExpiredInfo); i {
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
		file_move_v1_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RejectedOutcomeInfo); i {
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
	file_move_v1_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SpanData_SourceQueueName)(nil),
		(*SpanData_SourceTopicEndpointName)(nil),
		(*SpanData_DestinationQueueName)(nil),
		(*SpanData_DestinationTopicEndpointName)(nil),
		(*SpanData_MaxRedeliveriesInfo)(nil),
		(*SpanData_TtlExpiredInfo)(nil),
		(*SpanData_RejectedOutcomeInfo)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_move_v1_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_move_v1_proto_goTypes,
		DependencyIndexes: file_move_v1_proto_depIdxs,
		MessageInfos:      file_move_v1_proto_msgTypes,
	}.Build()
	File_move_v1_proto = out.File
	file_move_v1_proto_rawDesc = nil
	file_move_v1_proto_goTypes = nil
	file_move_v1_proto_depIdxs = nil
}

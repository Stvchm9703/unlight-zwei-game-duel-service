// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.10.0
// source: sr_message.proto

package proto

import (
	dtpb "ULZGameDuelService/proto"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// ANCHOR: message-struct for game-service
// SECTION: message.proto
// -------------------------------------------------------------
type SESkillCalReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncomeCard []*dtpb.EventCard  `protobuf:"bytes,1,rep,name=income_card,json=incomeCard,proto3" json:"income_card,omitempty"`
	Feat       []*dtpb.SkillSet   `protobuf:"bytes,2,rep,name=feat,proto3" json:"feat,omitempty"`
	FromCli    string             `protobuf:"bytes,3,opt,name=from_cli,json=fromCli,proto3" json:"from_cli,omitempty"`
	TargType   dtpb.EventCardType `protobuf:"varint,4,opt,name=targ_type,json=targType,proto3,enum=ULZProto.EventCardType" json:"targ_type,omitempty"`
}

func (x *SESkillCalReq) Reset() {
	*x = SESkillCalReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sr_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SESkillCalReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SESkillCalReq) ProtoMessage() {}

func (x *SESkillCalReq) ProtoReflect() protoreflect.Message {
	mi := &file_sr_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SESkillCalReq.ProtoReflect.Descriptor instead.
func (*SESkillCalReq) Descriptor() ([]byte, []int) {
	return file_sr_message_proto_rawDescGZIP(), []int{0}
}

func (x *SESkillCalReq) GetIncomeCard() []*dtpb.EventCard {
	if x != nil {
		return x.IncomeCard
	}
	return nil
}

func (x *SESkillCalReq) GetFeat() []*dtpb.SkillSet {
	if x != nil {
		return x.Feat
	}
	return nil
}

func (x *SESkillCalReq) GetFromCli() string {
	if x != nil {
		return x.FromCli
	}
	return ""
}

func (x *SESkillCalReq) GetTargType() dtpb.EventCardType {
	if x != nil {
		return x.TargType
	}
	return dtpb.EventCardType_NULL
}

type SESkillCalResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultVal    int32                `protobuf:"varint,1,opt,name=result_val,json=resultVal,proto3" json:"result_val,omitempty"`
	EffectResult []*dtpb.EffectResult `protobuf:"bytes,2,rep,name=effect_result,json=effectResult,proto3" json:"effect_result,omitempty"`
	TargType     dtpb.EventCardType   `protobuf:"varint,3,opt,name=targ_type,json=targType,proto3,enum=ULZProto.EventCardType" json:"targ_type,omitempty"`
}

func (x *SESkillCalResp) Reset() {
	*x = SESkillCalResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sr_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SESkillCalResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SESkillCalResp) ProtoMessage() {}

func (x *SESkillCalResp) ProtoReflect() protoreflect.Message {
	mi := &file_sr_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SESkillCalResp.ProtoReflect.Descriptor instead.
func (*SESkillCalResp) Descriptor() ([]byte, []int) {
	return file_sr_message_proto_rawDescGZIP(), []int{1}
}

func (x *SESkillCalResp) GetResultVal() int32 {
	if x != nil {
		return x.ResultVal
	}
	return 0
}

func (x *SESkillCalResp) GetEffectResult() []*dtpb.EffectResult {
	if x != nil {
		return x.EffectResult
	}
	return nil
}

func (x *SESkillCalResp) GetTargType() dtpb.EventCardType {
	if x != nil {
		return x.TargType
	}
	return dtpb.EventCardType_NULL
}

type SEDiceCalReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IncomeDice   int32                `protobuf:"varint,1,opt,name=income_dice,json=incomeDice,proto3" json:"income_dice,omitempty"`
	Act          int32                `protobuf:"varint,2,opt,name=act,proto3" json:"act,omitempty"`
	EffectResult []*dtpb.EffectResult `protobuf:"bytes,3,rep,name=effect_result,json=effectResult,proto3" json:"effect_result,omitempty"`
}

func (x *SEDiceCalReq) Reset() {
	*x = SEDiceCalReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sr_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SEDiceCalReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SEDiceCalReq) ProtoMessage() {}

func (x *SEDiceCalReq) ProtoReflect() protoreflect.Message {
	mi := &file_sr_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SEDiceCalReq.ProtoReflect.Descriptor instead.
func (*SEDiceCalReq) Descriptor() ([]byte, []int) {
	return file_sr_message_proto_rawDescGZIP(), []int{2}
}

func (x *SEDiceCalReq) GetIncomeDice() int32 {
	if x != nil {
		return x.IncomeDice
	}
	return 0
}

func (x *SEDiceCalReq) GetAct() int32 {
	if x != nil {
		return x.Act
	}
	return 0
}

func (x *SEDiceCalReq) GetEffectResult() []*dtpb.EffectResult {
	if x != nil {
		return x.EffectResult
	}
	return nil
}

type SEDiceCalResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiceResult []*DiceResultSet `protobuf:"bytes,1,rep,name=dice_result,json=diceResult,proto3" json:"dice_result,omitempty"`
}

func (x *SEDiceCalResp) Reset() {
	*x = SEDiceCalResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sr_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SEDiceCalResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SEDiceCalResp) ProtoMessage() {}

func (x *SEDiceCalResp) ProtoReflect() protoreflect.Message {
	mi := &file_sr_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SEDiceCalResp.ProtoReflect.Descriptor instead.
func (*SEDiceCalResp) Descriptor() ([]byte, []int) {
	return file_sr_message_proto_rawDescGZIP(), []int{3}
}

func (x *SEDiceCalResp) GetDiceResult() []*DiceResultSet {
	if x != nil {
		return x.DiceResult
	}
	return nil
}

type DiceResultSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []int32 `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
}

func (x *DiceResultSet) Reset() {
	*x = DiceResultSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sr_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiceResultSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiceResultSet) ProtoMessage() {}

func (x *DiceResultSet) ProtoReflect() protoreflect.Message {
	mi := &file_sr_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiceResultSet.ProtoReflect.Descriptor instead.
func (*DiceResultSet) Descriptor() ([]byte, []int) {
	return file_sr_message_proto_rawDescGZIP(), []int{4}
}

func (x *DiceResultSet) GetValue() []int32 {
	if x != nil {
		return x.Value
	}
	return nil
}

type SEEffectCalReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FromTime       *dtpb.EffectTiming `protobuf:"bytes,2,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	ToTime         *dtpb.EffectTiming `protobuf:"bytes,3,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
	GamesetInstant *dtpb.GameDataSet  `protobuf:"bytes,4,opt,name=gameset_instant,json=gamesetInstant,proto3" json:"gameset_instant,omitempty"`
	FromCli        string             `protobuf:"bytes,5,opt,name=from_cli,json=fromCli,proto3" json:"from_cli,omitempty"`
	Remark         string             `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
}

func (x *SEEffectCalReq) Reset() {
	*x = SEEffectCalReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sr_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SEEffectCalReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SEEffectCalReq) ProtoMessage() {}

func (x *SEEffectCalReq) ProtoReflect() protoreflect.Message {
	mi := &file_sr_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SEEffectCalReq.ProtoReflect.Descriptor instead.
func (*SEEffectCalReq) Descriptor() ([]byte, []int) {
	return file_sr_message_proto_rawDescGZIP(), []int{5}
}

func (x *SEEffectCalReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SEEffectCalReq) GetFromTime() *dtpb.EffectTiming {
	if x != nil {
		return x.FromTime
	}
	return nil
}

func (x *SEEffectCalReq) GetToTime() *dtpb.EffectTiming {
	if x != nil {
		return x.ToTime
	}
	return nil
}

func (x *SEEffectCalReq) GetGamesetInstant() *dtpb.GameDataSet {
	if x != nil {
		return x.GamesetInstant
	}
	return nil
}

func (x *SEEffectCalReq) GetFromCli() string {
	if x != nil {
		return x.FromCli
	}
	return ""
}

func (x *SEEffectCalReq) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type SEEffectCalResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FromCli       string             `protobuf:"bytes,2,opt,name=from_cli,json=fromCli,proto3" json:"from_cli,omitempty"`
	GamesetResult *dtpb.GameDataSet  `protobuf:"bytes,3,opt,name=gameset_result,json=gamesetResult,proto3" json:"gameset_result,omitempty"`
	ResultInfo    []string           `protobuf:"bytes,4,rep,name=result_info,json=resultInfo,proto3" json:"result_info,omitempty"`
	FromTime      *dtpb.EffectTiming `protobuf:"bytes,5,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	ToTime        *dtpb.EffectTiming `protobuf:"bytes,6,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
}

func (x *SEEffectCalResp) Reset() {
	*x = SEEffectCalResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sr_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SEEffectCalResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SEEffectCalResp) ProtoMessage() {}

func (x *SEEffectCalResp) ProtoReflect() protoreflect.Message {
	mi := &file_sr_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SEEffectCalResp.ProtoReflect.Descriptor instead.
func (*SEEffectCalResp) Descriptor() ([]byte, []int) {
	return file_sr_message_proto_rawDescGZIP(), []int{6}
}

func (x *SEEffectCalResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SEEffectCalResp) GetFromCli() string {
	if x != nil {
		return x.FromCli
	}
	return ""
}

func (x *SEEffectCalResp) GetGamesetResult() *dtpb.GameDataSet {
	if x != nil {
		return x.GamesetResult
	}
	return nil
}

func (x *SEEffectCalResp) GetResultInfo() []string {
	if x != nil {
		return x.ResultInfo
	}
	return nil
}

func (x *SEEffectCalResp) GetFromTime() *dtpb.EffectTiming {
	if x != nil {
		return x.FromTime
	}
	return nil
}

func (x *SEEffectCalResp) GetToTime() *dtpb.EffectTiming {
	if x != nil {
		return x.ToTime
	}
	return nil
}

var File_sr_message_proto protoreflect.FileDescriptor

var file_sr_message_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x01, 0x0a, 0x0d, 0x53, 0x45, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x34, 0x0a, 0x0b, 0x69, 0x6e,
	0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x43, 0x61, 0x72, 0x64, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x26, 0x0a, 0x04, 0x66, 0x65, 0x61, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x53,
	0x65, 0x74, 0x52, 0x04, 0x66, 0x65, 0x61, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d,
	0x5f, 0x63, 0x6c, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d,
	0x43, 0x6c, 0x69, 0x12, 0x34, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x08, 0x74, 0x61, 0x72, 0x67, 0x54, 0x79, 0x70, 0x65, 0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x53, 0x45,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x43, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x0d, 0x65,
	0x66, 0x66, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x16, 0x2e, 0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x65, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x34, 0x0a, 0x09, 0x74, 0x61, 0x72, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x55, 0x4c,
	0x5a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x72, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x74, 0x61, 0x72, 0x67, 0x54, 0x79, 0x70, 0x65, 0x22, 0x7e,
	0x0a, 0x0c, 0x53, 0x45, 0x44, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x5f, 0x64, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x63, 0x6f, 0x6d, 0x65, 0x44, 0x69, 0x63, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x63,
	0x74, 0x12, 0x3b, 0x0a, 0x0d, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x55, 0x4c, 0x5a, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x0c, 0x65, 0x66, 0x66, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x49,
	0x0a, 0x0d, 0x53, 0x45, 0x44, 0x69, 0x63, 0x65, 0x43, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x38, 0x0a, 0x0b, 0x64, 0x69, 0x63, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x44, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x52, 0x0a, 0x64,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x25, 0x0a, 0x0d, 0x44, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x53, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0xf9, 0x01, 0x0a, 0x0e, 0x53, 0x45, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x55, 0x4c, 0x5a, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x54, 0x69, 0x6d, 0x69, 0x6e,
	0x67, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0f, 0x67, 0x61, 0x6d,
	0x65, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x52, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x73,
	0x65, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x6f,
	0x6d, 0x5f, 0x63, 0x6c, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x72, 0x6f,
	0x6d, 0x43, 0x6c, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x22, 0x81, 0x02, 0x0a,
	0x0f, 0x53, 0x45, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x43, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x63, 0x6c, 0x69, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x43, 0x6c, 0x69, 0x12, 0x3c, 0x0a, 0x0e, 0x67,
	0x61, 0x6d, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x74, 0x52, 0x0d, 0x67, 0x61, 0x6d, 0x65,
	0x73, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x09, 0x66, 0x72,
	0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x66, 0x66, 0x65, 0x63, 0x74, 0x54,
	0x69, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x2f, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x66, 0x66, 0x65,
	0x63, 0x74, 0x54, 0x69, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65,
	0x42, 0x1b, 0x5a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x11, 0x55, 0x4c, 0x5a, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x6f, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sr_message_proto_rawDescOnce sync.Once
	file_sr_message_proto_rawDescData = file_sr_message_proto_rawDesc
)

func file_sr_message_proto_rawDescGZIP() []byte {
	file_sr_message_proto_rawDescOnce.Do(func() {
		file_sr_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_sr_message_proto_rawDescData)
	})
	return file_sr_message_proto_rawDescData
}

var file_sr_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sr_message_proto_goTypes = []interface{}{
	(*SESkillCalReq)(nil),     // 0: ULZProto.SESkillCalReq
	(*SESkillCalResp)(nil),    // 1: ULZProto.SESkillCalResp
	(*SEDiceCalReq)(nil),      // 2: ULZProto.SEDiceCalReq
	(*SEDiceCalResp)(nil),     // 3: ULZProto.SEDiceCalResp
	(*DiceResultSet)(nil),     // 4: ULZProto.DiceResultSet
	(*SEEffectCalReq)(nil),    // 5: ULZProto.SEEffectCalReq
	(*SEEffectCalResp)(nil),   // 6: ULZProto.SEEffectCalResp
	(*dtpb.EventCard)(nil),    // 7: ULZProto.EventCard
	(*dtpb.SkillSet)(nil),     // 8: ULZProto.SkillSet
	(dtpb.EventCardType)(0),   // 9: ULZProto.EventCardType
	(*dtpb.EffectResult)(nil), // 10: ULZProto.EffectResult
	(*dtpb.EffectTiming)(nil), // 11: ULZProto.EffectTiming
	(*dtpb.GameDataSet)(nil),  // 12: ULZProto.GameDataSet
}
var file_sr_message_proto_depIdxs = []int32{
	7,  // 0: ULZProto.SESkillCalReq.income_card:type_name -> ULZProto.EventCard
	8,  // 1: ULZProto.SESkillCalReq.feat:type_name -> ULZProto.SkillSet
	9,  // 2: ULZProto.SESkillCalReq.targ_type:type_name -> ULZProto.EventCardType
	10, // 3: ULZProto.SESkillCalResp.effect_result:type_name -> ULZProto.EffectResult
	9,  // 4: ULZProto.SESkillCalResp.targ_type:type_name -> ULZProto.EventCardType
	10, // 5: ULZProto.SEDiceCalReq.effect_result:type_name -> ULZProto.EffectResult
	4,  // 6: ULZProto.SEDiceCalResp.dice_result:type_name -> ULZProto.DiceResultSet
	11, // 7: ULZProto.SEEffectCalReq.from_time:type_name -> ULZProto.EffectTiming
	11, // 8: ULZProto.SEEffectCalReq.to_time:type_name -> ULZProto.EffectTiming
	12, // 9: ULZProto.SEEffectCalReq.gameset_instant:type_name -> ULZProto.GameDataSet
	12, // 10: ULZProto.SEEffectCalResp.gameset_result:type_name -> ULZProto.GameDataSet
	11, // 11: ULZProto.SEEffectCalResp.from_time:type_name -> ULZProto.EffectTiming
	11, // 12: ULZProto.SEEffectCalResp.to_time:type_name -> ULZProto.EffectTiming
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_sr_message_proto_init() }
func file_sr_message_proto_init() {
	if File_sr_message_proto != nil {
		return
	}
	dtpb.File_Data_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_sr_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SESkillCalReq); i {
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
		file_sr_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SESkillCalResp); i {
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
		file_sr_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SEDiceCalReq); i {
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
		file_sr_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SEDiceCalResp); i {
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
		file_sr_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DiceResultSet); i {
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
		file_sr_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SEEffectCalReq); i {
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
		file_sr_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SEEffectCalResp); i {
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
			RawDescriptor: file_sr_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sr_message_proto_goTypes,
		DependencyIndexes: file_sr_message_proto_depIdxs,
		MessageInfos:      file_sr_message_proto_msgTypes,
	}.Build()
	File_sr_message_proto = out.File
	file_sr_message_proto_rawDesc = nil
	file_sr_message_proto_goTypes = nil
	file_sr_message_proto_depIdxs = nil
}
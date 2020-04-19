// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message.proto

package proto

import (
	fmt "fmt"
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

// ANCHOR: message-struct for game-service
// SECTION: message.proto
// -------------------------------------------------------------
type SESkillCalReq struct {
	IncomeCard           []*EventCard `protobuf:"bytes,1,rep,name=income_card,json=incomeCard,proto3" json:"income_card,omitempty"`
	Feat                 []*SkillSet  `protobuf:"bytes,2,rep,name=feat,proto3" json:"feat,omitempty"`
	FromCli              string       `protobuf:"bytes,3,opt,name=from_cli,json=fromCli,proto3" json:"from_cli,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SESkillCalReq) Reset()         { *m = SESkillCalReq{} }
func (m *SESkillCalReq) String() string { return proto.CompactTextString(m) }
func (*SESkillCalReq) ProtoMessage()    {}
func (*SESkillCalReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{0}
}

func (m *SESkillCalReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SESkillCalReq.Unmarshal(m, b)
}
func (m *SESkillCalReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SESkillCalReq.Marshal(b, m, deterministic)
}
func (m *SESkillCalReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SESkillCalReq.Merge(m, src)
}
func (m *SESkillCalReq) XXX_Size() int {
	return xxx_messageInfo_SESkillCalReq.Size(m)
}
func (m *SESkillCalReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SESkillCalReq.DiscardUnknown(m)
}

var xxx_messageInfo_SESkillCalReq proto.InternalMessageInfo

func (m *SESkillCalReq) GetIncomeCard() []*EventCard {
	if m != nil {
		return m.IncomeCard
	}
	return nil
}

func (m *SESkillCalReq) GetFeat() []*SkillSet {
	if m != nil {
		return m.Feat
	}
	return nil
}

func (m *SESkillCalReq) GetFromCli() string {
	if m != nil {
		return m.FromCli
	}
	return ""
}

type SESkillCalResp struct {
	ResultVal            int32           `protobuf:"varint,1,opt,name=result_val,json=resultVal,proto3" json:"result_val,omitempty"`
	EffectResult         []*EffectResult `protobuf:"bytes,2,rep,name=effect_result,json=effectResult,proto3" json:"effect_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *SESkillCalResp) Reset()         { *m = SESkillCalResp{} }
func (m *SESkillCalResp) String() string { return proto.CompactTextString(m) }
func (*SESkillCalResp) ProtoMessage()    {}
func (*SESkillCalResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{1}
}

func (m *SESkillCalResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SESkillCalResp.Unmarshal(m, b)
}
func (m *SESkillCalResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SESkillCalResp.Marshal(b, m, deterministic)
}
func (m *SESkillCalResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SESkillCalResp.Merge(m, src)
}
func (m *SESkillCalResp) XXX_Size() int {
	return xxx_messageInfo_SESkillCalResp.Size(m)
}
func (m *SESkillCalResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SESkillCalResp.DiscardUnknown(m)
}

var xxx_messageInfo_SESkillCalResp proto.InternalMessageInfo

func (m *SESkillCalResp) GetResultVal() int32 {
	if m != nil {
		return m.ResultVal
	}
	return 0
}

func (m *SESkillCalResp) GetEffectResult() []*EffectResult {
	if m != nil {
		return m.EffectResult
	}
	return nil
}

type SEDiceCalReq struct {
	IncomeDice           int32    `protobuf:"varint,1,opt,name=income_dice,json=incomeDice,proto3" json:"income_dice,omitempty"`
	Act                  int32    `protobuf:"varint,2,opt,name=act,proto3" json:"act,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SEDiceCalReq) Reset()         { *m = SEDiceCalReq{} }
func (m *SEDiceCalReq) String() string { return proto.CompactTextString(m) }
func (*SEDiceCalReq) ProtoMessage()    {}
func (*SEDiceCalReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{2}
}

func (m *SEDiceCalReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SEDiceCalReq.Unmarshal(m, b)
}
func (m *SEDiceCalReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SEDiceCalReq.Marshal(b, m, deterministic)
}
func (m *SEDiceCalReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SEDiceCalReq.Merge(m, src)
}
func (m *SEDiceCalReq) XXX_Size() int {
	return xxx_messageInfo_SEDiceCalReq.Size(m)
}
func (m *SEDiceCalReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SEDiceCalReq.DiscardUnknown(m)
}

var xxx_messageInfo_SEDiceCalReq proto.InternalMessageInfo

func (m *SEDiceCalReq) GetIncomeDice() int32 {
	if m != nil {
		return m.IncomeDice
	}
	return 0
}

func (m *SEDiceCalReq) GetAct() int32 {
	if m != nil {
		return m.Act
	}
	return 0
}

type SEDiceCalResp struct {
	DiceResult           []*DiceResultSet `protobuf:"bytes,1,rep,name=dice_result,json=diceResult,proto3" json:"dice_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SEDiceCalResp) Reset()         { *m = SEDiceCalResp{} }
func (m *SEDiceCalResp) String() string { return proto.CompactTextString(m) }
func (*SEDiceCalResp) ProtoMessage()    {}
func (*SEDiceCalResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{3}
}

func (m *SEDiceCalResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SEDiceCalResp.Unmarshal(m, b)
}
func (m *SEDiceCalResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SEDiceCalResp.Marshal(b, m, deterministic)
}
func (m *SEDiceCalResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SEDiceCalResp.Merge(m, src)
}
func (m *SEDiceCalResp) XXX_Size() int {
	return xxx_messageInfo_SEDiceCalResp.Size(m)
}
func (m *SEDiceCalResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SEDiceCalResp.DiscardUnknown(m)
}

var xxx_messageInfo_SEDiceCalResp proto.InternalMessageInfo

func (m *SEDiceCalResp) GetDiceResult() []*DiceResultSet {
	if m != nil {
		return m.DiceResult
	}
	return nil
}

type DiceResultSet struct {
	Value                []int32  `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiceResultSet) Reset()         { *m = DiceResultSet{} }
func (m *DiceResultSet) String() string { return proto.CompactTextString(m) }
func (*DiceResultSet) ProtoMessage()    {}
func (*DiceResultSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{4}
}

func (m *DiceResultSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiceResultSet.Unmarshal(m, b)
}
func (m *DiceResultSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiceResultSet.Marshal(b, m, deterministic)
}
func (m *DiceResultSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiceResultSet.Merge(m, src)
}
func (m *DiceResultSet) XXX_Size() int {
	return xxx_messageInfo_DiceResultSet.Size(m)
}
func (m *DiceResultSet) XXX_DiscardUnknown() {
	xxx_messageInfo_DiceResultSet.DiscardUnknown(m)
}

var xxx_messageInfo_DiceResultSet proto.InternalMessageInfo

func (m *DiceResultSet) GetValue() []int32 {
	if m != nil {
		return m.Value
	}
	return nil
}

type SEEffectCalReq struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FromTime             *EffectTiming `protobuf:"bytes,2,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	ToTime               *EffectTiming `protobuf:"bytes,3,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
	GamesetInstant       *GameDataSet  `protobuf:"bytes,4,opt,name=gameset_instant,json=gamesetInstant,proto3" json:"gameset_instant,omitempty"`
	FromCli              string        `protobuf:"bytes,5,opt,name=from_cli,json=fromCli,proto3" json:"from_cli,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SEEffectCalReq) Reset()         { *m = SEEffectCalReq{} }
func (m *SEEffectCalReq) String() string { return proto.CompactTextString(m) }
func (*SEEffectCalReq) ProtoMessage()    {}
func (*SEEffectCalReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{5}
}

func (m *SEEffectCalReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SEEffectCalReq.Unmarshal(m, b)
}
func (m *SEEffectCalReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SEEffectCalReq.Marshal(b, m, deterministic)
}
func (m *SEEffectCalReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SEEffectCalReq.Merge(m, src)
}
func (m *SEEffectCalReq) XXX_Size() int {
	return xxx_messageInfo_SEEffectCalReq.Size(m)
}
func (m *SEEffectCalReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SEEffectCalReq.DiscardUnknown(m)
}

var xxx_messageInfo_SEEffectCalReq proto.InternalMessageInfo

func (m *SEEffectCalReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SEEffectCalReq) GetFromTime() *EffectTiming {
	if m != nil {
		return m.FromTime
	}
	return nil
}

func (m *SEEffectCalReq) GetToTime() *EffectTiming {
	if m != nil {
		return m.ToTime
	}
	return nil
}

func (m *SEEffectCalReq) GetGamesetInstant() *GameDataSet {
	if m != nil {
		return m.GamesetInstant
	}
	return nil
}

func (m *SEEffectCalReq) GetFromCli() string {
	if m != nil {
		return m.FromCli
	}
	return ""
}

type SEEffectCalResp struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FromCli              string       `protobuf:"bytes,2,opt,name=from_cli,json=fromCli,proto3" json:"from_cli,omitempty"`
	GamesetResult        *GameDataSet `protobuf:"bytes,3,opt,name=gameset_result,json=gamesetResult,proto3" json:"gameset_result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *SEEffectCalResp) Reset()         { *m = SEEffectCalResp{} }
func (m *SEEffectCalResp) String() string { return proto.CompactTextString(m) }
func (*SEEffectCalResp) ProtoMessage()    {}
func (*SEEffectCalResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_33c57e4bae7b9afd, []int{6}
}

func (m *SEEffectCalResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SEEffectCalResp.Unmarshal(m, b)
}
func (m *SEEffectCalResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SEEffectCalResp.Marshal(b, m, deterministic)
}
func (m *SEEffectCalResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SEEffectCalResp.Merge(m, src)
}
func (m *SEEffectCalResp) XXX_Size() int {
	return xxx_messageInfo_SEEffectCalResp.Size(m)
}
func (m *SEEffectCalResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SEEffectCalResp.DiscardUnknown(m)
}

var xxx_messageInfo_SEEffectCalResp proto.InternalMessageInfo

func (m *SEEffectCalResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SEEffectCalResp) GetFromCli() string {
	if m != nil {
		return m.FromCli
	}
	return ""
}

func (m *SEEffectCalResp) GetGamesetResult() *GameDataSet {
	if m != nil {
		return m.GamesetResult
	}
	return nil
}

func init() {
	proto.RegisterType((*SESkillCalReq)(nil), "ULZProto.SESkillCalReq")
	proto.RegisterType((*SESkillCalResp)(nil), "ULZProto.SESkillCalResp")
	proto.RegisterType((*SEDiceCalReq)(nil), "ULZProto.SEDiceCalReq")
	proto.RegisterType((*SEDiceCalResp)(nil), "ULZProto.SEDiceCalResp")
	proto.RegisterType((*DiceResultSet)(nil), "ULZProto.DiceResultSet")
	proto.RegisterType((*SEEffectCalReq)(nil), "ULZProto.SEEffectCalReq")
	proto.RegisterType((*SEEffectCalResp)(nil), "ULZProto.SEEffectCalResp")
}

func init() { proto.RegisterFile("message.proto", fileDescriptor_33c57e4bae7b9afd) }

var fileDescriptor_33c57e4bae7b9afd = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0x92, 0x65, 0x5b, 0x6f, 0x97, 0x0e, 0xcc, 0x57, 0x00, 0x21, 0xaa, 0x48, 0xa0, 0x3e,
	0x15, 0x89, 0xf1, 0x80, 0x04, 0x42, 0x1a, 0x5d, 0x85, 0x26, 0x0d, 0x09, 0xb9, 0x1b, 0x0f, 0x7d,
	0x89, 0x4c, 0x72, 0x5b, 0x59, 0xd8, 0x49, 0x88, 0xbd, 0x3e, 0xf0, 0x0b, 0xf8, 0x2d, 0xfc, 0x2b,
	0xfe, 0x09, 0xb2, 0x9d, 0xb4, 0x89, 0x40, 0x7d, 0xf3, 0xf5, 0x3d, 0xc7, 0xf7, 0x9c, 0x93, 0x1b,
	0x88, 0x24, 0x2a, 0xc5, 0xd6, 0x38, 0xad, 0xea, 0x52, 0x97, 0xe4, 0xf8, 0xe6, 0x6a, 0xf9, 0xc5,
	0x9c, 0x9e, 0xc0, 0x05, 0xd3, 0xcc, 0xdd, 0x26, 0xbf, 0x3c, 0x88, 0x16, 0xf3, 0xc5, 0x77, 0x2e,
	0xc4, 0x8c, 0x09, 0x8a, 0x3f, 0xc8, 0x1b, 0x18, 0xf2, 0x22, 0x2b, 0x25, 0xa6, 0x19, 0xab, 0xf3,
	0xd8, 0x1b, 0x07, 0x93, 0xe1, 0xeb, 0x7b, 0xd3, 0x96, 0x3d, 0x9d, 0x6f, 0xb0, 0xd0, 0x33, 0x56,
	0xe7, 0x14, 0x1c, 0xce, 0x9c, 0xc9, 0x4b, 0x38, 0x58, 0x21, 0xd3, 0xb1, 0x6f, 0xe1, 0x64, 0x07,
	0xb7, 0x4f, 0x2f, 0x50, 0x53, 0xdb, 0x27, 0x8f, 0xe1, 0x78, 0x55, 0x97, 0x32, 0xcd, 0x04, 0x8f,
	0x83, 0xb1, 0x37, 0x19, 0xd0, 0x23, 0x53, 0xcf, 0x04, 0x4f, 0x04, 0x8c, 0xba, 0x4a, 0x54, 0x45,
	0x9e, 0x01, 0xd4, 0xa8, 0x6e, 0x85, 0x4e, 0x37, 0x4c, 0xc4, 0xde, 0xd8, 0x9b, 0x84, 0x74, 0xe0,
	0x6e, 0xbe, 0x32, 0x41, 0xde, 0x41, 0x84, 0xab, 0x15, 0x66, 0x3a, 0x75, 0x77, 0xcd, 0xf0, 0x87,
	0x1d, 0xad, 0xb6, 0x4d, 0x6d, 0x97, 0x9e, 0x60, 0xa7, 0x4a, 0xce, 0xe1, 0x64, 0x31, 0xbf, 0xe0,
	0x19, 0x36, 0xb6, 0x9f, 0x6f, 0x6d, 0xe7, 0x3c, 0xc3, 0x66, 0x58, 0xe3, 0xd0, 0xc0, 0xc8, 0x1d,
	0x08, 0x58, 0x66, 0x66, 0x98, 0x86, 0x39, 0x26, 0x97, 0x26, 0xba, 0xed, 0x13, 0xaa, 0x22, 0x6f,
	0x61, 0x68, 0xc8, 0xad, 0x1c, 0x17, 0xdd, 0xa3, 0x9d, 0x1c, 0x83, 0x75, 0xe3, 0x4d, 0x20, 0x90,
	0x6f, 0xcb, 0xe4, 0x05, 0x44, 0xbd, 0x26, 0xb9, 0x0f, 0xe1, 0x86, 0x89, 0x5b, 0xb4, 0x8f, 0x84,
	0xd4, 0x15, 0xc9, 0x1f, 0xcf, 0x64, 0xe4, 0x5c, 0x35, 0xba, 0x47, 0xe0, 0xf3, 0xdc, 0xca, 0x1d,
	0x50, 0x9f, 0xe7, 0xe4, 0x0c, 0x06, 0x36, 0x60, 0xcd, 0x25, 0x5a, 0xb1, 0xff, 0x09, 0xe4, 0x9a,
	0x4b, 0x5e, 0xac, 0xa9, 0xfd, 0x12, 0xd7, 0x5c, 0x22, 0x79, 0x05, 0x47, 0xba, 0x74, 0x94, 0x60,
	0x2f, 0xe5, 0x50, 0x97, 0x96, 0xf0, 0x01, 0x4e, 0xd7, 0x4c, 0xa2, 0x42, 0x9d, 0xf2, 0x42, 0x69,
	0x56, 0xe8, 0xf8, 0xc0, 0x12, 0x1f, 0xec, 0x88, 0x9f, 0x98, 0x44, 0xb3, 0x69, 0xc6, 0xeb, 0xa8,
	0x41, 0x5f, 0x3a, 0x70, 0x6f, 0x0d, 0xc2, 0xfe, 0x1a, 0xfc, 0x84, 0xd3, 0x9e, 0x45, 0x55, 0xfd,
	0xe3, 0xb1, 0xcb, 0xf6, 0x7b, 0x6c, 0xf2, 0x1e, 0xda, 0x51, 0xed, 0x57, 0x08, 0xf6, 0xe9, 0x8a,
	0x1a, 0xb0, 0x0b, 0xfe, 0xe3, 0xd3, 0x65, 0x68, 0x7f, 0x8b, 0xdf, 0xfe, 0xdd, 0x9b, 0xab, 0xe5,
	0xb9, 0x52, 0xa8, 0xa7, 0x96, 0xf4, 0xb9, 0xcc, 0xbf, 0x1d, 0xda, 0xd6, 0xd9, 0xdf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x5e, 0x38, 0xee, 0x6c, 0x58, 0x03, 0x00, 0x00,
}
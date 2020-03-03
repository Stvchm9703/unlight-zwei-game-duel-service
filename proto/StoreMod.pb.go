// Code generated by protoc-gen-go. DO NOT EDIT.
// source: StoreMod.proto

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

// Move-Phase-SnapMod
type MovePhaseSnapMod struct {
	//  flaging
	Turns                int32        `protobuf:"varint,1,opt,name=turns,proto3" json:"turns,omitempty"`
	IsHostReady          bool         `protobuf:"varint,2,opt,name=is_host_ready,json=isHostReady,proto3" json:"is_host_ready,omitempty"`
	IsDuelReady          bool         `protobuf:"varint,3,opt,name=is_duel_ready,json=isDuelReady,proto3" json:"is_duel_ready,omitempty"`
	HostVal              int32        `protobuf:"varint,4,opt,name=host_val,json=hostVal,proto3" json:"host_val,omitempty"`
	DuelVal              int32        `protobuf:"varint,5,opt,name=duel_val,json=duelVal,proto3" json:"duel_val,omitempty"`
	HostOpt              MovePhaseOpt `protobuf:"varint,6,opt,name=host_opt,json=hostOpt,proto3,enum=ULZProto.MovePhaseOpt" json:"host_opt,omitempty"`
	DuelOpt              MovePhaseOpt `protobuf:"varint,7,opt,name=duel_opt,json=duelOpt,proto3,enum=ULZProto.MovePhaseOpt" json:"duel_opt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MovePhaseSnapMod) Reset()         { *m = MovePhaseSnapMod{} }
func (m *MovePhaseSnapMod) String() string { return proto.CompactTextString(m) }
func (*MovePhaseSnapMod) ProtoMessage()    {}
func (*MovePhaseSnapMod) Descriptor() ([]byte, []int) {
	return fileDescriptor_98fa4925a5e51b0e, []int{0}
}

func (m *MovePhaseSnapMod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovePhaseSnapMod.Unmarshal(m, b)
}
func (m *MovePhaseSnapMod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovePhaseSnapMod.Marshal(b, m, deterministic)
}
func (m *MovePhaseSnapMod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovePhaseSnapMod.Merge(m, src)
}
func (m *MovePhaseSnapMod) XXX_Size() int {
	return xxx_messageInfo_MovePhaseSnapMod.Size(m)
}
func (m *MovePhaseSnapMod) XXX_DiscardUnknown() {
	xxx_messageInfo_MovePhaseSnapMod.DiscardUnknown(m)
}

var xxx_messageInfo_MovePhaseSnapMod proto.InternalMessageInfo

func (m *MovePhaseSnapMod) GetTurns() int32 {
	if m != nil {
		return m.Turns
	}
	return 0
}

func (m *MovePhaseSnapMod) GetIsHostReady() bool {
	if m != nil {
		return m.IsHostReady
	}
	return false
}

func (m *MovePhaseSnapMod) GetIsDuelReady() bool {
	if m != nil {
		return m.IsDuelReady
	}
	return false
}

func (m *MovePhaseSnapMod) GetHostVal() int32 {
	if m != nil {
		return m.HostVal
	}
	return 0
}

func (m *MovePhaseSnapMod) GetDuelVal() int32 {
	if m != nil {
		return m.DuelVal
	}
	return 0
}

func (m *MovePhaseSnapMod) GetHostOpt() MovePhaseOpt {
	if m != nil {
		return m.HostOpt
	}
	return MovePhaseOpt_NO_MOVE
}

func (m *MovePhaseSnapMod) GetDuelOpt() MovePhaseOpt {
	if m != nil {
		return m.DuelOpt
	}
	return MovePhaseOpt_NO_MOVE
}

// AD-Phase-SnapMod
type ADPhaseSnapMod struct {
	Turns                int32          `protobuf:"varint,1,opt,name=turns,proto3" json:"turns,omitempty"`
	FirstAttack          PlayerSide     `protobuf:"varint,2,opt,name=first_attack,json=firstAttack,proto3,enum=ULZProto.PlayerSide" json:"first_attack,omitempty"`
	CurrPhase            PlayerSide     `protobuf:"varint,3,opt,name=curr_phase,json=currPhase,proto3,enum=ULZProto.PlayerSide" json:"curr_phase,omitempty"`
	EventPhase           EventHookPhase `protobuf:"varint,4,opt,name=event_phase,json=eventPhase,proto3,enum=ULZProto.EventHookPhase" json:"event_phase,omitempty"`
	HostVal              int32          `protobuf:"varint,5,opt,name=host_val,json=hostVal,proto3" json:"host_val,omitempty"`
	DuelVal              int32          `protobuf:"varint,6,opt,name=duel_val,json=duelVal,proto3" json:"duel_val,omitempty"`
	HostCard             []*EventCard   `protobuf:"bytes,7,rep,name=host_card,json=hostCard,proto3" json:"host_card,omitempty"`
	DuelCard             []*EventCard   `protobuf:"bytes,8,rep,name=duel_card,json=duelCard,proto3" json:"duel_card,omitempty"`
	HostTrigSkl          []*SkillSet    `protobuf:"bytes,9,rep,name=host_trig_skl,json=hostTrigSkl,proto3" json:"host_trig_skl,omitempty"`
	DuelTrigSkl          []*SkillSet    `protobuf:"bytes,10,rep,name=duel_trig_skl,json=duelTrigSkl,proto3" json:"duel_trig_skl,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ADPhaseSnapMod) Reset()         { *m = ADPhaseSnapMod{} }
func (m *ADPhaseSnapMod) String() string { return proto.CompactTextString(m) }
func (*ADPhaseSnapMod) ProtoMessage()    {}
func (*ADPhaseSnapMod) Descriptor() ([]byte, []int) {
	return fileDescriptor_98fa4925a5e51b0e, []int{1}
}

func (m *ADPhaseSnapMod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ADPhaseSnapMod.Unmarshal(m, b)
}
func (m *ADPhaseSnapMod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ADPhaseSnapMod.Marshal(b, m, deterministic)
}
func (m *ADPhaseSnapMod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ADPhaseSnapMod.Merge(m, src)
}
func (m *ADPhaseSnapMod) XXX_Size() int {
	return xxx_messageInfo_ADPhaseSnapMod.Size(m)
}
func (m *ADPhaseSnapMod) XXX_DiscardUnknown() {
	xxx_messageInfo_ADPhaseSnapMod.DiscardUnknown(m)
}

var xxx_messageInfo_ADPhaseSnapMod proto.InternalMessageInfo

func (m *ADPhaseSnapMod) GetTurns() int32 {
	if m != nil {
		return m.Turns
	}
	return 0
}

func (m *ADPhaseSnapMod) GetFirstAttack() PlayerSide {
	if m != nil {
		return m.FirstAttack
	}
	return PlayerSide_HOST
}

func (m *ADPhaseSnapMod) GetCurrPhase() PlayerSide {
	if m != nil {
		return m.CurrPhase
	}
	return PlayerSide_HOST
}

func (m *ADPhaseSnapMod) GetEventPhase() EventHookPhase {
	if m != nil {
		return m.EventPhase
	}
	return EventHookPhase_gameset_start
}

func (m *ADPhaseSnapMod) GetHostVal() int32 {
	if m != nil {
		return m.HostVal
	}
	return 0
}

func (m *ADPhaseSnapMod) GetDuelVal() int32 {
	if m != nil {
		return m.DuelVal
	}
	return 0
}

func (m *ADPhaseSnapMod) GetHostCard() []*EventCard {
	if m != nil {
		return m.HostCard
	}
	return nil
}

func (m *ADPhaseSnapMod) GetDuelCard() []*EventCard {
	if m != nil {
		return m.DuelCard
	}
	return nil
}

func (m *ADPhaseSnapMod) GetHostTrigSkl() []*SkillSet {
	if m != nil {
		return m.HostTrigSkl
	}
	return nil
}

func (m *ADPhaseSnapMod) GetDuelTrigSkl() []*SkillSet {
	if m != nil {
		return m.DuelTrigSkl
	}
	return nil
}

func init() {
	proto.RegisterType((*MovePhaseSnapMod)(nil), "ULZProto.MovePhaseSnapMod")
	proto.RegisterType((*ADPhaseSnapMod)(nil), "ULZProto.ADPhaseSnapMod")
}

func init() { proto.RegisterFile("StoreMod.proto", fileDescriptor_98fa4925a5e51b0e) }

var fileDescriptor_98fa4925a5e51b0e = []byte{
	// 428 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x40, 0xe5, 0xa4, 0x4e, 0x9c, 0x35, 0xb5, 0xc0, 0x44, 0xc8, 0x94, 0x8b, 0x95, 0x93, 0x4f,
	0x16, 0xa4, 0x12, 0x88, 0x63, 0x20, 0x48, 0x3d, 0xd4, 0x4a, 0x64, 0x53, 0x0e, 0xb9, 0x58, 0x4b,
	0x76, 0x69, 0x2d, 0xaf, 0xb2, 0xd6, 0xee, 0x38, 0x52, 0xff, 0x81, 0x2f, 0xe1, 0x8f, 0xf8, 0x1b,
	0x34, 0x63, 0xb7, 0xa1, 0x48, 0x89, 0xb8, 0xed, 0xcc, 0xbc, 0x37, 0xe3, 0x9d, 0x35, 0x0b, 0x0a,
	0xd0, 0x46, 0x66, 0x5a, 0xa4, 0x8d, 0xd1, 0xa0, 0x43, 0xef, 0xe6, 0x7a, 0xb3, 0xc6, 0xd3, 0x05,
	0x5b, 0x72, 0xe0, 0x5d, 0xf6, 0x62, 0xfa, 0x65, 0x2f, 0x77, 0x70, 0xa5, 0x75, 0xbd, 0xbe, 0xe3,
	0x56, 0x76, 0xd9, 0xd9, 0xcf, 0x01, 0x7b, 0x9e, 0xe9, 0xbd, 0xa4, 0x5c, 0xb1, 0xe3, 0x4d, 0xa6,
	0x45, 0x38, 0x65, 0x2e, 0xb4, 0x66, 0x67, 0x23, 0x27, 0x76, 0x12, 0x37, 0xef, 0x82, 0x70, 0xc6,
	0xce, 0x2b, 0x5b, 0xde, 0x69, 0x0b, 0xa5, 0x91, 0x5c, 0xdc, 0x47, 0x83, 0xd8, 0x49, 0xbc, 0xdc,
	0xaf, 0xec, 0x95, 0xb6, 0x90, 0x63, 0xaa, 0x67, 0x44, 0x2b, 0x55, 0xcf, 0x0c, 0x1f, 0x98, 0x65,
	0x2b, 0x55, 0xc7, 0xbc, 0x66, 0x1e, 0x35, 0xd9, 0x73, 0x15, 0x9d, 0xd1, 0x80, 0x31, 0xc6, 0xdf,
	0xb8, 0xc2, 0x12, 0xb9, 0x58, 0x72, 0xbb, 0x12, 0xc6, 0x58, 0x7a, 0xd7, 0x5b, 0xba, 0x81, 0x68,
	0x14, 0x3b, 0x49, 0x30, 0x7f, 0x95, 0x3e, 0xdc, 0x33, 0x7d, 0xbc, 0xc1, 0xaa, 0x81, 0xae, 0xdb,
	0xaa, 0x01, 0x54, 0xa8, 0x1b, 0x2a, 0xe3, 0xd3, 0x0a, 0x72, 0xab, 0x06, 0x66, 0xbf, 0x87, 0x2c,
	0x58, 0x2c, 0xff, 0x63, 0x19, 0x1f, 0xd8, 0xb3, 0x1f, 0x95, 0xb1, 0x50, 0x72, 0x00, 0xbe, 0xad,
	0x69, 0x17, 0xc1, 0x7c, 0x7a, 0xe8, 0xbf, 0x56, 0xfc, 0x5e, 0x9a, 0xa2, 0x12, 0x32, 0xf7, 0x89,
	0x5c, 0x10, 0x18, 0x5e, 0x32, 0xb6, 0x6d, 0x8d, 0x29, 0x1b, 0x9c, 0x41, 0xeb, 0x39, 0xa6, 0x4d,
	0x90, 0xa3, 0x4f, 0x09, 0x3f, 0x32, 0x5f, 0xe2, 0xeb, 0xf5, 0xd6, 0x19, 0x59, 0xd1, 0xc1, 0x7a,
	0xfa, 0xb4, 0x39, 0x23, 0xb8, 0x53, 0xff, 0xde, 0xb6, 0x7b, 0x7c, 0xdb, 0xa3, 0xa7, 0xdb, 0x7e,
	0xcb, 0x26, 0x64, 0x6d, 0xb9, 0x11, 0xd1, 0x38, 0x1e, 0x26, 0xfe, 0xfc, 0xe5, 0x3f, 0xe3, 0x3e,
	0x73, 0x23, 0x72, 0xea, 0x8d, 0x27, 0x34, 0xa8, 0x19, 0x19, 0xde, 0x09, 0x03, 0x29, 0x32, 0xde,
	0xb3, 0x73, 0x9a, 0x01, 0xa6, 0xba, 0x2d, 0x6d, 0xad, 0xa2, 0x09, 0x59, 0xe1, 0xc1, 0x2a, 0xea,
	0x4a, 0xa9, 0x42, 0x42, 0xee, 0x23, 0xf8, 0xd5, 0x54, 0xb7, 0x45, 0xad, 0xd0, 0xa3, 0x49, 0x8f,
	0x1e, 0x3b, 0xee, 0x21, 0xd8, 0x7b, 0x9f, 0xde, 0x6c, 0x5c, 0xfa, 0xe7, 0x7f, 0x0d, 0x5e, 0xdc,
	0x5c, 0x6f, 0x16, 0xd6, 0x4a, 0x48, 0x89, 0xcf, 0xb4, 0xf8, 0x3e, 0xa2, 0xd2, 0xe5, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xcc, 0x59, 0xc8, 0xda, 0x4c, 0x03, 0x00, 0x00,
}

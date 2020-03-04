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
	HostCard             []*EventCard `protobuf:"bytes,8,rep,name=host_card,json=hostCard,proto3" json:"host_card,omitempty"`
	DuelCard             []*EventCard `protobuf:"bytes,9,rep,name=duel_card,json=duelCard,proto3" json:"duel_card,omitempty"`
	HostTrigSkl          []*SkillSet  `protobuf:"bytes,10,rep,name=host_trig_skl,json=hostTrigSkl,proto3" json:"host_trig_skl,omitempty"`
	DuelTrigSkl          []*SkillSet  `protobuf:"bytes,11,rep,name=duel_trig_skl,json=duelTrigSkl,proto3" json:"duel_trig_skl,omitempty"`
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

func (m *MovePhaseSnapMod) GetHostCard() []*EventCard {
	if m != nil {
		return m.HostCard
	}
	return nil
}

func (m *MovePhaseSnapMod) GetDuelCard() []*EventCard {
	if m != nil {
		return m.DuelCard
	}
	return nil
}

func (m *MovePhaseSnapMod) GetHostTrigSkl() []*SkillSet {
	if m != nil {
		return m.HostTrigSkl
	}
	return nil
}

func (m *MovePhaseSnapMod) GetDuelTrigSkl() []*SkillSet {
	if m != nil {
		return m.DuelTrigSkl
	}
	return nil
}

// AD-Phase-SnapMod
type ADPhaseSnapMod struct {
	Turns                int32          `protobuf:"varint,1,opt,name=turns,proto3" json:"turns,omitempty"`
	FirstAttack          PlayerSide     `protobuf:"varint,2,opt,name=first_attack,json=firstAttack,proto3,enum=ULZProto.PlayerSide" json:"first_attack,omitempty"`
	CurrAttacker         PlayerSide     `protobuf:"varint,3,opt,name=curr_attacker,json=currAttacker,proto3,enum=ULZProto.PlayerSide" json:"curr_attacker,omitempty"`
	EventPhase           EventHookPhase `protobuf:"varint,4,opt,name=event_phase,json=eventPhase,proto3,enum=ULZProto.EventHookPhase" json:"event_phase,omitempty"`
	AttackVal            int32          `protobuf:"varint,5,opt,name=attack_val,json=attackVal,proto3" json:"attack_val,omitempty"`
	DefenceVal           int32          `protobuf:"varint,6,opt,name=defence_val,json=defenceVal,proto3" json:"defence_val,omitempty"`
	AttackCard           []*EventCard   `protobuf:"bytes,7,rep,name=attack_card,json=attackCard,proto3" json:"attack_card,omitempty"`
	DefenceCard          []*EventCard   `protobuf:"bytes,8,rep,name=defence_card,json=defenceCard,proto3" json:"defence_card,omitempty"`
	AttackTrigSkl        []*SkillSet    `protobuf:"bytes,9,rep,name=attack_trig_skl,json=attackTrigSkl,proto3" json:"attack_trig_skl,omitempty"`
	DefenceTrigSkl       []*SkillSet    `protobuf:"bytes,10,rep,name=defence_trig_skl,json=defenceTrigSkl,proto3" json:"defence_trig_skl,omitempty"`
	IsProcessed          bool           `protobuf:"varint,11,opt,name=is_processed,json=isProcessed,proto3" json:"is_processed,omitempty"`
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

func (m *ADPhaseSnapMod) GetCurrAttacker() PlayerSide {
	if m != nil {
		return m.CurrAttacker
	}
	return PlayerSide_HOST
}

func (m *ADPhaseSnapMod) GetEventPhase() EventHookPhase {
	if m != nil {
		return m.EventPhase
	}
	return EventHookPhase_gameset_start
}

func (m *ADPhaseSnapMod) GetAttackVal() int32 {
	if m != nil {
		return m.AttackVal
	}
	return 0
}

func (m *ADPhaseSnapMod) GetDefenceVal() int32 {
	if m != nil {
		return m.DefenceVal
	}
	return 0
}

func (m *ADPhaseSnapMod) GetAttackCard() []*EventCard {
	if m != nil {
		return m.AttackCard
	}
	return nil
}

func (m *ADPhaseSnapMod) GetDefenceCard() []*EventCard {
	if m != nil {
		return m.DefenceCard
	}
	return nil
}

func (m *ADPhaseSnapMod) GetAttackTrigSkl() []*SkillSet {
	if m != nil {
		return m.AttackTrigSkl
	}
	return nil
}

func (m *ADPhaseSnapMod) GetDefenceTrigSkl() []*SkillSet {
	if m != nil {
		return m.DefenceTrigSkl
	}
	return nil
}

func (m *ADPhaseSnapMod) GetIsProcessed() bool {
	if m != nil {
		return m.IsProcessed
	}
	return false
}

func init() {
	proto.RegisterType((*MovePhaseSnapMod)(nil), "ULZProto.MovePhaseSnapMod")
	proto.RegisterType((*ADPhaseSnapMod)(nil), "ULZProto.ADPhaseSnapMod")
}

func init() { proto.RegisterFile("StoreMod.proto", fileDescriptor_98fa4925a5e51b0e) }

var fileDescriptor_98fa4925a5e51b0e = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x95, 0xa6, 0xf9, 0x37, 0x4e, 0x4c, 0x31, 0x11, 0x32, 0x45, 0x88, 0x90, 0x53, 0x4e,
	0x11, 0x04, 0x04, 0x2a, 0xe2, 0x12, 0x08, 0x52, 0x0f, 0x8d, 0x1a, 0xd9, 0x94, 0x43, 0x2e, 0xd6,
	0x92, 0x9d, 0xb6, 0x96, 0xad, 0xac, 0xb5, 0xbb, 0x89, 0xd4, 0xb7, 0xe1, 0xcc, 0x53, 0xf1, 0x28,
	0x68, 0xc6, 0xeb, 0x04, 0x21, 0xa5, 0xf4, 0xe6, 0x9d, 0xf9, 0x7e, 0xf3, 0xd9, 0xfb, 0x4d, 0x02,
	0x7e, 0x6c, 0x95, 0xc6, 0xb9, 0x92, 0xe3, 0x42, 0x2b, 0xab, 0x82, 0xf6, 0xd5, 0xc5, 0x72, 0x41,
	0x4f, 0xa7, 0x30, 0x13, 0x56, 0x94, 0xd5, 0xd3, 0xfe, 0xd7, 0x2d, 0xae, 0xed, 0xb9, 0x52, 0xd9,
	0xe2, 0x56, 0x18, 0x2c, 0xab, 0xc3, 0xdf, 0x75, 0x38, 0x99, 0xab, 0x2d, 0x72, 0x2d, 0x5e, 0x8b,
	0x62, 0xae, 0x64, 0xd0, 0x87, 0x86, 0xdd, 0xe8, 0xb5, 0x09, 0x6b, 0x83, 0xda, 0xa8, 0x11, 0x95,
	0x87, 0x60, 0x08, 0xbd, 0xd4, 0x24, 0xb7, 0xca, 0xd8, 0x44, 0xa3, 0x90, 0x77, 0xe1, 0xd1, 0xa0,
	0x36, 0x6a, 0x47, 0x5e, 0x6a, 0xce, 0x95, 0xb1, 0x11, 0x95, 0x9c, 0x46, 0x6e, 0x30, 0x77, 0x9a,
	0x7a, 0xa5, 0x99, 0x6d, 0x30, 0x2f, 0x35, 0xcf, 0xa0, 0xcd, 0x43, 0xb6, 0x22, 0x0f, 0x8f, 0xd9,
	0xa0, 0x45, 0xe7, 0xef, 0x22, 0xa7, 0x16, 0xb3, 0xd4, 0x6a, 0x94, 0x2d, 0x3a, 0x53, 0xeb, 0x8d,
	0xa3, 0x54, 0x61, 0xc3, 0xe6, 0xa0, 0x36, 0xf2, 0x27, 0x4f, 0xc7, 0xd5, 0x77, 0x8e, 0x77, 0x5f,
	0x70, 0x59, 0xd8, 0x72, 0xda, 0x65, 0x61, 0x09, 0xe1, 0x69, 0x84, 0xb4, 0xee, 0x47, 0x48, 0x47,
	0xc8, 0x6b, 0xe8, 0xb0, 0xcb, 0x4a, 0x68, 0x19, 0xb6, 0x07, 0xf5, 0x91, 0x37, 0x79, 0xb2, 0x67,
	0xf8, 0x06, 0xbf, 0x08, 0x2d, 0x23, 0x7e, 0x17, 0x7a, 0x22, 0x82, 0x4d, 0x98, 0xe8, 0xdc, 0x43,
	0x90, 0x8a, 0x89, 0xf7, 0xd0, 0x63, 0x0f, 0xab, 0xd3, 0x9b, 0xc4, 0x64, 0x79, 0x08, 0x4c, 0x05,
	0x7b, 0x2a, 0xce, 0xd2, 0x3c, 0x8f, 0xd1, 0x46, 0x1e, 0x09, 0xbf, 0xe9, 0xf4, 0x26, 0xce, 0x72,
	0xe2, 0xd8, 0x69, 0xc7, 0x79, 0x87, 0x39, 0x12, 0x3a, 0x6e, 0xf8, 0xf3, 0x18, 0xfc, 0xe9, 0xec,
	0x01, 0x01, 0x7f, 0x80, 0xee, 0x75, 0xaa, 0x8d, 0x4d, 0x84, 0xb5, 0x62, 0x95, 0x71, 0xbe, 0xfe,
	0xa4, 0xbf, 0x9f, 0xbf, 0xc8, 0xc5, 0x1d, 0xea, 0x38, 0x95, 0x18, 0x79, 0xac, 0x9c, 0xb2, 0x30,
	0x38, 0x83, 0xde, 0x6a, 0xa3, 0xb5, 0xe3, 0x50, 0x73, 0xea, 0x87, 0xc8, 0x2e, 0x49, 0xa7, 0x4e,
	0x19, 0x9c, 0x81, 0x87, 0x74, 0x47, 0x49, 0x41, 0xef, 0xc7, 0xfb, 0xe0, 0x4f, 0xc2, 0x7f, 0x2e,
	0x70, 0xb7, 0xb4, 0x11, 0xb0, 0x98, 0x9f, 0x83, 0x17, 0x00, 0xa5, 0xe1, 0x5f, 0xeb, 0xd2, 0x29,
	0x2b, 0xb4, 0x30, 0x2f, 0xc1, 0x93, 0x78, 0x8d, 0xeb, 0x15, 0x72, 0xbf, 0xc9, 0x7d, 0x70, 0x25,
	0x12, 0xbc, 0x03, 0xcf, 0xf1, 0x9c, 0x5d, 0xeb, 0x70, 0x76, 0xce, 0xc7, 0xa5, 0xd7, 0xad, 0xc6,
	0xfe, 0x6f, 0x49, 0x2a, 0x7f, 0xe6, 0x3e, 0xc2, 0x23, 0xe7, 0xb6, 0xcb, 0xaf, 0x73, 0x30, 0xbf,
	0x5e, 0x29, 0xad, 0x92, 0xff, 0x04, 0x27, 0x95, 0xe7, 0x03, 0x96, 0xc6, 0x77, 0xda, 0x8a, 0x7e,
	0x05, 0xdd, 0xd4, 0x24, 0x85, 0x56, 0x2b, 0x34, 0x06, 0x65, 0xe8, 0x55, 0x3f, 0xc9, 0x45, 0x55,
	0xfa, 0xfc, 0x7c, 0xd9, 0xe0, 0xbf, 0x83, 0x5f, 0x47, 0x8f, 0xaf, 0x2e, 0x96, 0x53, 0x63, 0xd0,
	0x8e, 0x79, 0xea, 0x5c, 0xc9, 0x1f, 0x4d, 0x6e, 0xbd, 0xfd, 0x13, 0x00, 0x00, 0xff, 0xff, 0x83,
	0x32, 0xcb, 0x4e, 0x67, 0x04, 0x00, 0x00,
}

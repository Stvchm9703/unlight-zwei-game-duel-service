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
	IsHostReady          bool           `protobuf:"varint,12,opt,name=is_host_ready,json=isHostReady,proto3" json:"is_host_ready,omitempty"`
	IsDuelReady          bool           `protobuf:"varint,13,opt,name=is_duel_ready,json=isDuelReady,proto3" json:"is_duel_ready,omitempty"`
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

func (m *ADPhaseSnapMod) GetIsHostReady() bool {
	if m != nil {
		return m.IsHostReady
	}
	return false
}

func (m *ADPhaseSnapMod) GetIsDuelReady() bool {
	if m != nil {
		return m.IsDuelReady
	}
	return false
}

// // Effect-Status-SnapMod
type EffectNodeSnapMod struct {
	Turns                int32           `protobuf:"varint,1,opt,name=turns,proto3" json:"turns,omitempty"`
	PendingEf            []*EffectResult `protobuf:"bytes,2,rep,name=pending_ef,json=pendingEf,proto3" json:"pending_ef,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EffectNodeSnapMod) Reset()         { *m = EffectNodeSnapMod{} }
func (m *EffectNodeSnapMod) String() string { return proto.CompactTextString(m) }
func (*EffectNodeSnapMod) ProtoMessage()    {}
func (*EffectNodeSnapMod) Descriptor() ([]byte, []int) {
	return fileDescriptor_98fa4925a5e51b0e, []int{2}
}

func (m *EffectNodeSnapMod) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EffectNodeSnapMod.Unmarshal(m, b)
}
func (m *EffectNodeSnapMod) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EffectNodeSnapMod.Marshal(b, m, deterministic)
}
func (m *EffectNodeSnapMod) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EffectNodeSnapMod.Merge(m, src)
}
func (m *EffectNodeSnapMod) XXX_Size() int {
	return xxx_messageInfo_EffectNodeSnapMod.Size(m)
}
func (m *EffectNodeSnapMod) XXX_DiscardUnknown() {
	xxx_messageInfo_EffectNodeSnapMod.DiscardUnknown(m)
}

var xxx_messageInfo_EffectNodeSnapMod proto.InternalMessageInfo

func (m *EffectNodeSnapMod) GetTurns() int32 {
	if m != nil {
		return m.Turns
	}
	return 0
}

func (m *EffectNodeSnapMod) GetPendingEf() []*EffectResult {
	if m != nil {
		return m.PendingEf
	}
	return nil
}

func init() {
	proto.RegisterType((*MovePhaseSnapMod)(nil), "ULZProto.MovePhaseSnapMod")
	proto.RegisterType((*ADPhaseSnapMod)(nil), "ULZProto.ADPhaseSnapMod")
	proto.RegisterType((*EffectNodeSnapMod)(nil), "ULZProto.EffectNodeSnapMod")
}

func init() { proto.RegisterFile("StoreMod.proto", fileDescriptor_98fa4925a5e51b0e) }

var fileDescriptor_98fa4925a5e51b0e = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcd, 0x6e, 0xda, 0x40,
	0x10, 0xc7, 0x05, 0x84, 0xaf, 0x31, 0xb8, 0x89, 0x8b, 0x2a, 0x37, 0x55, 0x55, 0xca, 0x89, 0x13,
	0x6a, 0xe9, 0x97, 0x52, 0xf5, 0x42, 0x0b, 0x52, 0x0e, 0xa1, 0x41, 0x76, 0xd3, 0x03, 0x17, 0x77,
	0xeb, 0x5d, 0x13, 0x0b, 0xcb, 0x6b, 0xed, 0xae, 0x91, 0xf2, 0x4a, 0x7d, 0xaa, 0x1e, 0xfb, 0x18,
	0xd5, 0x8e, 0xd7, 0x10, 0x25, 0x22, 0xe1, 0xb6, 0x3b, 0xf3, 0xff, 0xcd, 0xdf, 0xbb, 0x33, 0x6b,
	0xb0, 0x7d, 0xc5, 0x05, 0x9b, 0x73, 0x3a, 0xca, 0x04, 0x57, 0xdc, 0x69, 0x5d, 0x5d, 0x2c, 0x17,
	0x7a, 0x75, 0x0a, 0x53, 0xa2, 0x48, 0x11, 0x3d, 0xed, 0xcd, 0x36, 0x2c, 0x55, 0xe7, 0x9c, 0xaf,
	0x17, 0xd7, 0x44, 0xb2, 0x22, 0x3a, 0xf8, 0x5b, 0x83, 0xe3, 0x39, 0xdf, 0x30, 0x8c, 0xf9, 0x29,
	0xc9, 0xe6, 0x9c, 0x3a, 0x3d, 0xa8, 0xab, 0x5c, 0xa4, 0xd2, 0xad, 0xf4, 0x2b, 0xc3, 0xba, 0x57,
	0x6c, 0x9c, 0x01, 0x74, 0x63, 0x19, 0x5c, 0x73, 0xa9, 0x02, 0xc1, 0x08, 0xbd, 0x71, 0xab, 0xfd,
	0xca, 0xb0, 0xe5, 0x59, 0xb1, 0x3c, 0xe7, 0x52, 0x79, 0x3a, 0x64, 0x34, 0x34, 0x67, 0x89, 0xd1,
	0xd4, 0x4a, 0xcd, 0x34, 0x67, 0x49, 0xa1, 0x79, 0x0e, 0x2d, 0x2c, 0xb2, 0x21, 0x89, 0x7b, 0x84,
	0x06, 0x4d, 0xbd, 0xff, 0x49, 0x12, 0x9d, 0x42, 0x56, 0xa7, 0xea, 0x45, 0x4a, 0xef, 0x75, 0xea,
	0xad, 0xa1, 0x78, 0xa6, 0xdc, 0x46, 0xbf, 0x32, 0xb4, 0xc7, 0xcf, 0x46, 0xe5, 0x39, 0x47, 0xdb,
	0x13, 0x5c, 0x66, 0xaa, 0xa8, 0x76, 0x99, 0x29, 0x8d, 0x60, 0x35, 0x8d, 0x34, 0x1f, 0x46, 0xb4,
	0x4e, 0x23, 0x6f, 0xa0, 0x8d, 0x2e, 0x21, 0x11, 0xd4, 0x6d, 0xf5, 0x6b, 0x43, 0x6b, 0xfc, 0x74,
	0xc7, 0xe0, 0x0d, 0x7e, 0x23, 0x82, 0x7a, 0xf8, 0x2d, 0x7a, 0xa5, 0x09, 0x34, 0x41, 0xa2, 0xfd,
	0x00, 0xa1, 0x55, 0x48, 0x7c, 0x84, 0x2e, 0x7a, 0x28, 0x11, 0xaf, 0x02, 0xb9, 0x4e, 0x5c, 0x40,
	0xca, 0xd9, 0x51, 0xfe, 0x3a, 0x4e, 0x12, 0x9f, 0x29, 0xcf, 0xd2, 0xc2, 0x1f, 0x22, 0x5e, 0xf9,
	0xeb, 0x44, 0x73, 0xe8, 0xb4, 0xe5, 0xac, 0xfd, 0x9c, 0x16, 0x1a, 0x6e, 0xf0, 0xef, 0x08, 0xec,
	0xc9, 0xf4, 0x80, 0x06, 0x7f, 0x82, 0x4e, 0x14, 0x0b, 0xa9, 0x02, 0xa2, 0x14, 0x09, 0xd7, 0xd8,
	0x5f, 0x7b, 0xdc, 0xdb, 0xd5, 0x5f, 0x24, 0xe4, 0x86, 0x09, 0x3f, 0xa6, 0xcc, 0xb3, 0x50, 0x39,
	0x41, 0xa1, 0x73, 0x06, 0xdd, 0x30, 0x17, 0xc2, 0x70, 0x4c, 0x60, 0xd7, 0xf7, 0x91, 0x1d, 0x2d,
	0x9d, 0x18, 0xa5, 0x73, 0x06, 0x16, 0xd3, 0x77, 0x14, 0x64, 0xfa, 0xfb, 0x70, 0x1e, 0xec, 0xb1,
	0x7b, 0xe7, 0x02, 0xb7, 0x43, 0xeb, 0x01, 0x8a, 0x71, 0xed, 0xbc, 0x04, 0x28, 0x0c, 0x6f, 0x8d,
	0x4b, 0xbb, 0x88, 0xe8, 0x81, 0x79, 0x05, 0x16, 0x65, 0x11, 0x4b, 0x43, 0x86, 0xf9, 0x06, 0xe6,
	0xc1, 0x84, 0xb4, 0xe0, 0x3d, 0x58, 0x86, 0xc7, 0xde, 0x35, 0xf7, 0xf7, 0xce, 0xf8, 0x98, 0xee,
	0x75, 0xca, 0xb2, 0x8f, 0x0d, 0x49, 0xe9, 0x8f, 0xdc, 0x67, 0x78, 0x62, 0xdc, 0xb6, 0xfd, 0x6b,
	0xef, 0xed, 0x5f, 0xb7, 0x90, 0x96, 0x9d, 0xff, 0x02, 0xc7, 0xa5, 0xe7, 0x01, 0x43, 0x63, 0x1b,
	0x6d, 0x49, 0xbf, 0x86, 0x4e, 0x2c, 0x83, 0x4c, 0xf0, 0x90, 0x49, 0xc9, 0xa8, 0x6b, 0x95, 0x4f,
	0x72, 0x51, 0x86, 0xee, 0x3f, 0xed, 0xce, 0x01, 0x4f, 0xbb, 0x7b, 0xef, 0x69, 0x0f, 0x7e, 0xc1,
	0xc9, 0x2c, 0x8a, 0x58, 0xa8, 0xbe, 0x73, 0xfa, 0xc8, 0xb0, 0x7d, 0x00, 0xc8, 0x58, 0x4a, 0xe3,
	0x74, 0x15, 0xb0, 0xc8, 0xad, 0xe2, 0x69, 0x6e, 0x3d, 0xcf, 0xa2, 0x8c, 0xc7, 0x64, 0x9e, 0x28,
	0xaf, 0x6d, 0x94, 0xb3, 0xe8, 0xeb, 0x8b, 0x65, 0x1d, 0x7f, 0x5c, 0x7f, 0xaa, 0x27, 0x57, 0x17,
	0xcb, 0x89, 0x94, 0x4c, 0x8d, 0x90, 0x98, 0x73, 0xfa, 0xbb, 0x81, 0xa9, 0x77, 0xff, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x09, 0x25, 0x77, 0x37, 0x11, 0x05, 0x00, 0x00,
}

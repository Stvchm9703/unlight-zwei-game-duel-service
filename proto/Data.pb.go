// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Data.proto

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

// -------------------------------------------------------------
// ANCHOR: data-struct for all game struct
// SECTION: Data.proto
type PlayerSide int32

const (
	PlayerSide_HOST   PlayerSide = 0
	PlayerSide_DUELER PlayerSide = 1
)

var PlayerSide_name = map[int32]string{
	0: "HOST",
	1: "DUELER",
}

var PlayerSide_value = map[string]int32{
	"HOST":   0,
	"DUELER": 1,
}

func (x PlayerSide) String() string {
	return proto.EnumName(PlayerSide_name, int32(x))
}

func (PlayerSide) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{0}
}

type EventCardType int32

const (
	EventCardType_NULL    EventCardType = 0
	EventCardType_ATTACK  EventCardType = 1
	EventCardType_DEFENCE EventCardType = 2
	EventCardType_GUN     EventCardType = 3
	EventCardType_MOVE    EventCardType = 4
	EventCardType_STAR    EventCardType = 5
)

var EventCardType_name = map[int32]string{
	0: "NULL",
	1: "ATTACK",
	2: "DEFENCE",
	3: "GUN",
	4: "MOVE",
	5: "STAR",
}

var EventCardType_value = map[string]int32{
	"NULL":    0,
	"ATTACK":  1,
	"DEFENCE": 2,
	"GUN":     3,
	"MOVE":    4,
	"STAR":    5,
}

func (x EventCardType) String() string {
	return proto.EnumName(EventCardType_name, int32(x))
}

func (EventCardType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{1}
}

type EventCardPos int32

const (
	EventCardPos_BLOCK   EventCardPos = 0
	EventCardPos_INSIDE  EventCardPos = 1
	EventCardPos_OUTSIDE EventCardPos = 2
	EventCardPos_DESTROY EventCardPos = 3
)

var EventCardPos_name = map[int32]string{
	0: "BLOCK",
	1: "INSIDE",
	2: "OUTSIDE",
	3: "DESTROY",
}

var EventCardPos_value = map[string]int32{
	"BLOCK":   0,
	"INSIDE":  1,
	"OUTSIDE": 2,
	"DESTROY": 3,
}

func (x EventCardPos) String() string {
	return proto.EnumName(EventCardPos_name, int32(x))
}

func (EventCardPos) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{2}
}

type RangeType int32

const (
	RangeType_SHORT  RangeType = 0
	RangeType_MIDDLE RangeType = 1
	RangeType_LONG   RangeType = 2
)

var RangeType_name = map[int32]string{
	0: "SHORT",
	1: "MIDDLE",
	2: "LONG",
}

var RangeType_value = map[string]int32{
	"SHORT":  0,
	"MIDDLE": 1,
	"LONG":   2,
}

func (x RangeType) String() string {
	return proto.EnumName(RangeType_name, int32(x))
}

func (RangeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{3}
}

type SignEq int32

const (
	SignEq_EQUAL   SignEq = 0
	SignEq_GREATER SignEq = 1
	SignEq_LESSER  SignEq = 2
)

var SignEq_name = map[int32]string{
	0: "EQUAL",
	1: "GREATER",
	2: "LESSER",
}

var SignEq_value = map[string]int32{
	"EQUAL":   0,
	"GREATER": 1,
	"LESSER":  2,
}

func (x SignEq) String() string {
	return proto.EnumName(SignEq_name, int32(x))
}

func (SignEq) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{4}
}

// GameDataSet
//      the whole instance game-duel set
type GameDataSet struct {
	//  room_key, see also RoomService/message.proto : Room
	RoomKey string `protobuf:"bytes,1,opt,name=room_key,json=roomKey,proto3" json:"room_key,omitempty"`
	//  host_id, the user-id of host player
	HostId string `protobuf:"bytes,2,opt,name=host_id,json=hostId,proto3" json:"host_id,omitempty"`
	//  dueler_id, the user-id of dueler player
	DuelId string `protobuf:"bytes,3,opt,name=duel_id,json=duelId,proto3" json:"duel_id,omitempty"`
	//  game_turn, the number of turns in game
	GameTurn int32 `protobuf:"varint,4,opt,name=game_turn,json=gameTurn,proto3" json:"game_turn,omitempty"`
	//  host_card_deck, the character card
	HostCardDeck []*CharCardSet `protobuf:"bytes,5,rep,name=host_card_deck,json=hostCardDeck,proto3" json:"host_card_deck,omitempty"`
	//  duel_card_deck, the character card
	DuelCardDeck []*CharCardSet `protobuf:"bytes,6,rep,name=duel_card_deck,json=duelCardDeck,proto3" json:"duel_card_deck,omitempty"`
	//  host_curr_card, current on-used char-card
	HostCurrCardKey int32 `protobuf:"varint,7,opt,name=host_curr_card_key,json=hostCurrCardKey,proto3" json:"host_curr_card_key,omitempty"`
	//  dueler_curr_card, current on-used char-card
	DuelCurrCardKey int32 `protobuf:"varint,8,opt,name=duel_curr_card_key,json=duelCurrCardKey,proto3" json:"duel_curr_card_key,omitempty"`
	//  host_event_card_deck
	HostEventCardDeck []*EventCard `protobuf:"bytes,9,rep,name=host_event_card_deck,json=hostEventCardDeck,proto3" json:"host_event_card_deck,omitempty"`
	//  duel_event_card_deck
	DuelEventCardDeck []*EventCard `protobuf:"bytes,10,rep,name=duel_event_card_deck,json=duelEventCardDeck,proto3" json:"duel_event_card_deck,omitempty"`
	//  curr_range
	Range RangeType `protobuf:"varint,11,opt,name=range,proto3,enum=ULZProto.RangeType" json:"range,omitempty"`
	// nvn
	Nvn       int32      `protobuf:"varint,12,opt,name=nvn,proto3" json:"nvn,omitempty"`
	PhaseAb   PlayerSide `protobuf:"varint,13,opt,name=phase_ab,json=phaseAb,proto3,enum=ULZProto.PlayerSide" json:"phase_ab,omitempty"`
	CurrPhase PlayerSide `protobuf:"varint,14,opt,name=curr_phase,json=currPhase,proto3,enum=ULZProto.PlayerSide" json:"curr_phase,omitempty"`
	//  event_phase, the event hook phase
	EventPhase EventHookPhase `protobuf:"varint,15,opt,name=event_phase,json=eventPhase,proto3,enum=ULZProto.EventHookPhase" json:"event_phase,omitempty"`
	//  hook_type, the event hook type
	HookType EventHookType `protobuf:"varint,16,opt,name=hook_type,json=hookType,proto3,enum=ULZProto.EventHookType" json:"hook_type,omitempty"`
	//  flaging
	IsHostReady bool `protobuf:"varint,17,opt,name=is_host_ready,json=isHostReady,proto3" json:"is_host_ready,omitempty"`
	IsDuelReady bool `protobuf:"varint,18,opt,name=is_duel_ready,json=isDuelReady,proto3" json:"is_duel_ready,omitempty"`
	//  Effect-result
	EffectCounter        []*EffectResult `protobuf:"bytes,19,rep,name=effect_counter,json=effectCounter,proto3" json:"effect_counter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GameDataSet) Reset()         { *m = GameDataSet{} }
func (m *GameDataSet) String() string { return proto.CompactTextString(m) }
func (*GameDataSet) ProtoMessage()    {}
func (*GameDataSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{0}
}

func (m *GameDataSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GameDataSet.Unmarshal(m, b)
}
func (m *GameDataSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GameDataSet.Marshal(b, m, deterministic)
}
func (m *GameDataSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GameDataSet.Merge(m, src)
}
func (m *GameDataSet) XXX_Size() int {
	return xxx_messageInfo_GameDataSet.Size(m)
}
func (m *GameDataSet) XXX_DiscardUnknown() {
	xxx_messageInfo_GameDataSet.DiscardUnknown(m)
}

var xxx_messageInfo_GameDataSet proto.InternalMessageInfo

func (m *GameDataSet) GetRoomKey() string {
	if m != nil {
		return m.RoomKey
	}
	return ""
}

func (m *GameDataSet) GetHostId() string {
	if m != nil {
		return m.HostId
	}
	return ""
}

func (m *GameDataSet) GetDuelId() string {
	if m != nil {
		return m.DuelId
	}
	return ""
}

func (m *GameDataSet) GetGameTurn() int32 {
	if m != nil {
		return m.GameTurn
	}
	return 0
}

func (m *GameDataSet) GetHostCardDeck() []*CharCardSet {
	if m != nil {
		return m.HostCardDeck
	}
	return nil
}

func (m *GameDataSet) GetDuelCardDeck() []*CharCardSet {
	if m != nil {
		return m.DuelCardDeck
	}
	return nil
}

func (m *GameDataSet) GetHostCurrCardKey() int32 {
	if m != nil {
		return m.HostCurrCardKey
	}
	return 0
}

func (m *GameDataSet) GetDuelCurrCardKey() int32 {
	if m != nil {
		return m.DuelCurrCardKey
	}
	return 0
}

func (m *GameDataSet) GetHostEventCardDeck() []*EventCard {
	if m != nil {
		return m.HostEventCardDeck
	}
	return nil
}

func (m *GameDataSet) GetDuelEventCardDeck() []*EventCard {
	if m != nil {
		return m.DuelEventCardDeck
	}
	return nil
}

func (m *GameDataSet) GetRange() RangeType {
	if m != nil {
		return m.Range
	}
	return RangeType_SHORT
}

func (m *GameDataSet) GetNvn() int32 {
	if m != nil {
		return m.Nvn
	}
	return 0
}

func (m *GameDataSet) GetPhaseAb() PlayerSide {
	if m != nil {
		return m.PhaseAb
	}
	return PlayerSide_HOST
}

func (m *GameDataSet) GetCurrPhase() PlayerSide {
	if m != nil {
		return m.CurrPhase
	}
	return PlayerSide_HOST
}

func (m *GameDataSet) GetEventPhase() EventHookPhase {
	if m != nil {
		return m.EventPhase
	}
	return EventHookPhase_gameset_start
}

func (m *GameDataSet) GetHookType() EventHookType {
	if m != nil {
		return m.HookType
	}
	return EventHookType_Instant
}

func (m *GameDataSet) GetIsHostReady() bool {
	if m != nil {
		return m.IsHostReady
	}
	return false
}

func (m *GameDataSet) GetIsDuelReady() bool {
	if m != nil {
		return m.IsDuelReady
	}
	return false
}

func (m *GameDataSet) GetEffectCounter() []*EffectResult {
	if m != nil {
		return m.EffectCounter
	}
	return nil
}

// CharCardSet
//
type CharCardSet struct {
	CharId               int32           `protobuf:"varint,1,opt,name=char_id,json=charId,proto3" json:"char_id,omitempty"`
	CardId               int32           `protobuf:"varint,2,opt,name=card_id,json=cardId,proto3" json:"card_id,omitempty"`
	HpInst               int32           `protobuf:"varint,3,opt,name=hp_inst,json=hpInst,proto3" json:"hp_inst,omitempty"`
	ApInst               int32           `protobuf:"varint,4,opt,name=ap_inst,json=apInst,proto3" json:"ap_inst,omitempty"`
	DpInst               int32           `protobuf:"varint,5,opt,name=dp_inst,json=dpInst,proto3" json:"dp_inst,omitempty"`
	Level                int32           `protobuf:"varint,6,opt,name=level,proto3" json:"level,omitempty"`
	StatusInst           []*StatusSet    `protobuf:"bytes,7,rep,name=status_inst,json=statusInst,proto3" json:"status_inst,omitempty"`
	EquSet               *CharCardEquSet `protobuf:"bytes,8,opt,name=equ_set,json=equSet,proto3" json:"equ_set,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CharCardSet) Reset()         { *m = CharCardSet{} }
func (m *CharCardSet) String() string { return proto.CompactTextString(m) }
func (*CharCardSet) ProtoMessage()    {}
func (*CharCardSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{1}
}

func (m *CharCardSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharCardSet.Unmarshal(m, b)
}
func (m *CharCardSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharCardSet.Marshal(b, m, deterministic)
}
func (m *CharCardSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharCardSet.Merge(m, src)
}
func (m *CharCardSet) XXX_Size() int {
	return xxx_messageInfo_CharCardSet.Size(m)
}
func (m *CharCardSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CharCardSet.DiscardUnknown(m)
}

var xxx_messageInfo_CharCardSet proto.InternalMessageInfo

func (m *CharCardSet) GetCharId() int32 {
	if m != nil {
		return m.CharId
	}
	return 0
}

func (m *CharCardSet) GetCardId() int32 {
	if m != nil {
		return m.CardId
	}
	return 0
}

func (m *CharCardSet) GetHpInst() int32 {
	if m != nil {
		return m.HpInst
	}
	return 0
}

func (m *CharCardSet) GetApInst() int32 {
	if m != nil {
		return m.ApInst
	}
	return 0
}

func (m *CharCardSet) GetDpInst() int32 {
	if m != nil {
		return m.DpInst
	}
	return 0
}

func (m *CharCardSet) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *CharCardSet) GetStatusInst() []*StatusSet {
	if m != nil {
		return m.StatusInst
	}
	return nil
}

func (m *CharCardSet) GetEquSet() *CharCardEquSet {
	if m != nil {
		return m.EquSet
	}
	return nil
}

type CharCardEquSet struct {
	EquId                int32    `protobuf:"varint,1,opt,name=equ_id,json=equId,proto3" json:"equ_id,omitempty"`
	Hp                   int32    `protobuf:"varint,2,opt,name=hp,proto3" json:"hp,omitempty"`
	Ap                   int32    `protobuf:"varint,3,opt,name=ap,proto3" json:"ap,omitempty"`
	Dp                   int32    `protobuf:"varint,4,opt,name=dp,proto3" json:"dp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CharCardEquSet) Reset()         { *m = CharCardEquSet{} }
func (m *CharCardEquSet) String() string { return proto.CompactTextString(m) }
func (*CharCardEquSet) ProtoMessage()    {}
func (*CharCardEquSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{2}
}

func (m *CharCardEquSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CharCardEquSet.Unmarshal(m, b)
}
func (m *CharCardEquSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CharCardEquSet.Marshal(b, m, deterministic)
}
func (m *CharCardEquSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CharCardEquSet.Merge(m, src)
}
func (m *CharCardEquSet) XXX_Size() int {
	return xxx_messageInfo_CharCardEquSet.Size(m)
}
func (m *CharCardEquSet) XXX_DiscardUnknown() {
	xxx_messageInfo_CharCardEquSet.DiscardUnknown(m)
}

var xxx_messageInfo_CharCardEquSet proto.InternalMessageInfo

func (m *CharCardEquSet) GetEquId() int32 {
	if m != nil {
		return m.EquId
	}
	return 0
}

func (m *CharCardEquSet) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *CharCardEquSet) GetAp() int32 {
	if m != nil {
		return m.Ap
	}
	return 0
}

func (m *CharCardEquSet) GetDp() int32 {
	if m != nil {
		return m.Dp
	}
	return 0
}

type EventCard struct {
	Id                   int32         `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UpOption             EventCardType `protobuf:"varint,2,opt,name=up_option,json=upOption,proto3,enum=ULZProto.EventCardType" json:"up_option,omitempty"`
	UpVal                int32         `protobuf:"varint,3,opt,name=up_val,json=upVal,proto3" json:"up_val,omitempty"`
	DownOption           EventCardType `protobuf:"varint,4,opt,name=down_option,json=downOption,proto3,enum=ULZProto.EventCardType" json:"down_option,omitempty"`
	DownVal              int32         `protobuf:"varint,5,opt,name=down_val,json=downVal,proto3" json:"down_val,omitempty"`
	Position             EventCardPos  `protobuf:"varint,6,opt,name=position,proto3,enum=ULZProto.EventCardPos" json:"position,omitempty"`
	IsInvert             bool          `protobuf:"varint,7,opt,name=is_invert,json=isInvert,proto3" json:"is_invert,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *EventCard) Reset()         { *m = EventCard{} }
func (m *EventCard) String() string { return proto.CompactTextString(m) }
func (*EventCard) ProtoMessage()    {}
func (*EventCard) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{3}
}

func (m *EventCard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventCard.Unmarshal(m, b)
}
func (m *EventCard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventCard.Marshal(b, m, deterministic)
}
func (m *EventCard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCard.Merge(m, src)
}
func (m *EventCard) XXX_Size() int {
	return xxx_messageInfo_EventCard.Size(m)
}
func (m *EventCard) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCard.DiscardUnknown(m)
}

var xxx_messageInfo_EventCard proto.InternalMessageInfo

func (m *EventCard) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EventCard) GetUpOption() EventCardType {
	if m != nil {
		return m.UpOption
	}
	return EventCardType_NULL
}

func (m *EventCard) GetUpVal() int32 {
	if m != nil {
		return m.UpVal
	}
	return 0
}

func (m *EventCard) GetDownOption() EventCardType {
	if m != nil {
		return m.DownOption
	}
	return EventCardType_NULL
}

func (m *EventCard) GetDownVal() int32 {
	if m != nil {
		return m.DownVal
	}
	return 0
}

func (m *EventCard) GetPosition() EventCardPos {
	if m != nil {
		return m.Position
	}
	return EventCardPos_BLOCK
}

func (m *EventCard) GetIsInvert() bool {
	if m != nil {
		return m.IsInvert
	}
	return false
}

type SkillSet struct {
	Id                   int32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FeatNo               int32            `protobuf:"varint,2,opt,name=feat_no,json=featNo,proto3" json:"feat_no,omitempty"`
	Pow                  int32            `protobuf:"varint,3,opt,name=pow,proto3" json:"pow,omitempty"`
	CondString           string           `protobuf:"bytes,4,opt,name=cond_string,json=condString,proto3" json:"cond_string,omitempty"`
	CondCard             []*SkillCardCond `protobuf:"bytes,5,rep,name=cond_card,json=condCard,proto3" json:"cond_card,omitempty"`
	CondRange            RangeType        `protobuf:"varint,6,opt,name=cond_range,json=condRange,proto3,enum=ULZProto.RangeType" json:"cond_range,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SkillSet) Reset()         { *m = SkillSet{} }
func (m *SkillSet) String() string { return proto.CompactTextString(m) }
func (*SkillSet) ProtoMessage()    {}
func (*SkillSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{4}
}

func (m *SkillSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SkillSet.Unmarshal(m, b)
}
func (m *SkillSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SkillSet.Marshal(b, m, deterministic)
}
func (m *SkillSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkillSet.Merge(m, src)
}
func (m *SkillSet) XXX_Size() int {
	return xxx_messageInfo_SkillSet.Size(m)
}
func (m *SkillSet) XXX_DiscardUnknown() {
	xxx_messageInfo_SkillSet.DiscardUnknown(m)
}

var xxx_messageInfo_SkillSet proto.InternalMessageInfo

func (m *SkillSet) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SkillSet) GetFeatNo() int32 {
	if m != nil {
		return m.FeatNo
	}
	return 0
}

func (m *SkillSet) GetPow() int32 {
	if m != nil {
		return m.Pow
	}
	return 0
}

func (m *SkillSet) GetCondString() string {
	if m != nil {
		return m.CondString
	}
	return ""
}

func (m *SkillSet) GetCondCard() []*SkillCardCond {
	if m != nil {
		return m.CondCard
	}
	return nil
}

func (m *SkillSet) GetCondRange() RangeType {
	if m != nil {
		return m.CondRange
	}
	return RangeType_SHORT
}

type SkillCardCond struct {
	Type                 EventCardType `protobuf:"varint,1,opt,name=type,proto3,enum=ULZProto.EventCardType" json:"type,omitempty"`
	Val                  int32         `protobuf:"varint,2,opt,name=val,proto3" json:"val,omitempty"`
	Compare              SignEq        `protobuf:"varint,3,opt,name=compare,proto3,enum=ULZProto.SignEq" json:"compare,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SkillCardCond) Reset()         { *m = SkillCardCond{} }
func (m *SkillCardCond) String() string { return proto.CompactTextString(m) }
func (*SkillCardCond) ProtoMessage()    {}
func (*SkillCardCond) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{5}
}

func (m *SkillCardCond) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SkillCardCond.Unmarshal(m, b)
}
func (m *SkillCardCond) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SkillCardCond.Marshal(b, m, deterministic)
}
func (m *SkillCardCond) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SkillCardCond.Merge(m, src)
}
func (m *SkillCardCond) XXX_Size() int {
	return xxx_messageInfo_SkillCardCond.Size(m)
}
func (m *SkillCardCond) XXX_DiscardUnknown() {
	xxx_messageInfo_SkillCardCond.DiscardUnknown(m)
}

var xxx_messageInfo_SkillCardCond proto.InternalMessageInfo

func (m *SkillCardCond) GetType() EventCardType {
	if m != nil {
		return m.Type
	}
	return EventCardType_NULL
}

func (m *SkillCardCond) GetVal() int32 {
	if m != nil {
		return m.Val
	}
	return 0
}

func (m *SkillCardCond) GetCompare() SignEq {
	if m != nil {
		return m.Compare
	}
	return SignEq_EQUAL
}

type StatusSet struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StatusId             string   `protobuf:"bytes,2,opt,name=status_id,json=statusId,proto3" json:"status_id,omitempty"`
	RemainCd             int32    `protobuf:"varint,3,opt,name=remain_cd,json=remainCd,proto3" json:"remain_cd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusSet) Reset()         { *m = StatusSet{} }
func (m *StatusSet) String() string { return proto.CompactTextString(m) }
func (*StatusSet) ProtoMessage()    {}
func (*StatusSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{6}
}

func (m *StatusSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusSet.Unmarshal(m, b)
}
func (m *StatusSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusSet.Marshal(b, m, deterministic)
}
func (m *StatusSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusSet.Merge(m, src)
}
func (m *StatusSet) XXX_Size() int {
	return xxx_messageInfo_StatusSet.Size(m)
}
func (m *StatusSet) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusSet.DiscardUnknown(m)
}

var xxx_messageInfo_StatusSet proto.InternalMessageInfo

func (m *StatusSet) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StatusSet) GetStatusId() string {
	if m != nil {
		return m.StatusId
	}
	return ""
}

func (m *StatusSet) GetRemainCd() int32 {
	if m != nil {
		return m.RemainCd
	}
	return 0
}

type Status_Effect struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FeatNo               int32    `protobuf:"varint,2,opt,name=feat_no,json=featNo,proto3" json:"feat_no,omitempty"`
	Pow                  int32    `protobuf:"varint,3,opt,name=pow,proto3" json:"pow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status_Effect) Reset()         { *m = Status_Effect{} }
func (m *Status_Effect) String() string { return proto.CompactTextString(m) }
func (*Status_Effect) ProtoMessage()    {}
func (*Status_Effect) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{7}
}

func (m *Status_Effect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status_Effect.Unmarshal(m, b)
}
func (m *Status_Effect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status_Effect.Marshal(b, m, deterministic)
}
func (m *Status_Effect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status_Effect.Merge(m, src)
}
func (m *Status_Effect) XXX_Size() int {
	return xxx_messageInfo_Status_Effect.Size(m)
}
func (m *Status_Effect) XXX_DiscardUnknown() {
	xxx_messageInfo_Status_Effect.DiscardUnknown(m)
}

var xxx_messageInfo_Status_Effect proto.InternalMessageInfo

func (m *Status_Effect) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Status_Effect) GetFeatNo() int32 {
	if m != nil {
		return m.FeatNo
	}
	return 0
}

func (m *Status_Effect) GetPow() int32 {
	if m != nil {
		return m.Pow
	}
	return 0
}

type EffectResult struct {
	// targer
	TarCard int32      `protobuf:"varint,1,opt,name=tar_card,json=tarCard,proto3" json:"tar_card,omitempty"`
	TarSide PlayerSide `protobuf:"varint,2,opt,name=tar_side,json=tarSide,proto3,enum=ULZProto.PlayerSide" json:"tar_side,omitempty"`
	// trigger-time
	EventPhase EventHookPhase `protobuf:"varint,3,opt,name=event_phase,json=eventPhase,proto3,enum=ULZProto.EventHookPhase" json:"event_phase,omitempty"`
	HookType   EventHookType  `protobuf:"varint,4,opt,name=hook_type,json=hookType,proto3,enum=ULZProto.EventHookType" json:"hook_type,omitempty"`
	SubCount   int32          `protobuf:"varint,5,opt,name=sub_count,json=subCount,proto3" json:"sub_count,omitempty"`
	// effect
	Hp            int32  `protobuf:"varint,6,opt,name=hp,proto3" json:"hp,omitempty"`
	Ap            int32  `protobuf:"varint,7,opt,name=ap,proto3" json:"ap,omitempty"`
	Dp            int32  `protobuf:"varint,8,opt,name=dp,proto3" json:"dp,omitempty"`
	DisableMove   bool   `protobuf:"varint,9,opt,name=disable_move,json=disableMove,proto3" json:"disable_move,omitempty"`
	DisableAtk    bool   `protobuf:"varint,10,opt,name=disable_atk,json=disableAtk,proto3" json:"disable_atk,omitempty"`
	DisableDef    bool   `protobuf:"varint,11,opt,name=disable_def,json=disableDef,proto3" json:"disable_def,omitempty"`
	DisableDraw   bool   `protobuf:"varint,12,opt,name=disable_draw,json=disableDraw,proto3" json:"disable_draw,omitempty"`
	DisableChange bool   `protobuf:"varint,13,opt,name=disable_change,json=disableChange,proto3" json:"disable_change,omitempty"`
	BindingFunc   string `protobuf:"bytes,14,opt,name=binding_func,json=bindingFunc,proto3" json:"binding_func,omitempty"`
	//
	RemainCd             int32    `protobuf:"varint,15,opt,name=remain_cd,json=remainCd,proto3" json:"remain_cd,omitempty"`
	AssignFrom           string   `protobuf:"bytes,16,opt,name=assign_from,json=assignFrom,proto3" json:"assign_from,omitempty"`
	SkillId              int32    `protobuf:"varint,17,opt,name=skill_id,json=skillId,proto3" json:"skill_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EffectResult) Reset()         { *m = EffectResult{} }
func (m *EffectResult) String() string { return proto.CompactTextString(m) }
func (*EffectResult) ProtoMessage()    {}
func (*EffectResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cccadeca1f8aa7b, []int{8}
}

func (m *EffectResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EffectResult.Unmarshal(m, b)
}
func (m *EffectResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EffectResult.Marshal(b, m, deterministic)
}
func (m *EffectResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EffectResult.Merge(m, src)
}
func (m *EffectResult) XXX_Size() int {
	return xxx_messageInfo_EffectResult.Size(m)
}
func (m *EffectResult) XXX_DiscardUnknown() {
	xxx_messageInfo_EffectResult.DiscardUnknown(m)
}

var xxx_messageInfo_EffectResult proto.InternalMessageInfo

func (m *EffectResult) GetTarCard() int32 {
	if m != nil {
		return m.TarCard
	}
	return 0
}

func (m *EffectResult) GetTarSide() PlayerSide {
	if m != nil {
		return m.TarSide
	}
	return PlayerSide_HOST
}

func (m *EffectResult) GetEventPhase() EventHookPhase {
	if m != nil {
		return m.EventPhase
	}
	return EventHookPhase_gameset_start
}

func (m *EffectResult) GetHookType() EventHookType {
	if m != nil {
		return m.HookType
	}
	return EventHookType_Instant
}

func (m *EffectResult) GetSubCount() int32 {
	if m != nil {
		return m.SubCount
	}
	return 0
}

func (m *EffectResult) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *EffectResult) GetAp() int32 {
	if m != nil {
		return m.Ap
	}
	return 0
}

func (m *EffectResult) GetDp() int32 {
	if m != nil {
		return m.Dp
	}
	return 0
}

func (m *EffectResult) GetDisableMove() bool {
	if m != nil {
		return m.DisableMove
	}
	return false
}

func (m *EffectResult) GetDisableAtk() bool {
	if m != nil {
		return m.DisableAtk
	}
	return false
}

func (m *EffectResult) GetDisableDef() bool {
	if m != nil {
		return m.DisableDef
	}
	return false
}

func (m *EffectResult) GetDisableDraw() bool {
	if m != nil {
		return m.DisableDraw
	}
	return false
}

func (m *EffectResult) GetDisableChange() bool {
	if m != nil {
		return m.DisableChange
	}
	return false
}

func (m *EffectResult) GetBindingFunc() string {
	if m != nil {
		return m.BindingFunc
	}
	return ""
}

func (m *EffectResult) GetRemainCd() int32 {
	if m != nil {
		return m.RemainCd
	}
	return 0
}

func (m *EffectResult) GetAssignFrom() string {
	if m != nil {
		return m.AssignFrom
	}
	return ""
}

func (m *EffectResult) GetSkillId() int32 {
	if m != nil {
		return m.SkillId
	}
	return 0
}

func init() {
	proto.RegisterEnum("ULZProto.PlayerSide", PlayerSide_name, PlayerSide_value)
	proto.RegisterEnum("ULZProto.EventCardType", EventCardType_name, EventCardType_value)
	proto.RegisterEnum("ULZProto.EventCardPos", EventCardPos_name, EventCardPos_value)
	proto.RegisterEnum("ULZProto.RangeType", RangeType_name, RangeType_value)
	proto.RegisterEnum("ULZProto.SignEq", SignEq_name, SignEq_value)
	proto.RegisterType((*GameDataSet)(nil), "ULZProto.GameDataSet")
	proto.RegisterType((*CharCardSet)(nil), "ULZProto.CharCardSet")
	proto.RegisterType((*CharCardEquSet)(nil), "ULZProto.CharCardEquSet")
	proto.RegisterType((*EventCard)(nil), "ULZProto.EventCard")
	proto.RegisterType((*SkillSet)(nil), "ULZProto.SkillSet")
	proto.RegisterType((*SkillCardCond)(nil), "ULZProto.SkillCardCond")
	proto.RegisterType((*StatusSet)(nil), "ULZProto.StatusSet")
	proto.RegisterType((*Status_Effect)(nil), "ULZProto.Status_Effect")
	proto.RegisterType((*EffectResult)(nil), "ULZProto.EffectResult")
}

func init() { proto.RegisterFile("Data.proto", fileDescriptor_4cccadeca1f8aa7b) }

var fileDescriptor_4cccadeca1f8aa7b = []byte{
	// 1337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0xdb, 0x6e, 0xdb, 0x46,
	0x13, 0xfe, 0x75, 0xe2, 0x61, 0x64, 0x29, 0x32, 0xe3, 0x24, 0xcc, 0x9f, 0x8b, 0xdf, 0xbf, 0x80,
	0x02, 0xae, 0x13, 0xb8, 0xa8, 0x93, 0x8b, 0x16, 0x45, 0x51, 0x28, 0x12, 0x63, 0xab, 0x91, 0x2d,
	0x97, 0x92, 0x52, 0x34, 0x37, 0xc4, 0x4a, 0xbb, 0xb2, 0x08, 0x49, 0x5c, 0x9a, 0x07, 0x19, 0xee,
	0x1b, 0xf4, 0xa6, 0x0f, 0xd2, 0xd7, 0xe8, 0x0b, 0xf4, 0x59, 0xfa, 0x04, 0xc5, 0xcc, 0x52, 0xc7,
	0x24, 0x46, 0xd1, 0x5e, 0x89, 0x33, 0xf3, 0xcd, 0xb7, 0x33, 0x3b, 0x87, 0x15, 0x40, 0x8b, 0x25,
	0xec, 0x24, 0x8c, 0x64, 0x22, 0x2d, 0x63, 0xd0, 0x79, 0x7f, 0x85, 0x5f, 0xff, 0x3d, 0x70, 0x16,
	0x22, 0x48, 0xce, 0xa5, 0x9c, 0x5e, 0x4d, 0x58, 0x2c, 0x94, 0xbd, 0xfe, 0xa7, 0x06, 0xe5, 0x33,
	0x36, 0x17, 0xe8, 0xd2, 0x13, 0x89, 0xf5, 0x14, 0x8c, 0x48, 0xca, 0xb9, 0x37, 0x15, 0x77, 0x76,
	0xee, 0x30, 0x77, 0x64, 0xba, 0x3a, 0xca, 0x6f, 0xc5, 0x9d, 0xf5, 0x04, 0xf4, 0x89, 0x8c, 0x13,
	0xcf, 0xe7, 0x76, 0x9e, 0x2c, 0x1a, 0x8a, 0x6d, 0x8e, 0x06, 0x9e, 0x8a, 0x19, 0x1a, 0x0a, 0xca,
	0x80, 0x62, 0x9b, 0x5b, 0xcf, 0xc0, 0xbc, 0x66, 0x73, 0xe1, 0x25, 0x69, 0x14, 0xd8, 0xc5, 0xc3,
	0xdc, 0x51, 0xc9, 0x35, 0x50, 0xd1, 0x4f, 0xa3, 0xc0, 0xfa, 0x06, 0xaa, 0x44, 0x37, 0x62, 0x11,
	0xf7, 0xb8, 0x18, 0x4d, 0xed, 0xd2, 0x61, 0xe1, 0xa8, 0x7c, 0xfa, 0xe8, 0x64, 0x19, 0xf2, 0x49,
	0x73, 0xc2, 0xa2, 0x26, 0x8b, 0x78, 0x4f, 0x24, 0xee, 0x1e, 0x82, 0x51, 0x68, 0x89, 0xd1, 0x14,
	0x9d, 0xe9, 0xc8, 0xb5, 0xb3, 0x76, 0xaf, 0x33, 0x82, 0x57, 0xce, 0xcf, 0xc1, 0x52, 0x27, 0xa7,
	0x51, 0xa4, 0x18, 0x30, 0x5b, 0x9d, 0xe2, 0x7b, 0x40, 0xc7, 0xa4, 0x11, 0xb9, 0x62, 0xd6, 0xcf,
	0xc1, 0x52, 0x27, 0x6d, 0x81, 0x0d, 0x05, 0x26, 0xda, 0x0d, 0x70, 0x0b, 0x0e, 0x88, 0x59, 0xe0,
	0x55, 0x6f, 0x04, 0x67, 0x52, 0x70, 0x0f, 0xd7, 0xc1, 0x51, 0x2d, 0xd0, 0xcb, 0xdd, 0x47, 0x87,
	0x95, 0x48, 0xf1, 0xb5, 0xe0, 0x80, 0x8e, 0xdc, 0x65, 0x81, 0x7b, 0x58, 0xd0, 0x61, 0x9b, 0xe5,
	0x73, 0x28, 0x45, 0x2c, 0xb8, 0x16, 0x76, 0xf9, 0x30, 0x77, 0x54, 0xdd, 0x74, 0x73, 0x51, 0xdd,
	0xbf, 0x0b, 0x85, 0xab, 0x10, 0x56, 0x0d, 0x0a, 0xc1, 0x22, 0xb0, 0xf7, 0x28, 0x29, 0xfc, 0xb4,
	0xbe, 0x00, 0x23, 0xc4, 0x2e, 0xf1, 0xd8, 0xd0, 0xae, 0x90, 0xff, 0xc1, 0xda, 0xff, 0x6a, 0xc6,
	0xee, 0x44, 0xd4, 0xf3, 0xb9, 0x70, 0x75, 0x42, 0x35, 0x86, 0xd6, 0x4b, 0x00, 0xba, 0x21, 0x92,
	0xed, 0xea, 0x3d, 0x2e, 0x26, 0xe2, 0xa8, 0x05, 0xad, 0xaf, 0xa1, 0xac, 0x72, 0x54, 0x5e, 0x0f,
	0xc8, 0xcb, 0xde, 0xc9, 0x6f, 0xd5, 0xb1, 0x2e, 0x10, 0x58, 0xb9, 0xbe, 0x02, 0x73, 0x22, 0xe5,
	0xd4, 0x4b, 0xee, 0x42, 0x61, 0xd7, 0xc8, 0xf1, 0xc9, 0x47, 0x1c, 0x29, 0x4b, 0x63, 0x92, 0x7d,
	0x59, 0x75, 0xa8, 0xf8, 0xb1, 0x47, 0x25, 0x8a, 0x04, 0xe3, 0x77, 0xf6, 0xfe, 0x61, 0xee, 0xc8,
	0x70, 0xcb, 0x7e, 0x7c, 0x2e, 0xe3, 0xc4, 0x45, 0x55, 0x86, 0xa1, 0x02, 0x28, 0x8c, 0xb5, 0xc4,
	0xb4, 0x52, 0x31, 0x53, 0x98, 0x6f, 0xa1, 0x2a, 0xc6, 0x63, 0x31, 0x4a, 0xbc, 0x91, 0x4c, 0x83,
	0x44, 0x44, 0xf6, 0x43, 0xaa, 0xcd, 0xe3, 0x8d, 0x10, 0xc8, 0xee, 0x8a, 0x38, 0x9d, 0x25, 0x6e,
	0x45, 0xa1, 0x9b, 0x0a, 0x5c, 0xff, 0x25, 0x0f, 0xe5, 0x8d, 0xf6, 0xc4, 0x01, 0x1a, 0x4d, 0x58,
	0x84, 0x03, 0x94, 0xa3, 0x1a, 0x68, 0x28, 0xaa, 0xc9, 0xa2, 0xf2, 0x67, 0x23, 0x87, 0x06, 0x16,
	0x71, 0x65, 0x98, 0x84, 0x9e, 0x1f, 0xc4, 0x09, 0x8d, 0x5c, 0xc9, 0xd5, 0x26, 0x61, 0x3b, 0x88,
	0x89, 0x8a, 0x65, 0x06, 0x35, 0x70, 0x1a, 0x5b, 0x19, 0x78, 0x66, 0x28, 0x29, 0x03, 0x57, 0x86,
	0x03, 0x28, 0xcd, 0xc4, 0x42, 0xcc, 0x6c, 0x8d, 0xd4, 0x4a, 0xb0, 0x5e, 0x41, 0x39, 0x4e, 0x58,
	0x92, 0xc6, 0xca, 0x45, 0xdf, 0x6d, 0xbd, 0x1e, 0x19, 0x71, 0xb6, 0x40, 0xe1, 0x88, 0xeb, 0x4b,
	0xd0, 0xc5, 0x4d, 0xea, 0xc5, 0x22, 0xa1, 0x09, 0x29, 0x6f, 0x16, 0x73, 0x99, 0xb0, 0x73, 0x93,
	0xa2, 0x9b, 0x26, 0xe8, 0xb7, 0xfe, 0x23, 0x54, 0xb7, 0x2d, 0xd6, 0x23, 0x40, 0xdb, 0xfa, 0x32,
	0x4a, 0xe2, 0x26, 0x6d, 0x73, 0xab, 0x0a, 0xf9, 0x49, 0x98, 0x5d, 0x43, 0x7e, 0x12, 0xa2, 0xcc,
	0xc2, 0x2c, 0xfb, 0x3c, 0x23, 0x99, 0x87, 0x59, 0xd2, 0x79, 0x1e, 0xd6, 0x7f, 0xcd, 0x83, 0xb9,
	0x9a, 0x08, 0xb4, 0xae, 0x08, 0xf3, 0x3e, 0xc7, 0xfe, 0x49, 0x43, 0x4f, 0x86, 0x89, 0x2f, 0x03,
	0x22, 0xfd, 0xb0, 0x7f, 0xd0, 0x4f, 0xf5, 0x4f, 0x1a, 0x76, 0x09, 0x88, 0xa1, 0xa5, 0xa1, 0xb7,
	0x60, 0xb3, 0xec, 0xdc, 0x52, 0x1a, 0xbe, 0x63, 0x33, 0xeb, 0x2b, 0x28, 0x73, 0x79, 0x1b, 0x2c,
	0xe9, 0x8a, 0xf7, 0xd3, 0x01, 0x62, 0x33, 0xc2, 0xa7, 0x60, 0x90, 0x27, 0x52, 0xaa, 0xb2, 0xe8,
	0x28, 0x23, 0xe9, 0x29, 0x18, 0xa1, 0x8c, 0x7d, 0x62, 0xd4, 0x88, 0xf1, 0xf1, 0x47, 0x18, 0xaf,
	0x64, 0xec, 0xae, 0x70, 0xb8, 0x70, 0x7d, 0xac, 0xd8, 0x42, 0x44, 0x09, 0x2d, 0x34, 0xc3, 0x35,
	0xfc, 0xb8, 0x4d, 0x72, 0xfd, 0x8f, 0x1c, 0x18, 0xbd, 0xa9, 0x3f, 0x9b, 0xe1, 0x25, 0xef, 0xde,
	0xc7, 0x13, 0xd0, 0xc7, 0x82, 0x25, 0x5e, 0x20, 0x97, 0x9d, 0x86, 0xe2, 0xa5, 0xc4, 0xdd, 0x10,
	0xca, 0xdb, 0x2c, 0x5f, 0xfc, 0xb4, 0xfe, 0x07, 0xe5, 0x91, 0x0c, 0xb8, 0x17, 0x27, 0x91, 0x1f,
	0x5c, 0x53, 0xb6, 0xa6, 0x0b, 0xa8, 0xea, 0x91, 0x06, 0xef, 0x96, 0x00, 0xd8, 0xab, 0xd9, 0x52,
	0xdf, 0xb8, 0x0c, 0x0a, 0x01, 0x43, 0x6f, 0xca, 0x80, 0xbb, 0x06, 0x22, 0xa9, 0x42, 0xa7, 0x40,
	0x1c, 0x9e, 0x5a, 0x5a, 0xda, 0xa7, 0x97, 0x16, 0x91, 0x93, 0x58, 0xff, 0x19, 0x2a, 0x5b, 0x74,
	0xd6, 0x73, 0x28, 0xd2, 0x46, 0xc8, 0xdd, 0x5f, 0x02, 0x02, 0x61, 0x6a, 0x78, 0xef, 0x2a, 0x5f,
	0xfc, 0xb4, 0x8e, 0x41, 0x1f, 0xc9, 0x79, 0xc8, 0x22, 0x41, 0x09, 0x57, 0x4f, 0x6b, 0x1b, 0x71,
	0xfb, 0xd7, 0x81, 0x73, 0xe3, 0x2e, 0x01, 0xf5, 0x01, 0x98, 0xab, 0x21, 0xf8, 0xe0, 0x3a, 0x9f,
	0x81, 0xb9, 0x1c, 0x9f, 0xe5, 0x6b, 0x69, 0x64, 0x73, 0x42, 0xc6, 0x48, 0xcc, 0x99, 0x1f, 0x78,
	0x23, 0x9e, 0x5d, 0xac, 0xa1, 0x14, 0x4d, 0x5e, 0xff, 0x1e, 0x2a, 0x8a, 0xd6, 0x53, 0x1b, 0xe4,
	0x5f, 0x54, 0xaa, 0xfe, 0x7b, 0x11, 0xf6, 0x36, 0xf7, 0x10, 0xb6, 0x5b, 0xc2, 0xd4, 0x33, 0x96,
	0x31, 0xea, 0x89, 0x9a, 0x3d, 0xdc, 0xf8, 0x68, 0x8a, 0x7d, 0x2e, 0xb2, 0x79, 0xf8, 0xc4, 0xc6,
	0x4f, 0x18, 0x7d, 0xec, 0x2e, 0xef, 0xc2, 0x3f, 0x5d, 0xde, 0xc5, 0xbf, 0xbb, 0xbc, 0xf1, 0x4e,
	0xd3, 0xa1, 0xda, 0xb8, 0xd9, 0xb0, 0x18, 0x71, 0x3a, 0xa4, 0xa5, 0x9a, 0x6d, 0x07, 0x6d, 0x67,
	0x3b, 0xe8, 0x3b, 0xdb, 0xc1, 0x58, 0x6e, 0x07, 0xeb, 0xff, 0xb0, 0xc7, 0xfd, 0x98, 0x0d, 0x67,
	0xc2, 0x9b, 0xcb, 0x85, 0xb0, 0x4d, 0xb5, 0xe4, 0x33, 0xdd, 0x85, 0x5c, 0x08, 0xec, 0xf3, 0x25,
	0x84, 0x25, 0xf8, 0xfa, 0x22, 0x02, 0x32, 0x55, 0x23, 0x99, 0x6e, 0x02, 0xb8, 0x18, 0xd3, 0x3b,
	0xbb, 0x06, 0xb4, 0xc4, 0x78, 0xf3, 0x10, 0x1e, 0xb1, 0x5b, 0x7a, 0x60, 0xd7, 0x87, 0xb4, 0x22,
	0x76, 0x6b, 0x7d, 0x06, 0xd5, 0x25, 0x64, 0x34, 0xa1, 0xce, 0xaf, 0x10, 0xa8, 0x92, 0x69, 0x9b,
	0xa4, 0x44, 0xa6, 0xa1, 0x1f, 0x70, 0x3f, 0xb8, 0xf6, 0xc6, 0x69, 0x30, 0xa2, 0x07, 0xd6, 0x74,
	0xcb, 0x99, 0xee, 0x4d, 0x1a, 0x8c, 0xb6, 0xbb, 0xea, 0xc1, 0x76, 0x57, 0x61, 0xa8, 0x2c, 0x8e,
	0xfd, 0xeb, 0xc0, 0x1b, 0x47, 0x72, 0x4e, 0x0f, 0xa6, 0xe9, 0x82, 0x52, 0xbd, 0x89, 0xe4, 0x1c,
	0x3b, 0x23, 0xc6, 0x49, 0xc2, 0x7e, 0xdd, 0x57, 0x9d, 0x41, 0x72, 0x9b, 0x1f, 0xd7, 0x01, 0xd6,
	0xf5, 0xb7, 0x0c, 0x28, 0x9e, 0x77, 0x7b, 0xfd, 0xda, 0x7f, 0x2c, 0x00, 0xad, 0x35, 0x70, 0x3a,
	0x8e, 0x5b, 0xcb, 0x1d, 0x77, 0xa1, 0xb2, 0x35, 0x61, 0x08, 0xbb, 0x1c, 0x74, 0x3a, 0x0a, 0xd6,
	0xe8, 0xf7, 0x1b, 0xcd, 0xb7, 0xb5, 0x9c, 0x55, 0x06, 0xbd, 0xe5, 0xbc, 0x71, 0x2e, 0x9b, 0x4e,
	0x2d, 0x6f, 0xe9, 0x50, 0x38, 0x1b, 0x5c, 0xd6, 0x0a, 0x88, 0xbd, 0xe8, 0xbe, 0x73, 0x6a, 0x45,
	0xfc, 0xea, 0xf5, 0x1b, 0x6e, 0xad, 0x74, 0xfc, 0x1d, 0xec, 0x6d, 0xee, 0x38, 0xcb, 0x84, 0xd2,
	0xeb, 0x4e, 0xb7, 0xf9, 0x56, 0x11, 0xb6, 0x2f, 0x7b, 0xed, 0x96, 0xa3, 0x08, 0xbb, 0x83, 0x3e,
	0x09, 0x79, 0xc5, 0xde, 0xeb, 0xbb, 0xdd, 0x9f, 0x6a, 0x85, 0xe3, 0x17, 0x60, 0xae, 0x56, 0x06,
	0x7a, 0xf7, 0xce, 0xbb, 0x6e, 0x16, 0xf5, 0x45, 0xbb, 0xd5, 0xea, 0xa0, 0xb7, 0x01, 0xc5, 0x4e,
	0xf7, 0xf2, 0xac, 0x96, 0x3f, 0x7e, 0x01, 0x9a, 0x9a, 0x6f, 0x84, 0x3a, 0x3f, 0x0c, 0x1a, 0x18,
	0x79, 0x19, 0xf4, 0x33, 0xd7, 0x69, 0xf4, 0x31, 0x43, 0xf4, 0xeb, 0x38, 0xbd, 0x9e, 0xe3, 0xd6,
	0xf2, 0xaf, 0x9f, 0xbd, 0x2f, 0xd1, 0xbf, 0xe7, 0xdf, 0xf2, 0xfb, 0x83, 0xce, 0xfb, 0x46, 0x1c,
	0x8b, 0xe4, 0x84, 0x7a, 0xf7, 0x42, 0xf2, 0xa1, 0x46, 0xa6, 0x97, 0x7f, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x18, 0xd0, 0x19, 0x2e, 0x86, 0x0b, 0x00, 0x00,
}

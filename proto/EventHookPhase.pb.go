// Code generated by protoc-gen-go. DO NOT EDIT.
// source: EventHookPhase.proto

package ULZProto

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

// ----------------------------------------------------
// ANCHOR EventHookPhase.proto
// SECTION: EventHookPhase.proto
type EventHookPhase int32

const (
	// start gameset
	EventHookPhase_gameset_start EventHookPhase = 0
	// -- turn lifecycle
	EventHookPhase_start_turn_phase EventHookPhase = 10
	// ---- Draw phase (ddp)
	EventHookPhase_refill_action_card_phase EventHookPhase = 20
	EventHookPhase_refill_event_card_phase  EventHookPhase = 25
	// ---- Move phase
	EventHookPhase_move_card_drop_phase         EventHookPhase = 30
	EventHookPhase_determine_move_phase         EventHookPhase = 40
	EventHookPhase_finish_move_phase            EventHookPhase = 50
	EventHookPhase_chara_change_phase           EventHookPhase = 60
	EventHookPhase_determine_chara_change_phase EventHookPhase = 70
	// ---- Atk/Def Phase
	EventHookPhase_attack_card_drop_phase            EventHookPhase = 80
	EventHookPhase_defence_card_drop_phase           EventHookPhase = 90
	EventHookPhase_determine_battle_point_phase      EventHookPhase = 100
	EventHookPhase_battle_result_phase               EventHookPhase = 110
	EventHookPhase_damage_phase                      EventHookPhase = 120
	EventHookPhase_dead_chara_change_phase           EventHookPhase = 125
	EventHookPhase_determine_dead_chara_change_phase EventHookPhase = 130
	EventHookPhase_change_initiative_phase           EventHookPhase = 140
	/// -- endof turn lifecycle
	EventHookPhase_finish_turn_phase EventHookPhase = 150
	/// endof game set
	EventHookPhase_gameset_end EventHookPhase = 160
)

var EventHookPhase_name = map[int32]string{
	0:   "gameset_start",
	10:  "start_turn_phase",
	20:  "refill_action_card_phase",
	25:  "refill_event_card_phase",
	30:  "move_card_drop_phase",
	40:  "determine_move_phase",
	50:  "finish_move_phase",
	60:  "chara_change_phase",
	70:  "determine_chara_change_phase",
	80:  "attack_card_drop_phase",
	90:  "defence_card_drop_phase",
	100: "determine_battle_point_phase",
	110: "battle_result_phase",
	120: "damage_phase",
	125: "dead_chara_change_phase",
	130: "determine_dead_chara_change_phase",
	140: "change_initiative_phase",
	150: "finish_turn_phase",
	160: "gameset_end",
}

var EventHookPhase_value = map[string]int32{
	"gameset_start":                     0,
	"start_turn_phase":                  10,
	"refill_action_card_phase":          20,
	"refill_event_card_phase":           25,
	"move_card_drop_phase":              30,
	"determine_move_phase":              40,
	"finish_move_phase":                 50,
	"chara_change_phase":                60,
	"determine_chara_change_phase":      70,
	"attack_card_drop_phase":            80,
	"defence_card_drop_phase":           90,
	"determine_battle_point_phase":      100,
	"battle_result_phase":               110,
	"damage_phase":                      120,
	"dead_chara_change_phase":           125,
	"determine_dead_chara_change_phase": 130,
	"change_initiative_phase":           140,
	"finish_turn_phase":                 150,
	"gameset_end":                       160,
}

func (x EventHookPhase) String() string {
	return proto.EnumName(EventHookPhase_name, int32(x))
}

func (EventHookPhase) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c9ab9311bf956e13, []int{0}
}

type EventHookType int32

const (
	EventHookType_Instant EventHookType = 0
	EventHookType_Before  EventHookType = 1
	EventHookType_After   EventHookType = 2
	EventHookType_Proxy   EventHookType = 3
)

var EventHookType_name = map[int32]string{
	0: "Instant",
	1: "Before",
	2: "After",
	3: "Proxy",
}

var EventHookType_value = map[string]int32{
	"Instant": 0,
	"Before":  1,
	"After":   2,
	"Proxy":   3,
}

func (x EventHookType) String() string {
	return proto.EnumName(EventHookType_name, int32(x))
}

func (EventHookType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c9ab9311bf956e13, []int{1}
}

func init() {
	proto.RegisterEnum("ULZProto.EventHookPhase", EventHookPhase_name, EventHookPhase_value)
	proto.RegisterEnum("ULZProto.EventHookType", EventHookType_name, EventHookType_value)
}

func init() { proto.RegisterFile("EventHookPhase.proto", fileDescriptor_c9ab9311bf956e13) }

var fileDescriptor_c9ab9311bf956e13 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x31, 0x6f, 0xd4, 0x40,
	0x10, 0x85, 0x71, 0x80, 0x00, 0x13, 0x82, 0xf6, 0x06, 0x73, 0x17, 0xe0, 0x84, 0xa0, 0x41, 0x28,
	0x45, 0x0a, 0x68, 0x11, 0x52, 0x22, 0x81, 0x40, 0x0a, 0x92, 0x0b, 0xd2, 0x5c, 0xb3, 0x9a, 0x78,
	0xc7, 0xf1, 0x2a, 0xf6, 0xae, 0xb5, 0x3b, 0x89, 0x92, 0x82, 0x86, 0x3a, 0x35, 0x35, 0x35, 0xbf,
	0x12, 0xd9, 0xe7, 0x73, 0xee, 0x74, 0x47, 0x37, 0x7a, 0xdf, 0xd3, 0xec, 0xd3, 0xce, 0x83, 0xf4,
	0xf3, 0x25, 0x3b, 0xf9, 0xea, 0xfd, 0x79, 0x56, 0x52, 0xe4, 0x83, 0x26, 0x78, 0xf1, 0xf8, 0xf0,
	0xe4, 0x78, 0x96, 0xb5, 0xd3, 0xfe, 0xcd, 0x3d, 0x78, 0xb2, 0x6a, 0xc1, 0x11, 0xec, 0x9e, 0x51,
	0xcd, 0x91, 0x45, 0x47, 0xa1, 0x20, 0xea, 0x0e, 0xa6, 0xa0, 0xba, 0x51, 0xcb, 0x45, 0x70, 0xba,
	0x69, 0x6d, 0x0a, 0x70, 0x0a, 0x7b, 0x81, 0x0b, 0x5b, 0x55, 0x9a, 0x72, 0xb1, 0xde, 0xe9, 0x9c,
	0x82, 0xe9, 0x69, 0x8a, 0x2f, 0x61, 0xd2, 0x53, 0x6e, 0xf7, 0x2f, 0xc3, 0xe7, 0xb8, 0x07, 0x69,
	0xed, 0x2f, 0x79, 0x2e, 0x9a, 0xe0, 0x9b, 0x9e, 0xbc, 0x6a, 0x89, 0x61, 0xe1, 0x50, 0x5b, 0xc7,
	0xba, 0xf3, 0xcc, 0xc9, 0x3b, 0x7c, 0x06, 0xa3, 0xc2, 0x3a, 0x1b, 0xcb, 0x65, 0xf9, 0x3d, 0x8e,
	0x01, 0xf3, 0x92, 0x02, 0xe9, 0xbc, 0x24, 0x77, 0xb6, 0xd0, 0x3f, 0xe2, 0x6b, 0x98, 0xde, 0x2e,
	0xda, 0xe0, 0xf8, 0x82, 0x2f, 0x60, 0x4c, 0x22, 0x94, 0x9f, 0xaf, 0xc5, 0xc8, 0xda, 0xf4, 0x86,
	0x0b, 0x76, 0xf9, 0x7a, 0xc6, 0xd9, 0xea, 0xea, 0x53, 0x12, 0xa9, 0x58, 0x37, 0xde, 0x3a, 0xe9,
	0x1d, 0x06, 0x27, 0xf0, 0xb4, 0xd7, 0x03, 0xc7, 0x8b, 0x6a, 0x01, 0x1c, 0x2a, 0x78, 0x6c, 0xa8,
	0xa6, 0x21, 0xc5, 0xd5, 0xfc, 0x25, 0x32, 0x9b, 0x22, 0xfe, 0xc4, 0xb7, 0xf0, 0xe6, 0xf6, 0xa5,
	0xff, 0xd9, 0x7e, 0x25, 0x38, 0x85, 0x49, 0x2f, 0x59, 0x67, 0xc5, 0x92, 0xd8, 0xe1, 0x87, 0x6e,
	0x12, 0x1c, 0x0f, 0x3f, 0xb7, 0x74, 0xbf, 0xdf, 0x09, 0x2a, 0xd8, 0x59, 0x5c, 0x9a, 0x9d, 0x51,
	0x7f, 0x92, 0xfd, 0x4f, 0xb0, 0x3b, 0xb4, 0xe1, 0xc7, 0x75, 0xc3, 0xb8, 0x03, 0x0f, 0xbe, 0xb9,
	0x28, 0xe4, 0xda, 0x1a, 0x00, 0x6c, 0x1f, 0x71, 0xe1, 0x03, 0xab, 0x04, 0x1f, 0xc1, 0xfd, 0xc3,
	0x42, 0x38, 0xa8, 0xad, 0x76, 0xcc, 0x82, 0xbf, 0xba, 0x56, 0x77, 0x8f, 0xd2, 0xbf, 0x5b, 0xa3,
	0x93, 0xe3, 0xd9, 0x61, 0x8c, 0x2c, 0x07, 0x5d, 0xc3, 0xbe, 0x7b, 0x73, 0xba, 0xdd, 0xb5, 0xee,
	0xc3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x5e, 0x60, 0x01, 0x8d, 0x02, 0x00, 0x00,
}
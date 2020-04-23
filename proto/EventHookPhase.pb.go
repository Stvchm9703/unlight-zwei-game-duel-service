// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.10.0
// source: EventHookPhase.proto

package proto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// ----------------------------------------------------
// ANCHOR EventHookPhase.proto
// SECTION: EventHookPhase.proto
type EventHookPhase int32

const (
	// start gameset
	EventHookPhase_gameset_start EventHookPhase = 0
	// -- turn lifecycle
	EventHookPhase_start_turn_phase EventHookPhase = 1
	// ---- Draw phase (ddp)
	EventHookPhase_refill_action_card_phase EventHookPhase = 2
	//   refill_event_card_phase                        = 3;
	// ---- Move phase
	EventHookPhase_move_card_drop_phase         EventHookPhase = 3
	EventHookPhase_determine_move_phase         EventHookPhase = 4
	EventHookPhase_finish_move_phase            EventHookPhase = 5
	EventHookPhase_chara_change_phase           EventHookPhase = 6
	EventHookPhase_determine_chara_change_phase EventHookPhase = 7
	// ---- Atk/Def Phase
	EventHookPhase_attack_card_drop_phase       EventHookPhase = 8
	EventHookPhase_defence_card_drop_phase      EventHookPhase = 9
	EventHookPhase_determine_battle_point_phase EventHookPhase = 10
	EventHookPhase_battle_result_phase          EventHookPhase = 11 //roll dice
	EventHookPhase_damage_phase                 EventHookPhase = 12
	// --- Any Raise Phase
	EventHookPhase_dead_chara_change_phase           EventHookPhase = 13
	EventHookPhase_determine_dead_chara_change_phase EventHookPhase = 14
	EventHookPhase_change_initiative_phase           EventHookPhase = 15
	/// -- endof turn lifecycle
	EventHookPhase_finish_turn_phase EventHookPhase = 16
	/// endof game set
	EventHookPhase_gameset_end EventHookPhase = 17
)

// Enum value maps for EventHookPhase.
var (
	EventHookPhase_name = map[int32]string{
		0:  "gameset_start",
		1:  "start_turn_phase",
		2:  "refill_action_card_phase",
		3:  "move_card_drop_phase",
		4:  "determine_move_phase",
		5:  "finish_move_phase",
		6:  "chara_change_phase",
		7:  "determine_chara_change_phase",
		8:  "attack_card_drop_phase",
		9:  "defence_card_drop_phase",
		10: "determine_battle_point_phase",
		11: "battle_result_phase",
		12: "damage_phase",
		13: "dead_chara_change_phase",
		14: "determine_dead_chara_change_phase",
		15: "change_initiative_phase",
		16: "finish_turn_phase",
		17: "gameset_end",
	}
	EventHookPhase_value = map[string]int32{
		"gameset_start":                     0,
		"start_turn_phase":                  1,
		"refill_action_card_phase":          2,
		"move_card_drop_phase":              3,
		"determine_move_phase":              4,
		"finish_move_phase":                 5,
		"chara_change_phase":                6,
		"determine_chara_change_phase":      7,
		"attack_card_drop_phase":            8,
		"defence_card_drop_phase":           9,
		"determine_battle_point_phase":      10,
		"battle_result_phase":               11,
		"damage_phase":                      12,
		"dead_chara_change_phase":           13,
		"determine_dead_chara_change_phase": 14,
		"change_initiative_phase":           15,
		"finish_turn_phase":                 16,
		"gameset_end":                       17,
	}
)

func (x EventHookPhase) Enum() *EventHookPhase {
	p := new(EventHookPhase)
	*p = x
	return p
}

func (x EventHookPhase) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventHookPhase) Descriptor() protoreflect.EnumDescriptor {
	return file_EventHookPhase_proto_enumTypes[0].Descriptor()
}

func (EventHookPhase) Type() protoreflect.EnumType {
	return &file_EventHookPhase_proto_enumTypes[0]
}

func (x EventHookPhase) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventHookPhase.Descriptor instead.
func (EventHookPhase) EnumDescriptor() ([]byte, []int) {
	return file_EventHookPhase_proto_rawDescGZIP(), []int{0}
}

type EventHookType int32

const (
	EventHookType_Instant EventHookType = 0
	EventHookType_Before  EventHookType = 1
	EventHookType_Proxy   EventHookType = 2
	EventHookType_After   EventHookType = 3
)

// Enum value maps for EventHookType.
var (
	EventHookType_name = map[int32]string{
		0: "Instant",
		1: "Before",
		2: "Proxy",
		3: "After",
	}
	EventHookType_value = map[string]int32{
		"Instant": 0,
		"Before":  1,
		"Proxy":   2,
		"After":   3,
	}
)

func (x EventHookType) Enum() *EventHookType {
	p := new(EventHookType)
	*p = x
	return p
}

func (x EventHookType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventHookType) Descriptor() protoreflect.EnumDescriptor {
	return file_EventHookPhase_proto_enumTypes[1].Descriptor()
}

func (EventHookType) Type() protoreflect.EnumType {
	return &file_EventHookPhase_proto_enumTypes[1]
}

func (x EventHookType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventHookType.Descriptor instead.
func (EventHookType) EnumDescriptor() ([]byte, []int) {
	return file_EventHookPhase_proto_rawDescGZIP(), []int{1}
}

var File_EventHookPhase_proto protoreflect.FileDescriptor

var file_EventHookPhase_proto_rawDesc = []byte{
	0x0a, 0x14, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x50, 0x68, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x55, 0x4c, 0x5a, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0xeb, 0x03, 0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x50, 0x68,
	0x61, 0x73, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x74, 0x75, 0x72, 0x6e, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18,
	0x72, 0x65, 0x66, 0x69, 0x6c, 0x6c, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61,
	0x72, 0x64, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x10, 0x02, 0x12, 0x18, 0x0a, 0x14, 0x6d, 0x6f,
	0x76, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x72, 0x6f, 0x70, 0x5f, 0x70, 0x68, 0x61,
	0x73, 0x65, 0x10, 0x03, 0x12, 0x18, 0x0a, 0x14, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e,
	0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x10, 0x04, 0x12, 0x15,
	0x0a, 0x11, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x70, 0x68,
	0x61, 0x73, 0x65, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x63, 0x68, 0x61, 0x72, 0x61, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x10, 0x06, 0x12, 0x20, 0x0a,
	0x1c, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x10, 0x07, 0x12,
	0x1a, 0x0a, 0x16, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x64,
	0x72, 0x6f, 0x70, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x10, 0x08, 0x12, 0x1b, 0x0a, 0x17, 0x64,
	0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x64, 0x72, 0x6f, 0x70,
	0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x10, 0x09, 0x12, 0x20, 0x0a, 0x1c, 0x64, 0x65, 0x74, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x65, 0x5f, 0x62, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x5f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x10, 0x0a, 0x12, 0x17, 0x0a, 0x13, 0x62, 0x61,
	0x74, 0x74, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x70, 0x68, 0x61, 0x73,
	0x65, 0x10, 0x0b, 0x12, 0x10, 0x0a, 0x0c, 0x64, 0x61, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x68,
	0x61, 0x73, 0x65, 0x10, 0x0c, 0x12, 0x1b, 0x0a, 0x17, 0x64, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x68,
	0x61, 0x72, 0x61, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65,
	0x10, 0x0d, 0x12, 0x25, 0x0a, 0x21, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x5f,
	0x64, 0x65, 0x61, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x61, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x10, 0x0e, 0x12, 0x1b, 0x0a, 0x17, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x70,
	0x68, 0x61, 0x73, 0x65, 0x10, 0x0f, 0x12, 0x15, 0x0a, 0x11, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68,
	0x5f, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x70, 0x68, 0x61, 0x73, 0x65, 0x10, 0x10, 0x12, 0x0f, 0x0a,
	0x0b, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x65, 0x6e, 0x64, 0x10, 0x11, 0x2a, 0x3e,
	0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x6f, 0x6f, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x66, 0x74, 0x65, 0x72, 0x10, 0x03, 0x42, 0x1d,
	0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x11, 0x55, 0x4c, 0x5a, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x4d, 0x6f, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_EventHookPhase_proto_rawDescOnce sync.Once
	file_EventHookPhase_proto_rawDescData = file_EventHookPhase_proto_rawDesc
)

func file_EventHookPhase_proto_rawDescGZIP() []byte {
	file_EventHookPhase_proto_rawDescOnce.Do(func() {
		file_EventHookPhase_proto_rawDescData = protoimpl.X.CompressGZIP(file_EventHookPhase_proto_rawDescData)
	})
	return file_EventHookPhase_proto_rawDescData
}

var file_EventHookPhase_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_EventHookPhase_proto_goTypes = []interface{}{
	(EventHookPhase)(0), // 0: ULZProto.EventHookPhase
	(EventHookType)(0),  // 1: ULZProto.EventHookType
}
var file_EventHookPhase_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_EventHookPhase_proto_init() }
func file_EventHookPhase_proto_init() {
	if File_EventHookPhase_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_EventHookPhase_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_EventHookPhase_proto_goTypes,
		DependencyIndexes: file_EventHookPhase_proto_depIdxs,
		EnumInfos:         file_EventHookPhase_proto_enumTypes,
	}.Build()
	File_EventHookPhase_proto = out.File
	file_EventHookPhase_proto_rawDesc = nil
	file_EventHookPhase_proto_goTypes = nil
	file_EventHookPhase_proto_depIdxs = nil
}

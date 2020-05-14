# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Protocol Documentation](#protocol-documentation)
  - [Table of Contents](#table-of-contents)
  - [StoreMod.proto](#storemodproto)
    - [ADPhaseSnapMod](#adphasesnapmod)
    - [EffectNodeSnapMod](#effectnodesnapmod)
    - [MovePhaseSnapMod](#movephasesnapmod)
    - [PhaseSnapMod](#phasesnapmod)
  - [Scalar Value Types](#scalar-value-types)
  
- [Scalar Value Types](#scalar-value-types)



<a name="StoreMod.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## StoreMod.proto



<a name="ULZProto.ADPhaseSnapMod"></a>

### ADPhaseSnapMod
AD-Phase-SnapMod


| Field            | Type                                       | Label    | Description |
| ---------------- | ------------------------------------------ | -------- | ----------- |
| turns            | [int32](#int32)                            |          |             |
| first_attack     | [PlayerSide](#ULZProto.PlayerSide)         |          |             |
| curr_attacker    | [PlayerSide](#ULZProto.PlayerSide)         |          |             |
| event_phase      | [EventHookPhase](#ULZProto.EventHookPhase) |          |             |
| attack_val       | [int32](#int32)                            |          |             |
| defence_val      | [int32](#int32)                            |          |             |
| attack_card      | [EventCard](#ULZProto.EventCard)           | repeated |             |
| defence_card     | [EventCard](#ULZProto.EventCard)           | repeated |             |
| attack_trig_skl  | [SkillSet](#ULZProto.SkillSet)             | repeated |             |
| defence_trig_skl | [SkillSet](#ULZProto.SkillSet)             | repeated |             |
| is_processed     | [bool](#bool)                              |          |             |






<a name="ULZProto.EffectNodeSnapMod"></a>

### EffectNodeSnapMod
Effect-Status-SnapMod


| Field      | Type                                   | Label    | Description |
| ---------- | -------------------------------------- | -------- | ----------- |
| turns      | [int32](#int32)                        |          |             |
| pending_ef | [EffectResult](#ULZProto.EffectResult) | repeated |             |






<a name="ULZProto.MovePhaseSnapMod"></a>

### MovePhaseSnapMod
Move-Phase-SnapMod


| Field         | Type                                   | Label    | Description |
| ------------- | -------------------------------------- | -------- | ----------- |
| turns         | [int32](#int32)                        |          | flaging     |
| host_val      | [int32](#int32)                        |          |             |
| duel_val      | [int32](#int32)                        |          |             |
| host_opt      | [MovePhaseOpt](#ULZProto.MovePhaseOpt) |          |             |
| duel_opt      | [MovePhaseOpt](#ULZProto.MovePhaseOpt) |          |             |
| host_card     | [EventCard](#ULZProto.EventCard)       | repeated |             |
| duel_card     | [EventCard](#ULZProto.EventCard)       | repeated |             |
| host_trig_skl | [SkillSet](#ULZProto.SkillSet)         | repeated |             |
| duel_trig_skl | [SkillSet](#ULZProto.SkillSet)         | repeated |             |






<a name="ULZProto.PhaseSnapMod"></a>

### PhaseSnapMod
PhaseInstance-SnapMod


| Field         | Type                                       | Label | Description                   |
| ------------- | ------------------------------------------ | ----- | ----------------------------- |
| turns         | [int32](#int32)                            |       |                               |
| event_phase   | [EventHookPhase](#ULZProto.EventHookPhase) |       |                               |
| hook_type     | [EventHookType](#ULZProto.EventHookType)   |       |                               |
| is_host_ready | [bool](#bool)                              |       |                               |
| is_duel_ready | [bool](#bool)                              |       |                               |
| first_attack  | [PlayerSide](#ULZProto.PlayerSide)         |       |                               |
| curr_attack   | [PlayerSide](#ULZProto.PlayerSide)         |       | EventHookType hook_type = 16; |





 

 

 

 



## Scalar Value Types

| .proto Type                    | Notes                                                                                                                                           | C++    | Java       | Python      | Go      | C#         | PHP            | Ruby                           |
| ------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- | ------ | ---------- | ----------- | ------- | ---------- | -------------- | ------------------------------ |
| <a name="double" /> double     |                                                                                                                                                 | double | double     | float       | float64 | double     | float          | Float                          |
| <a name="float" /> float       |                                                                                                                                                 | float  | float      | float       | float32 | float      | float          | Float                          |
| <a name="int32" /> int32       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="int64" /> int64       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="uint32" /> uint32     | Uses variable-length encoding.                                                                                                                  | uint32 | int        | int/long    | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64     | Uses variable-length encoding.                                                                                                                  | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.                            | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.                            | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="fixed32" /> fixed32   | Always four bytes. More efficient than uint32 if values are often greater than 2^28.                                                            | uint32 | int        | int         | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64   | Always eight bytes. More efficient than uint64 if values are often greater than 2^56.                                                           | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum                         |
| <a name="sfixed32" /> sfixed32 | Always four bytes.                                                                                                                              | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes.                                                                                                                             | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="bool" /> bool         |                                                                                                                                                 | bool   | boolean    | boolean     | bool    | bool       | boolean        | TrueClass/FalseClass           |
| <a name="string" /> string     | A string must always contain UTF-8 encoded or 7-bit ASCII text.                                                                                 | string | String     | str/unicode | string  | string     | string         | String (UTF-8)                 |
| <a name="bytes" /> bytes       | May contain any arbitrary sequence of bytes.                                                                                                    | string | ByteString | str         | []byte  | ByteString | string         | String (ASCII-8BIT)            |


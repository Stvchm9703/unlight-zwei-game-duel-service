# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Protocol Documentation](#protocol-documentation)
  - [Table of Contents](#table-of-contents)
  - [EventHookPhase.proto](#eventhookphaseproto)
    - [EventHookPhase](#eventhookphase)
    - [EventHookType](#eventhooktype)
  - [Scalar Value Types](#scalar-value-types)
  
- [Scalar Value Types](#scalar-value-types)



<a name="EventHookPhase.proto"></a>

<p align="right"><a href="#top">Top</a></p>

## EventHookPhase.proto

 

----------------------------------------------------

<a name="ULZProto.EventHookPhase"></a>

### EventHookPhase
ANCHOR EventHookPhase.proto
SECTION: EventHookPhase.proto

| Name                              | Number | Description          |
| --------------------------------- | ------ | -------------------- |
| gameset_start                     | 0      | start gameset        |
| start_turn_phase                  | 1      | turn lifecycle       |
| refill_action_card_phase          | 2      | Draw phase           |
| move_card_drop_phase              | 3      | Move phase           |
| determine_move_phase              | 4      |                      |
| finish_move_phase                 | 5      |                      |
| chara_change_phase                | 6      |                      |
| determine_chara_change_phase      | 7      |                      |
| attack_card_drop_phase            | 8      | Atk Phase            |
| defence_card_drop_phase           | 9      | Def Phase            |
| determine_battle_point_phase      | 10     |                      |
| battle_result_phase               | 11     | roll dice            |
| damage_phase                      | 12     |                      |
| dead_chara_change_phase           | 13     | Any Raise Phase      |
| determine_dead_chara_change_phase | 14     |                      |
| change_initiative_phase           | 15     |                      |
| finish_turn_phase                 | 16     | endof turn lifecycle |
| gameset_end                       | 17     | endof game set       |

```mermaid 
stateDiagram 
gameset_start --> Turn_cycle
Turn_cycle --> gameset_end

state Turn_cycle {
  [*] --> start_turn_phase 
  start_turn_phase --> Draw_Phase 
  Draw_Phase --> Move_Phase
  Move_Phase --> ADPhase
  AD_Phase --> change_initiative_phase 

  state change_initiative_phase <<fork>>
    change_initiative_phase --> AD_Phase
    change_initiative_phase --> finish_turn_phase

  finish_turn_phase --> [*]
  finish_turn_phase --> start_turn_phase
  
  state Draw_Phase {
    [*] -->   refill_action_card_phase 
    refill_action_card_phase --> [*]
  }

  state Move_Phase {
    [*] --> move_card_phase
    move_card_phase --> determine_move_phase
    determine_move_phase --> finish_move_phase
    state finish_move_phase <<fork>>
      finish_move_phase --> chara_change_phase
      chara_change_phase--> determine_chara_change_phase 
      determine_chara_change_phase --> [*]
      finish_move_phase --> [*]
  }

  state AD_Phase{
    [*] --> attack_card_drop_phase 
    attack_card_drop_phase --> defence_card_drop_phase
    defence_card_drop_phase --> determine_battle_point_phase
    determine_battle_point_phase --> battle_result_phase
    battle_result_phase --> damage_phase
    state damage_phase <<fork>>
      damage_phase --> dead_chara_change_phase 
      dead_chara_change_phase --> determine_dead_chara_change_phase
      determine_dead_chara_change_phase --> [*]
      damage_phase  --> [*]
  }
}

```


```mermaid 
stateDiagram 
gameset_start --> start_turn_phase 
start_turn_phase --> refill_action_card_phase 
refill_action_card_phase --> move_card_phase
move_card_phase --> determine_move_phase
determine_move_phase --> finish_move_phase
state finish_move_phase <<fork>>
  finish_move_phase --> chara_change_phase
  chara_change_phase--> determine_chara_change_phase 
  determine_chara_change_phase --> attack_card_drop_phase
  finish_move_phase --> attack_card_drop_phase


attack_card_drop_phase --> defence_card_drop_phase
defence_card_drop_phase --> determine_battle_point_phase
determine_battle_point_phase --> battle_result_phase
battle_result_phase --> damage_phase
state damage_phase <<fork>>
  damage_phase --> dead_chara_change_phase 
  dead_chara_change_phase --> determine_dead_chara_change_phase
  determine_dead_chara_change_phase --> change_initiative_phase
  damage_phase  --> change_initiative_phase


state change_initiative_phase <<fork>>
  change_initiative_phase --> attack_card_drop_phase
  change_initiative_phase --> finish_turn_phase

  finish_turn_phase --> [*]
  finish_turn_phase --> start_turn_phase
  
 


```

<a name="ULZProto.EventHookType"></a>

### EventHookType


| Name    | Number | Description |
| ------- | ------ | ----------- |
| Instant | 0      |             |
| Before  | 1      |             |
| Proxy   | 2      |             |
| After   | 3      |             |


 

 

 



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


# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Protocol Documentation](#protocol-documentation)
  - [Table of Contents](#table-of-contents)
  - [message.proto](#messageproto)
    - [GDCreateReq](#gdcreatereq)
    - [GDGetInfoReq](#gdgetinforeq)
  - [Broadcast Related](#broadcast-related)
    - [GDBroadcastResp](#gdbroadcastresp)
    - [ECShortHand](#ecshorthand)
    - [CastCmd](#castcmd)
    - [GDInstanceDT](#gdinstancedt)
  - [Change Character Phase](#change-character-phase)
    - [GDChangeConfirmReq](#gdchangeconfirmreq)
  - [Move Phase Handle](#move-phase-handle)
    - [GDMoveConfirmReq](#gdmoveconfirmreq)
    - [GDMoveConfirmResp](#gdmoveconfirmresp)
  - [Attack-Phase / Defernce-Phase Handle Related](#attack-phase--defernce-phase-handle-related)
    - [GDADConfirmReq](#gdadconfirmreq)
    - [GDADResultResp](#gdadresultresp)
    - [GDADDiceResult](#gdaddiceresult)
  - [Common Phase Handle](#common-phase-handle)
    - [GDPhaseConfirmReq](#gdphaseconfirmreq)
    - [GDPhaseConfirmResp](#gdphaseconfirmresp)
  - [Scalar Value Types](#scalar-value-types)



<a name="message.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## message.proto


-------------------------------------------------------------

<a name="ULZProto.GDCreateReq"></a>

### GDCreateReq
Create-Game Request


| Field          | Type                                 | Label    | Description |
| -------------- | ------------------------------------ | -------- | ----------- |
| room_key       | [string](#string)                    |          |             |
| host_id        | [string](#string)                    |          |             |
| dueler_id      | [string](#string)                    |          |             |
| nvn            | [int32](#int32)                      |          |             |
| host_card_deck | [CharCardSet](#ULZProto.CharCardSet) | repeated |             |
| duel_card_deck | [CharCardSet](#ULZProto.CharCardSet) | repeated |             |
| host_extra_ec  | [EventCard](#ULZProto.EventCard)     | repeated |             |
| duel_extra_ec  | [EventCard](#ULZProto.EventCard)     | repeated |             |






<a name="ULZProto.GDGetInfoReq"></a>

### GDGetInfoReq
Get-Game Request
 **general-struct


| Field          | Type                                       | Label | Description |
| -------------- | ------------------------------------------ | ----- | ----------- |
| room_key       | [string](#string)                          |       |             |
| side           | [PlayerSide](#ULZProto.PlayerSide)         |       |             |
| current_phase  | [EventHookPhase](#ULZProto.EventHookPhase) |       |             |
| is_watcher     | [bool](#bool)                              |       |             |
| income_user_id | [string](#string)                          |       |             |



-------------------------------------------------------------
## Broadcast Related

<a name="ULZProto.GDBroadcastResp"></a>


### GDBroadcastResp
Server-Broadcast-Response


| Field         | Type                                       | Label    | Description |
| ------------- | ------------------------------------------ | -------- | ----------- |
| room_key      | [string](#string)                          |          |             |
| msg           | [string](#string)                          |          |             |
| command       | [CastCmd](#ULZProto.CastCmd)               |          |             |
| current_phase | [EventHookPhase](#ULZProto.EventHookPhase) |          |             |
| phase_hook    | [EventHookType](#ULZProto.EventHookType)   |          |             |
| side          | [PlayerSide](#ULZProto.PlayerSide)         |          |             |
| instance_set  | [ECShortHand](#ULZProto.ECShortHand)       | repeated |             |
| effect_trig   | [EffectResult](#ULZProto.EffectResult)     | repeated |             |




<a name="ULZProto.ECShortHand"></a>

### ECShortHand
Sub-helper-struct, to reduce the data size during transmission. 

| Field    | Type                                   | Label | Description                   |
| -------- | -------------------------------------- | ----- | ----------------------------- |
| card_id  | [int32](#int32)                        |       | instance event-card id        |
| position | [EventCardPos](#ULZProto.EventCardPos) |       | position of event-card        |
| isInvert | [bool](#bool)                          |       | Invert flag of the event-card |


<a name="ULZProto.CastCmd"></a>

### CastCmd

| Name                   | Number | Description |
| ---------------------- | ------ | ----------- |
| GET_EFFECT_RESULT      | 0      |             |
| GET_DRAW_PHASE_RESULT  | 1      |             |
| GET_MOVE_PHASE_RESULT  | 2      |             |
| GET_ATK_PHASE_RESULT   | 3      |             |
| GET_DEF_PHASE_RESULT   | 4      |             |
| GET_INSTANCE_CARD      | 5      |             |
| GET_GAMESET_RESULT     | 6      |             |
| INSTANCE_DAMAGE        | 7      |             |
| INSTANCE_STATUS_CHANGE | 8      |             |


-------------------------------------------------------------

<a name="ULZProto.GDInstanceDT"></a>

### GDInstanceDT
instance-set-event-card


| Field         | Type                                       | Label    | Description |
| ------------- | ------------------------------------------ | -------- | ----------- |
| room_key      | [string](#string)                          |          |             |
| side          | [PlayerSide](#ULZProto.PlayerSide)         |          |             |
| current_phase | [EventHookPhase](#ULZProto.EventHookPhase) |          |             |
| update_card   | [ECShortHand](#ULZProto.ECShortHand)       | repeated |             |


-------------------------------------------------------------
## Change Character Phase

<a name="ULZProto.GDChangeConfirmReq"></a>

### GDChangeConfirmReq


| Field    | Type                               | Label | Description |
| -------- | ---------------------------------- | ----- | ----------- |
| room_key | [string](#string)                  |       |             |
| side     | [PlayerSide](#ULZProto.PlayerSide) |       |             |
| card_num | [int32](#int32)                    |       |             |


-------------------------------------------------------------

## Move Phase Handle

<a name="ULZProto.GDMoveConfirmReq"></a>

### GDMoveConfirmReq
Move-Phase-Confirm Request


| Field       | Type                                   | Label    | Description |
| ----------- | -------------------------------------- | -------- | ----------- |
| room_key    | [string](#string)                      |          |             |
| side        | [PlayerSide](#ULZProto.PlayerSide)     |          |             |
| update_card | [EventCard](#ULZProto.EventCard)       | repeated |             |
| move_opt    | [MovePhaseOpt](#ULZProto.MovePhaseOpt) |          |             |
| point       | [int32](#int32)                        |          |             |
| trigger_skl | [SkillSet](#ULZProto.SkillSet)         | repeated |             |


<a name="ULZProto.GDMoveConfirmResp"></a>

### GDMoveConfirmResp

| Field          | Type                             | Label | Description |
| -------------- | -------------------------------- | ----- | ----------- |
| room_key       | [string](#string)                |       |             |
| result_range   | [RangeType](#ULZProto.RangeType) |       |             |
| host_hp        | [int32](#int32)                  |       |             |
| duel_hp        | [int32](#int32)                  |       |             |
| host_curr_card | [int32](#int32)                  |       |             |
| duel_curr_card | [int32](#int32)                  |       |             |


-------------------------------------------------------------

## Attack-Phase / Defernce-Phase Handle Related


<a name="ULZProto.GDADConfirmReq"></a>

### GDADConfirmReq
ATK/DEF-Phase-Confirm Request
     Player send data


| Field         | Type                                       | Label    | Description                                                                                      |
| ------------- | ------------------------------------------ | -------- | ------------------------------------------------------------------------------------------------ |
| room_key      | [string](#string)                          |          |                                                                                                  |
| side          | [PlayerSide](#ULZProto.PlayerSide)         |          |                                                                                                  |
| current_phase | [EventHookPhase](#ULZProto.EventHookPhase) |          |                                                                                                  |
| trigger_skl   | [SkillSet](#ULZProto.SkillSet)             | repeated |                                                                                                  |
| update_card   | [EventCard](#ULZProto.EventCard)           | repeated | update-card : target event-card to used  !NOTE: suppose the card data is stored during card-move |



<a name="ULZProto.GDADResultResp"></a>

### GDADResultResp
ATK/DEF-Phase-Result Response


| Field         | Type                                       | Label | Description                           |
| ------------- | ------------------------------------------ | ----- | ------------------------------------- |
| room_key      | [string](#string)                          |       |                                       |
| side          | [PlayerSide](#ULZProto.PlayerSide)         |       |                                       |
| current_phase | [EventHookPhase](#ULZProto.EventHookPhase) |       |                                       |
| point         | [int32](#int32)                            |       | total Atk / Def Point after feat-func |


<a name="ULZProto.GDADDiceResult"></a>

### GDADDiceResult
ATK/DEF-Phase-DiceSet Response


| Field         | Type                                       | Label    | Description |
| ------------- | ------------------------------------------ | -------- | ----------- |
| room_key      | [string](#string)                          |          |             |
| turns         | [int32](#int32)                            |          |             |
| current_phase | [EventHookPhase](#ULZProto.EventHookPhase) |          |             |
| phase_ab      | [PlayerSide](#ULZProto.PlayerSide)         |          |             |
| atk_side      | [PlayerSide](#ULZProto.PlayerSide)         |          |             |
| atk_point     | [int32](#int32)                            |          |             |
| atk_skill_id  | [int32](#int32)                            | repeated |             |
| def_side      | [PlayerSide](#ULZProto.PlayerSide)         |          |             |
| def_point     | [int32](#int32)                            |          |             |
| def_skill_id  | [int32](#int32)                            | repeated |             |

-------------------------------------------------------------

## Common Phase Handle



<a name="ULZProto.GDPhaseConfirmReq"></a>

### GDPhaseConfirmReq


| Field         | Type                                       | Label | Description |
| ------------- | ------------------------------------------ | ----- | ----------- |
| room_key      | [string](#string)                          |       |             |
| side          | [PlayerSide](#ULZProto.PlayerSide)         |       |             |
| current_phase | [EventHookPhase](#ULZProto.EventHookPhase) |       |             |
| phase_hook    | [EventHookType](#ULZProto.EventHookType)   |       |             |






<a name="ULZProto.GDPhaseConfirmResp"></a>

### GDPhaseConfirmResp


| Field    | Type              | Label | Description          |
| -------- | ----------------- | ----- | -------------------- |
| room_key | [string](#string) |       | repeated PhaseEffect |




-------------------------------------------------------------


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


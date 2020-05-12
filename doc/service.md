# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Protocol Documentation](#protocol-documentation)
  - [Table of Contents](#table-of-contents)
  - [service.proto](#serviceproto)
    - [GameDuelService](#gameduelservice)
  - [Scalar Value Types](#scalar-value-types)
  
- [Scalar Value Types](#scalar-value-types)



<a name="service.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## service.proto


 

 

 


<a name="ULZProto.GameDuelService"></a>

### GameDuelService
ANCHOR: service-func for GameDuelService
SECTION: service.proto

Basic Server Function

| Method Name        | Request Type                                       | Response Type                                      | Description                                                                                                                                                                                                                                                                                                   |
| ------------------ | -------------------------------------------------- | -------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CreateGame         | [GDCreateReq](#ULZProto.GDCreateReq)               | [GameDataSet](#ULZProto.GameDataSet)               | rpc ServerBroadcast (GDGetInfoReq) returns (stream GDBroadcastResp);                                                                                                                                                                                                                                          |
| GetGameData        | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [GameDataSet](#ULZProto.GameDataSet)               |                                                                                                                                                                                                                                                                                                               |
| QuitGame           | [GDCreateReq](#ULZProto.GDCreateReq)               | [Empty](#ULZProto.Empty)                           |                                                                                                                                                                                                                                                                                                               |
| InstSetEventCard   | [GDInstanceDT](#ULZProto.GDInstanceDT)             | [Empty](#ULZProto.Empty)                           | GameSet Logic Function instance card move                                                                                                                                                                                                                                                                     |
| DrawPhaseConfirm   | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [Empty](#ULZProto.Empty)                           | Draw-phase : confirm NOTE: After Broadcast Send &lt;[ refill_action_card_phase ]&gt;, Client set the event-card by &lt; InstSetEventCard &gt; { from their deck to own hand }; then, send this &lt; DrawPhaseConfirm &gt; to notify the server that `client ready to start next phase [move_card_drop_phase]` |
| MovePhaseConfirm   | [GDMoveConfirmReq](#ULZProto.GDMoveConfirmReq)     | [Empty](#ULZProto.Empty)                           | Move-phase : confirm NOTE: After Broadcast Send &lt;[ move_card_drop_phase ]&gt;, Client set the event-card by &lt; InstSetEventCard &gt; { from their own hand to out-side }; Then send this &lt; MovePhaseConfirm &gt; to notify the server that `client ready to start next phase [determine_move_phase]`  |
| MovePhaseResult    | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [GDMoveConfirmResp](#ULZProto.GDMoveConfirmResp)   |                                                                                                                                                                                                                                                                                                               |
| ADPhaseConfirm     | [GDADConfirmReq](#ULZProto.GDADConfirmReq)         | [Empty](#ULZProto.Empty)                           |                                                                                                                                                                                                                                                                                                               |
| ADPhaseResult      | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [GDADResultResp](#ULZProto.GDADResultResp)         |                                                                                                                                                                                                                                                                                                               |
| ADPhaseDiceResult  | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [GDADDiceResult](#ULZProto.GDADDiceResult)         |                                                                                                                                                                                                                                                                                                               |
| ChangePhaseConfirm | [GDChangeConfirmReq](#ULZProto.GDChangeConfirmReq) | [Empty](#ULZProto.Empty)                           | ChangeCharaPhase : Confirm and Result FIXME : 3v3 may need it, but 1v1 is not implement;                                                                                                                                                                                                                      |
| ChangePhaseResult  | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [Empty](#ULZProto.Empty)                           | ChangeCharaPhase : Confirm and Result FIXME : 3v3 may need it, but 1v1 is not implement;                                                                                                                                                                                                                      |
| EventPhaseConfirm  | [GDPhaseConfirmReq](#ULZProto.GDPhaseConfirmReq)   | [Empty](#ULZProto.Empty)                           | Event-Phase : Confirm NOTE: Once the Server send any phase notify the client may send feedback to server that ready for phase                                                                                                                                                                                 |
| EventPhaseResult   | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [GDPhaseConfirmResp](#ULZProto.GDPhaseConfirmResp) | Event-Phase : Confirm NOTE: Once the Server send any phase notify the client may send feedback to server that ready for phase                                                                                                                                                                                 |

 



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


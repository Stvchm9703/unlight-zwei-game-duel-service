# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [Protocol Documentation](#protocol-documentation)
  - [Table of Contents](#table-of-contents)
  - [service.proto](#serviceproto)
    - [GameDuelService](#gameduelservice)
    - [General function](#general-function)
      - [CreateGame](#creategame)
        - [Work Flow](#work-flow)
      - [GetGameData](#getgamedata)
        - [Work Flow](#work-flow-1)
      - [QuitGame](#quitgame)
        - [Work Flow](#work-flow-2)
      - [InstSetEventCard](#instseteventcard)
    - [Event Phase : Common handle](#event-phase--common-handle)
      - [EventPhaseConfirm](#eventphaseconfirm)
      - [EventPhaseResult](#eventphaseresult)
      - [work flow](#work-flow-3)
    - [Draw Phase](#draw-phase)
      - [DrawPhaseConfirm](#drawphaseconfirm)
      - [Work Flow](#work-flow-4)
    - [Move PHase](#move-phase)
      - [MovePhaseConfirm](#movephaseconfirm)
      - [MovePhaseResult](#movephaseresult)
      - [Work Flow](#work-flow-5)
    - [Attack / Defence Phase](#attack--defence-phase)
      - [ADPhaseConfirm](#adphaseconfirm)
        - [Work Flow for Attacker (Attack-Card-Drop-Phase)](#work-flow-for-attacker-attack-card-drop-phase)
        - [Work Flow for Defence (Defence-Card-Drop-Phase)](#work-flow-for-defence-defence-card-drop-phase)
      - [ADPhaseResult](#adphaseresult)
        - [Work Flow](#work-flow-6)
      - [ADPhaseDiceResult](#adphasediceresult)
    - [Change Phase](#change-phase)
      - [ChangePhaseConfirm](#changephaseconfirm)
      - [ChangePhaseResult](#changephaseresult)
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

| Method Name        | Request Type                                       | Response Type                                      | Description                                                                                                                                                                                                                                                                                           |
| ------------------ | -------------------------------------------------- | -------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| CreateGame         | [GDCreateReq](#ULZProto.GDCreateReq)               | [GameDataSet](#ULZProto.GameDataSet)               |                                                                                                                                                                                                                                                                                                       |
| GetGameData        | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [GameDataSet](#ULZProto.GameDataSet)               |                                                                                                                                                                                                                                                                                                       |
| QuitGame           | [GDCreateReq](#ULZProto.GDCreateReq)               | [Empty](#ULZProto.Empty)                           |                                                                                                                                                                                                                                                                                                       |
| InstSetEventCard   | [GDInstanceDT](#ULZProto.GDInstanceDT)             | [Empty](#ULZProto.Empty)                           | GameSet Logic Function instance card move                                                                                                                                                                                                                                                             |
| DrawPhaseConfirm   | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [Empty](#ULZProto.Empty)                           | Draw-phase : confirm  NOTE: After Broadcast Send \< [ refill_action_card_phase ]\>, Client set the event-card by \<  InstSetEventCard \> { from their deck to own hand }; then, send this \<  DrawPhaseConfirm \> to notify the server that `client ready to start next phase [move_card_drop_phase]` |
| MovePhaseConfirm   | [GDMoveConfirmReq](#ULZProto.GDMoveConfirmReq)     | [Empty](#ULZProto.Empty)                           | Move-phase : confirm NOTE: After Broadcast Send \< [ move_card_drop_phase ]\>, Client set the event-card by \<  InstSetEventCard \> { from their own hand to out-side }; Then send this \<  MovePhaseConfirm \> to notify the server that `client ready to start next phase [determine_move_phase]`   |
| MovePhaseResult    | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [GDMoveConfirmResp](#ULZProto.GDMoveConfirmResp)   |                                                                                                                                                                                                                                                                                                       |
| ADPhaseConfirm     | [GDADConfirmReq](#ULZProto.GDADConfirmReq)         | [Empty](#ULZProto.Empty)                           |                                                                                                                                                                                                                                                                                                       |
| ADPhaseResult      | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [GDADResultResp](#ULZProto.GDADResultResp)         |                                                                                                                                                                                                                                                                                                       |
| ADPhaseDiceResult  | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [GDADDiceResult](#ULZProto.GDADDiceResult)         |                                                                                                                                                                                                                                                                                                       |
| ChangePhaseConfirm | [GDChangeConfirmReq](#ULZProto.GDChangeConfirmReq) | [Empty](#ULZProto.Empty)                           | ChangeCharaPhase : Confirm and Result FIXME : 3v3 may need it, but 1v1 is not implement;                                                                                                                                                                                                              |
| ChangePhaseResult  | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [Empty](#ULZProto.Empty)                           | ChangeCharaPhase : Confirm and Result FIXME : 3v3 may need it, but 1v1 is not implement;                                                                                                                                                                                                              |
| EventPhaseConfirm  | [GDPhaseConfirmReq](#ULZProto.GDPhaseConfirmReq)   | [Empty](#ULZProto.Empty)                           | Event-Phase : Confirm NOTE: Once the Server send any phase notify the client may send feedback to server that ready for phase                                                                                                                                                                         |
| EventPhaseResult   | [GDGetInfoReq](#ULZProto.GDGetInfoReq)             | [GDPhaseConfirmResp](#ULZProto.GDPhaseConfirmResp) | Event-Phase : Confirm NOTE: Once the Server send any phase notify the client may send feedback to server that ready for phase                                                                                                                                                                         |

---

### General function 

#### CreateGame         
create gameset data to Redis 
- Request : [GDCreateReq](#ULZProto.GDCreateReq)               
- Response : [GameDataSet](#ULZProto.GameDataSet)               

##### Work Flow
remark: the gameset not pack `HostEventCardDeck` , `DuelEventCardDeck` , `PhaseSnapMod`, `EffectNodeSnapMod`  into same Redis store position, since it will grow larger and usually get/set during the game.

```mermaid 
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant GDS as Game Duel Service
participant Rds as Redis

  GC -->> GDS: CreateGame ( GDCreateReq )

  Note over GC,Rds : check Room is Exist 
  GDS -->> Rds: Get RoomService-Room by Room key
  Rds -->> GDS: return room 
  opt Redis execution error
    GDS -->> GC: return Error (code = Not Found) 
  end

  Note over GC,Rds : Create the Data set
  GDS -->> GDS: Create GameDataSet
  GDS -->> Rds: Set the GameSet to Redis
  opt Redis execution error
    GDS -->> GC: return Error (code = Not Found) 
  end
  

  Note over GC,Rds : Create the other related data set
  GDS -->> GDS: Generate Host Event-Card Deck
  GDS -->> GDS: Generate Duel Event-Card Deck

  par: Host Event Card Deck Dataset  
    GDS -->> Rds: Set HostEventCardDeck (EventCardList) into Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and: Duel Event Card Deck Dataset 
    GDS -->> Rds: Set DuelEventCardDeck (EventCardList) into Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and: PhaseSnapMod Dataset 
    GDS -->> Rds: Set Phase Snap Mod into Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end
  
  and: EffectNodeMod Dataset 
    GDS -->> Rds: Set Effect Node Mod into Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  end

  Note over GC,Rds : Return the new gameset dataset
  GDS ->> GC: return GameDataSet  

  GDS -->> OGC: broadcast: "Gameset ready"

```




#### GetGameData        
get the gameset data from Redis
- Request : [GDGetInfoReq](#ULZProto.GDGetInfoReq)             
- Response : [GameDataSet](#ULZProto.GameDataSet)               


##### Work Flow


```mermaid 
sequenceDiagram 
%% participant OGC as Other Game Client
participant GC as Game Client
participant GDS as Game Duel Service
participant Rds as Redis

  GC -->> GDS: CreateGame ( GDCreateReq )

  Note over GC,Rds : Get the Data set if exist
  GDS -->> Rds: Get the GameSet to Redis
  opt Redis execution error
    GDS -->> GC: return Error (code = Not Found) 
  end

  Note over GC,Rds : Get the other Data set 
  par: Host Event Card Deck Dataset  
    GDS -->> Rds: Get HostEventCardDeck (EventCardList) into Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and: Duel Event Card Deck Dataset 
    GDS -->> Rds: Get DuelEventCardDeck (EventCardList) into Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and: PhaseSnapMod Dataset 
    GDS -->> Rds: Get Phase Snap Mod into Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end
  
  and: EffectNodeMod Dataset 
    GDS -->> Rds: Get Effect Node Mod into Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  end

  Note over GC,Rds : Pack into Gameset, return the gameset dataset
  GDS ->> GC: return GameDataSet  

```

#### QuitGame           
end the game, to clear the game data from Redis, broadcast the message if the game is sudden quit game. 
- Request : [GDCreateReq](#ULZProto.GDCreateReq)               
- Response : [Empty](#ULZProto.Empty)                           
##### Work Flow

```mermaid 
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant GDS as Game Duel Service
participant Rds as Redis

  GC -->> GDS: QuitGame ( GDCreateReq )

  Note over GC,Rds : Get the Data set if exist
  GDS -->> Rds: Get the GameSet to Redis
  opt Redis execution error
    GDS -->> GC: return Error (code = Not Found) 
  end

  Note over GC,Rds : Get the PhaseSnapMod Data set 
  GDS -->> Rds: Get PhaseSnapMod into Redis
  opt Redis execution error
    GDS -->> GC: return Error (code = Not Found) 
  end

  GDS -->> GDS: Check the gameset is normal GameQuit?
  alt EventHookPhase:GameEnd
    GDS -->> Rds: Remove the related dataset 
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end
  end
  alt EventHookPhase: (Other Phase)
    Note over GC,Rds: raise "GiveUpMessage"
    GDS -->> GC: broadcast the "GiveUpMessage"
    GDS -->> GC: broadcast the "GiveUpMessage"
    GDS -->> GDS: Set as EventHookPhase:GameEnd
  end
  GDS ->> GC: return GameDataSet  

```
 
#### InstSetEventCard    
- Request : [GDInstanceDT](#ULZProto.GDInstanceDT)            
- Response:  [Empty](#ULZProto.Empty)                           
instance event card move 
Remark: client should reduce the frequence of requesting this 
```mermaid 
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant GDS as Game Duel Service
participant Rds as Redis

  GC -->> GDS: InstSetEventCard ( GDInstanceDT )

  Note over GC,Rds : Set the income message to Redis
  GDS -->> Rds: Set the income message to Redis

  Note over GC,Rds : forward instance Event-card message to other game client  
  GDS -->> OGC: Broadcast instance Event-card message 
  
  Note over GC,Rds : Get the Host/Duel EventCardDeck 
  GDS -->> Rds: Get Host/Duel EventCardDeck into Redis
  opt Redis execution error
    GDS -->> GC: return Error (code = Not Found) 
  end

  Note over GC,Rds : Edit the Host/Duel EventCardDeck's card as income EventCard 
  GDS -->> GDS: update the eventcard data in Host/Duel-EventCardDeck
  GDS -->> Rds: Set Host/Duel EventCardDeck into Redis
  opt Redis execution error
    GDS -->> GC: return Error (code = InternalError) 
  end

  GDS -->> GC: return Complete Status (Empty)  

```
### Event Phase : Common handle 


#### EventPhaseConfirm  
- Request : [GDPhaseConfirmReq](#ULZProto.GDPhaseConfirmReq)   
- Response : [Empty](#ULZProto.Empty)                           

  Event-Phase : Confirm 
  <!-- NOTE: Once the Server send any phase notify the client may send feedback to server that ready for phase  -->
  Current no implement for it 


#### EventPhaseResult   
- Request : [GDGetInfoReq](#ULZProto.GDGetInfoReq)             
- Response : [GDPhaseConfirmResp](#ULZProto.GDPhaseConfirmResp) 

Event-Phase : Confirm 

NOTE: Once the Server send any phase notify the client may send feedback to server that ready for phase 

```c
// ----- any phase with hook-type : before . after
[any_phase:before] {
  executeEffectNode()
}

[any_phase:after] {
  executeEffectNode()
}


// ---- start_turn_phase
[start_turn_phase:proxy] {
  1. turn ++

}
// -------------------------------------
// draw-phase
// ---- refill_action_card_phase
[refill_action_card_phase:proxy] {

}

// ---- determine_battle_point_phase
[determine_battle_point_phase:proxy] {
  1. skill calculation
}

// ----- battle_result_phase
[battle_result_phase : proxy] {
  1. dice-roll from sub-client
  2. store dice-roll first-result
}
// ----- damage_phase -----
[damage_phase : proxy] {

}
// ----- dead_chara_change_phase -----
[dead_chara_change_phase : proxy] {

}
// ----- determine_dead_chara_change_phase -----
[determine_dead_chara_change_phase : proxy] {

}

```

#### work flow 

```mermaid
sequenceDiagram 
participant GC as Game Client
participant GDS as Game Duel Service
%%participant Skl as Skill 
participant Rds as Redis

  GC -->> GDS: InstSetEventCard ( GDGetInfoReq )

  Note over GC,Rds : Get the PhaseSnapMod 
  GDS -->> Rds: Get PhaseSnapMod into Redis
  opt Redis execution error
    GDS -->> GC: return Error (code = Not Found) 
  end

  Note over GC,Rds : check it is vaid 
  GDS -->> GDS: check EventHookPhase & EventHookType is same 
  opt Redis execution error
    GDS -->> GC: return Error (code = Invalid Arg) 
  end

  GDS -->> GDS: Set isReadyFlag by the request plasyer side 

  Note over GC,Rds : Set the Updated PhaseSnapMod EventCardDeck 
  GDS -->> Rds: Set Updated PhaseSnapMod EventCardDeck into Redis
  opt Redis execution error
    GDS -->> GC: return Error (code = Internal) 
  end

  alt if Both side ready 
    Note over GC,Rds : Both side ready, then move next phase in new routine process 
    par : get required game set data from Redis
      GDS->>Rds: Get GameSet Data Redis
    and : get effect-node-snap-mod data from Redis
      GDS->>Rds: Get EffectNodeSnapMod Data Redis
    end

    GDS -->> GDS : run moveNextPhase()
  end

  Note over GC,Rds : Finally, return Complete Status

  GDS -->> GC: return Complete Status (GDPhaseConformResp)  

```

---
### Draw Phase

Draw-phase : confirm  
NOTE: After Broadcast Send 
[refill_action_card_phase], Client set the event-card by \<InstSetEventCard\> { from their deck to own hand }; 
then, send this \<DrawPhaseConfirm\> to notify the server that `client ready to start next phase [move_card_drop_phase]` 



#### DrawPhaseConfirm   
- Request : [GDGetInfoReq](#ULZProto.GDGetInfoReq)             
- Response: [Empty](#ULZProto.Empty)                          

this comfirm request is for response the GameDuelService after GDS send the drawed cards to Host player and Duel player
Basically, it is similar with EventPhaseConfirm

#### Work Flow
```mermaid
sequenceDiagram 
participant GC as Game Client
participant GDS as Game Duel Service
%%participant Skl as Skill 
participant Rds as Redis

  GC -->> GDS: InstSetEventCard ( GDGetInfoReq )

  Note over GC,Rds : Get the PhaseSnapMod  
  GDS -->> Rds: Get PhaseSnapMod from Redis
  opt Redis execution error
    GDS -->> GC: return Error (code = Not Found) 
  end

  Note over GC,Rds : check it is vaid 
  GDS -->> GDS: check EventHookPhase & EventHookType is same 
  opt Redis execution error
    GDS -->> GC: return Error (code = Invalid Arg) 
  end

  GDS -->> GDS: Set isReadyFlag by the request plasyer side 

  Note over GC,Rds : Set the Updated PhaseSnapMod EventCardDeck 
  GDS -->> Rds: Set Updated PhaseSnapMod EventCardDeck into Redis
  opt Redis execution error
    GDS -->> GC: return Error (code = Internal) 
  end

  alt if Both side ready 
    Note over GC,Rds : Both side ready, then move next phase in new routine process 
    par : get required game set data from Redis
      GDS->>Rds: Get GameSet Data Redis
    and : get effect-node-snap-mod data from Redis
      GDS->>Rds: Get EffectNodeSnapMod Data Redis
    end

    GDS -->> GDS : run moveNextPhase()
  end

  Note over GC,Rds : Finally, return Complete Status

  GDS -->> GC: return Complete Status (GDPhaseConformResp)  
```




---
### Move PHase
 
Move-phase : confirm 
 
NOTE: After Broadcast Send \<[move_card_drop_phase]\>, Client set the event-card by \<InstSetEventCard\> { from their own hand to out-side }; Then send this \<MovePhaseConfirm\> to notify the server that `client ready to start next phase [determine_move_phase]`   



#### MovePhaseConfirm   
- Request: [GDMoveConfirmReq](#ULZProto.GDMoveConfirmReq)     
- Response: [Empty](#ULZProto.Empty)                           


#### MovePhaseResult    
- Request :  [GDGetInfoReq](#ULZProto.GDGetInfoReq)             
- Response: [GDMoveConfirmResp](#ULZProto.GDMoveConfirmResp)   


#### Work Flow
```mermaid
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant GDS as Game Duel Service
participant Rds as Redis
participant Skl as Skill Calculation Service 

  Note over GC,Skl : Suppose EventHookPhase is [move_card_drop_phase] 
  GC -->> GDS: MovePhaseConfirm ( GDMoveConfirmReq )

  Note over GC,Rds : Get the Required Dataset (with Async Waitgroup)
  par get GameSet Data 
    GDS -->> Rds: Get Gameset from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and get MovePhaseSnapMod
    GDS -->> Rds: Get MovePhaseSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and get the PhaseSnapMod  
    GDS -->> Rds: Get PhaseSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and get the EffectNodeSnapMod 
    GDS -->> Rds: Get EffectNodeSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end
  end
%%% ------------------------------------
  Note over GC,Rds : check it is vaid 
  GDS -->> GDS: check EventHookPhase & EventHookType is same 
  opt Redis execution error
    GDS -->> GC: return Error (code = Invalid Arg) 
  end
%%% ------------------------------------
  Note over GC,Rds : check the disable-skill is involved in the EffectNodeResult
  GDS -->> GDS: check disable-skill is involved
  alt is exist :
    GDS -->> GDS: set move-point as income net move-point 
  else not exist: 
    GDS -->> Skl: request Skill-Calculation
    opt Skill-Calculation execution error
      Skl -->> GDS: return Error (code = Invalid Arg) 
    end
    GDS -->> GDS: set the result value to move-point
    GDS -->> GDS: add the EffectNode if the Skill generate new EffectNode
  end
%%% ------------------------------------
  Note over GC,Rds : Calculate the EffectNode involved in move-drop-phase
  GDS -->> GDS: add move-point value from the EffectNode
%%% ------------------------------------
  Note over GC,Rds : update the ready flag
  GDS -->> GDS: Set isReadyFlag by the request plasyer side 
%%% ------------------------------------

  Note over GC,Rds : Set the updated Dataset (with Async Waitgroup)
  par Set GameSet Data 
    GDS -->> Rds: Set Gameset from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and set MovePhaseSnapMod
    GDS -->> Rds: set MovePhaseSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and set the PhaseSnapMod  
    GDS -->> Rds: set PhaseSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and set the EffectNodeSnapMod 
    GDS -->> Rds: set EffectNodeSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end
  end
%%% -------------------------------------
  Note over GC,Rds: Broadcast the "Host/Dueler Player is Ready"  
  GDS -->> GC: Broadcast the "Host/Dueler Player is Ready"
  GDS -->> OGC: Broadcast the "Host/Dueler Player is Ready"


%%% -------------------------------------
  alt if Both side ready 
    Note over GC,Rds : Both side ready, then move next phase in new routine process 
    Note over GC,Rds: Broadcast the "Both Side is Ready"  
    GDS -->> GC: Broadcast the "Host/Dueler Player is Ready"
    GDS -->> OGC: Broadcast the "Host/Dueler Player is Ready"
    GDS -->> GDS: run moveNextPhase( EventHookPhase:MoveDropPhase, EventHookType:After  )[see below]
  end

  Note over GC,Rds : Finally, return Complete Status
  GDS -->> GC: return Complete Status (Empty)  

```

```mermaid 
graph TB
  subgraph EventHookPhase
    id_1[MoveDropPhase,Proxy]
    id_2[MoveDropPhase,After]
    id_3[DetermineMovePhase,Before]
    id_4[DetermineMovePhase,Proxy]
    id_5[DetermineMovePhase,After]
    id_6[FinishMovePhase,Before]
    id_7[FinishMovePhase,Proxy]
    id_8[FinishMovePhase,After]
  end
  subgraph REquest

  hd_1[MovePhaseConfirm]
  hd_2[EventPhaseResult]
  hd_2_1[EventPhaseResult]
  hd_2_2[EventPhaseResult]
  hd_2_3[EventPhaseResult]
  hd_2_4[EventPhaseResult]
  hd_3[MovePhaseResult]
  end

  id_1 --> hd_1 --> id_2 --> hd_2 --> id_3  --> hd_2_1 -->id_4 --> hd_2_2 --> id_5 --> hd_2_3 --> id_6 -->hd_2_4 --> id_7 --> hd_3 --> id_8
```



---
### Attack / Defence Phase


#### ADPhaseConfirm     
- Request : [GDADConfirmReq](#ULZProto.GDADConfirmReq)         
- Response : [Empty](#ULZProto.Empty)                           


##### Work Flow for Attacker (Attack-Card-Drop-Phase) 
```mermaid
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant GDS as Game Duel Service
participant Rds as Redis
participant Skl as Skill Calculation Service 

  Note over GC,Skl : Suppose EventHookPhase is [attack_card_drop_phase] 
  GC -->> GDS: ADPhaseConfirm ( GDADConfirmReq )

  Note over GC,Rds : Get the Required Dataset (with Async Waitgroup)
  par get GameSet Data 
    GDS -->> Rds: Get Gameset from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and get ADPhaseSnapMod
    GDS -->> Rds: Get ADPhaseSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and get the PhaseSnapMod  
    GDS -->> Rds: Get PhaseSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and get the EffectNodeSnapMod 
    GDS -->> Rds: Get EffectNodeSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end
  end
%%% ------------------------------------
  Note over GC,Rds : check it is vaid 
  GDS -->> GDS: check EventHookPhase & EventHookType is same to PhaseSnapMod
  opt Redis execution error
    GDS -->> GC: return Error (code = Invalid Arg) 
  end

%%% ------------------------------------

  Note over GC,Rds: Attack Handle  
  GDS -->> GDS: set income data to ADPhaseSnapMod

  GDS -->> GDS: check disable-skill is involved
  alt is exist :
    GDS -->> GDS: set attack-point as income net attack-point 
  else not exist: 
    GDS -->> Skl: request Skill-Calculation
    opt Skill-Calculation execution error
      Skl -->> GDS: return Error (code = Invalid Arg) 
    end
    GDS -->> GDS: set the result value to move-point
    GDS -->> GDS: add the EffectNode if the Skill generate new EffectNode
  end
  %%% ------------------------------------
    Note over GC,Rds : Calculate the EffectNode involved in attack-card-drop-phase
    GDS -->> GDS: add attack-point value from the EffectNode
    GDS -->> GDS: check disable attack flag in EffectNode
  %%% ------------------------------------
  
    Note over GC,Rds : Set the updated Dataset (with Async Waitgroup)
    par Set GameSet Data 
      GDS -->> Rds: Set Gameset from Redis
      opt Redis execution error
        GDS -->> GC: return Error (code = Not Found) 
      end

    and set ADPhaseSnapMod
      GDS -->> Rds: set ADPhaseSnapMod from Redis
      opt Redis execution error
        GDS -->> GC: return Error (code = Not Found) 
      end

    and set the PhaseSnapMod  
      GDS -->> Rds: set PhaseSnapMod from Redis
      opt Redis execution error
        GDS -->> GC: return Error (code = Not Found) 
      end

    and set the EffectNodeSnapMod 
      GDS -->> Rds: set EffectNodeSnapMod from Redis
      opt Redis execution error
        GDS -->> GC: return Error (code = Not Found) 
      end
    end
  %%% -------------------------------------
    Note over GC,Rds: Broadcast the "AD_PHASE:ATK_RESULT:%value"  
    GDS -->> GC: Broadcast the "AD_PHASE:ATK_RESULT:%value"
    GDS -->> OGC: Broadcast the "AD_PHASE:ATK_RESULT:%value"


  Note over GC,Rds : Finally, return Complete Status
  GDS -->> GC: return Complete Status (Empty)  to Game Client( Attacker )

  Note over GC,Skl : Then, Wait for both side client requesting [ADPhaseResult] as Acknowledge 
```
---
##### Work Flow for Defence (Defence-Card-Drop-Phase) 
```mermaid
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant GDS as Game Duel Service
participant Rds as Redis
participant Skl as Skill Calculation Service 

  Note over GC,Skl : Suppose EventHookPhase is [defence_card_drop_phase] 
  GC -->> GDS: ADPhaseConfirm ( GDADConfirmReq )

  Note over GC,Rds : Get the Required Dataset (with Async Waitgroup)
  par get GameSet Data 
    GDS -->> Rds: Get Gameset from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and get ADPhaseSnapMod
    GDS -->> Rds: Get ADPhaseSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and get the PhaseSnapMod  
    GDS -->> Rds: Get PhaseSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and get the EffectNodeSnapMod 
    GDS -->> Rds: Get EffectNodeSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end
  end
%%% ------------------------------------
  Note over GC,Rds : check it is vaid 
  GDS -->> GDS: check EventHookPhase & EventHookType is same to PhaseSnapMod
  opt Redis execution error
    GDS -->> GC: return Error (code = Invalid Arg) 
  end

%%% ------------------------------------

  Note over GC,Rds: Defence Handle  
  GDS -->> GDS: set income data to ADPhaseSnapMod

  GDS -->> GDS: check disable-skill is involved
  alt is exist :
    GDS -->> GDS: set def-point as income net def-point 
  else not exist: 
    GDS -->> Skl: request Skill-Calculation
    opt Skill-Calculation execution error
      Skl -->> GDS: return Error (code = Invalid Arg) 
    end
    GDS -->> GDS: set the result value to move-point
    GDS -->> GDS: add the EffectNode if the Skill generate new EffectNode
  end
  %%% ------------------------------------
    Note over GC,Rds : Calculate the EffectNode involved in defence-card-drop-phase
    GDS -->> GDS: add defence-point value from the EffectNode
    GDS -->> GDS: check disable defence flag in EffectNode
  %%% ------------------------------------
  
    Note over GC,Rds : Set the updated Dataset (with Async Waitgroup)
    par Set GameSet Data 
      GDS -->> Rds: Set Gameset from Redis
      opt Redis execution error
        GDS -->> GC: return Error (code = Not Found) 
      end

    and set ADPhaseSnapMod
      GDS -->> Rds: set ADPhaseSnapMod from Redis
      opt Redis execution error
        GDS -->> GC: return Error (code = Not Found) 
      end

    and set the PhaseSnapMod  
      GDS -->> Rds: set PhaseSnapMod from Redis
      opt Redis execution error
        GDS -->> GC: return Error (code = Not Found) 
      end

    and set the EffectNodeSnapMod 
      GDS -->> Rds: set EffectNodeSnapMod from Redis
      opt Redis execution error
        GDS -->> GC: return Error (code = Not Found) 
      end
    end
  %%% -------------------------------------
    Note over GC,Rds: Broadcast the "AD_PHASE:DEF_RESULT:%value"  
    GDS -->> GC: Broadcast the "AD_PHASE:DEF_RESULT:%value"
    GDS -->> OGC: Broadcast the "AD_PHASE:DEF_RESULT:%value"

  Note over GC,Rds : Finally, return Complete Status
  GDS -->> GC: return Complete Status (Empty) to Game Client( Defencer )

  Note over GC,Skl : Then, Wait for both side client requesting [ADPhaseResult] as Acknowledge 
```

#### ADPhaseResult      
- Request : [GDGetInfoReq](#ULZProto.GDGetInfoReq)             
- Response : [GDADResultResp](#ULZProto.GDADResultResp)         

##### Work Flow
```mermaid
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant GDS as Game Duel Service
participant Rds as Redis
participant Skl as Skill Calculation Service 

  Note over GC,Skl : Suppose EventHookPhase is [***_card_drop_phase] 
  GC -->> GDS: ADPhaseResult ( GDGetInfoReq )

  Note over GC,Rds : Get the Required Dataset (with Async Waitgroup)
  par get ADPhaseSnapMod
    GDS -->> Rds: Get ADPhaseSnapMod from Redis
    %% activate Rds
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and get the PhaseSnapMod  
    GDS -->> Rds: Get PhaseSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end
  end
%%% ------------------------------------
  Note over GC,Rds : check it is vaid 
  GDS -->> GDS: check EventHookPhase & EventHookType is same to PhaseSnapMod
  opt Redis execution error
    GDS -->> GC: return Error (code = Invalid Arg) 
  end
%%% ------------------------------------
  Note over GC,Rds : update the ready flag
    GDS -->> Rds: Update PhaseSnapMod
    opt Redis execution error
      GDS -->> GC: return Error (code = Internal) 
    end
%%% ------------------------------------
    GDS -->> GDS: Setup return message

%%% -------------------------------------
  alt if Both side ready 
    Note over GC,Rds: Broadcast the "Both Side is Ready"  
    GDS -->> GC: Broadcast the "AD_PHASE:ACK_Both_SideResolve:"
    GDS -->> OGC: Broadcast the "AD_PHASE:ACK_Both_SideResolve:"

    Note over GC,Rds : Both side ready, then move next phase in new routine process 
    GDS -->> GDS: run moveNextPhase( EventHookPhase:MoveDropPhase, EventHookType:After  )[see below]
  end

  Note over GC,Rds : Finally, return packed message [ ADResultResp ]
  GDS -->> GC: return Complete Status (Empty) to Game Client( Defencer )
```





#### ADPhaseDiceResult  
- Request : [GDGetInfoReq](#ULZProto.GDGetInfoReq)             
- Response : [GDADDiceResult](#ULZProto.GDADDiceResult)         

Important: Code is not implement

```mermaid
sequenceDiagram 
%% participant OGC as Other Game Client
participant GC as Game Client
participant GDS as Game Duel Service
participant Rds as Redis
participant Skl as Skill Calculation Service 


  Note over GC,Skl: from MoveNextPhase [DeferenceCardDropPhase] -> [DetermineBattlePointPhase]
  Note over GC,Skl : Suppose EventHookPhase is [determine_battle_point_phase] 

  Note over GC,Rds : Get the Required Dataset (with Async Waitgroup)
  par update HostEventCardDeck Move {Outside} to {Destroy} 
    GDS -->> Rds: update HostEventCardDeck from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and update DuelEventCardDeck Move {Outside} to {Destroy}
    GDS -->> Rds: Update DuelEventCardDeck from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end
  end
%%% ------------------------------------
  Note over GC,Rds: Broadcast the "DetermineBattlePointPhase:Complete"  
  GDS -->> GC: Broadcast the "DetermineBattlePointPhase:Complete"
  %%GDS -->> OGC: Broadcast the "DetermineBattlePointPhase:Complete"
%%% ------------------------------------
  Note over GC,Skl : Wait for EventPhaseResult as Acknowledge 

```
```mermaid
sequenceDiagram 
%% participant OGC as Other Game Client
participant GC as Game Client
participant GDS as Game Duel Service
participant Rds as Redis
participant Skl as Skill Calculation Service 


  Note over GC,Skl: from MoveNextPhase  [DetermineBattlePointPhase] -> [BattleResultPhase]
  Note over GC,Skl : Suppose EventHookPhase is [determine_battle_point_phase] 

  Note over GC,Rds : Get the Required Dataset (with Async Waitgroup)
  par Dice Calculate for attacker  
    GDS -->> Skl: request DiceCalculate for Attack
  and Dice Calculate for defencer
    GDS -->> Skl: request DiceCalculate for Defence
  end

  Note over GC,Rds : Store Dice Result 
    GDS-->>Rds: Set GDADDiceResult to Redis

%%% ------------------------------------
  Note over GC,Rds: Broadcast the "DiceResult:%value"  
  GDS -->> GC: Broadcast the "DiceResult:%value"
%%% ------------------------------------
  Note over GC,Skl : Wait for ADDiceResult as Acknowledge 

```

```mermaid
sequenceDiagram 
participant OGC as Other Game Client
participant GC as Game Client
participant GDS as Game Duel Service
participant Rds as Redis
participant Skl as Skill Calculation Service 

  Note over GC,Skl : Suppose EventHookPhase is [battle_result_phase] 
  GC -->> GDS: ADPhaseDiceResult ( GDGetInfoReq )

  Note over GC,Rds : Get the Required Dataset (with Async Waitgroup)
  par get ADDiceResult
    GDS -->> Rds: Get ADDiceResult
    %% activate Rds
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end

  and get the PhaseSnapMod  
    GDS -->> Rds: Get PhaseSnapMod from Redis
    opt Redis execution error
      GDS -->> GC: return Error (code = Not Found) 
    end
  end
%%% ------------------------------------
  Note over GC,Rds : check it is vaid 
  GDS -->> GDS: check EventHookPhase & EventHookType is same to PhaseSnapMod
  opt Redis execution error
    GDS -->> GC: return Error (code = Invalid Arg) 
  end
%%% ------------------------------------
  Note over GC,Rds : update the ready flag
    GDS -->> Rds: Update PhaseSnapMod
    opt Redis execution error
      GDS -->> GC: return Error (code = Internal) 
    end
%%% ------------------------------------
    GDS -->> GC: return ADDiceResult

%%% -------------------------------------
  alt if Both side ready 
    Note over GC,Rds: Broadcast the "Both Side is Ready"  
    GDS -->> GC: Broadcast the "DiceResult:ACK_Both_SideResolve:"
    GDS -->> OGC: Broadcast the "DiceResult:ACK_Both_SideResolve:"

    Note over GC,Rds : Both side ready, then move next phase in new routine process 
    GDS -->> GDS: run moveNextPhase( EventHookPhase:battle_result_phase, EventHookType:After  )[see below]
  end

  Note over GC,Rds : Finally, return packed message [ ADResultResp ]
  
```



---
### Change Phase

#### ChangePhaseConfirm 
- Request : [GDChangeConfirmReq](#ULZProto.GDChangeConfirmReq) 
- Response : [Empty](#ULZProto.Empty)                           

ChangeCharaPhase : Confirm and Result 

FIXME : 3v3 may need it, but 1v1 is not implement;

#### ChangePhaseResult  
- Request : [GDGetInfoReq](#ULZProto.GDGetInfoReq)             
- Response : [Empty](#ULZProto.Empty)                           

ChangeCharaPhase : Confirm and Result 

FIXME : 3v3 may need it, but 1v1 is not  implement;

---


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


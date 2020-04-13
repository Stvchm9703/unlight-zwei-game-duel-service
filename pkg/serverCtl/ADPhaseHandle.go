package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"fmt"
	"log"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) attackPhaseHandle(roomKey *string, snapMod *pb.ADPhaseSnapMod, phaseMod *pb.PhaseSnapMod) {
	// do effect calculate
	// SECTION: skill-calculation
	result := 0
	// do update
	wkbox := this.searchAliveClient()
	snapMod.AttackVal = int32(result)
	snapMod.IsProcessed = true
	var snapModkey = *roomKey + ":ADPhMod"
	if _, err := (wkbox).SetPara(&snapModkey, snapMod); err != nil {
		log.Println(err)
	}
	// send ok message
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      *roomKey,
		Msg:          fmt.Sprintf("AD_PHASE:ATK_RESULT:", result),
		Command:      pb.CastCmd_GET_ATK_PHASE_RESULT,
		Side:         snapMod.CurrAttacker,
		CurrentPhase: pb.EventHookPhase_attack_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})

}

func (this *ULZGameDuelServiceBackend) defencePhaseHandle(roomKey *string, snapMod *pb.ADPhaseSnapMod, phaseMod *pb.PhaseSnapMod) {
	// do effect calculate
	// SECTION: skill-calculation
	result := 0
	// do update
	wkbox := this.searchAliveClient()
	snapMod.AttackVal = int32(result)
	var snapModkey = *roomKey + ":ADPhMod"
	if _, err := (wkbox).SetPara(&snapModkey, snapMod); err != nil {
		log.Println(err)
	}
	// send ok message
	side := snapMod.CurrAttacker
	if side == pb.PlayerSide_HOST {
		side = pb.PlayerSide_DUELER
	} else {
		side = pb.PlayerSide_HOST
	}
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      *roomKey,
		Msg:          fmt.Sprintf("AD_PHASE:DEF_RESULT:", result),
		Command:      pb.CastCmd_GET_DEF_PHASE_RESULT,
		Side:         side,
		CurrentPhase: pb.EventHookPhase_defence_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})
}

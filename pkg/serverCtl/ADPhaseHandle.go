package serverCtl

import (
	pb "ULZGameDuelService/proto"
	// svcpb "ULZGameDuelService/pkg/scriptRunner/proto"
	"fmt"
	"log"
	"sync"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) attackPhaseHandle(
	roomKey *string,
	gameSet *pb.GameDataSet,
	snapMod *pb.ADPhaseSnapMod,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
) {
	// do effect calculate
	// SECTION: skill-calculation
	wg := sync.WaitGroup{}
	wkbox := this.searchAliveClient()
	errch := make(chan error)
	wg.Add(1)
	var val *int32
	var eff []*pb.EffectResult
	go func() {
		var err error
		val, eff, err = this.skillClient.SkillCalculateWrap(
			snapMod.AttackCard,
			snapMod.AttackTrigSkl,
			"attack",
		)
		if err != nil {
			errch <- err
			return
		}
		wg.Done()
	}()
	wg.Wait()

	if errt := <-errch; errt != nil {
		fmt.Printf("attack-phase-handle::%v", errt)
		return
	}
	// do update

	snapMod.AttackVal = *val
	snapMod.IsProcessed = true
	go func() {
		var snapModkey = *roomKey + ":ADPhMod"
		if _, err := (wkbox).SetPara(&snapModkey, snapMod); err != nil {
			log.Println(err)
		}
	}()
	// send ok message
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      *roomKey,
		Msg:          fmt.Sprintf("AD_PHASE:ATK_RESULT:%s", *val),
		Command:      pb.CastCmd_GET_ATK_PHASE_RESULT,
		Side:         snapMod.CurrAttacker,
		CurrentPhase: pb.EventHookPhase_attack_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})

}

func (this *ULZGameDuelServiceBackend) defencePhaseHandle(
	roomKey *string,
	gameSet *pb.GameDataSet,
	snapMod *pb.ADPhaseSnapMod,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
) {
	// do effect calculate
	// SECTION: skill-calculation
	wg := sync.WaitGroup{}
	wkbox := this.searchAliveClient()
	errch := make(chan error)
	wg.Add(1)
	var val *int32
	var eff []*pb.EffectResult
	go func() {
		var err error
		val, eff, err = this.skillClient.SkillCalculateWrap(
			snapMod.DefenceCard,
			snapMod.DefenceTrigSkl,
			"defence",
		)
		if err != nil {
			errch <- err
			return
		}
		wg.Done()
	}()
	wg.Wait()
	if errt := <-errch; errt != nil {
		fmt.Printf("defence-phase-handle::%v", errt)
		return
	}

	// result := 0
	// do update
	snapMod.DefenceVal = *val
	snapMod.IsProcessed = true

	go func() {
		var snapModkey = *roomKey + snapMod.RdsKeyName()
		if _, err := (wkbox).SetPara(&snapModkey, snapMod); err != nil {
			log.Println(err)
		}
	}()
	// send ok message
	side := snapMod.CurrAttacker
	if side == pb.PlayerSide_HOST {
		side = pb.PlayerSide_DUELER
	} else {
		side = pb.PlayerSide_HOST
	}
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      *roomKey,
		Msg:          fmt.Sprintf("AD_PHASE:DEF_RESULT:%s", *val),
		Command:      pb.CastCmd_GET_DEF_PHASE_RESULT,
		Side:         side,
		CurrentPhase: pb.EventHookPhase_defence_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})
}

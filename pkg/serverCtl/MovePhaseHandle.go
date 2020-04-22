package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"fmt"
	"log"
	// cm "ULZGameDuelService/pkg/common"
	// "context"
	// "log"
	// "sync"
	// "time"
	// "github.com/gogo/status"
	// "google.golang.org/grpc/codes"
)

/**
 * movePhaseHandle :
 * 		for handle the determine_move_phase:proxy logic
 */
func (this *ULZGameDuelServiceBackend) movePhaseHandle(
	gameSet *pb.GameDataSet,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	moveMod *pb.MovePhaseSnapMod,
) {
	// go to request the move result
	// SECTION: skill-calculation

	// real act

	go func() {
		wkbox := this.searchAliveClient()
		snapModkey := gameSet.RoomKey + moveMod.RdsKeyName()
		if _, err := (wkbox).SetPara(&snapModkey, moveMod); err != nil {
			log.Println(err)
		}
		wkbox.Preserve(false)
	}()

	go func() {
		mbox := this.searchAliveClient()
		efKey := gameSet.RoomKey + effectMod.RdsKeyName()
		if _, err := (mbox).SetPara(&efKey, effectMod); err != nil {
			log.Println(err)
		}
		mbox.Preserve(false)
	}()
	// send ok message
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      gameSet.RoomKey,
		Msg:          fmt.Sprintf("MOVE:MOVE_RESULT:"),
		Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
		CurrentPhase: pb.EventHookPhase_attack_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})

}
func _effectCheckForMovePhase(gmSet *pb.GameDataSet, eflist *pb.EffectNodeSnapMod) {
	fmt.Println("Start Filter effect")
	/** TODO : concat the Effect node that for move-phase calculating
	 */
	// host current card
	var hostCurrentEF pb.EffectResult
	// duel current card
	var duelCurrentEF pb.EffectResult

	for _, v := range eflist.PendingEf {
		if v.TarSide == pb.PlayerSide_HOST &&
			v.TarCard == gmSet.HostCurrCardKey {
			hostCurrentEF.Mp += v.Mp
			if v.DisableMove {
				hostCurrentEF.DisableMove = v.DisableMove
			}
			if v.DisableChange {
				hostCurrentEF.DisableChange = v.DisableChange
			}
			if v.BindingFunc == "" {
				hostCurrentEF.BindingFunc += ";" + v.BindingFunc
			}
		} else if v.TarSide == pb.PlayerSide_DUELER &&
			v.TarCard == gmSet.HostCurrCardKey {
			duelCurrentEF.Mp += v.Mp
			if v.DisableMove {
				duelCurrentEF.DisableMove = v.DisableMove
			}
			if v.DisableChange {
				duelCurrentEF.DisableChange = v.DisableChange
			}
		}
	}
}

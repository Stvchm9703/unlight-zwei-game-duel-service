package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"fmt"
	"log"
	"sync"
	// cm "ULZGameDuelService/pkg/common"
	// "context"
	// "log"
	// "sync"
	// "time"
	// "github.com/gogo/status"
	// "google.golang.org/grpc/codes"
)

func (this *ULZGameDuelServiceBackend) movePhaseHandle(
	roomKey *string,
	gameSet *pb.GameDataSet,
	moveMod *pb.MovePhaseSnapMod,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
) {
	// go to request the move result
	// SECTION: skill-calculation
	// do effect calculate
	result := 0
	// do update

	var snapModkey = *roomKey + moveMod.RdsKeyName()

	wg := sync.WaitGroup{}
	errCh := make(chan error)
	// host
	var hostVal *int32
	var hostEff []*pb.EffectResult
	wg.Add(1)
	go func() {
		var err error
		hostVal, hostEff, err = this.skillClient.SkillCalculateWrap(moveMod.HostCard, moveMod.HostTrigSkl, "move")
		if err != nil {
			errCh <- err
			return
		}
		wg.Done()
	}()

	// duel
	var duelVal *int32
	var duelEff []*pb.EffectResult
	wg.Add(1)
	go func() {
		var err error
		duelVal, duelEff, err = this.skillClient.SkillCalculateWrap(moveMod.DuelCard, moveMod.DuelTrigSkl, "move")
		if err != nil {
			errCh <- err
			return
		}
		wg.Done()
	}()

	wg.Wait()

	if errt := <-errCh; errt != nil {
		fmt.Printf(" \n")
		return
	}

	effectMod.PendingEf = append(effectMod.PendingEf, hostEff...)
	effectMod.PendingEf = append(effectMod.PendingEf, duelEff...)

	// real act

	go func() {
		wkbox := this.searchAliveClient()
		if _, err := (wkbox).SetPara(&snapModkey, moveMod); err != nil {
			log.Println(err)
		}
		wkbox.Preserve(false)
	}()

	go func() {
		mbox := this.searchAliveClient()
		efKey := *roomKey + effectMod.RdsKeyName()
		if _, err := (mbox).SetPara(&efKey, effectMod); err != nil {
			log.Println(err)
		}
		mbox.Preserve(false)
	}()
	// send ok message
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      *roomKey,
		Msg:          fmt.Sprintf("MOVE:MOVE_RESULT:", result),
		Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
		CurrentPhase: pb.EventHookPhase_attack_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})

}
func _effectCheckForMovePhase(gmSet *pb.GameDataSet, eflist *pb.EffectNodeSnapMod) {
	fmt.Println("Start Filter effect")
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

func _moveResult(gmdata *pb.GameDataSet, mod *pb.MovePhaseSnapMod) {

}

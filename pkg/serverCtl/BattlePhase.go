package serverCtl

import (
	pb "ULZGameDuelService/proto"
	// "fmt"
	// "log"
	// "sync"
	// Static files
)

/** determineBattlePointPhase: event handle in determine-battle-point-phase
 * 		1. move the event card in evtcrdlst
 */
func (this *ULZGameDuelServiceBackend) determineBattlePointPhase(
	gameSet *pb.GameDataSet,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	adStoreMod *pb.ADPhaseSnapMod,
) {
	this.skillClient.EffectCalculateWrap(gameSet.RoomKey)

}

/**  battlePhaseHandle : event handle in battle-phase
 * 		1. dice the result set
 */
func (this *ULZGameDuelServiceBackend) battlePhaseHandle(
	gameSet *pb.GameDataSet,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	adStoreMod *pb.ADPhaseSnapMod,
) {

}

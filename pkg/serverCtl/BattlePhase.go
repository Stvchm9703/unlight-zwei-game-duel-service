package serverCtl

import (
	pb "ULZGameDuelService/proto"
	// "fmt"
	// "log"
	// "sync"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) battlePhaseHandle(
	gameSet *pb.GameDataSet,
	snapMod *pb.ADPhaseSnapMod,
	stateMod *pb.PhaseSnapMod,
	effectMod *pb.EffectNodeSnapMod,
	adStoreMod *pb.ADPhaseSnapMod,
) (bool, error) {

	return true, nil
}

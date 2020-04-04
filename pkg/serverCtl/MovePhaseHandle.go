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

func (this *ULZGameDuelServiceBackend) movePhaseHandle(roomKey *string, moveMod *pb.MovePhaseSnapMod, stateMod *pb.PhaseSnapMod) {
	// go to request the move result

	// do effect calculate
	result := 0
	// do update
	wkbox := this.searchAliveClient()
	var snapModkey = *roomKey + moveMod.RdsKeyName()
	if _, err := (wkbox).SetPara(&snapModkey, moveMod); err != nil {
		log.Println(err)
	}
	// send ok message
	go this.BroadCast(&pb.GDBroadcastResp{
		RoomKey:      *roomKey,
		Msg:          fmt.Sprintf("AD_PHASE:ATK_RESULT:", result),
		Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
		CurrentPhase: pb.EventHookPhase_attack_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
	})
}

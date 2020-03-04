package serverCtl

import (
	cm "ULZGameDuelService/pkg/common"
	pb "ULZGameDuelService/proto"
	"context"
	"log"
	"time"

	// "github.com/docker/docker/builder/builder-next/adapters/snapshot"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
	// Static files
	// _ "ULZGameDuelService/statik"
)

func (this *ULZGameDuelServiceBackend) MovePhaseConfirm(ctx context.Context, req *pb.GDMoveConfirmReq) (*pb.Empty, error) {
	cm.PrintReqLog(ctx, "Move-Phase-Confirm", req)
	start := time.Now()
	this.mu.Lock()
	wkbox := this.searchAliveClient()
	defer func() {
		wkbox.Preserve(false)
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Move-Phase-Confirm took %s", elapsed)
	}()

	var returner pb.GameDataSet
	if _, err := (wkbox).GetPara(&req.RoomKey, &returner); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	var snapMovekey = req.RoomKey + ":MvPhMod"
	var snapMove pb.MovePhaseSnapMod
	if _, err := (wkbox).GetPara(&snapMovekey, &snapMove); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if returner.EventPhase != pb.EventHookPhase_move_card_drop_phase {
		return nil, status.Error(codes.FailedPrecondition, "NOT_IN_MOVE_PHASE")
	}

	if (req.Side == pb.PlayerSide_HOST && snapMove.IsHostReady) ||
		(req.Side == pb.PlayerSide_DUELER && snapMove.IsDuelReady) {
		return nil, status.Error(codes.AlreadyExists, "ALREADY_READY")
	}

	if req.Side == pb.PlayerSide_HOST {
		snapMove.IsHostReady = true
		snapMove.HostVal = req.Point
		snapMove.HostOpt = req.MoveOpt
		snapMove.HostTrigSkl = req.TriggerSkl
		snapMove.HostCard = req.UpdateCard
		// snapMove.
		returner.IsHostReady = true

	} else if req.Side == pb.PlayerSide_DUELER {
		snapMove.IsDuelReady = true
		snapMove.DuelVal = req.Point
		snapMove.DuelOpt = req.MoveOpt
		snapMove.DuelTrigSkl = req.TriggerSkl
		snapMove.DuelCard = req.UpdateCard

		returner.IsDuelReady = true
	}
	// do snap-mod update
	if _, err := (wkbox).SetPara(&snapMovekey, snapMove); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	// broadcast ok ?
	side := req.Side.String()
	go this.BroadCast(&req.RoomKey, &side, &pb.GDBroadcastResp{
		RoomKey:      req.RoomKey,
		Msg:          req.Side.String() + "_READY",
		Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
		CurrentPhase: pb.EventHookPhase_move_card_drop_phase,
		PhaseHook:    pb.EventHookType_Proxy,
		Side:         req.Side,
		InstanceSet:  nil,
	})

	if snapMove.IsHostReady && snapMove.IsDuelReady {
		// broadcast for ready next phase
		returner.IsHostReady = true
		returner.IsDuelReady = true
		returner.EventPhase = pb.EventHookPhase_move_card_drop_phase
		returner.HookType = pb.EventHookType_After
		go this.BroadCast(&req.RoomKey, &side, &pb.GDBroadcastResp{
			RoomKey:      req.RoomKey,
			Msg:          "Both Ready",
			Command:      pb.CastCmd_GET_MOVE_PHASE_RESULT,
			CurrentPhase: pb.EventHookPhase_move_card_drop_phase,
			PhaseHook:    pb.EventHookType_After,
			InstanceSet:  nil,
		})
		// go store before move next
		if _, err := (wkbox).SetPara(&req.RoomKey, &returner); err != nil {
			log.Println(err)
			return nil, status.Errorf(codes.Internal, err.Error())
		}
		// both-ready; then move-next phase
		// go this.phase_run
	}
	return &pb.Empty{}, nil

}

/**
 * after the move-next to do [determine-move-phase]
 * it will request client to send move-phase-result to grep the result
 *
 * 	[move-phase]
 * 	|
 * 	{ both side ready }
 * 		| => ( go exec phase-runner )
 * 				| => [move_card_drop_phase:after] {
 * 				|		1. exec triggered skill
 * 				|			| -> get effect-func
 * 				|			| -> get instance value change
 * 				|		2. order effect-func in [ determine-move-phase:before,after ]
 * 				|		3. run the after-ef-result-node by order
 * 				|		4. return the list of node
 * 				|	}
 * 				|
 * 				|	phase-runner get af-ef-result-node list from exec-service:
 * 				|		-> send effent-phase-request
 * 				|
 * 				|	[determine-move-phase] is going end:{
 * 				|		1. store instance-result
 * 				|	 	2. send move-phase-result
 * 				|	}
 * 				|
 *  			| 	if the gameSet have dead result
 * 				|		if nvn == 1 :
 * 				|			game-end-phase
 * 				|		if nvn == 3 :
 * 				|			change-char-phase
 */

func (this *ULZGameDuelServiceBackend) MovePhaseResult(ctx context.Context, req *pb.GDGetInfoReq) (*pb.GDMoveConfirmResp, error) {
	cm.PrintReqLog(ctx, "Draw-Phase-Confirm", req)
	start := time.Now()
	this.mu.Lock()
	wkbox := this.searchAliveClient()
	defer func() {
		wkbox.Preserve(false)
		this.mu.Unlock()
		elapsed := time.Since(start)
		log.Printf("Draw-Phase-Confirm took %s", elapsed)
	}()

	var snapMovekey = req.RoomKey + ":MvPhModResult"
	var snapMove pb.GDMoveConfirmResp
	if _, err := (wkbox).GetPara(&snapMovekey, &snapMove); err != nil {
		log.Println(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return &snapMove, nil
}

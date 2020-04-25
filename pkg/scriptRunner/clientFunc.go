package scriptRunner

import (
	"context"
	"strings"
	"time"

	cm "ULZGameDuelService/pkg/common"
	Cf "ULZGameDuelService/pkg/config"

	pb "ULZGameDuelService/pkg/scriptRunner/proto"
	dataPb "ULZGameDuelService/proto"
	"flag"
	"fmt"
	"log"
	"sort"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// var (
// 	clientConn *grpc.ClientConn
// )

type SkillEffectSvcClient struct {
	clientName string
	conn       *grpc.ClientConn
	svcClient  pb.SkillEffectServiceClient
}

func ClientInit(conf *Cf.CfGrpcService) *SkillEffectSvcClient {
	flag.Parse()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.IP, conf.Port),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return &SkillEffectSvcClient{
		conn:      conn,
		svcClient: pb.NewSkillEffectServiceClient(conn),
	}
}

func ClientListInit(serviceName string, conf []Cf.CfGrpcService) *SkillEffectSvcClient {
	flag.Parse()

	var addrs []string
	for _, v := range conf {
		if v.ServiceName == serviceName {
			addrs = append(addrs, fmt.Sprintf("%s:%d", v.IP, v.Port))
		}
	}

	conn, err := grpc.Dial(
		serviceName,
		grpc.WithBalancer(
			grpc.RoundRobin(
				NewPseudoResolver(addrs),
			),
		), // This sets the initial balancing policy.
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}

	return &SkillEffectSvcClient{
		conn:      conn,
		svcClient: pb.NewSkillEffectServiceClient(conn),
	}
}
func (cli *SkillEffectSvcClient) ClientClose() error {
	fmt.Printf("client close:%s", time.Now())
	return cli.conn.Close()
}

func (cli *SkillEffectSvcClient) SkillInstCalc(req *pb.SESkillCalReq) (*pb.SESkillCalResp, error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cm.PrintReqLog(ctx, "Skill-Inst-Calc", req)
	defer func() {
		cancel()
		elapsed := time.Since(start)
		log.Printf("Draw-Phase-Confirm took %s", elapsed)
	}()
	var p peer.Peer
	res, err := cli.svcClient.SkillInstCalc(
		ctx, req,
		grpc.FailFast(false), // To wait a resolver returning addrs.
		grpc.Peer(&p),
	)
	if err != nil {
		fmt.Printf("Skill-Inst-Calc: cannot resolve:\n\t%v", err)
		return nil, err
	}
	fmt.Printf("Skill-Inst-Calc:resolve by %s\n", p.Addr)
	return res, nil
}
func (cli *SkillEffectSvcClient) SkillInstCalcWrap(
	incomeCard []*dataPb.EventCard,
	feat []*dataPb.SkillSet,
	requestPhase dataPb.EventCardType,
) (*int32, []*dataPb.EffectResult, error) {
	res, err := cli.SkillInstCalc(&pb.SESkillCalReq{
		IncomeCard: incomeCard,
		Feat:       feat,
		FromCli:    cli.clientName,
		TargType:   requestPhase,
	})
	return &res.ResultVal, res.EffectResult, err
}

// SkilCalculate:
func (cli *SkillEffectSvcClient) SkillCalculate(req *pb.SESkillCalReq) (*pb.SESkillCalResp, error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cm.PrintReqLog(ctx, "Skill-Calculate", req)
	defer func() {
		cancel()
		elapsed := time.Since(start)
		log.Printf("Draw-Phase-Confirm took %s", elapsed)
	}()
	var p peer.Peer

	res, err := cli.svcClient.SkillInstCalc(
		ctx, req,
		grpc.FailFast(false), // To wait a resolver returning addrs.
		grpc.Peer(&p),
	)
	if err != nil {
		fmt.Printf("Skill-Inst-Calc: cannot resolve:\n\t%v", err)
		return nil, err
	}
	fmt.Printf("Skill-Inst-Calc:resolve by %s\n", p.Addr)
	return res, nil
}
func (cli *SkillEffectSvcClient) SkillCalculateWrap(
	incomeCard []*dataPb.EventCard,
	feat []*dataPb.SkillSet,
	requestPhase dataPb.EventCardType,
) (*int32, []*dataPb.EffectResult, error) {
	res, err := cli.SkillCalculate(&pb.SESkillCalReq{
		IncomeCard: incomeCard,
		Feat:       feat,
		FromCli:    cli.clientName,
		TargType:   requestPhase,
	})
	return &res.ResultVal, res.EffectResult, err
}

// EffectCalculate: req-func
func (cli *SkillEffectSvcClient) EffectCalculate(req *pb.SEEffectCalReq) (*pb.SEEffectCalResp, error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cm.PrintReqLog(ctx, "Effect-Calculate", req)
	defer func() {
		cancel()
		elapsed := time.Since(start)
		log.Printf("Draw-Phase-Confirm took %s", elapsed)
	}()
	var p peer.Peer

	res, err := cli.svcClient.EffectCalculate(
		ctx, req,
		grpc.FailFast(false), // To wait a resolver returning addrs.
		grpc.Peer(&p),
	)
	if err != nil {
		fmt.Printf("Skill-Inst-Calc: cannot resolve:\n\t%v", err)
		return nil, err
	}
	fmt.Printf("Skill-Inst-Calc:resolve by %s\n", p.Addr)
	return res, nil
}

// EffectCalculateWrap
func (cli *SkillEffectSvcClient) EffectCalculateWrap(
	key string,
	fromTime *dataPb.EffectTiming, toTime *dataPb.EffectTiming,
	gameSet *dataPb.GameDataSet,
) (*dataPb.GameDataSet, error) {
	res, err := cli.EffectCalculate(&pb.SEEffectCalReq{
		Id:             key,
		FromTime:       fromTime,
		ToTime:         toTime,
		GamesetInstant: gameSet,
		FromCli:        cli.clientName,
	})
	return res.GamesetResult, err
}

func (cli *SkillEffectSvcClient) DiceCalculate(req *pb.SEDiceCalReq) (*pb.SEDiceCalResp, error) {
	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cm.PrintReqLog(ctx, "Dice-Calculate", req)
	defer func() {
		cancel()
		elapsed := time.Since(start)
		log.Printf("Draw-Phase-Confirm took %s", elapsed)
	}()
	var p peer.Peer

	res, err := cli.svcClient.DiceCalculate(
		ctx, req,
		grpc.FailFast(false), // To wait a resolver returning addrs.
		grpc.Peer(&p),
	)
	if err != nil {
		fmt.Printf("Skill-Inst-Calc: cannot resolve:\n\t%v", err)
		return nil, err
	}
	fmt.Printf("Skill-Inst-Calc:resolve by %s\n", p.Addr)
	return res, nil
}

type DiceResult [][]int

func (cli *SkillEffectSvcClient) DiceCalculateWrap(
	incomeDice int32,
	act int,
	InvolveEff []*dataPb.EffectResult,
) (DiceResult, error) {
	res, err := cli.DiceCalculate(&pb.SEDiceCalReq{
		IncomeDice:   incomeDice,
		Act:          int32(act),
		EffectResult: InvolveEff,
	})
	if err != nil {
		return nil, err
	}
	var resd DiceResult
	for k := range res.DiceResult {
		var asd []int
		for _, vb := range res.DiceResult[k].Value {
			asd = append(asd, int(vb))
		}
		resd = append(resd, asd)
	}
	return resd, nil
}

func (dr *DiceResult) ToString() string {
	tmpStr := ""
	for _, v := range *dr {
		tmpStr += strings.Replace(fmt.Sprint(v), " ", ",", -1) + ";"
	}
	return tmpStr
}

func (dr *DiceResult) ToTotal() int32 {
	var tmpInt []int32
	for _, v := range *dr {
		vv := int32(0)
		for _, vk := range v {
			vv += int32(vk)
		}
		tmpInt = append(tmpInt, vv)
	}
	sort.Slice(tmpInt, func(i, j int) bool {
		return tmpInt[i] > tmpInt[j]
	})
	return tmpInt[0]
}

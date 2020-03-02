package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

const DataMountPath []string = []string{"data", "ActionCard.csv"}

var EventCardSetArr []*pb.EventCard = nil

/**
public static const SWD:int = 1;
public static const ARW:int = 2;
public static const DEF:int = 3;
public static const MOVE:int = 4;
public static const SPC:int = 5;
public static const EVT:int = 6;
*/
/**
ACTION_EVENT_NO =[
	nil,                          # 0
	[:occur_chance_event, 1],     # 1 チャンスカードx1
	[:occur_chance_event, 2],     # 2 チャンスカードx2
	[:occur_chance_event, 3],     # 3 チャンスカードx3
	[:occur_chance_event, 4],     # 4 チャンスカードx4
	[:occur_chance_event, 5],     # 5 チャンスカードx5
	[:occur_heal_event, 1],       # 6 HP回復x1
	[:occur_heal_event, 2],       # 7 HP回復x2
	[:occur_heal_event, 3],       # 8 HP回復x3
	[:occur_cure_event],          # 9 パラメーター異常回復
	[:occur_quick_event],         # 10 必ずイニシアチブを取る
	[:occur_curse_event, 1],      # 11 カースカードx1
	[:occur_curse_event, 2],      # 12 カースカードx2
	[:occur_curse_event, 3],      # 13 カースカードx3
	[:occur_curse_event, 4],      # 14 カースカードx4
	[:occur_curse_event, 5],      # 15 カースカードx5
	[:occur_chalice_event, 1],    # 16 聖杯カード1
	[:occur_poison_event, 1],     # 17 毒杯カード1
	nil,                          # 18 ?
	[:occur_damage_event, 2],     # 19 ウイルス
]
*/

func init() {
	loadCardToMem()
}

func loadCardToMem() {
	csvFile, _ := os.Open("people.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		id_i, _ := strconv.ParseInt(line[0], 10, 32)
		u_vali, _ := strconv.ParseInt(line[2], 10, 32)
		b_vali, _ := strconv.ParseInt(line[4], 10, 32)
		EventCardSetArr = append(EventCardSetArr, &pb.EventCard{
			Id:         int32(id_i),
			UpOption:   convertNewType(&line[1]),
			UpVal:      int32(u_vali),
			DownOption: convertNewType(&line[3]),
			DownVal:    int32(b_vali),
		})
	}

}

func convertNewType(in *string) pb.EventCardType {
	switch *in {
	case "1":
		return pb.EventCardType_ATTACK
	case "2":
		return pb.EventCardType_GUN
	case "3":
		return pb.EventCardType_DEFENCE
	case "4":
		return pb.EventCardType_MOVE
	case "5":
		return pb.EventCardType_STAR
	}
	return pb.EventCardType_NULL
}

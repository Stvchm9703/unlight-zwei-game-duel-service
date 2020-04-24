package serverCtl

import (
	pb "ULZGameDuelService/proto"
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/golang/protobuf/proto"
)

var (
	DataMountPath   []string        = []string{"data", "ActionCards.csv"}
	DataSavePath    []string        = []string{"data", "ActionCardSet.json"}
	DataSavePathb   []string        = []string{"data", "ActionCardSet.pbc"}
	EventCardSetArr []*pb.EventCard = nil
)

func init() {
	fmt.Println("load pbc")
	if err := LoadFromFilePBC(true); err != nil {
		fmt.Println("load pbc fail")
		fmt.Println("load csv")
		LoadCardCSVToMem()
		fmt.Println("store pbc")
		SaveToFilePBC()
		LoadFromFilePBC(true)
	}

}

func LoadCardCSVToMem() {
	wd, _ := os.Getwd()
	wdl := []string{wd}
	wdl = append(wdl, DataMountPath...)
	csvFile, _ := os.Open(path.Join(wdl...))
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("load-card-csv", err)
		}
		fmt.Println(line)
		id_i, _ := strconv.ParseInt(line[0], 10, 32)
		u_vali, _ := strconv.ParseInt(line[2], 10, 32)
		b_vali, _ := strconv.ParseInt(line[4], 10, 32)
		mount, val := convCardFunc(&line[5])

		EventCardSetArr = append(EventCardSetArr, &pb.EventCard{
			Id:         int32(id_i),
			UpOption:   convCardType(&line[1]),
			UpVal:      int32(u_vali),
			DownOption: convCardType(&line[3]),
			DownVal:    int32(b_vali),
			FuncMount:  mount,
			MountVal:   val,
		})

	}
}

func SaveToFilePBC() {
	if len(EventCardSetArr) == 0 {
		return // no data!
	}
	wd, _ := os.Getwd()
	wdl := []string{wd}
	wdl = append(wdl, "data")
	if _, err := os.Stat(filepath.Join(wdl...)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(wdl...), os.ModePerm)
	}
	wdl = []string{wd}
	wdl = append(wdl, DataSavePathb...)
	fmt.Println("save in:", wdl)
	f, err := os.OpenFile(filepath.Join(wdl...), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	for _, v := range EventCardSetArr {
		tmp, _ := proto.Marshal(v)
		fmt.Println(string(tmp))
		fmt.Fprintln(f, string(tmp))
	}
	f.Close()
}

func LoadFromFilePBC(skip bool) error {
	EventCardSetArr = nil
	wd, _ := os.Getwd()
	wdl := []string{wd}
	wdl = append(wdl, DataSavePathb...)
	csvFile, err := os.Open(path.Join(wdl...))
	if err != nil {
		return err
	}
	reader := (bufio.NewReader(csvFile))
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
			return err
		}
		// fmt.Println(string(line))
		var y pb.EventCard
		proto.Unmarshal(line, &y)
		// fmt.Printf("%#v", y)
		if y.FuncMount == 0 && skip {
			EventCardSetArr = append(EventCardSetArr, &y)
		}
	}
	return nil
}

func convCardType(in *string) pb.EventCardType {
	/**
	public static const SWD:int = 1;
	public static const ARW:int = 2;
	public static const DEF:int = 3;
	public static const MOVE:int = 4;
	public static const SPC:int = 5;
	public static const EVT:int = 6;
	*/

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

func convCardFunc(inVal *string) (pb.EventCardFunc, int32) {
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

	switch *inVal {
	case "0":
		return pb.EventCardFunc_null, 0
	case "1":
		return pb.EventCardFunc_occur_chance_event, 1
	case "2":
		return pb.EventCardFunc_occur_chance_event, 2
	case "3":
		return pb.EventCardFunc_occur_chance_event, 3
	case "4":
		return pb.EventCardFunc_occur_chance_event, 4
	case "5":
		return pb.EventCardFunc_occur_chance_event, 5
	case "6":
		return pb.EventCardFunc_occur_heal_event, 1
	case "7":
		return pb.EventCardFunc_occur_heal_event, 2
	case "8":
		return pb.EventCardFunc_occur_heal_event, 3
	case "9":
		return pb.EventCardFunc_occur_cure_event, 0
	case "10":
		return pb.EventCardFunc_occur_quick_event, 0
	case "11":
		return pb.EventCardFunc_occur_curse_event, 1
	case "12":
		return pb.EventCardFunc_occur_curse_event, 2
	case "13":
		return pb.EventCardFunc_occur_curse_event, 3
	case "14":
		return pb.EventCardFunc_occur_curse_event, 4
	case "15":
		return pb.EventCardFunc_occur_curse_event, 5
	case "16":
		return pb.EventCardFunc_occur_chalice_event, 1
	case "17":
		return pb.EventCardFunc_occur_poison_event, 1
	case "18":
		return pb.EventCardFunc_null, 0
	case "19":
		return pb.EventCardFunc_occur_damage_event, 2
	}
	return pb.EventCardFunc_null, 0
}

func genCardSet(num int, costLimit int) (out []*pb.EventCard) {
	nm := len(EventCardSetArr)
	rand.Seed(int64(time.Now().UnixNano()))
	for i := 0; i <= num; i++ {
		out = append(
			out,
			EventCardSetArr[rand.Intn(nm)])
	}
	return
}

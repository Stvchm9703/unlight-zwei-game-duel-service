package main

import (
	ctl "ULZGameDuelService/pkg/serverCtl"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	HostTmp := ctl.GenCardSet(150, 0)
	time.Sleep(1200)
	DuelTmp := ctl.GenCardSet(150, 0)
	wd, _ := os.Getwd()
	wdl := []string{wd}
	wdl = append(wdl, "data")
	if _, err := os.Stat(filepath.Join(wdl...)); os.IsNotExist(err) {
		os.Mkdir(filepath.Join(wdl...), os.ModePerm)
	}
	wdl = []string{wd}
	wdl = append(wdl, "HostTmp.json")
	fmt.Println("save in:", wdl)
	f, err := os.OpenFile(filepath.Join(wdl...), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	HostTmpJson, _ := json.Marshal(HostTmp)
	fmt.Fprint(f, string(HostTmpJson))
	f.Close()

	wdl = []string{wd}
	wdl = append(wdl, "DuelTmp.json")
	f, err = os.OpenFile(filepath.Join(wdl...), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	DuelTmpJson, _ := json.Marshal(DuelTmp)
	fmt.Fprint(f, string(DuelTmpJson))
	f.Close()

}

package main

import (
	ctl "ULZGameDuelService/pkg/serverCtl"
	"fmt"
)

func main() {
	for k, v := range ctl.EventCardSetArr {
		fmt.Printf("%v : %#v \n", k, v)
	}
}

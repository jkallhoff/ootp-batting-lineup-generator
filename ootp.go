package main

import (
	"fmt"
)

type Player struct {
	Position, Name, Level, Bats, Overall, Potential string
	Age, Contact, Gap, Power, Eye, AvoidsStrikeOuts int64
	Catcher, FirstBase, SecondBase, ThirdBase       int64
	ShortStop, LeftField, RightField, CenterField   int64
	Speed, Stealing, BaseRunning, AB, HR, OPSPlus   int64
	AVG, OBP, SLG, WOBA, VORP, WAR, ZR              float64
}

func main() {

	players := fetchPlayers()

	//OrderedBy(players, sortByIncreasingAge, sortByDecreasingAtBats).Sort()

	for _, p := range players {
		fmt.Printf("%#v\n", p)
	}
}

package main

import (
	"log"
)

type Player struct {
	Name, Position, Bats                                         string
	Age, Experience, Contact, Gap, Power, Eye, Strikeout, AtBats int64
	Catcher, FirstBase, SecondBase, ThirdBase                    int64
	ShortStop, LeftField, RightField, CenterField                int64
	OnBase, Slugging, Ops                                        float64
	OpsPlus                                                      int64
}

func main() {

	players := fetchPlayers()

	OrderedBy(players, sortByIncreasingAge, sortByDecreasingAtBats).Sort()

	for _, p := range players {
		log.Printf("Name: %v Age: %v AtBats: %v \n", p.Name, p.Age, p.AtBats)
	}
}

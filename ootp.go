package main

import (
//"fmt"
)

var Players = make(map[string]*Player)

func main() {
	configuration := new(Configuration)
	configuration.Load()

	A := make(PlayersList, 0)
	AA := make(PlayersList, 0)
	AAA := make(PlayersList, 0)

	loadSamples("A")
	loadSamples("AA")
	loadSamples("AAA")

	//OrderedBy(Players, sortByIncreasingName).Sort()

	for _, player := range Players {
		player.Process(configuration)
		switch player.Level {
		case "A":
			A = append(A, player)
			break
		case "AA":
			AA = append(AA, player)
			break
		case "AAA":
			AAA = append(AAA, player)
			break
		}
	}

	A.Print("A")
	AA.Print("AA")
	AAA.Print("AAA")
}

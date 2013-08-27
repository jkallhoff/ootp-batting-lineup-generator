package main

import (
	"fmt"
)

const GREEN string = "\x1b[32m"
const RESET string = "\x1b[0m"
const WHITE string = "\x1b[37m"
const RED string = "\x1b[31m"
const YELLOW string = "\x1b[33m"
const BLUE string = "\x1b[34m"

var Players = make(map[string]*Player)

func main() {
	configuration := new(Configuration)
	configuration.Load()

	loadSamples("A")
	loadSamples("AA")
	loadSamples("AAA")

	//OrderedBy(Players, sortByIncreasingName).Sort()

	fmt.Println("Results")
	fmt.Println("---------------------")

	for _, player := range Players {
		fmt.Printf("\n %s - %s Age: %v\n", player.Name, player.Level, player.Age)
		fmt.Println("----------------------")
		for _, sample := range player.Samples {
			fmt.Printf("%#v\n", sample)
		}
	}
}

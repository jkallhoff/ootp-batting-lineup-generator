package main

import (
	"fmt"
)

//TypeDefs
type PlayerResults []string
type PlayersList []*Player

func (list PlayersList) Print(league string) {
	fmt.Printf("\n\n"+GREEN+"Output Results for League %s", league)
	fmt.Println("\n========================================" + RESET)
	for _, player := range list {
		if len(player.Output) > 0 {
			fmt.Printf("\n %s Age: %v\n", player.Name, player.Age)
			fmt.Println("----------------------")
			for _, message := range player.Output {
				fmt.Printf("%s", message)
			}
		}
	}
}

type Player struct {
	Position, Name, Level, Bats                                         string
	Age, Contact, Gap, Power, Eye, AvoidsStrikeOuts, Overall, Potential int64
	Catcher, FirstBase, SecondBase, ThirdBase                           int64
	ShortStop, LeftField, RightField, CenterField                       int64
	Speed, Stealing, BaseRunning                                        int64
	Samples                                                             SamplesList
	Output                                                              PlayerResults
}

func (p *Player) Process(config *Configuration) {
	for _, sample := range p.Samples {
		p.processLeague(sample, config)
	}
}

func (p *Player) processLeague(sample *LeagueSample, config *Configuration) {
	if sample.AB >= config.MinAtBats {
		message := fmt.Sprintf("Qualifies for review at the %s level due to %v AB", sample.League, sample.AB)
		p.AddNeutralMessage(message)
	}
}

//Output functions
func (p *Player) AddNeutralMessage(message string) {
	var outputMessage = fmt.Sprintf(WHITE+"%s"+RESET+"\n", message)
	p.Output = append(p.Output, outputMessage)
}

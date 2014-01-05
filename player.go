package main

import (
	"fmt"
)

//TypeDefs
type PlayerResults []string
type PlayersList []*Player

func (list PlayersList) Print(league string) {
	for _, player := range list {
		if len(player.Output) > 0 {
			fmt.Printf("\n (%s) %s Age: %v\n", player.Level, player.Name, player.Age)
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
	LeagueInt                                                           int64
}

func (p *Player) Process(config *Configuration) {
	for _, sample := range p.Samples {
		p.processLeague(sample, config)
	}
	p.generateOutputMessages()
}

func (p *Player) processLeague(sample *LeagueSample, config *Configuration) {
	if sample.AB >= config.MinAtBats {
		//message := fmt.Sprintf("Qualifies for review at the %s level due to %v AB", sample.League, sample.AB)
		//p.addNeutralMessage(message)
		p.processLeagueAge(sample, config)
		p.processLeagueBattingPerformance(sample, config)
		sample.Processed = true
	}
}

func (p *Player) processLeagueAge(sample *LeagueSample, config *Configuration) {
	var workingAgeLimit int64

	switch sample.League {
	case "A":
		workingAgeLimit = config.AMaxAge
		break
	case "AA":
		workingAgeLimit = config.AAAMaxAge
		break
	case "AAA":
		workingAgeLimit = config.AAAMaxAge
		break
	}

	gettingOldAage := workingAgeLimit - 2

	if p.Age < gettingOldAage {
		sample.AtAge = true
	} else if p.Age >= gettingOldAage && p.Age < workingAgeLimit {
		sample.GettingOld = true
	} else {
		sample.OverAge = true
	}
}

func (p *Player) processLeagueBattingPerformance(sample *LeagueSample, config *Configuration) {
	switch {
	case sample.WOBA >= config.TopwOBP:
		sample.KillingIt = true
		break
	case sample.WOBA >= config.MedianwOBP:
		sample.PlayingAverage = true
		break
	case sample.WOBA >= config.BottomwOBP:
		sample.PlayingBelowAverage = true
		break
	case sample.WOBA < config.BottomwOBP:
		sample.PlayingHorribly = true
		break
	}
}

func (p *Player) generateOutputMessages() {
	for _, sample := range p.Samples {
		if sample.Processed {
			switch {
			case sample.KillingIt:
				{
					switch {
					case p.LeagueInt == sample.LeagueInt:
						p.addGoodMessage("He is killing it at his current current level and should be moved up, if possible.")
						break
					case p.LeagueInt < sample.LeagueInt:
						p.addGoodMessage(fmt.Sprintf("He should be moved up right away; he's stuck in %s and should be higher than %s", p.Level, sample.League))
						break
					case p.LeagueInt > sample.LeagueInt:
						p.addGoodMessage(fmt.Sprintf("He killed it at %s and appears to be higher now", sample.League))
						break
					}
					break
				}
			case sample.PlayingAverage:
				{
					switch {
					case p.LeagueInt == sample.LeagueInt:
						switch {
						case sample.OverAge:
							p.addBadMessage("He's playing well but is too old for this level and needs to either advance or be cut")
							break
						case sample.GettingOld:
							p.addWarningMessage("He's playing well but isn't getting any younger")
							break
						default:
							p.addNeutralMessage("He's playing well and is at the level he should be at")
							break
						}
						break
					case p.LeagueInt < sample.LeagueInt:
						switch {
						case sample.AtAge:
							p.addGoodMessage(fmt.Sprintf("He's playing well at %s and thus should be moved up to %s", sample.League, sample.League))
							break
						default:
							p.addBadMessage(fmt.Sprintf("He's playing well at %s but is too old for that level, let alone %s. Move him WAY up or cut him", sample.League, p.Level))
							break
						}
						break
					case p.LeagueInt > sample.LeagueInt:
						switch {
						case sample.AtAge:
							p.addWarningMessage(fmt.Sprintf("He's playing well at %s but is currently at %s. Lets make sure he's where he should be", sample.League, p.Level))
							break
						default:
							p.addWarningMessage(fmt.Sprintf("He's playing well at %s but is too old for that level and thus should possibly be where he currently is or cut", sample.League, p.Level))
							break
						}
						break
					}
					break
				}
			}
		}
	}
}

//Output functions
func (p *Player) addNeutralMessage(message string) {
	var outputMessage = fmt.Sprintf("- "+WHITE+"%s"+RESET+"\n", message)
	p.Output = append(p.Output, outputMessage)
}

func (p *Player) addGoodMessage(message string) {
	var outputMessage = fmt.Sprintf("- "+GREEN+"%s"+RESET+"\n", message)
	p.Output = append(p.Output, outputMessage)
}

func (p *Player) addWarningMessage(message string) {
	var outputMessage = fmt.Sprintf("- "+YELLOW+"%s"+RESET+"\n", message)
	p.Output = append(p.Output, outputMessage)
}

func (p *Player) addColdMessage(message string) {
	var outputMessage = fmt.Sprintf("- "+BLUE+"%s"+RESET+"\n", message)
	p.Output = append(p.Output, outputMessage)
}

func (p *Player) addBadMessage(message string) {
	var outputMessage = fmt.Sprintf("- "+RED+"%s"+RESET+"\n", message)
	p.Output = append(p.Output, outputMessage)
}

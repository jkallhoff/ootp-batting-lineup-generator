package main

import (
	"fmt"
	"github.com/jkallhoff/gofig"
)

const GREEN string = "\x1b[32m"
const RESET string = "\x1b[0m"
const WHITE string = "\x1b[37m"
const RED string = "\x1b[31m"
const YELLOW string = "\x1b[33m"
const BLUE string = "\x1b[34m"

type Player struct {
	Position, Name, Level, Bats                                         string
	Age, Contact, Gap, Power, Eye, AvoidsStrikeOuts, Overall, Potential int64
	Catcher, FirstBase, SecondBase, ThirdBase                           int64
	ShortStop, LeftField, RightField, CenterField                       int64
	Speed, Stealing, BaseRunning, AB, HR, OPSPlus                       int64
	AVG, OBP, SLG, WOBA, VORP, WAR, ZR                                  float64
}

func (player *Player) NameAndPosition() string {
	return fmt.Sprintf("%s(%s %s %v stars)", player.Name, player.Position, player.Level, player.Potential)
}

func (player *Player) StatsLine() string {
	return fmt.Sprintf("wOBA: %0.3f WAR: %0.2f AB: %v", player.WOBA, player.WAR, player.AB)
}

type PlayersList []*Player

func (players *PlayersList) Process(metrics *Metrics) (results []string) {
	for _, p := range *players {
		if p.AB > metrics.MinAtBats {
			if p.WOBA >= metrics.TopwOBP && p.WAR >= metrics.GoodWAR {
				if p.Level == metrics.LeagueLevel {
					results = append(results, fmt.Sprintf(GREEN+"%s appears ready to advance. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
				}
				if metrics.OneLevelDown != "none" && p.Level == metrics.OneLevelDown {
					results = append(results, fmt.Sprintf(GREEN+"%s should be moved higher than %s immediately"+RESET, p.NameAndPosition(), metrics.LeagueLevel))
				}
				if metrics.OneLevelUp != "none" && p.Level == metrics.OneLevelUp {
					results = append(results, fmt.Sprintf(GREEN+"%s is where he should be in %s"+RESET, p.NameAndPosition(), metrics.OneLevelUp))
				}
			}
			if p.WOBA >= metrics.MedianwOBP && (p.WOBA < metrics.TopwOBP && p.WAR > 0) {
				if p.Age > metrics.MaxAge {
					results = append(results, fmt.Sprintf(GREEN+"%s is too old but playing at or above level. Should be moved up. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
				}
				if p.Age <= metrics.MaxAge && p.Age >= metrics.CalcedGettingOld {
					results = append(results, fmt.Sprintf("%s is reaching max age but playing at the right level. %s", p.NameAndPosition(), p.StatsLine()))
				}
				if p.Age < metrics.CalcedGettingOld {
					results = append(results, fmt.Sprintf("%s is playing where he should be. %s", p.NameAndPosition(), p.StatsLine()))
				}
			}
			if p.WOBA < metrics.MedianwOBP && (p.WOBA >= metrics.BottomwOBP && p.WAR >= metrics.MinWAR) {
				if p.Age > metrics.MaxAge {
					results = append(results, fmt.Sprintf(RED+"%s is too old and not playing well at this level. He should be cut. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
				} else {
					if metrics.OneLevelDown != "none" && p.Level == metrics.OneLevelDown {
						results = append(results, fmt.Sprintf(BLUE+"%s has not played well at %s and is at the right level now. %s"+RESET, p.NameAndPosition(), metrics.LeagueLevel, p.StatsLine()))
					} else {
						if p.Age <= metrics.MaxAge && p.Age >= metrics.CalcedGettingOld {
							if p.Potential >= metrics.MinStars {
								results = append(results, fmt.Sprintf(YELLOW+"%s is reaching max age and not playing well at this level. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
							}
							if p.Potential < metrics.MinStars {
								results = append(results, fmt.Sprintf(RED+"%s is reaching max age and not playing well at this level. He should be cut. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
							}
						}
						if p.Age < metrics.CalcedGettingOld {
							results = append(results, fmt.Sprintf(YELLOW+"%s is young and not playing well. Possibly drop a level until skillset grows a bit. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
						}
					}
				}
			}
			if p.WOBA < metrics.BottomwOBP || p.WAR < metrics.MinWAR {
				if p.Age >= metrics.CalcedGettingOld && p.Age <= metrics.MaxAge {
					if p.Potential >= metrics.MinStars {
						results = append(results, fmt.Sprintf(BLUE+"%s is playing horribly and needs to be dropped a level. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
					}
					if p.Potential < metrics.MinStars {
						results = append(results, fmt.Sprintf(RED+"%s is reaching max age and playing horribly. He should be cut. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
					}
				}
				if p.Age < metrics.CalcedGettingOld {
					results = append(results, fmt.Sprintf(BLUE+"%s is young and playing horribly. Drop a level to grow skills. %s"+RESET, p.NameAndPosition(), p.StatsLine()))
				}
			}
		}
	}

	return
}

type Metrics struct {
	LeagueLevel, OneLevelDown, OneLevelUp            string
	MinAtBats, MaxAge, MinStars, CalcedGettingOld    int64
	MedianwOBP, MinWAR, GoodWAR, BottomwOBP, TopwOBP float64
}

func (metrics *Metrics) Load() {
	metrics.LeagueLevel = gofig.Str("leagueLevel")
	metrics.MinAtBats = gofig.Int("minAtBats")
	metrics.MaxAge = gofig.Int("maxAge")
	metrics.MinStars = gofig.Int("minStars")
	metrics.MedianwOBP = gofig.Float("medianwOBP")
	metrics.MinWAR = gofig.Float("minWAR")
	metrics.GoodWAR = gofig.Float("goodWAR")
	metrics.BottomwOBP = gofig.Float("bottomwOBP")
	metrics.TopwOBP = gofig.Float("topwOBP")

	metrics.CalcedGettingOld = metrics.MaxAge - 2

	switch metrics.LeagueLevel {
	case "A":
		metrics.OneLevelDown = "none"
		metrics.OneLevelUp = "AA"
	case "AA":
		metrics.OneLevelDown = "A"
		metrics.OneLevelUp = "AAA"
	case "AAA":
		metrics.OneLevelDown = "AA"
		metrics.OneLevelUp = "ML"
	case "ML":
		metrics.OneLevelDown = "AAA"
		metrics.OneLevelUp = "none"
	}
}

func main() {
	metrics := new(Metrics)
	metrics.Load()

	players := fetchPlayers(metrics.LeagueLevel + ".html")

	//OrderedBy(players, sortByIncreasingAge, sortByDecreasingAtBats).Sort()

	fmt.Println("\nAnalyzing Level " + metrics.LeagueLevel)
	fmt.Println("---------------------")
	results := players.Process(metrics)

	for _, result := range results {
		fmt.Println(result)
	}
}

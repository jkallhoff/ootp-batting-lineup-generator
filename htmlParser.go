package main

import (
	"code.google.com/p/go.net/html"
	"log"
	"os"
	"strconv"
	"strings"
)

var leagueBeingProcessed string

func loadSamples(league string) {
	htmlPath := "Exports/" + league + ".html"
	leagueBeingProcessed = league

	file, err := os.Open(htmlPath)
	if err != nil {
		log.Fatalln("Could not open the roster file")
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatalln("Could not close the file")
		}
	}(file)

	node, _ := html.Parse(file)
	processNode(node)

	return
}

func processNode(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "tr" {
		pullTogetherSample(node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		processNode(child)
	}
}

func pullTogetherSample(node *html.Node) {
	positionHolder := ""
	player := new(Player)
	sample := new(LeagueSample)
	iterator := 1
	for column := node.FirstChild; column != nil; column = column.NextSibling {
		if column.FirstChild != nil {
			switch iterator {
			case 4:
				positionHolder = column.FirstChild.Data
			case 6:
				name := strings.TrimSpace(column.FirstChild.Data)
				if name == "" || name == "Name" {
					break
				}

				if _, ok := Players[name]; ok {
					player = Players[name]
				} else {
					player = &Player{Name: name}
					Players[name] = player
				}
				player.Position = positionHolder
			case 12:
				player.Level = column.FirstChild.Data
				player.LeagueInt = fetchLeagueInt(player.Level)
			case 14:
				player.Age = fetchInteger(column.FirstChild)
			case 16:
				player.Bats = column.FirstChild.Data
			case 18:
				overall := column.FirstChild.Data
				player.Overall, _ = strconv.ParseInt(overall[:1], 10, 8)
			case 20:
				potential := column.FirstChild.Data
				player.Potential, _ = strconv.ParseInt(potential[:1], 10, 8)
				//log.Printf("%d", player.Potential)
			case 22:
				player.Contact = fetchInteger(column.FirstChild)
			case 24:
				player.Gap = fetchInteger(column.FirstChild)
			case 26:
				player.Power = fetchInteger(column.FirstChild)
			case 28:
				player.Eye = fetchInteger(column.FirstChild)
			case 30:
				player.AvoidsStrikeOuts = fetchInteger(column.FirstChild)
			case 32:
				player.Catcher = fetchInteger(column.FirstChild)
			case 34:
				player.FirstBase = fetchInteger(column.FirstChild)
			case 36:
				player.SecondBase = fetchInteger(column.FirstChild)
			case 38:
				player.ThirdBase = fetchInteger(column.FirstChild)
			case 40:
				player.ShortStop = fetchInteger(column.FirstChild)
			case 42:
				player.LeftField = fetchInteger(column.FirstChild)
			case 44:
				player.CenterField = fetchInteger(column.FirstChild)
			case 46:
				player.RightField = fetchInteger(column.FirstChild)
			case 48:
				player.Speed = fetchInteger(column.FirstChild)
			case 50:
				player.Stealing = fetchInteger(column.FirstChild)
			case 52:
				player.BaseRunning = fetchInteger(column.FirstChild)
			case 54:
				sample.AB = fetchInteger(column.FirstChild)
			case 56:
				sample.HR = fetchInteger(column.FirstChild)
			case 58:
				sample.AVG = fetchFloat(column.FirstChild)
			case 60:
				sample.OBP = fetchFloat(column.FirstChild)
			case 62:
				sample.SLG = fetchFloat(column.FirstChild)
			case 64:
				sample.WOBA = fetchFloat(column.FirstChild)
			case 66:
				sample.OPSPlus = fetchInteger(column.FirstChild)
			case 68:
				sample.VORP = fetchFloat(column.FirstChild)
			case 70:
				sample.WAR = fetchFloat(column.FirstChild)
			case 72:
				sample.ZR = fetchFloat(column.FirstChild)
			}
		}
		iterator++
	}

	if player.Name != "" && player.Name != "Name" {
		sample.League = leagueBeingProcessed
		sample.LeagueInt = fetchLeagueInt(sample.League)
		player.Samples = append(player.Samples, sample)
	}
}

func fetchLeagueInt(league string) int64 {
	switch league {
	case "A":
		return 0
		break
	case "AA":
		return 1
		break
	case "AAA":
		return 2
		break
	case "ML":
		return 3
		break
	}
	return -1
}

func fetchInteger(node *html.Node) (num int64) {
	num, _ = strconv.ParseInt(node.Data, 10, 16)
	return
}

func fetchFloat(node *html.Node) (num float64) {
	num, _ = strconv.ParseFloat(node.Data, 64)
	return
}

package main

import (
	"code.google.com/p/go.net/html"
	"log"
	"os"
	"strconv"
)

func fetchPlayers() (players []*Player) {
	htmlPath := "2013-08-21-21-03-45.html"

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
	players = processNode(node, players)

	return
}

func processNode(node *html.Node, players []*Player) (returnPlayers []*Player) {
	if node.Type == html.ElementNode && node.Data == "tr" {
		players = pullTogetherPlayer(node, players)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		players = processNode(child, players)
	}
	return players
}

func pullTogetherPlayer(node *html.Node, players []*Player) (returnPlayers []*Player) {
	player := new(Player)

	iterator := 1
	for column := node.FirstChild; column != nil; column = column.NextSibling {
		if column.FirstChild != nil {
			switch iterator {
			case 4:
				player.Position = column.FirstChild.Data
			case 6:
				player.Name = column.FirstChild.Data
			case 12:
				player.Level = column.FirstChild.Data
			case 14:
				player.Age = fetchInteger(column.FirstChild)
			case 16:
				player.Bats = column.FirstChild.Data
			case 18:
				player.Overall = column.FirstChild.Data
			case 20:
				player.Potential = column.FirstChild.Data
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
				player.AB = fetchInteger(column.FirstChild)
			case 56:
				player.HR = fetchInteger(column.FirstChild)
			case 58:
				player.AVG = fetchFloat(column.FirstChild)
			case 60:
				player.OBP = fetchFloat(column.FirstChild)
			case 62:
				player.SLG = fetchFloat(column.FirstChild)
			case 64:
				player.WOBA = fetchFloat(column.FirstChild)
			case 66:
				player.OPSPlus = fetchInteger(column.FirstChild)
			case 68:
				player.VORP = fetchFloat(column.FirstChild)
			case 70:
				player.WAR = fetchFloat(column.FirstChild)
			case 72:
				player.ZR = fetchFloat(column.FirstChild)
			}
		}
		iterator++
	}

	if player.Name != "" && player.Name != "Name" {
		players = append(players, player)
	}

	return players
}

func fetchInteger(node *html.Node) (num int64) {
	num, _ = strconv.ParseInt(node.Data, 10, 8)
	return
}

func fetchFloat(node *html.Node) (num float64) {
	num, _ = strconv.ParseFloat(node.Data, 64)
	return
}

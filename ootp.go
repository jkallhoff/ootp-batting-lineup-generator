package main

import (
	"code.google.com/p/go.net/html"
	"log"
	"os"
	"strconv"
)

type Player struct {
	Name, Position, Bats                                         string
	Age, Experience, Contact, Gap, Power, Eye, Strikeout, AtBats int64
	Catcher, FirstBase, SecondBase, ThirdBase                    int64
	ShortStop, LeftField, RightField, CenterField                int64
	OnBase, Slugging, Ops                                        float64
	OpsPlus                                                      int64
}

var (
	players       []*Player
	headerSkipped bool
)

func main() {
	htmlPath := "2013-05-15-18-16-23.html"

	file, err := os.Open(htmlPath)
	if err != nil {
		log.Fatalln("Could not open the file")
	}

	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			log.Fatalln("Could not close the file")
		}
	}(file)

	node, _ := html.Parse(file)
	processNode(node)

	for _, p := range players {
		log.Printf("%#v \n", p)
	}
	log.Println("NUMBER OF PLAYERS: ", len(players))
}

func processNode(node *html.Node) {
	if node.Type == html.ElementNode && node.Data == "tr" {
		pullTogetherPlayer(node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		processNode(child)
	}
}

func pullTogetherPlayer(node *html.Node) {
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
				player.Age = fetchInteger(column.FirstChild)
			case 14:
				player.Bats = column.FirstChild.Data
			case 16:
				player.Experience = fetchInteger(column.FirstChild)
			case 22:
				player.Contact = fetchInteger(column.FirstChild)
			case 24:
				player.Gap = fetchInteger(column.FirstChild)
			case 26:
				player.Power = fetchInteger(column.FirstChild)
			case 28:
				player.Eye = fetchInteger(column.FirstChild)
			case 30:
				player.Strikeout = fetchInteger(column.FirstChild)
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
				player.AtBats = fetchInteger(column.FirstChild)
			case 50:
				player.OnBase = fetchFloat(column.FirstChild)
			case 52:
				player.Slugging = fetchFloat(column.FirstChild)
			case 54:
				player.Ops = fetchFloat(column.FirstChild)
			case 56:
				player.OpsPlus = fetchInteger(column.FirstChild)
			}
		}
		iterator++
	}
	if player.Name != "" && player.Name != "Name" {
		players = append(players, player)
	}
}

func fetchInteger(node *html.Node) (num int64) {
	num, _ = strconv.ParseInt(node.Data, 10, 8)
	return
}

func fetchFloat(node *html.Node) (num float64) {
	num, _ = strconv.ParseFloat(node.Data, 64)
	return
}

package main

/*import (
	"sort"
)

type playerLessFunc func(player1, player2 *Player) bool

type playerSorter struct {
	players     map[string]*Player
	sortMethods []playerLessFunc
}

func (sorter *playerSorter) Sort() {
	sort.Sort(sorter)
}

func OrderedBy(players map[string]*Player, sortMethods ...playerLessFunc) *playerSorter {
	return &playerSorter{
		players:     players,
		sortMethods: sortMethods,
	}
}

func (ps *playerSorter) Len() int {
	return len(ps.players)
}

func (ps *playerSorter) Swap(i, j int) {
	ps.players[i], ps.players[j] = ps.players[j], ps.players[i]
}

func (ps *playerSorter) Less(i, j int) bool {
	p, q := ps.players[i], ps.players[j]

	var k int
	for k = 0; k < len(ps.sortMethods)-1; k++ {
		sortMethod := ps.sortMethods[k]
		switch {
		case sortMethod(p, q):
			return true
		case sortMethod(q, p):
			return false
		}
	}
	return ps.sortMethods[k](p, q)
}

//Column Sorts

func sortByIncreasingAge(player1, player2 *Player) bool {
	return player1.Age < player2.Age
}

func sortByDecreasingAge(player1, player2 *Player) bool {
	return player1.Age > player2.Age
}

func sortByIncreasingName(player1, player2 *Player) bool {
	return player1.Name < player2.Name
}*/

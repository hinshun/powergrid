package powerutil

import "github.com/hinshun/powergrid"

func PlayerOwnedCities(player powergrid.Player, board powergrid.Board) uint {
	count := uint(0)
	nodes := board.Nodes()
	for _, node := range nodes {
		players := node.Cities()
		for _, playerCity := range players {
			if playerCity == player {
				count++
			}
		}
	}
	return count
}

type Collection interface {
	Len() int
	Swap(i, j int)
}

func Shuffle(collection Collection) {
	for i := 0; i < collection.Len(); i++ {
		j := rand.Intn(i + 1)
		collection.Swap(i, j)
	}
}

type PlayersByTurnOrder struct {
	Players []powergrid.Player
	Board   powergrid.Board
}

func (p PlayersByTurnOrder) Len() int      { return len(p.Players) }
func (p PlayersByTurnOrder) Swap(i, j int) { p.Players[i], p.Players[j] = p.Players[j], p.Players[i] }

// TODO: take into account more than just cities
func (p PlayersByTurnOrder) Less(i, j int) bool {
	return PlayerOwnedCities(p.Players[i], p.Board) < PlayerOwnedCities(p.Players[j], p.Board)
}

type PowerPlants struct {
	PowerPlants []powergrid.PowerPlant
}

func (pp *PowerPlants) Len() int { return len(p.PowerPlants) }

func (pp *PowerPlants) Swap(i, j int) {
	pp.PowerPlants[i], pp.PowerPlants[j] = pp.PowerPlants[j], pp.PowerPlants[i]
}

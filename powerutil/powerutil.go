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

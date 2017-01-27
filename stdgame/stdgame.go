package stdgame

import (
	"fmt"

	"github.com/hinshun/powergrid"
	"github.com/hinshun/powergrid/player"
)

type game struct {
	players      []powergrid.Player
	currentPhase phase
}

func New(numPlayers uint) powergrid.Game {
	players := []powergrid.Player{}
	for i := uint(0); i < numPlayers; i++ {
		players = append(players, player.New(i))
	}
	return &game{
		players:      players,
		currentPhase: playerOrder,
	}
}

func (g *game) GridInfo() string {
	return "power overwhelming"
}

func (g *game) Run() {
	defer fmt.Println("CLI terminated.")
	var input string
	fmt.Scanf("%s", &input)

	fmt.Printf("You entered %s.\n", input)
}

type phase uint

const (
	playerOrder phase = iota
	auctionPlants
	buyResources
	build
	powerYourShit
	resources
)

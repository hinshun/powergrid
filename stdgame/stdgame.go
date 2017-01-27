package stdgame

import (
	"fmt"
	"sort"

	"github.com/hinshun/powergrid"
	"github.com/hinshun/powergrid/player"
	"github.com/hinshun/powergrid/powerutil"
)

type game struct {
	players      []powergrid.Player
	currentPhase phase
	board        powergrid.Board
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

func (g *game) loop() {
	for p, f := range map[phase]func(){
		playerOrder:   g.playerOrder,
		auctionPlants: g.auctionPlants,
		buyResources:  g.buyResources,
		build:         g.build,
		powerYourShit: g.powerYourShit,
		resources:     g.resources,
	} {
		g.currentPhase = p
		f()
	}
}

func (g *game) playerOrder() {
	fmt.Println("Reordering players...")
	sort.Sort(powerutil.PlayersByTurnOrder{g.players, g.board})
	fmt.Printf("The player order is: %s\n", g.players)
}

func (g *game) auctionPlants() {
	fmt.Println("Auction time!")
	// TODO
}

func (g *game) buyResources() {
	fmt.Println("Buy resources!")
	for _, player := range g.players {
		fmt.Printf("It's player %s's turn to buy resources.\n", player.Name())
		g.buyResourcesPlayer(player)
	}
}
func (g *game) buyResourcesPlayer(player powergrid.Player) {
	// TODO
}

func (g *game) build() {
	fmt.Println("Build!")
	for _, player := range g.players {
		fmt.Printf("It's player %s's turn to build.\n", player.Name())
		g.buildPlayer(player)
	}
}

func (g *game) buildPlayer(player powergrid.Player) {
	// TODO
}

func (g *game) powerYourShit() {
	fmt.Println("Power Your Shit!")
	for _, player := range g.players {
		fmt.Printf("Player %s is powering his shit.\n", player.Name())
	}
	// TODO
}

func (g *game) resources() {
	// TODO replenish resources
}

func (g *game) Run() {
	defer fmt.Println("CLI terminated.")
	fmt.Println("Picking random player order...")
	fmt.Printf("The player order is: %s\n", g.players)

	g.loop()
	// TODO:
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

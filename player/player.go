package player

import "github.com/hinshun/powergrid"

type player struct {
	color   uint
	elektro powergrid.Elektro
}

func New(color uint) powergrid.Player {
	return player{color: color}
}

func (p player) Elektro() powergrid.Elektro {
	return p.elektro
}

func (p player) PowerPlants() []powergrid.PowerPlant {
	return nil
}

var _ = powergrid.Player((*player)(nil))

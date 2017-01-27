package player

import (
	"fmt"

	"github.com/hinshun/powergrid"
)

type player struct {
	color       uint
	elektro     powergrid.Elektro
	powerPlants []powergrid.PowerPlant
}

func New(color uint) powergrid.Player {
	return player{color: color}
}

func (p player) Name() string {
	return fmt.Sprint(p.color)
}

func (p player) Elektro() powergrid.Elektro {
	return p.elektro
}

func (p player) PowerPlants() []powergrid.PowerPlant {
	return p.powerPlants
}

var _ = powergrid.Player((*player)(nil))

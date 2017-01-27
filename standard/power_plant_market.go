package standard

import (
	"errors"

	"github.com/hinshun/powergrid"
)

var (
	ErrPowerPlantNotFound = errors.New("power plant not found")
)

type powerPlantMarket struct {
	powerPlants []powergrid.PowerPlant
}

func NewPowerPlantMarket(powerPlants []powergrid.PowerPlant) powergrid.PowerPlantMarket {
	market := &powerPlantMarket{powerPlants: powerPlants}

	for _, powerPlant := range powerPlants {
		market.Add(powerPlant)
	}

	return market
}

func (ppm *powerPlantMarket) Inventory() []powergrid.PowerPlant {
	return ppm.powerPlants
}

func (ppm *powerPlantMarket) Add(addPlant powergrid.PowerPlant) {
	for i, powerPlant := range ppm.powerPlants {
		if addPlant.Ordinal() < powerPlant.Ordinal() {
			ppm.powerPlants = append(
				append(ppm.powerPlants[:i], addPlant),
				ppm.powerPlants[i+1:]...,
			)
			return
		}
	}

	ppm.powerPlants = append(ppm.powerPlants, addPlant)
}

func (ppm *powerPlantMarket) Remove(removePlant powergrid.PowerPlant) error {
	for i, powerPlant := range ppm.powerPlants {
		if removePlant == powerPlant {
			ppm.powerPlants = append(ppm.powerPlants[:i], ppm.powerPlants[i+1:]...)
			return nil
		}
	}

	return ErrPowerPlantNotFound
}

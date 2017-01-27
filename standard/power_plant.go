package standard

import (
	"github.com/hinshun/powergrid"
)

// PowerPlant is an implementation of the powergrid.PowerPlant
// interface that works for the game's standard power plants
type standardPowerPlant struct {
	ordinal       uint
	fuelType      []powergrid.Resource
	fuelRequired  uint
	powerCapacity uint
	inventory     powergrid.ResourceSet
}

// NewPowerPlant constructs a standard PowerPlant
func NewPowerPlant(
	ordinal uint,
	fuelType []powergrid.Resource,
	fuelRequired uint,
	powerCapacity uint) powergrid.PowerPlant {
	return &standardPowerPlant{
		ordinal:       ordinal,
		fuelType:      fuelType,
		fuelRequired:  fuelRequired,
		powerCapacity: powerCapacity,
		inventory:     make(powergrid.ResourceSet),
	}
}

// Ordinal returns the ordinal of the PowerPlant
func (pp *standardPowerPlant) Ordinal() uint {
	return pp.ordinal
}

// FuelType returns a slice of acceptable Resources for the PowerPlant
func (pp *standardPowerPlant) FuelType() []powergrid.Resource {
	return pp.fuelType
}

// FuelRequired returns the quantity of fuel required by the PowerPlant
func (pp *standardPowerPlant) FuelRequired() uint {
	return pp.fuelRequired
}

// PowerCapacity returns the number of cities that the PowerPlant can power
func (pp *standardPowerPlant) PowerCapacity() uint {
	return pp.powerCapacity
}

// Inventory returns the quantity of each resource stored in the PowerPlant
func (pp *standardPowerPlant) Inventory() powergrid.ResourceSet {
	return pp.inventory
}

// AddFuel adds fuel to the PowerPlant
func (pp *standardPowerPlant) AddFuel(fuel powergrid.ResourceSet) {
	for fuelType, quantity := range fuel {
		currentQuantity, ok := pp.inventory[fuelType]
		if !ok {
			currentQuantity = 0
		}
		pp.inventory[fuelType] = currentQuantity + quantity
	}
}

// SpendFuel removes fuel from the PowerPlant
func (pp *standardPowerPlant) SpendFuel(fuel powergrid.ResourceSet) {
	negFuel := make(powergrid.ResourceSet)
	for fuelType, quantity := range fuel {
		negFuel[fuelType] = -quantity
	}
	pp.AddFuel(negFuel)
}

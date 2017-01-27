package standard

import (
	"github.com/hinshun/powergrid"
)

const (
	maxCoalOilGarbage uint = 24
	maxUranium             = 12
)

type resourceMarket struct {
	powergrid.ResourceSet
}

// NewResourceMarket creates an empty resource market ready to use for Germany
// or the United States.
func NewResourceMarket() powergrid.ResourceMarket {
	return resourceMarket{}
}

// Inventory returns a snapshot of the current market supply of each resource.
func (r resourceMarket) Inventory() powergrid.ResourceSet {
	// Return a copy of the internal set of resources.
	inventory := powergrid.ResourceSet{}
	for resource, supply := range r.ResourceSet {
		inventory[resource] = supply
	}
	return inventory
}

type marginalCostFunc func(supply uint) powergrid.Elektro

func marginalCostCoalOilGarbage(supply uint) powergrid.Elektro {
	return powergrid.Elektro((maxCoalOilGarbage-supply)/3 + 1)
}

func marginalCostUranium(supply uint) powergrid.Elektro {
	if supply <= 4 {
		return powergrid.Elektro(10 + (4-supply)*2)
	}

	return powergrid.Elektro(maxUranium - supply + 1)
}

var marginalCostFuncs = map[powergrid.Resource]marginalCostFunc{
	powergrid.Coal:    marginalCostCoalOilGarbage,
	powergrid.Oil:     marginalCostCoalOilGarbage,
	powergrid.Garbage: marginalCostCoalOilGarbage,
	powergrid.Uranium: marginalCostUranium,
}

// Cost returns the calculated price for the given order. If the order is not
// satisfiable then (0, false) is returned.
func (r resourceMarket) Cost(order powergrid.ResourceSet) (powergrid.Elektro, bool) {
	var runningCost powergrid.Elektro

	startingInventory := r.Inventory()

	for resource, demand := range order {
		if demand == 0 {
			// Don't want any? Don't bother!
			continue
		}

		marginalCost, ok := marginalCostFuncs[resource]
		if !ok {
			// We can't place the order if we can't price it.
			return 0, false
		}

		supply := startingInventory[resource]

		if demand > supply {
			// Can't place the order if we don't have the supply!
			return 0, false
		}

		for i := uint(0); i < demand; i++ {
			runningCost += marginalCost(supply - i)
		}
	}

	return runningCost, true
}

// PurchaseResources removes the given order from the current market supply of
// each resource.
func (r resourceMarket) PurchaseResources(order powergrid.ResourceSet) {
	for resource, demand := range order {
		supply, ok := r.ResourceSet[resource]
		if !ok {
			continue
		}

		if demand > supply {
			r.ResourceSet[resource] = 0
		} else {
			r.ResourceSet[resource] = supply - demand
		}
	}
}

// Replenish adds the given resources to the current market supply of each
// resource.
func (r resourceMarket) Replenish(resources powergrid.ResourceSet) {
	for resource, quantity := range resources {
		var maxQuantity uint

		supply := r.ResourceSet[resource]

		switch resource {
		case powergrid.Coal, powergrid.Oil, powergrid.Garbage:
			maxQuantity = maxCoalOilGarbage
		case powergrid.Uranium:
			maxQuantity = maxUranium
		default:
			maxQuantity = 0
		}

		if supply+quantity > maxQuantity {
			r.ResourceSet[resource] = maxQuantity
		} else {
			r.ResourceSet[resource] = supply + quantity
		}
	}
}

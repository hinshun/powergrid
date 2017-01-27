package standard

// TODO(rabrams) import and use testify
import (
	"reflect"
	"testing"

	"github.com/hinshun/powergrid"
)

func TestCreatePowerPlant(t *testing.T) {
	var types = []powergrid.Resource{powergrid.Coal, powergrid.Oil}
	plant := NewPowerPlant(1, types, 5, 10)
	if plant.Ordinal() != 1 {
		t.Fail()
	}
	if !reflect.DeepEqual(plant.FuelType(), types) {
		t.Fail()
	}
	if plant.FuelRequired() != 5 {
		t.Fail()
	}
	if plant.PowerCapacity() != 10 {
		t.Fail()
	}
}

func TestModifyPowerPlant(t *testing.T) {
	var types = []powergrid.Resource{powergrid.Coal, powergrid.Oil}
	plant := NewPowerPlant(1, types, 5, 10)

	var newFuel = make(powergrid.ResourceSet)
	newFuel[powergrid.Coal] = 2
	newFuel[powergrid.Oil] = 3
	plant.AddFuel(newFuel)
	plant.AddFuel(newFuel)
	if plant.Inventory()[powergrid.Coal] != 4 {
		t.Fail()
	}
	if plant.Inventory()[powergrid.Oil] != 6 {
		t.Fail()
	}
	newFuel[powergrid.Coal] = 1
	newFuel[powergrid.Oil] = 5
	plant.SpendFuel(newFuel)
	if plant.Inventory()[powergrid.Coal] != 3 {
		t.Fail()
	}
	if plant.Inventory()[powergrid.Oil] != 1 {
		t.Fail()
	}
}

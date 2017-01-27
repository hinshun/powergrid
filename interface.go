package powergrid

// Elektro is the game's unit of currency.
type Elektro uint

// A Resource is something.
type Resource uint8

// Resource Types
const (
	Coal Resource = iota
	Oil
	Garbage
	Uranium
)

// A ResourceSet is a collection of resource quantities.
type ResourceSet map[Resource]int

type Player interface {
	Elektro() Elektro
	PowerPlants() []PowerPlant
}

//
type ResourceMarket interface {
	Inventory() ResourceSet
	Cost(order ResourceSet) (Elektro, bool)
	PurchaseResources(order ResourceSet)
	Replenish(resources ResourceSet)
}

// A PowerPlant is a factory for electricity!
type PowerPlant interface {
	Ordinal() uint
	FuelType() []Resource
	FuelRequired() uint
	PowerCapacity() uint
	Inventory() ResourceSet
	AddFuel(fuel ResourceSet)
	SpendFuel(fuelUsed ResourceSet)
}

//
type PowerPlantMarket interface {
	Inventory() []PowerPlant
	Add(powerPlant PowerPlant)
	Remove(powerPlant PowerPlant)
}

type Bid struct {
	Bidder Player
	Price  Elektro
}

type Auction interface {
	Item() PowerPlant
	// Bidders returns a sorted list of players where the first is the one who bids next
	Bidders() []Player
	CurrentBidder() Player
	HighestBid() Bid
	Bid(bid Bid) error
	// Pass returns true if the round is over
	Pass() (bool, error)
}

type Map interface {
	Nodes() []Node
}

type Connection struct {
	Cost Elektro
	Node Node
}

type Node interface {
	Connections() []Connection
	Cities() []Player
}

type Game interface {
	Run()
	// GridInfo return type TBD
	GridInfo() string
}

type Stage interface {
	Replenish(market ResourceMarket) ResourceSet
	CanBuildCity(location Node) bool
	CanAuctionPowerPlant(powerPlant PowerPlant) bool
}

//var (
//	GermanyGameRules = Rules{
//		Stage1: GermanyStage1,
//		Stage2: GermanyStage2,
//		Stage3: GermanyStage3,
//		Map:    GermanyMap,
//	}
//
//	GermanyStage1 = StandardStage1{
//		ReplenishRate: ResourceSet{
//			Coal:    3,
//			Oil:     2,
//			Garbage: 3,
//			Uranium: 3,
//		},
//	}
//
//	GermanyStage2 = StandardStage2{
//		ReplenishRate: ResourceSet{
//			Coal:    4,
//			Oil:     3,
//			Garbage: 4,
//			Uranium: 5,
//		},
//	}
//
//	GermanyStage3 = StandardStage3{
//		ReplenishRate: ResourceSet{
//			Coal:    3,
//			Oil:     2,
//			Garbage: 3,
//			Uranium: 3,
//		},
//	}
//)
//
//type Stage1 struct{}
//type Stage2 struct{}
//type Stage2 struct{}

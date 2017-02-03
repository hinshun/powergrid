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

type Connection struct {
	Cost Elektro
	Node Node
}

type AuctionMove struct {
	Passes bool
	Bid Elektro
}

func (am *AuctionMove) Type() MoveType {
	return AuctionType
}

type BuyResourceMove struct {
	Resources ResourceSet
}

func (brm *BuyResourceMove) Type() MoveType {
	return BuyResourceType
}

type BuildHomesMove struct {
	Nodes []Node
}

func (bhm *BuildHomesMove) Type() MoveType {
	return BuildHomesType
}

type PowerMove struct {
	Plants []PowerPlant
}

func (pm *PowerMove) Type() MoveType {
	return PowerType
}

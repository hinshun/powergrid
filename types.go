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

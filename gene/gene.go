package gene

import "fmt"

type GeneKind int

// A gene is the arrow in a mechanism state machine diagram
// It performs a specific action or side effect to proceed to the next state of the mechanism on a condition
type Gene interface {
	Kind() GeneKind

	// Perform the side effect
	Perform(s Simulation, n Body)
}

const (
	// Movement genes perform a movement side effect in a direction
	Idle GeneKind = iota
	North
	South
	East
	West

	GeneCount
)

//go:generate stringer -type=GeneKind

func GeneFromKind(g GeneKind) Gene {
	switch g {
	// MovementGene
	case Idle:
		fallthrough
	case North:
		fallthrough
	case South:
		fallthrough
	case East:
		fallthrough
	case West:
		g := NewMovementGene(g)
		return &g

	default:
		panic(fmt.Sprintf("Cannot derive gene from kind %v!", g))
	}
}

type MovementGene struct {
	kind GeneKind
}

func NewMovementGene(g GeneKind) MovementGene {
	return MovementGene{g}
}

func (g *MovementGene) Kind() GeneKind { return g.kind }

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func (g *MovementGene) Perform(s Simulation, n Body) {
	offsetNode := func(x int, y int) {
		x2, y2 := n.Position()
		xo, yo := Max(0, Min(x2+x, SimulationWidth-1)), Max(0, Min(y2+y, SimulationHeight-1))

		n.SetPosition(xo, yo)
	}

	switch g.kind {
	case Idle:
		break
	case North:
		offsetNode(0, 1)
	case South:
		offsetNode(0, -1)
	case East:
		offsetNode(1, 0)
	case West:
		offsetNode(-1, 0)
	default:
		panic("Bad gene kind in movement gene")
	}
}

package gene

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
)

type MovementGene struct {
	kind GeneKind
}

func NewMovementGene(g GeneKind) MovementGene {
	return MovementGene{g}
}

func (g *MovementGene) Kind() GeneKind { return g.kind }

func (g *MovementGene) Perform(s Simulation, n Body) {
	offsetNode := func(x int, y int) {
		x2, y2 := n.Position()
		n.SetPosition(x2+x, y2+y)
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

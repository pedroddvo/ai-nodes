package gene

const (
	SimulationHeight = 16
	SimulationWidth  = 16
)

type Body interface {
	Position() (int, int)
	SetPosition(x, y int)
}

// Bottom simulation interface
// As genes will be performing side effects, they must perform it inside a boxed simulation
type Simulation interface {
	// Current simulation tick
	CurrentTick() int64
}

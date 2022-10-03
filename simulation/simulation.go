package simulation

const (
	SimulationHeight = 16
	SimulationWidth  = 16
)

// A body represents a simple entity in the simulation with a position
type Body struct {
	mechanism Mechanism

	x, y int
}

// A simulation is a sandboxed environment where each body can interact and change state
type Simulation struct {
	tick int64

	bodies []*Body
}

func NewBody(mechanism Mechanism, x int, y int) Body {
	return Body{mechanism: mechanism, x: x, y: y}
}

func NewSimulation(bodies []*Body) Simulation {
	return Simulation{tick: 0, bodies: bodies}
}

func (s *Body) Position() (int, int) {
	return s.x, s.y
}

func (s *Body) SetPosition(x, y int) {
	s.x = x
	s.y = y
}

func (s *Simulation) Tick() int64 {
	return s.tick
}

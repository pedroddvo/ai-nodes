package simulation

import (
	"math/rand"
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

func (s *Simulation) CurrentTick() int64 {
	return s.tick
}

// Activate a mechanism to potentially perform a side effect and change states
func (b *Body) Activate(sim *Simulation) {
	// Check if the condition is right
	var validConditions []int
	for i, c := range b.mechanism.currentState.conditions {
		if c.Determine(sim) {
			validConditions = append(validConditions, i)
		}
	}

	// If any conditions are right, pick a random condition then induce a side effect
	if len(validConditions) > 0 {
		connection := rand.Intn(len(validConditions))
		b.mechanism.currentState.effects[connection].Perform(sim, b)

		// ...And then change state
		b.mechanism.currentState = b.mechanism.currentState.connections[connection]
	}

}

// Simulate one tick
func (s *Simulation) Simulate() {
	for _, b := range s.bodies {
		b.Activate(s)
	}

	s.tick += 1
}

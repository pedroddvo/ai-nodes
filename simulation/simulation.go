package simulation

import (
	"math/rand"
	"sync"
)

const reproductionCooldown = 20
const lifeTime = 100

// A body represents a simple entity in the simulation with a position
type Body struct {
	mechanism Mechanism

	deathTimer           int
	reproductionCooldown int

	x, y int
}

// A simulation is a sandboxed environment where each body can interact and change state
type Simulation struct {
	Tick int64

	bodies []*Body
}

func NewBody(mechanism Mechanism, x int, y int) Body {
	return Body{mechanism: mechanism, x: x, y: y, reproductionCooldown: reproductionCooldown, deathTimer: lifeTime}
}

func NewSimulation(bodies []*Body) Simulation {
	return Simulation{Tick: 0, bodies: bodies}
}

func (s *Body) Position() (int, int) {
	return s.x, s.y
}

func (s *Body) SetPosition(x, y int) {
	s.x = x
	s.y = y
}

func (s *Simulation) CurrentTick() int64 {
	return s.Tick
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

func remove(slice []*Body, s int) []*Body {
	return append(slice[:s], slice[s+1:]...)
}

// Simulate one tick
func (s *Simulation) Simulate() {
	// Each body acts asynchronously
	var wg sync.WaitGroup
	for _, b := range s.bodies {
		wg.Add(1)
		go func(b *Body) {
			b.Activate(s)
			if b.reproductionCooldown > 0 {
				b.reproductionCooldown -= 1
			}

			b.deathTimer -= 1
			wg.Done()
		}(b)
	}
	wg.Wait()

	var deathList []int
	var children []*Body

	// Reproduction & Death
	for i, a := range s.bodies {
		if a.deathTimer < 0 {
			deathList = append(deathList, i)
		}
		for _, b := range s.bodies {
			if a != b && a.reproductionCooldown == 0 && b.reproductionCooldown == 0 && a.x == b.x && a.y == b.y {
				c := Reproduce(a, b)
				children = append(children, &c)

				a.reproductionCooldown = reproductionCooldown
				b.reproductionCooldown = reproductionCooldown
			}
		}
	}

	for c, i := range deathList {
		s.bodies = remove(s.bodies, c-i)
	}

	s.bodies = append(s.bodies, children...)

	s.Tick += 1
}

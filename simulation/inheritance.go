package simulation

import (
	"math/rand"

	"github.com/pedroddvo/ai-nodes/gene"
)

const mutationChance = 0.75

// Create a map with all states spread
func (s *State) States(visited map[*State]int) map[*State]int {
	if _, ok := visited[s]; ok {
		return visited
	}

	visited[s] = len(visited)

	for _, c := range s.connections {
		if _, ok := visited[c]; ok {
			continue
		}

		visited = c.States(visited)
	}

	return visited
}

func Min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func (s *State) Emplace(b *State, visited []*State) {
	for _, v := range visited {
		if s == v {
			return
		}
	}

	visited = append(visited, s)

	// Swap conditions and effects
	for i := 0; i < Min(len(s.conditions), len(b.conditions)); i++ {
		s.conditions[i] = b.conditions[i]
		if rand.Float32() < mutationChance {
			s.conditions[i] = GenerateCondition()
		}
	}

	for i := 0; i < Min(len(s.effects), len(b.effects)); i++ {
		s.effects[i] = b.effects[i]
		if rand.Float32() < mutationChance {
			s.effects[i] = GenerateGene()
		}
	}

	for i := 0; i < Min(len(s.connections), len(b.connections)); i++ {
		s.connections[i].Emplace(b.connections[i], visited)
	}
}

// Two bodies reproduce to produce a new body, combining both mechanisms into a new mechanism
// this procedure introduces potential mutations to the mechanisms in question
func Reproduce(a *Body, b *Body) Body {
	as := a.mechanism.zeroState
	bs := b.mechanism.zeroState

	ac := as.States(map[*State]int{})
	bc := bs.States(map[*State]int{})

	bigger, smaller := as, bs
	if len(bc) > len(ac) {
		bigger, smaller = bs, as
	}

	cpy := State{
		conditions:  make([]gene.Condition, len(bigger.conditions)),
		effects:     make([]gene.Gene, len(bigger.effects)),
		connections: make([]*State, len(bigger.connections)),
	}

	copy(cpy.conditions, bigger.conditions)
	copy(cpy.effects, bigger.effects)
	copy(cpy.connections, bigger.connections)

	cpy.Emplace(smaller, []*State{})

	x := NewBody(MechanismFromState(&cpy), a.x, a.y)
	return x
}

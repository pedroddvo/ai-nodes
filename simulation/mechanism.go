package simulation

import (
	"github.com/pedroddvo/ai-nodes/gene"
)

// A state of a mechanism does not perform any side effects on its own
// instead, it allows effects to bridge between connections on a certain condition
// Each effect connects to a connection hence len(effects) == len(connections) == len(conditions)
// States can be identified by their pointers
type State struct {
	conditions  []gene.Condition
	effects     []gene.Gene
	connections []*State
}

// A mechanism is a state machine where side effects are dictated by a gene
// In itself, it is an entry point to the state machine
// Mechanisms can be generated and mutated
type Mechanism struct {
	currentState *State
	zeroState    State // Entry point for state machine
}

func MechanismFromState(s State) Mechanism {
	return Mechanism{zeroState: s, currentState: &s}
}

// Creates a bridge from a->b
func (a *State) Bridge(b *State, effect gene.Gene, condition gene.Condition) {
	a.connections = append(a.connections, b)
	a.conditions = append(a.conditions, condition)
	a.effects = append(a.effects, effect)
}

// Activate a state to potentially perform a side effect and change states
func (s *State) Activate(sim *Simulation) {

}

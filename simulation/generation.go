package simulation

import (
	"math/rand"

	"github.com/pedroddvo/ai-nodes/gene"
)

func GenerateGene() gene.Gene {
	return gene.GeneFromKind(gene.GeneKind(rand.Intn(int(gene.GeneCount))))
}

func GenerateCondition() gene.Condition {
	return gene.ConditionFromKind(gene.ConditionKind(rand.Intn(int(gene.ConditionCount))))
}

func GenerateMechanism(maxDepth int) Mechanism {
	depth := 1 + rand.Intn(maxDepth-1)
	states := make([]*State, depth)

	for i := 0; i < depth; i++ {
		conCount := 1 + rand.Intn(maxDepth-1)

		states[i] = &State{
			conditions:  make([]gene.Condition, conCount),
			effects:     make([]gene.Gene, conCount),
			connections: make([]*State, conCount),
		}

	}

	for _, s := range states {
		for i := 0; i < len(s.connections); i++ {
			s.connections[i] = states[rand.Intn(depth)]
			s.conditions[i] = GenerateCondition()
			s.effects[i] = GenerateGene()
		}
	}

	return MechanismFromState(states[0])
}

package main

import (
	"fmt"

	"github.com/pedroddvo/ai-nodes/gene"
	"github.com/pedroddvo/ai-nodes/simulation"
)

func main() {
	a, b := simulation.State{}, simulation.State{}

	g := gene.NewMovementGene(gene.East)

	a.Bridge(&b, &g, &gene.DummyCondition{})
	b.Bridge(&a, &g, &gene.DummyCondition{})

	ba, bb := simulation.NewBody(simulation.MechanismFromState(a), 4, 4), simulation.NewBody(simulation.MechanismFromState(a), 4, 4)

	s := simulation.NewSimulation([]*simulation.Body{&ba, &bb})

	fmt.Println(s.Pretty())
}

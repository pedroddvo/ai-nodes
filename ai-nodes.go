package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pedroddvo/ai-nodes/gene"
	"github.com/pedroddvo/ai-nodes/simulation"
)

func main() {
	rand.Seed(time.Now().Unix())
	var bodies []*simulation.Body

	for i := 0; i < 8; i++ {
		ma := simulation.GenerateMechanism(5)
		ba := simulation.NewBody(ma, rand.Intn(gene.SimulationWidth), rand.Intn(gene.SimulationHeight))
		bodies = append(bodies, &ba)
	}

	s := simulation.NewSimulation(bodies)

	fmt.Println(s.Pretty(simulation.PrettyOpts{World: true, States: true}))
	for {
		var input string
		fmt.Scanln(&input)
		if input == "s" {
			fmt.Println(s.Pretty(simulation.PrettyOpts{States: true}))
			continue
		} else {
			s.Simulate()
			fmt.Println(s.Pretty(simulation.PrettyOpts{World: true}))
		}
	}
}

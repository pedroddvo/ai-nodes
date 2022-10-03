package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/pedroddvo/ai-nodes/simulation"
)

func main() {
	rand.Seed(time.Now().Unix())
	var bodies []*simulation.Body

	for i := 0; i < 10; i++ {
		ma := simulation.GenerateMechanism(5)
		ba := simulation.NewBody(ma, 5, 5)
		bodies = append(bodies, &ba)
	}

	s := simulation.NewSimulation(bodies)

	fmt.Println(s.Pretty(simulation.PrettyOpts{World: true, States: true}))
	for {
		fmt.Scanln()
		s.Simulate()
		fmt.Println(s.Pretty(simulation.PrettyOpts{World: true}))
	}
}

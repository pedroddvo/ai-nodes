package simulation

import (
	"fmt"
	"strings"

	"github.com/pedroddvo/ai-nodes/gene"
)

type PrettyOpts struct {
	World, States bool
}

func mapConnections(s *State, connectionMap map[*State]int) {
	if _, ok := connectionMap[s]; !ok {
		connectionMap[s] = len(connectionMap)
	} else {
		return
	}
	for _, con := range s.connections {
		mapConnections(con, connectionMap)
	}
}

func prettyConnections(connectionMap map[*State]int) string {
	if len(connectionMap) == 0 {
		return ""
	}

	buf := ""

	for s, id := range connectionMap {
		buf += fmt.Sprintf("s%d: ", id)
		if len(s.connections) == 0 {
			buf += "⭗" // empty state (end of state machine)
		}
		buf += "\n"
		for i, con := range s.connections {
			// Stretch the arrow to fit kind id
			k := fmt.Sprintf("%v", s.effects[i].Kind())
			buf += fmt.Sprintf("|    %v\n", s.effects[i].Kind())
			buf += fmt.Sprintf("|  s%d%s>s%d\n", id, repeat(len(k), "-"), connectionMap[con])
			buf += fmt.Sprintf("|  => %v\n\n", s.conditions[i].ConditionKind())
		}
	}

	return buf
}

func repeat(l int, s string) string {
	buf := ""
	for i := 0; i < l; i++ {
		buf += s
	}
	return buf
}

func (self *State) Pretty() string {
	connectionMap := make(map[*State]int)

	mapConnections(self, connectionMap)

	return prettyConnections(connectionMap)
}

func (self *Mechanism) Pretty() string {
	return self.zeroState.Pretty()
}

func (self *Simulation) Pretty(opts PrettyOpts) string {
	bodyIds := make(map[*Body]int)

	for _, b := range self.bodies {
		bodyIds[b] = len(bodyIds)
	}

	buf := ""
	char := byte('a')
	clumps := make(map[byte][]*Body)
	// map out simulation
	if opts.World {
		buf += " " + repeat(gene.SimulationWidth, "-") + "\n"
		for y := 0; y < gene.SimulationHeight; y++ {
			buf += "|"
			for x := 0; x < gene.SimulationWidth; x++ {
				var clump []*Body
				for b := range bodyIds {
					if b.x == x && b.y == y {
						clump = append(clump, b)
					}
				}

				if len(clump) > 1 {
					buf += string(char)
					clumps[char] = clump
					char++
				} else if len(clump) == 1 {
					buf += fmt.Sprint(bodyIds[clump[0]])
				} else {
					buf += " "
				}
			}
			buf += "|\n"
		}
		buf += " " + repeat(gene.SimulationWidth, "-")
	}

	// clumps
	if len(clumps) > 0 {
		buf += "\n"
		for c, bs := range clumps {
			buf += string(c) + ": "
			for _, b := range bs {
				buf += fmt.Sprint(bodyIds[b]) + " "
			}
		}
	}

	buf += "\n\n"

	if opts.States {
		// pretty print states
		buf += "-------- INDEX -------\n"
		for b, i := range bodyIds {
			buf += fmt.Sprint(i) + ":\n"
			buf += indent(1, b.mechanism.zeroState.Pretty())

			buf += "\n"
		}
	}
	return buf
}

func indent(l int, s string) string {
	i := repeat(l, " ")
	return i + strings.ReplaceAll(s, "\n", "\n"+i)
}

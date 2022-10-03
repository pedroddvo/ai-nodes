package gene

import "fmt"

type Condition interface {
	ConditionKind() ConditionKind

	Determine(s Simulation) bool
}

type ConditionKind int

const (
	NoCondition ConditionKind = iota
	TickIsEven

	ConditionCount
)

//go:generate stringer -type=ConditionKind

func ConditionFromKind(c ConditionKind) Condition {
	switch c {
	case NoCondition:
		return &NoConditionCondition{}
	case TickIsEven:
		return &TickIsEvenCondition{}
	default:
		panic(fmt.Sprintf("Cannot derive condition from kind %v!", c))
	}
}

type NoConditionCondition struct{}

func (d *NoConditionCondition) ConditionKind() ConditionKind { return NoCondition }
func (d *NoConditionCondition) Determine(s Simulation) bool  { return true }

type TickIsEvenCondition struct{}

func (d *TickIsEvenCondition) ConditionKind() ConditionKind { return TickIsEven }
func (d *TickIsEvenCondition) Determine(s Simulation) bool {
	return s.CurrentTick()%2 == 0
}

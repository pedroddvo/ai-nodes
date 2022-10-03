package gene

type Condition interface {
	Description() string

	Determine(s Simulation) bool
}

type DummyCondition struct{}

func (d *DummyCondition) Description() string { return "tick is odd" }

func (d *DummyCondition) Determine(s Simulation) bool {
	return s.Tick()%2 == 0
}

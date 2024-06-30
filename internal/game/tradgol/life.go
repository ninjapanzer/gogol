package tradgol

// LifeState Represents a cell in the game of life.
type LifeState struct {
	// Whether the cell is alive or dead.
	alive bool
}

type Life interface {
	State() bool
	SetState(bool)
}

func NewLife(b bool) *LifeState {
	return &LifeState{b}
}

func (l *LifeState) State() bool {
	return l.alive
}

func (l *LifeState) SetState(state bool) {
	l.alive = state
}

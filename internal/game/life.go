package game

// LifeState Represents a cell in the game of life.
type Life interface {
	State() bool
	SetState(bool)
}

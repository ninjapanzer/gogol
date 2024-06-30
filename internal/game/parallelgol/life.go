package parallelgol

import "github.com/ninjapanzer/gogol/internal/game"

// LifeState Represents a cell in the game of life.
type LifeState struct {
	// Whether the cell is alive or dead.
	alive     bool
	neighbors []*Neighbor
}

type Neighbor struct {
	position int
	depth    int
	cell     game.Life
}

const (
	TopLeft = iota + 1
	Top
	TopRight
	Left
	Right
	BottomLeft
	Bottom
	BottomRight
)

type Life interface {
	State() bool
	SetState(bool)
	RegisterNeighbor(*Neighbor)
	RegisterNeighbors([]*Neighbor)
	Neighbors() []*Neighbor
}

func NewNeighbor(position int, depth int, cell Life) *Neighbor {
	return &Neighbor{position, depth, cell}
}

func NewLife(b bool) Life {
	return &LifeState{b, nil}
}

func (l *Neighbor) Depth() int {
	return l.depth
}

func (l *LifeState) Neighbors() []*Neighbor {
	return l.neighbors
}

func (l *LifeState) RegisterNeighbor(cell *Neighbor) {
	l.neighbors = append(l.neighbors, cell)
}

func (l *LifeState) RegisterNeighbors(cells []*Neighbor) {
	l.neighbors = cells
}

func (l *LifeState) State() bool {
	return l.alive
}

func (l *LifeState) SetState(state bool) {
	var aliveNeighbors int
	for n := range l.neighbors {
		if l.neighbors[n].cell.State() {
			aliveNeighbors++
		}
	}

	if l.alive {
		if aliveNeighbors < 2 || aliveNeighbors > 3 {
			l.alive = false
		}
	} else {
		if aliveNeighbors == 3 {
			l.alive = true
		}
	}
}

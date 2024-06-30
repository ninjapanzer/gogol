package tradgol

import "github.com/ninjapanzer/gogol/internal/game"

type WorldState struct {
	cells [][]game.Life
}

type World interface {
	Bootstrap()
	ComputeState()
	Cells() [][]game.Life
}

func (w *WorldState) SetCells(cells [][]game.Life) {
	w.cells = cells
}

func (w *WorldState) Cells() [][]game.Life {
	return w.cells
}

func (w *WorldState) ComputeState() {
	for y, row := range w.cells {
		for x, cell := range row {
			aliveNeighbors := 0
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					if i == 0 && j == 0 {
						continue
					}
					if y+i < 0 || y+i >= len(w.cells) {
						continue
					}
					if x+j < 0 || x+j >= len(w.cells[y+i]) {
						continue
					}
					if w.cells[y+i][x+j] != nil && w.cells[y+i][x+j].State() {
						aliveNeighbors++
					}
				}
			}
			if cell.State() {
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					cell.SetState(false)
				}
			} else {
				if aliveNeighbors == 3 {
					cell.SetState(true)
				}
			}
		}
	}
}

func NewWorld(width, height int) World {
	w := &WorldState{}
	w.cells = make([][]game.Life, height)
	for i := range w.cells {
		w.cells[i] = make([]game.Life, width)
		for j := range w.cells[i] {
			w.cells[i][j] = NewLife(false) // Initialize each cell with a Life object with Alive set to false
		}
	}
	return w
}

// Bootstrap populate the world with a glider
func (w *WorldState) Bootstrap() {
	w.cells[0][1] = NewLife(true)
	w.cells[1][2] = NewLife(true)
	w.cells[2][0] = NewLife(true)
	w.cells[2][1] = NewLife(true)
	w.cells[2][2] = NewLife(true)
	w.cells[2][3] = NewLife(true)
	w.cells[2][4] = NewLife(true)
	w.cells[4][1] = NewLife(true)
	w.cells[4][2] = NewLife(true)
	w.cells[4][0] = NewLife(true)
	w.cells[4][1] = NewLife(true)
	w.cells[4][2] = NewLife(true)
	w.cells[4][3] = NewLife(true)
	w.cells[4][4] = NewLife(true)
	w.cells[5][1] = NewLife(true)
	w.cells[5][2] = NewLife(true)
	w.cells[5][0] = NewLife(true)
	w.cells[5][1] = NewLife(true)
	w.cells[5][2] = NewLife(true)
	w.cells[5][3] = NewLife(true)
	w.cells[5][4] = NewLife(true)
}

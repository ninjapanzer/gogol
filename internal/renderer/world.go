package renderer

import "github.com/ninjapanzer/gogol/internal/game"

type World struct {
	matrix  game.WorldState
	display Renderer
}

type WorldRenderer interface {
	Refresh()
}

func NewWorld(world *game.WorldState, display *Renderer) *World {
	w := &World{}
	w.matrix = *world
	w.display = *display
	return w
}

// loops through the matrix and draws the cells that are alive as 0 and dead as nothing
func (w *World) Refresh() {
	for y, row := range w.matrix.Cells() {
		for x, cell := range row {
			if cell.State() {
				w.display.Display.MovePrint(y, x, "0")
			} else {
				w.display.Display.MovePrint(y, x, "-")
			}
		}
	}
	w.display.Display.Refresh()
}

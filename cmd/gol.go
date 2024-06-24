package main

import (
	"github.com/ninjapanzer/gogol/internal/game"
	"github.com/ninjapanzer/gogol/internal/renderer"
	"time"
)

func main() {
	r := renderer.NewShellRenderer()
	defer r.End()
	r.Beep()
	r.Draw("Hello, world.go!")
	//r.Display.GetChar()

	y, x := r.Dimensions()
	world := game.NewWorld(x, y)

	world.Bootstrap()

	worldRenderer := renderer.NewWorld(world, r)

	for {
		time.Sleep(100 * time.Millisecond)
		world.ComputeState()
		r.Display.Clear()
		worldRenderer.Refresh()
		r.Display.Refresh()
	}
}

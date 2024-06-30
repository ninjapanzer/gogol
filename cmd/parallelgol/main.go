package main

import (
	"github.com/ninjapanzer/gogol/internal/game/parallelgol"
	"github.com/ninjapanzer/gogol/internal/renderer"
)

func main() {
	r := renderer.NewShellRenderer()
	defer r.End()
	r.Beep()
	r.Draw("Hello, world.go!")
	//r.Display.GetChar()

	y, x := r.Dimensions()
	world := parallelgol.NewWorld(x, y)

	world.Bootstrap()
	worldRenderer := renderer.NewWorld(&world, r)
}

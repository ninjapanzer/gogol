package main

import (
	"context"
	"github.com/ninjapanzer/gogol/internal/game/tradgol"
	"github.com/ninjapanzer/gogol/internal/renderer"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancelChan := make(chan os.Signal, 1)
	signal.Notify(cancelChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		r := renderer.NewShellRenderer()
		defer r.End()
		r.Beep()
		r.Draw("Hello, world.go!")
		//r.Display.GetChar()

		y, x := r.Dimensions()
		world := tradgol.NewWorld(x, y)

		world.Bootstrap()

		worldRenderer := renderer.NewWorld(world, *r)

		for {
			select {
			case <-ctx.Done():
				r.End()
				return
			default:
				time.Sleep(100 * time.Millisecond)
				world.ComputeState()
				r.Display.Clear()
				worldRenderer.Refresh()
				r.Display.Refresh()
			}
		}
	}()

	select {
	case <-cancelChan:
		cancel()
		println("Cancelled")
	case <-ctx.Done():
		println("Done")
	}
}

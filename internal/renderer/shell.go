package renderer

import (
	"github.com/gbin/goncurses"
)

// ShellRenderer is a renderer that outputs to the terminal.
type ShellRenderer interface {
	Beep()
	Draw(string)
	Dimensions() (int, int)
	Start()
	End()
}

type Renderer struct {
	screen  *goncurses.Window
	wrapper *goncurses.Window
	Display *goncurses.Window
	Padding int
}

func (s *Renderer) Start() {
	y, x := s.Dimensions()
	if s.wrapper != nil {
		if err := s.wrapper.Delete(); err != nil {
			panic(err)
		}
	}
	w, err := goncurses.NewWindow(y, x, 0, 0)
	if err != nil {
		panic(err)
	}
	if err := w.Box('|', '-'); err != nil {
		panic(err)
	}
	s.wrapper = w

	if s.Display != nil {
		if err := s.Display.Delete(); err != nil {
			panic(err)
		}
	}

	d, err := goncurses.NewWindow(y-(s.Padding*2), x-(s.Padding*2), s.Padding, s.Padding)
	if err != nil {
		panic(err)
	}
	s.Display = d

	s.wrapper.Refresh()
	s.Display.Refresh()
}

func (s *Renderer) End() {
	goncurses.End()
}

func end() {
	//TODO implement me
	panic("implement me")
}

func (s *Renderer) Dimensions() (int, int) {
	return s.screen.MaxYX()
}

func (s *Renderer) Draw(str string) {
	s.Display.Println(str)
}

func (s *Renderer) Beep() {
	goncurses.Beep()
}

// NewShellRenderer creates a new ShellRenderer.
func NewShellRenderer() *Renderer {
	screen, err := goncurses.Init()
	if err != nil {
		panic(err)
	}

	s := Renderer{
		screen:  screen,
		Padding: 2,
	}
	s.Start()
	return &s
}

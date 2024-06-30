package game

type World interface {
	Bootstrap()
	Cells() [][]Life
}

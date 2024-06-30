package parallelgol

type WorldState struct {
	startingCell Life
	maxWidth     int
	maxHeight    int
}

type World interface {
	Bootstrap()
	Cells() [][]Life
}

func NewWorld(width, height int) World {
	startingCell := NewLife(true)
	startingCell.RegisterNeighbors(
		[]*Neighbor{
			NewNeighbor(TopLeft, 1, NewLife(false)),
			NewNeighbor(Top, 1, NewLife(false)),
			NewNeighbor(TopRight, 1, NewLife(false)),
			NewNeighbor(Left, 1, NewLife(false)),
			NewNeighbor(Right, 1, NewLife(false)),
			NewNeighbor(BottomLeft, 1, NewLife(false)),
			NewNeighbor(Bottom, 1, NewLife(false)),
			NewNeighbor(BottomRight, 1, NewLife(false)),
		})

	return &WorldState{
		maxWidth:     width,
		maxHeight:    height,
		startingCell: startingCell,
	}
}

func (w WorldState) Bootstrap() {
	for c := range w.startingCell.Neighbors() {
		expandNeighborhood(w.startingCell, w.startingCell.Neighbors()[c])
		//if w.startingCell.Neighbors()[c].cell.Depth() < 2 {
		//	expandNeighborhood(w.startingCell, w.startingCell.Neighbors()[c].cell)
		//}
	}
}

func expandNeighborhood(source Life, neighbor *Neighbor) {
	for c := range neighbor.cell.Neighbors() {
		n := neighbor.cell.Neighbors()[c]
		depth := n.Depth() + 1
		switch n.position {
		case TopLeft:
			n.cell.RegisterNeighbor(NewNeighbor(BottomRight, depth, source))
		case Top:
			n.cell.RegisterNeighbor(NewNeighbor(Bottom, depth, source))
			for c := range source.Neighbors() {
				if source.Neighbors()[c].position == TopRight {
					n.cell.RegisterNeighbor(NewNeighbor(BottomRight, depth, source.Neighbors()[c].cell))
				}
				if source.Neighbors()[c].position == TopLeft {
					n.cell.RegisterNeighbor(NewNeighbor(BottomLeft, depth, source))
				}
			}
		case TopRight:
			n.cell.RegisterNeighbor(NewNeighbor(BottomLeft, depth, source))
		case Left:
			n.cell.RegisterNeighbor(NewNeighbor(Right, depth, source))
			for c := range source.Neighbors() {
				if source.Neighbors()[c].position == TopRight {
					n.cell.RegisterNeighbor(NewNeighbor(TopLeft, depth, source.Neighbors()[c].cell))
				}
				if source.Neighbors()[c].position == BottomRight {
					n.cell.RegisterNeighbor(NewNeighbor(BottomLeft, depth, source.Neighbors()[c].cell))
				}
			}
		case Right:
			n.cell.RegisterNeighbor(NewNeighbor(Left, depth, source))
			for c := range source.Neighbors() {
				if source.Neighbors()[c].position == TopRight {
					n.cell.RegisterNeighbor(NewNeighbor(TopLeft, depth, source.Neighbors()[c].cell))
				}
				if source.Neighbors()[c].position == BottomRight {
					n.cell.RegisterNeighbor(NewNeighbor(BottomLeft, depth, source.Neighbors()[c].cell))
				}
			}
		case BottomLeft:
			n.cell.RegisterNeighbor(NewNeighbor(TopRight, depth, source))
		case Bottom:
			n.cell.RegisterNeighbor(NewNeighbor(Top, depth, source))
			for c := range source.Neighbors() {
				if source.Neighbors()[c].position == BottomRight {
					n.cell.RegisterNeighbor(NewNeighbor(TopRight, depth, source.Neighbors()[c].cell))
				}
				if source.Neighbors()[c].position == BottomLeft {
					n.cell.RegisterNeighbor(NewNeighbor(TopLeft, depth, source.Neighbors()[c].cell))
				}

			}
		case BottomRight:
			n.cell.RegisterNeighbor(NewNeighbor(TopLeft, depth, source))
		}
	}

	foundPositions := make(map[int]bool)
	for c := range neighbor.cell.Neighbors() {
		foundPositions[neighbor.cell.Neighbors()[c].position] = true
	}

	for p := range foundPositions {
		if !foundPositions[p] {
			neighbor.cell.RegisterNeighbor(NewNeighbor(p, neighbor.depth+1, NewLife(false)))
		}
	}
}

func (w WorldState) Cells() [][]Life {
	//TODO implement me
	panic("implement me")
}

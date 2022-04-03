package grid

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/faces/grid"
	"log"
)

const (
	CELL_SIZE = 16
)

type Grid struct {
	storage map[int]map[int]grid.IGridCell
}

func NewGrid() *Grid {
	return &Grid{
		storage: map[int]map[int]grid.IGridCell{},
	}
}

func (g *Grid) GetCellAtPos(vec2 common.Vec2) grid.IGridCell {
	x, y := vec2.Scale(1 / CELL_SIZE).XYint()
	return g.GetCellAt((x), (y))
}

func (g *Grid) GetCellAt(x, y int) grid.IGridCell {
	if g.storage[x] == nil {
		g.storage[x] = map[int]grid.IGridCell{}
	}

	cell := g.storage[x][y]
	if cell == nil {
		cell = &GridCell{
			X: x,
			Y: y,
		}
		g.storage[x][y] = cell
	}

	return cell
}

func (g *Grid) GetCellAtSafe(x, y int) grid.IGridCell {
	if g.storage[x] == nil {
		g.storage[x] = map[int]grid.IGridCell{}
	}
	return g.storage[x][y]
}

func (g *Grid) SetCellAt(x, y int, cell grid.IGridCell) {
	if g.storage[x] == nil {
		g.storage[x] = map[int]grid.IGridCell{}
	}

	gc, ok := cell.(*GridCell)
	if !ok {
		log.Println("Non-GridCell is passed")
		return
	}
	gc.X, gc.Y = x, y
	g.storage[x][y] = gc
}

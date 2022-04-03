package pathing

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/faces/grid"
	"github.com/beefsack/go-astar"
	"log"
	"math"
)

type Node struct {
	length_left int
	grid        grid.IGrid
	int_vec
}

var nodes map[grid.IGridCell]*Node
var cells map[*Node]grid.IGridCell

func init() {
	nodes = map[grid.IGridCell]*Node{}
	cells = map[*Node]grid.IGridCell{}
}

func get_node_by_cell(grid grid.IGrid, cell grid.IGridCell) *Node {
	node := nodes[cell]
	if node != nil {
		return node
	}

	x, y := cell.GetPos()
	node = &Node{
		0,
		grid,
		int_vec{x, y},
	}
	nodes[cell] = node
	cells[node] = cell

	return node
}

func get_node_at_pos(grid grid.IGrid, pos common.Vec2) *Node {
	cell := grid.GetCellAtPos(pos)
	return get_node_by_cell(grid, cell)
}

func get_node_at_safe(grid grid.IGrid, x, y int) *Node {
	cell := grid.GetCellAtSafe(x, y)
	if cell != nil {
		return get_node_by_cell(grid, cell)
	}
	return nil
}

func FindWay(grid grid.IGrid, from, to common.Vec2, max_length int) []common.Vec2 {
	from_cell := get_node_at_pos(grid, from)
	to_cell := get_node_at_pos(grid, to)
	from_cell.length_left = max_length

	path, _, found := astar.Path(from_cell, to_cell)
	if found {
		points := make([]common.Vec2, len(path))
		for i, pather := range path {
			node, ok := pather.(*Node)
			if !ok {
				log.Fatalln("non-Node")
			}
			points[i] = node.GetWorldPos()
		}
	}

	return nil
}

type int_vec struct {
	x, y int
}

var n_shifts []int_vec

func init() {
	n_shifts = []int_vec{
		{-1, +0},
		{+1, +0},
		{-0, +1},
		{-0, -1},
	}
}

func (n *Node) GetWorldPos() common.Vec2 {
	return cells[n].GetPosWorld()
}

func (n *Node) PathNeighbors() []astar.Pather {
	if n.length_left == 0 {
		return nil
	}
	var pn []astar.Pather = nil

	for _, shift := range n_shifts {
		pos := int_vec{shift.x + n.x, shift.y + n.y}
		node := get_node_at_safe(n.grid, pos.x, pos.y)
		if node != nil {
			continue
		}
		node.length_left = n.length_left - 1
		pn = append(pn, node)
	}

	return pn
}

func (n *Node) PathNeighborCost(to astar.Pather) float64 {
	return 1
}

func (n *Node) PathEstimatedCost(to astar.Pather) float64 {
	t, ok := to.(*Node)
	if !ok {
		log.Println("got non-Node")
	}
	dx := math.Abs(float64(t.x - n.x))
	dy := math.Abs(float64(t.y - n.y))
	return dx + dy
}

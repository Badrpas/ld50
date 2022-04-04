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

type int_vec struct {
	x, y int
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

var start_pos_int int_vec

func FindWay(grid grid.IGrid, from, to common.Vec2, max_length int) []common.Vec2 {
	from_node := get_node_at_pos(grid, from)
	to_node := get_node_at_pos(grid, to)
	from_node.length_left = max_length
	start_pos_int = from_node.int_vec

	path, _, found := astar.Path(from_node, to_node)
	if found {
		point_count := len(path)
		points := make([]common.Vec2, point_count)
		for i, pather := range path {
			node, ok := pather.(*Node)
			if !ok {
				log.Fatalln("non-Node")
			}
			points[point_count-1-i] = node.GetWorldPos()
		}
		return points
	}

	return nil
}

var n_shifts []int_vec

func init() {
	n_shifts = []int_vec{
		{-1, +0},
		{+1, +0},
		{+0, +1},
		{+0, -1},
	}
}

func (n *Node) GetWorldPos() common.Vec2 {
	return cells[n].GetPosWorld()
}

func (n *Node) PathNeighbors() []astar.Pather {
	man_dist := n.manhat_dist_to(start_pos_int)

	//if n.length_left == 0 {
	//	//return nil
	//}
	if man_dist > 22 {
		return nil
	}
	//

	var pn []astar.Pather = nil

	for _, shift := range n_shifts {
		pos := int_vec{shift.x + n.x, shift.y + n.y}
		node := get_node_at_safe(n.grid, pos.x, pos.y)
		cell := cells[node]
		if node == nil || cell.GetHolder() != nil {
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
	return float64(t.manhat_dist_to(n.int_vec))
}

func (v int_vec) manhat_dist_to(o int_vec) int {
	dx := math.Abs(float64(o.x - v.x))
	dy := math.Abs(float64(o.y - v.y))
	return int(dx + dy)
}

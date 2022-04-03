package grid

import "github.com/badrpas/ld50/pkg/common"

type IGrid interface {
	GetCellAtPos(vec2 common.Vec2) IGridCell
	GetCellAt(x, y int) IGridCell
	GetCellAtSafe(x, y int) IGridCell
	SetCellAt(x, y int, cell IGridCell)
}

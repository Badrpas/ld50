package grid

import "github.com/badrpas/ld50/pkg/common"

type IGridCell interface {
	SetHolder(interface{})
	GetHolder() interface{}
	GetPos() (int, int)
	GetPosWorld() common.Vec2
}

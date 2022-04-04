package grid

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/config"
	"github.com/badrpas/ld50/pkg/entity"
	"log"
)

type GridCell struct {
	X, Y    int
	TakenBy interface{}
}

func (g *GridCell) GetHolder() interface{} {
	return g.TakenBy
}

func (c *GridCell) SetHolder(h interface{}) {
	if _, ok := h.(*entity.Entity); !ok && h != nil {
		log.Fatalln("Received non-Entity for GridCell.SetHolder()")
	}
	c.TakenBy = h
}

func (c *GridCell) GetPos() (int, int) {
	return c.X, c.Y
}

func (c *GridCell) GetPosWorld() common.Vec2 {
	return common.Vec2{
		float64(c.X*config.CELL_SIZE + 16),
		float64(c.Y*config.CELL_SIZE + 16),
	}
}

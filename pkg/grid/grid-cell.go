package grid

import (
	"github.com/badrpas/ld50/pkg/entity"
	"log"
)

type GridCell struct {
	Type    int
	TakenBy interface{}
}

func (g *GridCell) GetHolder() interface{} {
	return g.TakenBy
}

func (c *GridCell) SetHolder(h interface{}) {
	if _, ok := h.(*entity.Entity); !ok {
		log.Fatalln("Received non-Entity for GridCell.SetHolder()")
	}
	c.TakenBy = h
}

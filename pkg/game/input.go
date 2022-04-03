package game

import (
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"os"
)

func init_input(g *Game) {
	handler := &entity.Entity{
		Update: func(self *entity.Entity, dt float64) {
			if ebiten.IsKeyPressed(ebiten.KeyEscape) {
				os.Exit(0)
			}
		},
	}

	g.AddEntity(handler)
	g.SetUpdatePriority(handler, -10000)
}

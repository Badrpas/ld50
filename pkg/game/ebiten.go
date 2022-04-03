package game

import (
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

func (g *Game) Update() error {
	dt := 1. / 60. // Really disliking that

	for _, e := range g.update_order.roots {
		update_entity(e, dt)
	}

	return nil
}

func update_entity(e *entity.Entity, dt float64) {
	if e != nil {
		if e.Update != nil {
			e.Update(e, dt)
		}
		for _, child := range e.Children {
			update_entity(child, dt)
		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	cam_surface := g.Camera.Surface
	cam_surface.Clear()
	cam_surface.Fill(color.Black)

	z_count := len(g.zList)
	for i := 0; i < z_count; i++ {
		z := g.zList[i]
		for _, e := range g.zLevels[z] {
			if e != nil && e.Render != nil {
				e.Render(e, cam_surface)
			}
		}
	}

	g.Camera.Blit(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	c := g.Camera

	if c.Width != outsideWidth || c.Height != outsideHeight {
		c.Resize(outsideWidth, outsideHeight)
	}

	return outsideWidth, outsideHeight
}

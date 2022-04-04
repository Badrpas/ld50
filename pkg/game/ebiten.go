package game

import (
	"github.com/badrpas/ld50/pkg/controllers"
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"time"
)

var last_ms int64 = -1

func init() {
	last_ms = time.Now().UnixMilli()
}
func get_dt() float64 {
	now := time.Now().UnixMilli()
	if last_ms == -1 {
		last_ms = now - 40
	}
	dt := float64(now-last_ms) / 1000.0
	last_ms = now
	return dt
}

func (g *Game) Update() error {
	dt := get_dt()

	for _, e := range g.update_order.roots {
		update_entity(e, dt)
	}

	controllers.UpdatePhysics(dt)

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

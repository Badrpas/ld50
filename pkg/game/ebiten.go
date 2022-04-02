package game

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Update() error {
	dt := 1. / 60. // Really disliking that

	for _, e := range g.Entities {
		if e != nil && e.Update != nil {
			e.Update(e, dt)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	cam_surface := g.Camera.Surface
	cam_surface.Clear()

	for _, e := range g.Entities {
		if e != nil && e.Render != nil {
			e.Render(e, cam_surface)
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

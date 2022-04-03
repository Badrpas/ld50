package game

import (
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	camera "github.com/melonfunction/ebiten-camera"
)

func init_camera(g *Game) {
	width, height := ebiten.WindowSize()
	g.Camera = camera.NewCamera(width, height, 0, 0, 0, 1)

	const MouseMoveBtn = ebiten.MouseButtonRight

	sx, sy := 0, 0
	handler := &entity.Entity{
		Update: func(self *entity.Entity, dt float64) {
			cam_delta := dt * 400

			if ebiten.IsKeyPressed(ebiten.KeyA) {
				g.Camera.MovePosition(-cam_delta, 0)
			}
			if ebiten.IsKeyPressed(ebiten.KeyD) {
				g.Camera.MovePosition(cam_delta, 0)
			}
			if ebiten.IsKeyPressed(ebiten.KeyW) {
				g.Camera.MovePosition(0, -cam_delta)
			}
			if ebiten.IsKeyPressed(ebiten.KeyS) {
				g.Camera.MovePosition(0, cam_delta)
			}

			if inpututil.IsMouseButtonJustPressed(MouseMoveBtn) {
				sx, sy = ebiten.CursorPosition()
			}

			if ebiten.IsMouseButtonPressed(MouseMoveBtn) {
				x, y := ebiten.CursorPosition()
				dx, dy := -x+sx, -y+sy
				sx, sy = x, y
				g.Camera.MovePosition(float64(dx), float64(dy))
			}
		},
	}

	g.AddEntity(handler)
	g.SetUpdatePriority(handler, -10000+1)
}

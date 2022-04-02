package game

import (
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/hajimehoshi/ebiten/v2"
	camera "github.com/melonfunction/ebiten-camera"
)

func init_camera(g *Game) {
	width, height := ebiten.WindowSize()
	g.Camera = camera.NewCamera(width, height, 0, 0, 0, 1)

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
		},
	}
	g.AddEntity(handler)
}

package game

import (
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/hajimehoshi/ebiten/v2"
	camera "github.com/melonfunction/ebiten-camera"
)

type EntitiesStorage map[*entity.Entity]*entity.Entity

type Game struct {
	Entities EntitiesStorage
	Camera   *camera.Camera
}

func NewGame() *Game {
	width, height := ebiten.WindowSize()
	game := &Game{
		Entities: EntitiesStorage{},
		Camera:   camera.NewCamera(width, height, 0, 0, 0, 1),
	}

	return game
}

func (g *Game) AddEntity(e interface{}) {
	ent, ok := e.(*entity.Entity)
	if ok {
		g.Entities[ent] = ent
		ent.Game = g
	}
}

func (g *Game) RemoveEntity(e interface{}) {
	ent, ok := e.(*entity.Entity)
	if ok {
		delete(g.Entities, ent)
		ent.Game = nil
	}
}

func (g *Game) TranslateWithCamera(opts *ebiten.DrawImageOptions) {
	c := g.Camera
	w, h := c.Width, c.Height
	opts.GeoM.Translate(float64(w)/c.Scale/2, float64(h)/c.Scale/2)
	opts.GeoM.Translate(-c.X, -c.Y)
}

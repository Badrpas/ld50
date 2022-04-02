package game

import (
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/hajimehoshi/ebiten/v2"
	camera "github.com/melonfunction/ebiten-camera"
	"log"
)

type EntitiesStorage map[*entity.Entity]*entity.Entity

type Game struct {
	Entities EntitiesStorage
	Camera   *camera.Camera
}

func NewGame() *Game {
	game := &Game{
		Entities: EntitiesStorage{},
	}
	init_input(game)
	init_camera(game)

	return game
}

func (g *Game) AddEntity(e interface{}) {
	ent, ok := e.(*entity.Entity)
	if ok {
		g.Entities[ent] = ent
		ent.Game = g
	} else {
		log.Println("Received non entity in Game.AddEntity()")
	}
}

func (g *Game) RemoveEntity(e interface{}) {
	ent, ok := e.(*entity.Entity)
	if ok {
		delete(g.Entities, ent)
		ent.Game = nil
	} else {
		log.Println("Received non entity in Game.RemoveEntity()")
	}
}

func (g *Game) TranslateWithCamera(opts *ebiten.DrawImageOptions) {
	c := g.Camera
	w, h := c.Width, c.Height
	opts.GeoM.Translate(float64(w)/c.Scale/2, float64(h)/c.Scale/2)
	opts.GeoM.Translate(-c.X, -c.Y)
}

package game

import (
	"github.com/badrpas/ld50/pkg/entity"
	faces "github.com/badrpas/ld50/pkg/faces/grid"
	"github.com/badrpas/ld50/pkg/grid"
	"github.com/hajimehoshi/ebiten/v2"
	camera "github.com/melonfunction/ebiten-camera"
	"log"
)

type EntitiesStorage map[*entity.Entity]*entity.Entity

type Game struct {
	Entities EntitiesStorage
	Camera   *camera.Camera

	Grid *grid.Grid

	z_order
}

func NewGame() *Game {
	game := &Game{
		Entities: EntitiesStorage{},
		Grid:     grid.NewGrid(),
		z_order:  new_z_order(),
	}
	init_input(game)
	init_camera(game)

	return game
}

func (g *Game) AddEntitySafe(e *entity.Entity) {
	g.Entities[e] = e
	e.Game = g
	g.SetEntityZ(e, 0)
}

func (g *Game) AddEntity(e interface{}) {
	ent, ok := e.(*entity.Entity)
	if ok {
		g.AddEntitySafe(ent)
	} else {
		log.Println("Received non entity in Game.AddEntity()")
	}
}

func (g *Game) RemoveEntity(e interface{}) {
	ent, ok := e.(*entity.Entity)
	if ok {
		delete(g.Entities, ent)
		ent.Game = nil
		g.ClearEntityZ(ent)
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

func (g *Game) GetGrid() faces.IGrid {
	return g.Grid
}

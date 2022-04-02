package game

import (
	"github.com/badrpas/ld50/pkg/entity"
)

type EntitiesStorage map[*entity.Entity]*entity.Entity

type Game struct {
	Entities EntitiesStorage
}

func NewGame() *Game {
	game := &Game{
		Entities: EntitiesStorage{},
	}

	return game
}

func (g *Game) AddEntity(e interface{}) {
	ent, ok := e.(*entity.Entity)
	if ok {
		g.Entities[ent] = ent
	}
}

func (g *Game) RemoveEntity(e interface{}) {
	ent, ok := e.(*entity.Entity)
	if ok {
		delete(g.Entities, ent)
	}
}

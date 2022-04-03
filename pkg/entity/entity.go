package entity

import (
	"github.com/badrpas/ld50/pkg/faces/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	Game game.Game

	Update func(self *Entity, dt float64)
	Render func(self *Entity, screen *ebiten.Image)

	Heir interface{}
}

type EntitySet map[*Entity]*Entity

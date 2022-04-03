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

	Children []*Entity
	Parent   *Entity
}

type EntitySet map[*Entity]*Entity

func (e *Entity) AddChild(entity *Entity) {
	for _, child := range e.Children {
		if child == entity {
			return
		}
	}

	e.Children = append(e.Children, entity)
	entity.SetParent(e)
}

func (e *Entity) RemoveChild(entity *Entity) {
	for idx, child := range e.Children {
		if child == entity {
			e.Children = append(e.Children[:idx], e.Children[idx+1:]...)
			child.SetParent(nil)
		}
	}
}

func (e *Entity) SetParent(entity *Entity) {
	if e.Parent != nil {
		e.Parent.RemoveChild(e)
	}

	e.Parent = entity

	if e.Parent != nil {
		e.Parent.AddChild(e)
	}
}

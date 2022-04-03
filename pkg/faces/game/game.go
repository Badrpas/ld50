package game

import (
	"github.com/badrpas/ld50/pkg/faces/grid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game interface {
	ebiten.Game

	AddEntity(e interface{})
	RemoveEntity(e interface{})

	TranslateWithCamera(options *ebiten.DrawImageOptions)

	GetGrid() grid.IGrid
}

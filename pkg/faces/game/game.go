package game

import (
	"github.com/badrpas/ld50/pkg/faces/grid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
)

type Game interface {
	ebiten.Game

	AddEntity(e interface{})
	RemoveEntity(e interface{})

	EachEntity(func(e interface{}))

	TranslateWithCamera(options *ebiten.DrawImageOptions)

	GetGrid() grid.IGrid
	GetSpace() *resolv.Space
}

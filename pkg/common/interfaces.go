package common

import "github.com/hajimehoshi/ebiten/v2"

type Game interface {
	ebiten.Game

	AddEntity(e interface{})
	RemoveEntity(e interface{})

	TranslateWithCamera(options *ebiten.DrawImageOptions)
}

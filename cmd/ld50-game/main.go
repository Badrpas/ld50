package main

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/entities/sprite"
	"github.com/badrpas/ld50/pkg/game"
	_ "github.com/badrpas/ld50/pkg/img"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowResizable(true)

	e := sprite.NewSprite("check.png", common.Vec2{30, 30})
	g.AddEntity(e)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

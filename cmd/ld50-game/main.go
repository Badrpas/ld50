package main

import (
	"github.com/badrpas/ld50/pkg/game"
	_ "github.com/badrpas/ld50/pkg/img"
	gamemap "github.com/badrpas/ld50/pkg/map"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowResizable(true)

	gamemap.LoadMap("def.tmx", g)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

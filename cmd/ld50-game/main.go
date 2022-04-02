package main

import (
	"github.com/badrpas/ld50/pkg/game"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowResizable(true)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

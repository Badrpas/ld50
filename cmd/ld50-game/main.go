package main

import (
	"github.com/badrpas/ld50/pkg/game"
	_ "github.com/badrpas/ld50/pkg/imgrepo"
	gamemap "github.com/badrpas/ld50/pkg/map"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	g := game.NewGame()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowResizable(true)
	ebiten.SetWindowTitle("LD50")
	//ebiten.MaximizeWindow()

	err := gamemap.LoadMap("def.tmx", g)
	if err != nil {
		log.Fatal(err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

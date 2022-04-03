package main

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/entities/unit"
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

	for i := 0; i < 10; i++ {
		roach := unit.NewRoach(common.Vec2{float64(3000 + i*3), 3500})
		g.AddEntitySafe(roach.Entity)
		g.SetEntityZ(roach.Entity, 10)
	}

	roach2 := unit.NewRoach(common.Vec2{3104, 3104})
	g.AddEntitySafe(roach2.Entity)
	g.SetEntityZ(roach2.Entity, 10)

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

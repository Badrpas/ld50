package unit

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/entities/sprite"
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/badrpas/ld50/pkg/imgrepo"
	"github.com/hajimehoshi/ebiten/v2"
)

type Roach struct {
	*entity.Entity

	Sprite *sprite.Sprite
}

func NewRoach(pos common.Vec2) *Roach {
	image := imgrepo.ImgRepo["bug.png"]

	r := &Roach{
		Entity: nil,
		Sprite: sprite.NewSprite(image, pos),
	}
	r.Render = render
	r.Update = update
	r.Heir = r

	return r
}

func update(e *entity.Entity, dt float64) {
	roach, ok := e.Heir.(*Roach)
	if !ok {
		return
	}

	roach.Game.GetGrid().GetCellAtPos(roach.Sprite.Pos)
}

func render(e *entity.Entity, screen *ebiten.Image) {
	roach, ok := e.Heir.(*Roach)
	if ok {
		roach.Sprite.Render(e, screen)
	}
}

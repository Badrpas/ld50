package sprite

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	*entity.Entity

	Img *ebiten.Image
	Pos common.Vec2
}

func NewSprite(image *ebiten.Image, pos common.Vec2) *Sprite {
	s := &Sprite{
		&entity.Entity{
			Render: render,
		},
		image,
		pos,
	}

	s.Heir = s

	return s
}

func render(e *entity.Entity, screen *ebiten.Image) {
	if e.Game == nil {
		return
	}
	s, ok := e.Heir.(*Sprite)
	if !ok || s.Img == nil {
		return
	}

	opts := &ebiten.DrawImageOptions{}
	width, height := s.Img.Size()
	opts.GeoM.Translate(float64(width/-2), float64(height/-2))

	opts.GeoM.Translate(s.Pos.X, s.Pos.Y)

	e.Game.TranslateWithCamera(opts)

	screen.DrawImage(s.Img, opts)
}

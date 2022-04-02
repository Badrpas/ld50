package sprite

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/badrpas/ld50/pkg/img"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	*entity.Entity

	Img *ebiten.Image
	Pos common.Vec2
}

func NewSprite(img_name string, pos common.Vec2) *entity.Entity {
	s := &Sprite{
		&entity.Entity{
			Render: render,
		},
		img.ImgRepo[img_name],
		pos,
	}

	s.Heir = s

	return s.Entity
}

func render(e *entity.Entity, screen *ebiten.Image) {
	s, ok := e.Heir.(*Sprite)
	if !ok || s.Img == nil {
		return
	}

	opts := &ebiten.DrawImageOptions{}
	width, height := s.Img.Size()
	opts.GeoM.Translate(float64(width/2), float64(height/2))

	opts.GeoM.Translate(s.Pos.X, s.Pos.Y)

	screen.DrawImage(s.Img, opts)
}

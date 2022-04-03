package unit

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/controllers"
	"github.com/badrpas/ld50/pkg/entities/sprite"
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/badrpas/ld50/pkg/imgrepo"
	"github.com/hajimehoshi/ebiten/v2"
	"math/rand"
)

type Roach struct {
	*entity.Entity

	*sprite.Sprite

	target_pos common.Vec2
	Speed      float64

	controllers.VelocityComponent
}

func NewRoach(pos common.Vec2) *Roach {
	image := imgrepo.ImgRepo["bug.png"]

	r := &Roach{
		Entity:     &entity.Entity{},
		Sprite:     sprite.NewSprite(image, pos),
		target_pos: pos,
		Speed:      100,
	}
	r.Render = render
	r.Update = update
	r.Heir = r
	r.AddChild(r.Sprite.Entity)

	p := controllers.NewPhysicsLink(&r.VelocityComponent, &r.Positioned)
	r.AddChild(p.Entity)

	return r
}

func update(e *entity.Entity, dt float64) {
	roach, ok := e.Heir.(*Roach)
	if !ok {
		return
	}

	diff := roach.target_pos.Sub(roach.Pos)
	if diff.LengthSqr() > 0.2 {
		dir := diff.Normalize()
		step_length := dt * roach.Speed

		if diff.Length() < step_length {
			roach.Pos = roach.target_pos
		} else {
			roach.Velocity = dir.Scale(roach.Speed)
		}
	} else {
		roach.target_pos = roach.Pos.Add(common.Vec2{
			X: rand.Float64() - 0.5,
			Y: rand.Float64() - 0.5,
		}.Normalize().Scale(100))
	}

}

func render(e *entity.Entity, screen *ebiten.Image) {
	roach, ok := e.Heir.(*Roach)
	if ok {
		roach.Sprite.Render(roach.Sprite.Entity, screen)
	}
}

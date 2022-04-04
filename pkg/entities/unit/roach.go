package unit

import (
	. "github.com/badrpas/ld50/pkg/common"
	. "github.com/badrpas/ld50/pkg/components"
	. "github.com/badrpas/ld50/pkg/controllers"
	. "github.com/badrpas/ld50/pkg/entities/building"
	. "github.com/badrpas/ld50/pkg/entities/sprite"
	. "github.com/badrpas/ld50/pkg/entity"
	. "github.com/badrpas/ld50/pkg/imgrepo"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"math"
)

type Roach struct {
	*Entity

	*Sprite

	follower           *Follower
	TargetBuilding     *Building
	targetPosComponent Position

	VelocityComponent
	SpeedComponent
}

func NewRoach(pos Vec2) *Roach {
	image := ImgRepo["bug.png"]

	r := &Roach{
		Entity: &Entity{},
		Sprite: NewSprite(image, pos),
	}
	r.Speed = 100
	r.Render = render
	r.Update = update
	r.Init = init_roach
	r.Heir = r
	r.AddChild(r.Sprite.Entity)

	b := resolv.NewObject(pos.X, pos.Y, 4, 4)
	b.Data = r.Entity
	b.SetShape(resolv.NewCircle(0, 0, 16))
	AddResolvRegistrator(r.Entity, b)

	p := NewPhysicsLink(&r.VelocityComponent, &r.Position, &r.SpeedComponent, b)
	r.AddChild(p.Entity)

	r.follower = NewPosFollower(p, nil)
	r.AddChild(r.follower.Entity)

	return r
}

func init_roach(e *Entity) {
	r, _ := e.Heir.(*Roach)
	g := r.Game

	grid := g.GetGrid()
	cell := grid.GetCellAtPos(r.Pos)
	if cell.GetHolder() == nil {
		cell.SetHolder(r.Entity)
	} else {
		idx := 0
		shift := [2]int{}
		x, y := cell.GetPos()

		for true {
			cell := grid.GetCellAt(x+shift[0], y+shift[1])
			if cell.GetHolder() == nil {
				cell.SetHolder(r.Entity)
				r.Pos = cell.GetPosWorld()
				return
			}
			if idx%2 == 0 {
				shift[0]++
			} else {
				shift[1]++
			}
			idx++
		}
	}
}

func update(e *Entity, dt float64) {
	roach, ok := e.Heir.(*Roach)
	if !ok {
		return
	}
	g := roach.Game
	if g == nil {
		return
	}

	var closest *Building
	distance := math.MaxFloat64
	g.EachEntity(func(i interface{}) {
		e, ok := i.(*Entity)
		if !ok {
			return
		}

		b, ok := e.Heir.(*Building)
		if ok {
			dist := b.Pos.Sub(roach.Pos).LengthSqr()

			if dist < distance {
				distance = dist
				closest = b
			}
		}
	})

	roach.TargetBuilding = closest

	if closest != nil {
		roach.follower.Target = &roach.targetPosComponent
		roach.targetPosComponent.Pos = closest.GetNearestPoint(roach.Pos)
	}
}

func render(e *Entity, screen *ebiten.Image) {
	roach, ok := e.Heir.(*Roach)
	if ok {
		roach.Sprite.Render(roach.Sprite.Entity, screen)
	}
}

package controllers

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/solarlune/resolv"
	"math"
)

type VelocityComponent struct {
	Velocity common.Vec2
}

type PhysLink struct {
	obj *resolv.Object

	*VelocityComponent

	*common.Positioned

	*entity.Entity
}

var reg []*PhysLink

func NewPhysicsLink(velocityComponent *VelocityComponent, positioned *common.Positioned, obj *resolv.Object) *PhysLink {
	p := &PhysLink{
		obj,
		velocityComponent,
		positioned,
		&entity.Entity{},
	}
	p.Heir = p

	reg = append(reg, p)

	return p
}

func UpdatePhysics(dt float64) {
	l := len(reg)
	idx := 0
	for _, info := range reg {
		if info.Entity.Parent == nil {
			continue
		}
		reg[idx] = info
		idx++

		diff := info.Velocity.Scale(dt)
		if info.obj != nil {
			dx, dy := diff.XY()
			adx, ady := math.Abs(dx), math.Abs(dy)

			if c := info.obj.Check(dx, dy); c != nil {
				for _, object := range c.Objects {
					d := c.ContactWithObject(object)
					cx, cy := d.X(), d.Y()
					if math.Abs(cx) < adx {
						dx = cx
					}
					if math.Abs(cy) < ady {
						dy = cy
					}
				}
			}

			diff.X, diff.Y = dx, dy
		}

		info.Pos = info.Pos.Add(diff)

		if info.obj != nil {
			info.obj.X, info.obj.Y = info.Pos.XY()
			info.obj.Update()
		}
	}

	if l != idx {
		reg = reg[:idx]
	}
}

func AddResolvRegistrator(e *entity.Entity, object *resolv.Object) {
	temp := &entity.Entity{}
	temp.Update = func(self *entity.Entity, dt float64) {
		if self.Game == nil {
			return
		}
		self.Game.GetSpace().Add(object)
		self.SetParent(nil)
	}
	temp.Heir = temp

	e.AddChild(temp)
}

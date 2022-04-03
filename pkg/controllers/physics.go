package controllers

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/entity"
)

type VelocityComponent struct {
	Velocity common.Vec2
}

type PhysLink struct {
	*VelocityComponent

	*common.Positioned

	*entity.Entity
}

var reg []*PhysLink

func NewPhysicsLink(velocityComponent *VelocityComponent, positioned *common.Positioned) *PhysLink {
	p := &PhysLink{
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

		info.Pos = info.Pos.Add(info.Velocity.Scale(dt))
		reg[idx] = info
		idx++
	}

	if l != idx {
		reg = reg[:idx]
	}
}

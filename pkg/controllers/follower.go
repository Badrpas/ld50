package controllers

import (
	"github.com/badrpas/ld50/pkg/components"
	"github.com/badrpas/ld50/pkg/entity"
)

type Follower struct {
	*entity.Entity

	ActorLink *PhysLink

	Target *components.Position
}

func NewPosFollower(pLink *PhysLink, target *components.Position) *Follower {
	follower := &Follower{
		Entity:    &entity.Entity{},
		ActorLink: pLink,
		Target:    target,
	}
	follower.Heir = follower
	follower.Update = update

	return follower
}

func update(self *entity.Entity, dt float64) {
	follower, ok := self.Heir.(*Follower)
	if !ok {
		return
	}
	actor := follower.ActorLink
	target := follower.Target
	if target == nil {
		return
	}

	diff := target.Pos.Sub(actor.Pos)
	if diff.LengthSqr() > 0.2 {
		dir := diff.Normalize()
		step_length := dt * actor.Speed

		if diff.Length() < step_length {
			actor.Pos = target.Pos
		} else {
			actor.Vel = dir.Scale(actor.Speed)
		}
	}
}

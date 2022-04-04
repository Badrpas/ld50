package controllers

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/components"
	"github.com/badrpas/ld50/pkg/entity"
	"github.com/badrpas/ld50/pkg/pathing"
)

type Follower struct {
	*entity.Entity

	ActorLink *PhysLink

	Target *components.Position

	waypoints    []common.Vec2
	waypoint_idx int
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

	grid := follower.Game.GetGrid()
	current_cell := grid.GetCellAtPos(follower.ActorLink.Pos)

	if current_cell.GetHolder() == nil {
		current_cell.SetHolder(follower.Parent)
	}

	if len(follower.waypoints) <= follower.waypoint_idx {
		follower.waypoints = nil
	}

	if 0 == len(follower.waypoints) {
		follower.waypoints = pathing.FindWay(grid, actor.Pos, target.Pos, 50)
		if 0 != len(follower.waypoints) {
			follower.waypoint_idx = 0
		} else {
			follower.waypoint_idx = -1
		}
	}

	if follower.waypoint_idx == -1 {
		return
	}

	if follower.waypoint_idx >= 2 {
		to_clear_pos := follower.waypoints[follower.waypoint_idx-2]
		cell := grid.GetCellAtPos(to_clear_pos)
		if current_cell != cell && cell.GetHolder() == follower.Parent {
			cell.SetHolder(nil)
		}
	}

	target_pos := follower.waypoints[follower.waypoint_idx]

	diff := target_pos.Sub(actor.Pos)
	if diff.LengthSqr() > 0.2 {

		dir := diff.Normalize()
		step_length := dt * float64(actor.Speed)
		speed := actor.Speed

		length := diff.Length()
		if length < step_length {
			speed *= length / step_length
			follower.waypoint_idx++
		}

		target_cell := grid.GetCellAtPos(target_pos)
		if target_cell.GetHolder() == nil {
			target_cell.SetHolder(follower.Parent)
		}

		if target_cell.GetHolder() != follower.Parent {
			for _, waypoint := range follower.waypoints {
				cell := grid.GetCellAtPos(waypoint)
				if cell.GetHolder() == follower.Parent && cell != current_cell {
					cell.SetHolder(nil)
				}
			}

			follower.waypoints = nil

			actor.Vel = common.Vec2{}
		} else {
			actor.Vel = dir.Scale(speed)
		}

	} else {
		actor.Vel = common.Vec2{}
		follower.waypoint_idx++
	}

}

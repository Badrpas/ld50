package game

import (
	"github.com/badrpas/ld50/pkg/entity"
	"sort"
)

type update_order struct {
	roots      []*entity.Entity
	priorities map[*entity.Entity]float64
}

func new_update_order() update_order {
	return update_order{
		roots:      nil,
		priorities: map[*entity.Entity]float64{},
	}
}

func (u *update_order) AddRootEntity(e *entity.Entity) {
	for _, root_entity := range u.roots {
		if root_entity == e {
			return
		}
	}

	u.roots = append(u.roots, e)
}

func (u *update_order) ClearEntityUpdateOrder(e *entity.Entity) {
	for idx, root_entity := range u.roots {
		if root_entity == e {
			u.roots = append(u.roots[:idx], u.roots[idx+1:]...)
			break
		}
	}

	delete(u.priorities, e)
}

func (u *update_order) SetUpdatePriority(e *entity.Entity, value float64) {
	u.priorities[e] = value

	if e.Parent != nil {
		u.SortEntitiesByPrio(e.Parent.Children)
	} else {
		u.SortEntitiesByPrio(u.roots)
	}
}

func (u *update_order) SortEntitiesByPrio(arr []*entity.Entity) {
	sort.Slice(arr, func(i, j int) bool {
		arr := arr
		a, b := arr[i], arr[j]

		ap, _ := u.priorities[a]
		bp, _ := u.priorities[b]

		return ap < bp
	})
}

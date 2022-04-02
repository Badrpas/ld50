package game

import (
	"github.com/badrpas/ld50/pkg/entity"
	"sort"
)

type z_order struct {
	zList     []float64
	zLevels   map[float64]EntitiesStorage
	entitiesZ map[*entity.Entity]float64
}

func new_z_order() z_order {
	return z_order{
		zList:     []float64{},
		zLevels:   map[float64]EntitiesStorage{},
		entitiesZ: map[*entity.Entity]float64{},
	}
}

func (g *Game) SetEntityZ(e *entity.Entity, z float64) {
	storage := g.zLevels[z]
	if storage == nil {
		g.zLevels[z] = EntitiesStorage{}
		storage = g.zLevels[z]
		g.zList = append(g.zList, z)
		sort.Float64s(g.zList)
	}

	pre_z, exists := g.entitiesZ[e]
	if exists {
		delete(g.zLevels[pre_z], e)
	}

	g.entitiesZ[e] = z
	storage[e] = e
}

func (g *Game) ClearEntityZ(e *entity.Entity) {
	pre_z, exists := g.entitiesZ[e]
	if exists {
		delete(g.zLevels[pre_z], e)
		delete(g.entitiesZ, e)
	}
}

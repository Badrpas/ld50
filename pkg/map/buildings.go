package gamemap

import (
	"github.com/badrpas/ld50/pkg/common"
	. "github.com/badrpas/ld50/pkg/entities/building"
	"github.com/badrpas/ld50/pkg/faces/game"
	"github.com/lafriks/go-tiled"
)

func load_buildings(layer *tiled.ObjectGroup, g game.Game, file *tiled.Map) {
	for _, objInfo := range layer.Objects {
		prototile := GetTileByGid(objInfo.GID, file)
		img := GetImageFromTileImage(prototile.Image)

		s := NewBuilding(img, common.Vec2{objInfo.X, objInfo.Y})
		g.AddEntity(s.Entity)
	}
}

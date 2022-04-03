package gamemap

import (
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/entities/sprite"
	"github.com/badrpas/ld50/pkg/game"
	"github.com/lafriks/go-tiled"
)

func load_ground(layer *tiled.Layer, g *game.Game, file *tiled.Map) {
	cell_w := float64(file.TileWidth)
	cell_h := float64(file.TileHeight)

	for idx, tile := range layer.Tiles {
		if tile.IsNil() {
			continue
		}

		idx_x := (idx % (file.Width))
		idx_y := (idx / (file.Width))
		x := cell_w*float64(idx_x) + cell_w/2
		y := cell_h*float64(idx_y) + cell_h/2

		img := GetImageFromLayerTile(tile)
		s := sprite.NewSprite(img, common.Vec2{x, y})
		g.AddEntity(s.Entity)
		g.SetEntityZ(s.Entity, -100)

		g.GetGrid().GetCellAtPos(s.Pos)
	}

}

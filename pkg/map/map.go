package gamemap

import (
	"fmt"
	"github.com/badrpas/ld50/pkg/common"
	"github.com/badrpas/ld50/pkg/entities/sprite"
	"github.com/badrpas/ld50/pkg/game"
	imagerepo "github.com/badrpas/ld50/pkg/img"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lafriks/go-tiled"
	"log"
	"strings"
)

func LoadMap(name string, game *game.Game) error {
	path := fmt.Sprintf("assets/maps/%s", name)
	file, err := tiled.LoadFile(path)
	if err != nil {
		log.Println(err)
		return err
	}

	game.Camera.SetPosition(
		float64(file.Width*file.TileWidth/2),
		float64(file.Height*file.TileHeight/2),
	)

	for _, layer := range file.Layers {
		switch layer.Name {
		case "ground":
			loadGround(layer, game, file)
		}
	}

	return nil
}

func loadGround(layer *tiled.Layer, g *game.Game, file *tiled.Map) {
	cell_w := float64(file.TileWidth)
	cell_h := float64(file.TileHeight)

	for idx, tile := range layer.Tiles {
		if tile.IsNil() {
			continue
		}

		idx_x := (idx % (file.Width))
		idx_y := (idx / (file.Width))
		x := cell_w * float64(idx_x)
		y := cell_h * float64(idx_y)

		img := GetImage(tile)
		s := sprite.NewSpriteEntity(img, common.Vec2{x, y})
		g.AddEntity(s)
	}

}

func GetImage(t *tiled.LayerTile) *ebiten.Image {
	for _, prototile := range t.Tileset.Tiles {
		if prototile.ID == t.ID {
			img_name := prototile.Image.Source
			img_name = strings.Replace(img_name, "../img/", "", 1)
			img, ok := imagerepo.ImgRepo[img_name]
			if !ok {
				log.Println("Can't find image", img_name)
			}
			return img
		}
	}
	return nil
}

package gamemap

import (
	"fmt"
	"github.com/badrpas/ld50/pkg/game"
	imagerepo "github.com/badrpas/ld50/pkg/imgrepo"
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
			load_ground(layer, game, file)
		default:
			log.Printf("Unrecognized tile layer name %s", layer.Name)
		}
		log.Println("Processed tile layer", layer.Name)
	}

	for _, objectGroup := range file.ObjectGroups {
		switch objectGroup.Name {
		case "buildings":
			load_buildings(objectGroup, game, file)
		default:
			log.Printf("Unrecognized object layer name %s", objectGroup.Name)
		}
		log.Println("Processed tile layer", objectGroup.Name)
	}

	return nil
}

func GetTileByGid(gid uint32, file *tiled.Map) *tiled.TilesetTile {
	for _, tileset := range file.Tilesets {
		start_at := (tileset.FirstGID)
		count := uint32(tileset.TileCount)
		if start_at <= gid && gid < start_at+count {
			return tileset.Tiles[gid-start_at]
		}
	}

	return nil
}

func GetImageFromLayerTile(t *tiled.LayerTile) *ebiten.Image {
	for _, prototile := range t.Tileset.Tiles {
		if prototile.ID == t.ID {
			return GetImageFromTileImage(prototile.Image)
		}
	}

	return nil
}

func GetImageFromTileImage(image *tiled.Image) *ebiten.Image {
	img_name := image.Source
	img_name = strings.Replace(img_name, "../img/", "", 1)
	img, ok := imagerepo.ImgRepo[img_name]
	if !ok {
		log.Println("Can't find image", img_name)
	}
	return img
}

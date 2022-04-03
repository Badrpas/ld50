package imgrepo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	_ "image/png"
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
	"strings"
)

var ImgRepo = make(map[string]*ebiten.Image)

func init() {
	_ = filepath.WalkDir("assets/img", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		is_png, _ := regexp.MatchString("\\.png$", path)
		if !is_png {
			return nil
		}

		{
			img, _, err := ebitenutil.NewImageFromFile(path)
			if err != nil {
				return err
			}

			key := strings.ReplaceAll(path, "\\", "/")
			key = strings.Replace(key, "assets/img/", "", 1)
			ImgRepo[key] = img

			log.Println("Loaded image", key)
		}

		return nil
	})
}

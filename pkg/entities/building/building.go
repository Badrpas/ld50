package building

import (
	"github.com/badrpas/ld50/pkg/common"
	. "github.com/badrpas/ld50/pkg/components"
	. "github.com/badrpas/ld50/pkg/controllers"
	. "github.com/badrpas/ld50/pkg/entities/sprite"
	. "github.com/badrpas/ld50/pkg/entity"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/solarlune/resolv"
	"log"
)

type Building struct {
	*Sprite

	obj *resolv.Object

	*Entity

	HpComponent
}

func NewBuilding(img *ebiten.Image, pos common.Vec2) *Building {
	o := resolv.NewObject(pos.X, pos.Y, 4, 4)
	o.SetShape(resolv.NewCircle(0, 0, 16))

	b := &Building{
		Sprite: NewSprite(img, pos),
		obj:    o,
		Entity: &Entity{
			Init: building_init,
		},
	}
	o.Data = b.Entity
	b.Heir = b

	b.AddChild(b.Sprite.Entity)
	AddResolvRegistrator(b.Entity, o)

	return b
}

func building_init(entity *Entity) {
	building, ok := entity.Heir.(*Building)
	if !ok {
		log.Fatalln("Not a Building")
	}
	g := entity.Game
	cell := g.GetGrid().GetCellAtPos(building.Pos)
	cell.SetHolder(building.Entity)
}

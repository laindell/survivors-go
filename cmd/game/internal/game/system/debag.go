package system

import (
	"survivors-go/cmd/game/internal/arch"
	"survivors-go/cmd/game/internal/game/prefabs"

	"image/color"

	"github.com/go-glx/ecs/ecs"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Debug struct {
	ebiten arch.Engine
	limit  int
}

func NewDebug(ebt arch.Engine) *Debug {
	return &Debug{
		ebiten: ebt,
		limit:  10,
	}

}

func (d *Debug) TypeID() ecs.SystemTypeID {
	return "debug"
}

func (d *Debug) OnDraw(_ ecs.RuntimeWorld) {
	screen := d.ebiten.ScreenManager().Screen()

	screen.Fill(color.RGBA{255, 255, 255, 255})
}

func (d *Debug) OnUpdate(w ecs.RuntimeWorld) {
	if d.limit <= 0 {
		return
	}

	d.limit--

	w.AddPrefabEntity(prefabs.PrefabIDPlayer)
	w.AddPrefabEntity(prefabs.PrefabIDCamera)
}

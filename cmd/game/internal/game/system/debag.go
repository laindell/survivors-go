package system

import (
	"survivors-go/cmd/game/internal/game/system/internal"

	"image/color"

	"github.com/go-glx/ecs/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Debug struct {
	ebiten internal.Ebiten
}

func NewDebug(ebt internal.Ebiten) *Debug {
	return &Debug{
		ebiten: ebt,
	}

}

func (d *Debug) TypeID() ecs.SystemTypeID {
	return "debug"
}

func (d *Debug) OnDraw(_ ecs.RuntimeWorld) {
	screen := d.ebiten.Screen()

	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	ebitenutil.DebugPrint(screen, "Hello, Survivors Go!")
}

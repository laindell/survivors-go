package system

import (
	"survivors-go/cmd/game/internal/arch"

	"image/color"

	"github.com/go-glx/ecs/ecs"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Debug struct {
	ebiten arch.Engine
}

func NewDebug(ebt arch.Engine) *Debug {
	return &Debug{
		ebiten: ebt,
	}

}

func (d *Debug) TypeID() ecs.SystemTypeID {
	return "debug"
}

func (d *Debug) OnDraw(_ ecs.RuntimeWorld) {
	screen := d.ebiten.ScreenManager().Screen()

	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	ebitenutil.DebugPrint(screen, "Hello, Survivors Go!")
}

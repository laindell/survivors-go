package system

import (
	"survivors-go/cmd/game/internal/arch"
	"survivors-go/cmd/game/internal/game/component"

	"github.com/go-glx/ecs/ecs"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type BoxDrawer struct {
	ebiten arch.Engine
}

func NewBowDrawer(ebiten arch.Engine) *BoxDrawer {
	return &BoxDrawer{ebiten: ebiten}
}

func (b *BoxDrawer) TypeID() ecs.SystemTypeID {
	return "BoxDrawer"
}

func (b *BoxDrawer) OnDraw(w ecs.RuntimeWorld) {
	screen := b.ebiten.ScreenManager().Screen()
	found := ecs.NewFilter2[component.Transform, component.Color](w).Find()

	for found.Next() {
		_, transform, color := found.Get()

		vector.DrawFilledRect(
			screen,
			float32(transform.Position.X()),
			float32(transform.Position.Y()),
			float32(transform.Size.X()),
			float32(transform.Size.Y()),
			color.Color,
			false,
		)
	}
}

package system

import (
	"survivors-go/cmd/game/internal/arch"
	"survivors-go/cmd/game/internal/game" // Додаю імпорт для WorldToScreen
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

		// отримуємо координати світу
		x := transform.Pos.X()
		y := transform.Pos.Y()
		w := transform.Size.X()
		h := transform.Size.Y()
		unit := component.GameUnit
		screenX, screenY := game.WorldToScreen(x, y, unit)
		screenW := float32(w * unit)
		screenH := float32(h * unit)

		vector.DrawFilledRect(
			screen,
			float32(screenX),
			float32(screenY),
			screenW,
			screenH,
			color.Color,
			false,
		)
	}
}

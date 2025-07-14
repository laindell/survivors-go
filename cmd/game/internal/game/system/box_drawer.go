package system

import (
	"survivors-go/cmd/game/internal/game/component"

	"github.com/go-glx/ecs/ecs"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"survivors-go/cmd/game/internal/arch"
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

	// Отримуємо камеру
	cameraFilter := ecs.NewFilter1[component.Camera](w).Find()
	var cam *component.Camera
	for cameraFilter.Next() {
		_, c := cameraFilter.Get()
		cam = c
		break
	}

	// Отримуємо розміри екрану
	screenW := float64(screen.Bounds().Dx())
	screenH := float64(screen.Bounds().Dy())

	found := ecs.NewFilter2[component.Transform, component.Color](w).Find()

	for found.Next() {
		_, transform, color := found.Get()

		// Отримуємо позицію об'єкта в світових координатах
		worldX := transform.Pos.X()
		worldY := transform.Pos.Y()
		worldW := transform.Size.X()
		worldH := transform.Size.Y()

		var screenX, screenY float64
		var zoom float64 = 1.0

		if cam != nil {
			// Віднімаємо позицію камери від позиції об'єкта
			screenX = (worldX - cam.Pos.X()) * component.GameUnit * cam.Zoom
			screenY = (worldY - cam.Pos.Y()) * component.GameUnit * cam.Zoom

			// Зміщуємо в центр екрану
			screenX += screenW / 2
			screenY += screenH / 2

			zoom = cam.Zoom
		} else {
			// Якщо камери немає, просто конвертуємо в піксели
			screenX = worldX*component.GameUnit + screenW/2
			screenY = worldY*component.GameUnit + screenH/2
		}

		// Обчислюємо розмір з урахуванням зуму
		screenW_obj := float32(worldW * component.GameUnit * zoom)
		screenH_obj := float32(worldH * component.GameUnit * zoom)

		// Малюємо прямокутник
		vector.DrawFilledRect(
			screen,
			float32(screenX),
			float32(screenY),
			screenW_obj,
			screenH_obj,
			color.Color,
			false,
		)
	}
}

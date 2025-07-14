package system

import (
	"survivors-go/cmd/game/internal/arch"
	"survivors-go/cmd/game/internal/game/component"

	"github.com/go-glx/ecs/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteRenderer struct {
	ebiten arch.Engine
}

func NewSpriteRenderer(ebiten arch.Engine) *SpriteRenderer {
	return &SpriteRenderer{ebiten: ebiten}
}

func (s *SpriteRenderer) TypeID() ecs.SystemTypeID {
	return "SpriteRenderer"
}

func (s *SpriteRenderer) OnDraw(w ecs.RuntimeWorld) {
	screen := s.ebiten.ScreenManager().Screen()

	// Отримаємо камеру
	cameraFilter := ecs.NewFilter1[component.Camera](w).Find()
	var cam *component.Camera
	for cameraFilter.Next() {
		_, c := cameraFilter.Get()
		cam = c
		break
	}

	// Отримаємо розміри екрану
	screenW := float64(screen.Bounds().Dx())
	screenH := float64(screen.Bounds().Dy())

	// Рендеримо спрайти
	filter := ecs.NewFilter2[component.Transform, component.Sprite](w).Find()

	for filter.Next() {
		_, transform, sprite := filter.Get()

		if sprite.Image == nil {
			continue
		}

		// Отримаємо позицію об'єкта в світових координатах
		worldX := transform.Pos.X()
		worldY := transform.Pos.Y()

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
			// Якщо камери немає, просто конвертуємо в пікселі
			screenX = worldX*component.GameUnit + screenW/2
			screenY = worldY*component.GameUnit + screenH/2
		}

		// Створюємо субзображення для поточного кадру
		frameRect := sprite.GetFrameRect()
		frameImage := sprite.Image.SubImage(frameRect).(*ebiten.Image)

		// Налаштовуємо параметри малювання
		op := &ebiten.DrawImageOptions{}

		// Масштабування
		scaleX := zoom * transform.Scale.X()
		scaleY := zoom * transform.Scale.Y()
		op.GeoM.Scale(scaleX, scaleY)

		// Центрування спрайту
		op.GeoM.Translate(
			screenX-float64(sprite.FrameWidth)*scaleX/2,
			screenY-float64(sprite.FrameHeight)*scaleY/2,
		)

		// Поворот (якщо потрібно)
		if transform.Rotation != 0 {
			op.GeoM.Translate(-float64(sprite.FrameWidth)*scaleX/2, -float64(sprite.FrameHeight)*scaleY/2)
			op.GeoM.Rotate(transform.Rotation)
			op.GeoM.Translate(float64(sprite.FrameWidth)*scaleX/2, float64(sprite.FrameHeight)*scaleY/2)
		}

		// Малюємо спрайт
		screen.DrawImage(frameImage, op)
	}
}

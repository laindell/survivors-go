package system

import (
	"math"
	"survivors-go/cmd/game/internal/arch"
	"survivors-go/cmd/game/internal/game/component"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/go-glx/ecs/ecs"
)

type CameraSystem struct {
	ebiten arch.Engine
}

func NewCameraSystem(ebiten arch.Engine) *CameraSystem {
	return &CameraSystem{ebiten: ebiten}
}

func (s *CameraSystem) TypeID() ecs.SystemTypeID {
	return "CameraSystem"
}

func (s *CameraSystem) OnUpdate(w ecs.RuntimeWorld) {
	// Пошук гравця
	playerFilter := ecs.NewFilter2[component.Player, component.Transform](w).Find()
	var playerPos mgl64.Vec2
	found := false

	for playerFilter.Next() {
		_, _, transform := playerFilter.Get()
		playerPos = transform.Pos
		found = true
		break
	}
	if !found {
		return
	}

	// Оновлення камери
	cameraFilter := ecs.NewFilter1[component.Camera](w).Find()

	for cameraFilter.Next() {
		_, cam := cameraFilter.Get()

		// Встановлюємо цільову позицію камери на позицію гравця
		cam.Target = playerPos.Add(cam.Offset)

		// Плавне слідування за гравцем
		if cam.Speed >= 1.0 {
			// Миттєве слідування
			cam.Pos = cam.Target
		} else {
			// Плавне слідування з інтерполяцією
			diff := cam.Target.Sub(cam.Pos)
			cam.Pos = cam.Pos.Add(diff.Mul(cam.Speed))
		}

		// Застосовуємо обмеження, якщо вони є
		if cam.Bounds != nil {
			bufW := float64(s.ebiten.ScreenManager().BufW())
			bufH := float64(s.ebiten.ScreenManager().BufH())

			// Обчислюємо половину видимої області з урахуванням зуму
			halfScreen := mgl64.Vec2{bufW / 2 / cam.Zoom / component.GameUnit, bufH / 2 / cam.Zoom / component.GameUnit}

			min := cam.Bounds.Min.Add(halfScreen)
			max := cam.Bounds.Max.Sub(halfScreen)

			cam.Pos[0] = math.Max(min[0], math.Min(max[0], cam.Pos[0]))
			cam.Pos[1] = math.Max(min[1], math.Min(max[1], cam.Pos[1]))
		}
	}
}

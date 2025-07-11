package prefabs

import (
	"survivors-go/cmd/game/internal/game/component"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/go-glx/ecs/ecs"
)

func Camera() func() *ecs.Entity {
	return func() *ecs.Entity {
		cam := ecs.NewEntity(PrefabIDCamera)
		camera := component.NewCamera(
			mgl64.Vec2{0, 0}, // позиція todo: треба змінити щоб в напрямку куди йде гравець було зміщення на 1 системну одиницю
			0.1,              // швидкість слідування
			3.0,              // zoom
			mgl64.Vec2{0, 0}, // offset
			nil,              // bounds
		)
		cam.AddComponent(camera)
		return cam
	}
}

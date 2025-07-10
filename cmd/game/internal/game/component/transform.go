package component

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/go-glx/ecs/ecs"
)

const GameUnit = 32.0 // 1 ігрова одиниця = 32 пікселі

type Transform struct {
	Pos      mgl64.Vec2
	Size     mgl64.Vec2
	Rotation float64
	Scale    mgl64.Vec2
}

func NewTransform(pos, size mgl64.Vec2) *Transform {
	return &Transform{
		Pos:      pos,
		Size:     size,
		Rotation: 0,
		Scale:    mgl64.Vec2{1, 1},
	}
}

func (t Transform) TypeID() ecs.ComponentTypeID {
	return "Transform"
}

// Конвертація у пікселі для рендеру
func (t *Transform) ScreenPos() (float64, float64) {
	return t.Pos.X() * GameUnit, t.Pos.Y() * GameUnit
}

func (t *Transform) ScreenSize() (float64, float64) {
	return t.Size.X() * GameUnit, t.Size.Y() * GameUnit
}

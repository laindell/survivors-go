package component

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/go-glx/ecs/ecs"
)

type Transform struct {
	Position mgl64.Vec2
	Size     mgl64.Vec2
	Rotation float64
	Scale    mgl64.Vec2
}

func NewTransform(Position mgl64.Vec2) *Transform {
	return &Transform{
		Position: Position,
		Rotation: 0,
		Scale:    mgl64.Vec2{1, 1},
	}
}

func (t Transform) TypeID() ecs.ComponentTypeID {
	return "Transform"
}

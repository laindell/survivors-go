package container

import (
	"image/color"
	"survivors-go/cmd/game/internal/game/component"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/go-glx/ecs/ecs"
)

func (c *container) ecsComponents() []ecs.Component {
	return *static(c, func() *[]ecs.Component {
		return &[]ecs.Component{
			c.ecsComponentTransform(),
			c.ecsComponentColor(),
			c.ecsComponentPlayer(),
			c.ecsComponentCamera(),
		}
	})
}

func (c *container) ecsComponentTransform() *component.Transform {
	return static(c, func() *component.Transform {
		return component.NewTransform(mgl64.Vec2{}, mgl64.Vec2{})
	})
}

func (c *container) ecsComponentColor() *component.Color {
	return static(c, func() *component.Color {
		return component.NewColor(color.RGBA{255, 255, 255, 255})
	})
}

func (c *container) ecsComponentPlayer() *component.Player {
	return static(c, func() *component.Player {
		return component.NewPlayer(0)
	})
}
func (c *container) ecsComponentCamera() *component.Camera {
	return static(c, func() *component.Camera {
		return component.NewCamera(
			mgl64.Vec2{0, 0}, // початкова позиція
			0.1,              // швидкість слідування (0..1)
			1.0,              // zoom (масштаб)
			mgl64.Vec2{0, 0}, // offset (зсув)
			nil,              // bounds (межі, якщо не потрібні — nil)
		)
	})
}

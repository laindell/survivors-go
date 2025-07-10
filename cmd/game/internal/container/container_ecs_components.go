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

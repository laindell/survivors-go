package component

import (
	"github.com/go-glx/ecs/ecs"
	"image/color"
)

type Color struct {
	Color color.RGBA
}

func NewColor(clr color.RGBA) *Color {
	return &Color{
		Color: clr,
	}
}

func (c Color) TypeID() ecs.ComponentTypeID {
	return "Color"
}

package component

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/go-glx/ecs/ecs"
)

type Rect struct {
	Min, Max mgl64.Vec2
}

type Camera struct {
	Pos    mgl64.Vec2 // поточна позиція камери
	Target mgl64.Vec2 // цільова позиція
	Speed  float64    // швидкість слідування (0..1)
	Offset mgl64.Vec2 // зсув відносно цілі
	Zoom   float64    // масштаб
	Bounds *Rect      // межі (опціонально)
}

func NewCamera(pos mgl64.Vec2, speed float64, zoom float64, offset mgl64.Vec2, bounds *Rect) *Camera {
	return &Camera{
		Pos:    pos,
		Target: pos,
		Speed:  speed,
		Offset: offset,
		Zoom:   zoom,
		Bounds: bounds,
	}
}

func (c Camera) TypeID() ecs.ComponentTypeID {
	return "Camera"
}

func (c *Camera) Matrix(screenW, screenH float64) mgl64.Mat3 {
	scale := mgl64.Scale2D(c.Zoom, c.Zoom)
	translate := mgl64.Translate2D(-c.Pos[0], -c.Pos[1])
	center := mgl64.Translate2D(screenW/2, screenH/2)
	return center.Mul3(scale).Mul3(translate)
}

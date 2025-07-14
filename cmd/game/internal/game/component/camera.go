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

// Метод для отримання view матриці (більше не використовується в новій реалізації)
func (c *Camera) Matrix(screenW, screenH float64) mgl64.Mat3 {
	translate := mgl64.Translate2D(-c.Pos[0]*GameUnit, -c.Pos[1]*GameUnit)
	scale := mgl64.Scale2D(c.Zoom, c.Zoom)
	center := mgl64.Translate2D(screenW/2, screenH/2)

	return center.Mul3(scale).Mul3(translate)
}

// Допоміжні методи для конвертації координат
func (c *Camera) WorldToScreen(worldX, worldY, screenW, screenH float64) (float64, float64) {
	screenX := (worldX-c.Pos.X())*GameUnit*c.Zoom + screenW/2
	screenY := (worldY-c.Pos.Y())*GameUnit*c.Zoom + screenH/2
	return screenX, screenY
}

func (c *Camera) ScreenToWorld(screenX, screenY, screenW, screenH float64) (float64, float64) {
	worldX := (screenX-screenW/2)/(GameUnit*c.Zoom) + c.Pos.X()
	worldY := (screenY-screenH/2)/(GameUnit*c.Zoom) + c.Pos.Y()
	return worldX, worldY
}

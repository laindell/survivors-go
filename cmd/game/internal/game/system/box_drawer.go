package system

import (
	"survivors-go/cmd/game/internal/arch" // Додаю імпорт для WorldToScreen
	"survivors-go/cmd/game/internal/game/component"

	"github.com/go-glx/ecs/ecs"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type BoxDrawer struct {
	ebiten arch.Engine
}

func NewBowDrawer(ebiten arch.Engine) *BoxDrawer {
	return &BoxDrawer{ebiten: ebiten}
}

func (b *BoxDrawer) TypeID() ecs.SystemTypeID {
	return "BoxDrawer"
}

func (b *BoxDrawer) OnDraw(w ecs.RuntimeWorld) {
	screen := b.ebiten.ScreenManager().Screen()

	// Отримуємо камеру
	cameraFilter := ecs.NewFilter1[component.Camera](w).Find()
	var cam *component.Camera
	for cameraFilter.Next() {
		_, c := cameraFilter.Get()
		cam = c
		break
	}
	var camMatrix mgl64.Mat3
	if cam != nil {
		w, h := float64(screen.Bounds().Dx()), float64(screen.Bounds().Dy())
		camMatrix = cam.Matrix(w, h)
	} else {
		camMatrix = mgl64.Ident3()
	}

	found := ecs.NewFilter2[component.Transform, component.Color](w).Find()

	for found.Next() {
		_, transform, color := found.Get()

		// отримуємо координати світу
		x := transform.Pos.X()
		y := transform.Pos.Y()
		w := transform.Size.X()
		h := transform.Size.Y()
		unit := component.GameUnit

		// Трансформуємо позицію через матрицю камери
		pos := mgl64.Vec3{x * unit, y * unit, 1}
		pos = camMatrix.Mul3x1(pos)

		screenW := float32(w * unit * (cam.Zoom))
		screenH := float32(h * unit * (cam.Zoom))

		vector.DrawFilledRect(
			screen,
			float32(pos[0]),
			float32(pos[1]),
			screenW,
			screenH,
			color.Color,
			false,
		)
	}
}

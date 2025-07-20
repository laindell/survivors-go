package system

import (
	"survivors-go/cmd/game/internal/game/component"

	"survivors-go/cmd/game/internal/arch"

	"github.com/go-glx/ecs/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type WorldTilesRenderer struct {
	ebiten arch.Engine
}

func NewWorldTilesRenderer(ebiten arch.Engine) *WorldTilesRenderer {
	return &WorldTilesRenderer{ebiten: ebiten}
}

func (w *WorldTilesRenderer) OnDraw(world ecs.RuntimeWorld) {
	screen := w.ebiten.ScreenManager().Screen()
	// Знаходимо камеру
	cameraFilter := ecs.NewFilter1[component.Camera](world).Find()
	var cam *component.Camera
	for cameraFilter.Next() {
		_, c := cameraFilter.Get()
		cam = c
		break
	}
	if cam == nil {
		return
	}
	filter := ecs.NewFilter2[component.Transform, component.Sprite](world).Find()
	for filter.Next() {
		_, tr, spr := filter.Get()
		if spr.Image == nil {
			continue
		}
		if !spr.IsTile {
			continue // малюємо лише тайли
		}
		// Обчислюємо позицію на екрані з урахуванням камери
		screenX := (tr.Pos.X() - cam.Pos.X()) * component.GameUnit * cam.Zoom
		screenY := (tr.Pos.Y() - cam.Pos.Y()) * component.GameUnit * cam.Zoom
		screenX += float64(screen.Bounds().Dx()) / 2
		screenY += float64(screen.Bounds().Dy()) / 2
		// Малюємо тайл
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(cam.Zoom, cam.Zoom)
		op.GeoM.Translate(screenX, screenY)
		screen.DrawImage(spr.Image, op)
	}
}

func (w *WorldTilesRenderer) TypeID() ecs.SystemTypeID {
	return "WorldTilesRenderer"
}

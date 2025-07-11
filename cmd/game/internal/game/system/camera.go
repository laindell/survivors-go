package system

import (
	"math"
	"survivors-go/cmd/game/internal/game/component"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/go-glx/ecs/ecs"
)

type CameraSystem struct{}

func NewCameraSystem() *CameraSystem {
	return &CameraSystem{}
}

func (s *CameraSystem) TypeID() ecs.SystemTypeID {
	return "CameraSystem"
}

func (s *CameraSystem) OnUpdate(w ecs.RuntimeWorld) {
	// Пошук гравця
	playerFilter := ecs.NewFilter2[component.Player, component.Transform](w).Find()
	var playerPos mgl64.Vec2
	found := false
	for playerFilter.Next() {
		_, _, transform := playerFilter.Get()
		playerPos = transform.Pos
		found = true
		break
	}
	if !found {
		return
	}
	// Оновлення камери
	cameraFilter := ecs.NewFilter1[component.Camera](w).Find()
	for cameraFilter.Next() {
		_, cam := cameraFilter.Get()
		cam.Target = playerPos.Add(cam.Offset)
		delta := cam.Target.Sub(cam.Pos)
		cam.Pos = cam.Pos.Add(delta.Mul(cam.Speed))
		if cam.Bounds != nil {
			halfScreen := mgl64.Vec2{400 / cam.Zoom, 300 / cam.Zoom} // TODO: замінити на актуальні розміри екрану
			min := cam.Bounds.Min.Add(halfScreen)
			max := cam.Bounds.Max.Sub(halfScreen)
			cam.Pos[0] = math.Max(min[0], math.Min(max[0], cam.Pos[0]))
			cam.Pos[1] = math.Max(min[1], math.Min(max[1], cam.Pos[1]))
		}
	}
}

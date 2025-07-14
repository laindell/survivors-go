package system

import (
	"survivors-go/cmd/game/internal/game/component"

	"github.com/go-glx/ecs/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayerMoveSystem struct{}

func NewPlayerMoveSystem() *PlayerMoveSystem {
	return &PlayerMoveSystem{}
}

func (s *PlayerMoveSystem) TypeID() ecs.SystemTypeID {
	return "PlayerMoveSystem"
}

func (s *PlayerMoveSystem) OnUpdate(w ecs.RuntimeWorld) {
	filter := ecs.NewFilter2[component.Player, component.Transform](w).Find()
	for filter.Next() {
		_, player, transform := filter.Get()
		dx, dy := 0.0, 0.0
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			dy -= player.Speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			dy += player.Speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			dx -= player.Speed
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			dx += player.Speed
		}
		transform.Pos[0] += dx
		transform.Pos[1] += dy
	}
}

package system

import (
	"survivors-go/cmd/game/internal/game/component"
	"survivors-go/cmd/game/internal/game/system/internal"

	"github.com/go-glx/ecs/ecs"
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
		dx, dy := internal.GetMoveDirection()
		dx *= player.Speed
		dy *= player.Speed

		transform.Pos[0] += dx
		transform.Pos[1] += dy
	}
}

package prefabs

import (
	"image/color"
	"survivors-go/cmd/game/internal/game/component"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/go-glx/ecs/ecs"
)

func Player() func() *ecs.Entity {
	return func() *ecs.Entity {
		plr := ecs.NewEntity(PrefabIDPlayer)

		transform := component.NewTransform(
			mgl64.Vec2{
				1,
				1,
			},
			mgl64.Vec2{
				1,
				1,
			},
		)

		color := component.NewColor(color.RGBA{255, 255, 0, 250})

		player := component.NewPlayer(1.0)

		plr.AddComponent(transform)
		plr.AddComponent(color)
		plr.AddComponent(player)

		return plr
	}
}

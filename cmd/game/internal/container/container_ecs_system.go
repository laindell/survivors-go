package container

import (
	"survivors-go/cmd/game/internal/game/system"

	"github.com/go-glx/ecs/ecs"
)

func (c *container) ecsSystems() []ecs.System {
	return *static(c, func() *[]ecs.System {
		return &[]ecs.System{
			c.ecsSystemPlayerMove(),
			c.ecsSystemPlayerAnimation(),
			c.ecsSystemAnimation(),
			c.ecsSystemCamera(),
			c.ecsSystemBoxDrawer(),
			c.ecsSystemSpriteRenderer(),
		}
	})
}

func (c *container) ecsSystemBoxDrawer() *system.BoxDrawer {
	return static(c, func() *system.BoxDrawer {
		return system.NewBowDrawer(
			c.ebitenEngine(),
		)
	})
}

func (c *container) ecsSystemCamera() *system.CameraSystem {
	return static(c, func() *system.CameraSystem {
		return system.NewCameraSystem(
			c.ebitenEngine(),
		)
	})
}

func (c *container) ecsSystemPlayerMove() *system.PlayerMoveSystem {
	return static(c, func() *system.PlayerMoveSystem {
		return system.NewPlayerMoveSystem()
	})
}

func (c *container) ecsSystemSpriteRenderer() *system.SpriteRenderer {
	return static(c, func() *system.SpriteRenderer {
		return system.NewSpriteRenderer(
			c.ebitenEngine(),
		)
	})
}

func (c *container) ecsSystemAnimation() *system.AnimationSystem {
	return static(c, func() *system.AnimationSystem {
		return system.NewAnimationSystem()
	})
}

func (c *container) ecsSystemPlayerAnimation() *system.PlayerAnimationSystem {
	return static(c, func() *system.PlayerAnimationSystem {
		return system.NewPlayerAnimationSystem()
	})
}

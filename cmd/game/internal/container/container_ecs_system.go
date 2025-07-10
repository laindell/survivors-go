package container

import (
	"survivors-go/cmd/game/internal/game/system"

	"github.com/go-glx/ecs/ecs"
)

func (c *container) ecsSystems() []ecs.System {
	return *static(c, func() *[]ecs.System {
		return &[]ecs.System{
			c.ecsSystemDebug(),
			c.ecsSystemBoxDrawer(),
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

func (c *container) ecsSystemDebug() *system.Debug {
	return static(c, func() *system.Debug {
		return system.NewDebug(
			c.ebitenEngine(),
		)
	})
}

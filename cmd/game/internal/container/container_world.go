package container

import (
	"survivors-go/cmd/game/internal/game/system"

	"github.com/go-glx/ecs/ecs"
)

func (c *container) ecsWorld() *ecs.World {
	return static(c, func() *ecs.World {
		world := ecs.NewWorld(
			c.ecsRegistry(),
		)

		for _, sys := range c.ecsSystems() {
			world.AddSystem(sys.TypeID())
		}

		return world
	})
}

func (c *container) ecsRegistry() *ecs.Registry {
	return static(c, func() *ecs.Registry {
		reg := ecs.NewRegistry()

		for _, sys := range c.ecsSystems() {
			reg.RegisterSystem(sys)
		}

		return reg
	})
}

func (c *container) ecsSystems() []ecs.System {
	return *static(c, func() *[]ecs.System {
		return &[]ecs.System{
			c.ecsSystemDebug(),
		}
	})
}

func (c *container) ecsSystemDebug() *system.Debug {
	return static(c, func() *system.Debug {
		return system.NewDebug(
			c.ebitenScreenManager(),
		)
	})
}

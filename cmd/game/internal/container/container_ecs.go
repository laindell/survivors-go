package container

import (
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

		for _, cmp := range c.ecsComponents() {
			reg.RegisterComponent(cmp)
		}

		return reg
	})
}

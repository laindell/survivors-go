package container

import (
	"github.com/andygeiss/ecs"
)

func (c *container) Engine() ecs.Engine {
	return ecs.NewDefaultEngine(ecs.NewEntityManager(), ecs.NewSystemManager())
}

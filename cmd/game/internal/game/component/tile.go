package component

import "github.com/go-glx/ecs/ecs"

type Tile struct{}

func (t Tile) TypeID() ecs.ComponentTypeID {
	return "Tile"
}

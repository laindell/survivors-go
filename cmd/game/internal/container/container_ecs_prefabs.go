package container

import (
	"survivors-go/cmd/game/internal/game/prefabs"

	"github.com/go-glx/ecs/ecs"
)

func (c *container) ecsPrefabs() []ecs.Prefab {
	return *static(c, func() *[]ecs.Prefab {
		return &[]ecs.Prefab{
			*c.ecsPrefabPlayer(),
			*c.ecsPrefabCamera(),
		}

	})
}

func (c *container) ecsPrefabPlayer() *ecs.Prefab {
	return static(c, func() *ecs.Prefab {
		return ecs.NewPrefab(prefabs.PrefabIDPlayer, prefabs.Player())

	})
}

func (c *container) ecsPrefabCamera() *ecs.Prefab {
	return static(c, func() *ecs.Prefab {
		return ecs.NewPrefab(prefabs.PrefabIDCamera, prefabs.Camera())
	})
}

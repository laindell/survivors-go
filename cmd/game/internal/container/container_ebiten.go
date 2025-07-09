package container

import (
	"survivors-go/cmd/game/internal/arch"
	"survivors-go/cmd/game/internal/arch/screen"
)

func (c *container) ebitenEngine() *arch.EngineImpl {
	return static(c, func() *arch.EngineImpl {
		return arch.NewEngineImpl(
			c.ebitenScreenManager(),
		)
	})
}

func (c *container) ebitenScreenManager() *screen.Manager {
	return static(c, func() *screen.Manager {
		return screen.NewManager()
	})
}

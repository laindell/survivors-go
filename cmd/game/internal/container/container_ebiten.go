package container

import "survivors-go/cmd/game/internal/arch/screen"

func (c *container) ebitenScreenManager() *screen.Manager {
	return static(c, func() *screen.Manager {
		return screen.NewManager()
	})
}

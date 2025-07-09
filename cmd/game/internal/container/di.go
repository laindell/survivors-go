package container

import (
	"survivors-go/cmd/game/internal/game"
)

type DI struct {
	cont *container
}
type closer interface {
	EnqueueFree(fn func())
	EnqueueClose(fn func() error)
}

func NewDI(closer closer) *DI {
	return &DI{
		cont: newContainer(closer),
	}
}
func (c *DI) ProvideGame() *game.Game {
	return c.cont.game()
}

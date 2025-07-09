package container

import (
	"survivors-go/cmd/game/internal/game"
)

func (cont *container) game() *game.Game {
	return static(cont, func() *game.Game {
		return game.NewGame(cont.Engine())
	})
}

package internal

import (
	"log"
	"survivors-go/cmd/game/internal/container"

	"github.com/hajimehoshi/ebiten/v2"
	"survivors-go/cmd/game/internal/game"
)

func RunGame() int {
	game.InitWorldParams(-20, 20, -20, 20)
	Closer := NewCloser()
	di := container.NewDI(Closer)

	game := di.ProvideGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}

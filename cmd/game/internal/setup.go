package internal

import (
	"log"
	"survivors-go/cmd/game/internal/container"

	"github.com/hajimehoshi/ebiten/v2"
)

func RunGame() int {
	Closer := NewCloser()
	di := container.NewDI(Closer)

	game := di.ProvideGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}

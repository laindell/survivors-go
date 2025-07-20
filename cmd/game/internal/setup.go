package internal

import (
	"log"
	"survivors-go/cmd/game/internal/container"

	"github.com/hajimehoshi/ebiten/v2"
	"survivors-go/cmd/game/internal/arch"
)

func RunGame() int {
	arch.InitWorldParams(-10, 10, -10, 10)
	Closer := NewCloser()
	di := container.NewDI(Closer)

	game := di.ProvideGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
		return 1
	}
	return 0
}

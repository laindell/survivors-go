package main

import (
	"os"
	"survivors-go/cmd/game/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Survivor")

}

func main() {
	os.Exit(internal.RunGame())
}

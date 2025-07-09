package internal

import "github.com/hajimehoshi/ebiten/v2"

type Ebiten interface {
	Screen() *ebiten.Image
}

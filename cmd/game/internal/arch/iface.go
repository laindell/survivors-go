package arch

import "github.com/hajimehoshi/ebiten/v2"

type (
	Engine interface {
		ScreenManager() ScreenManager
	}
	ScreenManager interface {
		Screen() *ebiten.Image
	}
)

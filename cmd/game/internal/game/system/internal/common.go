package internal

import "github.com/hajimehoshi/ebiten/v2"

// Повертає напрямок руху гравця за натиснутими клавішами
func GetMoveDirection() (dx, dy float64) {
	switch {
	case ebiten.IsKeyPressed(ebiten.KeyW) && ebiten.IsKeyPressed(ebiten.KeyA):
		dy -= 0.5
		dx -= 0.5
	case ebiten.IsKeyPressed(ebiten.KeyW) && ebiten.IsKeyPressed(ebiten.KeyD):
		dy -= 0.5
		dx += 0.5
	case ebiten.IsKeyPressed(ebiten.KeyS) && ebiten.IsKeyPressed(ebiten.KeyA):
		dy += 0.5
		dx -= 0.5
	case ebiten.IsKeyPressed(ebiten.KeyS) && ebiten.IsKeyPressed(ebiten.KeyD):
		dy += 0.5
		dx += 0.5
	case ebiten.IsKeyPressed(ebiten.KeyW):
		dy -= 1
	case ebiten.IsKeyPressed(ebiten.KeyS):
		dy += 1
	case ebiten.IsKeyPressed(ebiten.KeyA):
		dx -= 1
	case ebiten.IsKeyPressed(ebiten.KeyD):
		dx += 1
	}
	return
}

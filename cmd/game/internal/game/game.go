package game

import (
	"errors"
	"github.com/andygeiss/ecs"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	engine         ecs.Engine
	prevF11Pressed bool
}

func NewGame(engine ecs.Engine) *Game {
	return &Game{
		engine: engine,
	}
}

func (game *Game) Update() error {
	game.engine.Tick()
	f11pressed := ebiten.IsKeyPressed(ebiten.KeyF11)
	if f11pressed && !game.prevF11Pressed {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
		ebiten.SetWindowDecorated(!ebiten.IsFullscreen())
	}

	if ebiten.IsKeyPressed(ebiten.KeyF1) {
		return errors.New("exit")
	}

	game.prevF11Pressed = f11pressed
	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, Survivors Go!")
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
}

func (game *Game) Layout(outsideWidth, outsideHeigxht int) (int, int) {
	return 500, 500
}

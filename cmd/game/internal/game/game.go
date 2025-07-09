package game

import (
	"errors"
	"github.com/go-glx/ecs/ecs"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	world          *ecs.World
	prevF11Pressed bool
}

func NewGame(world *ecs.World) *Game {
	return &Game{
		world: world,
	}
}

func (game *Game) Update() error {
	game.world.Update()
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
	// TODO: screen
	game.world.Draw()
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff})
	ebitenutil.DebugPrint(screen, "Hello, Survivors Go!")
}

func (game *Game) Layout(outsideWidth, outsideHeigxht int) (int, int) {
	return 500, 500
}

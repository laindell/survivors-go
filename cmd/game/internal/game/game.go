package game

import (
	"errors"

	"github.com/go-glx/ecs/ecs"

	//"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"survivors-go/cmd/game/internal/arch/screen"
	"survivors-go/cmd/game/internal/game/component"
)

type Game struct {
	world               *ecs.World
	prevF11Pressed      bool
	ebitenScreenManager *screen.Manager
}

func NewGame(world *ecs.World, ebitenScreenManager *screen.Manager) *Game {
	return &Game{
		world:               world,
		ebitenScreenManager: ebitenScreenManager,
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
	game.ebitenScreenManager.UpdateScreen(screen)
	game.world.Draw()

}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(WorldWidth() * component.GameUnit), int(WorldHeight() * component.GameUnit)
}

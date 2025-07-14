package game

import (
	"errors"

	"github.com/go-glx/ecs/ecs"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"survivors-go/cmd/game/internal/arch/screen"
	// Додаю імпорти для префабів
	"survivors-go/cmd/game/internal/game/assets"
	"survivors-go/cmd/game/internal/game/prefabs"
)

type Game struct {
	world               *ecs.World
	prevF11Pressed      bool
	ebitenScreenManager *screen.Manager
	assetManager        *assets.AssetManager
}

func NewGame(world *ecs.World, ebitenScreenManager *screen.Manager, assetManager *assets.AssetManager) *Game {
	// Створюємо анімованого гравця
	player := prefabs.AnimatedPlayer(assetManager)()
	camera := prefabs.Camera()()
	world.AddEntity(player)
	world.AddEntity(camera)

	return &Game{
		world:               world,
		ebitenScreenManager: ebitenScreenManager,
		assetManager:        assetManager,
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

	buffer := game.ebitenScreenManager.Buffer()
	buffer.Clear()
	oldScreen := game.ebitenScreenManager.Screen()
	game.ebitenScreenManager.UpdateScreen(buffer)
	game.world.Draw()
	game.ebitenScreenManager.UpdateScreen(oldScreen)

	screen.DrawImage(buffer, nil)
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Тепер ігрове поле завжди дорівнює розміру вікна
	game.ebitenScreenManager.SetBufferSize(outsideWidth, outsideHeight)
	return outsideWidth, outsideHeight
}

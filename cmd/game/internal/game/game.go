package game

import (
	"errors"

	"github.com/go-glx/ecs/ecs"

	"github.com/hajimehoshi/ebiten/v2"
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"image/color"
	"survivors-go/cmd/game/internal/arch/screen"

	// Додаю імпорти для префабів
	"survivors-go/cmd/game/internal/game/assets"
	"survivors-go/cmd/game/internal/game/component"
	"survivors-go/cmd/game/internal/game/prefabs"
	"survivors-go/cmd/game/internal/game/system"
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

	// === Додаємо тайли світу ===
	// Завантажуємо тайлсет (шлях з main_map.json)
	tilesetImg, err := assetManager.LoadImage("maps/PNG/Water_coasts_animation.png")
	if err != nil {
		println("[ERROR] Не вдалося завантажити tileset:", err.Error())
	} else {
		mapData, err := assetManager.LoadMap("maps/map_1.ldtk")
		if err != nil {
			println("[ERROR] Не вдалося завантажити карту:", err.Error())
		} else {
			transforms, sprites, err := system.ParseWorldTilesFromBytes(mapData, tilesetImg, 16, 16)
			if err != nil {
				println("[ERROR] Не вдалося розпарсити карту:", err.Error())
			} else {
				println("[INFO] Тайлів створено:", len(transforms))
				for i := 0; i < len(transforms) && i < 3; i++ {
					println("[INFO] Tile", i, "pos:", transforms[i].Pos.X(), transforms[i].Pos.Y(), "sprite nil:", sprites[i] == nil)
				}
				for i := range transforms {
					tileEntity := ecs.NewEntity("")
					tileEntity.AddComponent(transforms[i])
					tileEntity.AddComponent(sprites[i])
					tileEntity.AddComponent(&component.Tile{}) // Додаю Tile
					world.AddEntity(tileEntity)
				}
			}
		}
	}
	// === Кінець додавання тайлів ===

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

	buffer.Fill(color.RGBA{255, 105, 180, 255})
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

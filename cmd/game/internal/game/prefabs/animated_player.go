package prefabs

import (
	"survivors-go/cmd/game/internal/game/assets"
	"survivors-go/cmd/game/internal/game/component"
	"time"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/go-glx/ecs/ecs"
)

func AnimatedPlayer(assetManager *assets.AssetManager) func() *ecs.Entity {
	return func() *ecs.Entity {
		plr := ecs.NewEntity(PrefabIDPlayer)

		// Transform
		transform := component.NewTransform(
			mgl64.Vec2{0, 0},
			mgl64.Vec2{0, 0},
		)

		// Player
		player := component.NewPlayer(0.1)

		// Завантажуємо спрайт
		playerImage, err := assetManager.LoadImage("images/Assets/Player/Player_Walking.png")
		if err != nil {
			playerImage = nil
		}

		// Завантажуємо idle-спрайт
		idleImage, err := assetManager.LoadImage("images/Assets/Player/Player_Idle.png")
		if err != nil {
			idleImage = nil
		}

		// Sprite (walking)
		sprite := component.NewSprite(playerImage, 64, 128)
		sprite.IdleImage = idleImage
		sprite.WalkingImage = playerImage

		animation := component.NewAnimation(8, 100*time.Millisecond, true)

		plr.AddComponent(transform)
		plr.AddComponent(player)
		plr.AddComponent(sprite)
		plr.AddComponent(animation)

		return plr
	}
}

package container

import (
	"survivors-go/cmd/game/internal/game"
	"survivors-go/cmd/game/internal/game/assets"
)

func (c *container) game() *game.Game {
	return static(c, func() *game.Game {
		return game.NewGame(
			c.ecsWorld(),
			c.ebitenScreenManager(),
			c.assetManager(),
		)
	})
}

func (c *container) assetManager() *assets.AssetManager {
	return static(c, func() *assets.AssetManager {
		return assets.NewAssetManager()
	})
}

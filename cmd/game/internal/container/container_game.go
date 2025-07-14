package container

import (
	"survivors-go/cmd/game/internal/game"
	"survivors-go/cmd/game/internal/game/assets"
)

func (cont *container) game() *game.Game {
	return static(cont, func() *game.Game {
		return game.NewGame(
			cont.ecsWorld(),
			cont.ebitenScreenManager(),
			cont.assetManager(),
		)
	})
}

func (cont *container) assetManager() *assets.AssetManager {
	return static(cont, func() *assets.AssetManager {
		return assets.NewAssetManager()
	})
}

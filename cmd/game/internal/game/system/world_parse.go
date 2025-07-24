package system

import (
	"encoding/json"
	"image"
	"os"
	"survivors-go/cmd/game/internal/game/component"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/hajimehoshi/ebiten/v2"
)

type TileData struct {
	Px  [2]int `json:"px"`
	Src [2]int `json:"src"`
	F   int    `json:"f"`
	T   int    `json:"t"`
	D   []int  `json:"d"`
	A   int    `json:"a"`
}

type MapLayer struct {
	GridTiles []TileData `json:"gridTiles"`
}

type MapLevel struct {
	LayerInstances []MapLayer `json:"layerInstances"`
}

type MapData struct {
	Levels []MapLevel `json:"levels"`
}

func ParseWorldTiles(path string, tileset *ebiten.Image, tileW, tileH int) ([]*component.Transform, []*component.Sprite, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	return ParseWorldTilesFrom(file, tileset, tileW, tileH)
}

// Функція для парсингу карти з []byte (LDtk/JSON)
func ParseWorldTilesFrom(data []byte, tileset *ebiten.Image, tileW, tileH int) ([]*component.Transform, []*component.Sprite, error) {
	var mapData MapData
	if err := json.Unmarshal(data, &mapData); err != nil {
		return nil, nil, err
	}

	var transforms []*component.Transform
	var sprites []*component.Sprite

	// Діагностика структури даних
	println("[DEBUG] Levels count:", len(mapData.Levels))
	if len(mapData.Levels) == 0 {
		println("[DEBUG] No levels found in map data")
		return transforms, sprites, nil
	}

	println("[DEBUG] Layers count:", len(mapData.Levels[0].LayerInstances))
	if len(mapData.Levels[0].LayerInstances) == 0 {
		println("[DEBUG] No layers found in first level")
		return transforms, sprites, nil
	}

	// Знаходимо межі карти, проходячи по всіх тайлах
	minX, minY := 10, 10
	maxX, maxY := -10, -10
	totalTiles := 0

	// Спочатку знаходимо межі карти
	for _, level := range mapData.Levels {
		for _, layer := range level.LayerInstances {
			for _, tile := range layer.GridTiles {
				totalTiles++

				// Обчислюємо позицію тайла в сітці
				tileX := tile.Px[0] / tileW
				tileY := tile.Px[1] / tileH

				// Діагностика перших кількох тайлів
				if totalTiles <= 10 {
					println("[DEBUG] Tile", totalTiles, "px:", tile.Px[0], tile.Px[1], "grid:", tileX, tileY, "src:", tile.Src[0], tile.Src[1])
				}

				if tileX < minX {
					minX = tileX
				}
				if tileX > maxX {
					maxX = tileX
				}
				if tileY < minY {
					minY = tileY
				}
				if tileY > maxY {
					maxY = tileY
				}
			}
		}
	}

	println("[DEBUG] Total tiles found:", totalTiles)
	println("[DEBUG] Grid bounds - minX:", minX, "maxX:", maxX, "minY:", minY, "maxY:", maxY)
	println("[DEBUG] Tile size - tileW:", tileW, "tileH:", tileH)

	if totalTiles == 0 {
		println("[DEBUG] No tiles found in map data")
		return transforms, sprites, nil
	}

	// Розміри карти в тайлах
	mapWidth := maxX - minX + 1
	mapHeight := maxY - minY + 1

	// Обчислюємо offset для центрування карти
	offsetX := float64(mapWidth-1) / 2.0
	offsetY := float64(mapHeight-1) / 2.0

	println("[DEBUG] Map size - width:", mapWidth, "height:", mapHeight)
	println("[DEBUG] Center offset - offsetX:", offsetX, "offsetY:", offsetY)

	// Тепер створюємо тайли з правильним offset
	debugCount := 0
	for _, level := range mapData.Levels {
		for _, layer := range level.LayerInstances {
			for _, tile := range layer.GridTiles {
				// Позиція тайла в світових координатах з урахуванням центрування
				worldX := float64(tile.Px[0])/float64(tileW) - offsetX
				worldY := float64(tile.Px[1])/float64(tileH) - offsetY

				pos := mgl64.Vec2{worldX, worldY}

				// Створюємо підзображення з тайлсету
				srcRect := image.Rect(tile.Src[0], tile.Src[1], tile.Src[0]+tileW, tile.Src[1]+tileH)
				tileImg := tileset.SubImage(srcRect).(*ebiten.Image)

				// Діагностика перших кількох тайлів
				if debugCount < 5 {
					c := tileImg.At(0, 0)
					b := tileImg.Bounds()
					r, g, b2, a := c.RGBA()
					println("[DEBUG] Tile", debugCount, "world pos:", worldX, worldY,
						"srcRect:", srcRect.Min.X, srcRect.Min.Y, srcRect.Max.X, srcRect.Max.Y,
						"bounds:", b.Min.X, b.Min.Y, b.Max.X, b.Max.Y,
						"pixel(0,0):", r, g, b2, a)
					debugCount++
				}

				// Масштабуємо тайл до розміру GameUnit
				scaledTileImg := ebiten.NewImage(int(component.GameUnit), int(component.GameUnit))
				op := &ebiten.DrawImageOptions{}
				scaleX := component.GameUnit / float64(tileW)
				scaleY := component.GameUnit / float64(tileH)
				op.GeoM.Scale(scaleX, scaleY)
				scaledTileImg.DrawImage(tileImg, op)

				// Створюємо компоненти
				tr := component.NewTransform(pos, mgl64.Vec2{1, 1})
				spr := component.NewSprite(scaledTileImg, int(component.GameUnit), int(component.GameUnit))
				spr.IsTile = true

				transforms = append(transforms, tr)
				sprites = append(sprites, spr)
			}
		}
	}

	println("[DEBUG] Created", len(transforms), "tile transforms and", len(sprites), "tile sprites")

	return transforms, sprites, nil
}

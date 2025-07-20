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
	var mapData MapData
	if err := json.Unmarshal(file, &mapData); err != nil {
		return nil, nil, err
	}
	var transforms []*component.Transform
	var sprites []*component.Sprite
	// Центрування: визначаємо розмір карти в тайлах
	totalW, totalH := 0, 0
	if len(mapData.Levels) > 0 && len(mapData.Levels[0].LayerInstances) > 0 {
		layer := mapData.Levels[0].LayerInstances[0]
		for _, tile := range layer.GridTiles {
			if tile.Px[0] > totalW {
				totalW = tile.Px[0]
			}
			if tile.Px[1] > totalH {
				totalH = tile.Px[1]
			}
		}
		totalW = totalW/tileW + 1
		totalH = totalH/tileH + 1
		println("[DEBUG] totalW:", totalW, "totalH:", totalH)
	}
	offsetX := float64(totalW-1) / 2.0
	offsetY := float64(totalH-1) / 2.0
	println("[DEBUG] offsetX:", offsetX, "offsetY:", offsetY)
	debugCount := 0
	for _, level := range mapData.Levels {
		for _, layer := range level.LayerInstances {
			for _, tile := range layer.GridTiles {
				pos := mgl64.Vec2{
					float64(tile.Px[0])/float64(tileW) - offsetX,
					float64(tile.Px[1])/float64(tileH) - offsetY,
				}
				// Повертаємо вирізання з тайлсету
				srcRect := image.Rect(tile.Src[0], tile.Src[1], tile.Src[0]+tileW, tile.Src[1]+tileH)
				tileImg := tileset.SubImage(srcRect).(*ebiten.Image)
				if debugCount < 5 {
					c := tileImg.At(0, 0)
					b := tileImg.Bounds()
					r, g, b2, a := c.RGBA()
					println("[DEBUG] Tile pos:", pos.X(), pos.Y(), "srcRect:", srcRect.Min.X, srcRect.Min.Y, srcRect.Max.X, srcRect.Max.Y, "tileImg.Bounds:", b.Min.X, b.Min.Y, b.Max.X, b.Max.Y, "pixel(0,0):", r, g, b2, a)
					debugCount++
				}
				// Масштабуємо tileImg до 32x32 (GameUnit)
				scaledTileImg := ebiten.NewImage(int(component.GameUnit), int(component.GameUnit))
				op := &ebiten.DrawImageOptions{}
				scaleX := component.GameUnit / float64(tileW)
				scaleY := component.GameUnit / float64(tileH)
				op.GeoM.Scale(scaleX, scaleY)
				scaledTileImg.DrawImage(tileImg, op)
				tr := component.NewTransform(pos, mgl64.Vec2{1, 1}) // розмір 1x1 у системних координатах
				spr := component.NewSprite(scaledTileImg, int(component.GameUnit), int(component.GameUnit))
				spr.IsTile = true // Додаю маркування для тайлів
				transforms = append(transforms, tr)
				sprites = append(sprites, spr)
			}
		}
	}
	return transforms, sprites, nil
}

// Додаю функцію для парсингу карти з []byte (LDtk/JSON)
func ParseWorldTilesFromBytes(data []byte, tileset *ebiten.Image, tileW, tileH int) ([]*component.Transform, []*component.Sprite, error) {
	// Підтримує як main_map.json, так і LDtk-файли (map_1.ldtk)
	var mapData MapData
	if err := json.Unmarshal(data, &mapData); err != nil {
		return nil, nil, err
	}
	var transforms []*component.Transform
	var sprites []*component.Sprite
	// Центрування: визначаємо розмір карти в тайлах
	totalW, totalH := 0, 0
	if len(mapData.Levels) > 0 && len(mapData.Levels[0].LayerInstances) > 0 {
		layer := mapData.Levels[0].LayerInstances[0]
		for _, tile := range layer.GridTiles {
			tileX := tile.Px[0] / tileW
			tileY := tile.Px[1] / tileH
			if tileX > totalW {
				totalW = tileX
			}
			if tileY > totalH {
				totalH = tileY
			}
		}
		totalW = totalW + 1
		totalH = totalH + 1
		println("[DEBUG] totalW:", totalW, "totalH:", totalH)
	}
	offsetX := float64(totalW-1) / 2.0
	offsetY := float64(totalH-1) / 2.0
	println("[DEBUG] offsetX:", offsetX, "offsetY:", offsetY)
	debugCount := 0
	for _, level := range mapData.Levels {
		for _, layer := range level.LayerInstances {
			for _, tile := range layer.GridTiles {
				pos := mgl64.Vec2{
					float64(tile.Px[0]) / float64(tileW),
					float64(tile.Px[1]) / float64(tileH),
				}
				srcRect := image.Rect(tile.Src[0], tile.Src[1], tile.Src[0]+tileW, tile.Src[1]+tileH)
				tileImg := tileset.SubImage(srcRect).(*ebiten.Image)
				if debugCount < 5 {
					c := tileImg.At(0, 0)
					b := tileImg.Bounds()
					r, g, b2, a := c.RGBA()
					println("[DEBUG] Tile pos:", pos.X(), pos.Y(), "srcRect:", srcRect.Min.X, srcRect.Min.Y, srcRect.Max.X, srcRect.Max.Y, "tileImg.Bounds:", b.Min.X, b.Min.Y, b.Max.X, b.Max.Y, "pixel(0,0):", r, g, b2, a)
					debugCount++
				}
				// Масштабуємо tileImg до 32x32 (GameUnit)
				scaledTileImg := ebiten.NewImage(int(component.GameUnit), int(component.GameUnit))
				op := &ebiten.DrawImageOptions{}
				scaleX := component.GameUnit / float64(tileW)
				scaleY := component.GameUnit / float64(tileH)
				op.GeoM.Scale(scaleX, scaleY)
				scaledTileImg.DrawImage(tileImg, op)
				tr := component.NewTransform(pos, mgl64.Vec2{1, 1})
				spr := component.NewSprite(scaledTileImg, int(component.GameUnit), int(component.GameUnit))
				spr.IsTile = true // Додаю маркування для тайлів
				transforms = append(transforms, tr)
				sprites = append(sprites, spr)
			}
		}
	}
	return transforms, sprites, nil
}

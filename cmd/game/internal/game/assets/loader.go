package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed images/** maps/** maps/PNG/**
var assetsFS embed.FS

type AssetManager struct {
	images map[string]*ebiten.Image
}

func NewAssetManager() *AssetManager {
	return &AssetManager{
		images: make(map[string]*ebiten.Image),
	}
}

func (am *AssetManager) LoadImage(path string) (*ebiten.Image, error) {
	if img, exists := am.images[path]; exists {
		return img, nil
	}

	file, err := assetsFS.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	ebitenImg := ebiten.NewImageFromImage(img)
	am.images[path] = ebitenImg

	return ebitenImg, nil
}

// Альтернативний спосіб завантаження (якщо не використовуєш embed)
func (am *AssetManager) LoadImageFromFile(path string) (*ebiten.Image, error) {
	if img, exists := am.images[path]; exists {
		return img, nil
	}

	ebitenImg, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	am.images[path] = ebitenImg
	return ebitenImg, nil
}

// Додаю метод для читання карти з embed.FS (підтримує LDtk/JSON)
func (am *AssetManager) LoadMap(path string) ([]byte, error) {
	return assetsFS.ReadFile(path)
}

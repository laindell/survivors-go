package screen

import "github.com/hajimehoshi/ebiten/v2"

type Manager struct {
	screen *ebiten.Image
}

func NewManager() *Manager {
	return &Manager{
		screen: ebiten.NewImage(800, 600), // Default size
	}
}

func (m *Manager) UpdateScreen(screen *ebiten.Image) {
	m.screen = screen
}

func (m *Manager) Screen() *ebiten.Image {
	return m.screen
}

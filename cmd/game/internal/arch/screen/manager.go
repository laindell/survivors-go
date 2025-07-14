package screen

import "github.com/hajimehoshi/ebiten/v2"

type Manager struct {
	screen     *ebiten.Image // reference to екран (для масштабування)
	buffer     *ebiten.Image // власний буфер для рендеру ігрової сцени
	bufW, bufH int
}

func NewManager() *Manager {
	w, h := 800, 600 // Default size буде оновлено через SetBufferSize
	return &Manager{
		screen: nil,
		buffer: ebiten.NewImage(w, h),
		bufW:   w,
		bufH:   h,
	}
}

func (m *Manager) UpdateScreen(screen *ebiten.Image) {
	m.screen = screen
}

// Встановити розмір буфера (ігрового поля)
func (m *Manager) SetBufferSize(w, h int) {
	if w != m.bufW || h != m.bufH {
		m.buffer = ebiten.NewImage(w, h)
		m.bufW = w
		m.bufH = h
	}
}

// Повертає буфер для рендеру ігрової сцени
func (m *Manager) Buffer() *ebiten.Image {
	return m.buffer
}

// Повертає screen (екран для фінального рендеру)
func (m *Manager) Screen() *ebiten.Image {
	return m.screen
}

// Додаю геттери для розмірів буфера
func (m *Manager) BufW() int {
	return m.bufW
}

func (m *Manager) BufH() int {
	return m.bufH
}

package component

import (
	"image"

	"github.com/go-glx/ecs/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	Image        *ebiten.Image // Повний спрайтшит
	FrameWidth   int
	FrameHeight  int
	Direction    int           // 0 - вниз, 1 - вліво, 2 - вправо, 3 - вгору
	Frame        int           // Номер кадру анімації (0-7)
	IdleImage    *ebiten.Image // Idle-спрайт
	WalkingImage *ebiten.Image // Walking-спрайт
}

func NewSprite(img *ebiten.Image, frameW, frameH int) *Sprite {
	return &Sprite{
		Image:       img,
		FrameWidth:  frameW,
		FrameHeight: frameH,
		Direction:   0,
		Frame:       0,
	}
}

func (s *Sprite) GetFrameRect() image.Rectangle {
	x := s.Frame * s.FrameWidth
	y := s.Direction * s.FrameHeight
	return image.Rect(x, y, x+s.FrameWidth, y+s.FrameHeight)
}

func (s Sprite) TypeID() ecs.ComponentTypeID {
	return ecs.ComponentTypeID("Sprite")
}

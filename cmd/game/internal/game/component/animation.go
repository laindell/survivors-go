package component

import (
	"time"

	"github.com/go-glx/ecs/ecs"
)

type Animation struct {
	CurrentFrame int           // Поточний кадр анімації
	FrameCount   int           // Загальна кількість кадрів
	FrameTime    time.Duration // Тривалість одного кадру
	LastUpdate   time.Time     // Час останнього оновлення
	IsPlaying    bool          // Чи грає анімація
	Loop         bool          // Чи зациклена анімація
	Speed        float64       // Швидкість анімації (множник)
}

func NewAnimation(frameCount int, frameTime time.Duration, loop bool) *Animation {
	return &Animation{
		CurrentFrame: 0,
		FrameCount:   frameCount,
		FrameTime:    frameTime,
		LastUpdate:   time.Now(),
		IsPlaying:    true,
		Loop:         loop,
		Speed:        1.0,
	}
}

func (a Animation) TypeID() ecs.ComponentTypeID {
	return "Animation"
}

func (a *Animation) Play() {
	if !a.IsPlaying {
		a.IsPlaying = true
		a.LastUpdate = time.Now()
	}
}

func (a *Animation) Pause() {
	a.IsPlaying = false
}

func (a *Animation) Reset() {
	a.CurrentFrame = 0
	a.LastUpdate = time.Now()
}

func (a *Animation) SetSpeed(speed float64) {
	a.Speed = speed
}

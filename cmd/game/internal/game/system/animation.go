package system

import (
	"github.com/go-glx/ecs/ecs"
	"survivors-go/cmd/game/internal/game/component"
	"time"
)

type AnimationSystem struct{}

func NewAnimationSystem() *AnimationSystem {
	return &AnimationSystem{}
}

func (s *AnimationSystem) TypeID() ecs.SystemTypeID {
	return "AnimationSystem"
}

func (s *AnimationSystem) OnUpdate(w ecs.RuntimeWorld) {
	filter := ecs.NewFilter2[component.Animation, component.Sprite](w).Find()

	for filter.Next() {
		_, animation, sprite := filter.Get()

		if !animation.IsPlaying {
			continue
		}

		now := time.Now()
		frameTime := time.Duration(float64(animation.FrameTime) / animation.Speed)

		if now.Sub(animation.LastUpdate) >= frameTime {
			animation.CurrentFrame++

			if animation.CurrentFrame >= animation.FrameCount {
				if animation.Loop {
					animation.CurrentFrame = 0
				} else {
					animation.CurrentFrame = animation.FrameCount - 1
					animation.IsPlaying = false
				}
			}

			// Оновлюємо кадр у спрайті
			sprite.Frame = animation.CurrentFrame
			animation.LastUpdate = now
		}
	}
}

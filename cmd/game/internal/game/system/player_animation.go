package system

import (
	"math"
	"survivors-go/cmd/game/internal/game/component"
	"survivors-go/cmd/game/internal/game/system/internal"

	"github.com/go-glx/ecs/ecs"
)

type PlayerAnimationSystem struct{}

func NewPlayerAnimationSystem() *PlayerAnimationSystem {
	return &PlayerAnimationSystem{}
}

func (s *PlayerAnimationSystem) TypeID() ecs.SystemTypeID {
	return "PlayerAnimationSystem"
}

func (s *PlayerAnimationSystem) OnUpdate(w ecs.RuntimeWorld) {
	filter := ecs.NewFilter2[component.Player, component.Sprite](w).Find()
	for filter.Next() {
		_, _, sprite := filter.Get()
		// Шукаємо компонент Animation через world
		animationFilter := ecs.NewFilter2[component.Animation, component.Sprite](w).Find()
		var animation *component.Animation
		for animationFilter.Next() {
			_, animationCandidate, spriteCandidate := animationFilter.Get()
			if spriteCandidate == sprite {
				animation = animationCandidate
				break
			}
		}
		if animation == nil {
			continue
		}

		// Walking image для повернення
		walkingImage := sprite.WalkingImage

		dx, dy := internal.GetMoveDirection()

		if dx != 0 || dy != 0 {
			// Якщо рухається — повертаємо walking-спрайт
			if walkingImage != nil && sprite.Image != walkingImage {
				sprite.Image = walkingImage
				animation.FrameCount = 8        // Кількість кадрів у walking
				animation.FrameTime = 100 * 1e6 // 100ms
				animation.Loop = true
			}

			// Визначаємо напрямок для анімації
			if math.Abs(dx) > math.Abs(dy) {
				if dx > 0 {
					sprite.Direction = 2 // вправо
				} else {
					sprite.Direction = 1 // вліво
				}
			} else {
				if dy > 0 {
					sprite.Direction = 0 // вниз
				} else {
					sprite.Direction = 3 // вгору
				}
			}

			if !animation.IsPlaying {
				animation.Play()
			}
			animation.SetSpeed(1.0)

		} else {
			// Якщо стоїть — підміняємо на idle-спрайт
			if sprite.IdleImage != nil && sprite.Image != sprite.IdleImage {
				sprite.Image = sprite.IdleImage
				animation.FrameCount = 8        // Кількість кадрів у idle
				animation.FrameTime = 200 * 1e6 // 200ms
				animation.Loop = true
			}
			animation.Play()
		}
	}
}

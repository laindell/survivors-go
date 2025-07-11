package component

import "github.com/go-glx/ecs/ecs"

type Player struct {
	Speed float64
}

func NewPlayer(speed float64) *Player {
	return &Player{Speed: speed}
}

func (p Player) TypeID() ecs.ComponentTypeID {
	return "Player"
}

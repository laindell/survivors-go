package arch

import "survivors-go/cmd/game/internal/arch/screen"

type EngineImpl struct {
	screenManager *screen.Manager
}

func NewEngineImpl(screenManager *screen.Manager) *EngineImpl {
	return &EngineImpl{screenManager: screenManager}
}

func (e *EngineImpl) ScreenManager() ScreenManager {
	return e.screenManager
}

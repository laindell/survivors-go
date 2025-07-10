package game

// Глобальні змінні для меж світу
var (
	WorldMinX int
	WorldMaxX int
	WorldMinY int
	WorldMaxY int
)

func InitWorldParams(minX, maxX, minY, maxY int) {
	WorldMinX = minX
	WorldMaxX = maxX
	WorldMinY = minY
	WorldMaxY = maxY
}

func WorldWidth() int {
	return WorldMaxX - WorldMinX
}

func WorldHeight() int {
	return WorldMaxY - WorldMinY
}

// WorldToScreen переводить координати світу у координати екрану так, щоб (0,0) було в центрі
func WorldToScreen(x, y float64, unit float64) (float64, float64) {
	centerX := float64(WorldWidth()) * unit / 2
	centerY := float64(WorldHeight()) * unit / 2
	screenX := centerX + x*unit
	screenY := centerY + y*unit
	return screenX, screenY
}

package game

func Scence1(g *SpaceshipGame) {

	g.UpdateCounter++

	if shouldDo(g.GameConfig.TPS, g.UpdateCounter, g.GameConfig.EnemiesPerSecond) {
		g.PutNewEnemies(1)
	}

	if shouldDo(g.GameConfig.TPS, g.UpdateCounter, g.GameConfig.StarsPerSecond) {
		g.PutStars(1)
	}

	if shouldDo(g.GameConfig.TPS, g.UpdateCounter, g.GameConfig.AmmoPerSecond) {
		g.PutNewAmmos(1)
	}

	if g.UpdateCounter%10000 == 0 {
		g.UpdateCounter = 0
	}
}

func shouldDo(TPS float64, currentTick int, actionPerSecond float64) bool {

	if actionPerSecond == 0 {
		return false
	}

	var x int
	rate := TPS / actionPerSecond
	if rate > 1 {
		x = int(rate)
	} else {
		x = 1
	}

	return currentTick%x == 0
}

package game

func Scence1(g *SpaceshipGame) {

	g.TickCounter++

	if shouldDo(g.GameConfig.TPS, g.TickCounter, g.GameConfig.EnemiesPerSecond) {
		g.PutNewEnemies(1)
	}

	if shouldDo(g.GameConfig.TPS, g.TickCounter, g.GameConfig.StarsPerSecond) {
		g.PutStars(1)
	}

	if shouldDo(g.GameConfig.TPS, g.TickCounter, g.GameConfig.AmmoPerSecond) {
		g.PutNewAmmos(1)
	}

	if g.TickCounter%10000 == 0 {
		g.TickCounter = 0
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

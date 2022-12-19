package models

type GameConfig struct {
	TPS  float64
	MaxX int
	MaxY int
}

func GameConfig1() GameConfig {
	return GameConfig{
		TPS:  60,
		MaxX: 2000,
		MaxY: 1000,
	}
}

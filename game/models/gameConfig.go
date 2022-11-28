package models

type GameConfig struct {
	TPS float64

	BulletCountSpaceship1 int
	BulletCountSpaceship2 int

	HealthSpaceship1 int
	HealthSpaceship2 int

	BulletVelocity int
	StarVelocity   int

	StarsPerSecond   float64
	EnemiesPerSecond float64
	AmmoPerSecond    float64
}

func GameConfig1() GameConfig {

	return GameConfig{

		TPS: 60,

		BulletCountSpaceship1: 10,
		BulletCountSpaceship2: 10,

		HealthSpaceship1: 10,
		HealthSpaceship2: 10,

		BulletVelocity: 7,
		StarVelocity:   6,

		StarsPerSecond:   0,
		AmmoPerSecond:    0.3,
		EnemiesPerSecond: 0.1,
	}
}

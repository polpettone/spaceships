package game

type GameConfig struct {
	TPS float64

	BulletCountSpaceship1 int
	BulletCountSpaceship2 int

	HealthSpaceship1 int
	HealthSpaceship2 int

	StarsPerSecond   float64
	EnemiesPerSecond float64
	AmmoPerSecond    float64
}

func gameConfig1() GameConfig {

	return GameConfig{

		TPS: 60,

		BulletCountSpaceship1: 10,
		BulletCountSpaceship2: 10,

		HealthSpaceship1: 10,
		HealthSpaceship2: 10,

		StarsPerSecond: 5,

		AmmoPerSecond: 1,

		EnemiesPerSecond: 0,
	}
}

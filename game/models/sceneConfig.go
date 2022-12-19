package models

type SceneConfig struct {
	GameConfig GameConfig

	BulletCountSpaceship1 int
	BulletCountSpaceship2 int

	HealthSpaceship1 int
	HealthSpaceship2 int

	BulletVelocity int
	StarVelocity   int

	StarsPerSecond   float64
	EnemiesPerSecond float64
	AmmoPerSecond    float64

	InitialPosSpaceship1 Pos
	InitialPosSpaceship2 Pos
}

func SceneConfig1() SceneConfig {

	return SceneConfig{

		GameConfig: GameConfig{TPS: 60},

		BulletCountSpaceship1: 10,
		BulletCountSpaceship2: 10,

		HealthSpaceship1: 10,
		HealthSpaceship2: 10,

		BulletVelocity: 7,
		StarVelocity:   6,

		StarsPerSecond:   0,
		AmmoPerSecond:    0.3,
		EnemiesPerSecond: 0.1,

		InitialPosSpaceship1: NewPos(100, 300),
		InitialPosSpaceship2: NewPos(1900, 300),
	}
}

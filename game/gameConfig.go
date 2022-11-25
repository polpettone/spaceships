package game

type GameConfig struct {
	BulletCountSpaceship1 int
	BulletCountSpaceship2 int

	HealthSpaceship1 int
	HealthSpaceship2 int
}

func gameConfig1() GameConfig {

	return GameConfig{
		BulletCountSpaceship1: 10,
		BulletCountSpaceship2: 10,
		HealthSpaceship1:      10,
		HealthSpaceship2:      10,
	}
}

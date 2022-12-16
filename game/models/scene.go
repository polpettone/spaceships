package models

import "github.com/hajimehoshi/ebiten"

type Scene interface {
	GetMaxX() int
	GetMaxY() int

	AddGameObject(o GameObject)
	GetGameObjects() map[string]GameObject
	GetSpaceship1() *Spaceship
	GetSpaceship2() *Spaceship

	PutNewAmmos(count int)
	PutStars(count int)
	PutNewEnemies(count int)

	Update(screen *ebiten.Image) error
	Draw(screen *ebiten.Image)
	Reset()

	GetConfig() GameConfig

	GetTickCounter() int

	CheckGameOverCriteria() bool
}

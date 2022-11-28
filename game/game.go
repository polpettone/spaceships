package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/polpettone/gaming/spaceships/game/models"
)

type Game interface {
	GetMaxX() int
	GetMaxY() int
	AddGameObject(o models.GameObject)
	GetGameObjects() map[string]models.GameObject
	GetSpaceship1() *Spaceship
	GetSpaceship2() *Spaceship
	GetUpdateCounter() int
	Layout(outsideWidth, outsideHeight int) (int, int)
	Update(screen *ebiten.Image) error
}

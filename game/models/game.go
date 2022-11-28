package models

import (
	"github.com/hajimehoshi/ebiten"
)

type Game interface {
	GetMaxX() int
	GetMaxY() int
	AddGameObject(o GameObject)
	GetGameObjects() map[string]GameObject
	GetSpaceship1() *Spaceship
	GetSpaceship2() *Spaceship
	GetUpdateCounter() int
	Layout(outsideWidth, outsideHeight int) (int, int)
	Update(screen *ebiten.Image) error
}

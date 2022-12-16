package models

import "github.com/hajimehoshi/ebiten"

type Menu struct {
	Spaceship1  *Spaceship
	Spaceship2  *Spaceship
	GameObjects map[string]GameObject
	GameConfig  GameConfig
	TickCounter int
	MaxX        int
	MaxY        int
}

func (g *Menu) GetConfig() GameConfig {
	return g.GameConfig
}

func (g *Menu) GetTickCounter() int {
	return g.TickCounter
}

func (g *Menu) GetMaxX() int {
	return g.MaxX
}

func (g *Menu) GetMaxY() int {
	return g.MaxY
}

func (g *Menu) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Menu) Draw(screen *ebiten.Image) {
}

func (g *Menu) Reset() {
}

func (g *Menu) AddGameObject(o GameObject) {
}

func (g *Menu) GetGameObjects() map[string]GameObject {
	return g.GameObjects
}

func (g *Menu) GetSpaceship1() *Spaceship {
	return g.Spaceship1
}

func (g *Menu) GetSpaceship2() *Spaceship {
	return g.Spaceship2
}

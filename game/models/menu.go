package models

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/polpettone/gaming/spaceships/engine"
)

type Menu struct {
	Spaceship1  *Spaceship
	Spaceship2  *Spaceship
	GameObjects map[string]GameObject
	GameConfig  GameConfig
	TickCounter int
	MaxX        int
	MaxY        int
}

func NewMenu(config GameConfig) (*Menu, error) {

	return &Menu{
		GameConfig: config,
	}, nil
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

	t := fmt.Sprintf(
		`   SpaceShips 

		Hit Enter to start 
			Press Q for Quit`)

	text.Draw(screen, t, engine.MplusBigFont, 700, 300, color.White)
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

func (g *Menu) PutNewAmmos(count int)   {}
func (g *Menu) PutStars(count int)      {}
func (g *Menu) PutNewEnemies(count int) {}

func (g *Menu) CheckGameOverCriteria() bool {
	return false
}

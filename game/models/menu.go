package models

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/polpettone/gaming/spaceships/engine"
)

type Menu struct {
	GameConfig    GameConfig
	Scenes        map[string]Scene
	SelectedScene Scene
}

func NewMenu(config GameConfig, scenes map[string]Scene) (*Menu, error) {
	return &Menu{
		GameConfig:    config,
		Scenes:        scenes,
		SelectedScene: scenes["1"],
	}, nil
}

func (g *Menu) GetSelectedScene() Scene {
	return g.SelectedScene
}

func (g *Menu) GetConfig() GameConfig {
	return g.GameConfig
}

func (g *Menu) Update() (GameState, error) {

	if inpututil.IsKeyJustPressed(ebiten.Key1) {
		g.SelectedScene = g.Scenes["1"]
	}

	if inpututil.IsKeyJustPressed(ebiten.Key2) {
		g.SelectedScene = g.Scenes["2"]
	}

	return ShowMenu, nil
}

func (g *Menu) Draw(screen *ebiten.Image) {

	t := fmt.Sprintf(
		`   SpaceShips 

Selected: %s

Press Enter to start 
Press Q to quit`,
		g.GetSelectedScene().GetName())

	text.Draw(screen, t, engine.MplusBigFont, 700, 300, color.White)
}

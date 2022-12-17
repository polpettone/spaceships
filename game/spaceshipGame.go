package game

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/polpettone/gaming/spaceships/engine"
	"github.com/polpettone/gaming/spaceships/game/models"
)

const (
	screenWidth  = 2000
	screenHeight = 1000
)

type Scene func(g *SpaceshipGame)

type SpaceshipGame struct {
	MaxX            int
	MaxY            int
	GameConfig      models.GameConfig
	BackgroundImage *ebiten.Image
	DebugScreen     *DebugScreen

	DebugPrint bool

	BackgroundSound *audio.Player

	TickCounter int

	SoundOn bool

	State models.GameState

	GamepadIDs map[int]struct{}

	CurrentScene models.Scene
}

func NewGame(config models.GameConfig, scene models.Scene) (models.Game, error) {

	debugScreen, err := NewDebugScreen()
	if err != nil {
		return nil, err
	}

	g := &SpaceshipGame{
		GameConfig:   config,
		DebugScreen:  debugScreen,
		MaxX:         screenWidth,
		MaxY:         screenHeight,
		DebugPrint:   false,
		SoundOn:      false,
		State:        models.Running,
		GamepadIDs:   map[int]struct{}{},
		CurrentScene: scene,
	}

	return g, nil
}

func (g *SpaceshipGame) Reset() {

	g.CurrentScene.Reset()

	g.TickCounter = 0
}

func (g *SpaceshipGame) SetState(state models.GameState) {
	if g.State == state {
		return
	}
	fmt.Printf("Change State %d -> %d \n", g.State, state)
	g.State = state
}

func (g *SpaceshipGame) GetConfig() models.GameConfig {
	return g.GameConfig
}

func (g *SpaceshipGame) GetMaxX() int {
	return g.MaxX
}

func (g *SpaceshipGame) GetMaxY() int {
	return g.MaxY
}

func handleBackgroundSound(g *SpaceshipGame) {
	if g.BackgroundSound != nil {
		if g.State == models.Pause {
			g.BackgroundSound.Pause()
		}

		if !g.BackgroundSound.IsPlaying() && g.SoundOn {
			g.BackgroundSound.Play()
		}

		if g.BackgroundSound.IsPlaying() && !g.SoundOn {
			g.BackgroundSound.Pause()
		}
	}
}

func (g *SpaceshipGame) Update(screen *ebiten.Image) error {

	updateGamepads(g)

	state := handleControl(g.State)
	g.SetState(state)

	if g.State == models.Quit {
		os.Exit(0)
	}

	if g.State == models.Reset {
		g.Reset()
		g.SetState(models.Running)
		return nil
	}

	if g.State == models.MenuPreparation {
		g.Reset()
		nextScene, err := models.NewMenu(g.GameConfig)
		if err != nil {
			return err
		}
		g.CurrentScene = nextScene
		g.SetState(models.ShowMenu)
		return nil
	}

	if g.State == models.ScenePreparation {
		g.Reset()
		nextScene, err := models.NewScene1(g.GameConfig)
		if err != nil {
			return err
		}
		g.CurrentScene = nextScene
		g.SetState(models.Running)
		return nil
	}

	if g.State == models.GameOver {
		return nil
	}

	if g.State == models.Pause {
		return nil
	}

	g.SoundOn = handleSoundControl(g.SoundOn)
	g.DebugPrint = handleDebugPrintControl(g.DebugPrint)
	g.DebugScreen.Update(g.CurrentScene)

	state, err := g.CurrentScene.Update(screen)

	if err != nil {
		return err
	}

	g.SetState(state)

	return nil
}

func (g *SpaceshipGame) Draw(screen *ebiten.Image) {

	g.CurrentScene.Draw(screen)

	if g.DebugPrint {
		g.DebugScreen.Draw(screen, g)
	}

	if g.State == models.GameOver {
		drawGameOverScreen(g, screen)
	}
}

func (g *SpaceshipGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func drawGameOverScreen(g *SpaceshipGame, screen *ebiten.Image) {

	var livingSpaceship *models.Spaceship
	if g.CurrentScene.GetSpaceship1().Alive() {
		livingSpaceship = g.CurrentScene.GetSpaceship1()
	} else {
		livingSpaceship = g.CurrentScene.GetSpaceship1()
	}

	t := fmt.Sprintf(
		`GAME OVER 
%s has won
Press R for New Game
Press M for Menu
Press Q for Quit`,
		livingSpaceship.PilotName)

	text.Draw(screen, t, engine.MplusBigFont, 700, 300, color.White)
}

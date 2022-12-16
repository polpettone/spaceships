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

	Pause bool

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
		Pause:        false,
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
	g.Pause = false
	g.State = models.Running
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
		if g.Pause {
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

	if isQuitHit() {
		os.Exit(0)
	}

	if handleResetGameControl() && g.State == models.GameOver {
		g.Reset()
		g.State = models.Running
	}

	g.Pause = handlePauseControl(g.Pause)
	g.SoundOn = handleSoundControl(g.SoundOn)
	g.DebugPrint = handleDebugPrintControl(g.DebugPrint)

	if g.State == models.GameOver {
		return nil
	}

	if g.Pause {
		return nil
	}

	g.DebugScreen.Update(g.CurrentScene)

	state, err := g.CurrentScene.Update(screen)

	if err != nil {
		return err
	}

	g.State = state

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
Press Q for Quit`,
		livingSpaceship.PilotName)

	text.Draw(screen, t, engine.MplusBigFont, 700, 300, color.White)
}

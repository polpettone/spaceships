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

type Scene func(g *SpaceshipGame)

type SpaceshipGame struct {
	MaxX            int
	MaxY            int
	BackgroundImage *ebiten.Image
	DebugScreen     *DebugScreen

	DebugPrint bool

	BackgroundSound *audio.Player

	TickCounter int

	SoundOn bool

	State models.GameState

	GamepadIDs map[int]struct{}

	CurrentScene models.Scene

	Menu *models.Menu

	config models.GameConfig
}

func NewGame(menu *models.Menu, config models.GameConfig) (models.Game, error) {

	debugScreen, err := NewDebugScreen()

	if err != nil {
		return nil, err
	}

	g := &SpaceshipGame{
		DebugScreen:  debugScreen,
		MaxX:         config.MaxX,
		MaxY:         config.MaxY,
		DebugPrint:   false,
		SoundOn:      false,
		State:        models.ShowMenu,
		GamepadIDs:   map[int]struct{}{},
		CurrentScene: menu.GetSelectedScene(),
		Menu:         menu,
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

	state := handleStateControl(g.State)
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
		g.SetState(models.ShowMenu)
		return nil
	}

	if g.State == models.ScenePreparation {
		g.Reset()
		g.CurrentScene = g.Menu.GetSelectedScene()
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
	g.DebugScreen.ShortMode = handleDebugPrintModeControl(g.DebugScreen.ShortMode)
	g.DebugScreen.Update(g.CurrentScene)

	if state != models.ShowMenu {

		state, err := g.CurrentScene.Update(screen)

		if err != nil {
			return err
		}

		g.SetState(state)

	} else {
		g.Menu.Update()
	}

	return nil
}

func (g *SpaceshipGame) Draw(screen *ebiten.Image) {

	if g.State == models.ShowMenu {
		g.Menu.Draw(screen)
	} else {
		g.CurrentScene.Draw(screen)
	}

	if g.DebugPrint {
		g.DebugScreen.Draw(screen, g)
	}

	if g.State == models.GameOver {
		drawGameOverScreen(g, screen)
	}
}

func (g *SpaceshipGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.MaxX, g.MaxY
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
Press Q to quit`,
		livingSpaceship.PilotName)

	text.Draw(screen, t, engine.MplusBigFont, 700, 300, color.White)
}

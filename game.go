package main

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/polpettone/gaming/natalito/engine"
)

type GameState int64

const (
	Running GameState = iota
	GameOver
)

type Game struct {
	BackgroundImage *ebiten.Image
	Spaceship       *Spaceship
	GameObjects     map[string]GameObject
	DebugScreen     *DebugScreen
	MaxX            int
	MaxY            int
	DebugPrint      bool

	BackgroundSound *audio.Player

	UpdateCounter int

	Pause bool

	SoundOn bool

	KilledEnemies int

	State GameState
}

func NewGame() (*Game, error) {

	backgroundSound, err := engine.InitSoundPlayer(
		"assets/background-sound-1.mp3",
		audioContext)

	if err != nil {
		return nil, err
	}

	backgroundImage, _, err := ebitenutil.NewImageFromFile(
		"assets/earth-space-sunset.png",
		ebiten.FilterDefault)

	if err != nil {
		return nil, err
	}

	debugScreen, err := NewDebugScreen(500, screenHeight)
	if err != nil {
		return nil, err
	}

	spaceship, err := NewSpaceship(NewPos(100, 300))

	if err != nil {
		return nil, err
	}

	gameObjects := map[string]GameObject{}

	g := &Game{
		BackgroundImage: backgroundImage,
		GameObjects:     gameObjects,
		DebugScreen:     debugScreen,
		Spaceship:       spaceship,
		UpdateCounter:   0,
		MaxX:            screenWidth,
		MaxY:            screenHeight,
		DebugPrint:      false,
		Pause:           false,
		BackgroundSound: backgroundSound,
		SoundOn:         false,
		State:           Running,
	}

	return g, nil
}

func (g *Game) Reset() {
	g.GameObjects = map[string]GameObject{}
	g.Spaceship, _ = NewSpaceship(NewPos(100, 300))
	g.UpdateCounter = 0
	g.Pause = false
	g.State = Running
	g.KilledEnemies = 0
}

func (g *Game) Update(screen *ebiten.Image) error {

	checkGameOverCriteria(g)

	if isQuitHit() {
		os.Exit(0)
	}

	if handleResetGameControl() && g.State == GameOver {
		g.Reset()
		g.State = Running
	}

	g.Pause = handlePauseControl(g.Pause)
	g.SoundOn = handleSoundControl(g.SoundOn)

	if g.State == GameOver {
		return nil
	}

	if g.Pause {
		g.BackgroundSound.Pause()
		return nil
	}

	if !g.BackgroundSound.IsPlaying() && g.SoundOn {
		g.BackgroundSound.Play()
	}

	if g.BackgroundSound.IsPlaying() && !g.SoundOn {
		g.BackgroundSound.Pause()
	}

	g.DebugPrint = handleDebugPrintControl(g.DebugPrint)

	spaceshipCollisionDetection(g.Spaceship, g.GameObjects)

	bulletSkyObjectCollisionDetection(g)

	g.Spaceship.Update(g)

	for _, o := range g.GameObjects {
		o.Update(g)
	}

	g.DebugScreen.Update(g)

	putNewObjects(g)
	deleteObjectsOutOfView(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(g.BackgroundImage, op)

	for _, o := range g.GameObjects {
		o.Draw(screen)
	}

	g.Spaceship.Draw(screen)

	if g.DebugPrint {
		g.DebugScreen.Draw(screen)
	}

	drawGameState(g, screen)

	if g.State == GameOver {
		drawGameOverScreen(g, screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func checkGameOverCriteria(g *Game) {
	if g.Spaceship.Health < 0 {
		g.State = GameOver
	}
}

func bulletSkyObjectCollisionDetection(g *Game) {

	for _, o := range g.GameObjects {

		if o.GetType() == "bullet" {
			for k, x := range g.GameObjects {
				if x.GetType() == "skyObject" {
					oW, _ := o.GetSize()
					xW, _ := x.GetSize()
					if engine.CollisionDetection(
						o.GetPos().X,
						o.GetPos().Y,
						x.GetPos().X,
						x.GetPos().Y,
						oW,
						xW,
						0) {

						delete(g.GameObjects, k)
						g.KilledEnemies += 1

					}
				}
			}
		}
	}

}

func spaceshipCollisionDetection(s *Spaceship, gameObjects map[string]GameObject) {

	for _, o := range gameObjects {

		if o.GetType() == "skyObject" {

			sW, _ := s.GetSize()
			oW, _ := o.GetSize()
			if engine.CollisionDetection(
				s.Pos.X,
				s.Pos.Y,
				o.GetPos().X,
				o.GetPos().Y,
				sW,
				oW,
				0) {
				s.DamageCount += 1
				s.Health -= 1
			}
		}
	}
}

func drawGameState(g *Game, screen *ebiten.Image) {
	t := fmt.Sprintf(
		"Killed: %d  \n Health: %d \n Bullets %d",
		g.KilledEnemies,
		g.Spaceship.Health,
		g.Spaceship.BulletCount,
	)
	text.Draw(screen, t, engine.MplusNormalFont, 1800, 30, color.White)
}

func drawGameOverScreen(g *Game, screen *ebiten.Image) {
	t := fmt.Sprintf(
		"GAME OVER \n"+
			"You Killed %d Enemies \n"+
			"Press R for New Game \n"+
			"Press Q for Quit",
		g.KilledEnemies)
	text.Draw(screen, t, engine.MplusBigFont, 700, 300, color.White)
}

func putNewObjects(g *Game) {
	g.UpdateCounter++
	if g.UpdateCounter > 100 {
		g.UpdateCounter = 0
		newObjects := CreateSkyObjectAtRandomPosition(
			(screenWidth/3)*2, 0, screenWidth, screenHeight, 3)
		for _, nO := range newObjects {
			g.GameObjects[nO.GetID()] = nO
		}
	}
}

func deleteObjectsOutOfView(g *Game) {
	var ids []string
	for k, o := range g.GameObjects {
		x := o.GetPos().X
		y := o.GetPos().Y
		if x > g.MaxX || x < 0 || y > g.MaxX || y < 0 {
			ids = append(ids, k)
		}
	}
	for _, k := range ids {
		delete(g.GameObjects, k)
	}
}

package main

import (
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Game struct {
	BackgroundImage *ebiten.Image
	Spaceship       *Spaceship
	GameObjects     map[string]GameObject
	DebugScreen     *DebugScreen
	MaxX            int
	MaxY            int
	DebugPrint      bool
	AudioContext    *audio.Context

	BackgroundSound *audio.Player

	UpdateCounter int

	Pause bool
}

func NewGame() (*Game, error) {

	audioContext := audio.NewContext(44100)

	backgroundSound, err := InitSoundPlayer(
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

	spaceship, err := NewSpaceship(NewPos(100, 300))

	if err != nil {
		return nil, err
	}

	gameObjects := map[string]GameObject{}

	debugScreen, err := NewDebugScreen(500, screenHeight)
	if err != nil {
		return nil, err
	}

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
		AudioContext:    audioContext,
	}

	g.BackgroundSound.Play()

	return g, nil
}

func (g *Game) Update(screen *ebiten.Image) error {

	if isQuitHit() {
		os.Exit(0)
	}

	g.Pause = handlePauseControl(g.Pause)

	if g.Pause {
		g.BackgroundSound.Pause()
		return nil
	}

	if !g.BackgroundSound.IsPlaying() {
		g.BackgroundSound.Play()
	}

	for _, o := range g.GameObjects {

		oX := o.GetPos().X
		oY := o.GetPos().Y
		sX := g.Spaceship.Pos.X
		sY := g.Spaceship.Pos.Y
		sSize := g.Spaceship.Size

		if oX == sX && oY == sY ||
			((sX+sSize) == oX || sX == oX) && ((sY+sSize) == oY || (sY-sSize) == oY || sY == oY) {
			g.Spaceship.DamageCount += 1
		}

	}

	g.DebugPrint = handleDebugPrintControl(g.DebugPrint)

	g.Spaceship.Update(g)

	for _, o := range g.GameObjects {
		o.Update()
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
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
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

func isQuitHit() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return true
	}
	return false
}

func handlePauseControl(current bool) bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		return !current
	}
	return current
}

func handleDebugPrintControl(current bool) bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return !current
	}
	return current
}

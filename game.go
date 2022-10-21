package main

import (
	"image/color"

	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type Game struct {
	BackgroundImage *ebiten.Image
	Spaceship       *Spaceship
	GameObjects     map[string]GameObject
	DebugScreen     *DebugScreen

	UpdateCounter int
}

func NewGame() (*Game, error) {

	backgroundImage, err := ebiten.NewImage(
		screenWidth,
		screenHeight,
		ebiten.FilterDefault)

	if err != nil {
		return nil, err
	}
	backgroundImage.Fill(color.RGBA{240, 255, 240, 0xff})

	spaceship, err := NewSpaceship(NewPos(100, 300))

	if err != nil {
		return nil, err
	}

	gameObjects := map[string]GameObject{
		spaceship.ID: spaceship,
	}

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
	}

	return g, nil
}

func (g *Game) Update(screen *ebiten.Image) error {

	if isQuitHit() {
		os.Exit(0)
	}

	for _, o := range g.GameObjects {
		o.Update()
	}

	g.DebugScreen.Update(g)

	g.UpdateCounter++
	if g.UpdateCounter > 100 {
		g.UpdateCounter = 0

		newObjects := CreateSkyObjectAtRandomPosition(screenWidth, screenHeight, 3)

		for _, nO := range newObjects {
			g.GameObjects[nO.GetID()] = nO
		}

	}

	return nil
}

func isQuitHit() bool {

	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return true
	}

	return false
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	for _, o := range g.GameObjects {
		pos := o.GetPos()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(pos.X), float64(pos.Y))
		screen.DrawImage(o.GetImage(), op)
	}

	g.DebugScreen.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

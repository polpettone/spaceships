package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Game struct {
	BackgroundImage *ebiten.Image
	Spaceship       *Spaceship
	GameObjects     []GameObject
	DebugScreen     *DebugScreen
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

	gameObjects := []GameObject{
		spaceship,
	}
	gameObjects = append(gameObjects, CreateSkyObjects()...)

	debugScreen, err := NewDebugScreen(500, screenHeight)
	if err != nil {
		return nil, err
	}

	g := &Game{
		BackgroundImage: backgroundImage,
		GameObjects:     gameObjects,
		DebugScreen:     debugScreen,
	}

	return g, nil
}

func (g *Game) Update(screen *ebiten.Image) error {

	for _, o := range g.GameObjects {
		o.Update()
	}

	t :=
		`
	Debug Screen
	-----------------------------
	Game Object Count: %d
	`

	g.DebugScreen.SetText(
		fmt.Sprintf(t, len(g.GameObjects)),
	)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	for _, o := range g.GameObjects {
		pos := o.GetPos()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(pos.X), float64(pos.Y))
		screen.DrawImage(o.GetImage(), op)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(1500), float64(0))
	screen.DrawImage(g.DebugScreen.Image, op)

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

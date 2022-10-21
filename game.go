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
	MaxX            int
	MaxY            int
	DebugPrint      bool

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
	}

	return g, nil
}

func (g *Game) Update(screen *ebiten.Image) error {

	if isQuitHit() {
		os.Exit(0)
	}

	for _, o := range g.GameObjects {

		if o.GetPos().X == g.Spaceship.Pos.X &&
			o.GetPos().Y == g.Spaceship.Pos.Y {
			g.Spaceship.DamageCount += 1
		}

	}

	g.DebugPrint = handleDebugPrintControl(g.DebugPrint)

	g.Spaceship.Update(g.MaxX, g.MaxY)

	for _, o := range g.GameObjects {
		o.Update()
	}

	g.DebugScreen.Update(g)

	putNewObjects(g)
	deleteObjectsOutOfView(g)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

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
		newObjects := CreateSkyObjectAtRandomPosition((screenWidth/3)*2, 0, screenWidth, screenHeight, 3)
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

func handleDebugPrintControl(current bool) bool {

	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return !current
	} else {
		return current
	}

}

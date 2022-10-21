package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	natalitoImage *ebiten.Image
)

const (
	screenWidth  = 2000
	screenHeight = 700
)

func init() {
	var err error
	natalitoImage, _, err = ebitenutil.NewImageFromFile(
		"assets/natalito-front.png",
		ebiten.FilterDefault)

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Natalito")

	backgroundImage, err := ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
		return
	}
	backgroundImage.Fill(color.RGBA{240, 255, 240, 0xff})

	spaceship, err := NewSpaceship(NewPos(100, 300))
	if err != nil {
		log.Fatal(err)
		return
	}

	gameObjects := []GameObject{
		spaceship,
	}

	gameObjects = append(gameObjects, CreateSkyObjects()...)

	g := &Game{
		BackgroundImage: backgroundImage,
		GameObjects:     gameObjects,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

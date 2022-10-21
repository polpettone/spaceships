package main

import (
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

	g, err := NewGame()

	if err != nil {
		log.Fatal(err)
		return
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

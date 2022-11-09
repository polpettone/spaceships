package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

var (
	audioContext *audio.Context
)

const (
	screenWidth            = 1000
	screenHeight           = 1000
	spaceshipWallTolerance = 10
)

func init() {
	audioContext = audio.NewContext(44100)
}

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Natalito")
	ebiten.SetWindowResizable(true)

	g, err := NewGame()

	if err != nil {
		log.Fatal(err)
		return
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

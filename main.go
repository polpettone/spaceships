package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/polpettone/gaming/spaceships/game"
)

func main() {

	g, err := game.NewGame()

	if err != nil {
		log.Fatal(err)
		return
	}

	ebiten.SetWindowSize(g.GetMaxX(), g.GetMaxY())
	ebiten.SetWindowTitle("Spaceships")
	ebiten.SetWindowResizable(true)

	if err != nil {
		log.Fatal(err)
		return
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

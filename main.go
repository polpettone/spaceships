package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/polpettone/gaming/spaceships/game"
	"github.com/polpettone/gaming/spaceships/game/models"
)

func main() {

	scene, err := models.NewScene1(models.GameConfig1())

	g, err := game.NewGame(models.GameConfig1(), scene)

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

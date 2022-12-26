package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/polpettone/gaming/spaceships/game"
	"github.com/polpettone/gaming/spaceships/game/models"
)

func main() {

	simpleScene1, err := models.NewSimpleScene(
		"1 on 1 with enemies", models.SceneConfig1())
	if err != nil {
		log.Fatal(err)
		return
	}

	simpleScene2, err := models.NewSimpleScene(
		"Just one ship",
		models.SceneConfig2())
	if err != nil {
		log.Fatal(err)
		return
	}

	scene3, err := models.NewSimpleScene(
		"1 on 1 without enemies",
		models.SceneConfig3())
	if err != nil {
		log.Fatal(err)
		return
	}

	scene4, err := models.NewScene4(models.SceneConfig4())
	if err != nil {
		log.Fatal(err)
		return
	}

	scenes := map[string]models.Scene{
		"1": simpleScene1,
		"2": simpleScene2,
		"3": scene3,
		"4": scene4,
	}

	menu, err := models.NewMenu(scenes)
	if err != nil {
		log.Fatal(err)
		return
	}

	g, err := game.NewGame(menu, models.GameConfig1())

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

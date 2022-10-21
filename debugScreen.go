package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type DebugScreen struct {
	Text string
}

func NewDebugScreen(Width, Height int) (*DebugScreen, error) {
	return &DebugScreen{
		Text: "Debug Screen",
	}, nil
}

func (d *DebugScreen) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrintAt(
		screen,
		d.Text, 10, 10)
}

func (d *DebugScreen) Update(g *Game) {

	t :=
		`
	 TPS: %f2.2f
	 FPS: %f2.2f

	 UpdateCounter: %d 

	 SpaceShip Pos: %s

	 Game Object Count: %d

	 Game Objects: 
	 %s

	 `

	gameObjectsText := ""
	for _, o := range g.GameObjects {
		gameObjectsText += fmt.Sprintf(
			"%s - %s \n", o.GetID(), o.GetPos().Print(),
		)
	}

	spaceshipPos := "unknown"
	if g.Spaceship != nil {
		spaceshipPos = g.Spaceship.Pos.Print()
	}

	d.Text = fmt.Sprintf(
		t,
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
		g.UpdateCounter,
		spaceshipPos,
		len(g.GameObjects),
		gameObjectsText,
	)
}

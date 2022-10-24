package main

import (
	"fmt"
	"sort"

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

	 Spaceship Pos: %s

	 Spaceship Damage Count: %d

	 Game Object Count: %d

	 Game Objects: 
	 %s

	 `

	var keys []string

	for k, _ := range g.GameObjects {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	gameObjectsText := ""
	for _, k := range keys {
		gameObject := g.GameObjects[k]
		w, h := gameObject.GetSize()
		gameObjectsText += fmt.Sprintf(
			"%s - %s - (%d, %d) \n", gameObject.GetID(), gameObject.GetPos().Print(), w, h,
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
		g.Spaceship.DamageCount,
		len(g.GameObjects),
		gameObjectsText,
	)
}

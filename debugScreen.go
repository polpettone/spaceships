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

func (d *DebugScreen) Draw(screen *ebiten.Image, g *Game) {
	ebitenutil.DebugPrintAt(
		screen,
		d.Text, 10, 10)
	drawDebugCoordinate(screen, g)
}

func drawDebugCoordinate(screen *ebiten.Image, g *Game) {
	marginX := 55
	marginY := 20

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%d,%d", 0, 0),
		0, 0,
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%d,%d", g.MaxX/2, g.MaxY/2),
		g.MaxX/2-marginX, g.MaxY/2-marginY,
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%d,%d", 0, g.MaxY),
		0, g.MaxY-marginY,
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%d,%d", g.MaxX, 0),
		g.MaxX-marginX, 0,
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%d,%d", g.MaxX, g.MaxY),
		g.MaxX-marginX, g.MaxY-marginY,
	)
}

func (d *DebugScreen) Update(g *Game) {

	t :=
		`
	 TPS: %f2.2f
	 FPS: %f2.2f

	 Screen Size:  %d, %d

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
			"%s - %s - (%d, %d) - %s \n",
			gameObject.GetID(),
			gameObject.GetPos().Print(),
			w, h,
			gameObject.GetCentrePos().Print(),
		)
	}

	spaceshipPos := "n/a"
	centrePos := "n/a"
	if g.Spaceship != nil {
		spaceshipPos = g.Spaceship.Pos.Print()
		centrePos = g.Spaceship.GetCentrePos().Print()
	}
	sW, sH := g.Spaceship.GetSize()

	spaceshipText := fmt.Sprintf(
		"%s - %d,%d - %s", spaceshipPos, sW, sH, centrePos,
	)

	d.Text = fmt.Sprintf(
		t,
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
		g.MaxX,
		g.MaxY,
		g.UpdateCounter,
		spaceshipText,
		g.Spaceship.DamageCount,
		len(g.GameObjects),
		gameObjectsText,
	)
}

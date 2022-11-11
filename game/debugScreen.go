package game

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

func (d *DebugScreen) Draw(screen *ebiten.Image, g Game) {
	ebitenutil.DebugPrintAt(
		screen,
		d.Text, 10, 10)
	drawDebugCoordinate(screen, g)
}

func drawDebugCoordinate(screen *ebiten.Image, g Game) {
	marginX := 55
	marginY := 20

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%d,%d", 0, 0),
		0, 0,
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%d,%d", g.GetMaxX()/2, g.GetMaxY()/2),
		g.GetMaxX()/2-marginX, g.GetMaxY()/2-marginY,
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%d,%d", 0, g.GetMaxY()),
		0, g.GetMaxY()-marginY,
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%d,%d", g.GetMaxX(), 0),
		g.GetMaxX()-marginX, 0,
	)

	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%d,%d", g.GetMaxX(), g.GetMaxY()),
		g.GetMaxX()-marginX, g.GetMaxY()-marginY,
	)
}

func (d *DebugScreen) Update(g Game) {

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

	for k, _ := range g.GetGameObjects() {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	gameObjectsText := ""
	for _, k := range keys {
		gameObject := g.GetGameObjects()[k]
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
	if g.GetSpaceship() != nil {
		spaceshipPos = g.GetSpaceship().Pos.Print()
		centrePos = g.GetSpaceship().GetCentrePos().Print()
	}
	sW, sH := g.GetSpaceship().GetSize()

	spaceshipText := fmt.Sprintf(
		"%s - %d,%d - %s", spaceshipPos, sW, sH, centrePos,
	)

	d.Text = fmt.Sprintf(
		t,
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
		g.GetMaxX(),
		g.GetMaxY(),
		g.GetUpdateCounter(),
		spaceshipText,
		g.GetSpaceship().DamageCount,
		len(g.GetGameObjects()),
		gameObjectsText,
	)
}

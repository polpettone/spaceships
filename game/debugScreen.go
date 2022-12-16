package game

import (
	"fmt"
	"sort"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/polpettone/gaming/spaceships/game/models"
)

type DebugScreen struct {
	Text string
}

func NewDebugScreen() (*DebugScreen, error) {
	return &DebugScreen{
		Text: "Debug Screen",
	}, nil
}

func (d *DebugScreen) Draw(screen *ebiten.Image, g models.Game) {

	ebitenutil.DebugPrintAt(
		screen,
		d.Text, 10, 10)
	drawDebugCoordinate(screen, g)
}

func drawDebugCoordinate(screen *ebiten.Image, g models.Game) {
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

func (d *DebugScreen) Update(g models.Scene) {

	t :=
		`
	 TPS: %f2.2f
	 FPS: %f2.2f

	 Screen Size:  %d, %d

	 TickCounter: %d 

	 Spaceship 1 Pos: %s
	 Spaceship 2 Pos: %s

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

	d.Text = fmt.Sprintf(
		t,
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
		g.GetMaxX(),
		g.GetMaxY(),
		g.GetTickCounter(),
		spaceshipDebugInfos(g.GetSpaceship1()),
		spaceshipDebugInfos(g.GetSpaceship2()),
		len(g.GetGameObjects()),
		gameObjectsText,
	)
}

func spaceshipDebugInfos(s *models.Spaceship) string {
	if s == nil {
		return ""
	}
	spaceshipPos := "n/a"
	centrePos := "n/a"
	if s != nil {
		spaceshipPos = s.Pos.Print()
		centrePos = s.GetCentrePos().Print()
	}
	sW, sH := s.GetSize()
	spaceshipText := fmt.Sprintf(
		"%s - %d,%d - %s : %d - %d ", spaceshipPos, sW, sH, centrePos, s.XAxisForce, s.YAxisForce,
	)
	return spaceshipText
}

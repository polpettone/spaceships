package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type DebugScreen struct {
	Image  *ebiten.Image
	Width  int
	Height int
	Text   string
}

func NewDebugScreen(Width, Height int) (*DebugScreen, error) {
	image, err := ebiten.NewImage(
		Width,
		Height,
		ebiten.FilterDefault)

	image.Fill(color.RGBA{0, 0, 0, 0xff})

	if err != nil {
		return nil, err
	}

	return &DebugScreen{
		Image:  image,
		Width:  Width,
		Height: Height,
		Text:   "Debug Screen",
	}, nil
}

func (d *DebugScreen) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrintAt(
		d.Image,
		d.Text, 10, 10)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(1500), float64(0))
	screen.DrawImage(d.Image, op)
}

func (d *DebugScreen) Update(g *Game) {

	t :=
		`
	 TPS: %f2.2f
	 FPS: %f2.2f
	 Game Object Count: %d`

	d.Text = fmt.Sprintf(
		t,
		ebiten.CurrentTPS(),
		ebiten.CurrentFPS(),
		len(g.GameObjects))
}

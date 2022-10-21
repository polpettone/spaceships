package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
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

func (d *DebugScreen) SetText(t string) {
	d.Text = t
	text.Draw(d.Image, t, mplusNormalFont, 10, 10, color.White)
}

func (d *DebugScreen) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(1500), float64(0))
	screen.DrawImage(d.Image, op)
}

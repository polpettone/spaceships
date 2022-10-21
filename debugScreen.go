package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

type DebugScreen struct {
	BackgroundImage *ebiten.Image
	Width           int
	Height          int
	Text            string
}

func NewDebugScreen(Width, Height int) (*DebugScreen, error) {
	backgroundImage, err := ebiten.NewImage(
		Width,
		Height,
		ebiten.FilterDefault)

	backgroundImage.Fill(color.RGBA{0, 0, 0, 0xff})

	if err != nil {
		return nil, err
	}

	return &DebugScreen{
		BackgroundImage: backgroundImage,
		Width:           Width,
		Height:          Height,
		Text:            "Debug Screen",
	}, nil
}

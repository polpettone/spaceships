package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	spaceshipSize = 40
)

type Spaceship struct {
	Image *ebiten.Image
}

func NewSpaceship() (*Spaceship, error) {
	img, err := createSpaceshipImage()

	if err != nil {
		return nil, err
	}

	return &Spaceship{
		Image: img,
	}, nil

}

func createSpaceshipImage() (*ebiten.Image, error) {
	img, err := ebiten.NewImage(spaceshipSize, spaceshipSize, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	img.Fill(color.RGBA{0, 0, 0, 0xff})
	return img, nil
}

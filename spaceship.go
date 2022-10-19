package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	spaceShipSize = 10
)

type Spaceship struct {
	Image *ebiten.Image
}

type GameObject interface {
	GetPos() (int, int)
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
	img, err := ebiten.NewImage(spaceShipSize, spaceShipSize, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	img.Fill(color.RGBA{0, 0, 0, 0xff})
	return img, nil
}

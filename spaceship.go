package main

import (
	"image/color"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
)

const (
	spaceshipSize = 40
)

type Spaceship struct {
	Image *ebiten.Image
	Pos   Pos
	ID    string
}

func NewSpaceship(initialPos Pos) (*Spaceship, error) {
	img, err := createSpaceshipImage()

	if err != nil {
		return nil, err
	}

	return &Spaceship{
		Image: img,
		Pos:   initialPos,
		ID:    uuid.New().String(),
	}, nil
}

func (s *Spaceship) GetPos() Pos {
	return s.Pos
}

func (s *Spaceship) GetImage() *ebiten.Image {
	return s.Image
}

func (s *Spaceship) GetID() string {
	return s.ID
}

func (s *Spaceship) Update() {
	s.Pos.X += 1
}

func createSpaceshipImage() (*ebiten.Image, error) {
	img, err := ebiten.NewImage(spaceshipSize, spaceshipSize, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	img.Fill(color.RGBA{0, 0, 0, 0xff})
	return img, nil
}

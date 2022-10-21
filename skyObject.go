package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

const (
	skyObjectSize = 60
)

type SkyObject struct {
	Image    *ebiten.Image
	Pos      Pos
	Velocity int
}

func NewSkyObject(initialPos Pos) (*SkyObject, error) {
	img, err := createSkyObjectImage()

	if err != nil {
		return nil, err
	}

	return &SkyObject{
		Image:    img,
		Pos:      initialPos,
		Velocity: 2,
	}, nil
}

func (s *SkyObject) GetPos() Pos {
	return s.Pos
}

func (s *SkyObject) GetImage() *ebiten.Image {
	return s.Image
}

func createSkyObjectImage() (*ebiten.Image, error) {
	img, err := ebiten.NewImage(skyObjectSize, skyObjectSize, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	img.Fill(color.RGBA{192, 192, 192, 0xff})
	return img, nil
}

func (s *SkyObject) Update() {
	s.Pos.X += s.Velocity
}

func CreateSkyObjects() []GameObject {

	a, _ := NewSkyObject(NewPos(400, 100))
	b, _ := NewSkyObject(NewPos(700, 200))
	c, _ := NewSkyObject(NewPos(500, 200))

	skyObjects := []GameObject{
		a, b, c,
	}

	return skyObjects
}

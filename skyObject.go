package main

import (
	"image/color"
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
)

const (
	skyObjectSize = 30
)

type SkyObject struct {
	Image    *ebiten.Image
	Pos      Pos
	Velocity int
	ID       string
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
		ID:       uuid.New().String(),
	}, nil
}

func (s *SkyObject) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.Pos.X), float64(s.Pos.Y))
	screen.DrawImage(s.Image, op)
}

func (s *SkyObject) GetID() string {
	return s.ID
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
	s.Pos.X -= s.Velocity
}

func CreateSkyObjectAtRandomPosition(minX, minY, maxX, maxY, count int) []GameObject {

	skyObjects := []GameObject{}

	for n := 0; n < count; n++ {
		x := rand.Intn(maxX-minX) + minX
		y := rand.Intn(maxY-minY) + minY
		a, _ := NewSkyObject(NewPos(x, y))
		skyObjects = append(skyObjects, a)
	}

	return skyObjects
}

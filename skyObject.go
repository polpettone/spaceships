package main

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	skyObjectSize = 30
)

type SkyObject struct {
	Image    *ebiten.Image
	Pos      Pos
	Velocity int
	ID       string
	Size     int
}

func NewSkyObject(initialPos Pos) (*SkyObject, error) {
	img, err := createSkyObjectImageFromAsset()

	if err != nil {
		return nil, err
	}

	return &SkyObject{
		Image:    img,
		Pos:      initialPos,
		Velocity: 2,
		ID:       uuid.New().String(),
		Size:     skyObjectSize,
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

func (s *SkyObject) GetSize() int {
	return s.Size
}

func createSkyObjectImageFromAsset() (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(
		"assets/enemy-1.gif",
		ebiten.FilterDefault)

	if err != nil {
		return nil, err
	}
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

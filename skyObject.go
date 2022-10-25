package main

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type SkyObject struct {
	Image    *ebiten.Image
	Pos      Pos
	Velocity int
	ID       string
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
	}, nil
}

func (s *SkyObject) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
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

func (s *SkyObject) GetSize() (width, height int) {
	return s.Image.Size()
}

func (s *SkyObject) GetCentrePos() Pos {
	w, h := s.GetSize()
	x := (w / 2) + s.Pos.X
	y := (h / 2) + s.Pos.Y
	return NewPos(x, y)
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

func (s *SkyObject) Update(g *Game) {
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

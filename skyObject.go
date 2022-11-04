package main

import (
	"image/color"
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type SkyObject struct {
	CurrentImage   *ebiten.Image
	AliveImage     *ebiten.Image
	DestroyedImage *ebiten.Image
	Pos            Pos
	Velocity       int
	ID             string
	Alive          bool
}

func NewSkyObject(initialPos Pos) (*SkyObject, error) {

	img, err := createSkyObjectImageFromAsset()
	if err != nil {
		return nil, err
	}

	destroyedImg, err := createDestroyedImage(25)
	if err != nil {
		return nil, err
	}

	return &SkyObject{
		CurrentImage:   img,
		AliveImage:     img,
		DestroyedImage: destroyedImg,
		Pos:            initialPos,
		Velocity:       2,
		ID:             uuid.New().String(),
		Alive:          true,
	}, nil
}

func (s *SkyObject) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1, 1)
	op.GeoM.Translate(float64(s.Pos.X), float64(s.Pos.Y))
	screen.DrawImage(s.CurrentImage, op)
}

func (s *SkyObject) IsAlive() bool {
	return s.Alive
}

func (s *SkyObject) GetID() string {
	return s.ID
}

func (s *SkyObject) GetPos() Pos {
	return s.Pos
}

func (s *SkyObject) GetImage() *ebiten.Image {
	return s.CurrentImage
}

func (s *SkyObject) GetSize() (width, height int) {
	return s.CurrentImage.Size()
}

func (s *SkyObject) Destroy() {
	s.Alive = false
	s.CurrentImage = s.DestroyedImage
}

func (s *SkyObject) GetCentrePos() Pos {
	w, h := s.GetSize()
	x := (w / 2) + s.Pos.X
	y := (h / 2) + s.Pos.Y
	return NewPos(x, y)
}

func (s *SkyObject) GetType() string {
	return "skyObject"
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

func createDestroyedImage(size int) (*ebiten.Image, error) {
	img, err := ebiten.NewImage(size, size, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	img.Fill(color.RGBA{47, 79, 79, 0xff})
	return img, nil
}
func (s *SkyObject) Update(g *Game) {
	if s.Alive {
		s.Pos.X -= s.Velocity
	}
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

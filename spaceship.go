package main

import (
	"image/color"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
)

const (
	spaceshipSize = 10
)

type Spaceship struct {
	Image         *ebiten.Image
	Pos           Pos
	ID            string
	DownForce     int
	UpForce       int
	ForwardForce  int
	BackwardForce int
	DamageCount   int
	Size          int
}

func NewSpaceship(initialPos Pos) (*Spaceship, error) {
	img, err := createSpaceshipImageFromAsset()

	if err != nil {
		return nil, err
	}

	return &Spaceship{
		Image:         img,
		Pos:           initialPos,
		ID:            uuid.New().String(),
		DownForce:     0,
		UpForce:       0,
		ForwardForce:  0,
		BackwardForce: 0,
		DamageCount:   0,
		Size:          spaceshipSize,
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

func (s *Spaceship) Update(maxX, maxY int) {

	handleControls(s)

	if s.Pos.X < maxX-3 {
		s.Pos.X += s.ForwardForce
	}
	if s.Pos.X > 3 {
		s.Pos.X -= s.BackwardForce
	}
	if s.Pos.Y < maxY-3 {
		s.Pos.Y += s.UpForce
	}
	if s.Pos.Y > 3 {
		s.Pos.Y -= s.DownForce
	}
}

func (s *Spaceship) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.Pos.X), float64(s.Pos.Y))
	screen.DrawImage(s.Image, op)

}

func handleControls(s *Spaceship) {

	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		s.DownForce += 1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		s.UpForce += 1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		s.BackwardForce += 1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		s.ForwardForce += 1
	}

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		s.ForwardForce = 0
		s.BackwardForce = 0
		s.UpForce = 0
		s.DownForce = 0
	}

}

func createSpaceshipImage() (*ebiten.Image, error) {
	img, err := ebiten.NewImage(spaceshipSize, spaceshipSize, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	img.Fill(color.RGBA{0, 0, 0, 0xff})
	return img, nil
}

func createSpaceshipImageFromAsset() (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(
		"assets/spaceship-1.gif",
		ebiten.FilterDefault)

	if err != nil {
		return nil, err
	}
	return img, nil
}

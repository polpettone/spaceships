package main

import (
	"image/color"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
)

const (
	bulletSize = 2
)

type Bullet struct {
	Image    *ebiten.Image
	Pos      Pos
	Velocity int
	ID       string
}

func NewBullet(initialPos Pos) (*Bullet, error) {
	img, err := createBulletImage()

	if err != nil {
		return nil, err
	}

	return &Bullet{
		Image:    img,
		Pos:      initialPos,
		Velocity: 5,
		ID:       uuid.New().String(),
	}, nil
}

func (s *Bullet) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.Pos.X), float64(s.Pos.Y))
	screen.DrawImage(s.Image, op)
}

func (s *Bullet) GetID() string {
	return s.ID
}

func (s *Bullet) GetPos() Pos {
	return s.Pos
}

func (s *Bullet) GetImage() *ebiten.Image {
	return s.Image
}

func createBulletImage() (*ebiten.Image, error) {
	img, err := ebiten.NewImage(bulletSize, bulletSize, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	img.Fill(color.RGBA{255, 0, 0, 0xff})
	return img, nil
}

func (s *Bullet) Update() {
	s.Pos.X += s.Velocity
}
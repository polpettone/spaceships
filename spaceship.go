package main

import (
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/v2/audio"
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
	ShootBullet   bool

	ShootSound *audio.Player
}

func NewSpaceship(initialPos Pos) (*Spaceship, error) {
	img, err := createSpaceshipImageFromAsset()

	if err != nil {
		return nil, err
	}

	shootSound, err := InitSoundPlayer(
		"assets/gunshot.mp3",
		audioContext)

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
		ShootSound:    shootSound,
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

func (s *Spaceship) GetSize() (width, height int) {
	return s.Image.Size()
}

func (s *Spaceship) GetCentrePos() Pos {
	w, h := s.GetSize()
	x := (w / 2) + s.Pos.X
	y := (h / 2) + s.Pos.Y
	return NewPos(x, y)
}

//TODO: err handling
func (s *Spaceship) Update(g *Game) {

	handleControls(s)

	updatePosition(s, g)

	updateWeapons(s, g)

}

func (s *Spaceship) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.Pos.X), float64(s.Pos.Y))
	screen.DrawImage(s.Image, op)
}

//TODO: err handling
func updateWeapons(s *Spaceship, g *Game) {
	if s.ShootBullet {
		pos := NewPos(s.Pos.X, s.Pos.Y+20)
		bullet, _ := NewBullet(pos)
		s.ShootBullet = false
		g.GameObjects[bullet.ID] = bullet
		s.ShootSound.Rewind()
		s.ShootSound.Play()
	}
}

func updatePosition(s *Spaceship, g *Game) {
	if s.Pos.X < g.MaxX-3 {
		s.Pos.X += s.ForwardForce
	}
	if s.Pos.X > 3 {
		s.Pos.X -= s.BackwardForce
	}
	if s.Pos.Y < g.MaxY-3 {
		s.Pos.Y += s.UpForce
	}
	if s.Pos.Y > 3 {
		s.Pos.Y -= s.DownForce
	}
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

	if inpututil.IsKeyJustPressed(ebiten.KeyN) {
		s.ShootBullet = true
	}

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

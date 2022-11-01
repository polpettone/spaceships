package main

import (
	"math"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/polpettone/gaming/natalito/engine"
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

	BulletCount int
	Health      int

	ControlMap SpaceshipControlMap

	ImageScale float64
}

type SpaceshipControlMap struct {
	Up    ebiten.Key
	Down  ebiten.Key
	Left  ebiten.Key
	Right ebiten.Key
	Break ebiten.Key
	Shoot ebiten.Key
}

func NewSpaceship(initialPos Pos) (*Spaceship, error) {

	controlMap := SpaceshipControlMap{
		Up:    ebiten.KeyK,
		Down:  ebiten.KeyJ,
		Left:  ebiten.KeyH,
		Right: ebiten.KeyL,
		Break: ebiten.KeySpace,
		Shoot: ebiten.KeyN,
	}

	img, err := createSpaceshipImageFromAsset()

	if err != nil {
		return nil, err
	}

	shootSound, err := engine.InitSoundPlayer(
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
		Health:        1000,
		BulletCount:   30,
		ControlMap:    controlMap,
		ImageScale:    0.1,
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
	w, h := s.Image.Size()
	return int(s.ImageScale * float64(w)), int(s.ImageScale * float64(h))

}

func (s *Spaceship) GetCentrePos() Pos {
	w, h := s.GetSize()
	x := (w / 2) + s.Pos.X
	y := (h / 2) + s.Pos.Y
	return NewPos(x, y)
}

func (s *Spaceship) GetType() string {
	return "spaceship"
}

//TODO: err handling
func (s *Spaceship) Update(g *Game) {

	handleControls(s)

	updatePosition(s, g)

	updateWeapons(s, g)

}

func (s *Spaceship) Draw(screen *ebiten.Image) {

	w, h := s.Image.Size()
	op := &ebiten.DrawImageOptions{}

	// Move the image's center to the screen's upper-left corner.
	// This is a preparation for rotating. When geometry matrices are applied,
	// the origin point is the upper-left corner.
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// Rotate the image. As a result, the anchor point of this rotate is
	// the center of the image.

	op.GeoM.Rotate(float64(90%360) * 2 * math.Pi / 360)

	op.GeoM.Scale(s.ImageScale, s.ImageScale)
	op.GeoM.Translate(float64(s.Pos.X), float64(s.Pos.Y))
	screen.DrawImage(s.Image, op)
}

//TODO: err handling
func updateWeapons(s *Spaceship, g *Game) {
	if s.ShootBullet && s.BulletCount > 0 {
		pos := NewPos(s.Pos.X, s.Pos.Y+20)
		bullet, _ := NewBullet(pos)
		s.ShootBullet = false
		g.GameObjects[bullet.ID] = bullet
		s.ShootSound.Rewind()
		s.ShootSound.Play()
		s.BulletCount = s.BulletCount - 1
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
		s.Pos.Y -= s.UpForce
	}
	if s.Pos.Y > 3 {
		s.Pos.Y += s.DownForce
	}
}

func handleControls(s *Spaceship) {

	if inpututil.IsKeyJustPressed(s.ControlMap.Up) {
		s.UpForce += 1
	}

	if inpututil.IsKeyJustPressed(s.ControlMap.Down) {
		s.DownForce += 1
	}

	if inpututil.IsKeyJustPressed(s.ControlMap.Left) {
		s.BackwardForce += 1
	}

	if inpututil.IsKeyJustPressed(s.ControlMap.Right) {
		s.ForwardForce += 1
	}

	if inpututil.IsKeyJustPressed(s.ControlMap.Break) {
		s.ForwardForce = 0
		s.BackwardForce = 0
		s.UpForce = 0
		s.DownForce = 0
	}

	if inpututil.IsKeyJustPressed(s.ControlMap.Shoot) {
		s.ShootBullet = true
	}

}

func createSpaceshipImageFromAsset() (*ebiten.Image, error) {
	img, _, err := ebitenutil.NewImageFromFile(
		"assets/ship2.png",
		ebiten.FilterDefault)

	if err != nil {
		return nil, err
	}
	return img, nil
}

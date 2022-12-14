package models

import (
	"fmt"
	"image/color"
	"math"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/polpettone/gaming/spaceships/engine"
)

const (
	spaceshipSize          = 10
	spaceshipWallTolerance = 50
)

type Spaceship struct {
	PilotName string

	CurrentImage *GameObjectImage
	Image        *GameObjectImage
	DamageImage  *GameObjectImage

	ShootSound   *audio.Player
	ImpulseSound *audio.Player
	ImpactSound  *audio.Player

	ID        string
	Signature string

	Pos           Pos
	MoveDirection int
	XAxisForce    int
	YAxisForce    int
	Acceleration  bool
	ShootBullet   bool

	BulletCount int
	Health      int

	Size int

	Control SpaceshipControl

	SoundOn bool

	SimpleImpetusUnit SimpleImpetusUnit
}

func NewSpaceship(
	audioContext *audio.Context,
	health int,
	bulletCount int,
	pilotName string,
	initialPos Pos,
	control SpaceshipControl,
	img *GameObjectImage,
	damageImg *GameObjectImage,
	signature string) (*Spaceship, error) {

	shootSound, err := engine.InitSoundPlayer(
		"assets/sounds/Laser_shoot 39.wav",
		engine.TypeWAV,
		audioContext)

	if err != nil {
		return nil, err
	}

	impulseSound, err := engine.InitSoundPlayer(
		"assets/sounds/impulse.wav",
		engine.TypeWAV,
		audioContext)

	if err != nil {
		return nil, err
	}

	impactSound, err := engine.InitSoundPlayer(
		"assets/sounds/big-impact-7054.mp3",
		engine.TypeMP3,
		audioContext)

	if err != nil {
		return nil, err
	}

	return &Spaceship{
		PilotName:         pilotName,
		CurrentImage:      img,
		Image:             img,
		DamageImage:       damageImg,
		Pos:               initialPos,
		ID:                uuid.New().String(),
		Size:              spaceshipSize,
		ShootSound:        shootSound,
		ImpulseSound:      impulseSound,
		ImpactSound:       impactSound,
		Health:            health,
		BulletCount:       bulletCount,
		Control:           control,
		MoveDirection:     1,
		SoundOn:           false,
		Signature:         signature,
		SimpleImpetusUnit: SimpleImpetusUnit{},
	}, nil
}

func (s *Spaceship) Reset(health int, pos Pos, bulletCount int, moveDirection int) {
	s.Health = health
	s.Pos = pos
	s.MoveDirection = moveDirection
	s.BulletCount = bulletCount
}

func (s *Spaceship) GetSignature() string {
	return s.Signature
}

func (s *Spaceship) GetPos() Pos {
	return s.Pos
}

func (s *Spaceship) GetID() string {
	return s.ID
}

func (s *Spaceship) GetSize() (width, height int) {
	w, h := s.Image.Image.Size()
	return int(s.Image.Scale * float64(w)), int(s.Image.Scale * float64(h))

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

func (s *Spaceship) Alive() bool {
	return s.Health > 0
}

func updateSpaceshipPosition(
	s *Spaceship,
	g Scene,
	pos Pos,
	xAxisForce,
	yAxisForce int) {

	if pos.X < g.GetMaxX()-spaceshipWallTolerance {
		s.Pos.X = pos.X
	}

	if pos.X > spaceshipWallTolerance {
		s.Pos.X = pos.X
	}

	if pos.Y < g.GetMaxY()-spaceshipWallTolerance {
		s.Pos.Y = pos.Y
	}

	if pos.Y > spaceshipWallTolerance {
		s.Pos.Y = pos.Y
	}

	if pos.X-spaceshipWallTolerance < 0 ||
		pos.X+spaceshipWallTolerance > g.GetMaxX() {
		xAxisForce = 0
	}

	if pos.Y-spaceshipWallTolerance < 0 ||
		pos.Y+spaceshipWallTolerance > g.GetMaxY() {
		yAxisForce = 0
	}

	s.XAxisForce = xAxisForce
	s.YAxisForce = yAxisForce

}

func (s *Spaceship) Update(g Scene) {

	if s == nil {
		return
	}

	handleSpaceshipControl(s)

	if s.Acceleration {

		pos,
			xAxisForce,
			yAxisForce := s.SimpleImpetusUnit.Accelerate(
			s.Pos,
			s.XAxisForce,
			s.YAxisForce)

		updateSpaceshipPosition(s, g, pos, xAxisForce, yAxisForce)

	}

	updateWeapons(s, g)

	if g.GetTickCounter()%100 == 0 {
		s.CurrentImage = s.Image
	}

}

func (s *Spaceship) Damage() {
	s.Health -= 1

	s.CurrentImage = s.DamageImage

	if s.SoundOn {
		if !s.ImpactSound.IsPlaying() {
			s.ImpactSound.Rewind()
			s.ImpactSound.Play()
		}
	}
}

func (s *Spaceship) Draw(screen *ebiten.Image) {

	w, h := s.CurrentImage.Image.Size()
	op := &ebiten.DrawImageOptions{}

	// Move the image's center to the screen's upper-left corner.
	// This is a preparation for rotating. When geometry matrices are applied,
	// the origin point is the upper-left corner.
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// Rotate the image. As a result, the anchor point of this rotate is
	// the center of the image.

	op.GeoM.Rotate(float64(360%360) * 2 * math.Pi / 360)

	op.GeoM.Scale(float64(s.MoveDirection)*s.CurrentImage.Scale*float64(s.CurrentImage.Direction),
		s.CurrentImage.Scale)

	op.GeoM.Translate(float64(s.Pos.X), float64(s.Pos.Y))
	screen.DrawImage(s.CurrentImage.Image, op)
}

func (s *Spaceship) DrawState(screen *ebiten.Image, x int, y int) {
	t := fmt.Sprintf(
		`
%s
Health: %d
Bullets %d`,
		s.PilotName,
		s.Health,
		s.BulletCount,
	)
	text.Draw(screen, t, engine.MplusNormalFont, x, y, color.White)
}

// TODO: err handling
func updateWeapons(s *Spaceship, g Scene) {
	if s.ShootBullet && s.BulletCount > 0 {
		pos := NewPos(s.Pos.X, s.Pos.Y+20)
		bullet, _ := NewBullet(pos, s.MoveDirection, s.Signature, g.GetConfig().BulletVelocity)
		s.ShootBullet = false
		g.AddGameObject(bullet)

		if s.SoundOn {
			s.ShootSound.Rewind()
			s.ShootSound.Play()
		}

		s.BulletCount = s.BulletCount - 1
	}
}

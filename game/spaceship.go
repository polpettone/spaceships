package game

import (
	"fmt"
	"image/color"
	"math"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/polpettone/gaming/natalito/engine"
)

const (
	spaceshipSize = 10
)

type Spaceship struct {
	Image       *GameObjectImage
	Pos         Pos
	ID          string
	DamageCount int
	Size        int
	ShootBullet bool

	XAxisForce int
	YAxisForce int

	ShootSound   *audio.Player
	ImpulseSound *audio.Player
	ImpactSound  *audio.Player

	BulletCount   int
	Health        int
	KilledEnemies int

	KeyboardControlMap *SpaceshipKeyboardControlMap
	GamepadControlMap  *SpaceshipGamepadControlMap

	MoveDirection int

	SoundOn bool
}

type SpaceshipKeyboardControlMap struct {
	Up    ebiten.Key
	Down  ebiten.Key
	Left  ebiten.Key
	Right ebiten.Key
	Break ebiten.Key
	Shoot ebiten.Key
}

type SpaceshipGamepadControlMap struct {
	Up    ebiten.GamepadButton
	Down  ebiten.GamepadButton
	Left  ebiten.GamepadButton
	Right ebiten.GamepadButton
	Break ebiten.GamepadButton
	Shoot ebiten.GamepadButton
}

func NewSpaceship(initialPos Pos,
	keyboardControlMap *SpaceshipKeyboardControlMap,
	gamepadControlMap *SpaceshipGamepadControlMap) (*Spaceship, error) {

	img, err := createSpaceshipImageFromAsset()

	if err != nil {
		return nil, err
	}

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
		Image:              img,
		Pos:                initialPos,
		ID:                 uuid.New().String(),
		DamageCount:        0,
		Size:               spaceshipSize,
		ShootSound:         shootSound,
		ImpulseSound:       impulseSound,
		ImpactSound:        impactSound,
		Health:             20,
		BulletCount:        30,
		KeyboardControlMap: keyboardControlMap,
		GamepadControlMap:  gamepadControlMap,
		MoveDirection:      1,
		SoundOn:            false,
	}, nil
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

//TODO: err handling
func (s *Spaceship) Update(g Game) {

	handleGamepadControls(s)

	handleKeyboardControls(s)

	updatePosition(s, g)

	updateWeapons(s, g)

}

func (s *Spaceship) Damage() {
	s.DamageCount += 1
	s.Health -= 1

	if s.SoundOn {
		if !s.ImpactSound.IsPlaying() {
			s.ImpactSound.Rewind()
			s.ImpactSound.Play()
		}
	}
}

func (s *Spaceship) Draw(screen *ebiten.Image) {

	w, h := s.Image.Image.Size()
	op := &ebiten.DrawImageOptions{}

	// Move the image's center to the screen's upper-left corner.
	// This is a preparation for rotating. When geometry matrices are applied,
	// the origin point is the upper-left corner.
	op.GeoM.Translate(-float64(w)/2, -float64(h)/2)

	// Rotate the image. As a result, the anchor point of this rotate is
	// the center of the image.

	op.GeoM.Rotate(float64(360%360) * 2 * math.Pi / 360)

	op.GeoM.Scale(float64(s.MoveDirection)*s.Image.Scale*float64(s.Image.Direction),
		s.Image.Scale)

	op.GeoM.Translate(float64(s.Pos.X), float64(s.Pos.Y))
	screen.DrawImage(s.Image.Image, op)
}

func (s *Spaceship) DrawState(screen *ebiten.Image, x int, y int) {
	t := fmt.Sprintf(
		"\n Health: %d \n Bullets %d",
		s.Health,
		s.BulletCount,
	)
	text.Draw(screen, t, engine.MplusNormalFont, x, y, color.White)
}

//TODO: err handling
func updateWeapons(s *Spaceship, g Game) {
	if s.ShootBullet && s.BulletCount > 0 {
		pos := NewPos(s.Pos.X, s.Pos.Y+20)
		bullet, _ := NewBullet(pos, s.MoveDirection)
		s.ShootBullet = false
		g.AddGameObject(bullet)

		if s.SoundOn {
			s.ShootSound.Rewind()
			s.ShootSound.Play()
		}

		s.BulletCount = s.BulletCount - 1
	}
}

func updatePosition(s *Spaceship, g Game) {
	if s.Pos.X < g.GetMaxX()-spaceshipWallTolerance && s.XAxisForce > 0 {
		s.Pos.X += s.XAxisForce
	}
	if s.Pos.X > spaceshipWallTolerance && s.XAxisForce < 0 {
		s.Pos.X += s.XAxisForce
	}
	if s.Pos.Y < g.GetMaxY()-spaceshipWallTolerance && s.YAxisForce > 0 {
		s.Pos.Y += s.YAxisForce
	}
	if s.Pos.Y > spaceshipWallTolerance && s.YAxisForce < 0 {
		s.Pos.Y += s.YAxisForce
	}

	if s.XAxisForce != 0 {
		if s.SoundOn {
			if !s.ImpulseSound.IsPlaying() {
				s.ImpulseSound.Rewind()
				s.ImpulseSound.Play()
			}
		}
	}
	if s.XAxisForce == 0 {
		if s.SoundOn {
			if !s.ImpulseSound.IsPlaying() {
				s.ImpulseSound.Pause()
			}
		}
	}

}

func handleGamepadControls(s *Spaceship) {

	if s.GamepadControlMap == nil {
		return
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Up) {
		s.YAxisForce -= 1
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Down) {
		s.YAxisForce += 1
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Left) {
		s.XAxisForce -= 1
		s.MoveDirection = -1
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Right) {
		s.XAxisForce += 1
		s.MoveDirection = 1
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Break) {
		s.YAxisForce = 0
		s.XAxisForce = 0
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Shoot) {
		s.ShootBullet = true
	}

}

func handleKeyboardControls(s *Spaceship) {

	if s.KeyboardControlMap == nil {
		return
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Up) {
		s.YAxisForce -= 1
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Down) {
		s.YAxisForce += 1
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Left) {
		s.XAxisForce -= 1
		s.MoveDirection = -1
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Right) {
		s.XAxisForce += 1
		s.MoveDirection = 1
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Break) {
		s.YAxisForce = 0
		s.XAxisForce = 0
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Shoot) {
		s.ShootBullet = true
	}

}
func createSpaceshipImageFromAsset() (*GameObjectImage, error) {
	img, _, err := ebitenutil.NewImageFromFile(
		"assets/images/spaceships/star-wars-2.png",
		ebiten.FilterDefault)

	if err != nil {
		return nil, err
	}

	gameObjectImage := &GameObjectImage{
		Image:     img,
		Scale:     0.2,
		Direction: -1,
	}

	return gameObjectImage, nil
}

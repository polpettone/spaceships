package models

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
)

type SkyObject struct {
	CurrentImage   *GameObjectImage
	AliveImage     *GameObjectImage
	DestroyedImage *GameObjectImage
	Pos            Pos
	Velocity       int
	ID             string
	Alive          bool
	ImageScale     float64
	MoveDirection  int
}

func NewSkyObject(initialPos Pos, moveDirection int) (*SkyObject, error) {

	img, err := NewGameObjectImage(
		"assets/images/spaceships/star-wars-1.png",
		0.1,
		-1)

	if err != nil {
		return nil, err
	}

	destroyedImg, err := NewGameObjectImage(
		"assets/images/figures/pilot1.png",
		0.03,
		-1)
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
		ImageScale:     0.1,
		MoveDirection:  moveDirection,
	}, nil
}

func (s *SkyObject) GetSignature() string {
	return ""
}

func (s *SkyObject) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(s.CurrentImage.Direction)*s.CurrentImage.Scale, s.CurrentImage.Scale)
	op.GeoM.Translate(float64(s.Pos.X), float64(s.Pos.Y))
	screen.DrawImage(s.CurrentImage.Image, op)
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
	return s.CurrentImage.Image
}

func (s *SkyObject) GetSize() (width, height int) {
	w, h := s.CurrentImage.Image.Size()
	return int(s.CurrentImage.Scale * float64(w)), int(s.CurrentImage.Scale * float64(h))
}

func (s *SkyObject) Destroy() {
	s.Alive = false
	s.CurrentImage = s.DestroyedImage
	s.Velocity = 1
}

func (s *SkyObject) GetCentrePos() Pos {
	w, h := s.GetSize()
	x := (w / 2) + s.Pos.X
	y := (h / 2) + s.Pos.Y
	return NewPos(x, y)
}

func (s *SkyObject) GetType() GameObjectType {
	return Enemy
}

func (s *SkyObject) Update() {
	s.Pos.X += s.MoveDirection * s.Velocity
}

func CreateSkyObjectAtRandomPosition(minX, minY, maxX, maxY, count int) []GameObject {

	skyObjects := []GameObject{}

	for n := 0; n < count; n++ {
		x := rand.Intn(maxX-minX) + minX
		y := rand.Intn(maxY-minY) + minY

		random := rand.Intn(3)

		var moveDirection int
		switch random {
		case 0:
			moveDirection = 0
		case 1:
			moveDirection = 1
		case 2:
			moveDirection = -1
		}

		a, _ := NewSkyObject(NewPos(x, y), moveDirection)
		skyObjects = append(skyObjects, a)
	}

	return skyObjects
}

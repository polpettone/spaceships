package game

import (
	"image/color"
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
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
		ImageScale:     0.1,
	}, nil
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

func createSkyObjectImageFromAsset() (*GameObjectImage, error) {
	img, _, err := ebitenutil.NewImageFromFile(
		"assets/images/spaceships/star-wars-1.png",
		ebiten.FilterDefault)

	if err != nil {
		return nil, err
	}

	gameObjectImage := &GameObjectImage{
		Image:     img,
		Scale:     0.1,
		Direction: -1,
	}

	return gameObjectImage, nil
}

func createDestroyedImage(size int) (*GameObjectImage, error) {
	img, err := ebiten.NewImage(size, size, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	img.Fill(color.RGBA{47, 79, 79, 0xff})

	gameObjectImage := &GameObjectImage{
		Image:     img,
		Scale:     1,
		Direction: 1,
	}

	return gameObjectImage, nil
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

package game

import (
	"image/color"
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
)

type Star struct {
	Image         *ebiten.Image
	Pos           Pos
	Velocity      int
	ID            string
	Alive         bool
	MoveDirection int
}

func NewStar(initialPos Pos, moveDirection int) (*Star, error) {
	img, err := createStarImage(2)

	if err != nil {
		return nil, err
	}

	return &Star{
		Image:         img,
		Pos:           initialPos,
		Velocity:      10,
		ID:            uuid.New().String(),
		Alive:         true,
		MoveDirection: moveDirection,
	}, nil
}

func (s *Star) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.Pos.X), float64(s.Pos.Y))
	screen.DrawImage(s.Image, op)
}

func (s *Star) GetID() string {
	return s.ID
}

func (s *Star) GetPos() Pos {
	return s.Pos
}

func (s *Star) GetImage() *ebiten.Image {
	return s.Image
}

func (s *Star) GetSize() (width, height int) {
	return s.Image.Size()
}

func (s *Star) GetType() GameObjectType {
	return Passive
}

func (s *Star) Destroy() {
	s.Alive = false
}

func (s *Star) IsAlive() bool {
	return s.Alive
}

func (s *Star) GetCentrePos() Pos {
	w, h := s.GetSize()
	x := (w / 2) + s.Pos.X
	y := (h / 2) + s.Pos.Y
	return NewPos(x, y)
}

func createStarImage(size int) (*ebiten.Image, error) {
	img, err := ebiten.NewImage(size, size, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	img.Fill(color.RGBA{255, 255, 255, 0xff})
	return img, nil
}

func (s *Star) Update() {
	s.Pos.X += s.MoveDirection * s.Velocity
}

func CreateStarAtRandomPosition(minX, minY, maxX, maxY, count int) []GameObject {

	ammos := []GameObject{}

	for n := 0; n < count; n++ {
		x := rand.Intn(maxX-minX) + minX
		y := rand.Intn(maxY-minY) + minY
		a, _ := NewStar(NewPos(x, y), -1)
		ammos = append(ammos, a)
	}

	return ammos
}

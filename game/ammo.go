package game

import (
	"image/color"
	"math/rand"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
)

type Ammo struct {
	Image         *ebiten.Image
	Pos           Pos
	Velocity      int
	ID            string
	Alive         bool
	MoveDirection int
}

func NewAmmo(initialPos Pos) (*Ammo, error) {
	img, err := createAmmoImage(10)

	if err != nil {
		return nil, err
	}

	return &Ammo{
		Image:         img,
		Pos:           initialPos,
		Velocity:      1,
		ID:            uuid.New().String(),
		Alive:         true,
		MoveDirection: -1,
	}, nil
}

func (a *Ammo) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(a.Pos.X), float64(a.Pos.Y))
	screen.DrawImage(a.Image, op)
}

func (a *Ammo) GetID() string {
	return a.ID
}

func (a *Ammo) GetPos() Pos {
	return a.Pos
}

func (a *Ammo) GetImage() *ebiten.Image {
	return a.Image
}

func (a *Ammo) GetSize() (width, height int) {
	return a.Image.Size()
}

func (a *Ammo) GetType() GameObjectType {
	return Item
}

func (a *Ammo) Destroy() {
	a.Alive = false
}

func (a *Ammo) IsAlive() bool {
	return a.Alive
}

func (a *Ammo) GetCentrePos() Pos {
	w, h := a.GetSize()
	x := (w / 2) + a.Pos.X
	y := (h / 2) + a.Pos.Y
	return NewPos(x, y)
}

func createAmmoImage(size int) (*ebiten.Image, error) {
	img, err := ebiten.NewImage(size, size, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}
	img.Fill(color.RGBA{24, 0, 255, 0xff})
	return img, nil
}

func (a *Ammo) Update() {
	a.Pos.X += a.MoveDirection * a.Velocity
}

func CreateAmmoAtRandomPosition(minX, minY, maxX, maxY, count int) []GameObject {

	ammos := []GameObject{}

	for n := 0; n < count; n++ {
		x := rand.Intn(maxX-minX) + minX
		y := rand.Intn(maxY-minY) + minY
		a, _ := NewAmmo(NewPos(x, y))
		ammos = append(ammos, a)
	}

	return ammos
}

package models

import "github.com/hajimehoshi/ebiten"

type GameObjectType int64

const (
	Weapon GameObjectType = iota
	Enemy
	Item
	Passive
)

type GameObject interface {
	GetID() string

	GetPos() Pos
	GetCentrePos() Pos

	GetImage() *ebiten.Image

	GetType() GameObjectType

	GetSize() (width, height int)

	Destroy()

	Draw(screen *ebiten.Image)
	Update()

	IsAlive() bool

	GetSignature() string
}

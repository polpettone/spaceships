package main

import "github.com/hajimehoshi/ebiten"

type GameObjectType int64

const (
	Weapon GameObjectType = iota
	Enemy
	Item
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
	Update(g *Game)

	IsAlive() bool
}

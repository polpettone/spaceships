package main

import "github.com/hajimehoshi/ebiten"

type GameObject interface {
	GetID() string

	GetPos() Pos
	GetCentrePos() Pos

	GetImage() *ebiten.Image

	//TODO: think about
	GetType() string

	GetSize() (width, height int)

	Destroy()

	Draw(screen *ebiten.Image)
	Update(g *Game)

	IsAlive() bool
}

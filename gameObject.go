package main

import "github.com/hajimehoshi/ebiten"

type GameObject interface {
	GetPos() Pos
	GetImage() *ebiten.Image

	GetID() string

	//TODO: think about
	GetType() string

	GetSize() (width, height int)
	GetCentrePos() Pos

	Draw(screen *ebiten.Image)
	Destroy()

	Update(g *Game)
}

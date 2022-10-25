package main

import "github.com/hajimehoshi/ebiten"

type GameObject interface {
	GetPos() Pos
	GetImage() *ebiten.Image
	Update(g *Game)
	GetID() string
	Draw(screen *ebiten.Image)
	GetSize() (width, height int)
	GetCentrePos() Pos
}

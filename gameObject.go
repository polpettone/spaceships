package main

import "github.com/hajimehoshi/ebiten"

type GameObject interface {
	GetPos() Pos
	GetImage() *ebiten.Image
	Update()
	GetID() string
	Draw(screen *ebiten.Image)
}

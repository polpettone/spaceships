package main

import "github.com/hajimehoshi/ebiten"

type GameObject interface {
	GetPos() Pos
	GetImage() *ebiten.Image
}

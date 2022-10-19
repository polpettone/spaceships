package main

import "github.com/hajimehoshi/ebiten"

type GameObject interface {
	GetPos() (int, int)
	GetImage() *ebiten.Image
}

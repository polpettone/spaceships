package main

import "github.com/hajimehoshi/ebiten"

type Spaceship struct {
	Image *ebiten.Image
}

type GameObject interface {
	GetPos() (int, int)
}

package models

import "github.com/hajimehoshi/ebiten"

type GameObjectImage struct {
	Image     *ebiten.Image
	Scale     float64
	Direction int
}

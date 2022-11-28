package models

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type GameObjectImage struct {
	Image     *ebiten.Image
	Scale     float64
	Direction int
}

func NewGameObjectImage(filePath string, scale float64, direction int) (*GameObjectImage, error) {
	img, _, err := ebitenutil.NewImageFromFile(
		filePath,
		ebiten.FilterDefault)

	if err != nil {
		return nil, err
	}

	gameObjectImage := &GameObjectImage{
		Image:     img,
		Scale:     scale,
		Direction: direction,
	}

	return gameObjectImage, nil
}

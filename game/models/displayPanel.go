package models

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/polpettone/gaming/spaceships/engine"
)

type SpaceshipDisplay struct {
	Spaceship Spaceship
	Width     int
	Height    int
	Image     *ebiten.Image
}

func NewSpaceshipDisplay(s Spaceship) (*SpaceshipDisplay, error) {

	img, _, err := ebitenutil.NewImageFromFile(
		"assets/images/spaceshipDisplayBackground.png",
		ebiten.FilterDefault)

	if err != nil {
		return nil, err
	}

	return &SpaceshipDisplay{
		Spaceship: s,
		Image:     img,
	}, nil

}

func (d *SpaceshipDisplay) Draw() {
	t := fmt.Sprintf("Pilot: %s", d.Spaceship.PilotName)
	text.Draw(d.Image, t, engine.MplusBigFont, 0, 0, color.White)
}

package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	natalitoImage *ebiten.Image
)

const (
	screenWidth  = 1000
	screenHeight = 500
)

type Game struct {
	BackgroundImage *ebiten.Image
}

func init() {
	var err error
	natalitoImage, _, err = ebitenutil.NewImageFromFile("assets/natalito-front.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(natalitoImage, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Natalito")

	backgroundImage, err := ebiten.NewImage(screenWidth, screenHeight, ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
		return
	}
	backgroundImage.Fill(color.RGBA{240, 255, 240, 0xff})

	g := &Game{
		BackgroundImage: backgroundImage,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	natalitoImage *ebiten.Image
)

const (
	screenWidth  = 2000
	screenHeight = 700
)

type Game struct {
	BackgroundImage *ebiten.Image
	Spaceship       *Spaceship
	GameObjects     []GameObject
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

	for _, o := range g.GameObjects {
		x, y := o.GetPos()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(x), float64(y))
		screen.DrawImage(o.GetImage(), op)
	}

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
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

	spaceship, err := NewSpaceship()
	if err != nil {
		log.Fatal(err)
		return
	}

	gameObjects := []GameObject{
		spaceship,
	}

	g := &Game{
		BackgroundImage: backgroundImage,
		GameObjects:     gameObjects,
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	// Erstelle ein neues Bild mit einer Größe von 200x200 Pixeln
	img := image.NewRGBA(image.Rect(0, 0, 200, 200))

	// Setze den Hintergrund des Bildes auf schwarz
	draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

	// Zeichne das Raumschiff auf das Bild
	drawSpaceShip(img)

	// Erstelle eine neue Datei
	f, _ := os.Create("spaceship1.png")
	defer f.Close()

	// Schreibe das Bild als PNG in die Datei
	png.Encode(f, img)
}

func drawSpaceShip(img *image.RGBA) {
	// Setze die Farbe auf Grau
	c := color.RGBA{200, 200, 200, 255}

	// Zeichne den Rumpf des Raumschiffs als Oval
	draw.Draw(img, image.Rect(50, 100, 150, 150), &image.Uniform{c}, image.ZP, draw.Src)

	// Zeichne die beiden Flügel des Raumschiffs als Halbkreise
	draw.Draw(img, image.Rect(25, 50, 75, 100), &image.Uniform{c}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(125, 50, 175, 100), &image.Uniform{c}, image.ZP, draw.Src)

	// Zeichne den Kommandoturm des Raumschiffs als Quadrat
	draw.Draw(img, image.Rect(100, 0, 200, 100), &image.Uniform{c}, image.ZP, draw.Src)
}

func drawSpaceShip1(img *image.RGBA) {
	// Setze die Farbe auf Weiß
	c := color.RGBA{255, 255, 255, 255}

	// Zeichne den Rumpf des Raumschiffs als Quadrat
	draw.Draw(img, image.Rect(50, 150, 150, 50), &image.Uniform{c}, image.ZP, draw.Src)

	// Zeichne die beiden Flügel des Raumschiffs als Dreiecke
	draw.Draw(img, image.Rect(25, 100, 75, 50), &image.Uniform{c}, image.ZP, draw.Src)
	draw.Draw(img, image.Rect(125, 100, 175, 50), &image.Uniform{c}, image.ZP, draw.Src)

	// Zeichne den Kommandoturm des Raumschiffs als Kreis
	draw.Draw(img, image.Rect(100, 25, 200, 125), &image.Uniform{c}, image.ZP, draw.Src)
}

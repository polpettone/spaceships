package models

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/v2/audio"
)

type Game interface {
	GetMaxX() int
	GetMaxY() int

	Layout(outsideWidth, outsideHeight int) (int, int)
	Update(screen *ebiten.Image) error

	GetConfig() GameConfig
}

var (
	audioContext *audio.Context

	keyboardControlMap   *KeyboardControl
	gamepadControlMap    *GamepadControl
	ps3GamepadControlMap *GamepadControl
)

func init() {
	audioContext = audio.NewContext(44100)

	keyboardControlMap = &KeyboardControl{
		Up:           ebiten.KeyK,
		Down:         ebiten.KeyJ,
		Left:         ebiten.KeyH,
		Right:        ebiten.KeyL,
		Break:        ebiten.KeyB,
		Shoot:        ebiten.KeyN,
		Acceleration: ebiten.KeySpace,
	}

	gamepadControlMap = &GamepadControl{
		Up:           ebiten.GamepadButton11,
		Down:         ebiten.GamepadButton13,
		Left:         ebiten.GamepadButton14,
		Right:        ebiten.GamepadButton12,
		Break:        ebiten.GamepadButton4,
		Shoot:        ebiten.GamepadButton0,
		Acceleration: ebiten.GamepadButton5,
	}

	ps3GamepadControlMap = &GamepadControl{
		Up:           ebiten.GamepadButton13,
		Down:         ebiten.GamepadButton14,
		Left:         ebiten.GamepadButton15,
		Right:        ebiten.GamepadButton16,
		Break:        ebiten.GamepadButton6,
		Shoot:        ebiten.GamepadButton0,
		Acceleration: ebiten.GamepadButton5,
	}

}

const (
	screenWidth  = 2000
	screenHeight = 1000
)

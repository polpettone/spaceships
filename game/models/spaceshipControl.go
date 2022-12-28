package models

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

type SpaceshipControl interface {
	IsUp() bool
	IsAccelerationPressed() bool
	IsAccelerationReleased() bool
	IsDown() bool
	IsLeft() bool
	IsRight() bool
	IsBreak() bool
	IsShoot() bool
}

type KeyboardControl struct {
	Up           ebiten.Key
	Down         ebiten.Key
	Left         ebiten.Key
	Right        ebiten.Key
	Break        ebiten.Key
	Shoot        ebiten.Key
	Acceleration ebiten.Key
}

type GamepadControl struct {
	Up           ebiten.GamepadButton
	Down         ebiten.GamepadButton
	Left         ebiten.GamepadButton
	Right        ebiten.GamepadButton
	Break        ebiten.GamepadButton
	Shoot        ebiten.GamepadButton
	Acceleration ebiten.GamepadButton
}

func (s *GamepadControl) IsShoot() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Shoot)
}

func (s *GamepadControl) IsUp() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Up)
}

func (s *GamepadControl) IsAccelerationPressed() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Acceleration)
}

func (s *GamepadControl) IsAccelerationReleased() bool {
	return inpututil.IsGamepadButtonJustReleased(0, s.Acceleration)
}

func (s *GamepadControl) IsDown() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Down)
}

func (s *GamepadControl) IsLeft() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Left)
}

func (s *GamepadControl) IsRight() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Right)
}

func (s *GamepadControl) IsBreak() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Break)
}

func (s *KeyboardControl) IsShoot() bool {
	return inpututil.IsKeyJustPressed(s.Shoot)
}

func (s *KeyboardControl) IsUp() bool {
	return inpututil.IsKeyJustPressed(s.Up)
}

func (s *KeyboardControl) IsAccelerationPressed() bool {
	return inpututil.IsKeyJustPressed(s.Acceleration)
}

func (s *KeyboardControl) IsAccelerationReleased() bool {
	return inpututil.IsKeyJustReleased(s.Acceleration)
}

func (s *KeyboardControl) IsDown() bool {
	return inpututil.IsKeyJustPressed(s.Down)
}

func (s *KeyboardControl) IsLeft() bool {
	return inpututil.IsKeyJustPressed(s.Left)
}

func (s *KeyboardControl) IsRight() bool {
	return inpututil.IsKeyJustPressed(s.Right)
}

func (s *KeyboardControl) IsBreak() bool {
	return inpututil.IsKeyJustPressed(s.Break)
}

func handleSpaceshipControl(s *Spaceship) {

	if s.Control.IsShoot() {
		s.ShootBullet = true
	}

	if s.Control.IsUp() {
		s.YAxisForce -= 1
	}

	if s.Control.IsLeft() {
		s.XAxisForce -= 1
		s.MoveDirection = -1
	}

	if s.Control.IsRight() {
		s.XAxisForce += 1
		s.MoveDirection = 1
	}

	if s.Control.IsDown() {
		s.YAxisForce += 1
	}

	if s.Control.IsBreak() {
		s.YAxisForce = 0
		s.XAxisForce = 0
	}

	if s.Control.IsAccelerationPressed() {
		s.Acceleration = true
		s.XAxisForce += 1 * s.MoveDirection
	}

	if s.Control.IsAccelerationReleased() {
		s.Acceleration = false
		s.YAxisForce = 0
		s.XAxisForce = 0
	}

}

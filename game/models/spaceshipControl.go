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

type SpaceshipKeyboardControlMap struct {
	Up           ebiten.Key
	Down         ebiten.Key
	Left         ebiten.Key
	Right        ebiten.Key
	Break        ebiten.Key
	Shoot        ebiten.Key
	Acceleration ebiten.Key
}

type SpaceshipGamepadControlMap struct {
	Up           ebiten.GamepadButton
	Down         ebiten.GamepadButton
	Left         ebiten.GamepadButton
	Right        ebiten.GamepadButton
	Break        ebiten.GamepadButton
	Shoot        ebiten.GamepadButton
	Acceleration ebiten.GamepadButton
}

func (s *SpaceshipGamepadControlMap) IsShoot() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Shoot)
}

func (s *SpaceshipGamepadControlMap) IsUp() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Up)
}

func (s *SpaceshipGamepadControlMap) IsAccelerationPressed() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Acceleration)
}

func (s *SpaceshipGamepadControlMap) IsAccelerationReleased() bool {
	return inpututil.IsGamepadButtonJustReleased(0, s.Acceleration)
}

func (s *SpaceshipGamepadControlMap) IsDown() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Down)
}

func (s *SpaceshipGamepadControlMap) IsLeft() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Left)
}

func (s *SpaceshipGamepadControlMap) IsRight() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Right)
}

func (s *SpaceshipGamepadControlMap) IsBreak() bool {
	return inpututil.IsGamepadButtonJustPressed(0, s.Break)
}

func (s *SpaceshipKeyboardControlMap) IsShoot() bool {
	return inpututil.IsKeyJustPressed(s.Shoot)
}

func (s *SpaceshipKeyboardControlMap) IsUp() bool {
	return inpututil.IsKeyJustPressed(s.Up)
}

func (s *SpaceshipKeyboardControlMap) IsAccelerationPressed() bool {
	return inpututil.IsKeyJustPressed(s.Acceleration)
}

func (s *SpaceshipKeyboardControlMap) IsAccelerationReleased() bool {
	return inpututil.IsKeyJustReleased(s.Acceleration)
}

func (s *SpaceshipKeyboardControlMap) IsDown() bool {
	return inpututil.IsKeyJustPressed(s.Down)
}

func (s *SpaceshipKeyboardControlMap) IsLeft() bool {
	return inpututil.IsKeyJustPressed(s.Left)
}

func (s *SpaceshipKeyboardControlMap) IsRight() bool {
	return inpututil.IsKeyJustPressed(s.Right)
}

func (s *SpaceshipKeyboardControlMap) IsBreak() bool {
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

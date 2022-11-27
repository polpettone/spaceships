package game

import "github.com/hajimehoshi/ebiten/inpututil"

func handleGamepadControls(s *Spaceship) {

	if s.GamepadControlMap == nil {
		return
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Up) {
		s.YAxisForce -= 1
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Down) {
		s.YAxisForce += 1
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Left) {
		s.XAxisForce -= 1
		s.MoveDirection = -1
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Right) {
		s.XAxisForce += 1
		s.MoveDirection = 1
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Break) {
		s.YAxisForce = 0
		s.XAxisForce = 0
	}

	if inpututil.IsGamepadButtonJustPressed(0, s.GamepadControlMap.Shoot) {
		s.ShootBullet = true
	}

}

func handleKeyboardControls(s *Spaceship) {

	if s.KeyboardControlMap == nil {
		return
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Up) {
		s.YAxisForce -= 1
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Down) {
		s.YAxisForce += 1
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Left) {
		s.XAxisForce -= 1
		s.MoveDirection = -1
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Right) {
		s.XAxisForce += 1
		s.MoveDirection = 1
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Break) {
		s.YAxisForce = 0
		s.XAxisForce = 0
	}

	if inpututil.IsKeyJustPressed(s.KeyboardControlMap.Shoot) {
		s.ShootBullet = true
	}

}

package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

func handleSoundControl(current bool) bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
		return !current
	}
	return current
}

func isQuitHit() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return true
	}
	return false
}

func handleResetGameControl() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		return true
	}
	return false
}

func handlePauseControl(current bool) bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyO) {
		return !current
	}
	return current
}

func handleDebugPrintControl(current bool) bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) {
		return !current
	}
	return current
}

func updateGamepads(g *SpaceshipGame) {

	for _, id := range inpututil.JustConnectedGamepadIDs() {
		log.Printf("connected gamepad ID: %d", id)
		g.GamepadIDs[id] = struct{}{}
	}

	for id := range g.GamepadIDs {
		if inpututil.IsGamepadJustDisconnected(id) {
			log.Printf("disconnected gamepad ID: %d", id)
			delete(g.GamepadIDs, id)
		}
	}

	for id := range g.GamepadIDs {
		maxButton := ebiten.GamepadButton(ebiten.GamepadButtonNum(id))

		for b := ebiten.GamepadButton(id); b < maxButton; b++ {
			if inpututil.IsGamepadButtonJustPressed(id, b) {
				log.Printf("button pressed: id: %d, button: %d", id, b)
			}
		}
	}
}

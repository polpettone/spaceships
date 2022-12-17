package game

import (
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/polpettone/gaming/spaceships/game/models"
)

func handleSoundControl(current bool) bool {
	if inpututil.IsKeyJustPressed(ebiten.KeyS) {
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

func handleControl(currentState models.GameState) models.GameState {

	if (currentState == models.GameOver ||
		currentState == models.ShowMenu) &&
		inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return models.Quit
	}

	if currentState == models.Running &&
		inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return models.MenuPreparation
	}

	if currentState == models.GameOver &&
		inpututil.IsKeyJustPressed(ebiten.KeyR) {
		return models.Reset
	}

	if currentState != models.Pause &&
		inpututil.IsKeyJustPressed(ebiten.KeyO) {
		return models.Pause
	}

	if currentState == models.Pause &&
		inpututil.IsKeyJustPressed(ebiten.KeyO) {
		return models.Running
	}

	if currentState == models.GameOver &&
		inpututil.IsKeyJustPressed(ebiten.KeyM) {
		return models.MenuPreparation
	}

	if currentState == models.ShowMenu &&
		inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		return models.ScenePreparation
	}

	return currentState
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

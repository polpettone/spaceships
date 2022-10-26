package main

import (
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

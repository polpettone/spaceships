package models

import "github.com/hajimehoshi/ebiten/v2/audio"

func CreateHumanControledSpaceships(
	g GameConfig,
	audioContext *audio.Context,
	gamepadControlMap *SpaceshipGamepadControlMap,
	keyboardControlMap *SpaceshipKeyboardControlMap) (*Spaceship, *Spaceship, error) {

	img1, err := NewGameObjectImage("assets/images/spaceships/star-wars-2.png", 0.2, -1)

	if err != nil {
		return nil, nil, err
	}

	damageImg1, err := NewGameObjectImage("assets/images/spaceships/star-wars-2-red.png", 0.2, -1)
	if err != nil {
		return nil, nil, err
	}

	spaceship1, err := NewSpaceship(
		audioContext,
		g.HealthSpaceship1,
		g.BulletCountSpaceship1,
		"Player 1",
		NewPos(100, 300),
		nil,
		gamepadControlMap,
		img1,
		damageImg1,
		"s1")

	if err != nil {
		return nil, nil, err
	}

	img2, err := NewGameObjectImage("assets/images/spaceships/star-wars-3.png", 0.2, -1)
	if err != nil {
		return nil, nil, err
	}

	damageImg2, err := NewGameObjectImage("assets/images/spaceships/star-wars-3-red.png", 0.2, -1)
	if err != nil {
		return nil, nil, err
	}

	spaceship2, err := NewSpaceship(
		audioContext,
		g.HealthSpaceship2,
		g.BulletCountSpaceship2,
		"Player 2",
		NewPos(1900, 300),
		keyboardControlMap,
		nil,
		img2,
		damageImg2,
		"s2")
	spaceship2.MoveDirection *= -1

	if err != nil {
		return nil, nil, err
	}

	return spaceship1, spaceship2, nil
}

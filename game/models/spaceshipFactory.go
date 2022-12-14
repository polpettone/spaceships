package models

import "github.com/hajimehoshi/ebiten/v2/audio"

func CreateSpaceship1(
	g GameConfig,
	audioContext *audio.Context,
	spaceshipControl SpaceshipControl) (*Spaceship, error) {

	img, err := NewGameObjectImage("assets/images/spaceships/star-wars-2.png", 0.2, -1)

	if err != nil {
		return nil, err
	}

	damageImg, err := NewGameObjectImage(
		"assets/images/spaceships/star-wars-2-red.png", 0.2, -1)
	if err != nil {
		return nil, err
	}

	spaceship, err := NewSpaceship(
		audioContext,
		g.HealthSpaceship1,
		g.BulletCountSpaceship1,
		"Player 1",
		NewPos(100, 300),
		spaceshipControl,
		img,
		damageImg,
		"s1")

	if err != nil {
		return nil, err
	}

	return spaceship, nil
}

func CreateSpaceship2(
	g GameConfig,
	audioContext *audio.Context,
	spaceshipControl SpaceshipControl) (*Spaceship, error) {

	img, err := NewGameObjectImage("assets/images/spaceships/star-wars-3.png", 0.2, -1)

	if err != nil {
		return nil, err
	}

	damageImg, err := NewGameObjectImage(
		"assets/images/spaceships/star-wars-3-red.png", 0.2, -1)
	if err != nil {
		return nil, err
	}

	spaceship, err := NewSpaceship(
		audioContext,
		g.HealthSpaceship2,
		g.BulletCountSpaceship2,
		"Player 2",
		NewPos(1900, 300),
		spaceshipControl,
		img,
		damageImg,
		"s2")
	spaceship.MoveDirection *= -1

	if err != nil {
		return nil, err
	}

	return spaceship, nil
}

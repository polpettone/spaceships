package models

import (
	"github.com/hajimehoshi/ebiten"
)

type Scene1 struct {
	Spaceship1  *Spaceship
	Spaceship2  *Spaceship
	GameObjects map[string]GameObject
	GameConfig  GameConfig
	TickCounter int
	MaxX        int
	MaxY        int
}

func NewScene1(config GameConfig) (*Scene1, error) {

	spaceship1, err := CreateSpaceship1(
		config,
		audioContext,
		keyboardControlMap)

	if err != nil {
		return nil, err
	}

	spaceship2, err := CreateSpaceship2(
		config,
		audioContext,
		gamepadControlMap)

	if err != nil {
		return nil, err
	}

	gameObjects := map[string]GameObject{}

	return &Scene1{
		Spaceship1:  spaceship1,
		Spaceship2:  spaceship2,
		GameObjects: gameObjects,
		GameConfig:  config,
		TickCounter: 0,
		MaxX:        screenWidth,
		MaxY:        screenHeight,
	}, nil

}

func (g *Scene1) GetName() string {
	return "Scene 1"
}

func (g *Scene1) GetConfig() GameConfig {
	return g.GameConfig
}

func (g *Scene1) GetTickCounter() int {
	return g.TickCounter
}

func (g *Scene1) GetMaxX() int {
	return g.MaxX
}

func (g *Scene1) GetMaxY() int {
	return g.MaxY
}

func (g *Scene1) Update(screen *ebiten.Image) (GameState, error) {

	g.TickCounter++

	if checkCriteria(
		g.GameConfig.TPS,
		g.TickCounter,
		g.GameConfig.EnemiesPerSecond) {
		g.PutNewEnemies(1)
	}

	if checkCriteria(
		g.GameConfig.TPS,
		g.TickCounter,
		g.GameConfig.StarsPerSecond) {
		g.PutStars(1)
	}

	if checkCriteria(
		g.GameConfig.TPS,
		g.TickCounter,
		g.GameConfig.AmmoPerSecond) {
		g.PutNewAmmos(1)
	}

	if g.TickCounter%10000 == 0 {
		g.TickCounter = 0
	}

	SpaceshipCollisionDetection(g.Spaceship1, g.GameObjects)
	SpaceshipCollisionDetection(g.Spaceship2, g.GameObjects)

	BulletSkyObjectCollisionDetection(g.GetGameObjects())

	g.Spaceship1.Update(g)
	g.Spaceship2.Update(g)

	for _, o := range g.GameObjects {
		o.Update()
	}

	DeleteObjectsOutOfView(g)

	result := g.CheckGameOverCriteria()
	if result {
		return GameOver, nil
	}

	return Running, nil
}

func (g *Scene1) Draw(screen *ebiten.Image) {

	for _, o := range g.GameObjects {
		o.Draw(screen)
	}

	g.Spaceship1.Draw(screen)
	g.Spaceship2.Draw(screen)

	g.Spaceship1.DrawState(screen, 100, 10)
	g.Spaceship2.DrawState(screen, g.GetMaxX()-200, 10)

}

func (g *Scene1) Reset() {
	g.GameObjects = map[string]GameObject{}

	g.Spaceship1.Reset(
		g.GameConfig.HealthSpaceship1,
		g.GameConfig.InitialPosSpaceship1,
		g.GameConfig.BulletCountSpaceship1,
		1)

	g.Spaceship2.Reset(
		g.GameConfig.HealthSpaceship2,
		g.GameConfig.InitialPosSpaceship2,
		g.GameConfig.BulletCountSpaceship2,
		-1)
}

func (g *Scene1) AddGameObject(o GameObject) {
	g.GameObjects[o.GetID()] = o
}

func (g *Scene1) GetGameObjects() map[string]GameObject {
	return g.GameObjects
}

func (g *Scene1) GetSpaceship1() *Spaceship {
	return g.Spaceship1
}

func (g *Scene1) GetSpaceship2() *Spaceship {
	return g.Spaceship2
}

func (g *Scene1) PutStars(count int) {
	newSkyObjects := CreateStarAtRandomPosition(
		0,
		0,
		screenWidth,
		screenHeight,
		count,
		g.GameConfig.StarVelocity)

	for _, nO := range newSkyObjects {
		g.GameObjects[nO.GetID()] = nO
	}
}

func (g *Scene1) PutNewEnemies(count int) {
	newSkyObjects := CreateSkyObjectAtRandomPosition(
		0,
		0,
		screenWidth,
		screenHeight,
		count)

	for _, nO := range newSkyObjects {
		g.GameObjects[nO.GetID()] = nO
	}
}

func (g *Scene1) PutNewAmmos(count int) {
	newAmmos := CreateAmmoAtRandomPosition(
		0, 0, screenWidth, screenHeight, count)

	for _, nO := range newAmmos {
		g.GameObjects[nO.GetID()] = nO
	}
}

func (g *Scene1) CheckGameOverCriteria() bool {

	if !g.GetSpaceship1().Alive() || !g.GetSpaceship2().Alive() {
		return true
	}

	return false
}

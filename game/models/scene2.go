package models

import (
	"github.com/hajimehoshi/ebiten"
)

type Scene2 struct {
	Spaceship1  *Spaceship
	GameObjects map[string]GameObject
	SceneConfig SceneConfig
	TickCounter int
	MaxX        int
	MaxY        int
}

func NewScene2(config SceneConfig) (*Scene2, error) {

	spaceship1, err := CreateSpaceship1(
		config,
		audioContext,
		keyboardControlMap)

	if err != nil {
		return nil, err
	}

	gameObjects := map[string]GameObject{}

	return &Scene2{
		Spaceship1:  spaceship1,
		GameObjects: gameObjects,
		SceneConfig: config,
		TickCounter: 0,
		MaxX:        config.GameConfig.MaxX,
		MaxY:        config.GameConfig.MaxY,
	}, nil

}

func (g *Scene2) GetName() string {
	return "2 - One ship nothing more"
}

func (g *Scene2) GetConfig() SceneConfig {
	return g.SceneConfig
}

func (g *Scene2) GetTickCounter() int {
	return g.TickCounter
}

func (g *Scene2) GetMaxX() int {
	return g.MaxX
}

func (g *Scene2) GetMaxY() int {
	return g.MaxY
}

func (g *Scene2) Update(screen *ebiten.Image) (GameState, error) {
	g.TickCounter++

	g.Spaceship1.Update(g)

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

func (g *Scene2) Draw(screen *ebiten.Image) {

	for _, o := range g.GameObjects {
		o.Draw(screen)
	}

	g.Spaceship1.Draw(screen)

	g.Spaceship1.DrawState(screen, 100, 10)
}

func (g *Scene2) Reset() {
	g.GameObjects = map[string]GameObject{}

	g.Spaceship1.Reset(
		g.SceneConfig.HealthSpaceship1,
		g.SceneConfig.InitialPosSpaceship1,
		g.SceneConfig.BulletCountSpaceship1,
		1)
}

func (g *Scene2) AddGameObject(o GameObject) {
	g.GameObjects[o.GetID()] = o
}

func (g *Scene2) GetGameObjects() map[string]GameObject {
	return g.GameObjects
}

func (g *Scene2) GetSpaceship1() *Spaceship {
	return g.Spaceship1
}

func (g *Scene2) GetSpaceship2() *Spaceship {
	return nil
}

func (g *Scene2) CheckGameOverCriteria() bool {

	if !g.GetSpaceship1().Alive() {
		return true
	}

	return false
}

func (g *Scene2) PutNewAmmos(count int)   {}
func (g *Scene2) PutStars(count int)      {}
func (g *Scene2) PutNewEnemies(count int) {}

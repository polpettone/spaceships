package models

import (
	"github.com/hajimehoshi/ebiten"
)

type Scene4 struct {
	Spaceship1  *Spaceship
	GameObjects map[string]GameObject
	SceneConfig SceneConfig
	TickCounter int
	MaxX        int
	MaxY        int
	AIControl   *AIControl
}

func NewScene4(config SceneConfig) (*Scene4, error) {

	aiControl := createControl()

	spaceship1, err := CreateSpaceship1(
		config,
		audioContext,
		aiControl)

	if err != nil {
		return nil, err
	}

	gameObjects := map[string]GameObject{}

	return &Scene4{
		Spaceship1:  spaceship1,
		GameObjects: gameObjects,
		SceneConfig: config,
		TickCounter: 0,
		MaxX:        config.GameConfig.MaxX,
		MaxY:        config.GameConfig.MaxY,
		AIControl:   createControl(),
	}, nil

}

func createControl() *AIControl {
	return &AIControl{}
}

func (g *Scene4) GetName() string {
	return "4 - AI controled ship, nothing more"
}

func (g *Scene4) GetConfig() SceneConfig {
	return g.SceneConfig
}

func (g *Scene4) GetTickCounter() int {
	return g.TickCounter
}

func (g *Scene4) GetMaxX() int {
	return g.MaxX
}

func (g *Scene4) GetMaxY() int {
	return g.MaxY
}

func (g *Scene4) Update(screen *ebiten.Image) (GameState, error) {
	g.TickCounter++

	if g.TickCounter%60 == 0 {
		chanceControl(g.Spaceship1, g.GetMaxX(), g.GetMaxY())
	}

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

func (g *Scene4) Draw(screen *ebiten.Image) {

	for _, o := range g.GameObjects {
		o.Draw(screen)
	}

	g.Spaceship1.Draw(screen)

	g.Spaceship1.DrawState(screen, 100, 10)
}

func (g *Scene4) Reset() {
	g.GameObjects = map[string]GameObject{}

	g.Spaceship1.Reset(
		g.SceneConfig.HealthSpaceship1,
		g.SceneConfig.InitialPosSpaceship1,
		g.SceneConfig.BulletCountSpaceship1,
		1)
}

func (g *Scene4) AddGameObject(o GameObject) {
	g.GameObjects[o.GetID()] = o
}

func (g *Scene4) GetGameObjects() map[string]GameObject {
	return g.GameObjects
}

func (g *Scene4) GetSpaceship1() *Spaceship {
	return g.Spaceship1
}

func (g *Scene4) GetSpaceship2() *Spaceship {
	return nil
}

func (g *Scene4) CheckGameOverCriteria() bool {

	if !g.GetSpaceship1().Alive() {
		return true
	}

	return false
}

func (g *Scene4) PutNewAmmos(count int)   {}
func (g *Scene4) PutStars(count int)      {}
func (g *Scene4) PutNewEnemies(count int) {}

package models

import "github.com/hajimehoshi/ebiten"

type Scene3 struct {
	Spaceship1  *Spaceship
	Spaceship2  *Spaceship
	GameObjects map[string]GameObject
	SceneConfig SceneConfig
	TickCounter int
	MaxX        int
	MaxY        int
}

func NewScene3(config SceneConfig) (*Scene3, error) {

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

	return &Scene3{
		Spaceship1:  spaceship1,
		Spaceship2:  spaceship2,
		GameObjects: gameObjects,
		SceneConfig: config,
		TickCounter: 0,
		MaxX:        config.GameConfig.MaxX,
		MaxY:        config.GameConfig.MaxY,
	}, nil

}

func (g *Scene3) GetName() string {
	return "3 - Two ships and ammo"
}

func (g *Scene3) GetConfig() SceneConfig {
	return g.SceneConfig
}

func (g *Scene3) GetTickCounter() int {
	return g.TickCounter
}

func (g *Scene3) GetMaxX() int {
	return g.MaxX
}

func (g *Scene3) GetMaxY() int {
	return g.MaxY
}

func (g *Scene3) Update(screen *ebiten.Image) (GameState, error) {

	g.TickCounter++

	if checkCriteria(
		g.SceneConfig.GameConfig.TPS,
		g.TickCounter,
		g.SceneConfig.AmmoPerSecond) {
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

func (g *Scene3) Draw(screen *ebiten.Image) {

	for _, o := range g.GameObjects {
		o.Draw(screen)
	}

	g.Spaceship1.Draw(screen)
	g.Spaceship2.Draw(screen)

	g.Spaceship1.DrawState(screen, 100, 10)
	g.Spaceship2.DrawState(screen, g.GetMaxX()-200, 10)

}

func (g *Scene3) Reset() {
	g.GameObjects = map[string]GameObject{}

	g.Spaceship1.Reset(
		g.SceneConfig.HealthSpaceship1,
		g.SceneConfig.InitialPosSpaceship1,
		g.SceneConfig.BulletCountSpaceship1,
		1)

	g.Spaceship2.Reset(
		g.SceneConfig.HealthSpaceship2,
		g.SceneConfig.InitialPosSpaceship2,
		g.SceneConfig.BulletCountSpaceship2,
		-1)
}

func (g *Scene3) AddGameObject(o GameObject) {
	g.GameObjects[o.GetID()] = o
}

func (g *Scene3) GetGameObjects() map[string]GameObject {
	return g.GameObjects
}

func (g *Scene3) GetSpaceship1() *Spaceship {
	return g.Spaceship1
}

func (g *Scene3) GetSpaceship2() *Spaceship {
	return g.Spaceship2
}

func (g *Scene3) PutNewAmmos(count int) {
	newAmmos := CreateAmmoAtRandomPosition(
		0, 0, screenWidth, screenHeight, count)

	for _, nO := range newAmmos {
		g.GameObjects[nO.GetID()] = nO
	}
}

func (g *Scene3) CheckGameOverCriteria() bool {

	if !g.GetSpaceship1().Alive() || !g.GetSpaceship2().Alive() {
		return true
	}

	return false
}

func (g *Scene3) PutStars(count int)      {}
func (g *Scene3) PutNewEnemies(count int) {}

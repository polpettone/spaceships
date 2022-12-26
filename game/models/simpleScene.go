package models

import "github.com/hajimehoshi/ebiten"

type SimpleScene struct {
	Spaceship1  *Spaceship
	Spaceship2  *Spaceship
	GameObjects map[string]GameObject
	SceneConfig SceneConfig
	TickCounter int
	MaxX        int
	MaxY        int
	Name        string
}

func NewSimpleScene(
	name string,
	config SceneConfig) (*SimpleScene, error) {

	spaceship1, err := CreateSpaceship1(
		config,
		audioContext,
		keyboardControlMap)

	if err != nil {
		return nil, err
	}

	var spaceship2 *Spaceship
	if config.SecondShipEnabled {
		spaceship2, err = CreateSpaceship2(
			config,
			audioContext,
			gamepadControlMap)

		if err != nil {
			return nil, err
		}
	}

	gameObjects := map[string]GameObject{}

	return &SimpleScene{
		Spaceship1:  spaceship1,
		Spaceship2:  spaceship2,
		GameObjects: gameObjects,
		SceneConfig: config,
		TickCounter: 0,
		MaxX:        config.GameConfig.MaxX,
		MaxY:        config.GameConfig.MaxY,
		Name:        name,
	}, nil

}

func (g *SimpleScene) GetName() string {
	return g.Name
}

func (g *SimpleScene) GetConfig() SceneConfig {
	return g.SceneConfig
}

func (g *SimpleScene) GetTickCounter() int {
	return g.TickCounter
}

func (g *SimpleScene) GetMaxX() int {
	return g.MaxX
}

func (g *SimpleScene) GetMaxY() int {
	return g.MaxY
}

func (g *SimpleScene) Update(screen *ebiten.Image) (GameState, error) {

	g.TickCounter++

	if checkCriteria(
		g.SceneConfig.GameConfig.TPS,
		g.TickCounter,
		g.SceneConfig.EnemiesPerSecond) {
		g.PutNewEnemies(1)
	}

	if checkCriteria(
		g.SceneConfig.GameConfig.TPS,
		g.TickCounter,
		g.SceneConfig.StarsPerSecond) {
		g.PutStars(1)
	}

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

func (g *SimpleScene) Draw(screen *ebiten.Image) {

	for _, o := range g.GameObjects {
		o.Draw(screen)
	}

	g.Spaceship1.Draw(screen)
	g.Spaceship1.DrawState(screen, 100, 10)

	if g.SceneConfig.SecondShipEnabled {
		g.Spaceship2.Draw(screen)
		g.Spaceship2.DrawState(screen, g.GetMaxX()-200, 10)
	}

}

func (g *SimpleScene) Reset() {
	g.GameObjects = map[string]GameObject{}

	g.Spaceship1.Reset(
		g.SceneConfig.HealthSpaceship1,
		g.SceneConfig.InitialPosSpaceship1,
		g.SceneConfig.BulletCountSpaceship1,
		1)

	if g.SceneConfig.SecondShipEnabled {
		g.Spaceship2.Reset(
			g.SceneConfig.HealthSpaceship2,
			g.SceneConfig.InitialPosSpaceship2,
			g.SceneConfig.BulletCountSpaceship2,
			-1)
	}
}

func (g *SimpleScene) AddGameObject(o GameObject) {
	g.GameObjects[o.GetID()] = o
}

func (g *SimpleScene) GetGameObjects() map[string]GameObject {
	return g.GameObjects
}

func (g *SimpleScene) GetSpaceship1() *Spaceship {
	return g.Spaceship1
}

func (g *SimpleScene) GetSpaceship2() *Spaceship {
	return g.Spaceship2
}

func (g *SimpleScene) PutStars(count int) {
	newSkyObjects := CreateStarAtRandomPosition(
		0,
		0,
		screenWidth,
		screenHeight,
		count,
		g.SceneConfig.StarVelocity)

	for _, nO := range newSkyObjects {
		g.GameObjects[nO.GetID()] = nO
	}
}

func (g *SimpleScene) PutNewEnemies(count int) {
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

func (g *SimpleScene) PutNewAmmos(count int) {
	newAmmos := CreateAmmoAtRandomPosition(
		0, 0, screenWidth, screenHeight, count)

	for _, nO := range newAmmos {
		g.GameObjects[nO.GetID()] = nO
	}
}

func (g *SimpleScene) CheckGameOverCriteria() bool {

	if g.SceneConfig.SecondShipEnabled {
		if !g.GetSpaceship1().Alive() || !g.GetSpaceship2().Alive() {
			return true
		}
	} else {
		if !g.GetSpaceship1().Alive() {
			return true
		}
	}

	return false
}

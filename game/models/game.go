package models

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/polpettone/gaming/spaceships/engine"
)

type Game interface {
	GetMaxX() int
	GetMaxY() int

	Layout(outsideWidth, outsideHeight int) (int, int)
	Update(screen *ebiten.Image) error

	GetConfig() GameConfig
}

var (
	audioContext *audio.Context

	keyboardControlMap   *KeyboardControl
	gamepadControlMap    *GamepadControl
	ps3GamepadControlMap *GamepadControl
)

func init() {
	audioContext = audio.NewContext(44100)

	keyboardControlMap = &KeyboardControl{
		Up:           ebiten.KeyK,
		Down:         ebiten.KeyJ,
		Left:         ebiten.KeyH,
		Right:        ebiten.KeyL,
		Break:        ebiten.KeyB,
		Shoot:        ebiten.KeyN,
		Acceleration: ebiten.KeySpace,
	}

	gamepadControlMap = &GamepadControl{
		Up:           ebiten.GamepadButton11,
		Down:         ebiten.GamepadButton13,
		Left:         ebiten.GamepadButton14,
		Right:        ebiten.GamepadButton12,
		Break:        ebiten.GamepadButton4,
		Shoot:        ebiten.GamepadButton0,
		Acceleration: ebiten.GamepadButton5,
	}

	ps3GamepadControlMap = &GamepadControl{
		Up:           ebiten.GamepadButton13,
		Down:         ebiten.GamepadButton14,
		Left:         ebiten.GamepadButton15,
		Right:        ebiten.GamepadButton16,
		Break:        ebiten.GamepadButton6,
		Shoot:        ebiten.GamepadButton0,
		Acceleration: ebiten.GamepadButton5,
	}

}

type Scene interface {
	GetMaxX() int
	GetMaxY() int

	AddGameObject(o GameObject)
	GetGameObjects() map[string]GameObject
	GetSpaceship1() *Spaceship
	GetSpaceship2() *Spaceship

	PutNewAmmos(count int)
	PutStars(count int)
	PutNewEnemies(count int)

	Update(screen *ebiten.Image) error
	Draw(screen *ebiten.Image)
	Reset()

	GetConfig() GameConfig

	GetTickCounter() int
}

type Scene1 struct {
	Spaceship1  *Spaceship
	Spaceship2  *Spaceship
	GameObjects map[string]GameObject
	GameConfig  GameConfig
	TickCounter int
	MaxX        int
	MaxY        int
}

const (
	screenWidth  = 2000
	screenHeight = 1000
)

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

func (g *Scene1) Update(screen *ebiten.Image) error {

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

	spaceshipCollisionDetection(g.Spaceship1, g.GameObjects)
	spaceshipCollisionDetection(g.Spaceship2, g.GameObjects)

	bulletSkyObjectCollisionDetection(g.GetGameObjects())

	g.Spaceship1.Update(g)
	g.Spaceship2.Update(g)

	for _, o := range g.GameObjects {
		o.Update()
	}

	deleteObjectsOutOfView(g)

	return nil
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

func checkCriteria(
	TPS float64,
	currentTick int,
	actionPerSecond float64) bool {

	if actionPerSecond == 0 {
		return false
	}

	var x int
	rate := TPS / actionPerSecond
	if rate > 1 {
		x = int(rate)
	} else {
		x = 1
	}

	return currentTick%x == 0
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

func spaceshipCollisionDetection(s *Spaceship, gameObjects map[string]GameObject) {

	for k, o := range gameObjects {

		if o.GetType() == Enemy && o.IsAlive() {
			sW, _ := s.GetSize()
			oW, _ := o.GetSize()
			if engine.CollisionDetection(
				s.Pos.X,
				s.Pos.Y,
				o.GetPos().X,
				o.GetPos().Y,
				sW,
				oW,
				0) {
				s.Damage()
				o.Destroy()
			}
		}

		if o.GetType() == Item {
			sW, _ := s.GetSize()
			oW, _ := o.GetSize()
			if engine.CollisionDetection(
				s.Pos.X,
				s.Pos.Y,
				o.GetPos().X,
				o.GetPos().Y,
				sW,
				oW,
				0) {
				s.BulletCount += 10
				delete(gameObjects, k)
			}
		}

		if o.GetType() == Weapon && o.GetSignature() != s.GetSignature() {
			sW, _ := s.GetSize()
			oW, _ := o.GetSize()
			if engine.CollisionDetection(
				s.Pos.X,
				s.Pos.Y,
				o.GetPos().X,
				o.GetPos().Y,
				sW,
				oW,
				0) {
				s.Damage()
				delete(gameObjects, k)
			}
		}

	}
}

func deleteObjectsOutOfView(g Scene) {
	var ids []string
	for k, o := range g.GetGameObjects() {
		x := o.GetPos().X
		y := o.GetPos().Y
		if x > g.GetMaxX() || x < 0 || y > g.GetMaxX() || y < 0 {
			ids = append(ids, k)
		}
	}
	for _, k := range ids {
		delete(g.GetGameObjects(), k)
	}
}

func bulletSkyObjectCollisionDetection(gameObjects map[string]GameObject) {

	for k, o := range gameObjects {
		if o.GetType() == Weapon {
			for _, x := range gameObjects {
				if x.GetType() == Enemy && x.IsAlive() {
					oW, _ := o.GetSize()
					xW, _ := x.GetSize()
					if engine.CollisionDetection(
						o.GetPos().X,
						o.GetPos().Y,
						x.GetPos().X,
						x.GetPos().Y,
						oW,
						xW,
						0) {
						x.Destroy()
						delete(gameObjects, k)
					}
				}
			}
		}
	}
}

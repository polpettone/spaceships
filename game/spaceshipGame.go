package game

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/polpettone/gaming/spaceships/engine"
	"github.com/polpettone/gaming/spaceships/game/models"
)

const (
	screenWidth  = 2000
	screenHeight = 1000
)

var (
	audioContext *audio.Context

	keyboardControlMap   *models.KeyboardControl
	gamepadControlMap    *models.GamepadControl
	ps3GamepadControlMap *models.GamepadControl
)

func init() {
	audioContext = audio.NewContext(44100)

	keyboardControlMap = &models.KeyboardControl{
		Up:           ebiten.KeyK,
		Down:         ebiten.KeyJ,
		Left:         ebiten.KeyH,
		Right:        ebiten.KeyL,
		Break:        ebiten.KeyB,
		Shoot:        ebiten.KeyN,
		Acceleration: ebiten.KeySpace,
	}

	gamepadControlMap = &models.GamepadControl{
		Up:           ebiten.GamepadButton11,
		Down:         ebiten.GamepadButton13,
		Left:         ebiten.GamepadButton14,
		Right:        ebiten.GamepadButton12,
		Break:        ebiten.GamepadButton4,
		Shoot:        ebiten.GamepadButton0,
		Acceleration: ebiten.GamepadButton5,
	}

	ps3GamepadControlMap = &models.GamepadControl{
		Up:           ebiten.GamepadButton13,
		Down:         ebiten.GamepadButton14,
		Left:         ebiten.GamepadButton15,
		Right:        ebiten.GamepadButton16,
		Break:        ebiten.GamepadButton6,
		Shoot:        ebiten.GamepadButton0,
		Acceleration: ebiten.GamepadButton5,
	}

}

type GameState int64

const (
	Running GameState = iota
	GameOver
)

type Scene func(g *SpaceshipGame)

type SpaceshipGame struct {
	GameConfig      models.GameConfig
	BackgroundImage *ebiten.Image
	Spaceship1      *models.Spaceship
	Spaceship2      *models.Spaceship
	GameObjects     map[string]models.GameObject
	DebugScreen     *DebugScreen
	MaxX            int
	MaxY            int
	DebugPrint      bool

	BackgroundSound *audio.Player

	TickCounter int

	Pause bool

	SoundOn bool

	KilledEnemies int

	State GameState

	GamepadIDs map[int]struct{}
}

func NewGame(config models.GameConfig) (models.Game, error) {

	debugScreen, err := NewDebugScreen()
	if err != nil {
		return nil, err
	}

	spaceship1, err := models.CreateSpaceship1(
		config,
		audioContext,
		keyboardControlMap)

	if err != nil {
		return nil, err
	}

	spaceship2, err := models.CreateSpaceship2(
		config,
		audioContext,
		gamepadControlMap)

	if err != nil {
		return nil, err
	}

	gameObjects := map[string]models.GameObject{}

	g := &SpaceshipGame{
		GameConfig:  config,
		GameObjects: gameObjects,
		DebugScreen: debugScreen,
		Spaceship1:  spaceship1,
		Spaceship2:  spaceship2,
		TickCounter: 0,
		MaxX:        screenWidth,
		MaxY:        screenHeight,
		DebugPrint:  false,
		Pause:       false,
		SoundOn:     false,
		State:       Running,
		GamepadIDs:  map[int]struct{}{},
	}

	return g, nil
}

func (g *SpaceshipGame) Reset() {
	g.GameObjects = map[string]models.GameObject{}

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

	g.TickCounter = 0
	g.Pause = false
	g.State = Running
	g.KilledEnemies = 0
}

func (g *SpaceshipGame) GetConfig() models.GameConfig {
	return g.GameConfig
}

func (g *SpaceshipGame) GetMaxX() int {
	return g.MaxX
}

func (g *SpaceshipGame) GetMaxY() int {
	return g.MaxY
}

func (g *SpaceshipGame) GetTickCounter() int {
	return g.TickCounter
}

func (g *SpaceshipGame) AddGameObject(o models.GameObject) {
	g.GameObjects[o.GetID()] = o
}

func (g *SpaceshipGame) GetGameObjects() map[string]models.GameObject {
	return g.GameObjects
}

func (g *SpaceshipGame) GetSpaceship1() *models.Spaceship {
	return g.Spaceship1
}

func (g *SpaceshipGame) GetSpaceship2() *models.Spaceship {
	return g.Spaceship2
}

func handleBackgroundSound(g *SpaceshipGame) {
	if g.BackgroundSound != nil {
		if g.Pause {
			g.BackgroundSound.Pause()
		}

		if !g.BackgroundSound.IsPlaying() && g.SoundOn {
			g.BackgroundSound.Play()
		}

		if g.BackgroundSound.IsPlaying() && !g.SoundOn {
			g.BackgroundSound.Pause()
		}
	}
}

func (g *SpaceshipGame) Update(screen *ebiten.Image) error {

	Scene1(g)

	updateGamepads(g)

	checkGameOverCriteria(g)

	if isQuitHit() {
		os.Exit(0)
	}

	if handleResetGameControl() && g.State == GameOver {
		g.Reset()
		g.State = Running
	}

	g.Pause = handlePauseControl(g.Pause)
	g.SoundOn = handleSoundControl(g.SoundOn)

	if g.State == GameOver {
		return nil
	}

	if g.Pause {
		return nil
	}

	g.DebugPrint = handleDebugPrintControl(g.DebugPrint)

	spaceshipCollisionDetection(g.Spaceship1, g.GameObjects)
	spaceshipCollisionDetection(g.Spaceship2, g.GameObjects)

	bulletSkyObjectCollisionDetection(g)

	g.Spaceship1.Update(g)
	g.Spaceship2.Update(g)

	for _, o := range g.GameObjects {
		o.Update()
	}

	g.DebugScreen.Update(g)

	deleteObjectsOutOfView(g)

	return nil
}

func (g *SpaceshipGame) Draw(screen *ebiten.Image) {

	for _, o := range g.GameObjects {
		o.Draw(screen)
	}

	g.Spaceship1.Draw(screen)
	g.Spaceship2.Draw(screen)

	if g.DebugPrint {
		g.DebugScreen.Draw(screen, g)
	}

	g.Spaceship1.DrawState(screen, 100, 10)
	g.Spaceship2.DrawState(screen, g.GetMaxX()-200, 10)

	if g.State == GameOver {
		drawGameOverScreen(g, screen)
	}
}

func (g *SpaceshipGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func checkGameOverCriteria(g *SpaceshipGame) {

	if !g.Spaceship1.Alive() {
		g.State = GameOver
	}

	if !g.Spaceship2.Alive() {
		g.State = GameOver
	}

}

func bulletSkyObjectCollisionDetection(g *SpaceshipGame) {

	for k, o := range g.GameObjects {
		if o.GetType() == models.Weapon {
			for _, x := range g.GameObjects {
				if x.GetType() == models.Enemy && x.IsAlive() {
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
						g.KilledEnemies += 1
						delete(g.GameObjects, k)
					}
				}
			}
		}
	}
}

func spaceshipCollisionDetection(s *models.Spaceship, gameObjects map[string]models.GameObject) {

	for k, o := range gameObjects {

		if o.GetType() == models.Enemy && o.IsAlive() {
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

		if o.GetType() == models.Item {
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

		if o.GetType() == models.Weapon && o.GetSignature() != s.GetSignature() {
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

func drawGameOverScreen(g *SpaceshipGame, screen *ebiten.Image) {

	var livingSpaceship *models.Spaceship
	if g.Spaceship1.Alive() {
		livingSpaceship = g.Spaceship1
	} else {
		livingSpaceship = g.Spaceship2
	}

	t := fmt.Sprintf(
		`GAME OVER 
%s has won
Press R for New Game
Press Q for Quit`,
		livingSpaceship.PilotName)

	text.Draw(screen, t, engine.MplusBigFont, 700, 300, color.White)
}

func (g *SpaceshipGame) PutStars(count int) {
	newSkyObjects := models.CreateStarAtRandomPosition(0, 0, screenWidth, screenHeight, count, g.GameConfig.StarVelocity)

	for _, nO := range newSkyObjects {
		g.GameObjects[nO.GetID()] = nO
	}
}

func (g *SpaceshipGame) PutNewEnemies(count int) {
	newSkyObjects := models.CreateSkyObjectAtRandomPosition(
		0, 0, screenWidth, screenHeight, count)

	for _, nO := range newSkyObjects {
		g.GameObjects[nO.GetID()] = nO
	}
}

func (g *SpaceshipGame) PutNewAmmos(count int) {
	newAmmos := models.CreateAmmoAtRandomPosition(
		0, 0, screenWidth, screenHeight, count)

	for _, nO := range newAmmos {
		g.GameObjects[nO.GetID()] = nO
	}
}

func deleteObjectsOutOfView(g *SpaceshipGame) {
	var ids []string
	for k, o := range g.GameObjects {
		x := o.GetPos().X
		y := o.GetPos().Y
		if x > g.MaxX || x < 0 || y > g.MaxX || y < 0 {
			ids = append(ids, k)
		}
	}
	for _, k := range ids {
		delete(g.GameObjects, k)
	}
}

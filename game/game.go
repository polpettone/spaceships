package game

import (
	"fmt"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/polpettone/gaming/natalito/engine"
)

const (
	screenWidth            = 2000
	screenHeight           = 1000
	spaceshipWallTolerance = 10
)

var (
	audioContext *audio.Context

	keyboardControlMap *SpaceshipKeyboardControlMap

	gamepadControlMap    *SpaceshipGamepadControlMap
	ps3GamepadControlMap *SpaceshipGamepadControlMap
)

func init() {
	audioContext = audio.NewContext(44100)

	keyboardControlMap = &SpaceshipKeyboardControlMap{
		Up:    ebiten.KeyK,
		Down:  ebiten.KeyJ,
		Left:  ebiten.KeyH,
		Right: ebiten.KeyL,
		Break: ebiten.KeySpace,
		Shoot: ebiten.KeyN,
	}

	gamepadControlMap = &SpaceshipGamepadControlMap{
		Up:    ebiten.GamepadButton11,
		Down:  ebiten.GamepadButton13,
		Left:  ebiten.GamepadButton14,
		Right: ebiten.GamepadButton12,
		Break: ebiten.GamepadButton4,
		Shoot: ebiten.GamepadButton0,
	}

	ps3GamepadControlMap = &SpaceshipGamepadControlMap{
		Up:    ebiten.GamepadButton13,
		Down:  ebiten.GamepadButton14,
		Left:  ebiten.GamepadButton15,
		Right: ebiten.GamepadButton16,
		Break: ebiten.GamepadButton6,
		Shoot: ebiten.GamepadButton0,
	}

}

type GameState int64

const (
	Running GameState = iota
	GameOver
)

type Game interface {
	GetMaxX() int
	GetMaxY() int
	AddGameObject(o GameObject)
	GetGameObjects() map[string]GameObject
	GetSpaceship1() *Spaceship
	GetSpaceship2() *Spaceship
	GetUpdateCounter() int
	Layout(outsideWidth, outsideHeight int) (int, int)
	Update(screen *ebiten.Image) error
}

func (g *SpaceshipGame) GetMaxX() int {
	return g.MaxX
}

func (g *SpaceshipGame) GetMaxY() int {
	return g.MaxY
}

func (g *SpaceshipGame) GetUpdateCounter() int {
	return g.UpdateCounter
}

func (g *SpaceshipGame) AddGameObject(o GameObject) {
	g.GameObjects[o.GetID()] = o
}

func (g *SpaceshipGame) GetGameObjects() map[string]GameObject {
	return g.GameObjects
}

func (g *SpaceshipGame) GetSpaceship1() *Spaceship {
	return g.Spaceship1
}

func (g *SpaceshipGame) GetSpaceship2() *Spaceship {
	return g.Spaceship2
}

type SpaceshipGame struct {
	BackgroundImage *ebiten.Image
	Spaceship1      *Spaceship
	Spaceship2      *Spaceship
	GameObjects     map[string]GameObject
	DebugScreen     *DebugScreen
	MaxX            int
	MaxY            int
	DebugPrint      bool

	BackgroundSound *audio.Player

	UpdateCounter int

	Pause bool

	SoundOn bool

	KilledEnemies int

	State GameState

	GamepadIDs map[int]struct{}
}

func NewGame() (Game, error) {

	debugScreen, err := NewDebugScreen(500, screenHeight)
	if err != nil {
		return nil, err
	}

	spaceship1, err := NewSpaceship(NewPos(100, 300), nil, ps3GamepadControlMap)
	if err != nil {
		return nil, err
	}

	spaceship2, err := NewSpaceship(NewPos(800, 300), keyboardControlMap, nil)
	if err != nil {
		return nil, err
	}

	gameObjects := map[string]GameObject{}

	g := &SpaceshipGame{
		GameObjects:   gameObjects,
		DebugScreen:   debugScreen,
		Spaceship1:    spaceship1,
		Spaceship2:    spaceship2,
		UpdateCounter: 0,
		MaxX:          screenWidth,
		MaxY:          screenHeight,
		DebugPrint:    false,
		Pause:         false,
		SoundOn:       false,
		State:         Running,
		GamepadIDs:    map[int]struct{}{},
	}

	return g, nil
}

func (g *SpaceshipGame) Reset() {
	g.GameObjects = map[string]GameObject{}
	g.Spaceship1, _ = NewSpaceship(NewPos(100, 300), nil, gamepadControlMap)
	g.Spaceship2, _ = NewSpaceship(NewPos(100, 800), keyboardControlMap, nil)
	g.UpdateCounter = 0
	g.Pause = false
	g.State = Running
	g.KilledEnemies = 0
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

	//TODO: refactoring!
	g.UpdateCounter++
	if g.UpdateCounter%100 == 0 {
		putNewEnemies(g)
	}
	if g.UpdateCounter%40 == 0 {
		putStars(g)
	}
	if g.UpdateCounter%500 == 0 {
		putNewAmmos(g, 1)
	}
	if g.UpdateCounter%10000 == 0 {
		g.UpdateCounter = 0
	}

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
	if g.Spaceship1.Health < 0 {
		g.State = GameOver
	}
}

func bulletSkyObjectCollisionDetection(g *SpaceshipGame) {

	for k, o := range g.GameObjects {

		if o.GetType() == Weapon {
			for _, x := range g.GameObjects {
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
						g.KilledEnemies += 1
						delete(g.GameObjects, k)
					}
				}
			}
		}
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

	}
}

func drawGameOverScreen(g *SpaceshipGame, screen *ebiten.Image) {
	t := fmt.Sprintf(
		"GAME OVER \n"+
			"You Killed %d Enemies \n"+
			"Press R for New Game \n"+
			"Press Q for Quit",
		g.KilledEnemies)
	text.Draw(screen, t, engine.MplusBigFont, 700, 300, color.White)
}

func putStars(g *SpaceshipGame) {
	newSkyObjects := CreateStarAtRandomPosition(0, 0, screenWidth, screenHeight, 10)

	for _, nO := range newSkyObjects {
		g.GameObjects[nO.GetID()] = nO
	}
}

func putNewEnemies(g *SpaceshipGame) {
	newSkyObjects := CreateSkyObjectAtRandomPosition(
		(screenWidth/3)*2, 0, screenWidth, screenHeight, 1)

	for _, nO := range newSkyObjects {
		g.GameObjects[nO.GetID()] = nO
	}
}

func putNewAmmos(g *SpaceshipGame, count int) {
	newAmmos := CreateAmmoAtRandomPosition(
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

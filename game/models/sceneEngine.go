package models

import "github.com/polpettone/gaming/spaceships/engine"

func SpaceshipCollisionDetection(s *Spaceship, gameObjects map[string]GameObject) {

	if s == nil {
		return
	}

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

func DeleteObjectsOutOfView(g Scene) {
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

func BulletSkyObjectCollisionDetection(gameObjects map[string]GameObject) {

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

package models

import (
	"math/rand"
	"time"
)

func chanceControl(s *Spaceship, maxX, maxY int) {

	s.ShootBullet = false
	s.Acceleration = false
	s.XAxisForce = 1

	shootByChance(s, 50)
	accelerationByChance(s, 100)
	changeDirectionByChance(s, 50)
	changeXAxisForceByChance(s, 30, maxX, maxY)
	changeYAxisForceByChance(s, 50, maxX, maxY)
}

func changeXAxisForceByChance(s *Spaceship, chance int, maxX, maxY int) {

	if s.XAxisForce > 4 || s.XAxisForce < -4 {
		s.XAxisForce = 0
	}

	if s.Pos.X > maxX-20 {
		s.XAxisForce -= 5
		return
	}

	if s.Pos.X < 20 {
		s.XAxisForce += 5
		return
	}

	if chance > calcChance() {
		s.XAxisForce += 1
	} else {
		s.XAxisForce -= 1
	}
}

func changeYAxisForceByChance(s *Spaceship, chance int, maxX, maxY int) {

	if s.YAxisForce > 4 || s.YAxisForce < -4 {
		s.YAxisForce = 0
	}

	if s.Pos.Y > maxY-20 {
		s.YAxisForce -= 5
	}

	if s.Pos.Y < 20 {
		s.YAxisForce += 5
	}

	if chance > calcChance() {
		s.YAxisForce += 1
	} else {
		s.YAxisForce -= 1
	}
}

func changeDirectionByChance(s *Spaceship, chance int) {
	if chance > calcChance() {
		s.MoveDirection = s.MoveDirection * -1
	}
}

func shootByChance(s *Spaceship, chance int) {
	if chance > calcChance() {
		s.ShootBullet = true
	}
}

func accelerationByChance(s *Spaceship, chance int) {
	if chance > calcChance() {
		s.Acceleration = true
	}
}

func calcChance() int {
	max := 100
	min := 1
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

package models

type SimpleImpetusUnit struct {
}

func (i *SimpleImpetusUnit) UpdatePosition(s *Spaceship, maxX, maxY int) {
	if s.Acceleration {
		if s.Pos.X < maxX-spaceshipWallTolerance && s.XAxisForce > 0 {
			s.Pos.X += s.XAxisForce
		}
		if s.Pos.X > spaceshipWallTolerance && s.XAxisForce < 0 {
			s.Pos.X += s.XAxisForce
		}
		if s.Pos.Y < maxY-spaceshipWallTolerance && s.YAxisForce > 0 {
			s.Pos.Y += s.YAxisForce
		}
		if s.Pos.Y > spaceshipWallTolerance && s.YAxisForce < 0 {
			s.Pos.Y += s.YAxisForce
		}
	}

	if s.Pos.X-spaceshipWallTolerance < 0 ||
		s.Pos.X+spaceshipWallTolerance > maxX {
		s.XAxisForce = 0
	}

	if s.Pos.Y-spaceshipWallTolerance < 0 ||
		s.Pos.Y+spaceshipWallTolerance > maxY {
		s.YAxisForce = 0
	}
}

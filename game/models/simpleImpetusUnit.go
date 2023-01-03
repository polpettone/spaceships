package models

type SimpleImpetusUnit struct {
}

func (i *SimpleImpetusUnit) UpdatePosition(s *Spaceship, g Scene) {
	if s.Acceleration {
		if s.Pos.X < g.GetMaxX()-spaceshipWallTolerance && s.XAxisForce > 0 {
			s.Pos.X += s.XAxisForce
		}
		if s.Pos.X > spaceshipWallTolerance && s.XAxisForce < 0 {
			s.Pos.X += s.XAxisForce
		}
		if s.Pos.Y < g.GetMaxY()-spaceshipWallTolerance && s.YAxisForce > 0 {
			s.Pos.Y += s.YAxisForce
		}
		if s.Pos.Y > spaceshipWallTolerance && s.YAxisForce < 0 {
			s.Pos.Y += s.YAxisForce
		}
	}

	if s.Pos.X-spaceshipWallTolerance < 0 || s.Pos.X+spaceshipWallTolerance > g.GetMaxX() {
		s.XAxisForce = 0
	}

	if s.Pos.Y-spaceshipWallTolerance < 0 || s.Pos.Y+spaceshipWallTolerance > g.GetMaxY() {
		s.YAxisForce = 0
	}
}

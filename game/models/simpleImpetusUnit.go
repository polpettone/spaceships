package models

type SimpleImpetusUnit struct {
}

func (i *SimpleImpetusUnit) UpdatePosition(
	s *Spaceship,
	maxX,
	maxY,
	wallTolerance int) {

	if s.Acceleration {
		if s.Pos.X < maxX-wallTolerance && s.XAxisForce > 0 {
			s.Pos.X += s.XAxisForce
		}
		if s.Pos.X > wallTolerance && s.XAxisForce < 0 {
			s.Pos.X += s.XAxisForce
		}
		if s.Pos.Y < maxY-wallTolerance && s.YAxisForce > 0 {
			s.Pos.Y += s.YAxisForce
		}
		if s.Pos.Y > wallTolerance && s.YAxisForce < 0 {
			s.Pos.Y += s.YAxisForce
		}
	}

	if s.Pos.X-wallTolerance < 0 ||
		s.Pos.X+wallTolerance > maxX {
		s.XAxisForce = 0
	}

	if s.Pos.Y-wallTolerance < 0 ||
		s.Pos.Y+wallTolerance > maxY {
		s.YAxisForce = 0
	}
}

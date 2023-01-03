package models

type SimpleImpetusUnit struct {
}

func (i *SimpleImpetusUnit) UpdatePosition(
	acceleration bool,
	pos Pos,
	xAxisForce,
	yAxisForce,
	maxX,
	maxY,
	wallTolerance int) (Pos, int, int) {

	if acceleration {
		if pos.X < maxX-wallTolerance && xAxisForce > 0 {
			pos.X += xAxisForce
		}
		if pos.X > wallTolerance && xAxisForce < 0 {
			pos.X += xAxisForce
		}
		if pos.Y < maxY-wallTolerance && yAxisForce > 0 {
			pos.Y += yAxisForce
		}
		if pos.Y > wallTolerance && yAxisForce < 0 {
			pos.Y += yAxisForce
		}
	}

	if pos.X-wallTolerance < 0 ||
		pos.X+wallTolerance > maxX {
		xAxisForce = 0
	}

	if pos.Y-wallTolerance < 0 ||
		pos.Y+wallTolerance > maxY {
		yAxisForce = 0
	}
	return pos, xAxisForce, yAxisForce
}

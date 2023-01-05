package models

type SimpleImpetusUnit struct {
}

func (i *SimpleImpetusUnit) Accelerate(
	pos Pos,
	xAxisForce,
	yAxisForce int) (Pos, int, int) {

	pos.X += xAxisForce
	pos.Y += yAxisForce

	return pos, xAxisForce, yAxisForce
}

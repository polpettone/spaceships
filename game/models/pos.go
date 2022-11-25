package models

import "fmt"

type Pos struct {
	X int
	Y int
}

func NewPos(x, y int) Pos {
	return Pos{X: x, Y: y}
}

func (p Pos) Print() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

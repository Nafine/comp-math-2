package model

import "comp-math-2/internal/derivate"

type Function func(float64) float64
type Solution struct {
	X          float64
	Y          float64
	Iterations int
}

type NonlinearEquation struct {
	F   Function
	Eps float64
	A   float64
	B   float64
}

func (eq NonlinearEquation) RootExists() bool {
	return eq.F(eq.A)*eq.F(eq.B) < 0 && derivate.DerivAt(eq.F, eq.A)*derivate.DerivAt(eq.F, eq.B) > 0
}

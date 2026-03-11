package algo

import (
	"comp-math-2/internal/derivate"
	"comp-math-2/internal/numeric"
	"math"
)

func SolveNewton(eq numeric.NonlinearEquation) (numeric.Solution, error) {
	f := eq.F
	deriv := derivate.Derivate(f)
	a := eq.A
	b := eq.B
	eps := eq.Eps
	iterations := 0

	x := (b + a) / 2.0

	xPrev := x

	for {
		iterations++
		x = x - f(x)/deriv(x)

		if math.Abs(x-xPrev) <= eps {
			break
		}

		xPrev = x
	}

	return numeric.Solution{
		X:          x,
		Y:          f(x),
		Iterations: iterations,
	}, nil
}

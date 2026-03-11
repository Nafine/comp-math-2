package algo

import (
	"comp-math-2/internal/numeric"
	"math"
)

func SolveSplitting(eq numeric.NonlinearEquation) (numeric.Solution, error) {
	f := eq.F
	a := eq.A
	b := eq.B
	eps := eq.Eps
	iterations := 0

	prevX := (b + a) / 2.0

	x := prevX + eps

	for {
		iterations++
		temp := x
		x = x - (x-prevX)/(f(x)-f(prevX))*f(x)

		if math.Abs(x-prevX) <= eps {
			break
		}

		prevX = temp
	}

	return numeric.Solution{
		X:          x,
		Y:          f(x),
		Iterations: iterations,
		Method:     "Half Splitting",
	}, nil
}

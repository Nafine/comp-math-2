package algo

import (
	"comp-math-2/internal/numeric"
	"fmt"
	"math"
)

func SolveChord(eq numeric.NonlinearEquation) (numeric.Solution, error) {
	f := eq.F
	a := eq.A
	b := eq.B
	eps := eq.Eps

	x := a - (b-a)*f(a)/(f(b)-f(a))

	lastX := x

	for i := 1; i <= 10_000; i++ {
		if f(a)*f(x) < 0 {
			b = x
		} else {
			a = x
		}

		x = a - (b-a)*f(a)/(f(b)-f(a))

		if math.Abs(f(x)) <= eps && math.Abs(x-lastX) <= eps {
			return numeric.Solution{
				X:          x,
				Iterations: i,
			}, nil
		}

		lastX = x
	}

	return numeric.Solution{}, fmt.Errorf("maximum number of iterations reached")
}

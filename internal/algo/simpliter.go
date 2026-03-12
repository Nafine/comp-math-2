package algo

import (
	"comp-math-2/internal/derivate"
	"comp-math-2/internal/numeric"
	"fmt"
	"math"
)

func SolveSimpleIteration(eq numeric.NonlinearEquation) (numeric.Solution, error) {
	steps := 1000
	maxIters := 50_000

	f := eq.F
	x := (eq.A + eq.B) / 2.0

	maxDeriv := max(math.Abs(derivate.DerivAt(f, eq.A)), math.Abs(derivate.DerivAt(f, eq.B)))
	lbd := 1 / maxDeriv

	if derivate.DerivAt(f, x) > 0 {
		lbd *= -1
	}

	phi := func(x float64) float64 {
		return x + lbd*f(x)
	}

	for i := 1; i <= steps; i++ {
		if math.Abs(derivate.DerivAt(phi, eq.A+(eq.B-eq.A)/float64(steps)*float64(i))) >= 1 {
			return numeric.Solution{},
				fmt.Errorf("convergence condition failed: |phi'(x)| < 1 at x = %f", x)
		}
	}

	for i := 1; i <= maxIters; i++ {
		xPrev := x
		x = phi(x)

		if math.Abs(x-xPrev) <= eq.Eps {
			return numeric.Solution{
				X:          x,
				Iterations: i,
			}, nil
		}
	}

	return numeric.Solution{}, fmt.Errorf("maximum number of iterations reached")
}

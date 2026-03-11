package algo

import (
	"comp-math-2/internal/derivate"
	"comp-math-2/internal/numeric"
	"fmt"
	"math"
)

func SolveSimpleIteration(eq numeric.NonlinearEquation) (numeric.Solution, error) {
	steps := 100
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
				fmt.Errorf("не выполнено условие сходимости метода |phi'(x)| < 1 при x = %f", x)
		}
	}

	iterations := 0

	for {
		iterations++

		if iterations == maxIters {
			return numeric.Solution{}, fmt.Errorf("достигнуто %d итераций и ответ вы уже не получите", iterations)
		}

		xPrev := x
		x = phi(x)

		if math.Abs(x-xPrev) <= eq.Eps {
			break
		}
	}

	return numeric.Solution{
		X:          x,
		Y:          f(x),
		Iterations: iterations,
	}, nil
}

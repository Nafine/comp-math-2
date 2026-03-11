package algo

import (
	"comp-math-2/internal/derivate"
	"comp-math-2/internal/numeric"
	"fmt"
	"math"
)

func SolveSystem(eq numeric.NonlinearSystem) (numeric.Solution, error) {
	coords := numeric.Coordinates{X: eq.StartCoordinates.X, Y: eq.StartCoordinates.Y}
	maxIter := 1000

	for iter := 0; iter < maxIter; iter++ {
		f1 := eq.F1(coords)
		f2 := eq.F2(coords)

		// Проверка на сходимость
		if math.Abs(f1) < eq.Eps && math.Abs(f2) < eq.Eps {
			return numeric.Solution{
				X:          coords.X,
				Y:          coords.Y,
				Iterations: iter,
				Method:     "Simple Iteration",
			}, nil
		}

		J11 := derivate.DerivXAt(eq.F1, coords)
		J12 := derivate.DerivYAt(eq.F1, coords)
		J21 := derivate.DerivXAt(eq.F2, coords)
		J22 := derivate.DerivYAt(eq.F2, coords)

		det := J11*J22 - J12*J21

		if math.Abs(det) < 1e-12 {
			return numeric.Solution{}, fmt.Errorf("якобиан вырожден в точке (%f, %f)", coords.X, coords.Y)
		}

		deltaX := -(f1*J22 - f2*J12) / det
		deltaY := -(J11*f2 - J21*f1) / det

		coords.X += deltaX
		coords.Y += deltaY
	}

	return numeric.Solution{}, fmt.Errorf("достигнуто максимальное число итераций")
}

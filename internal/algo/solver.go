package algo

import (
	"comp-math-2/internal/derivate"
	"comp-math-2/internal/numeric"
	"errors"
)

var methods = []func(equation numeric.NonlinearEquation) (numeric.Solution, error){
	SolveChord,
	SolveSplitting,
	SolveSimpleIteration,
}

func SolveAllSingle(eq numeric.NonlinearEquation) ([]numeric.Solution, error) {
	solutions := make([]numeric.Solution, 0)

	if !(rootExists(eq)) {
		return nil, errors.New("function has no root in the given interval")
	}

	for _, method := range methods {
		solution, err := method(eq)

		if err != nil {
			return nil, err
		}
		solutions = append(solutions, solution)
	}

	return solutions, nil
}

func rootExists(eq numeric.NonlinearEquation) bool {
	return eq.F(eq.A)*eq.F(eq.B) < 0 && derivate.DerivAt(eq.F, eq.A)*derivate.DerivAt(eq.F, eq.B) > 0
}

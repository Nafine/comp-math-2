package algo

import (
	"comp-math-2/internal/derivate"
	"comp-math-2/internal/numeric"
	"errors"
)

var methods = map[string]func(equation numeric.NonlinearEquation) (numeric.Solution, error){
	"chord":     SolveChord,
	"secant":    SolveSecant,
	"iteration": SolveSimpleIteration,
}

func SolveSingle(method string, eq numeric.NonlinearEquation) (numeric.Solution, error) {
	if !rootExists(eq) {
		return numeric.Solution{}, errors.New("no roots exists on the given interval")
	}

	return methods[method](eq)
}

func rootExists(eq numeric.NonlinearEquation) bool {
	return eq.F(eq.A)*eq.F(eq.B) < 0 && derivate.DerivAt(eq.F, eq.A)*derivate.DerivAt(eq.F, eq.B) > 0
}

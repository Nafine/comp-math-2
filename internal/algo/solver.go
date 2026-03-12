package algo

import (
	"comp-math-2/internal/numeric"
	"errors"
)

var methods = map[string]func(equation numeric.NonlinearEquation) (numeric.Solution, error){
	"chord":     SolveChord,
	"secant":    SolveSecant,
	"iteration": SolveSimpleIteration,
}

func SolveSingle(method string, eq numeric.NonlinearEquation) (numeric.Solution, error) {
	if eq.A >= eq.B {
		return numeric.Solution{}, errors.New("a must be higher than b")
	}

	if eq.Eps <= 0 {
		return numeric.Solution{}, errors.New("eps must be greater than zero")
	}

	if !rootExists(eq) {
		return numeric.Solution{}, errors.New("no roots exists on the given interval")
	}

	if !monotonic(eq) {
		return numeric.Solution{}, errors.New("function is not monotonic on interval: multiple roots possible")
	}

	return methods[method](eq)
}

func monotonic(eq numeric.NonlinearEquation) bool {
	steps := 100
	prev := eq.F(eq.A)
	isIncreasing := eq.F(eq.A+(eq.B-eq.A)/float64(steps)) > prev

	for i := 1; i <= steps; i++ {
		x := eq.A + (eq.B-eq.A)/float64(steps)*float64(i)
		cur := eq.F(x)

		if isIncreasing && cur < prev {
			return false
		}

		if !isIncreasing && cur > prev {
			return false
		}

		prev = cur
	}

	return true
}

func rootExists(eq numeric.NonlinearEquation) bool {
	return eq.F(eq.A)*eq.F(eq.B) < 0
}

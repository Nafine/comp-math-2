package derivate

import (
	"comp-math-2/internal/numeric"
)

const h = 1e-5

func DerivAt(f func(float64) float64, x float64) float64 {
	return (f(x+h) - f(x-h)) / (2 * h)
}

func DerivXAt(f func(numeric.Coordinates) float64, coords numeric.Coordinates) float64 {
	return (f(numeric.Coordinates{X: coords.X + h, Y: coords.Y}) -
		f(numeric.Coordinates{X: coords.X - h, Y: coords.Y})) / (2 * h)
}

func DerivYAt(f func(numeric.Coordinates) float64, coords numeric.Coordinates) float64 {
	return (f(numeric.Coordinates{X: coords.X, Y: coords.Y + h}) -
		f(numeric.Coordinates{X: coords.X, Y: coords.Y - h})) / (2 * h)
}

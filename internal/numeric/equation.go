package numeric

import "math"

type Function struct {
	Name string
	Fn   func(float64) float64
}
type System struct {
	F1 func(Coordinates) float64
	F2 func(Coordinates) float64
}

var functions = []func(x float64) float64{
	func(x float64) float64 {
		return x*x*x + 2.84*x*x - 5.606*x - 14.766
	},
	func(x float64) float64 {
		return x*x*x - 1.89*x*x - 2*x + 1.76
	},
	func(x float64) float64 {
		return math.Sin(3*x) - 0.5
	},
}

var systems = []System{
	{
		F1: func(coords Coordinates) float64 {
			return math.Cos(coords.X-1) + coords.Y - 0.5
		},
		F2: func(coords Coordinates) float64 {
			return coords.X - math.Cos(coords.Y) - 3
		},
	},
	{
		F1: func(coords Coordinates) float64 {
			return math.Sin(coords.X+coords.Y) - 1.5*coords.X + 0.1
		},
		F2: func(coords Coordinates) float64 {
			return coords.X*coords.X + 2*coords.Y*coords.Y - 1
		},
	},
}

func GetFunction(index int) func(float64) float64 {
	return functions[index]
}

func GetSystem(index int) System {
	return systems[index]
}

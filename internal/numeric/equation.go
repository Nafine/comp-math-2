package numeric

import "math"

var singleEquations = map[string]func(float64) float64{
	"x^3 + 2.84x^2 - 5.606x - 14.766": func(x float64) float64 {
		return math.Pow(x, 3) + 2.84*math.Pow(x, 2) - 5.606*x - 14.766
	},
	"x^3 - 1.89x^2 - 2x + 1.76": func(x float64) float64 {
		return math.Pow(x, 3) - 1.89*math.Pow(x, 2) - 2*x + 1.76
	},
	"sin(3x) - 0.5": func(x float64) float64 {
		return math.Sin(3*x) - 0.5
	},
}

type system struct {
	f1 func(Coordinates) float64
	f2 func(Coordinates) float64
}

var systems = map[string]system{
	"cos(x-1) + y = 0.5\nx - cos(y) = 3": {
		f1: func(coordinates Coordinates) float64 {
			return math.Cos(coordinates.X) + coordinates.Y - 0.5
		},
		f2: func(coordinates Coordinates) float64 {
			return coordinates.X - math.Cos(coordinates.Y) - 3
		},
	},
	"sin(x+y) = 1.5x - 0.1\nx^2+2y^2=1": {
		f1: func(coordinates Coordinates) float64 {
			return math.Sin(coordinates.X+coordinates.Y) - 1.5*coordinates.X + 0.1
		},
		f2: func(coordinates Coordinates) float64 {
			return math.Pow(coordinates.X, 2) + 2*math.Pow(coordinates.Y, 2) - 1
		},
	},
}

func GetSingleEquations() []string {
	return []string{
		"x^3 + 2.84x^2 - 5.606x - 14.766",
		"x^3 - 1.89x^2 - 2x + 1.76",
		"sin(3x) - 0.5",
	}
}

func GetSystems() []string {
	return []string{
		"cos(x-1) + y = 0.5\nx - cos(y) = 3",
		"sin(x+y) = 1.5x - 0.1\nx^2+2y^2=1",
	}
}

func GetEquation(eq string) func(float64) float64 {
	return singleEquations[eq]
}

func GetSystem(system string) (func(coordinates Coordinates) float64, func(coordinates Coordinates) float64) {
	return systems[system].f1, systems[system].f2
}

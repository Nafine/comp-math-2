package numeric

import "math"

type Function struct {
	Name string
	Fn   func(float64) float64
}
type System struct {
	Name string
	F1   func(Coordinates) float64
	F2   func(Coordinates) float64
}

var functions = []Function{
	{
		Name: "x³ + 2.84x² - 5.606x - 14.766",
		Fn: func(x float64) float64 {
			return x*x*x + 2.84*x*x - 5.606*x - 14.766
		},
	},
	{
		Name: "x³ - 1.89x² - 2x + 1.76",
		Fn: func(x float64) float64 {
			return x*x*x - 1.89*x*x - 2*x + 1.76
		},
	},
	{
		Name: "sin(3x) - 0.5",
		Fn: func(x float64) float64 {
			return math.Sin(3*x) - 0.5
		},
	},
}

var systems = []System{
	{
		Name: "cos(x-1) + y = 0.5 & x - cos(y) = 3",
		F1: func(coords Coordinates) float64 {
			return math.Cos(coords.X-1) + coords.Y - 0.5
		},
		F2: func(coords Coordinates) float64 {
			return coords.X - math.Cos(coords.Y) - 3
		},
	},
	{
		Name: "sin(x+y) = 1.5x - 0.1 & x²+2y²=1",
		F1: func(coords Coordinates) float64 {
			return math.Sin(coords.X+coords.Y) - 1.5*coords.X + 0.1
		},
		F2: func(coords Coordinates) float64 {
			return coords.X*coords.X + 2*coords.Y*coords.Y - 1
		},
	},
}

func GetFunctionNames() []string {
	names := make([]string, len(functions))
	for i, eq := range functions {
		names[i] = eq.Name
	}
	return names
}

func GetSystemNames() []string {
	names := make([]string, len(systems))
	for i, s := range systems {
		names[i] = s.Name
	}
	return names
}

func GetFunction(index int) Function {
	if index < 0 || index >= len(functions) {
		return Function{}
	}
	return functions[index]
}

func GetSystem(index int) System {
	if index < 0 || index >= len(systems) {
		return System{}
	}
	return systems[index]
}

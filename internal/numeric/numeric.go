package numeric

type Solution struct {
	X          float64
	Iterations int
}

type SystemSolution struct {
	X          float64
	Y          float64
	Iterations int
}

type NonlinearEquation struct {
	F   func(float64) float64
	Eps float64
	A   float64
	B   float64
}

type NonlinearSystem struct {
	F1               func(Coordinates) float64
	F2               func(Coordinates) float64
	StartCoordinates Coordinates
	Eps              float64
}

type Coordinates struct {
	X, Y float64
}

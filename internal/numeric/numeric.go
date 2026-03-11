package numeric

import "fmt"

type Solution struct {
	X          float64
	Y          float64
	Iterations int
	Method     string
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

func (s Solution) String() string {
	return fmt.Sprintf("Method: %s\nX: %f, Y: %f\nIterations: %d\n", s.Method, s.X, s.Y, s.Iterations)
}

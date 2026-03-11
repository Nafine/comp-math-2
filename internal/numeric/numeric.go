package numeric

import "fmt"

type Solution struct {
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
	F1  func(Coordinates) float64
	F2  func(Coordinates) float64
	X0  float64
	Y0  float64
	Eps float64
}

type Coordinates struct {
	X, Y float64
}

func (s Solution) String() string {
	return fmt.Sprintf("X: %f, Y: %f\nIterations: %d\n", s.X, s.Y, s.Iterations)
}

//func (eq NonlinearEquation) RootExists() bool {
//	return eq.F(eq.A)*eq.F(eq.B) < 0 && derivate.DerivAt(eq.F, eq.A)*derivate.DerivAt(eq.F, eq.B) > 0
//}

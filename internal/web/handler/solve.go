package handler

import (
	"comp-math-2/internal/algo"
	"comp-math-2/internal/numeric"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SolveRequest struct {
	Type       string  `json:"type" binding:"required"`
	EquationId int     `json:"equationId"`
	Method     string  `json:"method" binding:"required"`
	Tolerance  float64 `json:"tolerance" binding:"required"`

	A *float64 `json:"a,omitempty"`
	B *float64 `json:"b,omitempty"`

	X0 *float64 `json:"x0,omitempty"`
	Y0 *float64 `json:"y0,omitempty"`
}

type SolveResponse struct {
	X              float64  `json:"x"`
	Y              *float64 `json:"y"`
	IterationCount int      `json:"iterationCount"`
}

func Solve() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SolveRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
			return
		}

		switch req.Type {
		case "single":
			if req.A == nil || req.B == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Fields 'a' and 'b' are required for single mode"})
				return
			}

			eq := numeric.NonlinearEquation{
				F:   numeric.GetFunction(req.EquationId),
				A:   *req.A,
				B:   *req.B,
				Eps: req.Tolerance,
			}

			solution, err := algo.SolveSingle(req.Method, eq)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, SolveResponse{
				X:              solution.X,
				IterationCount: solution.Iterations,
			})

		case "system":
			if req.X0 == nil || req.Y0 == nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Fields 'x0' and 'y0' are required for system mode"})
				return
			}

			system := numeric.GetSystem(req.EquationId)

			eq := numeric.NonlinearSystem{
				F1: system.F1,
				F2: system.F2,
				StartCoordinates: numeric.Coordinates{
					X: *req.X0,
					Y: *req.Y0,
				},
				Eps: req.Tolerance,
			}

			solution, err := algo.SolveSystem(eq)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, SolveResponse{
				X:              solution.X,
				Y:              &solution.Y,
				IterationCount: solution.Iterations,
			})

		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown type: " + req.Type})
		}
	}
}

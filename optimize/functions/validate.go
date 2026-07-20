package functions

import (
	"testing"
)

type function interface {
	Func(x []float64) float64
}

type gradient interface {
	Grad(grad, x []float64) []float64
}

type minimumer interface {
	function

	Minima() []Minimum
}

type Minimum struct {
	X []float64

	F float64

	Global bool
}

type funcTest struct {
	X []float64

	F float64

	Gradient []float64
}

const (
	defaultTol       = 1e-12
	defaultGradTol   = 1e-9
	defaultFDGradTol = 1e-5
)

func testFunction(f function, ftests []funcTest, t *testing.T) { _ = "STUB: not implemented"; return }

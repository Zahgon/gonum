package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dlanster interface {
	Dlanst(norm lapack.MatrixNorm, n int, d, e []float64) float64
	Dlanger
}

func DlanstTest(t *testing.T, impl Dlanster) { _ = "STUB: not implemented"; return }

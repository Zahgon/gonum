package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/lapack"
)

type Dlanger interface {
	Dlange(norm lapack.MatrixNorm, m, n int, a []float64, lda int, work []float64) float64
}

func DlangeTest(t *testing.T, impl Dlanger) { _ = "STUB: not implemented"; return }

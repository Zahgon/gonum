package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dgelser interface {
	Dgels(trans blas.Transpose, m, n, nrhs int, a []float64, lda int, b []float64, ldb int, work []float64, lwork int) bool
}

func DgelsTest(t *testing.T, impl Dgelser) { _ = "STUB: not implemented"; return }

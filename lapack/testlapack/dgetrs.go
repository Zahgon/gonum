package testlapack

import (
	"testing"

	"gonum.org/v1/gonum/blas"
)

type Dgetrser interface {
	Dgetrfer
	Dgetrs(trans blas.Transpose, n, nrhs int, a []float64, lda int, ipiv []int, b []float64, ldb int)
}

func DgetrsTest(t *testing.T, impl Dgetrser) { _ = "STUB: not implemented"; return }
